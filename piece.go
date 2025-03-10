package main

import (
	"math/rand"
	"time"
)

type Piece struct {
	// The piece's shape
	Shape    [][4][4]int
	Rotation int
	Color    int
	X, Y     int
}

var Pieces = []struct {
	Shape [][4][4]int
	Color int
}{
	// I
	{
		Shape: [][4][4]int{
			{{0, 0, 0, 0}, {1, 1, 1, 1}, {0, 0, 0, 0}, {0, 0, 0, 0}},
			{{0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}},
		},
		Color: 1,
	},
	// J
	{
		Shape: [][4][4]int{
			{{0, 0, 0, 0}, {1, 0, 0, 0}, {1, 1, 1, 0}, {0, 0, 0, 0}},
			{{0, 0, 0, 0}, {0, 1, 1, 0}, {0, 1, 0, 0}, {0, 1, 0, 0}},
			{{0, 0, 0, 0}, {0, 0, 0, 0}, {1, 1, 1, 0}, {0, 0, 1, 0}},
			{{0, 0, 0, 0}, {0, 1, 0, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}},
		},
		Color: 2,
	},
	// L
	{
		Shape: [][4][4]int{
			{{0, 0, 0, 0}, {0, 0, 1, 0}, {1, 1, 1, 0}, {0, 0, 0, 0}},
			{{0, 0, 0, 0}, {0, 1, 0, 0}, {0, 1, 0, 0}, {0, 1, 1, 0}},
			{{0, 0, 0, 0}, {0, 0, 0, 0}, {1, 1, 1, 0}, {1, 0, 0, 0}},
			{{0, 1, 1, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 0}},
		},
		Color: 3,
	},
	// O
	{
		Shape: [][4][4]int{
			{{0, 0, 0, 0}, {0, 1, 1, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}},
		},
		Color: 4,
	},
	// S
	{
		Shape: [][4][4]int{
			{{0, 0, 0, 0}, {0, 1, 1, 0}, {1, 1, 0, 0}, {0, 0, 0, 0}},
			{{0, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 0}},
		},
		Color: 5,
	},
	// T
	{
		Shape: [][4][4]int{
			{{0, 0, 0, 0}, {0, 1, 0, 0}, {1, 1, 1, 0}, {0, 0, 0, 0}},
			{{0, 1, 0, 0}, {0, 1, 1, 0}, {0, 1, 0, 0}, {0, 0, 0, 0}},
			{{0, 0, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {0, 0, 0, 0}},
			{{0, 0, 1, 0}, {0, 1, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 0}},
		},
		Color: 6,
	},
	// Z
	{
		Shape: [][4][4]int{
			{{0, 0, 0, 0}, {1, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}},
			{{0, 0, 1, 0}, {0, 1, 1, 0}, {0, 1, 0, 0}, {0, 0, 0, 0}},
		},
		Color: 7,
	},
}

func NewPiece() *Piece {
	// Usar uma fonte de rand única para evitar comportamento determinístico
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	pieceType := r.Intn(len(Pieces))
	p := &Piece{
		Shape:    Pieces[pieceType].Shape,
		Color:    Pieces[pieceType].Color,
		X:        BoardWidth/2 - 2, // Centralizar a peça no topo
		Y:        0,
		Rotation: 0,
	}
	return p
}

func (p *Piece) Rotate() {
	p.Rotation = (p.Rotation + 1) % len(p.Shape)
}
