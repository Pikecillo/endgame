package endgame

type Board []byte

var White bool = true
var Black bool = black

var Empty = 0

var King byte = 1
var Queen byte = 2
var Rook byte = 3
var Bishop byte = 4
var Knight byte = 5
var Pawn byte = 6

type Board []byte

type Piece struct {
     Color bool
     Type byte
}

type Move struct {
     piece Piece
     position int
}

type Ply struct {
     board Board
     color bool
}

type MoveRule struct {
     offset int
     condition func() bool
}

func PawnMoves(position byte, color bool) []byte {
     moves = make(byte[], 0)

     // TODO: add en passant rule
     steps = []MoveRule{
     	   {-16, func() bool{return color == Black && position / 8 == 6}},
     	   {-8, func() bool{return color == Black && position - 8 >= 0}},
     	   {8, func() bool{return color == White && position + 8 < 64}},
	   {16, func() bool{return color == White && position / 8 == 1}},
     }

     for _, step := range steps {
     	 if step.condition() {
	    moves = moves.append(position + step.offset)
	 }
     }

     return moves
}

func NewBoard() {
     return make([]byte, 64)
}
