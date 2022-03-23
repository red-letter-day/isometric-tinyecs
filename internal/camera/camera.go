package camera

var Cam Camera

func init() {
	Cam = Camera{-500, -500}
}

type Camera struct {
	PosX float64
	PosY float64
}
