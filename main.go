package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const msgMain = "Введите пример:"

func main() {
	sc := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println(msgMain)

		sc.Scan()

		s := strings.ReplaceAll(sc.Text(), " ", "")

		n1, n2, operator, isRoman, err := checkString(s)

		if err != nil {
			panic(err)
		}

		if err = checkLimits(n1, n2, operator, isRoman); err != nil {
			panic(err)
		}

		res := count(n1, n2, operator)

		if isRoman {
			fmt.Println(intToRoman(res))
		} else {
			fmt.Println(res)
		}
	}
}
