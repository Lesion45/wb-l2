package dev06

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type CutOptions struct {
	Fields    string
	Delimiter string
	Separated bool
}

func parseArgs() *CutOptions {
	fields := flag.String("f", "", "Выбрать поля (колонки)")
	delim := flag.String("d", "\t", "Использовать другой разделитель")
	sep := flag.Bool("s", false, "Только строки с разделителем")
	flag.Parse()

	return &CutOptions{
		Fields:    *fields,
		Delimiter: *delim,
		Separated: *sep,
	}
}

func readInput() ([]string, error) {
	var lines []string

	for {
		var line string
		_, err := fmt.Scanln(&line)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil, err
		}
		lines = append(lines, line)
	}
	return lines, nil
}

func processLine(line string, options *CutOptions) string {
	columns := strings.Split(line, options.Delimiter)

	if options.Separated && len(columns) == 1 {
		return ""
	}

	if options.Fields != "" {
		fieldIndexes := strings.Split(options.Fields, ",")
		var result []string
		for _, index := range fieldIndexes {
			i, err := strconv.Atoi(index)
			if err != nil || i <= 0 || i > len(columns) {
				continue
			}
			result = append(result, columns[i-1])
		}
		return strings.Join(result, options.Delimiter)
	}

	return line
}
