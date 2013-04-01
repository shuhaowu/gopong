package gt2d

import (
	"testing"
)

func TestVector2DAdd(t *testing.T){
	vector1 := Vector2D{1, 1}
	vector2 := Vector2D{5, 6}
	vector3 := vector1.Add(&vector2)
	if vector3.X != 6 || vector3.Y != 7 {
		t.FailNow()
	}

	vector1.AddIP(&vector2)
	if vector1.X != 6 || vector1.Y != 7 {
		t.FailNow()
	}
}

func TestVector2DSub(t *testing.T){
	vector1 := Vector2D{1, 1}
	vector2 := Vector2D{5, 6}
	vector3 := vector1.Sub(&vector2)

	if vector3.X != -4 || vector3.Y != -5 {
		t.FailNow()
	}

	vector1.SubIP(&vector2)
	if vector1.X != -4 || vector1.Y != -5 {
		t.FailNow()
	}
}

func TestVector2DDot(t *testing.T){
	vector1 := Vector2D{1, 1}
	vector2 := Vector2D{5, 6}
	if vector1.Dot(&vector2) != 11 {
		t.FailNow()
	}
}