package endgame

type Board []byte

type Piece struct {
     Position byte
     Type byte
}

type ChessPiece interface {
     ValidMoves(board Board) []byte
}

func (piece Piece) ValidMoves(board Board) []byte {
     return make([]byte, 0)
}

func Foo(a, b rune) rune {
     return a + b
}

