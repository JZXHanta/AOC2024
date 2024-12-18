package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

const INPUT string = "./input.txt"

func main() {
	leftList, rightList := readFile(INPUT)

	leftListSorted := sortSlice(leftList)
	rightListSorted := sortSlice(rightList)

	diffArray := getDifferences(leftListSorted, rightListSorted)

	finalAnswer := addAllArrayValues(diffArray)
	fmt.Println("FINAL ANSWER:", finalAnswer)
}

func readFile(filepath string) (list1, list2 []string) {
	file, err := os.Open(filepath)
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		left := strings.Split(line, "   ")[0]
		list1 = append(list1, left)
		right := strings.Split(line, "   ")[1]
		list2 = append(list2, right)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return
}

func sortSlice(list []string) []string {
	sort.Slice(list, func(i, j int) bool {
		return list[i] < list[j]
	})

	return list
}

func getDifferences(list1, list2 []string) (differenceArray []int) {

	for i := 0; i < len(list1); i++ {
		leftVal, err := strconv.Atoi(list1[i])
		if err != nil {
			panic(err)
		}
		rightVal, err := strconv.Atoi(list2[i])
		differenceArray = append(differenceArray, leftVal-rightVal)
	}

	return
}

func addAllArrayValues(list []int) (finalAnswer int) {
	for i := 0; i < len(list); i++ {
		finalAnswer += int(math.Abs(float64(list[i])))
	}

	return
}
