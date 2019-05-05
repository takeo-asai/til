package classifiers

import (
	"deeplearning/libs"
)

// MakeSamplePositions : 2.5 で使うサンプルデータを作成する
func MakeSamplePositions(class int, numberOfSamples int, dimension int) [][]float64 {
	g1 := libs.NewGaussian(-2.0, 1.0)
	g2 := libs.NewGaussian(+2.0, 1.0)
	data := make([][]float64, numberOfSamples)
	for i := 0; i < numberOfSamples; i++ {
		position := make([]float64, dimension)
		for j := 0; j < dimension; j++ {
			if (j+class)%2 == 0 {
				position[j] = g1.Distribution()
			} else {
				position[j] = g2.Distribution()
			}
		}
		data[i] = position
	}
	return data
}

// MakeSampleClasses : 2.5 で使うサンプル分類を作成する
func MakeSampleClasses(class int, numberOfSamples int) []int64 {
	data := make([]int64, numberOfSamples)
	for i := 0; i < numberOfSamples; i++ {
		if class == 1 {
			data[i] = 1
		} else {
			data[i] = -1
		}
	}
	return data
}
