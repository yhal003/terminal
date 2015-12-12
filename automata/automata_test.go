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
	// AlwaysReject does not accept string 'a'
	c.Assert(x.GetState(),Equals, automata.REJECTED)
}


func (s *AutomataSuite) TestSequenceEmpty1(c *C) {
	a := automata.MakeSequence([]byte{})
	// empty sequence accepts empty string 
	c.Assert(a.GetState(),Equals, automata.ACCEPTED)
}

func (s *AutomataSuite) TestSequenceEmpty2(c *C) {
	a := automata.MakeSequence([]byte{})
	a.Feed('a')
	a.Feed('b')
	// empty sequence does not accept non-empty string 'ab'
	c.Assert(a.GetState(),Equals, automata.REJECTED)
}

func (s *AutomataSuite) TestSequenceSimple1(c *C) {
	a := automata.MakeSequence([]byte{1})
	a.Feed(byte(1))
	// sequence automata accepts string '0x1'
	c.Assert(a.GetState(),Equals, automata.ACCEPTED)
}

func (s *AutomataSuite) TestSequenceSimple2(c *C) {
	a := automata.MakeSequence([]byte{1,3,4})
	a.Feed(byte(1))
	a.Feed(byte(3))
	a.Feed(byte(4))
	// sequence automata accepts string '0x1 0x3 0x4'
	c.Assert(a.GetState(),Equals, automata.ACCEPTED)
	a.Feed(byte(2))
	// sequence automata rejects string '0x1 0x3 0x4 0x2'
	c.Assert(a.GetState(),Equals, automata.REJECTED)
}

func (s *AutomataSuite) TestStarSimple(c *C){
	fa := func() automata.Automata {
		return automata.MakeSequence([]byte{'a'})
	}
	star := automata.MakeStar(fa)
	// a* must start with RUNNING state
	c.Assert(star.GetState(),Equals, automata.RUNNING)
	star.Feed('a')
	// a* must accept 'a'
	c.Assert(star.GetState(),Equals, automata.ACCEPTED)
	star.Feed('a')
	// a* must accept 'aa'
	c.Assert(star.GetState(),Equals, automata.ACCEPTED)
	star.Feed('b')
	// a* must reject 'aab'
	c.Assert(star.GetState(),Equals, automata.REJECTED)
	star.Feed('a')
	// a* must reject 'aaba'
	c.Assert(star.GetState(),Equals, automata.REJECTED)
}

func (s *AutomataSuite) TestStarABC(c *C) {
	fa := func() automata.Automata {
		return automata.MakeSequence([]byte{'a','b','c'})
	}
	star := automata.MakeStar(fa)
	// (abc)* must start with RUNNING state
	c.Assert(star.GetState(),Equals, automata.RUNNING)
	star.Feed('a')
	// (abc)* is running at 'a'
	c.Assert(star.GetState(),Equals, automata.RUNNING)
	star.Feed('b')
	// (abc)* is running at 'ab'
	c.Assert(star.GetState(),Equals, automata.RUNNING)
	star.Feed('c')
	// (abc)* is accepting at 'abc'
	c.Assert(star.GetState(),Equals, automata.ACCEPTED)
	star.Feed('a')
	// (abc)* is running at 'abca'
	c.Assert(star.GetState(),Equals, automata.RUNNING)
	star.Feed('d')
	// (abc)* is rejecting at 'abcad'
	c.Assert(star.GetState(),Equals, automata.REJECTED)
	
}


