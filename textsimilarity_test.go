package textsimilarity

import (
	"testing"
)

func TestSimilarity(t *testing.T) {

	t.Run("same", func(t *testing.T) {
		docA := "India's largest hotel chain OYO to expand into Britain"
		docB := "India's largest hotel chain OYO to expand into Britain"

		result, _ := Similarity(docA, docB)

		if result != 1.0 {
			t.Errorf("Similarity() did not return 1")
		}
	})

}
