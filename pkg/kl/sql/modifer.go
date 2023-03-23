package sql

type Modifier string

const ModifierPrefix = `@`
const BindParameterPrefix = `$`

const (
	ModifierCaller      Modifier = `caller`
	ModifierBlockHeight Modifier = `block.height`
)

var Modifiers = map[Modifier]bool{
	ModifierCaller:      true,
	ModifierBlockHeight: true,
}
