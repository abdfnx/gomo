package initx

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/abdfnx/gomo/constants"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/viper"
	"github.com/abdfnx/gomo/core/pipe/add"
)

func Init() {
	viper.AddConfigPath(".")
	viper.SetConfigName("gomo")
	viper.SetConfigType("json")

	viper.SetDefault("modules", []string{})
	viper.SetDefault("cmds.download", "go mod download")
	viper.SetDefault("cmds.update", "go get -u")

	if err := viper.SafeWriteConfig(); err != nil {
		if os.IsNotExist(err) {
			err = viper.WriteConfig()

			if err != nil {
				log.Fatal(err)
			}
		}
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal(err)
		}
	}

	if _, err := os.Stat("gomo.json"); err == nil {
		var style = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFF")).
			Background(lipgloss.Color(constants.GREEN_COLOR)).
			PaddingLeft(1).
			PaddingRight(1)

		fmt.Print(style.Render("SUCCESS"))
		fmt.Println(lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(constants.GREEN_COLOR)).SetString(" Initialization Successful").String())
	} else if errors.Is(err, os.ErrNotExist) {
		var style = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFF")).
			Background(lipgloss.Color(constants.RED_COLOR)).
			PaddingLeft(1).
			PaddingRight(1)

		fmt.Print(style.Render("ERROR"))
		fmt.Println(lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(constants.RED_COLOR)).SetString(" Initialization Failed, try again").String())
	}
}

func InitModule(mod, path string) {
	if path == "" || path == "." {
		path = mod[strings.LastIndex(mod, "/")+1:]
	}

	if err := os.Mkdir(path, os.ModePerm); err != nil {
        log.Fatal(err)
    }

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command("go", "mod", "init", mod)
	cmd.Dir = path
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {
		fmt.Print(stderr.String())
	}

	fmt.Print(stdout.String())

	if _, err := os.Stat(path); err == nil {
		var style = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFF")).
			Background(lipgloss.Color(constants.GREEN_COLOR)).
			PaddingLeft(1).
			PaddingRight(1)

		fmt.Print(style.Render("SUCCESS"))
		fmt.Println(lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(constants.GREEN_COLOR)).SetString(" " + mod + " Module initialized successfully").String())
	} else if errors.Is(err, os.ErrNotExist) {
		var style = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFF")).
			Background(lipgloss.Color(constants.RED_COLOR)).
			PaddingLeft(1).
			PaddingRight(1)

		fmt.Print(style.Render("ERROR"))
		fmt.Println(lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(constants.RED_COLOR)).SetString(" Initialization Failed, try again").String())
	}

	add.Add(path)
}
