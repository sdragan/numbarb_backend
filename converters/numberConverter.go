package converters

import "strings"

const outOfRangeMessage = "out of range"
const minConvertedInteger = 1
const maxConvertedInteger = 999999

var m map[int]string

func initMap() {
	m = make(map[int]string)
	m[0] = ""
	m[1] = "one"
	m[2] = "two"
	m[3] = "three"
	m[4] = "four"
	m[5] = "five"
	m[6] = "six"
	m[7] = "seven"
	m[8] = "eight"
	m[9] = "nine"
	m[10] = "ten"
	m[11] = "eleven"
	m[12] = "twelve"
	m[13] = "thirteen"
	m[14] = "fourteen"
	m[15] = "fifteen"
	m[16] = "sixteen"
	m[17] = "seventeen"
	m[18] = "eighteen"
	m[19] = "nineteen"
	m[20] = "twenty"
	m[30] = "thirty"
	m[40] = "forty"
	m[50] = "fifty"
	m[60] = "sixty"
	m[70] = "seventy"
	m[80] = "eighty"
	m[90] = "ninety"
}

// ConvertIntegerToWords converts integer between 1 and 999999 to words
func ConvertIntegerToWords(number int) string {
	initMap()

	if number < minConvertedInteger || number > maxConvertedInteger {
		return outOfRangeMessage
	}

	result := ""
	result += threeDigitNumberToWords(number/1000, "thousand", false)
	result += threeDigitNumberToWords(number%1000, "", number >= 1000)
	result = removeExcessiveWhitespaces(result)

	return result
}

func threeDigitNumberToWords(number int, groupName string, thereIsPrevious bool) string {
	if number == 0 {
		return ""
	}
	result := ""

	hundreds := hundredsToText(number)
	tens := tensToText(number)
	figures := figureToText(number)

	result = concatParts(hundreds, tens, figures, groupName, thereIsPrevious)

	return result
}

func hundredsToText(number int) string {
	if number > 99 {
		return m[number/100] + " hundred"
	}
	return ""
}

func tensToText(number int) string {
	rest := number % 100

	if rest >= 20 {
		return m[(rest/10)*10]
	} else if rest >= 10 {
		return m[rest]
	}

	return ""
}

func figureToText(number int) string {
	rest := number % 100
	if rest > 20 || rest < 10 {
		return m[number%10]
	}
	return ""
}

func concatParts(hundreds, tens, figures, groupName string, thereIsPrevious bool) string {
	result := ""
	if thereIsPrevious {
		result += " "
		if hundreds == "" && (tens != "" || figures != "") {
			result += "and "
		}
	}

	result += hundreds + " "
	if hundreds != "" && tens == "" && figures != "" {
		result += "and "
	}

	result += tens + " "
	result += figures + " "
	result += groupName + " "

	return result
}

func removeExcessiveWhitespaces(phrase string) string {

	for strings.Contains(phrase, "  ") {
		phrase = strings.ReplaceAll(phrase, "  ", " ")
	}

	phrase = strings.TrimPrefix(phrase, " ")
	phrase = strings.TrimSuffix(phrase, " ")

	return phrase
}
