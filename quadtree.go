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
	TopLeft     XY
	BottomRight XY
}

// creates a new aabb and returns its address
func NewAABB(center, halfDimension XY) *AABB {
	return &AABB{center, halfDimension}
}

// Contains Point returns true when the AABB contains the point given
func (aabb *AABB) ContainsPoint(p XY) bool {
	return false
}

// Intersects AABB returns true when the AABB intersects another AABB
func (aabb *AABB) IntersectsAABB(other AABB) bool {
	return false
}
