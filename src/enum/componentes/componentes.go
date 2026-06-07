package componentes

type ComponenteTipo int

const (
	CORPO ComponenteTipo = iota
	SUB_TIPO ComponenteTipo = iota
)

func (t ComponenteTipo) String() string {
	switch t {
	case CORPO:
		return "CORPO"
	case SUB_TIPO:
		return "SUB_TIPO"
	default:
		return "**DESCONHECIDO**"
	}
}
