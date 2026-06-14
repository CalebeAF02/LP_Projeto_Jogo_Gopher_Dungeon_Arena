# 1. Gera o recurso do ícone usando o caminho correto
rsrc -ico src\assets\icones\icone.ico -o ico.syso

# 2. Compila todos os arquivos da raiz (.) no executável final sem o terminal preto atrás
go build -ldflags="-H=windowsgui" -o GopherDungeonArena.exe .

Write-Host "Gopher Dungeon Arena atualizado e com ícone no .exe!" -ForegroundColor Green