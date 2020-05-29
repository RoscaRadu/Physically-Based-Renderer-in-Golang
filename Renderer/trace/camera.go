package trace

// Camera simulates a camera
type Camera struct {
	Width   int
	Height  int
	Lens    float64
	Sensor  float64
	origin  Vector3
	dir     Vector3
	focus   float64
	fStop   float64
	toWorld Matrix4
}

func NewCamera(width, height int , lens float64) Camera {
	return Camera{
		Width: width,
		Height: height,
		Lens: lens,
		Sensor: 0.024, // 24mm height, full frame equivalent
	}
}


