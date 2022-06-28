package excel

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"reflect"
)

func SetRecord2Xlsx(sheet string, records interface{}) *excelize.File {
	xlsx := excelize.NewFile()    // new file
	index := xlsx.NewSheet(sheet) // new sheet
	xlsx.SetActiveSheet(index)    // set active (default) sheet
	t := reflect.TypeOf(records)

	if t.Kind() != reflect.Slice {
		panic("records must be slice")
	}

	s := reflect.ValueOf(records)

	for i := 0; i < s.Len(); i++ {
		elem := s.Index(i).Interface()
		elemType := reflect.TypeOf(elem)
		elemValue := reflect.ValueOf(elem)
		if elemType.Kind() == reflect.Ptr {
			elemType = elemType.Elem()
			elemValue = elemValue.Elem()
		}

		if elemType.Kind() != reflect.Struct {
			panic("record in slice must be a struct")
		}

		for j := 0; j < elemType.NumField(); j++ {
			field := elemType.Field(j)
			tag := field.Tag.Get("xlsx")
			name := tag
			column, _ := excelize.ColumnNumberToName(j + 1)

			if tag == "-" {
				continue
			}

			if tag == "" {
				tag = field.Name
			}
			// 设置表头
			if i == 0 {
				xlsx.SetCellValue(sheet, fmt.Sprintf("%s%d", column, i+1), name)
			}
			// 设置内容
			xlsx.SetCellValue(sheet, fmt.Sprintf("%s%d", column, i+2), elemValue.Field(j).Interface())
		}
	}

	return xlsx
}

type File struct {
	*excelize.File
}

type Sheet struct {
}

func NewExcelFile() *File {
	return &File{
		excelize.NewFile(),
	}
}

// set header name from struct
func (f *File) SetStructHeader(sheet string, src interface{}) error {
	typ := reflect.TypeOf(src)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	if typ.Kind() != reflect.Struct {
		panic("src must be struct")
	}

	col := 0
	for i := 0; i < typ.NumField(); i++ {
		sf := typ.Field(i)
		if sf.Anonymous {
			continue
		}
		tag := sf.Tag.Get("excel")
		if tag == "-" {
			continue
		}

		if tag == "" {
			tag = sf.Name
		}

		colName, err := excelize.ColumnNumberToName(col + 1)
		if err != nil {
			return err
		}

		err = f.SetCellValue(sheet, fmt.Sprintf("%s%d", colName, 1), tag)
		if err != nil {
			return err
		}

		col++
	}

	return nil
}

func (f *File) SetStructRow(sheet string, row int, src interface{}) error {
	srcValue := reflect.ValueOf(src)
	if srcValue.Kind() == reflect.Ptr {
		srcValue = srcValue.Elem()
	}

	srcType := srcValue.Type()
	if srcType.Kind() != reflect.Struct {
		panic("src must be struct")
	}

	col := 0
	for i := 0; i < srcType.NumField(); i++ {
		sf := srcType.Field(i)
		sv := srcValue.Field(i)
		if sf.Anonymous {
			continue
		}
		tag := sf.Tag.Get("excel")
		if tag == "-" {
			continue
		}

		if tag == "" {
			tag = sf.Name
		}

		colName, err := excelize.ColumnNumberToName(col + 1)
		if err != nil {
			return err
		}

		err = f.SetCellValue(sheet, fmt.Sprintf("%s%d", colName, row+2), sv.Interface())
		if err != nil {
			return err
		}

		col++
	}

	return nil
}
