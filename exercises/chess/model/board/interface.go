// Package board will contain types and logic for handling chess board.
package board

import (
	"fmt"
)

// Board is an interface to a chess board.
// It defines methods for handling pieces ont it.
type Board interface {
	fmt.Stringer
	PieaceAt(at coord.ChessCoordinates) piece.Piece
}
