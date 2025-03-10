package main

import "time"

type GameState int

const (
	Starting GameState = iota
	Playing
	Paused
	GameOver
)

type Game struct {
	Board   *Board
	Current *Piece
	Next    *Piece
	Score   int
	Level   int
	Speed   time.Duration
	State   GameState
}

func NewGame() *Game {
	g := &Game{
		Board:   NewBoard(),
		Current: NewPiece(),
		Next:    NewPiece(),
		Score:   0,
		Level:   1,
		Speed:   500 * time.Millisecond,
		State:   Playing, // Começar diretamente no estado Playing
	}
	return g
}

func (g *Game) Update() {
	if g.State != Playing {
		return
	}

	// Verificar se a peça pode descer
	if !g.Board.CanPlacePiece(g.Current, g.Current.X, g.Current.Y+1) {
		// Peça chegou ao fundo ou colidiu
		g.Board.PlacePiece(g.Current, g.Current.X, g.Current.Y)
		lines := g.Board.ClearLines()
		g.updateScore(lines)

		// Trocar para a próxima peça
		g.Current = g.Next
		g.Next = NewPiece()
		g.Current.X = BoardWidth/2 - 2 // Centralizar a nova peça
		g.Current.Y = 0                // Começar no topo

		// Verificar se o jogo acabou
		if !g.Board.CanPlacePiece(g.Current, g.Current.X, g.Current.Y) {
			g.State = GameOver
		}
	} else {
		// Peça desce automaticamente
		g.Current.Y++
	}
}

func (g *Game) MovePiece(dx, dy int) {
	if g.State != Playing {
		return
	}
	newX := g.Current.X + dx
	newY := g.Current.Y + dy
	if g.Board.CanPlacePiece(g.Current, newX, newY) {
		g.Current.X = newX
		g.Current.Y = newY
	}
}

func (g *Game) RotatePiece() {
	if g.State != Playing {
		return
	}
	oldRotation := g.Current.Rotation
	g.Current.Rotate()
	if !g.Board.CanPlacePiece(g.Current, g.Current.X, g.Current.Y) {
		g.Current.Rotation = oldRotation
	}
}

func (g *Game) updateScore(lines int) {
	switch lines {
	case 1:
		g.Score += 40 * g.Level
	case 2:
		g.Score += 100 * g.Level
	case 3:
		g.Score += 300 * g.Level
	case 4:
		g.Score += 1200 * g.Level
	}
	g.Level = g.Score/1000 + 1
	g.Speed = time.Duration(500-(g.Level-1)*50) * time.Millisecond
	if g.Speed < 50*time.Millisecond {
		g.Speed = 50 * time.Millisecond
	}
}
