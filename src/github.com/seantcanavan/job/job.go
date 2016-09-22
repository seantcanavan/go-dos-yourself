package job

import (
	"fmt"
	"github.com/seantcanavan/config"
	"github.com/seantcanavan/routine"
	"time"
)

type Job struct {
	RunRoutine routine.Routine
	RunConfig  config.Config
}

func (j Job) Start() {
	count := uint64(0) // number of executions so far
	start := time.Now().Unix() // the instant we've begun executing

	// implement a floor for TimeBetween if Backoff is true
	// this allows lazy users to use Backoff without a specific TimeBetween
	if j.RunConfig.Backoff && (j.RunConfig.TimeBetween <= 0) {
		j.RunConfig.TimeBetween = 1
	}

	for ((count < j.RunConfig.MaxIterations) && j.RunConfig.MaxIterations != 0) &&
		((time.Now().Unix() - start) < j.RunConfig.Duration) {
		// execute our job
		err := j.RunRoutine.Run()

		// alert the user if we're sleeping and execute the sleep
		if j.RunConfig.TimeBetween > 0 {
			fmt.Println("Sleeping for", j.RunConfig.TimeBetween*time.Second)
			time.Sleep(j.RunConfig.TimeBetween * time.Second)
            if err != nil && j.RunConfig.Backoff { // if we error'd and backoff is set
                j.RunConfig.TimeBetween = j.RunConfig.TimeBetween * 2
            }
		}

		count++
		fmt.Println("Run", count, "finished")
	}

	end := time.Now().Unix()
	totalSeconds := float64(end - start)
	fmt.Printf("Total running seconds: %.2f \n", totalSeconds)

	// additional stats for nerds
	if j.RunConfig.TimeBetween > 0 { // if we're pausing between iterations
		activeSeconds := float64(totalSeconds) - ((j.RunConfig.TimeBetween * time.Second).Seconds() * float64(count))
		inactiveSeconds := totalSeconds - activeSeconds
		percentActive := activeSeconds / totalSeconds * 100
		percentInactive := inactiveSeconds / totalSeconds * 100

		fmt.Printf("Total active seconds: %.2f \n", activeSeconds)
		fmt.Printf("Total inactive seconds: %.2f \n", inactiveSeconds)
		fmt.Printf("Percent of time active: %.2f \n", percentActive)
		fmt.Printf("Percent of time inactive: %.2f \n", percentInactive)
	}

	if count < j.RunConfig.MaxIterations {
		fmt.Println("Executed longer than", j.RunConfig.Duration, "seconds. Exiting.")
	} else {
		fmt.Println("Executed", count, "of", j.RunConfig.MaxIterations, "runs. Exiting.")
	}
}
