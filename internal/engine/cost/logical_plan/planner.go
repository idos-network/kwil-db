package logical_plan

import "github.com/kwilteam/kwil-db/parse/sql/tree"

type LogicalPlanner interface {
	CreatePlan(node tree.AstNode) LogicalPlan
	CreateExpr(expr logical_plan.LogicalExpr, input logical_plan.LogicalPlan) VirtualExpr
}
