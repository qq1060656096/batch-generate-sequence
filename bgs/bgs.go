package bgs

import "os"

type TemplateReaderFunc func(*Template) error

type TemplateOutputSequenceFunc func(*Template) error

const (
	TYPE_CSV string = "csv"
	TYPE_EXCEL = "excel"
)

// 模板文件
type Template struct {
	file *os.File
	Head []string
	Body []string
	Config *TemplateConfig
	// 读取模板文件
	Reader TemplateReaderFunc
	// 输出模板生成的序列
	OutputSequence TemplateOutputSequenceFunc
}

// 模板配置
type TemplateConfig struct {
	Prefix string
	Count uint64
	StartCount uint64
	ExcludeIndexs []string
	OutputType string
}

// 新建模板
func New(file *os.File) *Template {
	return &Template{
		file: file,
	}
}

// 获取模板文件
func (t *Template) GetFile() *os.File {
	return t.file
}

// 运行模板
func (t *Template) Run() error {
	if err := t.Reader(t); err != nil {
		return err
	}
	if err := t.OutputSequence(t); err != nil {
		return err
	}
	return nil
}

// 开始序列
func (c *TemplateConfig) GetStartSequence() uint64 {
	return c.StartCount
}

// 结束序列
func (c *TemplateConfig) GetEndSequence() uint64 {
	return c.StartCount + c.Count
}

// 检测字符是否在数组中
func CheckStringInArray(search string, slices []string) bool {
	for _, v := range slices  {
		if v == search {
			return true
		}
	}
	return false
}