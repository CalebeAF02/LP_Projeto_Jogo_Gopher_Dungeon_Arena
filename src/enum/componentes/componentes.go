package componentes

import "Gopher_Dungeon_Arena/src/enum/entidades"

type ComponenteTipo int

const (
	CORPO    ComponenteTipo = iota
	SUB_TIPO ComponenteTipo = iota
	VIDA     ComponenteTipo = iota
	NIVEL    ComponenteTipo = iota
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

func (v *Vida) EstaVivo(tipo string) bool {
	if v.TipoOrganismo == entidades.JOGADOR.String() {
		if v.Quantidade > 0 {
			v.Status = true
			return v.Status
		}
		v.Status = false
		return v.Status
	} else if tipo == "BOT" {
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

func (v *Vida) PerdeSangue(rit int) {
	v.Sangue -= rit

	if v.Sangue <= 0 && v.TipoOrganismo == entidades.JOGADOR.String() {
		v.Renasce(3)
	}
}

func (v *Vida) Renasce(valor int) {
	if v.EstaVivo("JOGADOR") {
		v.TiraUmaVida()

		if v.Quantidade > 0 {
			v.ResetaSangue(valor)
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
