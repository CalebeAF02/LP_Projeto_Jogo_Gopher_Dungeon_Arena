package portal

import (
	"Gopher_Dungeon_Arena/src/config"
	"Gopher_Dungeon_Arena/src/ecs"
	"Gopher_Dungeon_Arena/src/entidades/geometria"
	"Gopher_Dungeon_Arena/src/entidades/objeto"
	"Gopher_Dungeon_Arena/src/entidades/outros"
	"Gopher_Dungeon_Arena/src/enum/componentes"
	"Gopher_Dungeon_Arena/src/enum/entidades"
	"Gopher_Dungeon_Arena/src/interfaces"
	"Gopher_Dungeon_Arena/src/sistema"
	"Gopher_Dungeon_Arena/src/utils"
	portais "Gopher_Dungeon_Arena/src/utils/sprits"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type CenaJogo struct {
	game             interfaces.IGame
	proximo          int
	mundo            *geometria.Retangulo
	entidades        map[ecs.EntidadeID]ecs.Entidade
	camera           *ecs.Camera
	miniMapa         *ecs.MiniMapa
	aleatorio        *rand.Rand
	sistemaAtualizar []interfaces.ISistemaAtualizar
	sistemaDesenhar  []interfaces.ISistemaDesenhar
	sistemaColisao   interfaces.ISistemaColisao
	bancoPortais     *portais.SpriteSheetPortal // 2. ARMAZENAR O BANCO DE IMAGENS DOS PORTAIS
}

func NovoCenaJogo(game interfaces.IGame) *CenaJogo {
	mundo := geometria.NovoRetangulo(0, 0, config.MUNDO_LARGURA, config.MUNDO_ALTURA)
	entidades := make(map[ecs.EntidadeID]ecs.Entidade)
	camera := ecs.NovaCamera(mundo)
	miniMapa := ecs.NovoMiniMapa(mundo, geometria.NovoPonto(config.POS_X_MAPA, config.POS_Y_MAPA), camera)
	aleatorio := config.GeradorAleatorio()

	cj := CenaJogo{game: game, mundo: mundo, entidades: entidades, aleatorio: aleatorio, sistemaColisao: &sistema.SistemaColisao{}}

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

	// 3. CARREGAR A TEXTURA E INICIALIZAR O BANCO DE PORTAIS NA MEMÓRIA
	// Certifique-se de ajustar o caminho do arquivo de acordo com a sua pasta de assets
	// --- CORREÇÃO DO CAMINHO DOS ASSETS ---
	imgImg, _, err := ebitenutil.NewImageFromFile("src/assets/sprits/portal/varios_portais.png")
	if err != nil {
		log.Fatal("Erro fatal ao carregar imagem dos portais: ", err)
	}
	cj.bancoPortais = portais.NovoSpriteSheetPortal(imgImg)

	sistemaSpaw := sistema.SistemaSpawn{}
	sistemaSpaw.SpawnJogadores(&cj)

	sistemaSpaw.SpawnParedesAoRedor(&cj, 20)
	//SpawnParedesEspecificas(&g)
	sistemaSpaw.SpawnLabirinto(&cj)

	sistemaSpaw.SpawnBotDeCadaTipo(&cj)

	// 4. CRIAR OS PORTAIS NO MAPA USANDO O BANCO DE IMAGENS
	cj.SpawnPortais()

	return &cj
}

// Método auxiliar criado para instanciar e posicionar os portais usando os novos sprites fatiados
func (cj *CenaJogo) SpawnPortais() {
	// Exemplo de pontos aleatórios no mapa que não colidem com paredes
	pontoEntrada := cj.OrganizaPosicaoAleatoriaBot()
	pontoSaida := cj.OrganizaPosicaoAleatoriaBot()

	if pontoEntrada != nil {
		// Pega o sprite Laranja do banco e cria o PortalSpriteEntrada
		spriteEntrada := cj.bancoPortais.ObterPortal(portais.PortalLaranja)
		idEntrada := cj.CriarEntidade()
		portalEntrada := objeto.NovoPortalSpriteEntrada(cj, int64(idEntrada), spriteEntrada)
		portalEntrada.SetPosicao(pontoEntrada.GetX(), pontoEntrada.GetY())
	}

	if pontoSaida != nil {
		// Pega o sprite Verde Claro do banco e cria o PortalSpriteSaida
		spriteSaida := cj.bancoPortais.ObterPortal(portais.PortalVerdeClaro)
		idSaida := cj.CriarEntidade()
		portalSaida := objeto.NovoPortalSpriteSaida(cj, int64(idSaida), spriteSaida)
		portalSaida.SetPosicao(pontoSaida.GetX(), pontoSaida.GetY())
	}
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

// Getter para caso algum sistema precise acessar o banco de portais externamente
func (cj *CenaJogo) GetBancoPortais() *portais.SpriteSheetPortal {
	return cj.bancoPortais
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
			}

		}
	}
}

func (cj *CenaJogo) GetNome() string {
	return "CENA_JOGO"
}
