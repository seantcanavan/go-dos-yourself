package routine

import (
    "path/filepath"
    "testing"

    "github.com/seantcanavan/target"
)

func TestGetReqRes1Pass(t *testing.T) {
    target, err := target.TargetFromFile(filepath.Join("..", "target", "samples", "get_reqres01.json"))

    if err != nil {
        t.Error(err)
    }

    var r Routine
    r = Get{RunTarget: target}

    err = r.Run()

    if err != nil {
        t.Error(err)
    }
}

func TestGetReqRes2Pass(t *testing.T) {
    target, err := target.TargetFromFile(filepath.Join("..", "target", "samples", "get_shodan01.json"))

    if err != nil {
        t.Error(err)
    }

    var r Routine
    r = Get{RunTarget: target}

    err = r.Run()

    if err != nil {
        t.Error(err)
    }
}

func TestGetMarkerApi01Pass(t *testing.T) {
    target, err := target.TargetFromFile(filepath.Join("..", "target", "samples", "get_markerapi01.json"))

    if err != nil {
        t.Error(err)
    }

    var r Routine
    r = Get{RunTarget: target}

    err = r.Run()

    if err != nil {
        t.Error(err)
    }
}
