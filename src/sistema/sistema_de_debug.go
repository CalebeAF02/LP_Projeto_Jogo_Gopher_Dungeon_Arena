package sistema

import (
	"Gopher_Dungeon_Arena/src/ecs"
	"Gopher_Dungeon_Arena/src/interfaces"
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
)

type SistemaDebug struct{}

func (self *SistemaDebug) Atualizar(cj interfaces.ICenaJogo) {
	// Atalho para debugar as entidades no terminal
	self.Input(cj)
}

func (self *SistemaDebug) Input(cj interfaces.ICenaJogo) {
	if ebiten.IsKeyPressed(ebiten.KeyF1) {
		self.ListarPrincipaisEntidades(cj)
	}

	if ebiten.IsKeyPressed(ebiten.KeyF2) {
		self.ListarEntidadesOrdenadas(cj)
	}
}

func (self *SistemaDebug) ListarEntidadesOrdenadas(cj interfaces.ICenaJogo) {

	contComidas := 0
	contParedes := 0
	contBots := 0

	fmt.Println("\n" + strings.Repeat("=", 40))
	fmt.Println("   RELATÓRIO DE ENTIDADES (ORDENADO POR ID)")
	fmt.Println(strings.Repeat("=", 40))

	entidades := cj.GetEntidades()
	if len(entidades) == 0 {
		fmt.Println("O mundo está vazio.")
		return
	}

	// 1. Extrair todas as chaves (IDs) do mapa
	ids := make([]int, 0, len(entidades))
	for id := range entidades {
		ids = append(ids, int(id))
	}

	// 2. Ordenar os IDs numericamente
	sort.Ints(ids)

	// 3. Percorrer os IDs já ordenados
	for _, idInt := range ids {
		id := ecs.EntidadeID(idInt)
		entidade := entidades[id]

		v := reflect.ValueOf(entidade)
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		t := v.Type()

		fmt.Printf("\n>>> ID: %d | CLASSE: %s\n", id, t.Name())

		if t.Name() == "Bot" {
			contBots++
		} else if t.Name() == "Parede" {
			contParedes++
		} else if t.Name() == "Comida" {
			contComidas++
		}

		// Listar Atributos
		for i := 0; i < v.NumField(); i++ {
			campoNome := t.Field(i).Name
			campoValor := v.Field(i)

			// Pula campos internos ou privados (que começam com letra minúscula) se desejar
			if t.Field(i).PkgPath != "" {
				continue
			}

			fmt.Printf("    %-18s : %v\n", campoNome, campoValor)
		}
	}
	fmt.Println("\n" + strings.Repeat("=", 40))
	fmt.Println("O Mundo possui : |", contParedes, " Paredes| , |", contBots, " Bots| e |", contComidas, " Comidas|")

	fmt.Println("\n" + strings.Repeat("=", 40))

}

func (self *SistemaDebug) ListarPrincipaisEntidades(cj interfaces.ICenaJogo) {
	entidades := cj.GetEntidades()
	if len(entidades) == 0 {
		fmt.Println("O mundo está vazio.")
		return
	}

	totalParedes := 0
	var idsOutros []int

	for id, e := range entidades {
		// Importante: verifique se seu GetTipo() retorna exatamente "PAREDE"
		// ou use a constante entidades.PAREDE.String()
		if e.GetTipo() == "PAREDE" {
			totalParedes++
		} else {
			idsOutros = append(idsOutros, int(id))
		}
	}

	sort.Ints(idsOutros)

	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Printf("   RELATÓRIO COMPLETO DE ATRIBUTOS (Paredes: %d)\n", totalParedes)
	fmt.Println(strings.Repeat("=", 60))

	for _, idInt := range idsOutros {
		id := ecs.EntidadeID(idInt)
		entidade := entidades[id]

		v := reflect.ValueOf(entidade)
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		t := v.Type()

		fmt.Printf("\n[ID: %d] CLASSE: %s\n", id, strings.ToUpper(t.Name()))
		fmt.Println(strings.Repeat("-", 30))

		for i := 0; i < v.NumField(); i++ {
			campoNome := t.Field(i).Name
			campoValor := v.Field(i)

			// Esta parte garante que apresente até valores complexos (como Ponto ou Retângulo)
			// de forma legível no terminal
			valorFormatado := fmt.Sprintf("%v", campoValor)

			// Se o valor for uma Interface ou Ponteiro interno, tentamos extrair o conteúdo real
			if campoValor.Kind() == reflect.Ptr && !campoValor.IsNil() {
				valorFormatado = fmt.Sprintf("%v", campoValor.Elem())
			}

			fmt.Printf("   %-20s : %s\n", campoNome, valorFormatado)
		}
	}

	fmt.Println("\n" + strings.Repeat("-", 60))
	fmt.Printf(" -> %d PAREDES FORAM AGRUPADAS PARA LIMPEZA VISUAL\n", totalParedes)
	fmt.Println(strings.Repeat("=", 60))
}
