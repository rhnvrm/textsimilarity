package textsimilarity

import (
	"errors"
	"math"
	"strings"
)

var ()

type TextSimilarity struct {
	corpus            []string
	documentFrequency map[string]int
}

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

// Tokenize uses Naive splitting along with our custom stopword
// filtering to return a list of tokens.
func Tokenize(s string) []string {
	// Iterate over the doc's tokens:
	tokens := []string{}
	for _, tok := range strings.Split(strings.ToLower(s), " ") {
		var exclude = false
		for _, v := range stopbytes {
			if string(v) == tok {
				exclude = true
			}
		}
		if exclude == false {
			tokens = append(tokens, tok)
		}
	}
	return tokens
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

func tfidf(v string, tokens []string, n int, documentFrequency map[string]int) float64 {
	tf := float64(count(v, tokens))
	idf := math.Log(float64(n) / (float64(documentFrequency[v] + 1)))
	return tf * idf
}

// New accepts a slice of documents and
// creates the internal corpus and document frequency mapping.
func New(documents []string) *TextSimilarity {
	var (
		allTokens []string
		ts        TextSimilarity
	)

	ts.documentFrequency = map[string]int{}

	for _, doc := range documents {
		allTokens = append(allTokens, Tokenize(doc)...)
	}
	// Generate a corpus.
	for _, t := range allTokens {
		if ts.documentFrequency[t] == 0 {
			ts.documentFrequency[t] = 1
			ts.corpus = append(ts.corpus, t)
		} else {
			ts.documentFrequency[t] = ts.documentFrequency[t] + 1
		}
	}

	return &ts
}

// Similarity returns the cosine similarity between two documents using
// Tf-Idf vectorization using the corpus.
func (ts *TextSimilarity) Similarity(a, b string) (float64, error) {
	tokensA := Tokenize(a)
	tokensB := Tokenize(b)
	// Populate the vectors using frequency in the corpus.
	n := len(ts.corpus)
	vectorA := make([]float64, n)
	vectorB := make([]float64, n)
	for k, v := range ts.corpus {
		vectorA[k] = tfidf(v, tokensA, n, ts.documentFrequency)
		vectorB[k] = tfidf(v, tokensB, n, ts.documentFrequency)
	}

	similarity, err := Cosine(vectorA, vectorB)
	if err != nil {
		return 0.0, err
	}
	return similarity, nil
}
