package initx

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/abdfnx/gomo/constants"
	"github.com/spf13/viper"
)

func Init() {
	viper.AddConfigPath(".")
	viper.SetConfigName("gomo")
	viper.SetConfigType("json")

	viper.SetDefault("modules", []string{})
	viper.SetDefault("cmds.download", "go mod download")

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
