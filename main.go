package main

import (
	"time"

	"github.com/gdamore/tcell/v2/termbox"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	game := NewGame()
	game.Run()
}

func (g *Game) Run() {
	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	ticker := time.NewTicker(g.Speed)
	defer ticker.Stop()

mainLoop:
	for {
		g.Render() // Renderizar a cada iteração
		select {
		case ev := <-eventQueue:
			if ev.Type == termbox.EventKey {
				switch ev.Key {
				case termbox.KeyEsc:
					if g.State == Playing {
						g.State = Paused
						ticker.Stop()
					} else if g.State == Paused {
						g.State = Playing
						ticker.Reset(g.Speed)
					}
				case termbox.KeyArrowLeft:
					g.MovePiece(-1, 0)
				case termbox.KeyArrowRight:
					g.MovePiece(1, 0)
				case termbox.KeyArrowDown:
					g.MovePiece(0, 1)
				case termbox.KeyArrowUp, termbox.KeySpace:
					g.RotatePiece()
				case termbox.KeyCtrlQ, termbox.KeyCtrlC:
					break mainLoop // Sair do jogo
				case termbox.KeyCtrlP:
					if g.State == Playing {
						g.State = Paused
						ticker.Stop()
					} else if g.State == Paused {
						g.State = Playing
						ticker.Reset(g.Speed)
					}
				default:
					// Verificar tecla 'R' ou 'r' para reiniciar
					if ev.Ch == 'r' || ev.Ch == 'R' {
						if g.State == GameOver {
							*g = *NewGame()       // Resetar para um novo jogo
							ticker.Reset(g.Speed) // Reiniciar o ticker
						}
					}
				}
			}
		case <-ticker.C:
			if g.State == Playing {
				g.Update()
				ticker.Reset(g.Speed)
			}
		}
	}
}
