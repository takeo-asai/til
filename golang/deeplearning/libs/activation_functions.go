package libs

import (
	"math"
)

func Step(x float64) int64 {
	if x >= 0 {
		return 1
	}
	return -1
}

func Sigmoid(x float64) float64 {
	return 1.0 / (1.0 + math.Pow(math.E, -x))
}

func DSigmoid(x float64) float64 {
	return x * (1.0 - x)
}

func Tanh(x float64) float64 {
	return math.Tanh(x)
}

func DTanh(x float64) float64 {
	return 1.0 - x*x
}

func ReLU(x float64) float64 {
	if x > 0 {
		return x
	}
	return 0.0
}

func DReLU(x float64) float64 {
	if x > 0 {
		return 1.0
	}
	return 0.0
}

func Softmax(xs []float64) []float64 {
	// max をとっておくことで、オーバーフローを防ぐ
	max := 0.0
	for _, x := range xs {
		if max < x {
			max = x
		}
	}

	sum := 0.0
	for _, x := range xs {
		sum += math.Exp(x - max)
	}

	values := []float64{}
	for _, x := range xs {
		values = append(values, math.Exp(x-max)/sum)
	}

	return values
}
