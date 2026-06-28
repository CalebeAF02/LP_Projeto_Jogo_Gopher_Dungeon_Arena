package nivel

import (
	"Gopher_Dungeon_Arena/src/componentes/movimentacao"
	"Gopher_Dungeon_Arena/src/entidades/geometria"
	"Gopher_Dungeon_Arena/src/entidades/objeto"
	"Gopher_Dungeon_Arena/src/entidades/outros"
	"Gopher_Dungeon_Arena/src/entidades/personagens"
	"Gopher_Dungeon_Arena/src/enum/cores"
	"Gopher_Dungeon_Arena/src/interfaces"
	"encoding/json"
	"fmt"
	"os"
)

func CarregarNivel(cj interfaces.ICenaJogo) {

	objetos, err := LerObjetos("src/assets/niveis/nivel_01.json")
	if err != nil {
		fmt.Println("Erro:", err)
	}

	var portaisSaida []*objeto.PortalSaida
	var portaisEntrada []*objeto.PortalEntrada

	// Mostrar os objetos carregados
	for _, objRaw := range objetos {

		var t struct {
			Tipo string `json:"tipo"`
		}
		if err := json.Unmarshal(objRaw, &t); err != nil {
			return
		}

		if t.Tipo == "JOGADOR" {

			var obj ObjetoSimples
			if err := json.Unmarshal(objRaw, &obj); err != nil {
				return
			}

			fmt.Printf("Tipo: %s, X: %d, Y: %d\n", obj.Tipo, obj.X, obj.Y)

			j1 := personagens.NovoJogador(cj, "Jogador 1")

			j1.SetPosicao(float64(obj.X), float64(obj.Y))

			j1.SetNivel(1)

			// Times
			t1 := outros.NovoTime(cj, "Vermelhao - Time_Azul", cores.AZUL)

			// Gerenciando
			t1.Adicionnar(j1)

			j1.CarregarPontuacao()
		} else if t.Tipo == "PAREDE" {

			var obj ObjetoSimples
			if err := json.Unmarshal(objRaw, &obj); err != nil {
				return
			}

			objeto.NovaParede(
				cj,
				geometria.NovoPonto(float64(obj.X), float64(obj.Y)),
			)

		} else if t.Tipo == "COMIDA" {

			var obj ObjetoSimples
			if err := json.Unmarshal(objRaw, &obj); err != nil {
				return
			}

			objeto.NovaComida(
				cj,
				geometria.NovoPonto(float64(obj.X), float64(obj.Y)),
			)

		} else if t.Tipo == "SAIDA" {

			var obj ObjetoSimples
			if err := json.Unmarshal(objRaw, &obj); err != nil {
				return
			}

			objeto.NovaSaida(cj, geometria.NovoPonto(float64(obj.X), float64(obj.Y)))

		} else if t.Tipo == "PORTAL_SAIDA" {

			var obj ObjetoSimples
			if err := json.Unmarshal(objRaw, &obj); err != nil {
				return
			}

			portalSaida := objeto.NovoPortalSaida(cj, 0)
			portalSaida.SetPosicao(float64(obj.X), float64(obj.Y))

			//bEntrada1.ConectarSaida(bSaida1)

			portaisSaida = append(portaisSaida, portalSaida)

		} else if t.Tipo == "PORTAL_ENTRADA" {

			var obj ObjetoSimples
			if err := json.Unmarshal(objRaw, &obj); err != nil {
				return
			}

			portalEntrada := objeto.NovoPortalEntrada(cj, 0)
			portalEntrada.SetPosicao(float64(obj.X), float64(obj.Y))

			//bEntrada1.ConectarSaida(bSaida1)

			portaisEntrada = append(portaisEntrada, portalEntrada)

		} else if t.Tipo == "BOT_SIMPLES" {

			var obj ObjetoSimples
			if err := json.Unmarshal(objRaw, &obj); err != nil {
				return
			}

			cj.SpawnarBot(cj, &movimentacao.MovimentadorSimples{}, geometria.NovoPonto(float64(obj.X), float64(obj.Y)))
		} else if t.Tipo == "BOT_VERTICAL" {

			var obj ObjetoSimples
			if err := json.Unmarshal(objRaw, &obj); err != nil {
				return
			}

			cj.SpawnarBot(cj, &movimentacao.MovimentadorVertical{}, geometria.NovoPonto(float64(obj.X), float64(obj.Y)))
		} else if t.Tipo == "BOT_VERTICAL_CONSTANTE" {

			var obj ObjetoSimples
			if err := json.Unmarshal(objRaw, &obj); err != nil {
				return
			}

			cj.SpawnarBot(cj, &movimentacao.MovimentadorVerticalConstante{}, geometria.NovoPonto(float64(obj.X), float64(obj.Y)))
		} else if t.Tipo == "BOT_HORIZONTAL" {

			var obj ObjetoSimples
			if err := json.Unmarshal(objRaw, &obj); err != nil {
				return
			}

			cj.SpawnarBot(cj, &movimentacao.MovimentadorHorizontal{}, geometria.NovoPonto(float64(obj.X), float64(obj.Y)))

		} else if t.Tipo == "PORTAL_CONECTADO" {

			var obj ObjetoConectado
			if err := json.Unmarshal(objRaw, &obj); err != nil {
				return
			}

			print("\n >> Conectando portal :: ", obj.X1, " | ", obj.Y1, " EM ", obj.X2, " | ", obj.Y2)

			// 1. Procurar o Portal de Entrada correspondente na lista (X1, Y1)
			var entradaEncontrada *objeto.PortalEntrada
			for _, pEntrada := range portaisEntrada {
				// Supondo que você tenha métodos GetX() e GetY() ou campos expostos.
				// Se GetPosicao() retornar float64, ajuste conforme sua struct:
				px, py := pEntrada.GetX1(), pEntrada.GetY1()

				if px == float64(obj.X1) && py == float64(obj.Y1) {
					entradaEncontrada = pEntrada
					break
				}
			}

			// 2. Procurar o Portal de Saída correspondente na lista (X2, Y2)
			var saidaEncontrada *objeto.PortalSaida
			for _, pSaida := range portaisSaida {
				px, py := pSaida.GetX1(), pSaida.GetY1()

				if px == float64(obj.X2) && py == float64(obj.Y2) {
					saidaEncontrada = pSaida
					break
				}
			}

			// 3. Se ambos forem encontrados na lista, realiza a conexão física
			if entradaEncontrada != nil && saidaEncontrada != nil {
				entradaEncontrada.ConectarSaida(saidaEncontrada)
				fmt.Printf("\nPortais vinculados com sucesso: Entrada(%d,%d) -> Saída(%d,%d)\n",
					obj.X1, obj.Y1, obj.X2, obj.Y2)
			} else {
				fmt.Printf("\nNão foi possível conectar: Entrada encontrada? %t | Saída encontrada? %t\n",
					entradaEncontrada != nil, saidaEncontrada != nil)
			}

		}

	}
}
func LerObjetos(caminho string) ([]json.RawMessage, error) {
	// Ler o arquivo
	data, err := os.ReadFile(caminho)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler arquivo: %w", err)
	}

	// Converter JSON para slice de objetos
	var rawItems []json.RawMessage
	err = json.Unmarshal(data, &rawItems)
	if err != nil {
		return nil, fmt.Errorf("erro ao converter JSON: %w", err)
	}

	return rawItems, nil
}
