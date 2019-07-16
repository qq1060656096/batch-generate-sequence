package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
	"strconv"
	"strings"
	"./bgs"
)

const (
	// 参数不够
	ErrorCodeRequireArgsNum = 0
	// 打开模板文件失败
	ErrorCodeTemplateFilePathOpen= 1
	// 条数类型错误
	ErrorCodeCountType = 2
)

var (
	OutputTypes = map[string]bgs.TemplateOutputSequenceFunc {
		bgs.TYPE_CSV: bgs.CsvOutputSequence,
	}
)

func main() {
	app := cli.NewApp()
	app.Version = "0.0.1"
	app.Name = "bgs"
	app.Usage = `generate sequence form template file path`
	app.HelpName = app.Name
	app.Authors = []cli.Author{
		{
			Name: "andy",
			Email: "1060656096@qq.com",
		},
	}
	app.CustomAppHelpTemplate = getAppHelpTemplate()
	app.Commands = []cli.Command{
		{
			Name: "csv",
			Usage: "The template file path format must be CSV",
			Action: csvAction,
		},
		{
			Name: "excel",
			Usage: "The template file path format must be Excel",
			Action: csvAction,
		},
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:"prefix, p",
			Value:"",
			Usage: "fields prefix",
		},
		cli.StringFlag{
			Name: "exclude-index,e",
			Value: "",
			Usage: `exclude field index list, excluding fields does not generate sequences(1 and 3 columns are not generate sequence example:"0,1,2")`,
		},
		cli.StringFlag{
			Name: "start-sequence, s",
			Value: "0",
			Usage: "generator start sequence",
		},
		cli.StringFlag{
			Name: "output-type, t",
			Value: "csv",
			Usage: "generator sequence output type",
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

// csv动作
func csvAction(ctx * cli.Context)  error {
	tc, err := checkOptions(ctx)
	if err != nil {
		return err
	}
	templateFilePath, count , err := checkArgs(ctx)
	if err != nil {
		return err
	}
	tc.Count = count
	outputSequenceFunc, _ := OutputTypes[tc.OutputType]
	file, err := os.Open(templateFilePath)
	t := bgs.New(file)
	t.Config = &tc
	t.Reader = bgs.CsvReader
	t.OutputSequence = outputSequenceFunc
	t.Run()
	tc.Count = count
	return nil
}

// 检测选项
func checkOptions(ctx *cli.Context) (tc bgs.TemplateConfig, err error) {
	var errMessage strings.Builder
	tc.Prefix = ctx.GlobalString("p")
	if tc.Prefix == "" {
		tc.Prefix = ctx.GlobalString("prefix")
	}

	excludeIndexs := ctx.GlobalString("e")
	if excludeIndexs == "" {
		excludeIndexs = ctx.GlobalString("exclude-index")
	}
	tc.ExcludeIndexs = strings.Split(excludeIndexs, ",")

	tc.StartCount = ctx.GlobalUint64("s")
	if tc.StartCount == 0 {
		tc.StartCount = ctx.GlobalUint64("start-sequence")
	}

	tc.OutputType = ctx.GlobalString("t")
	if tc.OutputType == "" {
		tc.OutputType = ctx.GlobalString("output-type")
	}
	_, ok := OutputTypes[tc.OutputType]
	if !ok {
		errMessage.WriteString(fmt.Sprintf("%s: output type error\n", ctx.App.Name))
		err = cli.NewExitError(errMessage.String(), ErrorCodeRequireArgsNum)
		return
	}
	return
}

// 检测参数
func checkArgs(ctx *cli.Context) (templateFilePath string, count uint64, err error) {
	var errMessage strings.Builder
	// 检测参数错误
	l := ctx.NArg()
	if l < 2 {
		errMessage.WriteString(fmt.Sprintf("%s: templateFilePath and count required\n", ctx.App.Name))
		err = cli.NewExitError(errMessage.String(), ErrorCodeRequireArgsNum)
		return
	}
	// 检测模板文件错误
	templateFilePath = ctx.Args().Get(0)
	file, err := os.Open(templateFilePath)
	if err != nil {
		errMessage.WriteString(fmt.Sprintf("%s: open templateFilePath fail\n", ctx.App.Name))
		errMessage.WriteString(fmt.Sprintf("system error %s", err))
		err = cli.NewExitError(errMessage.String(), ErrorCodeTemplateFilePathOpen)
		return
	}
	defer file.Close()
	// 检测count是否错误
	count, err = strconv.ParseUint(ctx.Args().Get(1), 10, 64)
	if err != nil {
		errMessage.WriteString(fmt.Sprintf("%s: count type error\n", ctx.App.Name))
		errMessage.WriteString(fmt.Sprintf("system error %s", err))
		err = cli.NewExitError(errMessage.String(), ErrorCodeCountType)
		return
	}
	return
}

func getAppHelpTemplate() string {
	return `NAME:
   {{.Name}} - {{.Usage}}
USAGE:
   {{.HelpName}} {{if .VisibleFlags}}[global options]{{end}}{{if .Commands}} command {{end}} <templateFilePath> <count>
   {{if len .Authors}}
AUTHOR:
   {{range .Authors}}{{ . }}{{end}}
   {{end}}{{if .Commands}}
COMMANDS:
{{range .Commands}}{{if not .HideHelp}}   {{join .Names ", "}}{{ "\t"}}{{.Usage}}{{ "\n" }}{{end}}{{end}}{{end}}{{if .VisibleFlags}}
GLOBAL OPTIONS:
   {{range .VisibleFlags}}{{.}}
   {{end}}{{end}}{{if .Copyright }}
COPYRIGHT:
   {{.Copyright}}
   {{end}}{{if .Version}}
VERSION:
   {{.Version}}
   {{end}}
`
}