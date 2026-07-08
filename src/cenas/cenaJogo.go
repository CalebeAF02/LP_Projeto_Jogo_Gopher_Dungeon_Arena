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
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type CenaJogo struct {
	game               interfaces.IGame
	proximo            int
	mundo              *geometria.Retangulo
	entidades          map[ecs.EntidadeID]ecs.Entidade
	entidadesLock      sync.RWMutex
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

func (self *CenaJogo) ReIniciar() {

	self.sistemaColisao.SetCenaJogo(self)

	self.sistemaAtualizar = []interfaces.ISistemaAtualizar{
		&sistema.SistemaInput{},
		&sistema.SistemaIA{},
		&sistema.SistemaSpawn{},
		&sistema.SistemaMovimento{},
		&sistema.SistemaEntidades{},
		&sistema.SistemaDebug{},
	}

	self.sistemaDesenhar = []interfaces.ISistemaDesenhar{
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

	self.entidadesLock.Lock()
	self.entidades = make(map[ecs.EntidadeID]ecs.Entidade)
	self.entidadesLock.Unlock()
	self.entrouNaSaida = false
	self.coletadoTudo = false
	self.contadorBotsMortos = 0

	nivel.CarregarNivel(self)

}

func (self *CenaJogo) SpawnarBot(cj interfaces.ICenaJogo, movendo interfaces.Movimentador, posicao *geometria.Ponto) {
	b := personagens.NovoBot(cj, 0)
	b.SetNivelAleatorio()
	b.SetPosicao(posicao.GetX(), posicao.GetY())
	b.SetMovimentacao(movendo)
	//fmt.Printf("BOT <%s> | X: %f | Y: %f\n", b.GetMovendoTipo(), b.GetX(), b.GetY())
}

func (self *CenaJogo) SetFonteCache(cache assets.FonteCache) {
}

func (self *CenaJogo) CriarEntidade() ecs.EntidadeID {
	entidade := ecs.EntidadeID(self.proximo)
	self.proximo++
	return entidade
}

func (self *CenaJogo) RemoverEntidade(entidade ecs.EntidadeID) {
	self.entidadesLock.Lock()
	delete(self.entidades, entidade)
	self.entidadesLock.Unlock()
}

func (self *CenaJogo) GetGame() interfaces.IGame {
	return self.game
}
func (self *CenaJogo) GetEntidades() map[ecs.EntidadeID]ecs.Entidade {
	self.entidadesLock.RLock()
	defer self.entidadesLock.RUnlock()

	copia := make(map[ecs.EntidadeID]ecs.Entidade, len(self.entidades))
	for k, v := range self.entidades {
		copia[k] = v
	}
	return copia
}

func (self *CenaJogo) GetTimes() []*outros.Time {
	listaTimes := []*outros.Time{}

	for _, e := range self.GetEntidades() {
		if e.GetTipo() == "TIME" {
			listaTimes = append(listaTimes, e.(*outros.Time))
		}
	}
	//fmt.Printf("Quantidade de Times %d\n", len(listaTimes))

	return listaTimes
}
func (self *CenaJogo) GetMundo() *geometria.Retangulo {
	return self.mundo
}
func (self *CenaJogo) GetAleatorio() *rand.Rand {
	return self.aleatorio
}
func (self *CenaJogo) GetLargura() float64 {
	return self.mundo.GetLargura()
}
func (self *CenaJogo) GetAltura() float64 {
	return self.mundo.GetAltura()
}
func (self *CenaJogo) GetCamera() *ecs.Camera {
	return self.camera
}

func (self *CenaJogo) GetMiniMapa() *ecs.MiniMapa {
	return self.miniMapa
}
func (self *CenaJogo) GetSistemaAtualizar() []interfaces.ISistemaAtualizar {
	return self.sistemaAtualizar
}
func (self *CenaJogo) GetSistemaDesenhar() []interfaces.ISistemaDesenhar {
	return self.sistemaDesenhar
}
func (self *CenaJogo) GetSistemaColisao() interfaces.ISistemaColisao {
	return self.sistemaColisao
}

func (self *CenaJogo) SetGame(game interfaces.IGame) {
	self.game = game
}
func (self *CenaJogo) SetEntidade(nEntidade ecs.EntidadeID, posicao ecs.Entidade) {
	self.entidadesLock.Lock()
	self.entidades[nEntidade] = posicao
	self.entidadesLock.Unlock()
}
func (self *CenaJogo) SetMiniMapa(miniMapa *ecs.MiniMapa) {
	self.miniMapa = miniMapa
}
func (self *CenaJogo) SetCamera(camera *ecs.Camera) {
	self.camera = camera
}

func (self *CenaJogo) OrganizarCamera() {
	// --- ATUALIZAÇÃO DA CÂMERA ---
	lTimes := self.GetTimes()

	if len(lTimes) > 0 && len(lTimes[0].GetJogadores()) > 0 {

		jogador := lTimes[0].GetJogador(0)

		self.GetCamera().OrganizarCameraPeloJogador(jogador.GetPosicao())
	}
}

func (self *CenaJogo) Input() {
	if ebiten.IsKeyPressed(ebiten.KeyP) {
		self.game.Pausar()
	}

	ctrlPressionado := ebiten.IsKeyPressed(ebiten.KeyControlLeft) || ebiten.IsKeyPressed(ebiten.KeyControlRight)

	if ctrlPressionado && inpututil.IsKeyJustPressed(ebiten.KeyM) {
		self.miniMapaExibir += 1

		if self.miniMapaExibir > 4 {
			self.miniMapaExibir = 1
		}

		if self.miniMapaExibir == 1 {
			self.miniMapa.SetPosicao(config.MM1_POS_X_MAPA, config.MM1_POS_Y_MAPA)
		}
		if self.miniMapaExibir == 2 {
			self.miniMapa.SetPosicao(config.MM2_POS_X_MAPA, config.MM2_POS_Y_MAPA)
		}

		if self.miniMapaExibir == 3 {
			self.miniMapa.SetPosicao(config.MM3_POS_X_MAPA, config.MM3_POS_Y_MAPA)
		}

		if self.miniMapaExibir == 4 {
			self.miniMapa.SetPosicao(config.MM4_POS_X_MAPA, config.MM4_POS_Y_MAPA)
		}

	} else if ctrlPressionado && inpututil.IsKeyJustPressed(ebiten.KeyO) {
		self.miniMapaVisivel = !self.miniMapaVisivel
	}

	if self.Concluiu() && self.entrouNaSaida && inpututil.IsKeyJustPressed(ebiten.KeyEscape) {

		self.game.ReiniciarMudarTelaMenuIniciar()

	}

}

func (self *CenaJogo) Update() error {
	self.Input()
	self.OrganizarCamera()
	for _, sistema := range self.sistemaAtualizar {
		sistema.Atualizar(self)
	}

	self.RemoverEntidadesMortas()

	return nil
}

func (self *CenaJogo) Draw(tela *ebiten.Image) {
	for _, sistema := range self.sistemaDesenhar {
		sistema.Desenhar(self, tela)
	}
}

func (self *CenaJogo) Pausar() {
	self.game.Pausar()
}

func (self *CenaJogo) OrganizaPosicaoAleatoriaBot() *geometria.Ponto {
	larguraBot := float64(utils.BOT_TAMANHO_MUNDO)
	alturaBot := float64(utils.BOT_TAMANHO_MUNDO)

	// Limitamos as tentativas para evitar loop infinito se o mapa estiver cheio
	for tentativas := 0; tentativas < 100; tentativas++ {
		x := float64(self.GetAleatorio().Intn(int(self.GetMundo().PosXmax(utils.BOT_TAMANHO_MUNDO))))
		y := float64(self.GetAleatorio().Intn(int(self.GetMundo().PosYmax(utils.BOT_TAMANHO_MUNDO))))

		// REUTILIZAÇÃO: Usamos diretamente o método do jogo para checar barreiras (paredes)
		corpoTemporario := geometria.NovoRetangulo(x, y, larguraBot, alturaBot)
		if !self.sistemaColisao.ColideComTipo(corpoTemporario, entidades.PAREDE.String()) {
			return geometria.NovoPonto(x, y)
		}
	}

	return nil
}

func (self *CenaJogo) OrganizaPosicaoAleatoriaComida() *geometria.Ponto {
	larguraComida := float64(utils.COMIDA_TAMANHO_MUNDO)
	alturaComida := float64(utils.COMIDA_TAMANHO_MUNDO)

	// Limitamos as tentativas para evitar loop infinito se o mapa estiver cheio
	for tentativas := 0; tentativas < 100; tentativas++ {
		x := float64(self.GetAleatorio().Intn(int(self.GetMundo().PosXmax(utils.COMIDA_TAMANHO_MUNDO))))
		y := float64(self.GetAleatorio().Intn(int(self.GetMundo().PosYmax(utils.COMIDA_TAMANHO_MUNDO))))

		// REUTILIZAÇÃO: Usamos diretamente o método do jogo para checar barreiras (paredes)
		corpoTemporario := geometria.NovoRetangulo(x, y, larguraComida, alturaComida)
		if !self.sistemaColisao.ColideComTipo(corpoTemporario, entidades.PAREDE.String()) {
			return geometria.NovoPonto(x, y)
		}
	}

	return nil
}

func (self *CenaJogo) RemoverEntidadesMortas() {
	for _, entidade := range self.GetEntidades() {
		if entidade.ExisteComponente(componentes.VIDA.String()) {

			vidaComp := entidade.GetComponente(componentes.VIDA.String())
			vida := vidaComp.(*componentes.Vida)

			if !vida.Status {
				self.RemoverEntidade(entidade.GetID())

				if vida.TipoOrganismo == "BOT" {
					self.ContarEntidadesMortas()
				}
			}
		} else if entidade.ExisteComponente(componentes.ENERGIA.String()) {

			energiaComp := entidade.GetComponente(componentes.ENERGIA.String())
			energia := energiaComp.(*componentes.Energia)

			if !energia.Status {
				self.RemoverEntidade(entidade.GetID())
			}
		}
	}
}
func (self *CenaJogo) GetContadorMortos() string {
	return strconv.Itoa(self.contadorBotsMortos)
}

func (self *CenaJogo) ContarEntidadesMortas() {
	self.contadorBotsMortos += 1
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

func (self *CenaJogo) GetNome() string {
	return "CENA_JOGO"
}
