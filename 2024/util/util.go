package util

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func GetInput(year, day int) []string {
	data, err := getData(year, day)
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	var input []string
	sc := bufio.NewScanner(data)
	for sc.Scan() {
		input = append(input, sc.Text())
	}

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	return input
}

func getData(year, day int) (io.ReadCloser, error) {
	dir := "input"
	filePath := fmt.Sprintf("%s/%d_%d.txt", dir, year, day)

	file, err := os.Open(filePath)
	if err == nil {
		return file, nil
	}

	if !errors.Is(err, os.ErrNotExist) {
		return nil, err
	}

	log.Printf("Fetching data, year %d day %d", year, day)
	req, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day), nil)
	if err != nil {
		return nil, err
	}

	if err = godotenv.Load(); err != nil {
		return nil, err
	}

	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: os.Getenv("SESSION_ID"),
	})

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(string(data))
	}

	// save
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return nil, err
		}
	}

	if err := os.WriteFile(filePath, data, 0666); err != nil {
		return nil, err
	}

	return io.NopCloser(bytes.NewReader(data)), nil
}

func IsInBounds[T any](pos [2]int, grid T) bool {
	switch v := any(grid).(type) {
	case [][]T:
		return pos[0] >= 0 && pos[0] < len(v) && pos[1] >= 0 && pos[1] < len(v[0])
	case []string:
		return pos[0] >= 0 && pos[0] < len(v) && pos[1] >= 0 && pos[1] < len(v[0])
	default:
		return false
	}
}
