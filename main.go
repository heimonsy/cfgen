package main

import (
	"fmt"
	"log"
	"os"
)

var (
	usage = `gentemp Pname`
	temp  = `package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func Print(args ...interface{})                 { out.WriteString(fmt.Sprint(args...)) }
func Printf(format string, args ...interface{}) { out.WriteString(fmt.Sprintf(format, args...)) }
func Println(args ...interface{})               { out.WriteString(fmt.Sprintln(args...)) }

func Scan(args ...interface{})   { fmt.Fscan(in, args...) }
func Scanln(args ...interface{}) { fmt.Fscanln(in, args...) }

func main() {
	defer out.Flush()
}
`
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println(usage)
	}
	err := os.Mkdir(args[1], 0755)
	errorCheck(err)
	f, err := os.OpenFile(args[1]+"/"+args[1]+".go", os.O_CREATE|os.O_RDWR, 0644)
	errorCheck(err)
	defer f.Close()
	_, err = f.WriteString(temp)
	errorCheck(err)
}

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
