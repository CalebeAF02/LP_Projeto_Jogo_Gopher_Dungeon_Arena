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
	ENERGIA                  ComponenteTipo = iota
	PONTUACAO                ComponenteTipo = iota
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
	case ENERGIA:
		return "ENERGIA"
	case PONTUACAO:
		return "PONTUACAO"
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

type Energia struct {
	Valor  int
	Status bool
}

type Pontuacao struct {
	Coletado      int
	Requisito     int
	EntreiNaSaida bool
}

func (self *Vida) EstaVivo() bool {
	if self.TipoOrganismo == entidades.JOGADOR.String() {
		if self.Quantidade > 0 {
			self.Status = true
			return self.Status
		}
		self.Status = false
		return self.Status
	} else if self.TipoOrganismo == entidades.BOT.String() {
		if self.Sangue > 0 {
			self.Status = true
			return self.Status
		}
		self.Status = false
		return self.Status
	}
	return false
}

func (self *Vida) CorrigeSangue(nivel int) {
	self.Sangue = 100 * nivel
}

func (self *Vida) PerdeSangue(rit int, nivel int) {
	self.Sangue -= rit

	if self.Sangue <= 0 && self.TipoOrganismo == entidades.JOGADOR.String() {
		self.Renasce(nivel)
	}
}

func (self *Vida) Renasce(nivel int) {
	if self.EstaVivo() {
		self.TiraUmaVida()

		if self.Quantidade > 0 {
			self.ResetaSangue(nivel)
		} else {
			//fmt.Println("O jogador " + v.nome + " morreu!")
		}

	} else {
		//fmt.Println("O jogador " + v.nome + " já está morto!")
	}
}

func (self *Vida) TiraUmaVida() {
	if self.Quantidade > 0 {
		self.Quantidade -= 1
	}
}

func (self *Vida) AcrescentaUmaVida() bool {
	if self.Quantidade < 3 {
		self.Quantidade += 1
		return true
	} else {
		return false
		//fmt.Println("A vida do jogador " + v.nome + " já está cheia!")
	}
}

func (self *Vida) Colisao() {
}

func (self *Vida) ResetaVida(valor int) {
	self.Quantidade = valor
}

func (self *Vida) ResetaSangue(nivel int) {
	self.CorrigeSangue(nivel)
}
