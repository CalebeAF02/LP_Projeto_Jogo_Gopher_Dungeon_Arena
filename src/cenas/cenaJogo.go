package cenas

import (
	"Gopher_Dungeon_Arena/src/assets"
	"Gopher_Dungeon_Arena/src/componentes"
	"Gopher_Dungeon_Arena/src/config"
	"Gopher_Dungeon_Arena/src/ecs"
	"Gopher_Dungeon_Arena/src/entidades/geometria"
	"Gopher_Dungeon_Arena/src/entidades/outros"
	"Gopher_Dungeon_Arena/src/entidades/personagens"
	"Gopher_Dungeon_Arena/src/enum/entidades"
	"Gopher_Dungeon_Arena/src/interfaces"
	"Gopher_Dungeon_Arena/src/nivel"
	"Gopher_Dungeon_Arena/src/sistema"
	"Gopher_Dungeon_Arena/src/utils"
	"math/rand"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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
	coletadoTudo       bool
	miniMapaExibir     int
	miniMapaVisivel    bool
	fonteCache         *assets.FonteCache
	pontuacaoFaltante  int
	entrouNaSaida      bool
}

func NovoCenaJogo(game interfaces.IGame) *CenaJogo {
	mundo := geometria.NovoRetangulo(0, 0, config.MUNDO_LARGURA, config.MUNDO_ALTURA)
	entidades := make(map[ecs.EntidadeID]ecs.Entidade)
	camera := ecs.NovaCamera(mundo)
	miniMapa := ecs.NovoMiniMapa(mundo, geometria.NovoPonto(config.MM2_POS_X_MAPA, config.MM2_POS_Y_MAPA), camera)
	aleatorio := config.GeradorAleatorio()

	cj := CenaJogo{game: game, mundo: mundo, entidades: entidades, aleatorio: aleatorio, sistemaColisao: &sistema.SistemaColisao{}, contadorBotsMortos: 0, coletadoTudo: false, miniMapaVisivel: true, miniMapaExibir: 1, fonteCache: assets.FonteCacheCriar(), entrouNaSaida: false}

	cj.SetMiniMapa(miniMapa)
	cj.SetCamera(camera)

	cj.ReIniciar()

	return &cj
}

func (cj *CenaJogo) ReIniciar() {

	cj.sistemaColisao.SetCenaJogo(cj)

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

	//sistemaSpaw := sistema.SistemaSpawn{}

	//sistemaSpaw.SpawnarPortais(&cj)

	//sistemaSpaw.SpawnParedesAoRedor(&cj, 20)
	//SpawnParedesEspecificas(&g)
	//sistemaSpaw.SpawnLabirinto(&cj)

	//sistemaSpaw.SpawnBotDeCadaTipo(&cj)

	//objeto.NovaComida(&cj, geometria.NovoPonto(150, 200))
	//objeto.NovaComida(&cj, geometria.NovoPonto(300, 450))
	//objeto.NovaComida(&cj, geometria.NovoPonto(500, 650))
	//objeto.NovaComida(&cj, geometria.NovoPonto(600, 700))
	//objeto.NovaComida(&cj, geometria.NovoPonto(800, 950))

	//objeto.NovaSaida(&cj, geometria.NovoPonto(300, 500))

	//sistemaSpaw.SpawnJogadores(&cj)

	cj.entidades = make(map[ecs.EntidadeID]ecs.Entidade)
	cj.entrouNaSaida = false
	cj.coletadoTudo = false
	cj.contadorBotsMortos = 0

	nivel.CarregarNivel(cj)

}

func (s *CenaJogo) SpawnarBot(cj interfaces.ICenaJogo, movendo interfaces.Movimentador, posicao *geometria.Ponto) {
	b := personagens.NovoBot(cj, 0)
	b.SetNivelAleatorio()
	b.SetPosicao(posicao.GetX(), posicao.GetY())
	b.SetMovimentacao(movendo)
	//fmt.Printf("BOT <%s> | X: %f | Y: %f\n", b.GetMovendoTipo(), b.GetX(), b.GetY())
}

func (cp *CenaJogo) SetFonteCache(cache assets.FonteCache) {
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

	ctrlPressionado := ebiten.IsKeyPressed(ebiten.KeyControlLeft) || ebiten.IsKeyPressed(ebiten.KeyControlRight)

	if ctrlPressionado && inpututil.IsKeyJustPressed(ebiten.KeyM) {
		cj.miniMapaExibir += 1

		if cj.miniMapaExibir > 4 {
			cj.miniMapaExibir = 1
		}

		if cj.miniMapaExibir == 1 {
			cj.miniMapa.SetPosicao(config.MM1_POS_X_MAPA, config.MM1_POS_Y_MAPA)
		}
		if cj.miniMapaExibir == 2 {
			cj.miniMapa.SetPosicao(config.MM2_POS_X_MAPA, config.MM2_POS_Y_MAPA)
		}

		if cj.miniMapaExibir == 3 {
			cj.miniMapa.SetPosicao(config.MM3_POS_X_MAPA, config.MM3_POS_Y_MAPA)
		}

		if cj.miniMapaExibir == 4 {
			cj.miniMapa.SetPosicao(config.MM4_POS_X_MAPA, config.MM4_POS_Y_MAPA)
		}

	} else if ctrlPressionado && inpututil.IsKeyJustPressed(ebiten.KeyO) {
		cj.miniMapaVisivel = !cj.miniMapaVisivel
	}

	if cj.Concluiu() && cj.entrouNaSaida && inpututil.IsKeyJustPressed(ebiten.KeyEscape) {

		cj.game.ReiniciarMudarTelaMenuIniciar()

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

func (cj *CenaJogo) OrganizaPosicaoAleatoriaComida() *geometria.Ponto {
	larguraComida := float64(utils.COMIDA_TAMANHO_MUNDO)
	alturaComida := float64(utils.COMIDA_TAMANHO_MUNDO)

	// Limitamos as tentativas para evitar loop infinito se o mapa estiver cheio
	for tentativas := 0; tentativas < 100; tentativas++ {
		x := float64(cj.GetAleatorio().Intn(int(cj.GetMundo().PosXmax(utils.COMIDA_TAMANHO_MUNDO))))
		y := float64(cj.GetAleatorio().Intn(int(cj.GetMundo().PosYmax(utils.COMIDA_TAMANHO_MUNDO))))

		// REUTILIZAÇÃO: Usamos diretamente o método do jogo para checar barreiras (paredes)
		corpoTemporario := geometria.NovoRetangulo(x, y, larguraComida, alturaComida)
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
		} else if entidade.ExisteComponente(componentes.ENERGIA.String()) {

			energiaComp := entidade.GetComponente(componentes.ENERGIA.String())
			energia := energiaComp.(*componentes.Energia)

			if !energia.Status {
				cj.RemoverEntidade(entidade.GetID())
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

func (self *CenaJogo) ColetadoTudo(status bool) {
	self.coletadoTudo = status
}

func (self *CenaJogo) CapturouTudo() bool {
	return self.coletadoTudo
}

func (self *CenaJogo) MiniMapaEstaVisivel() bool {
	return self.miniMapaVisivel
}

func (self *CenaJogo) ObterPontuacaoFaltante() int {
	return self.pontuacaoFaltante
}

func (self *CenaJogo) SetFaltaPontuacao(pontos int) {
	self.pontuacaoFaltante = pontos
}

func (self *CenaJogo) Concluiu() bool {
	return self.pontuacaoFaltante == 0 && self.entrouNaSaida
}

func (self *CenaJogo) EntreiNaSaida() {
	self.entrouNaSaida = true
}

func (self *CenaJogo) EntrouNaSaida() bool {
	return self.entrouNaSaida
}

func (self *CenaJogo) ObterFonteCache() *assets.FonteCache {
	return self.fonteCache
}

func (cj *CenaJogo) GetNome() string {
	return "CENA_JOGO"
}
