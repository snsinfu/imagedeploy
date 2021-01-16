package imagedeploy

import (
	"github.com/snsinfu/imagedeploy/docker"
	"github.com/snsinfu/imagedeploy/kubectl"
)

func Run(config Config) error {
	for _, image := range config.Images {
		if err := updateImage(image); err != nil {
			return err
		}
	}

	if err := updateKubernetes(config.Kubernetes); err != nil {
		return err
	}

	return nil
}

func updateImage(image ImageConfig) error {
	build := docker.Build{
		Name:    image.Name,
		Context: image.Build.Context,
		Args:    image.Build.Args,
	}

	trace("")
	trace("Building image...")
	trace(&build)

	if err := build.Run(); err != nil {
		return err
	}

	defer (func() {
		rmi := docker.ImageRemove{Name: image.Name}

		trace("")
		trace("Removing image...")
		trace(&rmi)

		if err := rmi.Run(); err != nil {
			// Non-critical. Just log.
			trace(err)
			return
		}
	})()

	push := docker.Push{Name: image.Name}

	trace("")
	trace("Uploading image...")
	trace(&push)

	if err := push.Run(); err != nil {
		return err
	}

	return nil
}

func updateKubernetes(kube KubernetesConfig) error {
	if len(kube.Apply) > 0 {
		apply := kubectl.Apply{
			Manifests: kube.Apply,
			Namespace: kube.Namespace,
		}

		trace("")
		trace("Applying manifests...")
		trace(&apply)

		if err := apply.Run(); err != nil {
			return err
		}
	}

	if len(kube.Restart) > 0 {
		restart := kubectl.RolloutRestart{
			Resources: kube.Restart,
			Namespace: kube.Namespace,
		}

		trace("")
		trace("Restarting resources...")
		trace(&restart)

		if err := restart.Run(); err != nil {
			return err
		}
	}

	return nil
}
