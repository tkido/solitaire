package main

import "fmt"

type Card uint64

func (c Card) IsOver(o Card) bool {
	if c <= S1 {
		return false
	}
	return c<<27&o != NULL
}

func (c Card) BitString() string {
	s := Reverse(fmt.Sprintf("%b", c)) + "00000000000000000000000000000000000000000000000000000000000000000"
	return s[:64]
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

const (
	NUM_NULL int = iota // 空白
	NUM_W1              // WHITE 白
	NUM_W2
	NUM_W3
	NUM_W4
	NUM_G1 // GREEN 発
	NUM_G2
	NUM_G3
	NUM_G4
	NUM_R1 // RED 中
	NUM_R2
	NUM_R3
	NUM_R4
	NUM_FL // FLOWER 花
	NUM_M1 // MANTSU
	NUM_P1 // PINTSU
	NUM_S1 // SOWTSU
	NUM_M2
	NUM_P2
	NUM_S2
	NUM_M3
	NUM_P3
	NUM_S3
	NUM_M4
	NUM_P4
	NUM_S4
	NUM_M5
	NUM_P5
	NUM_S5
	NUM_M6
	NUM_P6
	NUM_S6
	NUM_M7
	NUM_P7
	NUM_S7
	NUM_M8
	NUM_P8
	NUM_S8
	NUM_M9
	NUM_P9
	NUM_S9
)

const (
	BIT_W1 Card = 1 << iota
	BIT_W2
	BIT_W3
	BIT_W4
	BIT_G1
	BIT_G2
	BIT_G3
	BIT_G4
	BIT_R1
	BIT_R2
	BIT_R3
	BIT_R4
	BIT_FL
	BIT_M1
	BIT_P1
	BIT_S1
	BIT_M2
	BIT_P2
	BIT_S2
	BIT_M3
	BIT_P3
	BIT_S3
	BIT_M4
	BIT_P4
	BIT_S4
	BIT_M5
	BIT_P5
	BIT_S5
	BIT_M6
	BIT_P6
	BIT_S6
	BIT_M7
	BIT_P7
	BIT_S7
	BIT_M8
	BIT_P8
	BIT_S8
	BIT_M9
	BIT_P9
	BIT_S9
	BIT_NULL uint64 = 0 // 空白
)

const (
	W1 Card = 1 << iota
	W2
	W3
	W4
	G1
	G2
	G3
	G4
	R1
	R2
	R3
	R4
	FL
	M1
	P1
	S1
	M2
	P2
	S2
	M3   Card = BIT_M3 | BIT_P2<<27 | BIT_S2<<27
	P3   Card = BIT_P3 | BIT_S2<<27 | BIT_M2<<27
	S3   Card = BIT_S3 | BIT_M2<<27 | BIT_P2<<27
	M4   Card = BIT_M4 | BIT_P3<<27 | BIT_S3<<27
	P4   Card = BIT_P4 | BIT_S3<<27 | BIT_M3<<27
	S4   Card = BIT_S4 | BIT_M3<<27 | BIT_P3<<27
	M5   Card = BIT_M5 | BIT_P4<<27 | BIT_S4<<27
	P5   Card = BIT_P5 | BIT_S4<<27 | BIT_M4<<27
	S5   Card = BIT_S5 | BIT_M4<<27 | BIT_P4<<27
	M6   Card = BIT_M6 | BIT_P5<<27 | BIT_S5<<27
	P6   Card = BIT_P6 | BIT_S5<<27 | BIT_M5<<27
	S6   Card = BIT_S6 | BIT_M5<<27 | BIT_P5<<27
	M7   Card = BIT_M7 | BIT_P6<<27 | BIT_S6<<27
	P7   Card = BIT_P7 | BIT_S6<<27 | BIT_M6<<27
	S7   Card = BIT_S7 | BIT_M6<<27 | BIT_P6<<27
	M8   Card = BIT_M8 | BIT_P7<<27 | BIT_S7<<27
	P8   Card = BIT_P8 | BIT_S7<<27 | BIT_M7<<27
	S8   Card = BIT_S8 | BIT_M7<<27 | BIT_P7<<27
	M9   Card = BIT_M9 | BIT_P8<<27 | BIT_S8<<27
	P9   Card = BIT_P9 | BIT_S8<<27 | BIT_M8<<27
	S9   Card = BIT_S9 | BIT_M8<<27 | BIT_P8<<27
	NULL Card = 0 // 空白
)

func (c Card) String() string {
	switch c {
	case NULL:
		return "[]"
	case W1, W2, W3, W4:
		return "白"
	case G1, G2, G3, G4:
		return "発"
	case R1, R2, R3, R4:
		return "中"
	case FL:
		return "花"
	case M1:
		return "M1"
	case M2:
		return "M2"
	case M3:
		return "M3"
	case M4:
		return "M4"
	case M5:
		return "M5"
	case M6:
		return "M6"
	case M7:
		return "M7"
	case M8:
		return "M8"
	case M9:
		return "M9"
	case P1:
		return "P1"
	case P2:
		return "P2"
	case P3:
		return "P3"
	case P4:
		return "P4"
	case P5:
		return "P5"
	case P6:
		return "P6"
	case P7:
		return "P7"
	case P8:
		return "P8"
	case P9:
		return "P9"
	case S1:
		return "S1"
	case S2:
		return "S2"
	case S3:
		return "S3"
	case S4:
		return "S4"
	case S5:
		return "S5"
	case S6:
		return "S6"
	case S7:
		return "S7"
	case S8:
		return "S8"
	case S9:
		return "S9"
	default:
		return "MUST NOT HAPPEN!!"
	}
}

var cardSet = []Card{
	W1,
	W2,
	W3,
	W4,
	G1,
	G2,
	G3,
	G4,
	R1,
	R2,
	R3,
	R4,
	FL,
	M1,
	P1,
	S1,
	M2,
	P2,
	S2,
	M3,
	P3,
	S3,
	M4,
	P4,
	S4,
	M5,
	P5,
	S5,
	M6,
	P6,
	S6,
	M7,
	P7,
	S7,
	M8,
	P8,
	S8,
	M9,
	P9,
	S9,
}
