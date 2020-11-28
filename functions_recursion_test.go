package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"
)

func Fetch(url string) (io.Reader, error) {
	reply, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if reply.StatusCode != http.StatusOK {
		reply.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, reply.Status)
	}
	return reply.Body, nil
}

func Visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = Visit(links, c)
	}
	return links
}

func TestLinks(t *testing.T) {
	url := "http://golang.org"
	r, err := Fetch(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fetch %s error %s\n", url, err.Error())
		os.Exit(1)
	}
	doc, err := html.Parse(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Parse html error %s\n", err.Error())
		os.Exit(2)
	}
	for _, link := range Visit(nil, doc) {
		fmt.Println("extract link ==> ", link)
	}
}

// Exercise 5.1
func VisitV2(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	if n.FirstChild != nil {
		links = VisitV2(links, n.FirstChild)
	}

	if n.NextSibling != nil {
		links = VisitV2(links, n.NextSibling)
	}
	return links
}

func TestVisitV2(t *testing.T) {
	url := "http://golang.org"
	r, err := Fetch(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fetch %s error %s\n", url, err.Error())
		os.Exit(1)
	}
	doc, err := html.Parse(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Parse html error %s\n", err.Error())
		os.Exit(2)
	}
	for _, link := range VisitV2(nil, doc) {
		fmt.Println("extract link ==> ", link)
	}
}

// Ex 5.2
func Outline(stack []string, n *html.Node) {
	if n == nil {
		return
	}
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		Outline(stack, c)
	}
}

func TestOutline(t *testing.T) {
	url := "https://golang.org"
	r, err := Fetch(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fetch %s error %s\n", url, err.Error())
		os.Exit(1)
	}
	doc, err := html.Parse(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Parse html error %s\n", err.Error())
		os.Exit(2)
	}
	Outline(nil, doc)
}

func CalEntryMap(kv map[string]int, n *html.Node) {
	if n == nil {
		return
	}
	if n.Type == html.ElementNode {
		kv[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		CalEntryMap(kv, c)
	}
}

func TestMap(t *testing.T) {
	url := "https://golang.org"
	r, err := Fetch(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fetch %s error %s\n", url, err.Error())
		os.Exit(1)
	}
	doc, err := html.Parse(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Parse html error %s\n", err.Error())
		os.Exit(2)
	}
	kv := map[string]int{}
	CalEntryMap(kv, doc)
	for k, v := range kv {
		fmt.Printf("%s -> %d\n", k, v)
	}
}

func ExtractText(n *html.Node) {
	if n == nil {
		return
	}
	if n.Type == html.TextNode {
		if n.Data != "script" && n.Data != "style" {
			fmt.Printf("==> [%s]\n", strings.TrimSpace(n.Data))
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ExtractText(c)
	}
}

func TestText(t *testing.T) {
	url := "https://golang.org"
	r, err := Fetch(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fetch %s error %s\n", url, err.Error())
		os.Exit(1)
	}
	doc, err := html.Parse(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Parse html error %s\n", err.Error())
		os.Exit(2)
	}
	ExtractText(doc)
}

func ExtendVisit(links []string, n *html.Node) []string {
	for _, a := range n.Attr {
		if a.Key == "src" {
			links = append(links, a.Val)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = ExtendVisit(links, c)
	}
	return links
}

func TestExtentVisit(t *testing.T) {
	url := "http://www.baidu.com"
	r, err := Fetch(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fetch %s error %s\n", url, err.Error())
		os.Exit(1)
	}
	doc, err := html.Parse(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Parse html error %s\n", err.Error())
		os.Exit(2)
	}
	for _, link := range ExtendVisit(nil, doc) {
		fmt.Println(link)
	}
}