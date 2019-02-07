package luhn

import (
	"fmt"
	"strconv"
	"unicode"
)

func Valid(card string) bool {
	var secondDigits []int

	var formattedCard []int

	// format card to remove whitespace and anything not digit
	for _, num := range card {
		if unicode.IsDigit(num) {
			junk := string(num)
			numInt, _ := strconv.ParseInt(junk, 10, 64)

			formattedCard = append(formattedCard, int(numInt))
		} else if num != ' ' {
			return false
		}
	}

	// lengths of less than 1 are not valid
	if len(formattedCard) <= 1 {
		return false
	}

	fmt.Println("unformatted card: ", card)
	fmt.Println("formatted card: ", formattedCard)

	// double every other number while appending non doubled numbers
	for i, num := range formattedCard {

		if (len(card)-i)%2 == 0 {
			num += num

			fmt.Println("doubled: ", num)

			if num > 9 {
				num -= 9
			}
		}

		secondDigits = append(secondDigits, num)

		fmt.Println("appended digit: ", num)
	}

	fmt.Println("new card: ", secondDigits)

	var sum int

	for _, num := range secondDigits {
		sum += num

		fmt.Println(sum, num)
	}

	if sum%10 == 0 {
		return true
	}

	return false
}
