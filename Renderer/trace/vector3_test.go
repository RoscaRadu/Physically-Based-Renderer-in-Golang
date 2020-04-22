package trace

import (
	"math"
	"testing"
)

func TestMult(t *testing.T){
	testVector := Vector3{5, 5, 5}
	testVector2 := Vector3{3, 3, 3}
	expected := Vector3{15, 15, 15}
	if testVector.Mult(testVector2) != expected {
		t.Errorf(" expected %3f but got %3f",
			expected, testVector)
	}
}

func TestScale(t *testing.T) {
	testVector := Vector3{5, 5, 5}
	var scale float64 = 5
	expected := Vector3{25, 25, 25}
	if testVector.Scale(scale) != expected {
		t.Errorf(" expected %3f but got %3f",
			expected,testVector)
	}

}

func TestAdd(t *testing.T){
	testVector := Vector3{5, 5, 5}
	testVector2 := Vector3{3, 3, 3}
	expected:= Vector3{8,8,8,}
	if testVector.Add(testVector2) != expected {
		t.Errorf(" expected %3f but got %3f ",
			expected, testVector)
	}
}

/*
func TestRefract(t *testing.T){

}
*/


func TestDot(t *testing.T){
	testVector := Vector3{5, 5, 5}
	testVector2 := Vector3{3, 3, 3}
	var expected float64= 45
	if testVector.Dot(testVector2)!=expected{
		t.Errorf(" expected %f but got %f",
			expected,testVector.Dot(testVector2))
	}
}

func TestAverage(t *testing.T){
	testVector := Vector3{6, 10, 20}

	var expected float64= 12
	if testVector.Average()!=expected{
		t.Errorf(" expected %f but got %f",
			expected,testVector.Average())
	}
}
func TestMax(t *testing.T){
	testVector := Vector3{50, 25, 5}
	var expected float64= 50

	if testVector.Max()!=expected{
		t.Errorf(" expected %f but got %f",
			expected,testVector.Max())
	}
}


func TestArray(t *testing.T){
	testVector := Vector3{5,3,1,}
	expected := [3]float64 {5,3,1}

	if testVector.Array()!=expected{
		t.Errorf(" expected %3f but got %3f",
			expected,testVector.Array())
	}
}

func TestEnters(t *testing.T){
	testVector := Vector3{5,3,1}
	testVector2 := Vector3{5,3,1}

	var expected bool = false // true means smaller than 0 , false means larger than 0
	if testVector.Enters(testVector2)!=expected{
		t.Errorf(" expected %t but got %t",
			expected,testVector.Enters(testVector2))
	}
}

func TestCross(t *testing.T){
	testVector := Vector3{5,3,1}
	testVector2 := Vector3{5,3,1}

	expected := Vector3{0,0,0}
	if testVector.Cross(testVector2)!=expected{
		t.Errorf(" expected %3f but got %3f",
			expected,testVector.Cross(testVector2))
	}

}

func TestMinus(t *testing.T){
	testVector := Vector3{5,3,1}
	testVector2 := Vector3{5,3,1}

	expected := Vector3{0,0,0}
	if testVector.Minus(testVector2)!=expected{
		t.Errorf(" expected %3f but got %3f",
			expected,testVector.Minus(testVector2))
	}
}

func TestNormalize(t *testing.T){
	testVector := Vector3{6,3,9}
	d:=testVector.Length()
	expected := Vector3{testVector.X/d,testVector.Y/d,testVector.Z/d}
	if testVector.Normalize()!=expected{
		t.Errorf(" expected %3f but got %3f",
			expected,testVector.Normalize())
	}
}

func TestLength(t *testing.T){
	testVector := Vector3{4,3,2}
	var expected float64= 29
	if testVector.Length()!=math.Sqrt(expected){
		t.Errorf(" expected %f but got %f",
			expected,testVector.Length())
	}
}

func TestLerp(t *testing.T){
	testVector := Vector3{2,1,2}
	testVector2 := Vector3{1,2,3}
	var n float64
	n=3

	expected:=Vector3{-1,4,5}
	if testVector.Lerp(testVector2,n)!=expected{
		t.Errorf(" expected %3f but got %3f",
			expected,testVector.Lerp(testVector2,n))
	}
}