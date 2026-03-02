package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("salary.csv")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(file), "\n")

	var totalSalary, count, totalAge int

	for i, line := range lines {
		if i == 0 || line == "" {
			continue
		}

		columns := strings.Split(line, ",")

		salaryStr := columns[2]
		salaryStr = strings.TrimSpace(salaryStr)

		ageStr := columns[1]
		ageStr = strings.TrimSpace(ageStr)

		salary, err := strconv.Atoi(salaryStr)
		if err != nil {
			fmt.Println("Ошибка конвертации в строке:", i, err)
			continue
		}

		age, err := strconv.Atoi(ageStr)
		if err != nil {
			fmt.Println("Ошибка конвертации в строке:", i, err)
			continue
		}

		totalSalary += salary
		totalAge += age
		count++

	}

	if count > 0 {
		fmt.Printf("Всего сотрудников: %d\n", count)
		fmt.Printf("Средний возраст сотрудников: %d\n", totalAge/count)
		fmt.Printf("Общий возраст сотрудников: %d\n", totalAge)
		fmt.Printf("Общая сумма зарплат: %d\n", totalSalary)
		fmt.Printf("Средняя зарплата: %d\n", totalSalary/count)
	}
}
