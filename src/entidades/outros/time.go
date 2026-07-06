package outros

import (
	"Gopher_Dungeon_Arena/src/ecs"
	"Gopher_Dungeon_Arena/src/entidades/geometria"
	"Gopher_Dungeon_Arena/src/entidades/personagens"
	"Gopher_Dungeon_Arena/src/enum/entidades"
	"Gopher_Dungeon_Arena/src/interfaces"
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Time struct {
	cenaJogo    interfaces.ICenaJogo
	entidadeID  ecs.EntidadeID
	nome        string
	jogadores   []*personagens.Jogador
	cor         color.Color
	Componentes map[string]interface{}
}

func NovoTime(cenaJogo interfaces.ICenaJogo, n string, cor color.Color) *Time {
	nEntidade := cenaJogo.CriarEntidade()
	nTime := Time{cenaJogo: cenaJogo, entidadeID: nEntidade, nome: n, cor: cor}
	cenaJogo.SetEntidade(nEntidade, &nTime)
	return &nTime
}

func (self *Time) GetID() ecs.EntidadeID {
	return self.entidadeID
}

func (self *Time) Adicionnar(jogador *personagens.Jogador) {
	jogador.SetCor(self.cor)
	self.jogadores = append(self.jogadores, jogador)
}

func (self *Time) EstaVivo() bool {
	for _, jogador := range self.jogadores {
		if jogador.ObterVida().Status {
			return true
		}
	}
	return false
}

func (self *Time) GetNome() string {
	return self.nome
}
func (self *Time) GetJogador(id int) *personagens.Jogador {
	return self.jogadores[id]
}
func (self *Time) GetJogadores() []*personagens.Jogador {
	listaJogadore := []*personagens.Jogador{}

	for _, e := range self.cenaJogo.GetEntidades() {
		if e.GetTipo() == "JOGADOR" {
			listaJogadore = append(listaJogadore, e.(*personagens.Jogador))
		}
	}
	//fmt.Printf("Quantidade de Jogador %d\n", len(listaJogadore))

	return listaJogadore
}

func (self *Time) Posicoes() {
	for i, j := range self.jogadores {
		erro := fmt.Sprintf("Jogador_%d: %s esta na posicao %s\n", i+1, j.GetNome(), j.GetPosicao().ToString())
		fmt.Println(erro)
	}
}

func (self *Time) GetTipo() string {
	return entidades.TIME.String()
}
func (self *Time) GetQuantidade() int {
	return len(self.jogadores)
}
func (self *Time) GetPosicaoTime() *geometria.Ponto {
	jogadores := self.GetJogadores()

	for _, j := range jogadores {
		//fmt.Println("Jogador : ", j.GetNome(), " esta: ", j.EstaVivo())

		if j.EstaVivo() {
			return j.GetPosicao()
		}
	}
	fmt.Println("Erro: Posicao nao encontrada!")
	return nil
}

func (self *Time) Atualizar() {
	self.GetQuantidade()
}
func (self *Time) Desenhar(tela *ebiten.Image) {
	self.GetQuantidade()
}
func (self *Time) DesenharMapa(tela *ebiten.Image, mapaX float64, mapaY float64) {
	self.GetQuantidade()
}

func (self *Time) GetComponente(id string) interface{} {
	return self.Componentes[id]
}

func (self *Time) AdicionarComponente(id string, comp interface{}) {
	if self.Componentes == nil {
		self.Componentes = make(map[string]interface{})
	}
	self.Componentes[id] = comp
}

func (self *Time) ExisteComponente(id string) bool {
	_, existe := self.Componentes[id]
	return existe
}
