package schema

import (
	"fmt"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
)

type UpdateDef struct {
	Name    string
	Table   string
	Columns ColumnMap
	Where   []where_predicate
}

func (q *UpdateDef) Type() QueryType {
	return Update
}

func (q *UpdateDef) Prepare(db *Database) (*ExecutableQuery, error) {
	// Create a preparedStatement value and initialize its fields
	tbl, ok := db.Tables[q.Table]
	if !ok {
		return nil, fmt.Errorf("table %s not found", q.Table)
	}

	qry := ExecutableQuery{
		Statement:  "",
		Args:       make(map[int]arg),
		UserInputs: make([]*requiredInput, 0),
	}

	// Create a statementInput value for each column and add it to the preparedStatement's inputs slice
	statement := UpdateBuilder(db.SchemaName() + "." + q.Table)
	i := 1
	for name, val := range q.Columns {
		statement.Column(name)
		fillable := false
		if val == "" {
			fillable = true

			qry.UserInputs = append(qry.UserInputs, &requiredInput{
				Column: name,
				Type:   tbl.Columns[name].Type.String(),
			})
		}

		pgType, ok := Types[tbl.Columns[name].Type]
		if !ok {
			return nil, fmt.Errorf("unsupported type: %s", tbl.Columns[name].Type)
		}

		qry.Args[i] = arg{
			Column:   name,
			Default:  val,
			Type:     pgType.String(),
			Fillable: fillable,
		}
		i++
	}

	// Now create the WHERE clauses
	for _, where := range q.Where {
		statement.Where(&where)
		fillable := false
		if where.Default == "" {
			fillable = true

			qry.UserInputs = append(qry.UserInputs, &requiredInput{
				Column: where.Column,
				Type:   tbl.Columns[where.Column].Type.String(),
			})
		}

		pgType, ok := Types[tbl.Columns[where.Column].Type]
		if !ok {
			return nil, fmt.Errorf("unsupported type: %s", tbl.Columns[where.Column].Type)
		}

		qry.Args[i] = arg{
			Column:   where.Column,
			Default:  where.Default,
			Type:     pgType.String(),
			Fillable: fillable,
		}
		i++

	}

	// Now create the SQL statement
	stmt, err := statement.ToSQL()
	if err != nil {
		return nil, err
	}
	qry.Statement = stmt

	// Return a pointer to the preparedStatement value
	return &qry, nil
}

type updateBuilder struct {
	stmt   *goqu.UpdateDataset
	sets   goqu.Record
	wheres []goqu.Expression
}

func UpdateBuilder(table string) *updateBuilder {
	return &updateBuilder{
		stmt: goqu.Dialect("postgres").Update(table).Prepared(true),
		sets: make(goqu.Record),
	}
}

func (b *updateBuilder) Column(column string) *updateBuilder {
	var i interface{}
	b.sets[column] = i
	return b
}

func (b *updateBuilder) Where(where *where_predicate) *updateBuilder {
	gexp := WhereConversions.StringToPredicate(where.Operator, where.Column)
	if gexp == nil {
		return b
	}
	b.wheres = append(b.wheres, gexp) // passing empty interface doesn't work with goqu here for some reason so I pass empty string
	return b
}

func (b *updateBuilder) ToSQL() (string, error) {
	b.stmt = b.stmt.Set(b.sets)
	b.stmt = b.stmt.Where(b.wheres...)
	stmt, _, err := b.stmt.ToSQL()
	return stmt, err
}