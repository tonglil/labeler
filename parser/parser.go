package parser

import (
	"fmt"
	"log"

	"github.com/tonglil/labeler/types"

	"gopkg.in/yaml.v2"
)

var labels = `
repo: tonglil/labeler
labels:
  - name: bug
    color: fc2929
  - name: fix
    color: fbca04
    # if label doesn't exist
    #   if from exists: rename from to it
    #   else create it
  - name: feature
    color: 0052cc
    from: help wanted
  - name: enhancement
    color: 0e8a16
`

func main() {
	f := types.File{}

	test := labels

	err := yaml.Unmarshal([]byte(test), &f)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("--- f:\n%+v\n\n", f)

	d, err := yaml.Marshal(&f)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- d dump:\n%s\n", string(d))

	fmt.Printf("--- original:\n%s\n", string(test))
}
