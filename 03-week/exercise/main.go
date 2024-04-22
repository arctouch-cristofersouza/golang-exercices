package main

type shapes interface {
	volume() float64
}

func calculateVolume(shapes shapes) float64 {
	return shapes.volume()
}
