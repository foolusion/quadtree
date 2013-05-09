package quadtree

import "testing"

var containsPointsTests = []struct {
	box AABB
	p   XY
	exp bool
}{
	{AABB{XY{0, 0}, XY{1, 1}}, XY{0, 0}, true},
}

func TestAABBContainsPoint(t *testing.T) {
	for i, v := range containsPointsTests {
		out := v.box.ContainsPoint(v.p)
		if out != v.exp {
			t.Errorf("%d. %v with input = %v: output %v expected %v", i, v.box, v.p, out, v.exp)
		}
	}
}

var intersectsAABBTests = []struct {
	a, b AABB
	exp  bool
}{
	{AABB{XY{0, 0}, XY{1, 1}}, AABB{XY{0, 0}, XY{2, 1}}, true},
}

func TestAABBIntersctsAABB(t *testing.T) {
	for i, v := range intersectsAABBTests {
		out := v.a.IntersectsAABB(v.b)
		if out != v.exp {
			t.Errorf("%d. %v with inpute = %v: output %v expected %v", i, v.a, v.b, out, v.exp)
		}
	}
}
