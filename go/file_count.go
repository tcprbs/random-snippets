package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

func main() {

	dir := os.Args[1]
	count := 0
	var wg sync.WaitGroup

	ch := make(chan string)

	go func() {
		wg.Add(1)
		traverse(dir, &wg, ch)
		wg.Wait()
		close(ch)
	}()

	for range ch {
		count = count + 1
	}

	fmt.Println(count)
}

func traverse(root string, wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()
	files, err := ioutil.ReadDir(root)

	if err != nil {
		fmt.Println(err)
	}

	for _, f := range files {
		if f.IsDir() {
			wg.Add(1)
			name := filepath.Join(root, f.Name())
			go traverse(name, wg, ch)
		} else {
			if f.Mode().IsRegular() {
				ch <- f.Name()
			}
		}
	}

}
