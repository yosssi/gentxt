package main

import (
	"bufio"
	"flag"
	"math/rand"
	"os"
	"time"
)

const txts = "abcdefghijklmnopqrstuvwxyz"

var rownum = flag.Int("r", 0, "row number")
var maxc = flag.Int("c", 1, "max column number")
var path = flag.String("p", "data.txt", "file path")

func init() {
	rand.Seed(time.Now().UnixNano())
	flag.Parse()
}

func main() {
	f, err := os.Create(*path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	for r := 0; r < *rownum; r++ {
		for i, l := 0, rand.Intn(*maxc)+1; i < l; i++ {
			if err := w.WriteByte(txts[rand.Intn(len(txts))]); err != nil {
				panic(err)
			}
		}
		if err := w.WriteByte(0x0a); err != nil {
			panic(err)
		}
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}
}
