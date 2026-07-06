package sistema

import (
	"Gopher_Dungeon_Arena/src/componentes"
	"Gopher_Dungeon_Arena/src/ecs"
	"Gopher_Dungeon_Arena/src/entidades/funcionalidades"
	"Gopher_Dungeon_Arena/src/entidades/geometria"
	"Gopher_Dungeon_Arena/src/enum/entidades"
	"Gopher_Dungeon_Arena/src/interfaces"
)

type SistemaColisao struct {
	cenaJogo interfaces.ICenaJogo
}

func (self *SistemaColisao) SetCenaJogo(cj interfaces.ICenaJogo) {
	self.cenaJogo = cj
}

func (self *SistemaColisao) CriarRespostaColisao(status bool, tipo string, subTipo string) *ecs.RespostaColisao {
	nRespostaColisao := ecs.RespostaColisao{Status: status, Tipo: tipo, SubTipo: subTipo}

	return &nRespostaColisao
}

func (self *SistemaColisao) EsseTipoColide(tipo string) bool {

	switch tipo {
	case entidades.PAREDE.String():
		return true
	case entidades.JOGADOR.String():
		return true
	case entidades.BOT.String():
		return true
	case entidades.PORTAL_ENTRADA.String():
		return true
	case entidades.PORTAL_SAIDA.String():
		return true
	case entidades.COMIDA.String():
		return true
	case entidades.SAIDA.String():
		return true
	}

	return false
}

func (self *SistemaColisao) VaiColidir(origem string, origemEntidade ecs.Entidade, meuCorpoAtual *geometria.Retangulo, proximoCorpo *geometria.Retangulo) *ecs.RespostaColisao {
	for _, entidadeColidida := range self.cenaJogo.GetEntidades() {
		colididoTipo := entidadeColidida.GetTipo()
		if self.EsseTipoColide(colididoTipo) {

			estaVivo := true

			if colididoTipo == entidades.JOGADOR.String() || colididoTipo == entidades.BOT.String() {
				vidaComp := entidadeColidida.GetComponente(componentes.VIDA.String())
				vida := vidaComp.(*componentes.Vida)
				estaVivo = vida.Status
			}

			if !estaVivo {
				continue
			}

			if corpoEntidade := entidadeColidida.GetComponente(componentes.CORPO.String()); corpoEntidade != nil {
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

					if funcionalidades.Simetria(origem, colididoTipo, entidades.JOGADOR.String(), entidades.BOT.String()) {
						if origem == entidades.BOT.String() && colididoTipo == entidades.JOGADOR.String() {
							funcionalidades.CombateJogadorBot(entidadeColidida, origemEntidade)
						} else if origem == entidades.JOGADOR.String() && colididoTipo == entidades.BOT.String() {
							funcionalidades.CombateJogadorBot(origemEntidade, entidadeColidida)
						}

					} else if funcionalidades.Simetria(origem, colididoTipo, entidades.JOGADOR.String(), entidades.COMIDA.String()) {

						if origem == entidades.COMIDA.String() && colididoTipo == entidades.JOGADOR.String() {
							funcionalidades.EncherBucho(entidadeColidida, origemEntidade)
						} else if origem == entidades.JOGADOR.String() && colididoTipo == entidades.COMIDA.String() {
							funcionalidades.EncherBucho(origemEntidade, entidadeColidida)
						}

					} else if funcionalidades.Simetria(origem, colididoTipo, entidades.BOT.String(), entidades.PORTAL_ENTRADA.String()) {

						//	fmt.Printf("Bot bateu no portal de Entrada !!!1 #slc\n")

						if origem == entidades.BOT.String() && colididoTipo == entidades.PORTAL_ENTRADA.String() {
							funcionalidades.TeleTransporta(entidadeColidida, origemEntidade)
						} else if origem == entidades.PORTAL_ENTRADA.String() && colididoTipo == entidades.BOT.String() {
							funcionalidades.TeleTransporta(origemEntidade, entidadeColidida)
						}

					} else if funcionalidades.Simetria(origem, colididoTipo, entidades.JOGADOR.String(), entidades.SAIDA.String()) {

						if origem == entidades.JOGADOR.String() {
							funcionalidades.ConcluirPartida(origemEntidade)
						} else if colididoTipo == entidades.JOGADOR.String() {
							funcionalidades.ConcluirPartida(entidadeColidida)
						}

					}

					if colididoTipo == entidades.BOT.String() {
						if sub_tipo := entidadeColidida.GetComponente(componentes.SUB_TIPO.String()); sub_tipo != nil {

							sub_tipo_valor := sub_tipo.(*componentes.SubTipo)
							return self.CriarRespostaColisao(true, colididoTipo, sub_tipo_valor.Valor)
						} else {
							return self.CriarRespostaColisao(true, colididoTipo, "")
						}
					} else {
						return self.CriarRespostaColisao(true, colididoTipo, "")

					}
				}
			}
		}
	}
	return self.CriarRespostaColisao(false, "", "")
}

// ColideComTipo isola uma busca específica (útil para o Spawn ou lógicas de IA direcionadas)
func (self *SistemaColisao) ColideComTipo(eu *geometria.Retangulo, tipoDesejado string) bool {
	for _, e := range self.cenaJogo.GetEntidades() {
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
func (self *SistemaColisao) ColideComBarreiras(eu *geometria.Retangulo) bool {
	return self.ColideComTipo(eu, entidades.PAREDE.String())
}

func (self *SistemaColisao) ColideComJogador(eu *geometria.Retangulo) bool {
	return self.ColideComTipo(eu, entidades.JOGADOR.String())
}

func (self *SistemaColisao) ColideComBot(eu *geometria.Retangulo) bool {
	return self.ColideComTipo(eu, entidades.BOT.String())
}
