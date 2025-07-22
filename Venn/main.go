package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Track struct {
	ID   int
	BPM  int
	Hour int
}

func parseHour(hourStr string) int {
	parts := strings.Split(hourStr, ":")
	if len(parts) < 2 {
		return 0
	}
	hour, _ := strconv.Atoi(parts[0])
	if strings.Contains(hourStr, "PM") && hour != 12 {
		hour += 12
	} else if strings.Contains(hourStr, "AM") && hour == 12 {
		hour = 0
	}
	return (hour % 12) // 12-hour clock for Oli's system
}

func isValidTransition(a, b Track) bool {
	// Hour rule
	hourDiff := int(math.Abs(float64(a.Hour - b.Hour)))
	hourValid := hourDiff == 0 || hourDiff == 1 || hourDiff == 11 // wraparound
	if !hourValid {
		return false
	}

	// BPM rule
	bpmDiff := b.BPM - a.BPM
	return bpmDiff >= -2 && bpmDiff <= 5
}

func loadTracks(filename string) []Track {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var tracks []Track
	firstLine := true

	for scanner.Scan() {
		line := scanner.Text()
		if firstLine {
			firstLine = false
			continue // skip header
		}

		parts := strings.Split(line, "\t")
		if len(parts) < 4 {
			continue
		}

		id, _ := strconv.Atoi(parts[0])
		bpm, _ := strconv.Atoi(parts[2])
		hour := parseHour(parts[3])

		tracks = append(tracks, Track{ID: id, BPM: bpm, Hour: hour})
	}

	return tracks
}

func buildGraph(tracks []Track) map[int][]int {
	graph := make(map[int][]int)
	for i, a := range tracks {
		for j, b := range tracks {
			if i != j && isValidTransition(a, b) {
				graph[a.ID] = append(graph[a.ID], b.ID)
			}
		}
	}
	return graph
}

func dfs(graph map[int][]int, tracksMap map[int]Track, path []int, visited map[int]bool, targetLen int) []int {
	if len(path) == targetLen {
		return path
	}

	last := path[len(path)-1]
	for _, next := range graph[last] {
		if visited[next] {
			continue
		}
		visited[next] = true
		newPath := dfs(graph, tracksMap, append(path, next), visited, targetLen)
		if newPath != nil {
			return newPath
		}
		visited[next] = false
	}
	return nil
}

func main() {
	tracks := loadTracks("DJTrackLibrary.csv")
	graph := buildGraph(tracks)

	tracksMap := make(map[int]Track)
	for _, t := range tracks {
		tracksMap[t.ID] = t
	}

	for _, t := range tracks {
		visited := make(map[int]bool)
		visited[t.ID] = true
		result := dfs(graph, tracksMap, []int{t.ID}, visited, 33)
		if result != nil {
			fmt.Println("Found valid setlist:")
			for i, id := range result {
				if i > 0 {
					fmt.Print(",")
				}
				fmt.Print(id)
			}
			fmt.Println()
			return
		}
	}

	fmt.Println("No valid setlist found.")
}
