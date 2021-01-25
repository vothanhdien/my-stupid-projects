package main

import (
	"log"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

const (
	csv  = "bibs.csv"
	xls  = "bibs.xls"
	xlsx = "bibs.xlsx"
)

func main() {
	//if _, err := excelize.OpenFile(csv); err != nil {
	//	log.Fatalf("open csv err %v", err)
	//}
	//if _, err := excelize.OpenFile(xls); err != nil {
	//	log.Fatalf("open xls err %v", err)
	//}
	if _, err := excelize.OpenFile(xlsx); err != nil {
		log.Fatalf("open xlsx err %v", err)
	}
}
