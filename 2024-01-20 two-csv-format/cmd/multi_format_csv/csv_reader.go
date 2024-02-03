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

func (c *CsvReader) ReadFile(path string) (reader.ControllableFile, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	scanner := bufio.NewScanner(file)

	figureType, err := c.findFigureType(scanner)

	switch figureType {
	case reader.FigureType01:
		return c.readFigure1(scanner)
	case reader.FigureType02:
		return c.readFigure2(scanner)
	default:
		return nil, errors.New("not implemented")
	}
}

func (c *CsvReader) findFigureType(scanner *bufio.Scanner) (reader.FigureType, error) {

	if !scanner.Scan() {
		return "", errors.New("first line is empty")
	}

	line := scanner.Text()
	cols := strings.Split(line, ",")

	if len(cols) == 4 {
		return reader.FigureType01, nil
	} else if len(cols) == 5 {
		return reader.FigureType02, nil
	} else {
		return "", errors.New("not implemented")
	}

}

func (*CsvReader) readFigure1(scanner *bufio.Scanner) (reader.ControllableFile, error) {
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

func (c *CsvReader) readFigure2(scanner *bufio.Scanner) (reader.ControllableFile, error) {
	var figures reader.Figure02s

	for scanner.Scan() {
		line := scanner.Text()
		cols := strings.Split(line, ",")

		figure := reader.Figure02{
			Name:       cols[0],
			Surname:    cols[1],
			Areacode:   cols[2],
			Tel:        cols[3],
			LastAccess: cols[4],
		}

		figures = append(figures, figure)
	}

	return figures, scanner.Err()
}
