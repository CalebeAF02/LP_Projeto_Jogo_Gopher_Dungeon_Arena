package entidades

type EntidadeTipo int

const (
	POSICAO EntidadeTipo = iota
	BOT
	JOGADOR
	TIME
	PAREDE
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
	default:
		return "**DESCONHECIDO**"
	}
}
