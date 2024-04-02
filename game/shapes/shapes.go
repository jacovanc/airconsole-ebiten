package shapes

type Vector struct {
	X, Y float64
}

type Rectangle struct {
	Width, Height float64
}

type CollisionBox struct {
	Position Vector
	Box      Rectangle
}