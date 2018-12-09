package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
)

var logger = log.New(os.Stdout, "", log.LstdFlags)

type Packers struct {
	Root string
	Ps   []*Packer
}

type Packer struct {
	Path string
	Data string
}

var rootCMD = &cobra.Command{
	Use:   "packr",
	Short: "packr - packer your static file into binary file",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ps := make([]*Packer, 0, 10)
		err := filepath.Walk(args[0], func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}

			data, err := encodeFile(path)
			if err != nil {
				return err
			}

			ps = append(ps, &Packer{
				Path: strings.Trim(strings.TrimPrefix(path, args[0]), "/"),
				Data: data,
			})

			return nil
		})

		if err != nil {
			return err
		}

		s := bytes.NewBufferString("")
		t := template.Must(template.New("listing").Parse(temp))

		packers := &Packers{
			Root: args[0],
			Ps:   ps,
		}
		err = t.Execute(s, packers)
		if err != nil {
			return err
		}

		err = ioutil.WriteFile("./packer_generate.go", s.Bytes(), 0644)
		if err != nil {
			return err
		}

		return nil
	},
}

const temp = `
package main

import "github.com/overbool/packr"

func init() {
	{{range .Ps}} 
	packer.PackData("{{ $.Root }}", "{{ .Path }}", "{{ .Data }}")
	{{end}}
}
`

func encodeFile(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}

	reader := bufio.NewReader(f)
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", err
	}

	// Encode as base64.
	encoded := base64.StdEncoding.EncodeToString(content)

	return encoded, nil
}
