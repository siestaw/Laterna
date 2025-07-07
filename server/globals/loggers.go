package globals

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var (
	HTTPLogger *log.Logger
	DBLogger   *log.Logger
	flags      int
)

func InitLoggers() {
	var logOutput io.Writer = os.Stdout

	if AppConfig.FileLogging {

		os.Mkdir("logs", os.ModePerm)

		timestamp := time.Now().Format("2006-01-02_15-04-05")
		filename := fmt.Sprintf("logs/log_%s.txt", timestamp)

		file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal("An error occured while loading logs.txt")
		}

		logOutput = io.MultiWriter(os.Stdout, file)
		log.SetOutput(logOutput)
	}

	if AppConfig.VerboseLogging {
		flags = log.Ldate | log.Ltime | log.Lshortfile
	}

	HTTPLogger = log.New(logOutput, "[HTTP] ", flags)
	DBLogger = log.New(logOutput, "[DATABASE] ", flags)
}
