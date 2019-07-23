package bgs

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"os"
	"strconv"
	"strings"
)

// excel读取模板
func ExcelReader(t *Template) error {
	// 1. 读取csv文件记录
	// 2. 设置模板头部
	// 3. 设置模板内容
	f, err := excelize.OpenReader(t.GetFile())
	if err != nil {
		return err
	}
	sheetName := f.GetSheetName(1)
	rows, err := f.Rows(sheetName)
	index := 0
	var records [][]string
	for rows.Next() {
		index ++
		row, err := rows.Columns()
		if err != nil {
			return err
		}
		records = append(records, row)
	}

	if len(records) < 2 {
		t.Body = records[0]
		return nil
	}
	t.Head = records[0]
	t.Body = records[1]
	return nil
}

// 输出序列
func ExcelOutputSequence(t *Template) error {
	// 1. 输出头部和模板
	f := excelize.NewFile()
	sheetName := "Sheet1"
	rowsIndex := 1
	if len(t.Head) > 0 {
		axis := "A" + strconv.Itoa(rowsIndex)
		err := f.SetSheetRow(sheetName,  axis, &t.Head)
		if err != nil {
			return err
		}
	}
	startSequence := t.Config.GetStartSequence()
	endSequence := t.Config.GetEndSequence()
	columnLen := len(t.Body)
	prefix := t.Config.Prefix
	for ; startSequence <= endSequence; startSequence ++{
		v := make([]string, columnLen)
		copy(v, t.Body)
		for i := 0; i < columnLen; i++ {
			columnValue := v[i]
			columnValue = strings.Trim(columnValue, prefix)
			iStr := strconv.Itoa(i)
			if CheckStringInArray(iStr, t.Config.ExcludeIndexs) {
				v[i] = prefix+ columnValue
				continue
			}

			if newValue, err := strconv.ParseInt(columnValue, 10, 64); err == nil {
				v[i] = fmt.Sprintf("%s%d", prefix,newValue + int64(startSequence))
				continue
			} else if newValue, err := strconv.ParseFloat(columnValue, 64); err == nil {
				v[i] = fmt.Sprintf("%s%d", prefix, int64(newValue + float64( startSequence)))
			} else  {
				v[i] = fmt.Sprintf("%s%s%d", prefix, columnValue, startSequence)
			}
		}
		rowsIndex ++
		axis := "A" + strconv.Itoa(rowsIndex)
		f.SetSheetRow(sheetName, axis, &v)
	}
	f.WriteTo(os.Stdout)
	f.WriteToBuffer()
	return nil
}