package automata_test

import ("github.com/yhal003/terminal/automata"
	"testing"
	. "gopkg.in/check.v1"
)

func Test(t *testing.T) {TestingT(t)}

type AutomataSuite struct {}
var _ = Suite(&AutomataSuite{})


func (s *AutomataSuite) TestReject1(c *C) {
	x := &automata.AlwaysReject{}
	// AlwaysReject does not accept empty string
	c.Assert(x.GetState(),Equals, automata.REJECTED)
}

func (s *AutomataSuite) TestReject2(c *C) {
	x := &automata.AlwaysReject{}
	x.Feed('a')
	// AlwaysReject goes not accept string 'a'
	c.Assert(x.GetState(),Equals, automata.REJECTED)
}


func TestSequenceEmpty1(t *testing.T) {
	a := automata.MakeSequence([]byte{})
	if (a.GetState() != automata.ACCEPTED) {
		t.Fatalf("empty sequence must init in accepted state")
	}
}

func TestSequenceEmpty2(t *testing.T) {
	a := automata.MakeSequence([]byte{})
	a.Feed('a')
	a.Feed('b')
	if (a.GetState() != automata.REJECTED) {
		t.Fatalf("empty sequence must not accept non-empty string")
	}
}

func TestSequenceSimple1(t *testing.T) {
	a := automata.MakeSequence([]byte{1})
	a.Feed(byte(1))
	if (a.GetState() != automata.ACCEPTED) {
		t.Fatalf("simple sequence must accept byte 1 ")
	}
}

func TestSequenceSimple2(t *testing.T) {
	a := automata.MakeSequence([]byte{1,3,4})
	a.Feed(byte(1))
	a.Feed(byte(3))
	a.Feed(byte(4))
	if (a.GetState() != automata.ACCEPTED) {
		t.Fatalf("simple sequence must accept bytes 1 3 4 ")
	}
	a.Feed(byte(2))
	if (a.GetState() != automata.REJECTED) {
		t.Fatalf("simple sequence reject byte 2 ")
	}
}

func TestStarSimple(t *testing.T){
	fa := func() automata.Automata {
		return automata.MakeSequence([]byte{'a'})
	}
	s := automata.MakeStar(fa)
	if (s.GetState() != automata.RUNNING) {
		t.Fatalf("a* must start from running")
	}
	s.Feed('a')
	if (s.GetState() != automata.ACCEPTED) {
		t.Fatalf("a* must accept  a")
	}
}


