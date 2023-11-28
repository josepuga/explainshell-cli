// explainshell.com CLI.
// (c) José Puga. 2023. GPL3 License.

package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

const version = "0.1.0"
const usage = `ShellExplain CLI. Go Edition. v%s
(c)José Puga 2023. Under GPL 3 License.
Usage: %s <Any shell command with or without parameters>
`

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Printf(usage, version, filepath.Base(args[0]))
		os.Exit(1)
	}
	_url := "https://explainshell.com/explain?cmd="
	params := url.QueryEscape(strings.Join(args[1:], " "))
	data, err := http.Get(_url + params)
	if err != nil {
		fmt.Println("HTTP request error:", err)
		os.Exit(1)
	}
	defer data.Body.Close()
	if err != nil {
		fmt.Println("Error extracting HTTP Body:", err)
		os.Exit(1)
	}
	doc, err := html.Parse(data.Body)
	if err != nil {
		fmt.Println("Error parsing data:", err)
		os.Exit(1)
	}
	for _, line := range parse(doc) {
		println(line)
	}
	//io.Copy(os.Stdout, data.Body)
}

func parse(n *html.Node) []string {
	var result []string

	if n.Type == html.ElementNode && n.Data == "pre" {
		// Extract text between <pre></pre>
		preText := extractTextFromNode(n)
		// Remove any HTML tag (like <p>) before append...
		result = append(result, removeHTMLTags(preText))
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result = append(result, parse(c)...)
	}
	return result
}

// Extract text from node and childrens
func extractTextFromNode(n *html.Node) string {
	var buffer bytes.Buffer

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			buffer.WriteString(c.Data)
		} else {
			buffer.WriteString(extractTextFromNode(c))
		}
	}

	return buffer.String()
}

func removeHTMLTags(input string) string {
	re := regexp.MustCompile(`<[^>]*>`)
	return re.ReplaceAllString(input, "")
}
