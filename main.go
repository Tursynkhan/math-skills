package main

import (
	"bufio"
	"fmt"
	"log"
	"math-skills/calculate/avg"
	"math-skills/calculate/median"
	"math-skills/calculate/standarddev"
	"math-skills/calculate/strtofloat"
	"math-skills/calculate/variance"
	"os"
	"strings"
)

func main() {
	arg := os.Args[1]
	if len(os.Args) != 2 {
		return
	}
	file2, err := os.Open(arg)
	if err != nil {
		log.Fatal(err)
	}
	defer file2.Close()
	scanner := bufio.NewScanner(file2)

	var text string
	for scanner.Scan() {
		text += scanner.Text() + " "
	}
	slice := strings.Fields(text)
	slice2, err := strtofloat.StrToFloat(slice)
	if err != nil {
		fmt.Printf("Str to int error:%v\n", err)
		return
	}
	avg := avg.Average(slice2)
	median.Median(slice2)
	variance := variance.Variance(slice2, avg)
	standarddev.StandardDeviation(slice2, variance)
}
