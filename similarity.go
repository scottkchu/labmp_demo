package main

import (
	"math"
)

// computeCosineSimilarity computes the cosine similarity score between two embeddings.
func computeCosineSimilarity(embedding1, embedding2 []float64) float64 {
	// Compute dot product
	dotProduct := 0.0
	for i := range embedding1 {
		dotProduct += embedding1[i] * embedding2[i]
	}

	// Compute magnitudes
	magnitude1 := 0.0
	for i := range embedding1 {
		magnitude1 += math.Pow(embedding1[i], 2)
	}
	magnitude1 = math.Sqrt(magnitude1)

	magnitude2 := 0.0
	for i := range embedding2 {
		magnitude2 += math.Pow(embedding2[i], 2)
	}
	magnitude2 = math.Sqrt(magnitude2)

	// Compute cosine similarity score
	if magnitude1 > 0 && magnitude2 > 0 {
		return dotProduct / (magnitude1 * magnitude2)
	}
	return 0.0
}

func GetBestMatch(questionEmbedding []float64, faqEmbedding [][]float64) (int, float64) {
	var bestMatchIndex int
	var bestMatchScore float64

	for i, embedding := range faqEmbedding {
		score := computeCosineSimilarity(questionEmbedding, embedding)
		if score > bestMatchScore {
			bestMatchScore = score
			bestMatchIndex = i
		}
	}

	return bestMatchIndex, bestMatchScore
}

// func main() {
// 	// Example usage
// 	embedding1 := []float64{1.0, 2.0, 3.0}
// 	embedding2 := []float64{4.0, 5.0, 6.0}
// 	score := ComputeCosineSimilarity(embedding1, embedding2)
// 	fmt.Printf("Cosine similarity score: %f\n", score)
// }
// In this example, the ComputeCosineSimilarity function takes two embeddings as input, represented as slices of float64 values. The function computes the dot product of the two embeddings, as well as their magnitudes, using the math package from the standard library. The cosine similarity score is then computed as the dot product divided by the product of the magnitudes. If either magnitude is zero, the function returns a similarity score of zero.

// This implementation assumes that the two embeddings have the same length. If this is not the case, the function will panic at runtime when attempting to access out-of-bounds elements in one of the slices.





