package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

var INPUT string = "./input.txt"

func main() {
	fmt.Println("Hello, AOC!")
	leftList, rightList := readFile(INPUT)

	leftListSorted := sortSlice(leftList)
	rightListSorted := sortSlice(rightList)

	fmt.Println(leftListSorted)
	fmt.Println(rightListSorted)
}

func readFile(filepath string) ([]string, []string) {
	var list1, list2 []string
	file, err := os.Open(filepath)
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		line := scanner.Text()
		left := strings.Split(line, "   ")[0]
		list1 = append(list1, left)
		right := strings.Split(line, "   ")[1]
		list2 = append(list2, right)
		// fmt.Println("LEFT:", left, "RIGHT:", right, "\n")

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return list1, list2

}

func sortSlice(list []string) []string {
	sort.Slice(list, func(i, j int) bool {
		return list[i] < list[j]
	})

	return list
}
