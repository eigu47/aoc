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

func GetInput(year, day int) ([]string, error) {
	data, err := getData(year, day)
	if err != nil {
		return nil, err
	}
	defer data.Close()

	var input []string
	sc := bufio.NewScanner(data)
	for sc.Scan() {
		input = append(input, sc.Text())
	}

	if err := sc.Err(); err != nil {
		return nil, err
	}

	return input, nil
}

func getData(year, day int) (io.ReadCloser, error) {
	dir := "input"
	filePath := fmt.Sprintf("%s/%d_%02d.txt", dir, year, day)

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
		return nil, fmt.Errorf("http %d: %s", res.StatusCode, data)
	}

	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}

	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return nil, err
	}

	return io.NopCloser(bytes.NewReader(data)), nil
}

var Directions = [4][2]int{
	{-1, 0}, // UP
	{0, 1},  // RGT
	{1, 0},  // DWN
	{0, -1}, // LFT
}

var Diagonals = [4][2]int{
	{-1, -1}, // UP LFT
	{-1, 1},  // UP RGT
	{1, -1},  // DWN LFT
	{1, 1},   // DWN RGT
}
