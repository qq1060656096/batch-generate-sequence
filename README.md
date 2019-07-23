# batch-generate-sequence
batch generate sequence

## compiling
```sh
# compiling mac linux win
make all

# compiling win 64bit
make win64
```

## bgs command help
```sh
$ bin/mac/64/bgs
NAME:
   bgs - generate sequence form template file path
USAGE:
   bgs [global options] command  <templateFilePath> <count>
   
AUTHOR:
   andy <1060656096@qq.com>
   
COMMANDS:
   csv      The template file path format must be CSV
   excel    The template file path format must be Excel
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --prefix value, -p value          fields prefix
   --exclude-index value, -e value   exclude field index list, excluding fields does not generate sequences(1 and 3 columns are not generate sequence example:"0,1,2")
   --start-sequence value, -s value  generator start sequence (default: "0")
   --output-type value, -t value     generator sequence output type (default: "csv")
   --help, -h                        show help
   --version, -v                     print the version
   
VERSION:
   0.0.1
```

## Examples
```sh
# Generate the sequence from the CSV file and export the CSV format
bin/mac/64/bgs -s 10 -e "2,3,4,5,8,9,10,11,12,13,14,15,16" csv bin/templates/dhb.client.csv 20 > bin/templates/dhb.client.new.csv

# Generate the sequence from the excel file and export the excel format
bin/mac/64/bgs -t excel -s 10 -e "2,3,4,5,8,9,10,11,12,13,14,15,16" excel bin/templates/dhb.client.xls 20 > bin/templates/dhb.client.new.xlsx
```