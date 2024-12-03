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
		log.Fatalf("%v", err)
	}
	defer data.Close()

	var strs []string
	sc := bufio.NewScanner(data)
	for sc.Scan() {
		strs = append(strs, sc.Text())
	}

	if err := sc.Err(); err != nil {
		log.Fatalf("%v", err)
	}

	return strs
}

func getData(year, day int) (io.ReadCloser, error) {
	dir := "input"
	fileName := fmt.Sprintf("%d_%d.txt", year, day)
	filePath := fmt.Sprintf("%s/%s", dir, fileName)

	file, err := os.Open(filePath)
	if err == nil {
		return file, nil
	}

	if !errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("failed to open file: %w", err)
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

	if res.StatusCode != 200 {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(body))
	}

	// save
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return nil, err
		}
	}

	if err := os.WriteFile(filePath, data, 0666); err != nil {
		log.Print("Could not save on disk", err)
	}

	return io.NopCloser(bytes.NewReader(data)), nil
}

// type Input struct {
// 	session string
// }

// func NewInput() (*Input, error) {
// 	err := godotenv.Load()
// 	if err != nil {
// 		return nil, err
// 	}

// 	session := os.Getenv("SESSION_ID")
// 	if session == "" {
// 		return nil, errors.New("no session in environment")
// 	}

// 	return &Input{session: session}, nil
// }

// func (i *Input) Reader(year, day int) (io.ReadCloser, error) {
// 	dir := "input"
// 	fileName := fmt.Sprintf("%d_%d.txt", year, day)
// 	file, err := os.Open(fmt.Sprintf("%s/%s", dir, fileName))

// 	if errors.Is(err, os.ErrNotExist) {
// 		// fetch
// 		log.Printf("Fetching data, year %d day %d", year, day)
// 		req, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day), nil)
// 		if err != nil {
// 			return nil, err
// 		}

// 		req.AddCookie(&http.Cookie{
// 			Name:  "session",
// 			Value: i.session,
// 		})

// 		res, err := http.DefaultClient.Do(req)
// 		if err != nil {
// 			return nil, err
// 		}
// 		defer res.Body.Close()

// 		if res.StatusCode != 200 {
// 			body, err := io.ReadAll(res.Body)
// 			if err != nil {
// 				return nil, err
// 			}
// 			return nil, errors.New(string(body))
// 		}

// 		// save
// 		data, err := io.ReadAll(res.Body)
// 		if err != nil {
// 			return nil, err
// 		}

// 		if _, err := os.Stat(dir); os.IsNotExist(err) {
// 			os.MkdirAll(dir, os.ModePerm)
// 		}

// 		if err := os.WriteFile(fmt.Sprintf("%s/%s", dir, fileName), data, 0666); err != nil {
// 			log.Print("Could not save on disk", err)
// 		}

// 		return io.NopCloser(bytes.NewReader(data)), nil

// 	} else if err != nil {
// 		return nil, err

// 	} else {
// 		// log.Printf("Read from file %s", fileName)
// 		return file, nil
// 	}
// }

// func (i *Input) Bytes(year, day int) ([]byte, error) {
// 	rc, err := i.Reader(year, day)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rc.Close()

// 	return io.ReadAll(rc)
// }

// func (i *Input) Strings(year, day int) ([]string, error) {
// 	rc, err := i.Reader(year, day)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rc.Close()

// 	strs := make([]string, 0)
// 	sc := bufio.NewScanner(rc)
// 	for sc.Scan() {
// 		strs = append(strs, sc.Text())
// 	}

// 	return strs, nil
// }

// func (i *Input) Ints(year, day int) ([]int, error) {
// 	rc, err := i.Reader(year, day)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rc.Close()

// 	ints := make([]int, 0)
// 	sc := bufio.NewScanner(rc)
// 	for sc.Scan() {
// 		i, err := strconv.Atoi(sc.Text())
// 		if err != nil {
// 			return nil, fmt.Errorf("at index %d: %w", len(ints), err)
// 		}
// 		ints = append(ints, i)
// 	}

// 	return ints, nil
// }

// func (i *Input) Floats(year, day int) ([]float64, error) {
// 	rc, err := i.Reader(year, day)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rc.Close()

// 	floats := make([]float64, 0)
// 	sc := bufio.NewScanner(rc)
// 	for sc.Scan() {
// 		i, err := strconv.ParseFloat(sc.Text(), 64)
// 		if err != nil {
// 			return nil, fmt.Errorf("at index %d: %w", len(floats), err)
// 		}
// 		floats = append(floats, i)
// 	}

// 	return floats, nil
// }

// func (i *Input) Int64s(year, day, base int) ([]int64, error) {
// 	rc, err := i.Reader(year, day)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rc.Close()

// 	int64s := make([]int64, 0)
// 	sc := bufio.NewScanner(rc)
// 	for sc.Scan() {
// 		i, err := strconv.ParseInt(sc.Text(), base, 64)
// 		if err != nil {
// 			return nil, fmt.Errorf("at index %d: %w", len(int64s), err)
// 		}
// 		int64s = append(int64s, i)
// 	}

// 	return int64s, nil
// }
