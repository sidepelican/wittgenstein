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
	"strings"
)

func runCommand(command string) string {
	fmt.Println("\x1b[36m❯", command, "\x1b[0m")

	out, _ := exec.Command("sh", "-c", command).CombinedOutput()
	// if err != nil {
	// 	fmt.Println("\x1b[31m❯ Error:", err, "\x1b[0m")
	// 	return err.Error()
	// }

	return string(out)
}

var beginRe = regexp.MustCompile(`(//|#)\sWITTGENSTEIN_BEGIN\s` + "`(.+)`")
var endRe = regexp.MustCompile(`(//|#)\sWITTGENSTEIN_END`)

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

		endReResult := endRe.FindAllStringSubmatch(line, -1)
		if endReResult != nil {
			skip = false
		}

		if !skip {
			fw.WriteString(line)
			fw.WriteString("\n")
		}

		beginReResult := beginRe.FindAllStringSubmatch(line, -1)
		if beginReResult != nil && len(beginReResult[0]) >= 3 {
			command := beginReResult[0][2]
			commandResult := runCommand(command)

			fw.WriteString(commandResult)
			skip = true
		}
	}

	if err = scanner.Err(); err != nil {
		return err
	}

	diff := strings.TrimRight(runCommand("diff "+filename+" "+fw.Name()), "\n")
	if len(diff) > 0 {
		fmt.Print("updated: ")
		fmt.Println(diff)
	}

	return nil
}

func main() {
	flag.Parse()
	args := flag.Args()

	for _, filename := range args {
		err := replace(filename)
		fmt.Println(err)
	}
}
