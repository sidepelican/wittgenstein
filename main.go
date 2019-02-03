package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
)

var debug bool = false

func runCommand(command string) string {
	if debug {
		fmt.Println("\x1b[36m❯", command, "\x1b[0m")
	}

	out, _ := exec.Command("sh", "-c", command).CombinedOutput()
	// if err != nil {
	// 	fmt.Println("\x1b[31m❯ Error:", err, "\x1b[0m")
	// 	return err.Error()
	// }

	return string(out)
}

func isSameFile(file1, file2 string) bool {
	diff := runCommand("diff " + file1 + " " + file2)
	return len(diff) == 0
}

var beginRe = regexp.MustCompile(`(//|#)\s*WITTGENSTEIN_BEGIN\s*` + "`(.+)`")
var endRe = regexp.MustCompile(`(//|#)\s*WITTGENSTEIN_END`)

func replace(filename string) error {
	fp, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fp.Close()

	fw, err := ioutil.TempFile("", filepath.Base(filename)+".tmp")
	if err != nil {
		return err
	}
	defer fw.Close()

	scanner := bufio.NewScanner(fp)
	var skip = false
	for scanner.Scan() {
		line := scanner.Text()

		if skip {
			endReResult := endRe.FindAllStringSubmatch(line, -1)
			if endReResult != nil {
				skip = false
			}
		}

		if !skip {
			fw.WriteString(line)
			fw.WriteString("\n")
		}

		if !skip {
			beginReResult := beginRe.FindAllStringSubmatch(line, -1)
			if beginReResult != nil && len(beginReResult[0]) >= 3 {
				command := beginReResult[0][2]
				commandResult := runCommand(command)

				fw.WriteString(commandResult)
				skip = true
			}
		}
	}

	if err = scanner.Err(); err != nil {
		return err
	}

	if !isSameFile(filename, fw.Name()) {
		runCommand(`\cp -f ` + fw.Name() + " " + filename)
	}

	return nil
}

func main() {
	flag.BoolVar(&debug, "d", false, "show debug log")
	flag.Parse()
	args := flag.Args()

	hasError := false
	for _, filename := range args {
		err := replace(filename)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			hasError = true
		}
	}

	if hasError {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}
