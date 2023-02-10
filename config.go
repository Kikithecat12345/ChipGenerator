// this file is used to read the config.ini file.
// if the config file is not found, the program will use example.config.ini

package ChipGenerator

import (
	"os"

	"gopkg.in/ini.v1"
)

type Config struct {
	IO
	Chip
}
type IO struct {
	ChipTexture   string
	PlaqueTexture string
	Palette       string
	OutputFolder  string
	Format        string
}

// any marked with a * comment are actually big.Ints, but you can't import them with go-ini so they're strings
type Chip struct {
	Mode                string
	Minimum             string
	Maximum             string
	ChipSet             []string
	ChipSetMode         int
	UsePalette          bool
	Patterns            [][]int
	PatternSwitchPoints []string
	PatternMode         int
	PlaqueStart         string
	PlaqueMode          uint8
}
type Text struct {
	ChipTextCenter   []int
	PlaqueTextCenter []int
	ChipTextSize     []int
	PlaqueTextSize   []int
	ChipTextFont     string
	PlaqueTextFont   string
}

func GetConfig() (a *Config, err error) {
	// check if config.ini exists
	// if it doesn't, use example.config.ini
	if _, err := os.Open("config.ini"); os.IsNotExist(err) {
		// use config.ini
		err = ini.MapTo(&a, "config.ini")
		return a, err
	} else {
		// use example.config.ini and throw a warning to the console
		err = ini.MapTo(&a, "example.config.ini")
		return a, err
	}

}
