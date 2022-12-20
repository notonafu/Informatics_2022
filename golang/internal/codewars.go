package internal

import (
	"math"
	"strings"
)

func EvenOrOdd(num int64) string {
	if num%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}
func CountingSheeps(arrcs []bool) int {
	var csres int
	var i int
	for ; i < len(arrcs); i++ {
		if arrcs[i] {
			csres++
		}
	}
	return csres
}
func CountingMonkeys(n int) []int {
	var arrcm = make([]int, 0)
	for i := 1; i <= n; i++ {
		arrcm = append(arrcm, i)
	}
	return arrcm
}
func SchoolPaperwork(n int, m int) int {
	if (n >= 0) && (m >= 0) {
		return n * m
	} else {
		return 0
	}
}
func SchoolPaperwork1(n int, m int) int {
	if (n < 0) || (m < 0) {
		return 0
	} else {
		return m * n
	}
}
func IsHe(bullets int, dragons int) string {
	if bullets >= dragons*2 {
		return "True"
	} else {
		return "False"
	}
}

// ПОЛЬСКАЯ КОРОВА ЗАХОТЕЛА
func PolishA(text string) string {
	var changing [][]string = [][]string{
		{"ą", "a"},
		{"ć", "c"},
		{"ę", "e"},
		{"ł", "l"},
		{"ń", "n"},
		{"ó", "o"},
		{"ś", "s"},
		{"ź", "z"},
		{"ż", "z"},
	}
	for i := 0; i < len(changing); i++ {
		text = strings.Replace(text, changing[i][0], changing[i][1], 1)
	}
	return text
}
func PolishB(text string) string { //БЕЗ ПАКЕТА STRINGS
	var changingA string = "ąćęłńóśźż"
	var changingB string = "acelnoszz"
	conv := []byte(text)
	var count int8
	for i := 0; i < len(conv); i++ {
		for k := 0; k < len(changingA); k++ {
			count++
			j := k / 2
			if (count % 2) == 0 {
				if conv[i] == changingA[k] {
					conv[i] = changingB[j]
					conv = append(conv[:i-1], conv[i:]...)
				}
			}
		}
	}
	return string(conv)
}
func FindAll(arrFa []int, n int) []int {
	var res []int
	for i := 0; i < len(arrFa); i++ {
		if arrFa[i] == n {
			res = append(res, i)
		}
	}
	return res
}
func SumOfMin(arrA [][]int, m int, n int) int {
	var res int = 0
	for i := 0; i < (m); i++ {
		min := math.Inf(1)
		for k := 0; k < (n); k++ {
			if float64(arrA[i][k]) < min {
				min = float64(arrA[i][k])
			}
		}
		res = res + int(min)
	}
	return res
}

// text - string
// text[i] - byte