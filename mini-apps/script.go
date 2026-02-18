package main

import (
	"cmp"
	"fmt"
	"slices"
)

const (
	Reset  = "\033[0m"
	Bold   = "\033[1m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
)

func main() {

	// fmt.Printf("%sHello This is Red%s\n", Red, Reset)
	// fmt.Printf("%sHello This is Green%s\n", Green, Reset)
	// fmt.Println(Bold + Yellow + "Hello this is Yellow" + Reset)

	// price := 250.45555
	// 250.46
	// strFormatedPrice := fmt.Sprintf("%.2f", price)
	// fmt.Println(strFormatedPrice)

	// slog.SetLogLoggerLevel(slog.LevelDebug)
	// logger := slog.New(
	// 	slog.NewJSONHandler(os.Stdout, nil),
	// )
	// logger.Debug("hello this is debug log")
	// logger.Info("hello this is info log")
	// logger.Warn("hi this is warning log")
	// log.Println("ok")

	// port := cmp.Or(
	// 	getFromEnv(),
	// 	getPortFromFlag(),
	// 	"7070",
	// )
	// fmt.Println("application start on port:", port)

	// ternary operator
	// age := 22
	// canDrink := func() string {
	// 	if age >= 18 {
	// 		return "Yes"
	// 	}
	// 	return "No"
	// }()
	// fmt.Println("can drink ?", canDrink)

	// userName := cmp.Or(
	// 	getUserNameFromInput(),
	// 	"Anonymous",
	// )
	// fmt.Printf("Hello %s welcome to secret panel :)\n", userName)

	employees := []employee{
		{name: "Pouya", age: 23},
		{name: "Ali", age: 22},
		{name: "Mona", age: 18},
		{name: "Hassan", age: 35},
		{name: "Behnam", age: 22},
	}
	fmt.Println("Employess before sorting:")
	fmt.Println(employees)

	fmt.Println("Employess after sorting:")
	fmt.Println(sortEmployees(employees))

}

type employee struct {
	name string
	age  uint
}

func sortEmployees(employees []employee) []employee {
	sortedEmployess := employees
	slices.SortFunc(
		sortedEmployess,
		func(a, b employee) int {
			return cmp.Or(
				cmp.Compare(a.age, b.age),
				cmp.Compare(a.name, b.name),
			)
		},
	)

	return sortedEmployess
}

// func getPortFromFlag() string {
// 	return "8080"
// }

// func getFromEnv() string {
// 	return os.Getenv("PORT")
// }

// func getUserNameFromInput() string {
// 	fmt.Print("Enter your name (or press enter to use 'Anonymous'): ")
// 	var name string
// 	fmt.Scanln(&name)
// 	return strings.TrimSpace(name)
// }
