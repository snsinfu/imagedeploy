package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/docopt/docopt-go"
	"gopkg.in/yaml.v3"

	"github.com/snsinfu/imagedeploy"
)

const usage = `
Build container image and deploy.

Usage: imagedeploy [-h] [<config>]

  <config>  Config file describing what images to build and how to
            deploy them. Defaults to Deploy.yaml if not specified.

Options:
  -h, --help  Show this message and exit
`

const defaultConfig = "Deploy.yaml"

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}

func run() error {
	opts, err := docopt.ParseDoc(usage)
	if err != nil {
		return err
	}

	filename, ok := opts["<config>"].(string)
	if !ok {
		filename = defaultConfig
	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	var config imagedeploy.Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return err
	}

	return imagedeploy.Run(config)
}
