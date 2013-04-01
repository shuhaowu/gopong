package gt2d

type Line2D struct {
	Start, End Vector2D
}

// This is a short hand for Line2D{Vector2D{x1, y1}, Vector2D{x2, y2}}
func NewLine(x1, y1, x2, y2 int) Line2D {
	return Line2D{Vector2D{x1, y1}, Vector2D{x2, y2}}
}
