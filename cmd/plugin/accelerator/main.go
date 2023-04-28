package main

import (
	"os"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/log"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/plugin"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/plugin/buildinfo"
)

var descriptor = plugin.PluginDescriptor{
	Name:        "accelerator",
	Description: "manage accelerators in a kubernetes cluster",
	Target:      types.TargetK8s,
	Version:     buildinfo.Version,
	BuildSHA:    buildinfo.SHA,
	Group:       plugin.RunCmdGroup,
	Aliases:     []string{"acc"},
}

func main() {
	p, err := plugin.NewPlugin(&descriptor)
	if err != nil {
		log.Fatal(err, "")
	}
	p.AddCommands(
	// Add commands
	)
	if err := p.Execute(); err != nil {
		os.Exit(1)
	}
}
