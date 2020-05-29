package trace

import (
	"math"
	"math/rand"
)

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

func  NewLight(r,g,b float64) Material {
	return Material{Light: Vector3{r,g,b}}
}

func NewPlastic (r,g,b float64, gloss float64) Material{
	return Material{
		Color: Vector3{r,g,b},
		Fresnel: Vector3{0.04,0.04,0.04},
		Opacity: 1,
		Gloss: gloss,
	}
}

func NewMetal(r,g,b float64, gloss float64) Material{
	return Material{
		Fresnel: Vector3{r,g,b},
		Opacity: 1,
		Gloss: gloss,
		Metal: 1,
	}
}

func NewGlass(r,g,b, opacity float64,gloss float64) Material{
	return Material{
		Color: Vector3{r,g,b},
		Fresnel: Vector3{0.04,0.04,0.04},
		Refract: 1.52,
		Opacity: opacity,
		Gloss: gloss,
	}
}

func (m *Material) Bsdf(normal Vector3, incident Vector3, dist float64, rnd *rand.Rand) (next bool, dir Vector3, signal Vector3){
	if incident.Enters(normal){
		//is reflected
		reflect:=m.schlick(normal,incident)
		if rnd.Float64() <=reflect.Average() {
			tint:= Vector3{1,1,1}.Lerp(m.Fresnel,m.Metal)
			return true,incident.Reflect(normal).Cone(1-m.Gloss,rnd), tint
		}
		//transmitted
		if rnd.Float64() > m.Opacity{
			refracted, dir:= incident.Refract(normal,1,m.Refract)
			return refracted,dir.Cone(1-m.Gloss,rnd),Vector3{1,1,1}
		}
		//absorbed
		if rnd.Float64()<m.Metal{
			return false, incident, Vector3{0, 0, 0}
		}
		return true, normal.RandHemiCos(rnd),m.Color.Scale(1/math.Pi)
	}
	exited, dir := incident.Refract(normal.Scale(-1),m.Refract,1)
	volume:=math.Min(m.Opacity*dist*dist,1)
	tint:=Vector3{1,1,1}.Lerp(m.Color,volume)
	return exited, dir, tint
}

// http://blog.selfshadow.com/publications/s2015-shading-course/hoffman/s2015_pbs_physics_math_slides.pdf
// http://graphics.stanford.edu/courses/cs348b-10/lectures/reflection_i/reflection_i.pdf

func (m *Material) schlick(incident Vector3, normal Vector3) Vector3{
	cos:= incident.Scale(-1).Dot(normal)
	invFresnel := Vector3{1,1,1}.Minus(m.Fresnel)
	scaled := invFresnel.Scale(math.Pow(1-cos,5))
	return m.Fresnel.Add(scaled)
}


func (m *Material) Emit(normal Vector3, dir Vector3) Vector3{
	cos := math.Max(normal.Dot(dir.Scale(-1)),0)
	return m.Light.Scale(cos)
}
