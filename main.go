package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
	"strings"
)

//"golang.org/x/net/html"

type Link struct {
	Href string
	Text string
}

// TODO loadHTML
func loadHTML(name string) string {
	file, _ := os.ReadFile(name)
	return string(file)
}

// TODO parseHTML
func parseHTML(file string) []Link {
	doc, err := html.Parse(strings.NewReader(file))
	if err != nil {
		fmt.Println("Error parsing HTML:", err)
		panic("File Error")
	}
	var sl []Link
	buildSlice(doc, 0, &sl)
	return sl
}

func buildSlice(n *html.Node, depth int, sl *[]Link) {
	if n.Type == html.ElementNode {
		if n.DataAtom.String() == "a" {
			text := strings.TrimSpace(n.FirstChild.Data)
			text = strings.TrimRight(text, "\n")
			l := Link{
				Href: n.Attr[0].Val,
				Text: text,
			}
			*sl = append(*sl, l)
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		buildSlice(c, depth+1, sl)
	}
}

func main() {
	file := loadHTML("ex3.html")
	links := parseHTML(file)
	for _, l := range links {
		fmt.Printf("Href: %s \n", l.Href)
		fmt.Printf("Text: %s \n", l.Text)
	}
}
