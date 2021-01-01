package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"
	"time"
)

// WaitForServer attempts to contact the server of a URL.
// It tries  for one minute using exponential back-off.
// It reports an error if all attempts fail.
func WaitForServer(url string) error {
	const timeout = 6 * time.Second
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil // success
		}
		log.Printf("server not responding (%s); retrying %dth time ...", err, tries + 1)
		time.Sleep(time.Second << uint(tries)) // exponential back-off
	}
	return fmt.Errorf("server %s failed to reponse after %v", url, timeout.Seconds())
}

func TestNormalFetch(t *testing.T) {
	fmt.Println(WaitForServer("http://www.baidu.com"))
}

func TestBadFetch(t *testing.T) {
	fmt.Println(WaitForServer("unkown.xx"))
}

func TestDown(t *testing.T) {
	if err := WaitForServer("xx"); err != nil {
		fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
		// os.Exit(1)
	}
}

// Always fail which is normal.
// Fatalf will abort the program.
func TestFatal(t *testing.T) {
	if err := WaitForServer("xx"); err != nil {
		log.Fatalf("Site is down: %v\n", err)
	}
}

func TempDir(dest string) error {
	dir, err := ioutil.TempDir("", dest)
	if err != nil {
		return fmt.Errorf("failed to create temp dir: %v", err)
	}
	// ... use temp dir ...
	os.RemoveAll(dir)
	return nil
}

func TestIgnoreError(t *testing.T) {
	TempDir("scratch")
}

// End of File(EOF)
func ReadRune() error {
	in := bufio.NewReader(os.Stdin)
	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			fmt.Println("EOF")
			break
		}
		if err != nil {
			return fmt.Errorf("read failed: %v", err)
		}
		fmt.Printf("rune=[%v]", r)
		break
	}
	return nil
}

func TestReadRune(t *testing.T) {
	ReadRune()
}
