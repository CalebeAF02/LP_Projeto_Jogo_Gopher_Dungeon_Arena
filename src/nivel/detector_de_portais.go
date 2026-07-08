package nivel

import (
	"Gopher_Dungeon_Arena/src/utils"
	"fmt"
	"image"
)

type Point struct {
	X, Y int
}

type Portal struct {
	ColorName string
	Pixels    []Point
	Center    Point
	RGB       [3]uint8
}

// Suas funções de suporte estruturadas
func COR_RGB_8(r int, g int, b int) [3]uint8 {
	return [3]uint8{uint8(r), uint8(g), uint8(b)}
}

func isColorMatch(c1, c2 [3]uint8) bool {
	return c1[0] == c2[0] && c1[1] == c2[1] && c1[2] == c2[2]
}

func getRGB8(img image.Image, x, y int) [3]uint8 {
	r, g, b, _ := img.At(x, y).RGBA()
	return [3]uint8{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8)}
}

func isAnyPortalColor(c [3]uint8, pathColor [3]uint8) bool {
	if isColorMatch(c, pathColor) {
		return false
	}
	// Portal verde
	if isColorMatch(c, COR_RGB_8(34, 177, 76)) {
		return true
	}
	// Portal vermelho
	if isColorMatch(c, COR_RGB_8(237, 28, 36)) {
		return true
	}
	return false
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Verifica se o portal tem caminho cinza colado
func portalTemCaminho(img image.Image, portal Portal, pathColor [3]uint8) bool {
	dx := []int{0, 0, -1, 1}
	dy := []int{-1, 1, 0, 0}

	for i := 0; i < 4; i++ {
		nx, ny := portal.Center.X+dx[i], portal.Center.Y+dy[i]
		c := getRGB8(img, nx, ny)
		if isColorMatch(c, pathColor) {
			return true
		}
	}
	return false
}

func procurarPortais(img image.Image, dados []interface{}) []interface{} {

	colorPath := COR_RGB_8(127, 127, 127)

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	// Matriz de controle para agrupar portais
	visitedPortals := make([][]bool, width)
	for i := range visitedPortals {
		visitedPortals[i] = make([]bool, height)
	}

	var portais []Portal

	dx := []int{0, 0, -1, 1}
	dy := []int{-1, 1, 0, 0}

	// 1. PRIMEIRO ENCONTRAR TODOS OS PORTAIS DA IMAGEM
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if visitedPortals[x][y] {
				continue
			}

			c := getRGB8(img, x, y)
			if isAnyPortalColor(c, colorPath) {
				// BFS local para achar todos os pixels desta bolinha de portal
				var pList []Point
				queue := []Point{{X: x, Y: y}}
				visitedPortals[x][y] = true
				sumX, sumY := 0, 0

				for len(queue) > 0 {
					curr := queue[0]
					queue = queue[1:]
					pList = append(pList, curr)
					sumX += curr.X
					sumY += curr.Y

					for i := 0; i < 4; i++ {
						nx, ny := curr.X+dx[i], curr.Y+dy[i]
						if nx >= 0 && nx < width && ny >= 0 && ny < height && !visitedPortals[nx][ny] {
							if isColorMatch(getRGB8(img, nx, ny), c) {
								visitedPortals[nx][ny] = true
								queue = append(queue, Point{X: nx, Y: ny})
							}
						}
					}
				}

				portais = append(portais, Portal{
					Center: Point{X: sumX / len(pList), Y: sumY / len(pList)},
					RGB:    c,
				})
			}
		}
	}

	fmt.Printf("Passo 1: %d portais detectados na malha do labirinto.\n", len(portais))
	fmt.Println("Passo 2: Rastreando caminhos a partir de cada portal...")

	// Para evitar exibir conexões duplicadas (A->B e B->A)
	printedPairs := make(map[string]bool)

	// 2. PROCURAR SE AO REDOR DAS 4 DIREÇÕES EXISTE CINZA E SEGUIR
	for i, portalA := range portais {
		visitedPath := make([][]bool, width)
		for idx := range visitedPath {
			visitedPath[idx] = make([]bool, height)
		}

		if !portalTemCaminho(img, portalA, colorPath) {
			continue // ignora portais sem cinza colado
		}

		// BFS segue a partir do cinza colado
		queue := []Point{}
		for d := 0; d < 4; d++ {
			nx, ny := portalA.Center.X+dx[d], portalA.Center.Y+dy[d]
			if isColorMatch(getRGB8(img, nx, ny), colorPath) {
				queue = append(queue, Point{X: nx, Y: ny})
			}
		}
		// Rastreamento do caminho cinza
		for len(queue) > 0 {
			curr := queue[0]
			queue = queue[1:]

			// Verifica se este ponto cinza encostou nas 4 direções de QUALQUER OUTRO portal (Portal B)
			foundConnection := false
			for j, portalB := range portais {
				if i == j {
					continue
				}

				// Verifica se o pixel atual do caminho está adjacente ao centro/borda do portal B
				if abs(curr.X-portalB.Center.X) <= 2 && abs(curr.Y-portalB.Center.Y) <= 2 {
					pairKey := fmt.Sprintf("%d-%d", min(i, j), max(i, j))
					if !printedPairs[pairKey] {
						printedPairs[pairKey] = true
						fmt.Printf("🔗 Par Conectado Encontrado:\n")
						fmt.Printf("   • Portal Entrada: RGB(%d,%d,%d) em (X: %d, Y: %d)\n", portalA.RGB[0], portalA.RGB[1], portalA.RGB[2], portalA.Center.X, portalA.Center.Y)
						fmt.Printf("   • Portal Saída  : RGB(%d,%d,%d) em (X: %d, Y: %d)\n\n", portalB.RGB[0], portalB.RGB[1], portalB.RGB[2], portalB.Center.X, portalB.Center.Y)

						valor1 := getRGB8(img, portalA.Center.X, portalA.Center.Y)
						valor2 := getRGB8(img, portalB.Center.X, portalB.Center.Y)

						fmt.Printf(" >> RGB(%d,%d,%d)\n", valor1[0], valor1[1], valor1[2])
						fmt.Printf(" >> RGB(%d,%d,%d)\n", valor2[0], valor2[1], valor2[2])

						entradaX := 0
						entradaY := 0

						saidaX := 0
						saidaY := 0

						// Decide o tipo de portal com base na cor
						if isColorMatch(valor1, COR_RGB_8(237, 28, 36)) {
							dados = append(dados, ObjetoSimples{
								Tipo: "PORTAL_ENTRADA",
								X:    portalA.Center.X * utils.PAREDE_TAMANHO_MUNDO,
								Y:    portalA.Center.Y * utils.PAREDE_TAMANHO_MUNDO,
							})

							entradaX = portalA.Center.X * utils.PAREDE_TAMANHO_MUNDO
							entradaY = portalA.Center.Y * utils.PAREDE_TAMANHO_MUNDO

						} else if isColorMatch(valor1, COR_RGB_8(34, 177, 76)) {
							dados = append(dados, ObjetoSimples{
								Tipo: "PORTAL_SAIDA",
								X:    portalA.Center.X * utils.PAREDE_TAMANHO_MUNDO,
								Y:    portalA.Center.Y * utils.PAREDE_TAMANHO_MUNDO,
							})

							saidaX = portalA.Center.X * utils.PAREDE_TAMANHO_MUNDO
							saidaY = portalA.Center.Y * utils.PAREDE_TAMANHO_MUNDO

						}

						if isColorMatch(valor2, COR_RGB_8(237, 28, 36)) {
							dados = append(dados, ObjetoSimples{
								Tipo: "PORTAL_ENTRADA",
								X:    portalB.Center.X * utils.PAREDE_TAMANHO_MUNDO,
								Y:    portalB.Center.Y * utils.PAREDE_TAMANHO_MUNDO,
							})

							entradaX = portalB.Center.X * utils.PAREDE_TAMANHO_MUNDO
							entradaY = portalB.Center.Y * utils.PAREDE_TAMANHO_MUNDO

						} else if isColorMatch(valor2, COR_RGB_8(34, 177, 76)) {
							dados = append(dados, ObjetoSimples{
								Tipo: "PORTAL_SAIDA",
								X:    portalB.Center.X * utils.PAREDE_TAMANHO_MUNDO,
								Y:    portalB.Center.Y * utils.PAREDE_TAMANHO_MUNDO,
							})

							saidaX = portalB.Center.X * utils.PAREDE_TAMANHO_MUNDO
							saidaY = portalB.Center.Y * utils.PAREDE_TAMANHO_MUNDO

						}

						dados = append(dados, ObjetoConectado{
							Tipo: "PORTAL_CONECTADO",
							X1:   entradaX,
							Y1:   entradaY,
							X2:   saidaX,
							Y2:   saidaY,
						})

					}
					foundConnection = true
					break
				}
			}

			if foundConnection {
				break
			}

			// Se não conectou ainda, continua seguindo a linha cinza nas 4 direções retas
			for d := 0; d < 4; d++ {
				nx, ny := curr.X+dx[d], curr.Y+dy[d]
				if nx >= 0 && nx < width && ny >= 0 && ny < height && !visitedPath[nx][ny] {
					if isColorMatch(getRGB8(img, nx, ny), colorPath) {
						visitedPath[nx][ny] = true
						queue = append(queue, Point{X: nx, Y: ny})
					}
				}
			}
		}
	}

	return dados
}

func isPath(c [3]uint8) bool {
	return c[0] == 127 && c[1] == 127 && c[2] == 127
}
