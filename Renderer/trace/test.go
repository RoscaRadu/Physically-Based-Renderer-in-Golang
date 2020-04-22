package trace

import "testing"

func TestScale(t *testing.T) {
	testVector := Vector3{5, 5, 5}
	var scale float64 = 5
	expected := Vector3{25, 25, 25}
	if testVector.Scale(scale) != expected {
		t.Errorf(" expected %b %b %b but got %b %b %b",
			expected.X, expected.Y, expected.Z, testVector.X, testVector.Y, testVector.Z)
	}

}
