package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("day7_input")
	if err != nil {
		log.Fatal("failed to open file")
	}

	data := strings.Split(string(input), "\n")
	if err != nil {
		log.Fatalf("failed split the input lines")
	}

	part1(data)
	part2(data)
}

func part1(data []string) {
	var cmd []string
	files := make(map[string]int64)
	folders := make(map[string]bool)

	for _, line := range data {
		if strings.HasPrefix(line, "$") {
			if strings.HasPrefix(line, "$ cd") {
				args := line[5:]
				switch args {
				case "..":
					if len(cmd) > 0 {
						cmd = cmd[:len(cmd)-1]
					}
				case "/":
					cmd = nil
				default:
					for _, dir := range strings.Split(args, "/") {
						cmd = append(cmd, dir)
					}
				}
				continue
			}
		} else {
			fileData := strings.Split(line, " ")
			fileSizeOrDir := fileData[0]
			fileName := fileData[1]
			if fileSizeOrDir == "dir" {
				continue
			}
			fileSize, err := strconv.ParseInt(fileData[0], 10, 64)
			if err != nil {
				log.Fatalf("failed to parse file size")
			}
			files[fmt.Sprintf("%s/%s", filepath.Join(cmd...), fileName)] = fileSize
		}

		if len(cmd) > 0 {
			folders[filepath.Join(cmd...)] = true
		}
	}

	res := int64(0)
	for folder, _ := range folders {
		folderSize := int64(0)
		for file, size := range files {
			if strings.HasPrefix(filepath.Clean(file), fmt.Sprintf("%s/", folder)) {
				folderSize += size
			}
		}
		if folderSize <= 100000 {
			res += folderSize
		}
	}

	fmt.Println(res)
}

func part2(data []string) {
	var cmd []string
	files := make(map[string]int64)
	folders := make(map[string]bool)
	folderSizes := make(map[string]int64)

	for _, line := range data {
		if strings.HasPrefix(line, "$") {
			if strings.HasPrefix(line, "$ cd") {
				args := line[5:]
				switch args {
				case "..":
					if len(cmd) > 0 {
						cmd = cmd[:len(cmd)-1]
					}
				case "/":
					cmd = nil
				default:
					for _, dir := range strings.Split(args, "/") {
						cmd = append(cmd, dir)
					}
				}
				continue
			}
		} else {
			fileData := strings.Split(line, " ")
			fileSizeOrDir := fileData[0]
			fileName := fileData[1]
			if fileSizeOrDir == "dir" {
				continue
			}
			fileSize, err := strconv.ParseInt(fileData[0], 10, 64)
			if err != nil {
				log.Fatalf("failed to parse file size")
			}
			files[fmt.Sprintf("%s/%s", filepath.Join(cmd...), fileName)] = fileSize
		}

		if len(cmd) > 0 {
			folders[filepath.Join(cmd...)] = true
		}
	}

	rootFolderSize := int64(0)
	for _, size := range files {
		rootFolderSize += size
	}

	sizeRemaining := int64(70000000) - rootFolderSize
	sizeRequiredForUpgrade := int64(30000000) - sizeRemaining

	res := int64(0)
	for _, size := range folderSizes {
		if size >= sizeRequiredForUpgrade {
			if size < res || res == 0 {
				res = size
			}
		}
	}

	fmt.Println(res)
}
