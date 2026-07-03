package nivel

import (
	"encoding/json"
	"fmt"
	"os"
)

type Progresso struct {
	NivelCorrente int `json:"nivel_corrente"`
}

func CarregarProgresso() Progresso {
	data, err := os.ReadFile("progresso.json")
	if err != nil {
		fmt.Println("Erro ao ler arquivo:", err)
		return Progresso{}
	}

	var progresso Progresso
	err = json.Unmarshal(data, &progresso)
	if err != nil {
		fmt.Println("Erro ao converter JSON:", err)
		return Progresso{}
	}

	return progresso
}

func SalvarProgresso(progresso Progresso) error {
	data, err := json.MarshalIndent(progresso, "", "  ")
	if err != nil {
		return fmt.Errorf("erro ao converter para JSON: %w", err)
	}

	// Escreve no arquivo
	err = os.WriteFile("progresso.json", data, 0644)
	if err != nil {
		return fmt.Errorf("erro ao salvar arquivo: %w", err)
	}

	return nil
}
