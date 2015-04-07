package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"sort"
)

const DEV = false
const n = 37

type DesignEntity struct {
	Name         string   `yaml:"name"` // Affects YAML field names too.
	Requirements []string `yaml:"requirements"`
}

// for sorting
type DesignEntitys []DesignEntity

func (a DesignEntitys) Len() int               { return len(a) }
func (slice DesignEntitys) Less(i, j int) bool { return slice[i].Name < slice[j].Name }
func (slice DesignEntitys) Swap(i, j int)      { slice[i], slice[j] = slice[j], slice[i] }

type DesignDoc struct {
	Reqs        []map[string]string `yaml:"Requirements"`
	Controllers []DesignEntity      `yaml:"Controllers"`
	Models      []DesignEntity      `yaml:"Models"`
	Views       []DesignEntity      `yaml:"Views"`
}

// // This will chunk up the requirements ever n requiremens
// func chunk(n int, ar []map[string]string) [][]map[string]string {
//
// 	var arOut [][]map[string]string
//
// 	for i := 0; i < len(ar); i += n {
// 		maxSlice := i + n
//
// 		if i+n > len(ar) {
// 			maxSlice = len(ar)
// 		}
//
// 		arOut = append(arOut, ar[i:maxSlice])
//
// 	}
//
// 	return arOut
// }

func toCSVTable(title string, entitys []DesignEntity, reqs []map[string]string) {

	fmt.Println("" + title + "")
	fmt.Println("")

	for _, entity := range entitys {
		fmt.Printf(`,%s`, entity.Name)
	}

	fmt.Println()

	// each reqirement row
	for _, req := range reqs {
		fmt.Printf(`%s,`, req["name"])

		// each reqirement in a design entity
		for _, entity := range entitys {

			has_it := false

			for _, dreq := range entity.Requirements {
				if req["name"] == dreq {
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

func makeHeader(mid bool, entitys []DesignEntity) {
	if mid {
		fmt.Println("</table>")
		fmt.Println(`<p style="page-break-after:always;"></p>`)
	}
	fmt.Println("<table>")

	fmt.Print(`<tr><th class="req-name"></th>`)

	for _, entity := range entitys {
		var badClass string

		if len(entity.Requirements) == 0 {
			badClass = " bad"
		} else {
			badClass = ""
		}

		if DEV {
			fmt.Printf(`<th class="rotate%s"><div><span>%s (%d)</span></div></th>`, badClass, entity.Name, len(entity.Requirements))
		} else {
			fmt.Printf(`<th class="rotate"><div><span>%s</span></div></th>`, entity.Name)
		}
	}
	fmt.Print("</tr>\n")
}

/*
This will convert a slice of design to an HTML Table
*/
func toHTMLTable(title string, entitys []DesignEntity, reqs []map[string]string) {

	// fmt.Println("<h2>" + title + "</h2>")

	makeHeader(false, entitys)

	// each reqirement row
	for i, req := range reqs {

		if ((i + 1) % n) == 0 {
			makeHeader(true, entitys)
		}

		fmt.Print("<tr>")
		if DEV {
			fmt.Printf(`<td class="req-name"><a title="%s" href="#%s">%s</a></td>`, req["description"], req["name"], req["name"])
		} else {
			fmt.Printf(`<td class="req-name">%s</td>`, req["name"])
		}

		// each reqirement in a design entity
		for _, entity := range entitys {

			has_it := false

			for _, dreq := range entity.Requirements {
				if req["name"] == dreq {
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
	fmt.Println(`<p style="page-break-after:always;"></p>`)

}

func toMarkdownTable(title string, entitys []DesignEntity, reqs []map[string]string) {

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
	for _, req := range reqs {
		fmt.Printf("| %s |", req["name"])

		// each reqirement in a design entity
		for _, entity := range entitys {

			has_it := false

			for _, dreq := range entity.Requirements {
				if req["name"] == dreq {
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
			sort.Sort(DesignEntitys(doc.Controllers))
			sort.Sort(DesignEntitys(doc.Models))
			sort.Sort(DesignEntitys(doc.Views))
			// sort.Sort(doc.Reqs)

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

				if DEV {
					for _, req := range doc.Reqs {
						fmt.Printf(`<h3 id="%s"><strong>%s</strong>: %s</h3>`, req["name"], req["name"], req["description"])
						fmt.Printf(`<p><strong>Rationale:</strong> %s</p>`, req["rationale"])
					}
				}
			}
		}
	}
}
