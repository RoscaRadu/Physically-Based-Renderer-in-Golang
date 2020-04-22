package trace

type Material struct {
	Color   Vector3
	Light   Vector3
	Fresnel Vector3

	Metal   float64
	Opacity float64
	Gloss   float64
	Refract float64
}

//TODO  NewLight, NewPlastic, NewMetal, NewGlass functions.
