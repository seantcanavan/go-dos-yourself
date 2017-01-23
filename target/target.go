package target

import (
	"encoding/json"
	"io/ioutil"
)

type Target struct {
	// Required Values
	Name         string // The canonical name of the target. E.G. "Load Balancer 3"
	URI			 string // The URI where the target can be reached. This can be a URL, IPV4 address, or IPV6 address.

	// Optional Values
	Port         uint16 // The port to use when resolving the target, if any.
	IPV6 		 bool   // whether to use IPV6 routing technologies or not. Default is to use IPV4 technologies.

	// Port Range Values. The routine will target all the ports between min and max if PortRange is set to true.
	PortRange	 bool   // Whether to cover a range of ports or not.
	PortRangeMin string // The starting port to target. E.G. 320
	PortRangeMax string // The ending port to target. E.G. 5620

	// IP Range Values. The routine will target all IP addresses between min and max if IPRange is set to true.
	IPRange	     bool   // Whether to cover a range of IP address or not.
	IPRangeMin	 string // The starting range of IP addresses to target. E.G. 124.103.224.100
	IPRangeMax	 string // The ending range of IP addresses to target. E.G. 124.103.224.132

	//REST variables
	Method		 string // The REST method to use: {GET,PUT,POST,DELETE}.
	ContentType  string // The content type you're sending via REST. E.G. application/json
	PayloadPath  string // The file path containing the payload to send. Should match ContentType.
}

// DefaultTarget returns a sample default target pointing at your localhost.
func DefaultTarget() Target {
	t := Target{}
	t.Name = "my 1337 machine"
	t.URI = "127.0.0.1"
	t.Port = 8080
	t.IPV6 = false
	t.PortRange = false
	t.IPRange = false
	return t
}

// TargetFromFile will load the given file and construct an instance of Target{} from the JSON data inside.
func TargetFromFile(filePath string) (Target, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return DefaultTarget(), err
	}
	var t Target
	err = json.Unmarshal(data, &t)
	if err != nil {
		return DefaultTarget(), err
	}
	return t, nil
}
