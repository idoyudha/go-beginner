package main

import "fmt"

func main() {
	// runApp(0)
	runPanic(false)
	runPanic(true)
}

func logging() {
	fmt.Println("Call logging!")
}

func runApp(value int) {
	defer logging() // run this line whatever the runApp condition (error/success)
	fmt.Println("Run app!")
	result := 10 / value
	fmt.Println("Result", result)
}

func endApp() {
	message := recover() // catch error message and then stop panic -> run app after panic line
	if message != nil {
		fmt.Println("Error with message: ", message)
	}
	fmt.Println("End App!")
}

func runPanic(error bool) {
	defer endApp() // recover must be write in defer function
	fmt.Println("App running!")
	if error {
		panic("ERROR")
	}
	// fmt.Println("App running!")
}
