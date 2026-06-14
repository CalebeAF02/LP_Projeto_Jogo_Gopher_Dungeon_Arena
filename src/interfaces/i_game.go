package interfaces

type IGame interface {
	GetCena() ICena
	GetCenaJogo() ICenaJogo
	SetCena(cena ICena)
	IniciarJogo()
	Pausar()
	Voltar()
	Sair()
	MudarTelaMenuIniciar()
	MudarTelaProgresso()
}
