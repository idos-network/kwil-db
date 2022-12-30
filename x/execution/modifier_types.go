package execution

type ModifierType int

// Modifiers
const (
	INVALID_MODIFIER ModifierType = iota
	NO_MODIFIER
	CALLER
	END_MODIFIER
)

func (m ModifierType) String() string {
	switch m {
	case CALLER:
		return "caller"
	}
	return "unknown"
}

func (m ModifierType) Int() int {
	return int(m)
}

func (m ModifierType) IsValid() bool {
	return m > INVALID_MODIFIER && m < END_MODIFIER
}
