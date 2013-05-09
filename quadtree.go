// Package quadtree implements methods for a quadtree spatial partitioning data
// structure.
//
// Code is based on the Wikipedia article
// http://en.wikipedia.org/wiki/Quadtree.
package quadtree

// node_capacity is the maximum number of points allowed in a quadtree node
var node_capacity int = 4

// XY is a simple coordinate structure for points and vectors
type XY struct {
	X float64
	Y float64
}

// NewXY creates a new point and returns address
func NewXY(x, y float64) *XY {
	return &XY{x, y}
}

// AABB represents an Axis-Aligned bounding box structure with center and half
// dimension
type AABB struct {
	center  XY
	halfDim XY
}

// NewAABB creates a new axis-aligned bounding box and returns its address
func NewAABB(center, halfDim XY) *AABB {
	return &AABB{center, halfDim}
}

// ContainsPoint returns true when the AABB contains the point given
func (aabb *AABB) ContainsPoint(p XY) bool {
	if p.X < aabb.center.X-aabb.halfDim.X {
		return false
	}
	if p.Y < aabb.center.Y-aabb.halfDim.Y {
		return false
	}
	if p.X > aabb.center.X+aabb.halfDim.X {
		return false
	}
	if p.Y > aabb.center.Y+aabb.halfDim.Y {
		return false
	}

	return true
}

// IntersectsAABB returns true when the AABB intersects another AABB
func (aabb *AABB) IntersectsAABB(other *AABB) bool {
	if other.center.X+other.halfDim.X < aabb.center.X-aabb.halfDim.X {
		return false
	}
	if other.center.Y+other.halfDim.Y < aabb.center.Y-aabb.halfDim.Y {
		return false
	}
	if other.center.X-other.halfDim.X > aabb.center.X+aabb.halfDim.X {
		return false
	}
	if other.center.Y-other.halfDim.Y > aabb.center.Y+aabb.halfDim.Y {
		return false
	}

	return true
}

// QuadTree represents the quadtree data structure.
type QuadTree struct {
	boundary  AABB
	points    []*XY
	northWest *QuadTree
	northEast *QuadTree
	southWest *QuadTree
	southEast *QuadTree
}

// New creates a new quadtree node that is bounded by boundary and contains
// node_capacity points.
func New(boundary AABB) *QuadTree {
	points := make([]*XY, 0, node_capacity)
	qt := &QuadTree{boundary: boundary, points: points}
	return qt
}

// Insert adds a point to the quadtree. It returns true if it was successful
// and false otherwise.
func (qt *QuadTree) Insert(p XY) bool {
	return false
}
