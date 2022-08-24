package utils

import (
	"fmt"
	"fyneapp/types"
	"gopkg.in/ini.v1"
	"os"
	"strconv"
)

func trovaEnvironment() (bool, string) {
	s := os.Getenv("FYNESSHGO_CONFPATH")
	if s != "" {
		types.ConfigPath = s
	} else {
		types.ConfigPath, _ = os.UserHomeDir()
	}
	s = os.Getenv("FYNESSHGO_DATAPATH")
	if s != "" {
		types.DataPath = s
	} else {
		homedir, _ := os.UserHomeDir()
		_, err := os.ReadDir(homedir + "/.local")
		if err != nil {
			fmt.Println(err.Error())
			err = os.MkdirAll(homedir+"/.fynesshgo.data", 0755)
			types.DataPath = homedir + "/.fynesshgo.data"
		} else {
			err = os.MkdirAll(homedir+"/.local/share/fynesshgo", 0755)
			types.DataPath = homedir + "/.local/share/fynesshgo"
		}
		if err != nil {
			fmt.Println(err)
			return false, "Non posso creare la directory " + types.DataPath
		}
		//types.ConfigPath, _ =
	}
	_, err := os.ReadFile(types.DataPath + "/elencodati.json")
	if err != nil {
		fmt.Println("Creo file: " + types.DataPath + "/elencodati.json")
		os.WriteFile(types.DataPath+"/elencodati.json", []byte(""), 0666)
	}
	return true, ""
}
func ConfigurazionePresente() int {
	ok, msg := trovaEnvironment()
	if !ok {
		fmt.Println(msg)
		return -3
	}

	cfg, err := ini.Load(types.ConfigPath + "/.fynesshgo.config")
	if err != nil {
		fmt.Printf("Creo il file: %v", types.ConfigPath+"/.fynesshgo.config")
		s := "owner = fynesshgo\n"
		os.WriteFile(types.ConfigPath+"/.fynesshgo.config", []byte(s), 0666)
		return -1
	}
	//fmt.Println(cfg)
	if cfg.Section("").Key("owner").Value() != "fynesshgo" {
		fmt.Printf("il file: %v non è valido", types.ConfigPath+"/.fynesshgo.config")
		return -2
	}
	if cfg.Section("").Key("default").Value() != "" {
		types.FileJSON = cfg.Section("").Key("default").Value()
		types.DefaultFileJson = types.FileJSON
		fmt.Println("default: " + types.FileJSON)
	}
	if cfg.Section("").Key("comandopass").Value() != "" {
		ComandoPass = cfg.Section("").Key("comandopass").Value()
	}
	if cfg.Section("").Key("passpath").Value() != "" {
		PassPath = cfg.Section("").Key("passpath").Value()
	}
	if cfg.Section("").Key("savetype").Value() != "" {
		types.SaveTypeDef, _ = strconv.Atoi(cfg.Section("").Key("savetype").Value())
	}
	return 0
}
func SalvaConf(key string, valore string) {
	cfg, _ := ini.Load(types.ConfigPath + "/.fynesshgo.config")
	cfg.Section("").Key(key).SetValue(valore)
	cfg.SaveTo(types.ConfigPath + "/.fynesshgo.config")
}
