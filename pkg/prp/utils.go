package prp

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// getFileContent loads the local content of a file and return the target name
// of the file in the target repository and its contents.
func getFileContent(fileArg string) (targetName string, b []byte, err error) {
	var localFile string
	files := strings.Split(fileArg, ":")

	switch {
	case len(files) < 1:
		return "", nil, errors.New("empty `-files` parameter")
	case len(files) == 1:
		localFile = files[0]
		targetName = files[0]
	default:
		localFile = files[0]
		targetName = files[1]
	}

	b, err = os.ReadFile(localFile)
	return targetName, b, err
}

func getFileContentArray(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// prompt user for confimation
func confirmPrompt(msg string) bool {
	r := bufio.NewReader(os.Stdin)
	var input string

	for {
		fmt.Println(msg)
		input, _ = r.ReadString('\n')
		input = strings.TrimSpace(input)
		input = strings.ToLower(input)

		if input == "y" || input == "yes" || input == "" {
			return true
		}
		if input == "n" || input == "no" {
			return false
		}
	}
}

func runCommand(cwd string, output bool, mainCmd string, cmd ...string) error {
	newCmd := exec.Command(mainCmd, cmd...)

	if len(cwd) > 0 { // change current dir if needed
		newCmd.Dir = cwd
	}

	if output { // needs to stream output
		newCmd.Stdout = os.Stdout
		newCmd.Stderr = os.Stderr
	}

	return newCmd.Run()
}

func saveCommandToFile(outDir *os.File, mainCmd string, cmd ...string) error {
	newCmd := exec.Command(mainCmd, cmd...)

	newCmd.Stdout = outDir
	newCmd.Stderr = os.Stderr

	return newCmd.Run()
}
