package classifiers

import (
	"deeplearning/libs"
)

type Perceptrons struct {
	weights       []float64
	numberOfInput int
}

func NewPerceptrons(n int) *Perceptrons {
	w := make([]float64, n)
	return &Perceptrons{w, n}
}

func (p *Perceptrons) train(position []float64, objective int64, learningRate float64) int64 {
	if len(position) != p.numberOfInput {
		panic("Invalid argements are passed to Perceptrons.train()")
	}

	// 分類: c > 0 で正しく分類されている
	c := 0.0
	for i := 0; i < p.numberOfInput; i++ {
		c += p.weights[i] * position[i] * float64(objective)
	}
	if c > 0 {
		return 1
	}

	// 分類に失敗した場合は重みを更新する
	for i := 0; i < p.numberOfInput; i++ {
		p.weights[i] += learningRate * position[i] * float64(objective)
	}

	return 0
}

func (p *Perceptrons) Fit(numberOfIterations int, positions [][]float64, objectives []int64, learningRate float64) {
	for epoch := 0; epoch < numberOfIterations; epoch++ {
		for i := 0; i < len(positions); i++ {
			p.train(positions[i], objectives[i], learningRate)
		}
	}
}

func (p *Perceptrons) Predict(positions [][]float64) []int64 {
	results := make([]int64, len(positions))
	for i, position := range positions {
		if len(position) != p.numberOfInput {
			panic("Invalid argements are passed to Perceptrons.Predict()")
		}

		c := 0.0
		for i := 0; i < p.numberOfInput; i++ {
			c += p.weights[i] * position[i]
		}
		results[i] = libs.Step(c)
	}
	return results
}
