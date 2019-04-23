package libs

import (
	"math"
	"math/rand"
	"time"
)

// Gaussian : 正規分布
func Gaussian(mean float64, variance float64) float64 {
	rand.Seed(time.Now().UnixNano())
	return rand.NormFloat64()*variance + mean
}

// BoxMuller : 正規分布(https://ja.wikipedia.org/wiki/ボックス＝ミュラー法)
func BoxMuller(mean float64, variance float64) float64 {
	rand.Seed(time.Now().UnixNano())
	r := rnd()
	c := math.Sqrt(-2.0 * math.Log(r))
	if rnd() < 0.5 {
		return c * math.Sin(2.0*math.Pi*rnd()*variance+mean)
	}
	return c * math.Cos(2.0*math.Pi*rnd()*variance+mean)
}

func rnd() float64 {
	return rand.Float64()
}
