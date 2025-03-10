package main

const (
	BoardWidth  = 10
	BoardHeight = 20
)

type Board struct {
	// The board's contents
	Grid [BoardHeight][BoardWidth]int
}

func NewBoard() *Board {
	return &Board{}
}

func (b *Board) CanPlacePiece(p *Piece, x, y int) bool {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if p.Shape[p.Rotation][i][j] == 1 {
				// Check if the piece is out of bounds
				newX := x + j
				newY := y + i
				if newX < 0 || newX >= BoardWidth || newY >= BoardHeight || (newY >= 0 && b.Grid[newY][newX] != 0) {
					return false
				}
			}
		}
	}
	return true
}

func (b *Board) PlacePiece(p *Piece, x, y int) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if p.Shape[p.Rotation][i][j] == 1 {
				b.Grid[y+i][x+j] = p.Color
			}
		}
	}
}

// Remove linhas completas e retorna o número de linhas removidas
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
