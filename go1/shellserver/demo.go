package main

import (
	"bytes"
	"os"
	"os/exec"
	"os/user"
	"path"
	"strings"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second * 30)
	for range ticker.C {
		u, _ := user.Current()

		result, err := shell("ps -ef | grep ./" + os.Args[1] + " | grep -v \"$0\" | grep -w " + u.Username + " | grep -v \"grep\" | awk '{print $2}'")
		if err != nil {
			continue
		}

		result = strings.Replace(result, "\n", " ", -1)
		if result == "" {
			shstart("./start.sh")
		}

		_, err = os.Stat(path.Join("./data", os.Args[1]+".zip"))
		if err != nil {
			continue
		}

		result, err = shell("ps -ef | grep ./" + os.Args[1] + " | grep -v \"$0\" | grep -w " + u.Username + " | grep -v \"grep\" | awk '{print $2}'")
		if err != nil {
			continue
		}
		if result != "" {
			result, err = shell("kill -9 " + strings.Replace(result, "\n", " ", -1))
			if err != nil {
				continue
			}
		}
		shell("unzip -o ./data/" + os.Args[1] + ".zip")
		shell("chmod 755 -R *")
		shell("rm -rf ./data/" + os.Args[1] + ".zip")
		shstart("./start.sh")
	}
}

func shell(s string) (string, error) {
	cmd := exec.Command("/bin/bash", "-c", s)

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	return out.String(), err
}

func shstart(s string) {
	cmd := exec.Command("/bin/bash", "-c", s)
	cmd.Run()
}
