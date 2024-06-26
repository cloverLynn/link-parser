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
			text := getAllChildrenNodes(n)
			l := Link{
				Href: n.Attr[0].Val,
				Text: strings.Join(text, " "),
			}
			*sl = append(*sl, l)
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		buildSlice(c, depth+1, sl)
	}
}

func getAllChildrenNodes(n *html.Node) []string {
	var children []string
	hasSibling := true
	currentChild := n.FirstChild
	for hasSibling == true {
		children = append(children, pullText(currentChild))
		if currentChild.NextSibling != nil {
			currentChild = currentChild.NextSibling
		} else {
			hasSibling = false
		}
	}
	return children
}

func pullText(n *html.Node) string {
	if n.Type == html.CommentNode {
		return ""
	}
	if n.Type == html.ElementNode {
		if n.FirstChild != nil {
			return strings.TrimRight(strings.TrimSpace(n.FirstChild.Data), "\n")
		} else {
			return ""
		}
	} else {
		return strings.TrimRight(strings.TrimSpace(n.Data), "\n")

	}
}

func main() {
	file := loadHTML("ex4.html")
	links := parseHTML(file)
	for _, l := range links {
		fmt.Printf("Href: %s \n", l.Href)
		fmt.Printf("Text: %s \n", l.Text)
	}
}
