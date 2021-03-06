package routine

import (
	"path/filepath"
	"testing"

	"github.com/seantcanavan/go-dos-yourself/target"
)

func TestPingPass(t *testing.T) {
	target, err := target.TargetFromFile(filepath.Join("..", "target", "samples", "ping_google01.json"))

	if err != nil {
		t.Error(err)
	}

	var r Routine
	r = Ping{RunTarget: target}

	err = r.Run()

	if err != nil {
		t.Error(err)
	}
}
