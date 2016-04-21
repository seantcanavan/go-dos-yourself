package main

import "fmt"
import "os"
import "bufio"
import "runtime"
import "github.com/seantcanavan/threads"
//import "github.com/seantcanavan/go-dos-yourself/jobs"

func main() {

	req_params := 1
	param_names := [1]string{"Target IP"};
	threads.PrintItem()

	os.Args = os.Args[1:] // delete the first goland argument which is useless. significantly cleans up later code.

	fmt.Println("Go DOS Yourself.")
	fmt.Println("Found ", len(os.Args), " arguments")

	if len(os.Args) < req_params {
		fmt.Println("Invalid number of parameters. Expected: ", req_params, " Found: ", len(os.Args))
		os.Exit(1)
	}

	for index, element := range os.Args {
		fmt.Println("Index: ", index, "Param: ", param_names[index], " Argument: ", element)
	}

	fmt.Println("Main Menu")
	fmt.Println("Detected OS: ", runtime.GOOS)

	next_input := ""
	reader := bufio.NewReader(os.Stdin)

	for next_input != "7" {
		fmt.Println("\n\n\n")
		fmt.Println("1) Spin up a thread (automatically executes default job)\n")
		fmt.Println("2) Spin down a thread (closes gracefully)\n")
		fmt.Println("3) Change the default job\n")
		fmt.Println("4) Change job for a single thread\n")
		fmt.Println("5) Change jobs for all threads\n")
		fmt.Println("6) List all active threads and their associated jobs\n")
		fmt.Println("7) Exit\n")
		fmt.Println("Enter your input:\n")
		text, error := reader.ReadString('\n')
		//TODO(canavan) add better error handling
		if error != nil {
			continue
		}
		fmt.Println("Detected your input: ", text)
	}
}
