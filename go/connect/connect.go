package connect

import (
	"errors"
)

// identifiers
const (
	OID = 'O'
	XID = 'X'
)

// neighbours coordinates
var neighbours [6][2]int = [6][2]int{
	{-1, 0},
	{-1, 1},
	{0, 1},
	{1, 0},
	{1, -1},
	{0, -1},
}

// game structure
type game struct {
	lines      []string
	numLines   int
	numColumns int
	visited    map[[2]int]bool
}

func (g *game) getNeighbours(point [2]int) [][2]int {
	var adj [][2]int
	for _, n := range neighbours {
		line, col := point[0], point[1]
		line += n[0]
		col += n[1]

		if line >= 0 && line < len(g.lines) && col >= 0 && col < len(g.lines[0]) {
			adj = append(adj, [2]int{line, col})
		}
	}

	return adj
}

// canWin checks if O can make through top to bottom, or if X can make through left to right
func (g *game) canWin(point [2]int) bool {
	line, col := point[0], point[1]
	player := g.lines[line][col]

	if (player == OID && line == g.numLines-1) || (player == XID && col == len(g.lines[0])-1) {
		return true
	}

	g.visited[point] = true

	nb := g.getNeighbours(point)
	for _, n := range nb {
		if g.lines[n[0]][n[1]] == player {
			if _, ok := g.visited[n]; ok {
				continue
			}

			if g.canWin(n) {
				return true
			}
		}
	}

	return false
}

// oCanWin checks if O can win the game
func (g *game) oCanWin() bool {
	for col := range g.lines[0] {
		if g.lines[0][col] == OID && g.canWin([2]int{0, col}) {
			return true
		}
	}

	return false
}

// xCanWin checks if X can win the game
func (g *game) xCanWin() bool {
	for line, cols := range g.lines {
		if cols[0] == XID && g.canWin([2]int{line, 0}) {
			return true
		}
	}

	return false
}

// ResultOf return the winner or "" if there is no winner
func ResultOf(board []string) (string, error) {
	if len(board) == 0 {
		return "", errors.New("empty board")
	}

	g := &game{
		lines:      board,
		numLines:   len(board),
		numColumns: len(board[0]),
	}

	g.visited = make(map[[2]int]bool)
	if g.oCanWin() {
		return string(OID), nil
	}

	g.visited = make(map[[2]int]bool)
	if g.xCanWin() {
		return string(XID), nil
	}

	return "", nil
}
