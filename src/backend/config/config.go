package config

import (
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port          string `yaml:"port"`
		Host          string `yaml:"host"`
		CookieHashKey string `yaml:"cookieHashKey"`
	} `yaml:"server"`
	DatabaseMigration struct {
		Username string `yaml:"user"`
		Password string `yaml:"pass"`
	} `yaml:"databaseMigration"`
	Database struct {
		Name     string `yaml:"name"`
		Username string `yaml:"user"`
		Password string `yaml:"pass"`
	} `yaml:"database"`
}

func LoadConfig(path string) Config {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		panic(err)
	}

	return cfg
}

var rootPath string = ""

func GetRootPath() string {
	if len(rootPath) == 0 {
		rootPath = findModuleRootPath()
		return rootPath
	}

	return rootPath
}

func findModuleRootPath() (roots string) {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	pathBefore := ""
	// Look for enclosing go.mod.
	for {
		files, _ := os.ReadDir(path)

		if err != nil {
			log.Fatal(err)
		}
		for _, f := range files {
			if f.Name() == "go.mod" {
				return path
			}
		}

		pathBefore = path
		path = filepath.Dir(path)

		if path == pathBefore {
			panic("ERROR Not found root path")
		}
	}
}
