# Tetris em Go 🎮

Bem-vindo ao **Tetris em Go**! Este é um projeto open-source que recria o clássico jogo Tetris usando a linguagem de programação Go (Golang). Com uma interface baseada em texto no terminal, o jogo traz todas as funcionalidades que você espera: tabuleiro 10x20, sete peças icônicas, rotação, pontuação e níveis de dificuldade crescentes. Desenvolvido para ser simples, eficiente e divertido, este Tetris é perfeito para entusiastas de programação e fãs de jogos retrô.

O objetivo deste repositório é compartilhar uma implementação funcional de Tetris, demonstrando o poder de Go para criar jogos leves e performáticos. Se você é iniciante em Go ou quer explorar desenvolvimento de jogos, este projeto é um ótimo ponto de partida. Vamos aprender como rodar o jogo e mergulhar nos detalhes!

---

## Screenshots

![Tetris](pictures/01.png)
![Tetris](pictures/02.png)
![Tetris](pictures/03.png)

--- 

## Características

- **Tabuleiro**: Grade padrão de 10 colunas por 20 linhas.
- **Peças**: Inclui as sete formas clássicas (I, J, L, O, S, T, Z) com rotações em 90 graus.
- **Mecânicas**: Queda automática das peças, movimento lateral, rotação e limpeza de linhas completas.
- **Pontuação**: 40 pontos por uma linha, 100 por duas, 300 por três, 1200 por quatro, com aumento de velocidade por nível.
- **Interface**: UI colorida no terminal, centralizada, com contorno e informações (score, nível, próxima peça).
- **Controles**:
  - Setas esquerda/direita: mover peça.
  - Seta para baixo: acelerar queda.
  - Seta para cima ou espaço: rotacionar peça.
  - 'P' ou 'Esc': pausar/retomar.
  - 'R': reiniciar após "Game Over".
  - 'Q' ou 'Ctrl+C': sair.

---

## Pré-requisitos

Para rodar o jogo, você precisará de:
- **Go**: Versão 1.16 ou superior. [Baixe aqui](https://golang.org/dl/).
- **Git**: Para clonar o repositório. [Instale aqui](https://git-scm.com/downloads).
- **Terminal**: Qualquer terminal compatível (Linux, macOS, Windows com WSL ou PowerShell).

Nenhuma experiência prévia em Go é necessária — vamos guiá-lo passo a passo!

---

## Como Rodar o Jogo (Guia para Iniciantes)

Siga estas instruções para jogar Tetris no seu computador. Não se preocupe, é simples!

### 1. Clone o Repositório
Abra seu terminal e clone este repositório:

```bash
git clone https://github.com/chmenegatti/tetris.git
cd tetris
```

Isso baixa o código para uma pasta chamada `tetris` e entra nela.

### 2. Instale a Dependência
O jogo usa a biblioteca `termbox-go` (via `tcell`) para renderizar no terminal. Instale-a com:

```bash
go get github.com/gdamore/tcell/v2/termbox
```

Se você estiver usando Go Modules (recomendado), esse comando atualizará o arquivo `go.mod` automaticamente.

### 3. Compile o Jogo
Agora, compile o código em um executável:

```bash
go build
```

Isso cria um arquivo chamado `tetris` (ou `tetris.exe` no Windows) na pasta atual.

### 4. Execute o Jogo
Rode o jogo com:

```bash
./tetris
```

No Windows, use:

```bash
tetris.exe
```

Alternativamente, você pode pular a compilação e rodar diretamente:

```bash
go run .
```

### 5. Jogue!
- Use as setas para mover e rotacionar as peças.
- Pressione 'P' ou 'Esc' para pausar, 'R' para reiniciar após "Game Over", e 'Q' para sair.
- Tente alcançar a maior pontuação possível!

---

## Estrutura do Código

O projeto é organizado em cinco arquivos principais, cada um com uma responsabilidade clara:

- **`main.go`**: Ponto de entrada. Gerencia o loop principal, eventos de teclado e o ticker para a queda das peças.
  - Exemplo: O loop usa uma goroutine para capturar eventos enquanto o ticker atualiza o jogo.
- **`board.go`**: Define o tabuleiro 10x20 como uma matriz `[20][10]int`. Inclui funções para verificar colisões (`CanPlacePiece`), fixar peças (`PlacePiece`) e limpar linhas (`ClearLines`).
- **`piece.go`**: Contém a definição das sete peças como matrizes 4x4, com rotações e cores. A função `NewPiece` gera peças aleatórias.
- **`game.go`**: Gerencia a lógica do jogo: estados (`Playing`, `Paused`, `GameOver`), movimento, rotação e pontuação.
  - Exemplo: `Update` controla a descida automática e verifica "Game Over".
- **`render.go`**: Renderiza o tabuleiro, peças e informações no terminal usando `termbox-go`. Inclui ajustes para proporção e cores.

---

## Detalhes Técnicos

Aqui está uma visão mais profunda de como o jogo funciona:

### Tabuleiro (`board.go`)
O tabuleiro é uma matriz bidimensional onde `0` representa células vazias e valores de 1 a 7 indicam peças fixadas com suas respectivas cores. A função `ClearLines` é um destaque:

```go
func (b *Board) ClearLines() int {
    linesCleared := 0
    for i := 0; i < BoardHeight; i++ {
        full := true
        for j := 0; j < BoardWidth; j++ {
            if b.Grid[i][j] == 0 {
                full = false
                break
            }
        }
        if full {
            linesCleared++
            for k := i; k > 0; k-- {
                b.Grid[k] = b.Grid[k-1]
            }
            b.Grid[0] = [BoardWidth]int{}
        }
    }
    return linesCleared
}
```

### Peças (`piece.go`)
As peças são definidas como slices de matrizes 4x4, com cada rotação pré-calculada. Por exemplo, a peça "I":

```go
{
    Shape: [][4][4]int{
        {{0, 0, 0, 0}, {1, 1, 1, 1}, {0, 0, 0, 0}, {0, 0, 0, 0}}, // Horizontal
        {{0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}}, // Vertical
    },
    Color: 1, // Ciano
},
```

### Lógica do Jogo (`game.go`)
O método `Update` é o coração do jogo, controlando a descida e a transição para a próxima peça:

```go
func (g *Game) Update() {
    if !g.Board.CanPlacePiece(g.Current, g.Current.X, g.Current.Y+1) {
        g.Board.PlacePiece(g.Current, g.Current.X, g.Current.Y)
        lines := g.Board.ClearLines()
        g.updateScore(lines)
        g.Current = g.Next
        g.Next = NewPiece()
        g.Current.X, g.Current.Y = BoardWidth/2-2, 0
        if !g.Board.CanPlacePiece(g.Current, g.Current.X, g.Current.Y) {
            g.State = GameOver
        }
    } else {
        g.Current.Y++
    }
}
```

### Pontuação (`game.go`)
A pontuação segue o padrão clássico, com velocidade ajustada por nível:

```go
func (g *Game) updateScore(lines int) {
    switch lines {
    case 1: g.Score += 40 * g.Level
    case 2: g.Score += 100 * g.Level
    case 3: g.Score += 300 * g.Level
    case 4: g.Score += 1200 * g.Level
    }
    g.Level = g.Score/1000 + 1
    g.Speed = time.Duration(500-(g.Level-1)*50) * time.Millisecond
}
```

### Renderização (`render.go`)
A renderização usa `termbox-go` para desenhar o tabuleiro com duas células por bloco, garantindo proporção:

```go
for y := 0; y < BoardHeight; y++ {
    for x := 0; x < BoardWidth; x++ {
        if g.Board.Grid[y][x] != 0 {
            color := getColor(g.Board.Grid[y][x])
            termbox.SetCell(boardXOffset+x*2, boardYOffset+y, '█', color, termbox.ColorBlack)
            termbox.SetCell(boardXOffset+x*2+1, boardYOffset+y, '█', color, termbox.ColorBlack)
        }
    }
}
```

---

## Contribuindo

Quer melhorar o Tetris? Sinta-se à vontade para contribuir! Aqui está como:

1. Faça um fork do repositório.
2. Crie uma branch para sua feature (`git checkout -b minha-feature`).
3. Faça suas alterações e commit (`git commit -m "Adicionei tal coisa"`).
4. Envie para seu fork (`git push origin minha-feature`).
5. Abra um Pull Request descrevendo suas mudanças.

Sugestões de melhorias:
- Adicionar um menu inicial.
- Implementar um ranking de pontuação.
- Suporte a temas visuais no terminal.

---

## Problemas Comuns e Soluções

- **Erro "module not found"**: Certifique-se de rodar `go get github.com/gdamore/tcell/v2/termbox`.
- **Jogo não abre**: Verifique se o Go está instalado (`go version`) e se você compilou com `go build`.
- **Cores estranhas**: Depende do terminal; tente em um terminal moderno como WSL ou iTerm2.

---

## Licença

Este projeto é licenciado sob a [MIT License](LICENSE). Use, modifique e distribua à vontade!

---

## Agradecimentos

Obrigado por conferir este Tetris em Go! Este projeto foi uma jornada divertida de aprendizado e nostalgia. Se você gostou, deixe uma ⭐ no repositório e compartilhe suas ideias. Divirta-se jogando e programando!

## Licença

Este projeto é licenciado sob a licença MIT - veja o arquivo [LICENSE](LICENSE) para mais detalhes.
