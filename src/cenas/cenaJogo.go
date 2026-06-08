package cenas

import (
	"Gopher_Dungeon_Arena/src/config"
	"Gopher_Dungeon_Arena/src/ecs"
	"Gopher_Dungeon_Arena/src/entidades/geometria"
	"Gopher_Dungeon_Arena/src/entidades/outros"
	"Gopher_Dungeon_Arena/src/enum/componentes"
	"Gopher_Dungeon_Arena/src/enum/entidades"
	"Gopher_Dungeon_Arena/src/interfaces"
	"Gopher_Dungeon_Arena/src/sistema"
	"Gopher_Dungeon_Arena/src/utils"
	"math/rand"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
)

type CenaJogo struct {
	game               interfaces.IGame
	proximo            int
	mundo              *geometria.Retangulo
	entidades          map[ecs.EntidadeID]ecs.Entidade
	camera             *ecs.Camera
	miniMapa           *ecs.MiniMapa
	aleatorio          *rand.Rand
	sistemaAtualizar   []interfaces.ISistemaAtualizar
	sistemaDesenhar    []interfaces.ISistemaDesenhar
	sistemaColisao     interfaces.ISistemaColisao
	contadorBotsMortos int
}

func NovoCenaJogo(game interfaces.IGame) *CenaJogo {
	mundo := geometria.NovoRetangulo(0, 0, config.MUNDO_LARGURA, config.MUNDO_ALTURA)
	entidades := make(map[ecs.EntidadeID]ecs.Entidade)
	camera := ecs.NovaCamera(mundo)
	miniMapa := ecs.NovoMiniMapa(mundo, geometria.NovoPonto(config.POS_X_MAPA, config.POS_Y_MAPA), camera)
	aleatorio := config.GeradorAleatorio()

	cj := CenaJogo{game: game, mundo: mundo, entidades: entidades, aleatorio: aleatorio, sistemaColisao: &sistema.SistemaColisao{}, contadorBotsMortos: 0}

	cj.sistemaColisao.SetCenaJogo(&cj)

	cj.sistemaAtualizar = []interfaces.ISistemaAtualizar{
		&sistema.SistemaInput{},
		&sistema.SistemaIA{},
		&sistema.SistemaSpawn{},
		&sistema.SistemaMovimento{},
		&sistema.SistemaEntidades{},
		&sistema.SistemaDebug{},
	}

	cj.sistemaDesenhar = []interfaces.ISistemaDesenhar{
		&sistema.SistemaDesenhar{},
	}

	cj.SetMiniMapa(miniMapa)
	cj.SetCamera(camera)

	sistemaSpaw := sistema.SistemaSpawn{}
	sistemaSpaw.SpawnJogadores(&cj)

	sistemaSpaw.SpawnarPortais(&cj)

	sistemaSpaw.SpawnParedesAoRedor(&cj, 20)
	//SpawnParedesEspecificas(&g)
	sistemaSpaw.SpawnLabirinto(&cj)

	sistemaSpaw.SpawnBotDeCadaTipo(&cj)

	return &cj
}

func (cj *CenaJogo) CriarEntidade() ecs.EntidadeID {
	entidade := ecs.EntidadeID(cj.proximo)
	cj.proximo++
	return entidade
}

func (cj *CenaJogo) RemoverEntidade(entidade ecs.EntidadeID) {
	delete(cj.entidades, entidade)
}

func (cj *CenaJogo) GetGame() interfaces.IGame {
	return cj.game
}
func (cj *CenaJogo) GetEntidades() map[ecs.EntidadeID]ecs.Entidade {
	return cj.entidades
}

func (cj *CenaJogo) GetTimes() []*outros.Time {
	listaTimes := []*outros.Time{}

	for _, e := range cj.GetEntidades() {
		if e.GetTipo() == "TIME" {
			listaTimes = append(listaTimes, e.(*outros.Time))
		}
	}
	//fmt.Printf("Quantidade de Times %d\n", len(listaTimes))

	return listaTimes
}
func (cj *CenaJogo) GetMundo() *geometria.Retangulo {
	return cj.mundo
}
func (cj *CenaJogo) GetAleatorio() *rand.Rand {
	return cj.aleatorio
}
func (cj *CenaJogo) GetLargura() float64 {
	return cj.mundo.GetLargura()
}
func (cj *CenaJogo) GetAltura() float64 {
	return cj.mundo.GetAltura()
}
func (cj *CenaJogo) GetCamera() *ecs.Camera {
	return cj.camera
}

func (cj *CenaJogo) GetMiniMapa() *ecs.MiniMapa {
	return cj.miniMapa
}
func (cj *CenaJogo) GetSistemaAtualizar() []interfaces.ISistemaAtualizar {
	return cj.sistemaAtualizar
}
func (cj *CenaJogo) GetSistemaDesenhar() []interfaces.ISistemaDesenhar {
	return cj.sistemaDesenhar
}
func (cj *CenaJogo) GetSistemaColisao() interfaces.ISistemaColisao {
	return cj.sistemaColisao
}

func (cj *CenaJogo) SetGame(game interfaces.IGame) {
	cj.game = game
}
func (cj *CenaJogo) SetEntidade(nEntidade ecs.EntidadeID, posicao ecs.Entidade) {
	cj.entidades[nEntidade] = posicao
}
func (cj *CenaJogo) SetMiniMapa(miniMapa *ecs.MiniMapa) {
	cj.miniMapa = miniMapa
}
func (cj *CenaJogo) SetCamera(camera *ecs.Camera) {
	cj.camera = camera
}

func (cj *CenaJogo) OrganizarCamera() {
	// --- ATUALIZAÇÃO DA CÂMERA ---
	lTimes := cj.GetTimes()

	if len(lTimes) > 0 && len(lTimes[0].GetJogadores()) > 0 {

		jogador := lTimes[0].GetJogador(0)

		cj.GetCamera().OrganizarCameraPeloJogador(jogador.GetPosicao())
	}
}

func (cj *CenaJogo) Input() {
	if ebiten.IsKeyPressed(ebiten.KeyP) {
		cj.game.Pausar()
	}
}

func (cj *CenaJogo) Update() error {
	cj.Input()
	cj.OrganizarCamera()
	for _, sistema := range cj.sistemaAtualizar {
		sistema.Atualizar(cj)
	}

	cj.RemoverEntidadesMortas()

	return nil
}

func (cj *CenaJogo) Draw(tela *ebiten.Image) {
	for _, sistema := range cj.sistemaDesenhar {
		sistema.Desenhar(cj, tela)
	}
}

func (cj *CenaJogo) Pausar() {
	cj.game.Pausar()
}

func (cj *CenaJogo) OrganizaPosicaoAleatoriaBot() *geometria.Ponto {
	larguraBot := float64(utils.BOT_TAMANHO_MUNDO)
	alturaBot := float64(utils.BOT_TAMANHO_MUNDO)

	// Limitamos as tentativas para evitar loop infinito se o mapa estiver cheio
	for tentativas := 0; tentativas < 100; tentativas++ {
		x := float64(cj.GetAleatorio().Intn(int(cj.GetMundo().PosXmax(utils.BOT_TAMANHO_MUNDO))))
		y := float64(cj.GetAleatorio().Intn(int(cj.GetMundo().PosYmax(utils.BOT_TAMANHO_MUNDO))))

		// REUTILIZAÇÃO: Usamos diretamente o método do jogo para checar barreiras (paredes)
		corpoTemporario := geometria.NovoRetangulo(x, y, larguraBot, alturaBot)
		if !cj.sistemaColisao.ColideComTipo(corpoTemporario, entidades.PAREDE.String()) {
			return geometria.NovoPonto(x, y)
		}
	}

	return nil
}

func (cj *CenaJogo) RemoverEntidadesMortas() {
	for _, entidade := range cj.GetEntidades() {
		if entidade.ExisteComponente(componentes.VIDA.String()) {

			vidaComp := entidade.GetComponente(componentes.VIDA.String())
			vida := vidaComp.(*componentes.Vida)

			if !vida.Status {
				cj.RemoverEntidade(entidade.GetID())

				if vida.TipoOrganismo == "BOT" {
					cj.ContarEntidadesMortas()
				}
			}

		}
	}
}
func (cj *CenaJogo) GetContadorMortos() string {
	return strconv.Itoa(cj.contadorBotsMortos)
}

func (cj *CenaJogo) ContarEntidadesMortas() {
	cj.contadorBotsMortos += 1
}

func (cj *CenaJogo) GetNome() string {
	return "CENA_JOGO"
}
