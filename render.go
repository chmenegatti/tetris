package main

import (
	"strconv"

	"github.com/gdamore/tcell/v2/termbox"
)

func (g *Game) Render() {
	termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)

	// Obter dimensões do terminal e centralizar
	width, height := termbox.Size()
	visualWidth := BoardWidth * 2
	boardXOffset := (width - visualWidth) / 2
	boardYOffset := (height - BoardHeight) / 2

	// Função auxiliar para mapear cores
	getColor := func(color int) termbox.Attribute {
		switch color {
		case 1:
			return termbox.ColorCyan
		case 2:
			return termbox.ColorBlue
		case 3:
			return termbox.ColorYellow
		case 4:
			return termbox.ColorGreen
		case 5:
			return termbox.ColorRed
		case 6:
			return termbox.ColorMagenta
		case 7:
			return termbox.ColorWhite
		default:
			return termbox.ColorDefault
		}
	}

	// Desenhar contorno do tabuleiro
	for x := boardXOffset - 1; x <= boardXOffset+visualWidth; x++ {
		termbox.SetCell(x, boardYOffset-1, '═', termbox.ColorWhite, termbox.ColorBlack)
		termbox.SetCell(x, boardYOffset+BoardHeight, '═', termbox.ColorWhite, termbox.ColorBlack)
	}
	for y := boardYOffset - 1; y <= boardYOffset+BoardHeight; y++ {
		termbox.SetCell(boardXOffset-1, y, '║', termbox.ColorWhite, termbox.ColorBlack)
		termbox.SetCell(boardXOffset+visualWidth, y, '║', termbox.ColorWhite, termbox.ColorBlack)
	}
	termbox.SetCell(boardXOffset-1, boardYOffset-1, '╔', termbox.ColorWhite, termbox.ColorBlack)
	termbox.SetCell(boardXOffset+visualWidth, boardYOffset-1, '╗', termbox.ColorWhite, termbox.ColorBlack)
	termbox.SetCell(boardXOffset-1, boardYOffset+BoardHeight, '╚', termbox.ColorWhite, termbox.ColorBlack)
	termbox.SetCell(boardXOffset+visualWidth, boardYOffset+BoardHeight, '╝', termbox.ColorWhite, termbox.ColorBlack)

	// Renderizar o tabuleiro
	for y := 0; y < BoardHeight; y++ {
		for x := 0; x < BoardWidth; x++ {
			if g.Board.Grid[y][x] != 0 {
				color := getColor(g.Board.Grid[y][x])
				termbox.SetCell(boardXOffset+x*2, boardYOffset+y, '█', color, termbox.ColorBlack)
				termbox.SetCell(boardXOffset+x*2+1, boardYOffset+y, '█', color, termbox.ColorBlack)
			} else {
				termbox.SetCell(boardXOffset+x*2, boardYOffset+y, ' ', termbox.ColorDefault, termbox.ColorBlack)
				termbox.SetCell(boardXOffset+x*2+1, boardYOffset+y, ' ', termbox.ColorDefault, termbox.ColorBlack)
			}
		}
	}

	// Renderizar a peça atual
	if g.Current != nil {
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				if g.Current.Shape[g.Current.Rotation][i][j] == 1 {
					color := getColor(g.Current.Color)
					termbox.SetCell(boardXOffset+g.Current.X*2+j*2, boardYOffset+g.Current.Y+i, '█', color, termbox.ColorBlack)
					termbox.SetCell(boardXOffset+g.Current.X*2+j*2+1, boardYOffset+g.Current.Y+i, '█', color, termbox.ColorBlack)
				}
			}
		}
	}

	// Renderizar informações
	infoX := boardXOffset + visualWidth + 2
	drawText(infoX, boardYOffset, "Score: "+strconv.Itoa(g.Score))
	drawText(infoX, boardYOffset+1, "Level: "+strconv.Itoa(g.Level))
	drawText(infoX, boardYOffset+3, "Next:")
	if g.Next != nil {
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				if g.Next.Shape[0][i][j] == 1 {
					color := getColor(g.Next.Color)
					termbox.SetCell(infoX+j*2, boardYOffset+4+i, '█', color, termbox.ColorBlack)
					termbox.SetCell(infoX+j*2+1, boardYOffset+4+i, '█', color, termbox.ColorBlack)
				}
			}
		}
	}

	// Renderizar estado do jogo
	switch g.State {
	case Paused:
		drawText(boardXOffset+visualWidth/2-3, boardYOffset+BoardHeight/2, "PAUSED")
	case GameOver:
		drawText(boardXOffset+visualWidth/2-4, boardYOffset+BoardHeight/2, "GAME OVER")
		drawText(boardXOffset+visualWidth/2-6, boardYOffset+BoardHeight/2+1, "Press R to restart")
	}

	termbox.Flush()
}

func drawText(x, y int, text string) {
	for i, ch := range text {
		termbox.SetCell(x+i, y, ch, termbox.ColorWhite, termbox.ColorBlack)
	}
}
