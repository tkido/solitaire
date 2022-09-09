package main

import (
	"fmt"
	"testing"
)

func TestCardsToString(t *testing.T) {
	fmt.Println(NULL)
	fmt.Println(W1)
	fmt.Println(G2)
	fmt.Println(R3)
	fmt.Println(FL)
	fmt.Println(M1)
	fmt.Println(P2)
	fmt.Println(S3)
	fmt.Println(M7)
	fmt.Println(P8)
	fmt.Println(S9)
	fmt.Println(cardSet)

}

func TestBitString(t *testing.T) {
	cases := map[string]struct {
		Card Card
		Want string
	}{
		"NULL": {NULL, "0000000000000000000000000000000000000000000000000000000000000000"},
		"W1":   {W1, "1000000000000000000000000000000000000000000000000000000000000000"},
		"M8":   {M8, "0000000000000000000000000000000000100000000000000000000000011000"},
		"M9":   {M9, "0000000000000000000000000000000000000100000000000000000000000011"},
		"S9":   {S9, "0000000000000000000000000000000000000001000000000000000000000110"},
	}
	for name, tt := range cases {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := tt.Card.BitString()
			if got != tt.Want {
				t.Errorf("want = %v, but got = %v", tt.Want, got)
			}
		})
	}
}

func TestIsOver(t *testing.T) {
	cases := map[string]struct {
		A    Card
		B    Card
		Want bool
	}{
		"S4/W1": {S4, W1, false},
		"M5/W1": {M5, W1, false},
		"W1/S4": {W1, S4, false},
		"W1/M5": {W1, M5, false},
		"M1/M1": {M1, M1, false},
		"M1/P2": {M1, P2, false},
		"M1/S2": {M1, S2, false},
		"M2/P3": {M2, P3, true},
		"M2/S3": {M2, S3, true},
		"P8/M9": {P8, M9, true},
		"P8/S9": {P8, S9, true},
		"M9/P1": {M9, P1, false},
		"M9/S1": {M9, S1, false},
	}

	for name, tt := range cases {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := tt.A.IsOver(tt.B)
			if got != tt.Want {
				t.Errorf("want = %v, but got = %v", tt.Want, got)
			}
		})
	}
}
