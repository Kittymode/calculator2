package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// карта для перевода арабских чисел в римские
var arabToRome = map[int]string{
	1:  "I",
	2:  "II",
	3:  "III",
	4:  "IV",
	5:  "V",
	6:  "VI",
	7:  "VII",
	8:  "VIII",
	9:  "IX",
	10: "X",
}

// карта для перевода римских чисел в арабские
var romeToArab = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
}

// функция преобразования арабских чисел в римские
func arabicToRoman(num int) string {
	if num >= 1 && num <= 10 {
		return arabToRome[num]
	}
	var result strings.Builder
	values := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	numerals := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	for i := 0; i < len(values); i++ {
		for num >= values[i] {
			result.WriteString(numerals[i])
			num -= values[i]
		}
	}
	return result.String()
}

// функция для проверки, является ли строка римским числом
func isRoman(input string) bool {
	_, exist := romeToArab[input]
	return exist
}

// функция для выполнения операции над двумя числами
func calc(a, b int, operator string) int {
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("Ошибка: неизвестная операция")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите выражение в формате: 1 + 1 или I + II")

	// чтение строки целиком
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(fmt.Sprintf("Ошибка при чтении ввода: %v", err))
	}

	// убираем символ новой строки и лишние пробелы
	input = strings.TrimSpace(input)

	// разделяем строку по пробелам для получения операндов и оператора
	parts := strings.Split(input, " ")
	if len(parts) != 3 {
		panic("Ошибка: неверный формат ввода")
	}

	// проверка, являются ли числа римскими
	isRomanInput := isRoman(parts[0]) && isRoman(parts[2])

	var a, b int
	var err1, err2 error
	operator := parts[1]

	if isRomanInput {
		// преобразуем римские числа в арабские
		a = romeToArab[parts[0]]
		b = romeToArab[parts[2]]
	} else {
		// преобразуем арабские числа
		a, err1 = strconv.Atoi(parts[0])
		b, err2 = strconv.Atoi(parts[2])
		if err1 != nil || err2 != nil {
			panic("Ошибка: неверный формат чисел")
		}
	}

	// проверка на диапазон чисел
	if a < 1 || b < 1 {
		panic("Ошибка: числа должны быть в диапазоне от 1 до 10")
	}

	// выполняем операцию
	result := calc(a, b, operator)

	// выводим результат
	if isRomanInput {
		// проверка на возможность представления результата в римских числах
		if result <= 0 {
			panic("Ошибка: результат в римских числах должен быть положительным")
		}
		fmt.Println("Результат:", arabicToRoman(result))
	} else {
		fmt.Println("Результат:", result)
	}
}
