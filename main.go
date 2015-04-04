package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"sort"
)

const DEV = true

type DesignEntity struct {
	Name         string   `yaml:"name"` // Affects YAML field names too.
	Requirements []string `yaml:"requirements"`
}

type DesignEntitys []DesignEntity

func (a DesignEntitys) Len() int { return len(a) }

func (slice DesignEntitys) Less(i, j int) bool {
	return slice[i].Name < slice[j].Name
}

func (slice DesignEntitys) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

type DesignDoc struct {
	Reqs        map[string]map[string]string `yaml:"Requirements"`
	Controllers []DesignEntity               `yaml:"Controllers"`
	Models      []DesignEntity               `yaml:"Models"`
	Views       []DesignEntity               `yaml:"Views"`
}

func toCSVTable(title string, entitys []DesignEntity, reqs map[string]map[string]string) {

	fmt.Println("" + title + "")
	fmt.Println("")

	for _, entity := range entitys {
		fmt.Printf(`,%s`, entity.Name)
	}

	fmt.Println()

	// each reqirement row
	for req_name := range reqs {
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

func sortedKeys(m map[string]map[string]string) []string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return keys
}

/*
This will convert a slice of design to an HTML Table
*/
func toHTMLTable(title string, entitys []DesignEntity, reqs map[string]map[string]string) {

	fmt.Println("<h2>" + title + "<h2>")
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

	// each reqirement row
	for _, req_name := range sortedKeys(reqs) {
		fmt.Print("<tr>")
		if DEV {
			fmt.Printf(`<td class="req-name"><a href="#%s">%s</a></td>`, req_name, req_name)
		} else {
			fmt.Printf(`<td class="req-name">%s</td>`, req_name)
		}

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
	fmt.Println(`<p style="page-break-after:always;"></p>`)

}

func toMarkdownTable(title string, entitys []DesignEntity, reqs map[string]map[string]string) {

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
	for req_name, _ := range reqs {
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
					for _, name := range sortedKeys(doc.Reqs) {
						contents := doc.Reqs[name]
						fmt.Printf(`<h2 id="%s">%s:</h2>`, name, name)
						fmt.Printf(`<p><strong>description:</strong>%s</p>`, contents["description"])
						fmt.Printf(`<p><strong>rationale:</strong>%s</p>`, contents["rationale"])
					}
				}
			}
		}
	}
}
