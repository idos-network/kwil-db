package dto

import (
	"kwil/x/crypto"
	"kwil/x/utils"
)

type Database struct {
	Owner      string      `json:"owner"`
	Name       string      `json:"name"`
	Tables     []*Table    `json:"tables"`
	Roles      []*Role     `json:"roles"`
	SQLQueries []*SQLQuery `json:"sql_queries"`
	Indexes    []*Index    `json:"indexes"`
}

func (d *Database) GetSchemaName() string {

	return "x" + crypto.Sha224Str(utils.JoinBytes([]byte(d.Name), []byte(d.Owner)))
}

func (d *Database) GetQuery(q string) *SQLQuery {
	for _, qry := range d.SQLQueries {
		if qry.Name == q {
			return qry
		}
	}
	return nil
}

func (d *Database) GetTable(t string) *Table {
	for _, tbl := range d.Tables {
		if tbl.Name == t {
			return tbl
		}
	}
	return nil
}

func (d *Database) GetRole(r string) *Role {
	for _, role := range d.Roles {
		if role.Name == r {
			return role
		}
	}
	return nil
}

func (d *Database) GetIndex(i string) *Index {
	for _, idx := range d.Indexes {
		if idx.Name == i {
			return idx
		}
	}
	return nil
}
