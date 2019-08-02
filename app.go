package termapp

import (
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/chzyer/readline"
)

func tokenize(line string) []string {
	line = strings.Trim(line, " ")
	re, _ := regexp.CompilePOSIX(" +")
	return re.Split(line, -1)
}

func Loop(instance *readline.Instance,
	action func(tokens []string) (interface{}, bool)) interface{} {

	for {
		line, err := instance.Readline()
		if err == readline.ErrInterrupt {
			if len(line) == 0 {
				os.Exit(0)
			} else {
				continue
			}
		} else if err == io.EOF {
			os.Exit(0)
		}

		tokens := tokenize(line)
		if rt, quit := action(tokens); quit {
			return rt
		}
	}
}
