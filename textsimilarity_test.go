package textsimilarity

import (
	"reflect"
	"testing"
)

func TestSimilarity(t *testing.T) {

	ts := New([]string{
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
	})

	t.Run("same", func(t *testing.T) {
		cases := []struct {
			docA, docB string
			result     float64
		}{
			{
				docA:   "BSE, NSE to foray into commodity derivatives from Oct 1, to start with gold",
				docB:   "BSE, NSE to foray into commodity derivatives from Oct 1, to start with gold",
				result: 1.0,
			},
			{
				docA:   "Gold prices edge up as easing trade concerns hurt dollar",
				docB:   "Global gold prices edge up as easing trade concerns hurt dollar",
				result: 0.9,
			},
		}

		for _, tc := range cases {
			result, _ := ts.Similarity(tc.docA, tc.docB)
			if result < tc.result {
				t.Errorf("Similarity(%v,%v) did not return %v, got = %v", tc.docA, tc.docB, tc.result, result)
			}
		}

	})

}

func TestTokenize(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test1",
			args: args{
				s: "Hello to the World",
			},
			want: []string{"hello", "world"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Tokenize(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tokenize() = %v, want %v", got, tt.want)
			}
		})
	}
}
