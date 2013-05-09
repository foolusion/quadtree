package quadtree

// XY is a simple coordinate structure for points and vectors
type XY struct {
	X float64
	Y float64
}

// creates a new point and returns address
func NewXY(x, y float64) *XY {
	return &XY{x, y}
}

// Axis-Aligned bounding box structure with center and half dimension
type AABB struct {
	center  XY
	halfDim XY
}

// creates a new aabb and returns its address
func NewAABB(center, halfDim XY) *AABB {
	return &AABB{center, halfDim}
}

// Contains Point returns true when the AABB contains the point given
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

// Intersects AABB returns true when the AABB intersects another AABB
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
