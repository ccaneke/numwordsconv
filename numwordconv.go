package numwordsconv

var Thousands map[int]string = map[int]string{1000: "one thousand", 2000: "two thousand", 3000: "three thousand", 4000: "four thousand",
	5000: "five thousand", 6000: "six thousand", 7000: "seven thousand", 8000: "eight thousand", 9000: "nine thousand"}

var Hundreds map[int]string = map[int]string{100: "one hundred", 200: "two hundred", 300: "three hundred", 400: "four hundred", 500: "five hundred",
	600: "six hundred", 700: "seven hundred", 800: "eight hundred", 900: "nine hundred"}

var Tens map[int]string = map[int]string{10: "ten", 20: "twenty", 30: "thirty", 40: "forty", 50: "fifty", 60: "sixty", 70: "seventy", 80: "eighty",
	90: "ninety"}

var UptoNineteen map[int]string = map[int]string{1: "one", 2: "two", 3: "three", 4: "four", 5: "five", 6: "six", 7: "seven",
	8: "eight", 9: "nine", 10: "ten", 11: "eleven", 12: "twelve", 13: "thirteen", 14: "fourteen", 15: "fifteen", 16: "sixteen", 17: "seventeen",
	18: "eighteen", 19: "nineteen"}

func IntegerToWords(n int) string {

	var result string
	var integer int
	if n > 999 {
		integer = n / 1000
		result += Thousands[integer*1000]

		if n < 1100 {
			result += " and "
		} else {
			result += " "
		}
		var remainder int = n % 1000
		if remainder >= 100 {
			integer = remainder / 100
			result += Hundreds[integer*100]
			result += " and "
			remainder = remainder % 100
			integer = remainder / 10
			result += Tens[integer*10]
			result += " "
			remainder = remainder % 10
			result += UptoNineteen[remainder]
		} else if remainder > 19 {
			integer = remainder / 10
			result += Tens[integer*10]
			result += " "
			var remainder = remainder % 10
			result += UptoNineteen[remainder]
		} else {
			result += UptoNineteen[remainder]
		}
	} else if n > 99 {
		// E.g. 126
		integer = n / 100
		result += Hundreds[integer*100]
		result += " and "
		var remainder int = n % 100
		if remainder > 19 {
			integer = remainder / 10
			result += Tens[integer*10]
			result += " "
			remainder = remainder % 10
			result += UptoNineteen[remainder]
		} else {
			result += UptoNineteen[remainder]
		}
	} else if n > 19 {
		// E.g. 56
		integer = n / 10
		result += Tens[integer*10]
		result += " "
		var remainder int = n % 10
		result += UptoNineteen[remainder]
	} else {
		result += UptoNineteen[n]
	}

	return result
}
