package main

var (
	romanNums   = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	romanDozens = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
)

func intToRoman(n int) string {
	if n == 100 {
		return "C"
	}
	return romanDozens[n/10] + romanNums[n%10]
}

func romanToInt(s string) (int, error) {
	switch s {
	case "I":
		return 1, nil
	case "II":
		return 2, nil
	case "III":
		return 3, nil
	case "IV":
		return 4, nil
	case "V":
		return 5, nil
	case "VI":
		return 6, nil
	case "VII":
		return 7, nil
	case "VIII":
		return 8, nil
	case "IX":
		return 9, nil
	case "X":
		return 10, nil
	default:
		return 0, errWrongInput
	}
}
