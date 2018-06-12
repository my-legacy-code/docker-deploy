package core

import (
	"log"
	"os"
	"time"
)

const prefix = "[DOCKER DEPLOY] "
const flag = log.Ldate | log.Ltime | log.Llongfile | log.LUTC

func setupLogger() {
	log.SetPrefix(prefix)
	log.SetFlags(flag)
}

func makeErrLogger() *log.Logger {
	fileName := time.Now().Format(time.RFC3339) + ".err"
	file, err := os.Create("errors/" + fileName)
	if err != nil {
		log.Fatal(err)
	}
	return log.New(file, prefix, flag)
}
