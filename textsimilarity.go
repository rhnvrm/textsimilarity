package textsimilarity

import (
	"errors"
	"math"
	"sort"
	"strings"

	tokenize "github.com/AlasdairF/Tokenize"
)

type kv struct {
	Key   string
	Value float64
}

// TextSimilarity is a struct containing internal
// data to be re-used by the package.
type TextSimilarity struct {
	stopwords         [][]byte
	corpus            []string
	documents         []string
	documentFrequency map[string]int

	// Configuration Options
	useBiGrams bool
}

// Option type describes functional options that
// allow modification of the internals of TextSimilarity
// before initialization. They are optional, and not using them
// allows you to use the defaults.
type Option func(TextSimilarity) TextSimilarity

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
	tf := float64(count(v, tokens)) / float64(documentFrequency[v])
	idf := math.Log(float64(n) / (float64(documentFrequency[v])))
	return tf * idf
}

func union(a, b []string) []string {
	m := make(map[string]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; !ok {
			a = append(a, item)
		}
	}
	return a
}

func minMaxKvSlice(s []kv) (min, max float64) {
	min = math.Inf(0)
	max = math.Inf(-1)
	for _, v := range s {
		max = math.Max(v.Value, max)
		min = math.Min(v.Value, min)
	}
	return min, max
}

func filter(vs []kv, f func(kv) bool) []kv {
	var vsf []kv
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func bigrams(g []string) []string {
	var b []string

	for i := 0; i < len(g)-1; i++ {
		b = append(b, strings.Join(g[i:i+2], " "))
	}

	return b
}

// New accepts a slice of documents and
// creates the internal corpus and document frequency mapping.
func New(documents []string, options ...Option) *TextSimilarity {
	var (
		allTokens []string
	)

	ts := TextSimilarity{
		stopwords:  stopbytes,
		documents:  documents,
		useBiGrams: false,
	}

	for _, option := range options {
		ts = option(ts)
	}

	ts.documentFrequency = map[string]int{}

	for _, doc := range documents {
		allTokens = append(allTokens, ts.Tokenize(doc)...)
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

// Tokenize splits the string into tokens and filters based on
// our custom stopword list.
func (ts *TextSimilarity) Tokenize(s string) []string {
	// Iterate over the doc's tokens:
	tokens := []string{}
	result := []string{}

	wordfn := func(word []byte) {
		tokens = append(tokens, string(word))
	}

	lowercase, stripAccents, stripContractions, stripNumbers, stripForeign := true, true, true, true, true
	tokenize.AllInOne([]byte(s), wordfn, lowercase, stripAccents, stripContractions, stripNumbers, stripForeign)

	// Filter Tokens using stopwords list
	for _, tok := range tokens {
		var exclude = false
		for _, v := range ts.stopwords {
			if string(v) == tok {
				exclude = true
			}
		}
		if exclude == false {
			result = append(result, tok)
		}
	}

	if ts.useBiGrams {
		result = append(result, bigrams(result)...)
	}

	return result
}

// Similarity returns the cosine similarity between two documents using
// Tf-Idf vectorization using the corpus.
func (ts *TextSimilarity) Similarity(a, b string) (float64, error) {
	tokensA := ts.Tokenize(a)
	tokensB := ts.Tokenize(b)
	combinedTokens := union(tokensA, tokensB)
	// Populate the vectors using frequency in the corpus.
	n := len(combinedTokens)
	vectorA := make([]float64, n)
	vectorB := make([]float64, n)
	for k, v := range combinedTokens {
		vectorA[k] = tfidf(v, tokensA, n, ts.documentFrequency)
		vectorB[k] = tfidf(v, tokensB, n, ts.documentFrequency)
	}

	similarity, err := Cosine(vectorA, vectorB)
	if err != nil {
		return 0.0, err
	}
	return similarity, nil
}

// Keywords accepts thresholds, which can be used to filter keyswords that
// are either they are too common or too unique and returns a sorted list of
// keywords (index 0 being the lower tf-idf score). Play with the thresholds
// according to your corpus.
func (ts *TextSimilarity) Keywords(threshLower, threshUpper float64) []string {
	var (
		docKeywords = []kv{}
		result      = []string{}
	)

	for _, doc := range ts.documents {
		tokens := ts.Tokenize(doc)
		n := len(tokens)
		mapper := map[string]float64{}

		for _, v := range tokens {
			val := tfidf(v, tokens, n, ts.documentFrequency)
			mapper[v] = val
		}

		// Convert to a kv pair for convenience.
		i := 0
		vector := make([]kv, len(mapper))
		for k, v := range mapper {
			vector[i] = kv{
				Key:   k,
				Value: v,
			}
			i++
		}

		// Filter tf-idf, using threshold.
		vector = filter(vector, func(v kv) bool {
			return v.Value >= threshLower && v.Value <= threshUpper
		})

		// Select the most common words relative to the corpus for this doc.
		min, _ := minMaxKvSlice(vector)
		vector = filter(vector, func(v kv) bool {
			return (v.Value == min)
		})

		docKeywords = append(docKeywords, vector...)
	}

	// Sort the vector based on tf-idf scores
	sort.Slice(docKeywords, func(i, j int) bool {
		return docKeywords[i].Value < docKeywords[j].Value
	})

	// Convert back to slice.
	for _, word := range docKeywords {
		result = append(result, word.Key)
	}
	return result
}

// Customization Options

// WithCustomStopwords can be used to replace the stopwords with
// a custom list of stopwords.
// eg.
//
//    ts := New(test_corpus, textsimilarity. WithCustomStopwords([][]byte{
//      []byte(`hello`),
//      []byte(`world`),
//    })
func WithCustomStopwords(wordList [][]byte) Option {
	return func(s TextSimilarity) TextSimilarity {
		s.stopwords = wordList
		return s
	}
}

// WithExtraStopwords can be used to augment stopwords with
// a custom list of stopwords.
// eg.
//
//    ts := New(test_corpus, textsimilarity.WithExtraStopwords([][]byte{
//      []byte(`hello`),
//    }))
func WithExtraStopwords(wordList [][]byte) Option {
	return func(s TextSimilarity) TextSimilarity {
		s.stopwords = append(s.stopwords, wordList...)
		return s
	}
}

// WithBiGrams can be used to augment tokens with
// bigrams along with 1-grams.
// eg.
//
//    ts := New(test_corpus, textsimilarity.WithBiGrams())
func WithBiGrams() Option {
	return func(s TextSimilarity) TextSimilarity {
		s.useBiGrams = true
		return s
	}
}
