package domain

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

// FileSource to load values from file
type FileValueSource struct {
	filePath string
}

func (fs *FileValueSource) Load() ([]int, error) {
	log.Infof("Loading Source File: %s", fs.filePath)
	file, err := os.Open(fs.filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var values []int

	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			return nil, fmt.Errorf("failed to convert '%s' to int: %v", line, err)
		}
		values = append(values, num)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return values, nil
}

func NewFileSource(FilePath string) *FileValueSource {
	return &FileValueSource{filePath: FilePath}
}
