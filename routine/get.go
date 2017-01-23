package routine

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/seantcanavan/target"
)

type Get struct {
	RunTarget target.Target
}

func (g Get) SetTarget(target target.Target) {
	g.RunTarget = target
}

func (g Get) Run() error {
	res, err := http.Get(g.RunTarget.URI)

	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		return err
	}

	ct := res.Header.Get("Content-Type")

	fmt.Println("GET target: ", g.RunTarget.URI)

	if ct != "" {
		if strings.Contains(strings.ToLower(ct), "json") {
			var prettyJSON bytes.Buffer
			err = json.Indent(&prettyJSON, body, "", "    ")
			if err != nil {
				return nil
			}
			fmt.Println(string(prettyJSON.Bytes()))
		}
	} else {
		fmt.Printf("%s", body)
	}

	return nil
}
