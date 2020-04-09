package tools

import (
    "encoding/csv"
    "fmt"
    "io"
	"os"
	
)
func ReadCsv(path string) [][]string  {
	file, err := os.Open(path)
	var arr [][]string
    if err != nil {
        fmt.Println("Error:", err)
        //return
    }
    defer file.Close()
    reader := csv.NewReader(file)
    for {
        record, err := reader.Read()
        if err == io.EOF {
            break
        } else if err != nil {
            fmt.Println("Error:", err)
            //return
		}
		arr = append(arr, record)
        //fmt.Println(arr) // record has the type []string
	}
	//fmt.Println(arr)
	
	return arr
}