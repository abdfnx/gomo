package install

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/tidwall/gjson"
)

func Install(isTidy bool) {
	gomoFile, err := os.ReadFile("gomo.json")

	if err != nil {
		fmt.Println(err)
	}

    modules := gjson.Get(string(gomoFile), "modules.#")

	for i := 0; i < int(modules.Int()); i++ {
		mod := gjson.Get(string(gomoFile), "modules." + fmt.Sprint(i)).String()

		var stdout bytes.Buffer
		var stderr bytes.Buffer

		cmd := exec.Command("")
		installCmd := gjson.Get(string(gomoFile), "cmds.install").String()

		if isTidy {
			installCmd = "go mod tidy"
		}

		if runtime.GOOS == "windows" {
			cmd = exec.Command("powershell.exe", installCmd)
		} else {
			cmd = exec.Command("bash", "-c", installCmd)
		}

		cmd.Dir = mod
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr

		err := cmd.Run()

		if err != nil {
			fmt.Print(stderr.String())
		}

		fmt.Print(stdout.String())
	}
}
