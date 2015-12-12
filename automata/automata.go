package automata

const (
	RUNNING = iota
	REJECTED
	ACCEPTED
)

type Automata interface {
	GetState() int
	Feed(b byte)
}


// simple automata that always rejects a string
type AlwaysReject struct {
}

func (r *AlwaysReject) GetState() int {
	return REJECTED
}

func (r *AlwaysReject) Feed(b byte) {
	return
}

// automata that accepts particular byte string

type Sequence struct {
	bs []byte
	pos int
	state int
}

func MakeSequence(bs []byte) *Sequence {
	if (len(bs) > 0) {
		return &Sequence{bs,0,RUNNING}
	} else {
		return &Sequence{bs,0,ACCEPTED}
	}
}

func (a *Sequence) GetState() int {
	return a.state
}

func (a *Sequence) Feed(b byte) {
	if (a.state != RUNNING) {
		a.state = REJECTED
		return
	}
	if (a.bs[a.pos] == b) {
		a.pos++
	} else {
		a.state = REJECTED
		return
	}
	if (a.pos == len(a.bs)) {
		a.state = ACCEPTED
		return
	}
}

// "star" operator

type Star struct {
	fa func() Automata 
	a Automata  
} 

func MakeStar(fa func() Automata) Star{
	return Star{fa,fa()}
}

func (s *Star) GetState() int{
	return s.a.GetState()
}

func (s *Star) Feed(b byte) {
	if (s.a.GetState() == ACCEPTED) {
		s.a = s.fa()
	} 
	s.a.Feed(b)
}
