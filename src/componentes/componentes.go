package componentes

import (
	"Gopher_Dungeon_Arena/src/ecs"
	"Gopher_Dungeon_Arena/src/enum/entidades"
	"Gopher_Dungeon_Arena/src/interfaces"
	"image/color"
)

type ComponenteTipo int

const (
	CORPO                    ComponenteTipo = iota
	SUB_TIPO                 ComponenteTipo = iota
	VIDA                     ComponenteTipo = iota
	NIVEL                    ComponenteTipo = iota
	ENVIANDO_TELETRANSPORTE  ComponenteTipo = iota
	ATIVIDADE                ComponenteTipo = iota
	RECEBENDO_TELETRANSPORTE ComponenteTipo = iota
	MOVIMENTO                ComponenteTipo = iota
)

const (
	AIVIDADE_MOVIMENTO      int = iota
	AIVIDADE_TELETRANSPORTE int = iota
)

func (t ComponenteTipo) String() string {
	switch t {
	case CORPO:
		return "CORPO"
	case SUB_TIPO:
		return "SUB_TIPO"
	case VIDA:
		return "VIDA"
	case NIVEL:
		return "NIVEL"
	case ENVIANDO_TELETRANSPORTE:
		return "ENVIANDO_TELETRANSPORTE"
	case RECEBENDO_TELETRANSPORTE:
		return "RECEBENDO_TELETRANSPORTE"
	case ATIVIDADE:
		return "ATIVIDADE"
	case MOVIMENTO:
		return "MOVIMENTO"
	default:
		return "**DESCONHECIDO**"
	}
}

type SubTipo struct {
	Valor string
}

type Vida struct {
	TipoOrganismo string
	Quantidade    int
	Status        bool
	Sangue        int
}

type Nivel struct {
	Valor      int
	Progressao int
}

type EnviandoTeletransporte struct {
	TemBot         bool
	Bot            ecs.Entidade
	Contagem       int
	ConectadoSaida ecs.Entidade
}

type Atividade struct {
	Acao int
}

type RecebendoTeletransporte struct {
	TemBot   bool
	Bot      ecs.Entidade
	Contagem int
}

type Movimento struct {
	Tipo interfaces.Movimentador
	Cor  color.Color
}

func (v *Vida) EstaVivo() bool {
	if v.TipoOrganismo == entidades.JOGADOR.String() {
		if v.Quantidade > 0 {
			v.Status = true
			return v.Status
		}
		v.Status = false
		return v.Status
	} else if v.TipoOrganismo == entidades.BOT.String() {
		if v.Sangue > 0 {
			v.Status = true
			return v.Status
		}
		v.Status = false
		return v.Status
	}
	return false
}

func (v *Vida) CorrigeSangue(nivel int) {
	v.Sangue = 100 * nivel
}

func (v *Vida) PerdeSangue(rit int, nivel int) {
	v.Sangue -= rit

	if v.Sangue <= 0 && v.TipoOrganismo == entidades.JOGADOR.String() {
		v.Renasce(nivel)
	}
}

func (v *Vida) Renasce(nivel int) {
	if v.EstaVivo() {
		v.TiraUmaVida()

		if v.Quantidade > 0 {
			v.ResetaSangue(nivel)
		} else {
			//fmt.Println("O jogador " + v.nome + " morreu!")
		}

	} else {
		//fmt.Println("O jogador " + v.nome + " já está morto!")
	}
}

func (v *Vida) TiraUmaVida() {
	if v.Quantidade > 0 {
		v.Quantidade -= 1
	}
}

func (v *Vida) AcrescentaUmaVida() bool {
	if v.Quantidade < 3 {
		v.Quantidade += 1
		return true
	} else {
		return false
		//fmt.Println("A vida do jogador " + v.nome + " já está cheia!")
	}
}

func (v *Vida) Colisao() {
}

func (v *Vida) ResetaVida(valor int) {
	v.Quantidade = valor
}

func (v *Vida) ResetaSangue(nivel int) {
	v.CorrigeSangue(nivel)
}
