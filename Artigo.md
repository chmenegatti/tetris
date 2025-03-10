## Tetris em Go: Uma Jornada Completa no Desenvolvimento de Jogos com Golang 🎮

Você já parou para pensar no que torna Tetris tão atemporal? Aquela mistura de simplicidade e desafio, com blocos caindo em um tabuleiro enquanto você tenta alinhá-los sob pressão, é pura magia dos games. Agora, imagine recriar esse clássico usando Go (Golang), uma linguagem moderna conhecida por sua eficiência e simplicidade. Foi exatamente isso que fiz: desenvolvi um Tetris completo em Go, rodando no terminal, com todas as funcionalidades que você espera — e mais algumas surpresas ao longo do caminho.

Neste artigo, vou guiá-lo por essa aventura de programação. Desde o tabuleiro 10x20 até as sete peças icônicas, passando pela lógica de rotação, pontuação e uma interface colorida, compartilho como construí esse jogo, os desafios técnicos que enfrentei e as soluções que encontrei. Go se provou uma escolha perfeita para o projeto, com sua sintaxe limpa, desempenho robusto e suporte a concorrência. Se você é um desenvolvedor curioso ou um entusiasta de jogos, prepare-se para mergulhar nos detalhes técnicos dessa recriação — e talvez se inspirar a criar seu próprio game em Go!

Por que Go? Além de ser uma linguagem que adoro, ela oferece tipagem estática, uma biblioteca padrão poderosa e um garbage collector eficiente, tornando-a ideal para projetos que exigem lógica clara e execução rápida. Tetris, com sua necessidade de manipulação de estados em tempo real, foi o campo de testes perfeito para explorar essas características. Vamos destrinchar como esse clássico ganhou vida em código.

---

### Visão Geral do Projeto

O Tetris que desenvolvi é uma implementação fiel ao original, com todos os elementos que o tornam reconhecível:
- **Tabuleiro**: Uma grade retangular de 10 colunas por 20 linhas, representada como uma matriz bidimensional.
- **Peças**: As sete formas clássicas (I, J, L, O, S, T, Z), cada uma com rotações em 90 graus.
- **Mecânicas**: Queda automática das peças, movimento lateral, rotação, detecção de colisões e limpeza de linhas completas.
- **Pontuação**: Sistema clássico com 40 pontos por uma linha, 100 por duas, 300 por três e 1200 por quatro, com níveis que aumentam a velocidade.
- **Interface**: Uma UI baseada em texto no terminal, centralizada, com contorno, cores distintas para as peças e informações como pontuação e próxima peça.

Para a interface, utilizei a biblioteca `termbox-go`, especificamente a versão compatível em `github.com/gdamore/tcell/v2/termbox`, que permite desenhar caracteres e cores diretamente no terminal. O resultado é um Tetris jogável, nostálgico e funcional, tudo dentro da simplicidade de um console. Vamos explorar cada componente em detalhes.

---

### Desafios Técnicos e Soluções

Desenvolver Tetris em Go trouxe uma série de desafios técnicos interessantes. Aqui está como lidei com eles, com foco em cada parte do código.

#### Construção das Peças (piece.go)
As sete peças do Tetris foram definidas como matrizes 4x4, com rotações pré-calculadas. O desafio foi representar todas as formas e suas variações de maneira eficiente. Usei uma slice de matrizes para armazenar as rotações, e cada peça tem um campo `Color` para a renderização:

```go
type Piece struct {
    Shape    [][4][4]int // Todas as rotações da peça
    Rotation int         // Rotação atual
    Color    int         // Cor para renderização
    X, Y     int         // Posição no tabuleiro
}

var Pieces = []struct {
    Shape [][4][4]int
    Color int
}{
    { // Peça I
        Shape: [][4][4]int{
            {{0, 0, 0, 0}, {1, 1, 1, 1}, {0, 0, 0, 0}, {0, 0, 0, 0}}, // Horizontal
            {{0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}}, // Vertical
        },
        Color: 1, // Ciano
    },
    // Outras peças (J, L, O, S, T, Z) seguem o mesmo padrão
}
```

A função `NewPiece` usa `rand.New` para selecionar uma peça aleatória, aproveitando a biblioteca padrão de Go para aleatoriedade segura. A rotação é gerenciada com um simples incremento modular, mas com verificação de colisão para evitar sobreposições.

#### Lógica do Tabuleiro (board.go)
O tabuleiro é uma matriz `[20][10]int`, onde `0` indica vazio e valores maiores representam cores das peças fixadas. A detecção de colisões foi um desafio crucial:

```go
func (b *Board) CanPlacePiece(p *Piece, x, y int) bool {
    for i := 0; i < 4; i++ {
        for j := 0; j < 4; j++ {
            if p.Shape[p.Rotation][i][j] == 1 {
                newX, newY := x+j, y+i
                if newX < 0 || newX >= BoardWidth || newY >= BoardHeight || (newY >= 0 && b.Grid[newY][newX] != 0) {
                    return false // Fora dos limites ou colisão
                }
            }
        }
    }
    return true
}
```

A função `ClearLines` percorre o tabuleiro, remove linhas cheias e desloca as superiores para baixo, uma tarefa simplificada pela manipulação direta de arrays em Go. A eficiência da linguagem ajudou a manter essas operações rápidas.

#### Mecânica do Jogo (game.go)
A lógica principal reside em `game.go`, onde gerencio o estado do jogo (`Playing`, `Paused`, `GameOver`) e a queda das peças. O método `Update` controla a descida automática:

```go
func (g *Game) Update() {
    if g.State != Playing {
        return
    }
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
        g.Current.Y++ // Peça desce
    }
}
```

A concorrência de Go brilhou aqui, com uma goroutine capturando eventos de teclado enquanto o ticker atualiza o jogo em intervalos definidos por `g.Speed`.

#### Sistema de Pontuação
O sistema de pontuação segue as regras clássicas, com a velocidade aumentando por nível:

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
    if g.Speed < 50*time.Millisecond {
        g.Speed = 50 * time.Millisecond
    }
}
```

Ajustar o `ticker` com `Reset` garantiu que a velocidade refletisse o nível atual, uma solução elegante usando os recursos temporais de Go.

#### Renderização (render.go)
A renderização no terminal foi desafiadora devido à proporção dos caracteres. Usei dois caracteres por célula (`██`) para um tabuleiro mais quadrado e mapeei cores com uma função auxiliar:

```go
getColor := func(color int) termbox.Attribute {
    switch color {
    case 1: return termbox.ColorCyan
    case 2: return termbox.ColorBlue
    case 3: return termbox.ColorYellow
    case 4: return termbox.ColorGreen
    case 5: return termbox.ColorRed
    case 6: return termbox.ColorMagenta
    case 7: return termbox.ColorWhite
    default: return termbox.ColorDefault
    }
}
```

O contorno do tabuleiro (`╔═╗`, etc.) e o fundo preto (`termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)`) melhoraram a estética, aproveitando a simplicidade de `termbox-go`.

#### Controles e Estados (main.go)
Os controles foram mapeados com eventos de tecla, como `KeyArrowLeft` para mover e 'R' para reiniciar após `GameOver`:

```go
case ev.Ch == 'r' || ev.Ch == 'R':
    if g.State == GameOver {
        *g = *NewGame()
        ticker.Reset(g.Speed)
    }
```

A migração para `github.com/gdamore/tcell/v2/termbox` exigiu ajustes na captura de eventos, mas a compatibilidade foi mantida.

---

### Lições Aprendidas

Esse projeto foi uma aula prática em Go. Aprendi a estruturar código modular, separando responsabilidades em arquivos como `board.go`, `piece.go`, `game.go`, `render.go` e `main.go`. A simplicidade de Go facilitou criar funções claras e reutilizáveis, como `CanPlacePiece` e `ClearLines`.

Também aprofundei meu entendimento de concorrência. Usar goroutines para eventos de teclado e um `time.Ticker` para a lógica de jogo me mostrou como Go lida com tarefas assíncronas de forma natural. Trabalhar com o terminal me desafiou a pensar criativamente sobre interfaces visuais, reforçando a importância de adaptar soluções aos recursos disponíveis.

Outro insight valioso foi a depuração iterativa. Resolver problemas como a peça não renderizando após chegar ao fundo ou o reinício falhando me ensinou a testar hipóteses e refinar o código passo a passo — uma habilidade essencial para qualquer desenvolvedor.

---

### Desempenho e Otimizações

Go entregou um desempenho excepcional para esse Tetris. Como linguagem compilada, o jogo roda em um executável leve, com resposta instantânea aos comandos do jogador. A manipulação da matriz do tabuleiro e a renderização no terminal foram fluidas, mesmo em níveis altos onde as peças caem a cada 50ms.

Otimizações incluíram:
- **Proporção do Tabuleiro**: Usar dois caracteres por célula (`x*2` nas coordenadas) corrigiu o visual “esticado” sem sobrecarregar a renderização.
- **Velocidade Dinâmica**: O ajuste do `ticker` com `Reset` evitou recriações desnecessárias, mantendo o uso de CPU mínimo.
- **Fundo Preto**: Mudar para `termbox.ColorBlack` melhorou a estética sem impacto no desempenho, já que a limpeza da tela é uma operação constante.

A ausência de alocações frequentes (graças à reutilização de structs como `Piece`) e a eficiência do garbage collector de Go mantiveram o jogo leve. Para um projeto mais complexo, poderia explorar profiling com `pprof`, mas para Tetris, o desempenho já era mais que suficiente.

---

### Conclusão

Recriar Tetris em Go foi uma experiência gratificante. Desde a definição das peças em `piece.go` até a renderização colorida em `render.go`, cada componente destacou os pontos fortes de Go: simplicidade, desempenho e concorrência. Os desafios — como rotação, colisões e pontuação — foram superados com código claro e soluções práticas, resultando em um jogo que captura a essência do clássico enquanto exibe o poder da linguagem.

Se você é desenvolvedor, experimente criar seu próprio jogo em Go! Tetris é um ótimo ponto de partida, e o código está disponível em [github.com/chmenegatti/tetris](https://github.com/chmenegatti/tetris). Clone o repositório, jogue uma partida e me conte o que achou nos comentários. Que tal compartilhar sua própria experiência com Go ou perguntar algo sobre o projeto? Vamos trocar ideias e inspirar uns aos outros! 🚀

---

### Hashtags
#GoLang #DesenvolvimentoDeJogos #Tetris #Programação #Tecnologia #DesenvolvimentoDeSoftware #OpenSource


