package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func runCommand(command string) (result string) {
	fmt.Println("\x1b[36m❯", command, "\x1b[0m")

	defer func() {
		result = strings.TrimRight(result, "\n")
	}()

	out, err := exec.Command("sh", "-c", command).CombinedOutput()
	if err != nil {
		fmt.Println("\x1b[31mError:", err, "\x1b[0m")
		result = err.Error()
		return
	}
	result = string(out)
	return
}

var beginRe = regexp.MustCompile(`(//|#)\sWITTGENSTEIN_BEGIN\s` + "`(.+)`")
var endRe = regexp.MustCompile(`(//|#)\sWITTGENSTEIN_END`)

func replace(filepath string) error {
	if !exists(filepath) {
		return fmt.Errorf("%v is not found", filepath)
	}

	fp, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		beginReResult := beginRe.FindAllStringSubmatch(line, -1)
		if beginReResult != nil && len(beginReResult[0]) >= 3 {
			command := beginReResult[0][2]
			fmt.Println(runCommand(command))
		} else {
			endReResult := endRe.FindAllStringSubmatch(line, -1)
			if endReResult != nil {
				fmt.Println(endReResult)
			}
		}
	}

	if err = scanner.Err(); err != nil {
		return err
	}

	return nil
}

func main() {
	flag.Parse()
	args := flag.Args()

	for _, filepath := range args {
		replace(filepath)
	}
}

func exists(filepath string) bool {
	_, err := os.Stat(filepath)
	return err == nil
}
