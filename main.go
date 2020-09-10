package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// sort by aisle
// sort by x coord
// write to xlsx
type floor struct {
	level_1 [][]string
	level_2 [][]string
	level_3 [][]string
	level_4 [][]string
}

func main() {
	args := os.Args
	if len(args) < 2 {
		usage()
	}

	fi := new(floor)

	f, err := excelize.OpenFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get all the rows in the Sheet1.
	rows := f.GetRows("Sheet1")
	for _, row := range rows {
		// for _, colCell := range row {
		// 	fmt.Print(colCell, "\t")
		// }
		// fmt.Println("whole row", row)
		// fmt.Println("index", row[1])
		level(row[1], fi, row)
	}
	fmt.Println("first floor\n", fi.level_1, "\n",
		"second floor\n", fi.level_2, "\n",
		"third floor\n", fi.level_3, "\n",
		"fourth floor\n", fi.level_4, "\n")
}

func level(s string, f *floor, r []string) {

	//NRA4103X66240Y02Z25
	// Floor is always going to be the 7th char
	sn := string(s[6])
	switch sn {
	case "1":
		fmt.Println("First floor", s)
		f.level_1 = append(f.level_1, r)
	case "2":
		fmt.Println("Second floor", s)
		f.level_2 = append(f.level_2, r)
	case "3":
		fmt.Println("Third floor", s)
		f.level_3 = append(f.level_3, r)
	case "4":
		fmt.Println("Fourth floor", s)
		f.level_4 = append(f.level_4, r)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s <inputfile>\n", os.Args[0])
	flag.PrintDefaults()
	fmt.Scanln()
	os.Exit(2)
}
