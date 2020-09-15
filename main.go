package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

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

	file, err := excelize.OpenFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get all the rows in the Sheet1.
	rows := file.GetRows("Sheet1")
	for _, row := range rows {
		level(row[1], flr, row)
	}

	sort.Slice(flr.f1, func(j, k int) bool {
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

	// Create new file and write sorted rows
	sortedFile := excelize.NewFile()
	if err != nil {
		fmt.Println(err)
		return
	}

	var rowIndex int

	for idx, sortedFlr := range flr.f4 {
		// Sheet to write, row index to write, convert slice to interface
		sortedFile.SetSheetRow("Sheet1", "A"+strconv.Itoa(idx+1), &[]interface{}{sortedFlr[1], sortedFlr[2], sortedFlr[3]})
		rowIndex = idx + 1
		fmt.Println("f4", sortedFlr)
	}

	for _, sortedFlr := range flr.f3 {
		rowIndex++
		// Sheet to write, row index to write, convert slice to interface
		sortedFile.SetSheetRow("Sheet1", "A"+strconv.Itoa(rowIndex), &[]interface{}{sortedFlr[1], sortedFlr[2], sortedFlr[3]})
		fmt.Println("f3", sortedFlr)
	}

	for _, sortedFlr := range flr.f2 {
		rowIndex++
		// Sheet to write, row index to write, convert slice to interface
		sortedFile.SetSheetRow("Sheet1", "A"+strconv.Itoa(rowIndex), &[]interface{}{sortedFlr[1], sortedFlr[2], sortedFlr[3]})
		fmt.Println("f2", sortedFlr)
	}

	for _, sortedFlr := range flr.f1 {
		rowIndex++
		// Sheet to write, row index to write, convert slice to interface
		sortedFile.SetSheetRow("Sheet1", "A"+strconv.Itoa(rowIndex), &[]interface{}{sortedFlr[1], sortedFlr[2], sortedFlr[3]})
		fmt.Println("f1", sortedFlr)
	}

	fmt.Println("Row index here", rowIndex)
	fmt.Println(os.Args[1])

	//if err := sortedFile.SaveAs("Sorted_" + os.Args[1]); err != nil {
	if err := sortedFile.SaveAs("Sorted_Blockadelist.xlsx"); err != nil {
		fmt.Println(err)
	}

}

func level(s string, flr *floor, r []string) {
	// check for empty row
	if len(s) < 1 || len(string(s)) < 7 {
		return
	}

	// "NRA4103X66240Y02Z25" Floor is always going to be the 7th char
	sn := string(s[6])

	isfloor := regexp.MustCompile(`1|2|3|4`)

	if isfloor.MatchString(sn) {
		switch sn {
		case "1":
			flr.f1 = append(flr.f1, r)
		case "2":
			flr.f2 = append(flr.f2, r)
		case "3":
			flr.f3 = append(flr.f3, r)
		case "4":
			flr.f4 = append(flr.f4, r)
		}
	} else {
		return
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s <inputflrle>\n", os.Args[0])
	flag.PrintDefaults()
	fmt.Scanln()
	os.Exit(2)
}
