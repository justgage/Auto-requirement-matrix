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
	Reqs        []string       `yaml:"Reqirements"`
	Controllers []DesignEntity `yaml:"Controllers"`
	Models      []DesignEntity `yaml:"Models"`
}

/*
This will convert a slice of design to an HTML Table
*/
func toTable(entitys []DesignEntity, reqs []string) {

	fmt.Println("<table>")

	fmt.Print(`<tr><th class="req-name"></th>`)

	for _, entity := range entitys {
		fmt.Printf(`<th class="rotate"><div><span>%s</span></div></th>`, entity.Name)
	}
	fmt.Print("</tr>\n")

	// each reqirement row
	for _, req_name := range reqs {
		fmt.Print("<tr>")
		fmt.Printf(`<td class="req-name">%s</td>`, req_name)

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
	fmt.Println(`<p style="page-break-after:always;"></p>
	`)
}

func cssRender(css []byte) {
	fmt.Printf("<style>\n%s\n</style>", css)
}

func main() {

	if len(os.Args) == 1 {
		panic("Please enter a filename!")
	}

	yaml_stuff, err := ioutil.ReadFile(os.Args[1])
	css, err := ioutil.ReadFile("style.css")

	cssRender(css)

	if err != nil {
		panic(err)
	}

	var doc DesignDoc

	err = yaml.Unmarshal([]byte(yaml_stuff), &doc)

	if err != nil {
		fmt.Println(err)
	} else {
		toTable(doc.Controllers, doc.Reqs)
		toTable(doc.Models, doc.Reqs)
	}
}
