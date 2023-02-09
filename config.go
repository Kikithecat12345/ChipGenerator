package ChipGenerator

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

type Config struct {
	IO
	Chip
}
type IO struct {
	test1 int
	test2 string
	test3 float32
}
type Chip struct {
	test4 int
	test5 string
	test6 float32
}

func getConfig() *Config {
	a := new(Config)
	err := ini.MapTo(a, "config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	return a
}
