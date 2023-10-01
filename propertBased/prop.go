package prop

import (
	"strings"
)

type (
	romanNumrals []romanNumral

	romanNumral struct {
		Value  int
		Symbol string
	}

	windowedRoman string
)

func (w windowedRoman) Symbols() (symbols [][]byte) {

	for i := 0; i < len(w); i++ {
		symbol := w[i]
		notAtEnd := i+1 < len(w)
		if notAtEnd && isSubtractive(symbol) && allRomanNumrals.Exists(symbol, w[i+1]) {
			symbols = append(symbols, []byte{symbol, w[i+1]})
			i++
		} else {
			symbols = append(symbols, []byte{symbol})
		}
	}
	return
}

func isSubtractive(symbol uint8) bool {
	return symbol == 'I' || symbol == 'X' || symbol == 'C'
}

var allRomanNumrals = romanNumrals{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func (r romanNumrals) ValueOf(symbols ...byte) int {
	symbol := string(symbols)
	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}
	return 0
}

func (r romanNumrals) Exists(symbols ...byte) bool {
	symbol := string(symbols)
	for _, s := range r {
		if s.Symbol == symbol {
			return true
		}
	}
	return false
}

func ConvertToRoman(number int) (romanNumber string) {
	var result strings.Builder

	for _, numral := range allRomanNumrals {
		for number >= numral.Value {
			result.WriteString(numral.Symbol)
			number = number - numral.Value
		}
	}
	romanNumber = result.String()
	return romanNumber
}

func ConvertToArabic(roman string) (total int) {

	for _, symbols := range windowedRoman(roman).Symbols() {
		total += allRomanNumrals.ValueOf(symbols...)
	}
	return
}

// func ConvertToArabic(roman string) int {
// 	total := 0
//
// 	for i := 0; i < len(roman); i++ {
// 		symbol := roman[i]
//
// 		if couldBeSubtractive(i, symbol, roman) {
// 			nextSymbol := roman[i+1]
//
// 			// build the two character string
// 			potentialNumber := string([]byte{symbol, nextSymbol})
//
// 			// get the value of the two character string
// 			value := allRomanNumrals.ValueOf(potentialNumber)
//
// 			if value != 0 {
// 				total += value
// 				i++ // move past this character too for the next loop
// 			} else {
// 				total++
// 			}
// 		} else {
// 			total++
// 		}
// 	}
// 	return total
// }
//
// func (r RomanNumrals) ValueOf(symbol string) int {
// 	for _, s := range r {
// 		if s.Symbol == symbol {
// 			return s.Value
// 		}
// 	}
// 	return 0
// }
//
// func couldBeSubtractive(index int, currentSymbol uint8, roman string) bool {
// 	return index+1 < len(roman) && currentSymbol == 'I'
// }
