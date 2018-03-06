package main

import (
    "bufio"
    "fmt"
    "os"
    "io"
)

func main() {
    inputFile, _ := os.Open("goprogram")
    outputFile, _ := os.OpenFile("goprogramT", os.O_WRONLY|os.O_CREATE, 0666)
    defer inputFile.Close()
    defer outputFile.Close()
    inputReader := bufio.NewReader(inputFile)
    outputWriter := bufio.NewWriter(outputFile)
    for {
        inputString, _, readerError := inputReader.ReadLine()
        if readerError == io.EOF {
            fmt.Println("EOF")
            break
        }
        outputString := string(inputString[2:5])+"\n"
		fmt.Println(outputString)
        _, err := outputWriter.WriteString(outputString)
        if err != nil {
            fmt.Println(err)
            continue
        }
    }
	outputWriter.Flush()
    fmt.Println("Conversion done")
}
