package utils

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func GetTopPath() string {
	value, exists := os.LookupEnv("TOP_PATH")
	if exists {
		return value
	} else {
		InputTopPathToEnv()
		return GetTopPath()
	}
}

func InputTopPathToEnv() {
	cmd := exec.Command("go", "env", "GOMOD")
	out, err := cmd.Output()
	if err != nil {
		log.Println(err)
		return
	}
	gomod := strings.TrimSpace(string(out))
	newPath := strings.TrimSuffix(gomod, "go.mod")
	if err := os.Setenv("TOP_PATH", newPath); err != nil {
		log.Println(err)
	}
}
