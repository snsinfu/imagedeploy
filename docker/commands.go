// Package kubectl wraps kubectl cli commands.
package docker

import (
	"github.com/snsinfu/imagedeploy/command"
)

// Build holds arguments for `docker build`.
type Build struct {
	Context string
	Name    string
	Args    map[string]string
}

func (opts *Build) Run() error {
	return opts.command().Run()
}

func (opts *Build) String() string {
	return opts.command().String()
}

func (opts *Build) command() command.Line {
	cmd := command.Line{"docker", "build"}

	if opts.Name != "" {
		cmd = append(cmd, "-t", opts.Name)
	}
	for key, value := range opts.Args {
		cmd = append(cmd, "--build-arg", key+"="+value)
	}
	cmd = append(cmd, opts.Context)

	return cmd
}

// Build holds arguments for `docker push`.
type Push struct {
	Name string
}

func (opts *Push) Run() error {
	return opts.command().Run()
}

func (opts *Push) String() string {
	return opts.command().String()
}

func (opts *Push) command() command.Line {
	return command.Line{"docker", "push", opts.Name}
}

// Build holds arguments for `docker image rm` (or `docker rmi`).
type ImageRemove struct {
	Name string
}

func (opts *ImageRemove) Run() error {
	return opts.command().Run()
}

func (opts *ImageRemove) String() string {
	return opts.command().String()
}

func (opts *ImageRemove) command() command.Line {
	return command.Line{"docker", "image", "rm", opts.Name}
}
