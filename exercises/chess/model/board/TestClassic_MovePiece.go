package board

import (
	"testing"

	"github.com/cdebriey/4eail40_2020/exercises/chess/model/coord"
)

type mockPiece struct {
}

func (r mockPiece) String() string {
	panic("implement me")
}

func (r mockPiece) Moves() map[coord.Coordinates]bool {
	panic("implement me")
}

func TestClassic_MovePiece(t *testing.T) {
	class := Classic{}
	piece := mockPiece{}
	piece1 := mockPiece{}
	coordin := coord.NewCartesian(0, 0)
	x, _ := coordin.Coord(0)
	y, _ := coordin.Coord(1)
	class[x][y] = piece
	coordin1 := coord.NewCartesian(7, 0)
	x1, _ := coordin1.Coord(0)
	y1, _ := coordin1.Coord(1)
	class[x1][y1] = pi1

	type args struct {
		from coord.Coordinates
		to   coord.Coordinates
	}
	tests := []struct {
		name    string
		c       Classic
		args    args
		wantErr bool // true if there is an error, false if there is not.
	}{
		{
			"no piece at from",
			class,
			args{from: coord.NewCartesian(3, 5), to: coordin},
			true,
		},
		{
			"occupied at to",
			class,
			args{from: coordin1, to: coordin},
			true,
		},
		{
			"move accepted",
			class,
			args{from: coordin, to: coord.NewCartesian(4, 4)},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.MovePiece(tt.args.from, tt.args.to); (err != nil) != tt.wantErr {
				t.Errorf("Classic.MovePiece() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
