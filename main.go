package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// sort group floors
// sort by aisle
// sort by x coord
// write to xlsx

func main() {
	args := os.Args
	if len(args) < 2 {
		usage()
	}

	f, err := excelize.OpenFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get all the rows in the Sheet1.
	rows := f.GetRows("Sheet1")
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}

}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s <inputfile>\n", os.Args[0])
	flag.PrintDefaults()
	fmt.Scanln()
	os.Exit(2)
}
