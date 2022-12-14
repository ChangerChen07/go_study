/* package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func testReadAll() {
	// r := strings.NewReader("Hello World!!")
	f, _ := os.Open("test.txt")
	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("string(b): %v\n", string(b))
	}
}

func testReadDir(name string) {
	fi, err := ioutil.ReadDir("a")
	if err != nil {
		log.Fatal(err)
	} else {
		for _, v := range fi {
			fmt.Printf("v.Name(): %v\n", v.Name())

		}
	}
}

func testWalk() {
	var files []string
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		log.Fatal(err)
		// panic(err)
	}
	for _, v := range files {
		fmt.Println(v)
	}
}

func testReadFile() {
	b, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("string(b): %v\n", string(b))
}

func testWriteFile() {
	// func WriteFile(filename string, data []byte, perm fs.FileMode) error
	message := []byte("Hello, Gophers!")
	err := ioutil.WriteFile("hello", message, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func testTempFile() {
	content := []byte("temporary file's content")
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tmpfile.Name(): %v\n", tmpfile.Name())
	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
}
func main() {
	// testWalk()
	// testReadFile()
	// testWriteFile()
	testTempFile()
}
*/