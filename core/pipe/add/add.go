package add

import (
	"bytes"
	"fmt"
	"os"

	"github.com/abdfnx/gomo/constants"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/viper"
	"github.com/tidwall/sjson"
)

func Add(mod string) {
	gomoFile, err := os.ReadFile("gomo.json")

	if err != nil {
		fmt.Println(err)
	}

    addMod, errx := sjson.Set(string(gomoFile), "modules.-1", mod)

	if errx != nil {
		fmt.Println(err)
	}

	viper.SetConfigType("json")

	viper.ReadConfig(bytes.NewBuffer([]byte(addMod)))

	fileErr := viper.WriteConfigAs("gomo.json")

	if fileErr != nil {
		fmt.Println(fileErr)
	}

	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FFF")).
		Background(lipgloss.Color(constants.GREEN_COLOR)).
		PaddingLeft(1).
		PaddingRight(1)

	fmt.Print(style.Render("SUCCESS"))
	fmt.Println(lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(constants.GREEN_COLOR)).SetString(" " + mod +" Added Successfully").String())
}
