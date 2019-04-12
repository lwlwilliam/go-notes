package geometry

// A Path is a journey connecting the points with straight lines.
type Path []Point

// Distance return s the distance traveled along the path.
func (path Path) Distance() float64  {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}
