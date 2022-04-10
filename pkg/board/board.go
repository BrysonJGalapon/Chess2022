package board

type Board interface {
	Make(Move) error
	Copy() Board
}

type board struct {
}

func (b *board) Make(m Move) error {
	// TODO
	return nil
}

func (b *board) Copy() Board {
	// TODO
	return &board{}
}

func Standard() Board {
	return &board{}
}
