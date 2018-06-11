package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var re1 = regexp.MustCompile(`([^ <>="+])=`)
	var re2 = regexp.MustCompile(`=([^ "=]\\)`)

	f, err := os.Open("./main.c")
	check(err)
	defer f.Close()

	w, err := os.OpenFile("./new_main.c", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)
	defer w.Close()

	r4 := bufio.NewReader(f)
	w4 := bufio.NewWriter(w)
	for s, err := r4.ReadString('\n'); err != io.EOF; s, err = r4.ReadString('\n') {
		t1 := re1.ReplaceAllString(s, `$1 =`)
		t2 := re2.ReplaceAllString(t1, `= $1`)
		_, err = w4.WriteString(t2)
		check(err)
		fmt.Println(t2)
	}
	w4.Flush()
}
