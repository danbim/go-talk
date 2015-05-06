package vertex

// hidden because of non-capital letter
type hiddenVertex struct {
	A, B float64
}

// exported because of capital letter
type Vertex struct {
	X, Y float64
}

// exported because of capital letter
func New(x, y float64) Vertex {
	return Vertex{x, y}
}
