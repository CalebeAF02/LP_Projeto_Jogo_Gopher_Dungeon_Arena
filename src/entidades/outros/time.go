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

func (j *Time) GetID() ecs.EntidadeID {
	return j.entidadeID
}

func (t *Time) Adicionnar(jogador *personagens.Jogador) {
	jogador.SetCor(t.cor)
	t.jogadores = append(t.jogadores, jogador)
}

func (t *Time) EstaVivo() bool {
	for _, jogador := range t.jogadores {
		if jogador.ObterVida().Status {
			return true
		}
	}
	return false
}

func (t *Time) GetNome() string {
	return t.nome
}
func (t *Time) GetJogador(id int) *personagens.Jogador {
	return t.jogadores[id]
}
func (t *Time) GetJogadores() []*personagens.Jogador {
	listaJogadore := []*personagens.Jogador{}

	for _, e := range t.cenaJogo.GetEntidades() {
		if e.GetTipo() == "JOGADOR" {
			listaJogadore = append(listaJogadore, e.(*personagens.Jogador))
		}
	}
	//fmt.Printf("Quantidade de Jogador %d\n", len(listaJogadore))

	return listaJogadore
}

func (t *Time) Posicoes() {
	for i, j := range t.jogadores {
		erro := fmt.Sprintf("Jogador_%d: %s esta na posicao %s\n", i+1, j.GetNome(), j.GetPosicao().ToString())
		fmt.Println(erro)
	}
}

func (t *Time) GetTipo() string {
	return entidades.TIME.String()
}
func (t *Time) GetQuantidade() int {
	return len(t.jogadores)
}
func (t *Time) GetPosicaoTime() *geometria.Ponto {
	jogadores := t.GetJogadores()

	for _, j := range jogadores {
		//fmt.Println("Jogador : ", j.GetNome(), " esta: ", j.EstaVivo())

		if j.EstaVivo() {
			return j.GetPosicao()
		}
	}
	fmt.Println("Erro: Posicao nao encontrada!")
	return nil
}

func (t *Time) Atualizar() {
	t.GetQuantidade()
}
func (t *Time) Desenhar(tela *ebiten.Image) {
	t.GetQuantidade()
}
func (t *Time) DesenharMapa(tela *ebiten.Image, mapaX float64, mapaY float64) {
	t.GetQuantidade()
}

func (e *Time) GetComponente(id string) interface{} {
	return e.Componentes[id]
}

func (e *Time) AdicionarComponente(id string, comp interface{}) {
	if e.Componentes == nil {
		e.Componentes = make(map[string]interface{})
	}
	e.Componentes[id] = comp
}

func (e *Time) ExisteComponente(id string) bool {
	_, existe := e.Componentes[id]
	return existe
}
