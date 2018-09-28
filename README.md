# Text Similarity

[![GoDoc](https://godoc.org/github.com/rhnvrm/textsimilarity?status.svg)](https://godoc.org/github.com/rhnvrm/textsimilarity)

A package that provides similarity between two string documents
using cosine similarity and tf-idf along with various
other useful things.

## Usage

```go
    docs := []string{
        "Samsung Galaxy Tab A 10.5 review: Slate will impress multimedia enthusiasts",
        "Why banks are upgrading your debit, credit cards",
        "Investments in HDFC AMC shares are subject to regulatory risks",
        "Along with equity, it's time to put money in fixed income, gold and real estate: UR Bhat, Dalton Capital Advisors",
        "Global gold prices edge up as easing trade concerns hurt dollar",
        "Global Markets: Shares inch up as trade woes take backseat to buoyant U.S. markets",
        "Gold prices edge up as easing trade concerns hurt dollar",
        "Bearish bets remain as edgy investors retreat from risky Asian currencies: Reuters poll",
        "Asian currencies firm as investors look past trade concerns",
        "BSE, NSE to foray into commodity derivatives from Oct 1, to start with gold",
        "Buy ITC, target Rs 382: HDFC Securities",
        "Buy JSW Steel, target Rs 484: Nomura, India",
    }

    ts := textsimilarity.New(docs)

    docA := "Gold prices edge up as easing trade concerns hurt dollar"
    docB := "Global gold prices edge up as easing trade concerns hurt dollar"
    result, _ := ts.Similarity(docA, docB)

    keywords := ts.Keywords(0.2, 0.5)

    fmt.Println(result)
    fmt.Println(keywords)
```