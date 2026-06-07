package cenas

import (
	"Gopher_Dungeon_Arena/src/config"
	"Gopher_Dungeon_Arena/src/ecs"
	"Gopher_Dungeon_Arena/src/entidades/geometria"
	"Gopher_Dungeon_Arena/src/entidades/outros"
	"Gopher_Dungeon_Arena/src/entidades/personagens"
	"Gopher_Dungeon_Arena/src/enum/componentes"
	"Gopher_Dungeon_Arena/src/enum/entidades"
	"Gopher_Dungeon_Arena/src/interfaces"
	"Gopher_Dungeon_Arena/src/sistema"
	"Gopher_Dungeon_Arena/src/utils"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type CenaJogo struct {
	proximo          int
	mundo            *geometria.Retangulo
	entidades        map[ecs.EntidadeID]ecs.Entidade
	camera           *ecs.Camera
	miniMapa         *ecs.MiniMapa
	aleatorio        *rand.Rand
	sistemaAtualizar []interfaces.ISistemaAtualizar
	sistemaDesenhar  []interfaces.ISistemaDesenhar
}

func NovoCenaJogo() *CenaJogo {
	mundo := geometria.NovoRetangulo(0, 0, config.MUNDO_LARGURA, config.MUNDO_ALTURA)
	entidades := make(map[ecs.EntidadeID]ecs.Entidade)
	camera := ecs.NovaCamera(mundo)
	miniMapa := ecs.NovoMiniMapa(mundo, geometria.NovoPonto(config.POS_X_MAPA, config.POS_Y_MAPA), camera)
	aleatorio := config.GeradorAleatorio()

	cj := CenaJogo{mundo: mundo, entidades: entidades, aleatorio: aleatorio}

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

	sj := sistema.SistemaSpawn{}
	sj.SpawnJogadores(&cj)

	sj.SpawnParedesAoRedor(&cj, 20)
	//SpawnParedesEspecificas(&g)
	sj.SpawnLabirinto(&cj)

	sj.SpawnBotDeCadaTipo(&cj)

	return &cj
}


func (cj *CenaJogo) CriarEntidade() ecs.EntidadeID {
	entidade := ecs.EntidadeID(cj.proximo)
	cj.proximo++
	return entidade
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

func (cj *CenaJogo) Update() error {
	cj.OrganizarCamera()
	for _, sistema := range cj.sistemaAtualizar {
		sistema.Atualizar(cj)
	}
	return nil
}

func (cj *CenaJogo) Draw(tela *ebiten.Image) {
	for _, sistema := range cj.sistemaDesenhar {
		sistema.Desenhar(cj, tela)
	}
}

func (cj *CenaJogo) Sair() {
	// Limpando entidades
	cj.entidades = nil

	// Limpando estruturas principais
	cj.mundo = nil
	cj.camera = nil
	cj.miniMapa = nil
	cj.aleatorio = nil

	// Zerando contadores
	cj.proximo = 0
}

func (cj *CenaJogo) CriarRespostaColisao(status bool, tipo string, subTipo string) *ecs.RespostaColisao {
	nRespostaColisao := ecs.RespostaColisao{Status: status, Tipo: tipo, SubTipo: subTipo}

	return &nRespostaColisao
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
		if !cj.ColideComTipo(corpoTemporario, entidades.PAREDE.String()) {
			return geometria.NovoPonto(x, y)
		}
	}

	return nil
}

func (cj *CenaJogo) VaiColidir(meuCorpoAtual *geometria.Retangulo, proximoCorpo *geometria.Retangulo) *ecs.RespostaColisao {
	for _, e := range cj.GetEntidades() {
		tipo := e.GetTipo()
		if tipo == entidades.PAREDE.String() || tipo == entidades.JOGADOR.String() || tipo == entidades.BOT.String() {
			if corpoEntidade := e.GetComponente(componentes.CORPO.String()); corpoEntidade != nil {
				corpo := corpoEntidade.(*geometria.Retangulo)

				// EVITA AUTO-COLISÃO REAL:
				// Se a entidade da lista tiver exatamente a mesma posição X e Y do meu corpo atual,
				// significa que essa entidade SOU EU MESMO na tabela do ECS. Ignoramos!
				if corpo.GetX() == meuCorpoAtual.GetX() && corpo.GetY() == meuCorpoAtual.GetY() {
					continue
				}

				// Agora sim, testa se a minha PRÓXIMA posição vai bater em OUTRA entidade
				if proximoCorpo.Colide(corpo) {
					if meuCorpoAtual.Colide(corpo) {
						continue
					}
					if tipo == entidades.BOT.String() {
						if sub_tipo := e.GetComponente(componentes.SUB_TIPO.String()); sub_tipo != nil {

							sub_tipo_valor := sub_tipo.(*personagens.SubTipo)
							return cj.CriarRespostaColisao(true, tipo, sub_tipo_valor.Valor)
						} else {
							return cj.CriarRespostaColisao(true, tipo, "")
						}
					} else {
						return cj.CriarRespostaColisao(true, tipo, "")

					}
				}
			}
		}
	}
	return cj.CriarRespostaColisao(false, "", "")
}

// ColideComTipo isola uma busca específica (útil para o Spawn ou lógicas de IA direcionadas)
func (cj *CenaJogo) ColideComTipo(eu *geometria.Retangulo, tipoDesejado string) bool {
	for _, e := range cj.GetEntidades() {
		if e.GetTipo() == tipoDesejado {
			if corpoEntidade := e.GetComponente(componentes.CORPO.String()); corpoEntidade != nil {
				if eu.Colide(corpoEntidade.(*geometria.Retangulo)) {
					return true
				}
			}
		}
	}
	return false
}

// Métodos auxiliares semanticamente limpos, reaproveitando a função genérica
func (cj *CenaJogo) ColideComBarreiras(eu *geometria.Retangulo) bool {
	return cj.ColideComTipo(eu, entidades.PAREDE.String())
}

func (cj *CenaJogo) ColideComJogador(eu *geometria.Retangulo) bool {
	return cj.ColideComTipo(eu, entidades.JOGADOR.String())
}

func (cj *CenaJogo) ColideComBot(eu *geometria.Retangulo) bool {
	return cj.ColideComTipo(eu, entidades.BOT.String())
}
