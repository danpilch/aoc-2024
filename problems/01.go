package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readPairsFromFile(filePath string) ([]int, []int, error) {
	var aValues, bValues []int
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Split by whitespace
		fields := strings.Fields(line)
		if len(fields) != 2 {
			return nil, nil, fmt.Errorf("invalid line format: %s", line)
		}

		// Convert to int
		a, err1 := strconv.Atoi(fields[0])
		b, err2 := strconv.Atoi(fields[1])
		if err1 != nil || err2 != nil {
			return nil, nil, fmt.Errorf("failed to parse: %s", line)
		}

		aValues = append(aValues, a)
		bValues = append(bValues, b)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return aValues, bValues, nil
}

func computeMaxDistance(aValues []int, bValues []int) int {
	sort.Ints(aValues)
	sort.Ints(bValues)

	totalDistance := 0

	for i := range aValues {
		totalDistance += abs(aValues[i] - bValues[i])
	}

	return totalDistance
}

func computeSimilarityScore(aValues []int, bValues []int) int {
	similarity := 0
	for _, n := range aValues {
		oc := countOccurances(n, bValues)
		if oc > 0 {
			similarity += abs(n * oc)
		}
	}
	return similarity
}

func countOccurances(target int, slice []int) int {
	count := 0

	for _, n := range slice {
		if n == target {
			count++
		}
	}

	return count
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func main() {
	aValues, bValues, err := readPairsFromFile("../inputs/01")
	if err != nil {
		log.Fatalf("Failed to parse file: %v", err)
	}

	maxDistance := computeMaxDistance(aValues, bValues)
	fmt.Printf("The furthest distance is: %d\n", maxDistance)

	similarity := computeSimilarityScore(aValues, bValues)
	fmt.Printf("The similarity score is: %d\n", similarity)
}
