package main

import "fmt"
import "os"

func main() {
	required_params := 2
	param_names := string[]{"test"}
	fmt.Println("Go DOS Yourself.")
	fmt.Println("Found ", len(os.Args), " arguments")
	if len(os.Args) < req_params {
		fmt.Println("Invalid number of parameters. Expected: ", req_params, " Found: ", len(os.Args))
		os.Exit(1)
	}
	for index, element := range os.Args {
		fmt.Println("Index: ", index, "Param: ", param_names[index], " Argument: ", element)
	}
	fmt.Println("hello world")
}
