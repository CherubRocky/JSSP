package cli

import (
	"strings"
	"fmt"
	"io"
	"os"
)

func MyReadFile(path string) (string, error) {
	f, err := os.Open(path);
	if err != nil {
		return "", err
	}
	builder := strings.Builder{}
	buff := make([]byte, 1024)
	count := 0
	for {
		n, err := f.Read(buff)
		builder.Write(buff[0: n])
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		count += n
	}
	return builder.String(), nil
}
