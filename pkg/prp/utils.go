package prp

import (
	"errors"
	"os"
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