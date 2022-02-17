package initfuro

import (
	"bufio"
	"fmt"
	"github.com/eclipse/eclipsefuro/furo/pkg/util"
	"github.com/markbates/pkger"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

type TemplateData struct {
	RepositoryName string
	FuroVersion    string
}

func Run(cmd *cobra.Command, args []string) {
	// check if .furo file exist
	if util.FileExists(".furo") {
		fmt.Println("Project already exists")
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Furo spec project init")
	fmt.Println("-----------------------")
	fmt.Println("Please enter your repository name:")
	fmt.Println("This is mostly something like github.com/yourname/sample-specs")

	tpldata := TemplateData{
		RepositoryName: "",
		FuroVersion:    cmd.Flag("version").Value.String(),
	}

	tpldata.RepositoryName = cmd.Flag("repository").Value.String()
	if tpldata.RepositoryName == "" {
		// ask for a repository name
		for {
			fmt.Print("-> ")
			text, _ := reader.ReadString('\n')
			// convert CRLF to LF
			text = strings.Replace(text, "\n", "", -1)

			if strings.Compare("done", text) == 0 {
				os.Exit(0)
			}

			if strings.Compare("exit", text) == 0 {
				os.Exit(0)
			}

			if strings.Compare("", text) == 0 {
				fmt.Println("A empty string is not allowed.")
				fmt.Println("Type done or exit to quit")
			} else {
				tpldata.RepositoryName = text
				break
			}

		}
	}

	fmt.Println("Creating project for ", tpldata.RepositoryName)

	pkger.Walk("/furo/template/", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		p, err := pkger.Parse(path)
		if err != nil {
			return err
		}
		// itself
		if p.Name == "/furo/template" {
			return nil
		}
		targetPath := strings.Replace(p.Name, "/furo/template", ".", 1)

		if info.IsDir() {
			err := os.Mkdir(targetPath, info.Mode())
			if err != nil {
				return nil
			}
		} else {
			f, err := pkger.Open(p.Name)
			if err != nil {
				return err
			}
			var data = make([]byte, info.Size())
			_, err = f.Read(data)
			if err != nil {
				return err
			}

			return ioutil.WriteFile(targetPath, data, info.Mode())
		}

		return nil
	})

	// create .furo with defaults
	parsedTemplate, _ := template.ParseFiles(".furo")
	f, err := os.Create(".furo")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	err = parsedTemplate.Execute(f, tpldata)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}
	e := os.Rename("./go.mod.template", "./go.mod")
	if e != nil {
		log.Fatal(e)
	}
	// create .furo with defaults
	modTemplate, _ := template.ParseFiles("go.mod")
	modf, err := os.Create("go.mod")
	defer modf.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	err = modTemplate.Execute(modf, tpldata)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}

	fmt.Println("Init completed")
	fmt.Println("--------------")
	fmt.Println("")
	fmt.Println("Do not forget to run a furo install to get the correct spec dependencies.")
}
