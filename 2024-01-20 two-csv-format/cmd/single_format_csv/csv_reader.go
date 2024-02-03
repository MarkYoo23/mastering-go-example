package main

import (
	"bufio"
	"errors"
	"mastering-go-example/internal/reader"
	"os"
	"strings"
)

type CsvReader struct {
}

func NewCsvReader() *CsvReader {
	return &CsvReader{}
}

func (c *CsvReader) ReadFile(path string) (reader.Figure01s, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	scanner := bufio.NewScanner(file)

	err = c.isFigure01(scanner)
	if err != nil {
		return nil, err
	}

	return c.readFigure1(scanner)
}

func (c *CsvReader) isFigure01(scanner *bufio.Scanner) error {

	if !scanner.Scan() {
		return errors.New("first line is empty")
	}

	line := scanner.Text()
	cols := strings.Split(line, ",")

	if len(cols) == 4 {
		return nil
	} else {
		return errors.New("not implemented")
	}

}

func (*CsvReader) readFigure1(scanner *bufio.Scanner) (reader.Figure01s, error) {
	var figures reader.Figure01s

	for scanner.Scan() {
		line := scanner.Text()
		cols := strings.Split(line, ",")

		figure := reader.Figure01{
			Name:       cols[0],
			Surname:    cols[1],
			Tel:        cols[2],
			LastAccess: cols[3],
		}

		figures = append(figures, figure)
	}

	return figures, scanner.Err()
}
