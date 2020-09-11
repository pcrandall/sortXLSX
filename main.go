package main

import (
	"flag"
	"fmt"
	"os"
	"sort"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// sort by aisle
// sort by x coord
// write to xlsx
type floor struct {
	f1 [][]string
	f2 [][]string
	f3 [][]string
	f4 [][]string
}

func main() {
	args := os.Args
	if len(args) < 2 {
		usage()
	}

	//flr := &floor{} is same as flr := new(floor)
	flr := &floor{}

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
		level(row[1], flr, row)
	}
	// fmt.Println("first floor\n", flr.f1, "\n",
	// 	"second floor\n", flr.f2, "\n",
	// 	"third floor\n", flr.f3, "\n",
	// 	"fourth floor\n", flr.f4, "\n")

	sort.Slice(flr.f1, func(j, k int) bool {
		// J, _ := strconv.Atoi(flr.f1[j][0])
		// K, _ := strconv.Atoi(flr.f1[k][0])
		// fmt.Println("J", flr.f1[j][0], "K", flr.f1[k][0])
		fmt.Println("J", flr.f1[j][1], "K", flr.f1[k][1], flr.f1[j][1] < flr.f1[j][1])
		return flr.f1[j][1] > flr.f1[k][1]
	})
	sort.Slice(flr.f2, func(j, k int) bool {
		return flr.f2[j][1] > flr.f2[k][1]
	})
	sort.Slice(flr.f3, func(j, k int) bool {
		return flr.f3[j][1] > flr.f3[k][1]
	})
	sort.Slice(flr.f4, func(j, k int) bool {
		return flr.f4[j][1] > flr.f4[k][1]
	})

	fmt.Println("first floor\n", flr.f1, "\n",
		"second floor\n", flr.f2, "\n",
		"third floor\n", flr.f3, "\n",
		"fourth floor\n", flr.f4, "\n")

}

func level(s string, flr *floor, r []string) {

	//NRA4103X66240Y02Z25
	// Floor is always going to be the 7th char
	sn := string(s[6])
	fmt.Println(sn)
	switch sn {
	case "1":
		fmt.Println("First floor", s)
		flr.f1 = append(flr.f1, r)
	case "2":
		fmt.Println("Second floor", s)
		flr.f2 = append(flr.f2, r)
	case "3":
		fmt.Println("Third floor", s)
		flr.f3 = append(flr.f3, r)
	case "4":
		fmt.Println("Fourth floor", s)
		flr.f4 = append(flr.f4, r)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s <inputflrle>\n", os.Args[0])
	flag.PrintDefaults()
	fmt.Scanln()
	os.Exit(2)
}
