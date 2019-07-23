package bgs

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// csv读取模板
func CsvReader(t *Template) error {
	// 1. 读取csv文件记录
	// 2. 设置模板头部
	// 3. 设置模板内容
	reader := csv.NewReader(t.GetFile())
	var records[][]string
	for i := 0; i < 3; i ++ {
		record, err := reader.Read()
		if err == io.EOF {
			break;
		}
		if err != nil {
			return err
		}
		records = append(records, record)
	}
	if len(records) == 1 {
		t.Body = records[0]
		return nil
	}
	t.Head = records[0]
	t.Body = records[1]
	return nil
}

// 输出序列
func CsvOutputSequence(t *Template) error {
	// 1. 输出头部和模板
	writer := csv.NewWriter(os.Stdout)
	if len(t.Head) > 0 {
		writer.Write(t.Head)
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
		writer.Write(v)
	}
	writer.Flush()
	return nil
}