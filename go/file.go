package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main(){
	inputFile, inputError := os.Open("for.go")
	if inputError != nil {
		fmt.Println("An error occurred on opening the File :",inputError)
		return 
	}
	defer inputFile.Close()
	
	inputReader := bufio.NewReader(inputFile)
	for {
		inputString ,readError := inputReader.ReadString('\n')
		fmt.Print(inputString)
		if readError ==io.EOF {
			return
		}
	}
	
}

