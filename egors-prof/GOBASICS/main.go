package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	var err error
	var result int
	if b == 0 {
		err = errors.New("b==0\nна ноль делить нельзя")
	} else if a < 0 || b < 0 {
		err = errors.New("одно из чисел меньше нуля")

	} else {
		result = a / b

		switch result {
		case 0:
			fmt.Println("Результат маленький или ноль")

		case 1, 2, 3, 4, 5, 6, 7, 8, 9:
			fmt.Println("Результат средний")
		default:
			fmt.Println("Результат большой")
		}
	}
	return result, err
}
func main() {
	for {
		fmt.Println("Введите start,чтобы начать или exit, чтобы выйти")
		var userInput string
		_, err := fmt.Scanln(&userInput)
		if err != nil {
			fmt.Println("Повторите заново!")
			continue
		}
		if userInput == "start" {
			fmt.Println("Введите два числа")
			var (
				number1 int
				number2 int
			)
			_, err = fmt.Scanln(&number1, &number2)
			if err != nil {
				fmt.Println(err, "\nВозникла ошибка при вводе данных\nпроверьте введённое на правильность")
			}
			fmt.Println(divide(number1, number2))
		} else if userInput == "exit" {
			fmt.Println("Выход из программы")
			break
		} else {
			fmt.Println("Повторите заново")
		}

	}

}
