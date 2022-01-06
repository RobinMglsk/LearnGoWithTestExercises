package clockface

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

const (
	secondsInClock  = 60
	secondsInHalfClock = secondsInClock / 2
	minutesInClock  = 60
	minutesInHalfClock = minutesInClock / 2
	hoursInClock  = 12
	hoursInHalfClock = hoursInClock / 2

)

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}

func secondHandPoint(t time.Time) Point {
	angle := secondsInRadians(t)
	return angleToPoint(angle)
}

func secondsInRadians(t time.Time) float64 {
	return (math.Pi / (secondsInHalfClock / float64(t.Second())))
}

func minuteHandPoint(t time.Time) Point {
	angle := minutesInRadians(t)
	return angleToPoint(angle)
}

func minutesInRadians(t time.Time) float64 {
	return (secondsInRadians(t) / secondsInClock ) + (math.Pi / (minutesInHalfClock / float64(t.Minute())))
}

func hoursInRadians(t time.Time) float64 {
	return (minutesInRadians(t) / hoursInClock ) + (math.Pi / (hoursInHalfClock / float64(t.Hour() % hoursInClock)))
}

func hourHandPoint(t time.Time) Point {
	angle := hoursInRadians(t)
	return angleToPoint(angle)
}