# Text Similarity

A naive package that provides similarity between two string documents
using cosine similarity.

## Usage 
```go
docA := "Danske Bank boss quits over $234 billion money laundering scandal"
docB := "Danske Bank CEO quits in $234 billion money laundering scandal"
result, err := Similarity(docA, docB)
```