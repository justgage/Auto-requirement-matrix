package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type DesignEntity struct {
	Name         string   `yaml:"name"` // Affects YAML field names too.
	Requirements []string `yaml:"reqirements"`
}

type DesignDoc struct {
	Reqs []string       `yaml:"Reqirements"`
	DEs  []DesignEntity `yaml:"Design Entities"`
}

/*
This will convert a slice of design to an HTML Table
*/
func toHtml(entitys []DesignEntity, reqs []string) {

	fmt.Println("<table>")

	fmt.Print("<tr><th></th>")

	for _, entity := range entitys {
		fmt.Printf("<th>%s</th>", entity.Name)
	}
	fmt.Print("</tr>\n")

	// each reqirement row
	for _, req_name := range reqs {
		fmt.Print("<tr>")
		fmt.Printf("<td>%s</td>", req_name)

		// each reqirement in a design entity
		for _, entity := range entitys {

			has_it := false

			for _, dreq := range entity.Requirements {
				if req_name == dreq {
					has_it = true
					break
				}
			}

			if has_it {
				fmt.Printf("<td>x</td>\n")
			} else {
				fmt.Printf("<td> </td>\n")
			}
		}

		fmt.Printf("</tr>\n")
	}

	fmt.Println("</table>")
}

func main() {

	if len(os.Args) == 1 {
		panic("Please enter a filename!")
	}

	y, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	var doc DesignDoc

	err = yaml.Unmarshal([]byte(y), &doc)

	if err != nil {
		fmt.Println(err)
	} else {
		toHtml(doc.DEs, doc.Reqs)
	}
}
