package txsvc

import (
	"fmt"

	txpb "github.com/kwilteam/kwil-db/core/rpc/protobuf/tx/v1"
	"github.com/kwilteam/kwil-db/core/types/transactions"
	engineTypes "github.com/kwilteam/kwil-db/internal/engine/types"
)

func convertSchemaFromEngine(schema *engineTypes.Schema) (*txpb.Schema, error) {
	actions, err := convertActionsFromEngine(schema.Procedures)
	if err != nil {
		return nil, err
	}
	return &txpb.Schema{
		Owner:   schema.Owner,
		Name:    schema.Name,
		Tables:  convertTablesFromEngine(schema.Tables),
		Actions: actions,
	}, nil
}

func convertTablesFromEngine(tables []*engineTypes.Table) []*txpb.Table {
	convTables := make([]*txpb.Table, len(tables))
	for i, table := range tables {
		convTable := &txpb.Table{
			Name:    table.Name,
			Columns: convertColumnsFromEngine(table.Columns),
			Indexes: convertIndexesFromEngine(table.Indexes),
		}
		convTables[i] = convTable
	}

	return convTables
}

func convertColumnsFromEngine(columns []*engineTypes.Column) []*txpb.Column {
	convColumns := make([]*txpb.Column, len(columns))
	for i, column := range columns {
		convColumn := &txpb.Column{
			Name:       column.Name,
			Type:       column.Type.String(),
			Attributes: convertAttributesFromEngine(column.Attributes),
		}
		convColumns[i] = convColumn
	}

	return convColumns
}

func convertAttributesFromEngine(attributes []*engineTypes.Attribute) []*txpb.Attribute {
	convAttributes := make([]*txpb.Attribute, len(attributes))
	for i, attribute := range attributes {
		convAttribute := &txpb.Attribute{
			Type:  attribute.Type.String(),
			Value: fmt.Sprint(attribute.Value),
		}
		convAttributes[i] = convAttribute
	}

	return convAttributes
}

func convertIndexesFromEngine(indexes []*engineTypes.Index) []*txpb.Index {
	convIndexes := make([]*txpb.Index, len(indexes))
	for i, index := range indexes {
		convIndexes[i] = &txpb.Index{
			Name:    index.Name,
			Columns: index.Columns,
			Type:    index.Type.String(),
		}
	}

	return convIndexes
}

func convertActionsFromEngine(actions []*engineTypes.Procedure) ([]*txpb.Action, error) {

	convActions := make([]*txpb.Action, len(actions))
	for i, action := range actions {
		mutability, auxiliaries, err := convertModifiersFromEngine(action.Modifiers)
		if err != nil {
			return nil, err
		}

		convActions[i] = &txpb.Action{
			Name:        action.Name,
			Public:      action.Public,
			Mutability:  mutability,
			Auxiliaries: auxiliaries,
			Inputs:      action.Args,
			Statements:  action.Statements,
		}
	}

	return convActions, nil
}

func convertModifiersFromEngine(mods []engineTypes.Modifier) (mutability string, auxiliaries []string, err error) {
	auxiliaries = make([]string, 0)
	mutability = transactions.MutabilityUpdate.String()
	for _, mod := range mods {
		switch mod {
		case engineTypes.ModifierAuthenticated:
			auxiliaries = append(auxiliaries, transactions.AuxiliaryTypeMustSign.String())
		case engineTypes.ModifierView:
			mutability = transactions.MutabilityView.String()
		case engineTypes.ModifierOwner:
			auxiliaries = append(auxiliaries, transactions.AuxiliaryTypeOwner.String())
		default:
			return "", nil, fmt.Errorf("unknown modifier type: %v", mod)
		}
	}

	return mutability, auxiliaries, nil
}