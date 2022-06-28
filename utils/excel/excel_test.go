package excel

import (
	"testing"
)

type xlsxTestRecord struct {
	Name string `excel:"name"`
	Age  int    `excel:"age"`
}

func TestSplitExcel(t *testing.T) {
	f := NewExcelFile()
	sheet := "test"
	sheet1Index := f.NewSheet(sheet)
	f.SetActiveSheet(sheet1Index)
	err := f.SetStructHeader(sheet, &xlsxTestRecord{})
	if err != nil {
		t.Fatal(err)
	}
	err = f.SetStructRow(sheet, 1, xlsxTestRecord{
		Name: "name1",
		Age:  10,
	})
	if err != nil {
		t.Fatal(err)
	}
	err = f.SetStructRow(sheet, 2, xlsxTestRecord{
		Name: "name2",
		Age:  20,
	})
	if err != nil {
		t.Fatal(err)
	}
	err = f.SaveAs("test.xlsx")
	if err != nil {
		t.Fatal(err)
	}
}

func TestWholeExcel(t *testing.T) {
	tt := []xlsxTestRecord{{"name", 20}}
	f := SetRecord2Xlsx("test", tt)
	err := f.SaveAs("test.xlsx")
	if err != nil {
		t.Fatal(err)
	}
}
