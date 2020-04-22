package trace

var yAxis Vector3

func init() {
	yAxis = Vector3{0, 1, 0}
}

//Matrix4 is for storing matrix data and matrix operations
//It is stored in a Column-major way
type Matrix4 struct {
	el [4][4]float64
}

//Creates a new matrix with the given information
func NewMatrix4(a [4]float64,b [4]float64,c [4]float64,d [4]float64) (m Matrix4) {
	m.el[0] = a
	m.el[1] = b
	m.el[2] = c
	m.el[3] = d
	return
}

//Creates a new Identity matrix
func NewIdentityMatrix4() (m Matrix4) {
	m.el[0] = [4]float64{1, 0, 0, 0}
	m.el[1] = [4]float64{0, 1, 0, 0}
	m.el[2] = [4]float64{0, 0, 1, 0}
	m.el[3] = [4]float64{0, 0, 0, 1}
	return
}

//Creates a view matrix from the world position of the camera(eye) and a target point (the place we Look at)
func NewLookAtMatrix4(eye Vector3, target Vector3) Matrix4 {
	forward := eye.Minus(target).Normalize()
	right := yAxis.Cross(forward).Normalize()
	u := forward.Cross(right).Normalize()

	a:=[4]float64 {right.X, u.X, forward.X, 0}
	b:=[4]float64 {right.Y, u.Y, forward.Y, 0}
	c:=[4]float64 {right.Z, u.Z, forward.Z, 0}
	d:=[4]float64 {0, 0, 0, 1}

	return NewMatrix4(a,b,c,d)

}

//Multiplies two given matrices
func (a *Matrix4) Mult(b Matrix4) (result Matrix4) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 4; k++ {
				result.el[j][i] += a.el[k][i] * b.el[j][k]
			}
		}
	}
	return
}

//Multiplies a matrix4 with a given vector, including translation
func (a *Matrix4) ApplyPoint(v Vector3) (result Vector3) {
	result.X = v.X*a.el[0][0] + v.Y*a.el[1][0] + v.Z*a.el[2][0] + a.el[3][0]
	result.Y = v.X*a.el[0][1] + v.Y*a.el[1][1] + v.Z*a.el[2][1] + a.el[3][1]
	result.Z = v.X*a.el[0][2] + v.Y*a.el[1][2] + v.Z*a.el[2][2] + a.el[3][2]
	// final row assumed to be [0,0,0,1]
	return
}

//Multiplies a matrix4 with a given vector, excluding translation
func (a *Matrix4) ApplyDir(v Vector3) (result Vector3) {
	result.X = v.X*a.el[0][0] + v.Y*a.el[1][0] + v.Z*a.el[2][0]
	result.Y = v.X*a.el[0][1] + v.Y*a.el[1][1] + v.Z*a.el[2][1]
	result.Z = v.X*a.el[0][2] + v.Y*a.el[1][2] + v.Z*a.el[2][2]
	return
}
