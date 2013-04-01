package gt2d

// Also known as a point
type Vector2D struct {
	X, Y int
}

func (this *Vector2D) Add(other *Vector2D) Vector2D {
	return Vector2D{this.X + other.X, this.Y + other.Y}
}

func (this *Vector2D) AddIP(other *Vector2D) {
	this.X += other.X
	this.Y += other.Y
}

func (this *Vector2D) Sub(other *Vector2D) Vector2D {
	return Vector2D{this.X - other.X, this.Y - other.Y}
}

func (this *Vector2D) SubIP(other *Vector2D) {
	this.X -= other.X
	this.Y -= other.Y
}

func (this *Vector2D) Dot(other *Vector2D) int {
	return this.X * other.X + this.Y * other.Y
}
