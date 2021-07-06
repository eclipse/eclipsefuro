package util

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"path"
	"regexp"
	"strings"
)

func GetDependencyList() []string {
	deps := []string{}
	for _, d := range viper.GetStringSlice("dependencies") {
		dep := ParseDependency(d)

		// should only have 2 parts (repo, version)
		if dep.Kind == UNKNOWN {
			fmt.Println(ScanForStringPosition(d, "./.furo"), "Config Error")
			log.Fatal("Config error or dependency not installed. Maybe you should run spectools install")
		}

		// todo check for existence of p and give spectools install hint

		// load config to resolve Message and Service dirs
		depconf := viper.New()
		depconf.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
		depconf.AddConfigPath(dep.DependencyPath)
		depconf.SetConfigName(".furo")
		err := depconf.ReadInConfig()
		if err == nil {

			sdir := depconf.GetString("specDir")
			if sdir != "" {
				if _, err := os.Stat(path.Join(dep.DependencyPath, sdir)); !os.IsNotExist(err) {
					deps = append(deps, path.Join(dep.DependencyPath, sdir))
				}
			}

		} else {
			// no .furo config in target dir, use the complete path
			deps = append(deps, dep.DependencyPath)
		}

	}
	return deps
}

type DependencyKind int

const (
	UNKNOWN DependencyKind = iota
	GIT
	FILESYSTEM
)

type Dependency struct {
	Kind           DependencyKind
	Path           string
	DependencyPath string
	Repository     string
	Version        string
	Original       string
}

func ParseDependency(dep string) Dependency {
	d := Dependency{Original: dep, Kind: UNKNOWN}
	// parse dep string  (.*[:]\/*(.*).git)\s(.*)  => 1:Repository, 2:Path, 3: Version

	regex := regexp.MustCompile(`(.*(:\/\/|@)(.*).git)\s([^\s]*)`)
	matches := regex.FindStringSubmatch(dep)
	if len(matches) == 0 {
		// is file
		d.Kind = FILESYSTEM
		d.DependencyPath = path.Join(viper.GetString("dependenciesDir"), dep)
		d.Path = dep
	}
	if len(matches) == 5 {
		d.Kind = GIT
		d.Repository = matches[1]
		if matches[2] == "@" {
			d.DependencyPath = path.Join(viper.GetString("dependenciesDir"), strings.Replace(matches[3], ":", "/", 1))
			d.Path = strings.Replace(matches[3], ":", "/", 1)
		} else {
			d.DependencyPath = path.Join(viper.GetString("dependenciesDir"), matches[3])
			d.Path = matches[3]
		}

		d.Version = matches[4]
	}

	return d
}
