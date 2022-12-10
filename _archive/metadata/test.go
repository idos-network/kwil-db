package metadata

import (
	"context"
	"ksl/ast"
	"ksl/lift"
	"ksl/postgres"
	"ksl/sqlmigrate"
	"ksl/sqlschema"
)

type dbinfo struct {
	Wallet   string
	Database string
}

type testservice struct {
	databases map[dbinfo]sqlschema.Database
}

func NewTestService() Service {
	return &testservice{
		databases: make(map[dbinfo]sqlschema.Database),
	}
}

func (s *testservice) Plan(ctx context.Context, req SchemaRequest) (Plan, error) {
	ksch := ast.Parse(req.SchemaData, "<schema>")

	if ksch.HasErrors() {
		return Plan{}, ksch.Diagnostics
	}
	target := lift.Sql(ksch, req.Database)

	current, ok := s.databases[dbinfo{Wallet: req.Wallet, Database: req.Database}]
	if !ok {
		current = sqlschema.NewDatabase(req.Database)
	}

	differ := sqlmigrate.NewDiffer(postgres.Backend{})
	changes, err := differ.Diff(current, target)
	if err != nil {
		return Plan{}, err
	}

	planner := postgres.Planner{}
	plan, err := planner.Plan(sqlmigrate.Migration{Before: current, After: target, Changes: changes})
	if err != nil {
		return Plan{}, err
	}

	return convertPlan(plan), nil
}

func (s *testservice) Apply(ctx context.Context, req SchemaRequest) error {
	ksch := ast.Parse(req.SchemaData, "<schema>")

	if ksch.HasErrors() {
		return ksch.Diagnostics
	}

	target := lift.Sql(ksch, req.Database)
	s.databases[dbinfo{Wallet: req.Wallet, Database: req.Database}] = target
	return nil
}

func (s *testservice) GetMetadata(ctx context.Context, req RequestMetadata) (Metadata, error) {
	db, ok := s.databases[dbinfo(req)]
	if !ok {
		return Metadata{}, ErrDatabaseNotFound
	}

	return convertDatabase(db), nil
}