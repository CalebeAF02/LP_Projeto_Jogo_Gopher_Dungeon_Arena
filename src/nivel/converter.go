package nivel

import (
	"Gopher_Dungeon_Arena/src/enum/cores"
	"Gopher_Dungeon_Arena/src/utils"
	"encoding/json"
	"fmt"
	"image"
	"os"
)

func Converter() {

	dados := []interface{}{}

	file, err := os.Open("src/assets/niveis/nivel_01.png")
	if err != nil {
		fmt.Println("Erro ao abrir imagem:", err)
		return
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Erro ao decodificar imagem:", err)
		return
	}

	bounds := img.Bounds()

	setCores := make(map[string]struct{})

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			rx, gx, bx, _ := img.At(x, y).RGBA()
			// Normaliza para 0–255

			r, g, b := uint8(rx>>8), uint8(gx>>8), uint8(bx>>8)

			if r != 0 && g != 0 && b != 0 {
				//print("COR = [", r, " | ", g, " | ", b, "]\n")
			}

			corStr := fmt.Sprintf("COR = [%d | %d | %d]\n", r, g, b)
			setCores[corStr] = struct{}{}

			// Considera preto se todos os canais forem baixos
			if r == 0 && g == 0 && b == 0 {
				dados = append(dados, ObjetoSimples{
					Tipo: "PAREDE",
					X:    x * utils.PAREDE_TAMANHO_MUNDO,
					Y:    y * utils.PAREDE_TAMANHO_MUNDO,
				})
			} else if r == 163 && g == 73 && b == 164 {

			} else if r == 34 && g == 177 && b == 76 {

				//dados = append(dados, ObjetoSimples{"PORTAL_SAIDA", x * utils.PAREDE_TAMANHO_MUNDO, y * utils.PAREDE_TAMANHO_MUNDO})

				print("PORTAL SAIDA :: ", x, " | ", y, "\n")

			} else if r == 237 && g == 28 && b == 36 {

				//dados = append(dados, ObjetoSimples{"PORTAL_ENTRADA", x * utils.PAREDE_TAMANHO_MUNDO, y * utils.PAREDE_TAMANHO_MUNDO})
				print("PORTAL ENTRADA :: ", x, " | ", y, "\n")

			} else if r == 237 && g == 28 && b == 36 {

				//} else if r == 255 && g == 201 && b == 14 {
				//	dados = append(dados, ObjetoSimples{"BOT_SIMPLES", x * utils.PAREDE_TAMANHO_MUNDO, y * utils.PAREDE_TAMANHO_MUNDO})

			} else if cores.IguaisComponentes(cores.MARROM, r, g, b) {
				dados = append(dados, ObjetoSimples{"BOT_HORIZONTAL", x * utils.PAREDE_TAMANHO_MUNDO, y * utils.PAREDE_TAMANHO_MUNDO})
			} else if cores.IguaisComponentes(cores.VERDE_LIMAO, r, g, b) {
				dados = append(dados, ObjetoSimples{"BOT_SIMPLES", x * utils.PAREDE_TAMANHO_MUNDO, y * utils.PAREDE_TAMANHO_MUNDO})
			} else if cores.IguaisComponentes(cores.AMARELO, r, g, b) {
				dados = append(dados, ObjetoSimples{"BOT_VERTICAL", x * utils.PAREDE_TAMANHO_MUNDO, y * utils.PAREDE_TAMANHO_MUNDO})
			} else if cores.IguaisComponentes(cores.AMARELO_ESCURO, r, g, b) {
				dados = append(dados, ObjetoSimples{"BOT_VERTICAL_CONSTANTE", x * utils.PAREDE_TAMANHO_MUNDO, y * utils.PAREDE_TAMANHO_MUNDO})

			} else if r == 255 && g == 89 && b == 150 {

				dados = append(dados, ObjetoSimples{"COMIDA", x * utils.PAREDE_TAMANHO_MUNDO, y * utils.PAREDE_TAMANHO_MUNDO})

			} else if r == 200 && g == 200 && b == 200 {

				dados = append(dados, ObjetoSimples{"SAIDA", x * utils.PAREDE_TAMANHO_MUNDO, y * utils.PAREDE_TAMANHO_MUNDO})

			} else if r == 0 && g == 162 && b == 232 {
				dados = append(dados, ObjetoSimples{"JOGADOR", x * utils.PAREDE_TAMANHO_MUNDO, y * utils.PAREDE_TAMANHO_MUNDO})
			}
		}
	}

	dados = procurarPortais(img, dados)

	// Imprime todas as strings únicas
	print("------------ CORES ----------------\n")
	for s := range setCores {
		print("\t >>> ", s)
	}

	// Converter para JSON com indentação
	jsonBytes, err := json.MarshalIndent(dados, "", "  ")
	if err != nil {
		fmt.Println("Erro ao converter:", err)
		return
	}

	// Salvar em arquivo
	err = os.WriteFile("src/assets/niveis/nivel_01.json", jsonBytes, 0644)
	if err != nil {
		fmt.Println("Erro ao salvar arquivo:", err)
		return
	}

	fmt.Println("Arquivo salvo com sucesso!")

}
