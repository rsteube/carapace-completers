package action

import (
	"os/exec"
	"strings"

	"github.com/rsteube/carapace"
)

func ActionApiResources() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		if output, err := exec.Command("kubectl", "api-resources", "--output=name", "--cached").Output(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			lines := strings.Split(string(output), "\n")
			for index, line := range lines {
				lines[index] = strings.SplitN(line, ".", 2)[0]
			}
			return carapace.ActionValues(lines[:len(lines)-1]...)
		}
	})
}

func ActionResources(namespace string, types string) carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		if output, err := exec.Command("kubectl", "--namespace", namespace, "get", "-o", "go-template={{range .items}}{{.metadata.name}}\n{{.kind}}\n{{end}}", types).Output(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			lines := strings.Split(string(output), "\n")
			return carapace.ActionValuesDescribed(lines[:len(lines)-1]...)
		}
	})
}

func ActionDryRunModes() carapace.Action {
	return carapace.ActionValues("none", "server", "client")
}

func ActionOutputFormats() carapace.Action {
	return carapace.ActionValues("json", "yaml", "name", "go-template", "go-template-file", "template", "templatefile", "jsonpath", "jsonpath-as-json", "jsonpath-file")
}

func ActionResourceVerbs() carapace.Action {
	return carapace.ActionValues("get", "list", "create", "update", "patch", "watch", "delete", "deletecollection")
}

func ActionNamespaceServiceAccounts() carapace.Action {
	return carapace.ActionMultiParts(":", func(args, parts []string) carapace.Action {
		switch len(parts) {
		case 0:
			return ActionResources("", "namespaces").Invoke(args).Suffix(":").ToA()
		case 1:
			return ActionResources(parts[0], "serviceaccounts")
		default:
			return carapace.ActionValues()
		}
	})
}

func ActionApiResourceResources() carapace.Action {
	return carapace.ActionMultiParts("/", func(args, parts []string) carapace.Action {
		switch len(parts) {
		case 0:
			return ActionApiResources().Invoke(args).Suffix("/").ToA()
		case 1:
			return ActionResources("", parts[0])
		default:
			return carapace.ActionValues()
		}
	})
}

func ActionContainers(namespace string, resource string) carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		if output, err := exec.Command("kubectl", "--namespace", namespace, "get", "-o", "go-template={{range .spec.containers}}{{.name}}\n{{end}}", resource).Output(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			lines := strings.Split(string(output), "\n")
			return carapace.ActionValues(lines[:len(lines)-1]...)
		}
	})
}
