package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func checkPossibleColor(c int, cubes []string) bool {
	for i := 0; i < len(cubes); i++ {
		val, err := strconv.Atoi(cubes[i])
		if err != nil {
			panic(err)
		}

		if val > c {
			return false
		}
	}

	return true
}

func checkPossibleGame(r int, g int, b int, s string) int {
	game := strings.Split(s, ":")

	reRedCubes := regexp.MustCompile("[0-9]+ red")
	reGreenCubes := regexp.MustCompile("[0-9]+ green")
	reBlueCubes := regexp.MustCompile("[0-9]+ blue")

	redCubesShown := reRedCubes.FindAllString(game[1], -1)
	greenCubesShown := reGreenCubes.FindAllString(game[1], -1)
	blueCubesShown := reBlueCubes.FindAllString(game[1], -1)

	reCubes := regexp.MustCompile("[0-9]+")
	redCubes := reCubes.FindAllString(strings.Join(redCubesShown, " "), -1)
	greenCubes := reCubes.FindAllString(strings.Join(greenCubesShown, " "), -1)
	blueCubes := reCubes.FindAllString(strings.Join(blueCubesShown, " "), -1)

	if !checkPossibleColor(r, redCubes) ||
		!checkPossibleColor(g, greenCubes) ||
		!checkPossibleColor(b, blueCubes) {
		return 0
	}

	reGameId := regexp.MustCompile("[0-9]+")
	gameId := reGameId.FindAllString(game[0], -1)

	id, err := strconv.Atoi(gameId[0])
	if err != nil {
		panic(err)
	}

	return id
}

func CubeGameCount() {
	fmt.Println("Starting Cube Game....")

	data, err := os.ReadFile("cubegame.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Data from all games played:")
	fmt.Println(string(data))

	r, g, b := 12, 13, 14
	fmt.Println("Maximum cubes per color:")
	fmt.Printf("Red: %v cubes; Green: %v cubes; Blue: %v cubes\n", r, g, b)

	var sum int = 0
	games := strings.Split(string(data), "\n")
	for i := 0; i < len(games); i++ {
		id := checkPossibleGame(r, g, b, games[i])
		sum += id
	}

	fmt.Printf("Sum of valid game ID's is: %v\n", sum)
}
