package trace

import (
	"testing"
)

func TestMult(t *testing.T){
	testVector := Vector3{5, 5, 5}
	testVector2 := Vector3{3, 3, 3}
	expected := Vector3{15, 15, 15}
	if testVector.Mult(testVector2) != expected {
		t.Errorf(" expected %f %f %f but got %f %f %f",
			expected.X, expected.Y, expected.Z, testVector.X, testVector.Y, testVector.Z)
	}
}

func TestScale(t *testing.T) {
	testVector := Vector3{5, 5, 5}
	var scale float64 = 5
	expected := Vector3{25, 25, 25}
	if testVector.Scale(scale) != expected {
		t.Errorf(" expected %f %f %f but got %f %f %f",
			expected.X, expected.Y, expected.Z, testVector.X, testVector.Y, testVector.Z)
	}

}

func TestAdd(t *testing.T){
	testVector := Vector3{5, 5, 5}
	testVector2 := Vector3{3, 3, 3}
	expected:= Vector3{8,8,8,}
	if testVector.Add(testVector2) != expected {
		t.Errorf(" expected %f %f %f but got %f %f %f",
			expected.X, expected.Y, expected.Z, testVector.X, testVector.Y, testVector.Z)
	}
}

/*
func TestRefract(t *testing.T){

}
*/


func TestDot(t *testing.T){
	testVector := Vector3{5, 5, 5}
	testVector2 := Vector3{3, 3, 3}
	var expected float64= 43
	if testVector.Dot(testVector2)!=expected{
		t.Errorf(" expected %f but got %f",
			expected,testVector.Dot(testVector2))
	}
}


