package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
	"unicode"
)

func wordCounts(filename string) (int, int, int, int, map[string]int, error) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return 0, 0, 0, 0, nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()
	// Initialize variables
	wordCounts := make(map[string]int)
	totalWords := 0
	totalCharsWithSpaces := 0
	totalCharsWithoutSpaces := 0
	// Define a regex for tokenizing words
	wordRegex := regexp.MustCompile(`\b\w+\b`)
	// Read file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		totalCharsWithSpaces += len(line)
		for _, char := range line {
			if !unicode.IsSpace(char) {
				totalCharsWithoutSpaces++
			}
		}
		words := wordRegex.FindAllString(strings.ToLower(line), -1) // Convert to lowercase
		for _, word := range words {
			wordCounts[word]++
			totalWords++
		}
	}
	if err := scanner.Err(); err != nil {
		return 0, 0, 0, 0, nil, fmt.Errorf("error reading file: %w", err)
	}

	distinctWords := len(wordCounts)
	return totalWords, distinctWords, totalCharsWithSpaces, totalCharsWithoutSpaces, wordCounts, nil
}

func getTopNWords(wordCounts map[string]int, n int) []string {
	type wordFreq struct {
		word  string
		count int
	}
	wordFreqs := make([]wordFreq, 0, len(wordCounts))
	for word, count := range wordCounts {
		wordFreqs = append(wordFreqs, wordFreq{word, count})
	}
	// Sort by count (descending) and alphabetically for ties
	sort.Slice(wordFreqs, func(i, j int) bool {
		if wordFreqs[i].count == wordFreqs[j].count {
			return wordFreqs[i].word < wordFreqs[j].word
		}
		return wordFreqs[i].count > wordFreqs[j].count
	})
	topN := make([]string, 0, n)
	for i := 0; i < len(wordFreqs) && i < n; i++ {
		topN = append(topN, fmt.Sprintf("%s: %d", wordFreqs[i].word, wordFreqs[i].count))
	}

	return topN
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: wordcount <filename> [-n number]")
		return
	}
	filename := os.Args[1]
	var topN int
	flagSet := flag.NewFlagSet("wordcount", flag.ContinueOnError)
	flagSet.IntVar(&topN, "n", 5, "Number of top occurring words to display")
	err := flagSet.Parse(os.Args[2:])
	if err != nil {
		fmt.Println("Usage: wordcount <filename> [-n number]")
		return
	}
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Printf("Error: File '%s' not found.\n", filename)
		return
	}
	totalWords, distinctWords, charsWithSpaces, charsWithoutSpaces, wordCounts, err := wordCounts(filename)
	if err != nil {
		fmt.Printf("Error processing file: %v\n", err)
		return
	}
	fmt.Printf("Total word count (including repeated words): %d\n", totalWords)
	fmt.Printf("Total number of distinct words: %d\n", distinctWords)
	fmt.Printf("Total number of characters (including spaces): %d\n", charsWithSpaces)
	fmt.Printf("Total number of characters (excluding spaces): %d\n", charsWithoutSpaces)
	if topN > 0 {
		fmt.Printf("Top %d word frequencies:\n", topN)
		topWords := getTopNWords(wordCounts, topN)
		for _, word := range topWords {
			fmt.Println("    " + word)
		}
	} else {
		fmt.Println("Word frequencies (sorted alphabetically):")
		keys := make([]string, 0, len(wordCounts))
		for key := range wordCounts {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for _, key := range keys {
			fmt.Printf("    %s: %d\n", key, wordCounts[key])
		}
	}
}
