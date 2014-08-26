package quadtree

import "testing"

var containsPointsTests = []struct {
	box *AABB
	p   XY
	exp bool
}{
	{&AABB{XY{0, 0}, XY{1, 1}}, XY{0, 0}, true},
	{&AABB{XY{0, 0}, XY{1, 1}}, XY{1, 0}, true},
	{&AABB{XY{0, 0}, XY{1, 1}}, XY{1, 1}, true},
	{&AABB{XY{0, 0}, XY{1, 1}}, XY{2, 0}, false},
	{&AABB{XY{0, 0}, XY{1, 1}}, XY{0, 2}, false},
	{&AABB{XY{0, 0}, XY{1, 1}}, XY{-2, 0}, false},
	{&AABB{XY{0, 0}, XY{1, 1}}, XY{0, 2}, false},
}

func TestAABBContainsPoint(t *testing.T) {
	for i, v := range containsPointsTests {
		out := v.box.ContainsPoint(&v.p)
		if out != v.exp {
			t.Errorf("%d. %v with input = %v: output %v expected %v", i, v.box, v.p, out, v.exp)
		}
	}
}

var intersectsAABBTests = []struct {
	a, b *AABB
	exp  bool
}{
	{&AABB{XY{0, 0}, XY{1, 1}}, &AABB{XY{0, 0}, XY{2, 2}}, true},   // 1 inside 2
	{&AABB{XY{0, 0}, XY{1, 1}}, &AABB{XY{0, 0}, XY{.5, .5}}, true}, // 1 contains 2
	{&AABB{XY{0, 0}, XY{1, 1}}, &AABB{XY{2, 0}, XY{2, .5}}, true},  // overlap on the right
	{&AABB{XY{0, 0}, XY{1, 1}}, &AABB{XY{0, 2}, XY{4, 2}}, true},   // overlap on top
	{&AABB{XY{0, 0}, XY{1, 1}}, &AABB{XY{-2, 0}, XY{2, 3}}, true},  // overlap on the left
	{&AABB{XY{0, 0}, XY{1, 1}}, &AABB{XY{0, -2}, XY{0, 3}}, true},  // overlap on the bottom
	{&AABB{XY{0, 0}, XY{1, 1}}, &AABB{XY{0, 0}, XY{2, .5}}, true},  // overlap on left and right
	{&AABB{XY{0, 0}, XY{1, 1}}, &AABB{XY{0, 0}, XY{.5, 2}}, true},  // overlap on top and bottom
	{&AABB{XY{0, 0}, XY{1, 1}}, &AABB{XY{-3, 0}, XY{1, 1}}, false}, // 1 right of 2
	{&AABB{XY{0, 0}, XY{1, 1}}, &AABB{XY{0, -3}, XY{1, 1}}, false}, // 1 above 2
	{&AABB{XY{0, 0}, XY{1, 1}}, &AABB{XY{3, 0}, XY{1, 1}}, false},  // 1 left of 2
	{&AABB{XY{0, 0}, XY{1, 1}}, &AABB{XY{0, 3}, XY{1, 1}}, false},  // 1 under 2

}

func TestAABBIntersctsAABB(t *testing.T) {
	for i, v := range intersectsAABBTests {
		out := v.a.IntersectsAABB(v.b)
		if out != v.exp {
			t.Errorf("%d. %v with inpute = %v: output %v expected %v", i, v.a, v.b, out, v.exp)
		}
	}
}

var qtRoot = New(AABB{XY{0, 0}, XY{10, 10}}, 4)

var qtInsertTests = []struct {
	pt  *XY
	exp bool
}{
	{&XY{5, 5}, true},
	{&XY{-5, 5}, true},
	{&XY{-5, -5}, true},
	{&XY{5, -5}, true},
	{&XY{0, 0}, true},
	{&XY{11, 0}, false},
	{&XY{-11, 0}, false},
	{&XY{0, -11}, false},
	{&XY{0, 11}, false},
}

func TestQTInsert(t *testing.T) {
	for i, v := range qtInsertTests {
		out := qtRoot.Insert(v.pt)
		if out != v.exp {
			t.Errorf("%d. %v with input = %v: output %v expected %v", i, qtRoot, v.pt, out, v.exp)
		}
	}
}

var qtSearchAreaTests = []struct {
	area *AABB
	exp  *XY
}{
	{&AABB{XY{5, 5}, XY{1, 1}}, &XY{5, 5}},
	{&AABB{XY{-5, 5}, XY{1, 1}}, &XY{-5, 5}},
	{&AABB{XY{-5, -5}, XY{1, 1}}, &XY{-5, -5}},
	{&AABB{XY{5, -5}, XY{1, 1}}, &XY{5, -5}},
	{&AABB{XY{11, 0}, XY{1, 1}}, nil},
	{&AABB{XY{-11, 0}, XY{1, 1}}, nil},
	{&AABB{XY{0, 11}, XY{1, 1}}, nil},
	{&AABB{XY{0, -11}, XY{1, 1}}, nil},
	{&AABB{XY{0, -11}, XY{1, 1}}, nil},
}

func TestQTSearchArea(t *testing.T) {
	for i, v := range qtSearchAreaTests {
		out := qtRoot.SearchArea(v.area)
		if len(out) == 0 {
			if v.exp == nil {
				continue
			} else {
				t.Errorf("%d. %v with input = %v: output %v expected %v", i, qtRoot, v.area, out, v.exp)
			}
			continue
		}
		if *out[0] != *v.exp {
			t.Errorf("%d. %v with input = %v: output %v expected %v", i, qtRoot, v.area, *out[0], *v.exp)
		}
	}
}
