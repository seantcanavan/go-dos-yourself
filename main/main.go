package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/abiosoft/ishell"
	"github.com/seantcanavan/go-dos-yourself/config"
	"github.com/seantcanavan/go-dos-yourself/job"
	"github.com/seantcanavan/go-dos-yourself/routine"
	"github.com/seantcanavan/go-dos-yourself/target"
)

var configs []config.Config
var jobs []job.Job
var targets []target.Target
var routines []routine.Routine

func main() {
	// create new shell.
	// by default, new shell includes 'exit', 'help' and 'clear' commands.
	shell := ishell.New()

	// display welcome info.
	shell.Println(`
	go-dos-yourself: Network security and stress tester. Use at your own risk.
	You assume all liabilities of executing and automating large amounts of
	correct or incorrect network traffic at any target(s) you select.`)

	// register a function for "greet" command.
	shell.Register("load", load)
	shell.Register("set", set)
	shell.Register("print", printThings)
	shell.Register("start", start)
	shell.Register("stop", stop)
	shell.Register("pause", pause)

	// start shell
	shell.Start()
}

func load(args ...string) (string, error) {
	if len(args) != 2 {
		fmt.Println(`
		Incorrect arguments to load - returning
		Try 'load [command] [inputSource]'
		E.G. 'load targets targets.json'
		This will load the data from targets.json as a new target`)
		return "", errors.New("invalid parameters")
	}

	action := args[0]
	param := args[1]

	switch action {
	case "target":
		newTarget, err := target.TargetFromFile(param)
		if err != nil {
			return fmt.Sprintf("error loading new target object from file:", param), err
		}
		targets = append(targets, newTarget)
		fmt.Println("Successfully added target:", newTarget.Name)
	case "config":
		newConfig, err := config.ConfigFromFile(param)
		if err != nil {
			return fmt.Sprintf("error loading new config object from file:", param), err
		}
		configs = append(configs, newConfig)
		fmt.Println("Successfully added config:", param)
	default:
		fmt.Println("Unknown command:", action, "try again or refer to docs")
	}
	return "", nil
}

func printThings(args ...string) (string, error) {
	if len(args) != 1 {
		fmt.Println("Not enough arguments to print - returning")
		fmt.Println("Try 'print [command]'")
		fmt.Println("Example: 'print jobs'")
		fmt.Println("This will print all the current jobs that are running")
		fmt.Println("Example: 'print targets'")
		fmt.Println("This will print all the current targets that are available")
		return "", errors.New("invalid parameters")
	}

	target := args[0]

	switch target {
	case "configs":
		for index, element := range configs {
			fmt.Printf("%v) %+v \n", index, element)
		}
	case "targets":
		for index, element := range targets {
			fmt.Printf("%v) %+v \n", index, element)
		}
	case "routines":
		for index, element := range routines {
			fmt.Printf("%v) %+v \n", index, element)
		}
	case "jobs":
		for index, element := range jobs {
			fmt.Printf("%v) %+v \n", index, element)
		}
	default:
		fmt.Println("Unknown object to print", target)
		fmt.Println("Supported values are: configs, targets, routines, jobs")
	}
	return "", nil
}

func set(args ...string) (string, error) {
	if len(args) != 3 {
		fmt.Println("Not enough arguments to set - returning")
		fmt.Println("Try 'set [command] [parameter]'")
		fmt.Println("Example: 'set target 127.0.0.1'")
		fmt.Println("This will set the default routine target to 127.0.0.1")
	}
	return "", nil
}

// start will start the given routine with the given config and target
func start(args ...string) (string, error) {
	if len(args) != 3 {
		fmt.Println("Not enough arguments to start - returning")
		fmt.Println("Try 'start [routineName] [config#] [target#]'")
		fmt.Println("Example: 'start get 1 2'")
		fmt.Println("The above examples starts a get routine with the 1st (0 based) config loaded in the system and the 2nd (0th based) target loaded in the system")
		fmt.Println("You can use 'load' command to load a config or target from a file. You can use the 'print' command to print all values already loaded into the system")
		fmt.Println("Try 'start [routineName] [configPath] [targetPath]'")
		fmt.Println("Example 'start ping ../src/github.com/seantcanavan/config/samples/ping01.json ../src/github.com/seantcanavan/target/samples/ping_google01.json")
		fmt.Println("The above example starts a ping job and loads both the config and target directly from the given file.")
		return "", errors.New("invalid parameters")
	}

	routineName := args[0]
	configVal := args[1]
	targetVal := args[2]

	var r routine.Routine
	var c config.Config
	var t target.Target

	r, err := routine.GetRoutinebyName(routineName)
	if err != nil {
		return "", err
	}

	// first try to load the config as a number
	configNum, configErr := strconv.Atoi(configVal)
	// if we can't load the config as a number, load it as a file path
	if configErr != nil {
		c, configErr = config.ConfigFromFile(configVal)
		if configErr != nil {
			return "", configErr
		}
	} else { // else load the config from the global array of config values
		c = configs[configNum]
	}

	// first try to load the target as a number
	targetNum, targetErr := strconv.Atoi(targetVal)
	// if we can't load the target as a number, load it as a file path
	if targetErr != nil {
		t, targetErr = target.TargetFromFile(targetVal)
		if targetErr != nil {
			return "", targetErr
		}
	} else {
		t = targets[targetNum]
	}

	r.SetTarget(t)

	j := job.Job{RunRoutine: r, RunConfig: c}
	jobs = append(jobs, j)
	go j.Start()

	return "", nil
}

// stop will stop the given routine
func stop(args ...string) (string, error) {
	return "", nil
}

// pause will pause the given routine
func pause(args ...string) (string, error) {
	return "", nil
}

// monitor will monitor the given routine's log standard out
func monitor(args ...string) (string, error) {
	return "", nil
}
