package numbers

import (
	"errors"
	"math"
)

var (
	InvalidAccuracy = errors.New("invalid accuracy (accuracy >= 1)")
)

// Округление вниз с указанием знаков после запятой
func Floor(num float64, accuracy int) float64 {
	if accuracy < 0 {
		panic(InvalidAccuracy)
	}

	pow := math.Pow(10.0, float64(accuracy))

	return math.Floor(num*pow) / pow
}

// Округление к ближайшему с указанием знаков после запятой
func Round(num float64, accuracy int) float64 {
	if accuracy < 0 {
		panic(InvalidAccuracy)
	}

	pow := math.Pow(10.0, float64(accuracy))

	return math.Round(num*pow) / pow
}

// Округление вверх с указанием знаков после запятой
func Ceil(num float64, accuracy int) float64 {
	if accuracy < 0 {
		panic(InvalidAccuracy)
	}

	pow := math.Pow(10.0, float64(accuracy))

	return math.Ceil(num*pow) / pow
}
