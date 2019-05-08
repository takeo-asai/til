package main

import (
	"bufio"
	"flag"
	"fmt"
	"image"
	"os"
	"sort"

	"github.com/disintegration/imaging"
	"github.com/pfnet-research/go-menoh"
)

func main() {
	const (
		batch   = 1
		channel = 3
		width   = 224
		height  = 224

		conv1_1InName  = "Input_0"
		softmaxOutName = "Softmax_0"
	)
	var (
		inputImagePath  = flag.String("input-image", "data/Light_sussex_hen.jpg", "input image path")
		onnxModelPath   = flag.String("model", "data/vgg16.onnx", "ONNX model path")
		synsetWordsPath = flag.String("synset-words", "data/synset_words.txt", "synset words file path")
	)
	flag.Parse()

	// prepare input data
	imageFile, err := os.Open(*inputImagePath)
	if err != nil {
		panic(err)
	}
	defer imageFile.Close()
	img, _, err := image.Decode(imageFile)
	if err != nil {
		panic(err)
	}
	resizedImg := imaging.Resize(img, width, height, imaging.Linear)
	bgrMean := []float32{103.939, 116.779, 123.68} // VGG16の平均画素を引く必要がある
	resizedImgTensor := &menoh.FloatTensor{
		Dims:  []int32{batch, channel, height, width},
		Array: toOneHotFloats(resizedImg, channel, bgrMean),
	}

	// build model runner
	runner, err := menoh.NewRunner(menoh.Config{
		ONNXModelPath: *onnxModelPath,
		Backend:       menoh.TypeMKLDNN,
		BackendConfig: "",
		Inputs: []menoh.InputConfig{
			{
				Name:  conv1_1InName,
				Dtype: menoh.TypeFloat,
				Dims:  []int32{batch, channel, height, width},
			},
		},
		Outputs: []menoh.OutputConfig{
			{
				Name:         softmaxOutName,
				Dtype:        menoh.TypeFloat,
				FromInternal: false,
			},
		},
	})
	if err != nil {
		panic(err)
	}
	defer runner.Stop()

	// run ONNX model with input and get output result
	if err := runner.RunWithTensor(conv1_1InName, resizedImgTensor); err != nil {
		panic(err)
	}
	softmaxOutTensor, err := runner.GetOutput(softmaxOutName)
	if err != nil {
		panic(err)
	}
	softmaxOutData, _ := softmaxOutTensor.FloatArray()

	// evalute image detection
	categories, err := loadCategoryList(*synsetWordsPath)
	if err != nil {
		panic(err)
	}
	topKIndices := extractTopKIndexList(softmaxOutData, 5)
	fmt.Println("top 5 categories are")
	for _, idx := range topKIndices {
		fmt.Printf(" -%3d: %.5f, %s\n", idx, softmaxOutData[idx], categories[idx])
	}
}

// ================================================================================

func toOneHotFloats(img image.Image, channel int, bgrMean []float32) []float32 {
	bounds := img.Bounds()
	w, h := bounds.Dx(), bounds.Dy()
	floats := make([]float32, channel*h*w)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			floats[0*(w*h)+y*w+x] = float32(r/257) - bgrMean[2]
			floats[1*(w*h)+y*w+x] = float32(g/257) - bgrMean[1]
			floats[2*(w*h)+y*w+x] = float32(b/257) - bgrMean[0]
		}
	}
	return floats
}

// synset_words.txt から categories を取り出す
func loadCategoryList(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()

	categories := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		categories = append(categories, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return []string{}, err
	}
	return categories, nil
}

// category の順番と values の index が一致
// values の値で高いものから, その index(= category_id)を返すようにすると,
// txt から取り出した (category_id, category_name) から category_name を取り出せる
func extractTopKIndexList(values []float32, k int) []int {
	type pair struct {
		index int
		value float32
	}
	pairs := make([]pair, len(values))
	for i, f := range values {
		pairs[i] = pair{
			index: i,
			value: f,
		}
	}
	sort.SliceStable(pairs, func(i, j int) bool {
		return pairs[i].value > pairs[j].value
	})
	topKIndices := make([]int, k)
	for i := 0; i < k; i++ {
		topKIndices[i] = pairs[i].index
	}
	return topKIndices
}
