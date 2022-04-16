package camera

var Cam Camera

func init() {
	Cam = Camera{-200, -200}
}

type Camera struct {
	PosX float64
	PosY float64
}
