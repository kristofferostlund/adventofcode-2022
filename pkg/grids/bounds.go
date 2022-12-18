package grids

type Bounds struct {
	minX, maxX int
	minY, maxY int
}

func NewBounds(minX, maxX, minY, maxY int) Bounds {
	return Bounds{
		minX: minX,
		maxX: maxX,
		minY: minY,
		maxY: maxY,
	}
}

func (b Bounds) MaxX() int {
	return b.maxX
}

func (b Bounds) MinX() int {
	return b.minX
}

func (b Bounds) MaxY() int {
	return b.maxY
}

func (b Bounds) MinY() int {
	return b.minY
}

func (b Bounds) Extend(loc Loc) Bounds {
	x, y := loc.XY()

	if b.maxX < x {
		b.maxX = x
	}
	if x < b.minX {
		b.minX = x
	}

	if b.maxY < y {
		b.maxY = y
	}
	if y < b.minY {
		b.minY = y
	}

	return b
}

func (b Bounds) IsInside(loc Loc) bool {
	x, y := loc.XY()

	minX, maxX := b.minX, b.maxX
	minY, maxY := b.minY, b.maxY

	return minX <= x && x <= maxX &&
		minY <= y && y <= maxY
}
