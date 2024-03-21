package main

import (
	"math/rand"
	"time"
)

// Perceptron represents the perceptron model
type Perceptron struct {
	weights      []float64
	bias         float64
	learningRate float64
}

// NewPerceptron creates a new perceptron with random weights and bias
func NewPerceptron(inputSize int, learningRate float64) *Perceptron {
	rand.Seed(time.Now().UnixNano())
	weights := make([]float64, inputSize)
	for i := range weights {
		weights[i] = rand.Float64()*2 - 1 // Initialize weights between -1 and 1
	}
	return &Perceptron{
		weights:      weights,
		bias:         rand.Float64()*2 - 1,
		learningRate: learningRate,
	}
}

// Activate performs the weighted sum and applies the activation function
func (p *Perceptron) Activate(inputs []float64) int {
	sum := p.bias
	for i := range inputs {
		sum += inputs[i] * p.weights[i]
	}
	if sum > 0 {
		return 1
	}
	return -1
}

// Train adjusts the weights and bias based on the provided inputs and target
func (p *Perceptron) Train(inputs []float64, target int) {
	guess := p.Activate(inputs)
	error := target - guess

	// Adjust weights
	for i := range p.weights {
		p.weights[i] += p.learningRate * float64(error) * inputs[i]
	}

	// Adjust bias
	p.bias += p.learningRate * float64(error)
}

// GetWeightsAndBias returns the current weights and bias of the perceptron
func (p *Perceptron) GetWeightsAndBias() ([]float64, float64) {
	return p.weights, p.bias
}

/*func main() {
	// Example usage
	inputSize := 3
	learningRate := 0.1
	perceptron := NewPerceptron(inputSize, learningRate)

	// Training data
	trainingData := [][]float64{
		{0, 0, rand.Float64()*2 - 1},
		{0, 0, rand.Float64()*2 - 1},
		{0, 1, rand.Float64()*2 - 1},
		{0, 1, rand.Float64()*2 - 1},
		{1, 0, rand.Float64()*2 - 1},
		{1, 0, rand.Float64()*2 - 1},
		{1, 1, rand.Float64()*2 - 1},
		{1, 1, rand.Float64()*2 - 1},
	}
	targets := []int{-1, -1, 1, 1, 1, 1, 1, 1}
	println("OR_GATE")
	// Training the perceptron
	for epoch := 0; epoch < 10000; epoch++ {
		for i, inputs := range trainingData {
			target := targets[i]
			perceptron.Train(inputs, target)
		}
	}

	// Get the trained weights and bias
	weights, bias := perceptron.GetWeightsAndBias()
	fmt.Printf("Trained Weights: %v\n", weights)
	fmt.Printf("Trained Bias: %v\n", bias)

	// Testing the trained perceptron
	testData := [][]float64{
		{0, 0, 0},
		{0, 0, 1},
		{0, 1, 0},
		{0, 1, 1},
		{1, 0, 0},
		{1, 0, 1},
		{1, 1, 0},
		{1, 1, 1},
	}

	fmt.Println("Results after training:")
	for _, inputs := range testData {
		prediction := perceptron.Activate(inputs)
		fmt.Printf("Inputs: %v, Prediction: %d\n", inputs, prediction)
	}
}
*/
