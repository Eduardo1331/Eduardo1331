package main

import (
	"bufio"
	"fmt"5
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	number1, operator, number2, isRoman, err := parseInput()
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	result, err := calculate(number1, operator, number2, isRoman)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	if isRoman {
		fmt.Printf("Результат: %s %s %s = %s\n", toRoman(number1), operator, toRoman(number2), toRoman(int(result)))
	} else {
		fmt.Printf("Результат: %d %s %d = %s\n", number1, operator, number2, formatResult(result))
	}
}

func parseInput() (int, string, int, bool, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите выражение: ")
	expression, _ := reader.ReadString('\n')
	expression = strings.TrimSpace(expression)

	re := regexp.MustCompile(`^(\w+)\s*([\+\-\*/])\s*(\w+)$`)
	matches := re.FindStringSubmatch(expression)

	if len(matches) != 4 {
		return 0, "", 0, false, fmt.Errorf("Введите выражение с двумя числами и одним из математических операторов! (+, -, /, *).")
	}

	isRoman1 := isRomanNumeral(matches[1])
	isRoman2 := isRomanNumeral(matches[3])

	if isRoman1 != isRoman2 {
		return 0, "", 0, false, fmt.Errorf("Используются одновременно разные системы счисления.")
	}

	isRoman := isRoman1 && isRoman2

	var number1, number2 int
	var err1, err2 error

	if isRoman {
		number1 = fromRoman(matches[1])
		number2 = fromRoman(matches[3])
	} else {
		number1, err1 = strconv.Atoi(matches[1])
		number2, err2 = strconv.Atoi(matches[3])
		if err1 != nil || err2 != nil {
			return 0, "", 0, false, fmt.Errorf("Введите арабские или римские цифры!")
		}
	}

	return number1, matches[2], number2, isRoman, nil
}

func calculate(number1 int, operator string, number2 int, isRoman bool) (float64, error) {
	var result int
	switch operator {
	case "+":
		result = number1 + number2
	case "-":
		result = number1 - number2
	case "*":
		result = number1 * number2
	case "/":
		if number2 == 0 {
			return 0, fmt.Errorf("Деление на ноль")
		}
		return float64(number1) / float64(number2), nil
	default:
		return 0, fmt.Errorf("Неизвестный оператор: %s", operator)
	}

	if isRoman && result < 0 {
		return 0, fmt.Errorf("В римской системе нет отрицательных чисел")
	}

	return float64(result), nil
}

func formatResult(result float64) string {
	if result == float64(int(result)) {
		return fmt.Sprintf("%d", int(result))
	}
	return fmt.Sprintf("%.2f", result)
}

func isRomanNumeral(s string) bool {
	re := regexp.MustCompile(`^[IVXLCDM]+$`)
	return re.MatchString(s)
}

func fromRoman(roman string) int {
	romanToInt := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total := 0
	prev := 0
	for _, char := range roman {
		curr := romanToInt[char]
		if curr > prev {
			total += curr - 2*prev
		} else {
			total += curr
		}
		prev = curr
	}
	return total
}

func toRoman(num int) string {
	if num <= 0 {
		panic("В римской системе нет отрицательных чисел или нуля")
	}

	val := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	syb := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	roman := ""
	for i := 0; i < len(val); i++ {
		for num >= val[i] {
			num -= val[i]
			roman += syb[i]
		}
	}
	return roman
}
