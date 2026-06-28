package nivel

type ObjetoSimples struct {
	Tipo string `json:"tipo"`
	X    int    `json:"x"`
	Y    int    `json:"y"`
}

type ObjetoConectado struct {
	Tipo string `json:"tipo"`
	X1   int    `json:"x1"`
	Y1   int    `json:"y1"`

	X2 int `json:"x2"`
	Y2 int `json:"y2"`
}
