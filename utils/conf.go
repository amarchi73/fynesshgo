package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

func Configurazione() {
	userdir, _ := os.UserHomeDir()
	userdir = "."
	cfg, err := ini.Load(userdir + "/.fynesshgo.config")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	fmt.Println(cfg)

	fmt.Println(cfg.Section("").Key("owner").Value())
}
