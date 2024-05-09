package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func getCalibrationUnits(s string) string {
	re := regexp.MustCompile("[0-9]+")
	var x, y string

	digits := re.FindAllString(s, -1)

	if len(digits) == 0 {
		return "0"
	}

	x = digits[0]
	y = digits[len(digits)-1]

	return string(x[0]) + string(y[len(y)-1])
}

func Calibration() {
	fmt.Println("Starting Calibration....")

	data, err := os.ReadFile("calibration.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Data from calibration document:")
	fmt.Println(string(data))

	index := strings.Split((string(data)), "\n")

	var sum int
	for i := 0; i < len(index); i++ {
		val, err := strconv.Atoi(getCalibrationUnits(index[i]))
		if err != nil {
			panic(err)
		}

		sum += val
	}

	fmt.Printf("Finished Calibration with value: %v", sum)
}
