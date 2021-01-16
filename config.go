package imagedeploy

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

// Config describes what images to build and how to deploy them.
type Config struct {
	Images     []ImageConfig
	Kubernetes KubernetesConfig
}

// ImageConfig describes a container image.
type ImageConfig struct {
	Name  string
	Build BuildConfig
}

// BuildConfig describes how to build a container image.
type BuildConfig struct {
	Context string
	Args    map[string]string
}

func (bc *BuildConfig) UnmarshalYAML(node *yaml.Node) error {
	value := resolveAlias(node)

	if value.Kind == yaml.ScalarNode {
		var str string
		if err := value.Decode(&str); err != nil {
			return err
		}
		*bc = BuildConfig{Context: str}
		return nil
	}

	if value.Kind == yaml.MappingNode {
		var data struct {
			Context string
			Args    map[string]string
		}
		if err := value.Decode(&data); err != nil {
			return err
		}
		*bc = BuildConfig{Context: data.Context, Args: data.Args}
		return nil
	}

	return fmt.Errorf("unmarshal error:\n"+
		"  line %d: expected string or map", node.Line)
}

// KubernetesConfig describes app deployment on a Kubernetes cluster.
type KubernetesConfig struct {
	Namespace string
	Apply     anyString
	Restart   anyString
}

// anyString is array of strings that is unmarshable from a sequence of strings
// or a scalar string in YAML document. A scalar string is read as an array of
// single element.
type anyString []string

func (ss *anyString) UnmarshalYAML(node *yaml.Node) error {
	value := resolveAlias(node)

	if value.Kind == yaml.ScalarNode {
		var str string
		if err := value.Decode(&str); err != nil {
			return err
		}
		*ss = anyString{str}
		return nil
	}

	if value.Kind == yaml.SequenceNode {
		var seq []string
		if err := value.Decode(&seq); err != nil {
			return err
		}
		*ss = seq
		return nil
	}

	return fmt.Errorf("unmarshal error:\n"+
		"  line %d: expected string or sequence of strings", node.Line)
}

// resolveAlias resolves alias node to the node pointed to.
func resolveAlias(node *yaml.Node) *yaml.Node {
	for node.Kind == yaml.AliasNode {
		node = node.Alias
	}
	return node
}
