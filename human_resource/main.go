package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


type Employee struct {
	name string
	age int
	dept string
	position string
	salary int
}

func DisplayEmployee(e *Employee){
	fmt.Printf(`---Employee Details:---
Name: %s
Age: %d
Department: %s
Position: %s
Salary: %d
`, e.name, e.age, e.dept, e.position, e.salary)
}

func main() {
	for {
		fmt.Print("Employee Info (one line): ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		employeeText := strings.TrimSpace(scanner.Text())
		if strings.ToLower(strings.TrimSpace(employeeText)) == "exit" {
			fmt.Println("Bye Bye! Will miis you!")
			break
		}

		comps := strings.Split(employeeText, " ")

		var parts []string
		for _, e := range comps{
			e = strings.TrimSpace(e)
			if len(e) != 0 {
				parts = append(parts, e)
			}
		}

		if len(parts) != 5 {
			fmt.Println("Wrong number of input found: ", len(parts), parts)
			continue
		}

		name := parts[0]
		age, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("Wrong age input. Input an integer: ", age)
		}
		dept := parts[2]
		position := parts[3]
		salary, err := strconv.Atoi(parts[4])
		if err != nil {
			fmt.Println("Wrong salary input. Input an integer: ", salary)
		}

		employee := Employee{
			name:     name,
			age:      age,
			dept:     dept,
			position: position,
			salary:   salary,
		}
		DisplayEmployee(&employee)
	}
}
