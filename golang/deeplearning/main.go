package main

import (
	"deeplearning/classifiers"
	"fmt"
)

func main() {
	dimension := 2
	epochs := 100
	learningRate := 1.0

	// テストデータ作成
	numberOfTrains := 1000
	numberOfTests := 2000
	trainsX := append(classifiers.MakeSamplePositions(0, numberOfTrains/2, dimension), classifiers.MakeSamplePositions(1, numberOfTrains/2, dimension)...)
	trainsY := append(classifiers.MakeSampleClasses(0, numberOfTrains/2), classifiers.MakeSampleClasses(1, numberOfTrains/2)...)
	testsX := append(classifiers.MakeSamplePositions(0, numberOfTests/2, dimension), classifiers.MakeSamplePositions(1, numberOfTests/2, dimension)...)
	testsY := append(classifiers.MakeSampleClasses(0, numberOfTests/2), classifiers.MakeSampleClasses(1, numberOfTests/2)...)

	// 分類
	classifier := classifiers.NewPerceptrons(dimension)
	classifier.Fit(epochs, trainsX, trainsY, learningRate)
	predicted := classifier.Predict(testsX)

	// 検証
	matrix := [2][2]int64{}
	for i := range predicted {
		if testsY[i] > 0 {
			if predicted[i] > 0 {
				matrix[0][0] += 1.0
			} else {
				matrix[0][1] += 1.0
			}
		} else {
			if predicted[i] > 0 {
				matrix[1][0] += 1.0
			} else {
				matrix[1][1] += 1.0
			}
		}
	}
	accuracy := (float64(matrix[0][0]) + float64(matrix[1][1])) / (float64(matrix[0][0]) + float64(matrix[0][1]) + float64(matrix[1][0]) + float64(matrix[1][1]))
	precision := float64(matrix[0][0]) / (float64(matrix[0][0]) + float64(matrix[1][0]))
	recall := float64(matrix[0][0]) / (float64(matrix[0][0]) + float64(matrix[0][1]))
	fmt.Printf("accuracy: %v\n", accuracy)
	fmt.Printf("precision: %v\n", precision)
	fmt.Printf("recall: %v\n", recall)
}
