package routine

import (
    "fmt"
    "strings"
    "github.com/seantcanavan/target"
)

// Routine defines an interface through which any abstract set of network commands can be executed.
type Routine interface {
	Run() error
    SetTarget(target target.Target)
}

func PrintallRoutineNames() {
    fmt.Println("get")
    fmt.Println("ping")
    fmt.Println("ssh")
}

func GetRoutinebyName(name string) (Routine, error) {
    name = strings.ToLower(name)
    if name == "get" {
        return Get{}, nil
    } else if name == "ping" {
        return Ping{}, nil
    } else if name == "ssh" {
        return SSH{}, nil
    } else {
        return nil, fmt.Errorf("The value you provided:", name, "does not correspond to any known routines in the system")
    }
}
