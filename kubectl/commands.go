// Package kubectl wraps kubectl cli commands.
package kubectl

import (
	"github.com/snsinfu/imagedeploy/command"
)

// Apply holds arguments for `kubectl apply`.
type Apply struct {
	Manifests []string
	Kustomize string
	Namespace string
}

func (opts *Apply) Run() error {
	return opts.command().Run()
}

func (opts *Apply) String() string {
	return opts.command().String()
}

func (opts *Apply) command() command.Line {
	cmd := command.Line{"kubectl", "apply"}

	for _, manifest := range opts.Manifests {
		cmd = append(cmd, "-f", manifest)
	}
	if opts.Kustomize != "" {
		cmd = append(cmd, "-k", opts.Kustomize)
	}
	if opts.Namespace != "" {
		cmd = append(cmd, "-n", opts.Namespace)
	}

	return cmd
}

// Apply holds arguments for `kubectl rollout restart`.
type RolloutRestart struct {
	Resources []string
	Namespace string
}

func (opts *RolloutRestart) Run() error {
	return opts.command().Run()
}

func (opts *RolloutRestart) String() string {
	return opts.command().String()
}

func (opts *RolloutRestart) command() command.Line {
	cmd := command.Line{"kubectl", "rollout", "restart"}

	if opts.Namespace != "" {
		cmd = append(cmd, "-n", opts.Namespace)
	}
	cmd = append(cmd, opts.Resources...)

	return cmd
}
