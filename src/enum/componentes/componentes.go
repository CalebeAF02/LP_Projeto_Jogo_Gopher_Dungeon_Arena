package componentes

type ComponenteTipo int

const (
	CORPO ComponenteTipo = iota
)

func (t ComponenteTipo) String() string {
	switch t {
	case CORPO:
		return "CORPO"
	default:
		return "**DESCONHECIDO**"
	}
}
