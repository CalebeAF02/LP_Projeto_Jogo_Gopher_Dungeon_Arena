package entidades

type EntidadeTipo int

const (
	POSICAO EntidadeTipo = iota
	BOT
	JOGADOR
	TIME
	PAREDE
	PORTAL_ENTRADA
	PORTAL_SAIDA
	COMIDA
	SAIDA
)

func (t EntidadeTipo) String() string {
	switch t {
	case POSICAO:
		return "POSICAO"
	case BOT:
		return "BOT"
	case JOGADOR:
		return "JOGADOR"
	case TIME:
		return "TIME"
	case PAREDE:
		return "PAREDE"
	case PORTAL_ENTRADA:
		return "PORTAL_ENTRADA"
	case PORTAL_SAIDA:
		return "PORTAL_SAIDA"
	case COMIDA:
		return "COMIDA"
	case SAIDA:
		return "SAIDA"
	default:
		return "**DESCONHECIDO**"
	}
}
