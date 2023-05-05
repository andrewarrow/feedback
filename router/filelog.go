package router

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/andrewarrow/feedback/files"
)

var openFile *os.File

func (c *Context) Filelog(fields ...any) {
	home := files.UserHomeDir()
	filename := home + "/feedback.log"
	if openFile == nil {
		openFile, _ = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	}
	t := fmt.Sprintf("%d", time.Now().Unix())
	buffer := []string{t}
	for _, thing := range fields {
		buffer = append(buffer, fmt.Sprintf("%v", thing))
	}

	openFile.WriteString(strings.Join(buffer, " ") + "\n")
}
