package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	l := len(os.Args)
	if l != 3 {
		//fmt.Println(l)
		fmt.Println("Usage: grep.exe file \"string\"")
		return
	}
	file := os.Args[1]
	rgx := os.Args[2]
	f, e := os.Open(file)
	defer f.Close()
	if e != nil {
		panic(e)
	}
	buf := bufio.NewReader(f)

	for {
		b, e := buf.ReadBytes('\n')
		if e != nil {
			if e == io.EOF {
				break
			}
			panic(e)
		}
		s := string(b)
		//fmt.Println(s)
		r := regexp.MustCompile(rgx)
		//		r := regexp.MustCompile("f([a-z]+)g")
		if r.MatchString(s) {
			fmt.Print(s)
			//fmt.Println("a")
		}
		//time.Sleep(time.Second * 1)
	}

}
