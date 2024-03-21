package main

import (
	"fmt"
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

func main() {
	// Example usage
	inputSize := 2
	learningRate := 0.1
	perceptron := NewPerceptron(inputSize, learningRate)

	// Training data
	trainingData := [][]float64{
		{1, 0},
		{0, 1},
		{1, 1},
		{0, 0},
	}
	targets := []int{1, 1, -1, -1}

	// Training the perceptron
	for epoch := 0; epoch < 1000; epoch++ {
		for i, inputs := range trainingData {
			target := targets[i]
			perceptron.Train(inputs, target)
		}
	}

	// Testing the trained perceptron
	testData := [][]float64{
		{1, 0},
		{0, 1},
		{1, 1},
		{0, 0},
	}

	fmt.Println("Results after training:")
	for _, inputs := range testData {
		prediction := perceptron.Activate(inputs)
		fmt.Printf("Inputs: %v, Prediction: %d\n", inputs, prediction)
	}
}
