package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
	"testing"
)

func findLinks(url string) ([]string, error) {
	reply, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if reply.StatusCode != http.StatusOK {
		reply.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, reply.Status)
	}
	doc, err := html.Parse(reply.Body)
	reply.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	return Visit(nil, doc), nil
}

func TestMultipleReturnValues(t *testing.T) {
	links, err := findLinks("https://golang.org")
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
		os.Exit(1)
	}
	for _, link := range links {
		fmt.Println(link)
	}
}

func TestSole(t *testing.T) {
	links, err := findLinks("https://golang.org")
	fmt.Println(links, err)
	fmt.Println(findLinks("https://golang.org"))
}

func CountWordsAndImages(url string) (words, images int, err error) {
	reply, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(reply.Body)
	reply.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

// Exercise 5.5
func countWordsAndImages(n *html.Node) (words, images int) {
	if n == nil {
		return
	}
	if n.Type == html.ElementNode {
		if n.Data == "style" || n.Data == "script" {
			return
		} else if n.Data == "img" {
			images++
		}
	} else if n.Type == html.TextNode {
		words += CountWords(n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		w, i := countWordsAndImages(c)
		words += w
		images += i
	}
	return words, images
}

func TestCountWordsAndImages(t *testing.T) {
	fmt.Println(CountWordsAndImages("http://golang.org"))
}

func CountWords(text string) int {
	total := 0
	for _, line := range strings.Split(text, "\n") {
		fixLine := strings.TrimSpace(line)
		if fixLine != "" {
			total += len(strings.Split(fixLine, " "))
		}
	}
	return total
}

func TestTrimSpace(t *testing.T) {
	text := " z ab x \n sdfsa \n   \n\n why \n\n "
	fixText := strings.TrimSpace(text)
	assert.Equal(t, "z ab x \n sdfsa \n   \n\n why", fixText)
	fmt.Printf("[%s]", fixText)
	assert.Equal(t, 5, CountWords(text))
}