package config

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

type Config struct {
	Backoff       bool          // Toggle simple back off after unsuccessful attempts. For {N, N+1, N+2, N+3} iterations, delay = X, 2X, 4X, 8X, etc.
	Duration      int64         // The routine ends after Duration in seconds, unless MaxIteratoins is hit first. Set to 0 for infinity.
	MaxIterations uint64        // The Routine ends after MaxIterations, unless Duration is hit first. Set to 0 for infinity.
	TimeBetween   time.Duration // The delay between each iteration in seconds.
	Routine       string        // The name of the routine passed in via json. E.G. {GET, PING, SSH}
	WriteToFile	  bool 			// Whether or not to write to a file instead of standard out.
}

// DefaultConfig returns a sample default config that will execute the ping routine.
func DefaultConfig() Config {
	c := Config{}
	c.Backoff = true
	c.Duration = 200
	c.MaxIterations = 10
	c.TimeBetween = 3
	c.Routine = "ping"
	c.WriteToFile = false
	return c
}

// ConfigFromFile will load the given file and construct an instance of Config{} from the JSON data inside.
func ConfigFromFile(filePath string) (Config, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return DefaultConfig(), err
	}
	var c Config
	err = json.Unmarshal(data, &c)
	if err != nil {
		return DefaultConfig(), err
	}
	return c, nil
}
