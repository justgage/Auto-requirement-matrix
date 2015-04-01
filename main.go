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
	Views       []DesignEntity `yaml:"Views"`
}

// func splitReqs(reqs []string) {
// reqs
// }

func toCSVTable(title string, entitys []DesignEntity, reqs []string) {

	fmt.Println("" + title + "")
	fmt.Println("")

	for _, entity := range entitys {
		fmt.Printf(`,%s`, entity.Name)
	}

	fmt.Println()

	// each reqirement row
	for _, req_name := range reqs {
		fmt.Printf(`%s,`, req_name)

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
				fmt.Printf("x,")
			} else {
				fmt.Printf(" ,")
			}
		}

		fmt.Println()
	}

	fmt.Println("")
}

/*
This will convert a slice of design to an HTML Table
*/
func toHTMLTable(title string, entitys []DesignEntity, reqs []string) {

	fmt.Println("<h2>" + title + "<h2>")
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

func toMarkdownTable(title string, entitys []DesignEntity, reqs []string) {

	fmt.Println("# " + title + "")
	fmt.Println()
	fmt.Print("||")

	for _, entity := range entitys {
		fmt.Printf("  %s  |", entity.Name)
	}
	fmt.Print("\n|------|")

	for _ = range entitys {
		fmt.Printf("-----|")
	}

	fmt.Print("\n")

	// each reqirement row
	for _, req_name := range reqs {
		fmt.Printf("| %s |", req_name)

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
				fmt.Printf(" x |")
			} else {
				fmt.Printf("   |")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Print("\n")
}

func cssRender(css []byte) {
	fmt.Printf("<style>\n%s\n</style>", css)
}

func main() {

	if len(os.Args) == 1 {
		fmt.Println("Please enter a filename!")
	} else {

		yaml_stuff, err := ioutil.ReadFile(os.Args[1])
		typeOfFile := os.Args[2]

		if err != nil {
			panic(err)
		}

		var doc DesignDoc

		err = yaml.Unmarshal([]byte(yaml_stuff), &doc)

		if err != nil {
			fmt.Println(err)

		} else { // no error reading the yaml

			if typeOfFile == "markdown" {
				toMarkdownTable("Controllers", doc.Controllers, doc.Reqs)
				toMarkdownTable("Models", doc.Models, doc.Reqs)
				toMarkdownTable("Views", doc.Views, doc.Reqs)

			} else if typeOfFile == "csv" {
				toCSVTable("Controllers", doc.Controllers, doc.Reqs)
				toCSVTable("Models", doc.Models, doc.Reqs)
				toCSVTable("Views", doc.Views, doc.Reqs)

			} else { // default to HTML
				css, err := ioutil.ReadFile("style.css")

				if err != nil {
					panic(err)
				}

				cssRender(css)
				toHTMLTable("Controllers", doc.Controllers, doc.Reqs)
				toHTMLTable("Models", doc.Models, doc.Reqs)
				toHTMLTable("Views", doc.Views, doc.Reqs)
			}
		}
	}
}
