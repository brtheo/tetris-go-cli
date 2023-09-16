package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Tetrimino struct {
	self Matrix
	x    int
	y    int
}
type Matrix [][]string

const (
	BOARD_ROWS int = 25
	BOARD_COLS     = 10
	CENTER         = BOARD_COLS / 2
)

const (
	X string = "　" //"　"
	R        = "🟥"
	O        = "🟧"
	Y        = "🟨"
	B        = "🟦"
	G        = "🟩"
	P        = "🟪"
	M        = "🟫"
	D        = "🔲"
)

func main() {
	selfP := Matrix{
		{P, P, P},
		{X, P, X},
	}
	selfO := Matrix{
		{O, O, O},
		{O, X, X},
	}
	selfB := Matrix{
		{B, B, B},
		{X, X, B},
	}
	selfR := Matrix{
		{R, R, X},
		{X, R, R},
	}
	selfG := Matrix{
		{X, G, G},
		{G, G, X},
	}
	selfM := Matrix{
		{M},
		{M},
		{M},
		{M},
	}
	selfY := Matrix{
		{Y, Y},
		{Y, Y},
	}

	tetriminos := []Tetrimino{
		{self: selfP},
		{self: selfO},
		{self: selfB},
		{self: selfR},
		{self: selfG},
		{self: selfM},
		{self: selfY},
	}

	board := NewMatrix(25, 10)

	// debugone := ms[len(ms)-1]
	curr := pick(tetriminos)
	reader := bufio.NewReader(os.Stdin)
	for true {
		char, _, _ := reader.ReadRune()
		switch char {
		case 'q':
			curr.x--
		case 'd':
			curr.x++
		case ' ':
			curr.self = curr.self.rotate()
		}

		loop(curr, board)
		curr.fall()
	}

}
func loop(t Tetrimino, board Matrix) {
	fmt.Print("\033[H\033[2J")
	board.merge(t)
	board.output()
	time.Sleep(1 * time.Second / 4)
}

func (t *Tetrimino) fall() {
	t.y++
}

func pick(arr []Tetrimino) Tetrimino {
	size := len(arr)
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	v := r.Intn(size)
	tetri := arr[v]
	mRows, _ := tetri.self.size()
	tetri.x = CENTER - mRows/2
	tetri.y = 0
	return tetri
}

func NewMatrix(rows, cols int) Matrix {
	m := make([][]string, rows)
	for i := range m {
		m[i] = make([]string, cols)
	}
	return m
}

func (m Matrix) size() (int, int) {
	return len(m[0]), len(m)
}

func (m Matrix) output() {
	row, col := m.size()
	for i := 0; i < col; i++ {
		for j := 0; j < row; j++ {
			fmt.Print(m[i][j])
		}
		fmt.Println(" ")
	}
}

func (board Matrix) merge(t Tetrimino) {
	// _, rows := board.size()
	mcols, mrows := t.self.size()
	x := t.x
	y := t.y
	board.fill(D)
	for i := 0; i < mrows; i++ {
		for j := 0; j < mcols; j++ {
			if t.self[i][j] != X && y+i < 25 {
				board[y+i][x+j] = t.self[i][j]
			}
		}
	}
}

func (m Matrix) rotate() Matrix {
	temp := m
	_, tC := (temp).size()
	for k := 0; k < tC; k++ {
		for i, j := 0, len(temp[k])-1; i < j; i, j = i+1, j-1 {
			temp[k][i], temp[k][j] = temp[k][j], temp[k][i]
		}
	}
	return temp.transpose()
}

// func (m Matrix) rotate() {
// 	rows, cols := m.size()

// 	for i := 0; i < rows/2; i++ {
// 		for j := i; j < cols-i-1; j++ {
// 			Xm := m[i][j]
// 			m[i][j] = m[j][rows-1-i]
// 			m[j][rows-1-i] = m[rows-1-i][rows-1-j]
// 			m[cols-1-i][cols-1-j] = m[cols-1-j][i]
// 			m[cols-1-j][i] = Xm
// 		}
// 	}
// }
func (m Matrix) transpose() Matrix {
	tR, tC := m.size()
	nt := NewMatrix(tR, tC)
	for i := 0; i < tR; i++ {
		for j := 0; j < tC; j++ {
			nt[i][j] = m[j][i]
		}
	}
	return nt
}

func (m Matrix) fill(c string) {
	rows, cols := m.size()
	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			m[i][j] = c
		}
	}
}
