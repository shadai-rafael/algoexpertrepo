package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDepthFirstSearch(t *testing.T) {

	var k Node
	k.Name = "K"

	var j Node
	j.Name = "J"

	var i Node
	i.Name = "I"

	var h Node
	h.Name = "H"

	var g Node
	g.Name = "G"
	g.Children = append(g.Children, &k)

	var f Node
	f.Name = "F"
	f.Children = append(f.Children, &i, &j)

	var e Node
	e.Name = "E"

	var d Node
	d.Name = "D"
	d.Children = append(d.Children, &g, &h)

	var c Node
	c.Name = "C"

	var b Node
	b.Name = "B"
	b.Children = append(b.Children, &e, &f)

	var a Node
	a.Name = "A"
	a.Children = append(a.Children, &b, &c, &d)

	var array []string
	array = a.DepthFirstSearch(array)

	resultExpected := []string{"A", "B", "E", "F", "I", "J", "C", "D", "G", "K", "H"}

	for i, v := range resultExpected {
		assert.Equal(t, v, array[i])
	}

}
