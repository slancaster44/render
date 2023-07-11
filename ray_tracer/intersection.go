package ray_tracer

type Intersection struct {
	Shape  *Triangle
	TValue float64
}

func NewIntersection(s *Triangle, t float64) *Intersection {
	return &Intersection{s, t}
}

//Assumes all intersections are from the same ray
func GetHit(list []*Intersection) *Intersection {
	var retVal *Intersection = nil
	for _, val := range list {
		if val.TValue >= 0 && (retVal == nil || retVal.TValue > val.TValue) {
			retVal = val
		}
	}

	return retVal
}
