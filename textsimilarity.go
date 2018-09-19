package textsimilarity

import (
	"errors"
	"math"
	"strings"
)

// Cosine returns the Cosine Similarity between two vectors.
func Cosine(a, b []float64) (float64, error) {
	count := 0
	lengthA := len(a)
	lengthB := len(b)
	if lengthA > lengthB {
		count = lengthA
	} else {
		count = lengthB
	}
	sumA := 0.0
	s1 := 0.0
	s2 := 0.0
	for k := 0; k < count; k++ {
		if k >= lengthA {
			s2 += math.Pow(b[k], 2)
			continue
		}
		if k >= lengthB {
			s1 += math.Pow(a[k], 2)
			continue
		}
		sumA += a[k] * b[k]
		s1 += math.Pow(a[k], 2)
		s2 += math.Pow(b[k], 2)
	}
	if s1 == 0 || s2 == 0 {
		return 0.0, errors.New("null vector")
	}
	return sumA / (math.Sqrt(s1) * math.Sqrt(s2)), nil
}

// Tokenize is a naive implementation that currently only splits by spaces.
func Tokenize(s string) []string {
	return strings.Split(strings.ToLower(s), " ")
}

func count(key string, a []string) int {
	count := 0
	for _, s := range a {
		if key == s {
			count = count + 1
		}
	}
	return count
}

// Similarity returns the cosine similarity between two documents.
func Similarity(a, b string) (float64, error) {
	var corpus []string

	tokensA := Tokenize(a)
	tokensB := Tokenize(b)
	allTokens := append(tokensA, tokensB...)
	encountered := map[string]bool{}

	// Generate a corpus.
	for _, t := range allTokens {
		if encountered[t] != true {
			encountered[t] = true
			corpus = append(corpus, t)
		}
	}

	// Populate the vectors using frequency in the corpus.
	vectorA := make([]float64, len(corpus))
	vectorB := make([]float64, len(corpus))
	for k, v := range corpus {
		vectorA[k] = float64(count(v, tokensA))
		vectorB[k] = float64(count(v, tokensB))
	}

	similarity, err := Cosine(vectorA, vectorB)
	if err != nil {
		return 0.0, err
	}
	return similarity, nil
}
