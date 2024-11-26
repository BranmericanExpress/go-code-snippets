type size struct {
	width  int
	height int
}

type point struct {
	x int
	y int
}

type rectangle struct {
	size	size
	point	point
}

var rect rectangle {
	size: size{width: 10, height: 20},
	point: point{x: 5, y: 5},
}