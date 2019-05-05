package libs

import (
	"math"
	"math/rand"
	"time"
)

// Gaussian : 正規分布
type Gaussian struct {
	Mean     float64
	Variance float64
}

// NewGaussian : constructor
func NewGaussian(mean float64, variance float64) *Gaussian {
	return &Gaussian{Mean: mean, Variance: variance}
}

// Distribution : golang 標準の方法で正規分布を返す
func (g *Gaussian) Distribution() float64 {
	rand.Seed(time.Now().UnixNano())
	return rand.NormFloat64()*g.Variance + g.Mean
}

// BoxMuller : 正規分布(https://ja.wikipedia.org/wiki/ボックス＝ミュラー法)
func (g *Gaussian) BoxMuller() float64 {
	rand.Seed(time.Now().UnixNano())
	r := rand.Float64()
	c := math.Sqrt(-2.0 * math.Log(r))
	if rand.Float64() < 0.5 {
		return c * math.Sin(2.0*math.Pi*rand.Float64()*g.Variance+g.Mean)
	}
	return c * math.Cos(2.0*math.Pi*rand.Float64()*g.Variance+g.Mean)
}
