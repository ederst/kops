/*
Copyright 2021 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"io"

	"github.com/spf13/cobra"
	"k8s.io/kops/cmd/kops/util"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	unsetLong = templates.LongDesc(i18n.T(`Unset a configuration field.

        kops unset does not update the cloud resources; to apply the changes use "kops update cluster".
    `))

	unsetExample = templates.Examples(i18n.T(`
	kops unset cluster k8s-cluster.example.com spec.iam.allowContainerRegistry
	kops unset instancegroup --name k8s-cluster.example.com nodes-1a spec.maxSize
	`))
)

func NewCmdUnset(f *util.Factory, out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "unset",
		Short:   i18n.T("Unset fields on clusters and other resources."),
		Long:    unsetLong,
		Example: unsetExample,
	}

	// create subcommands
	cmd.AddCommand(NewCmdUnsetCluster(f, out))
	cmd.AddCommand(NewCmdUnsetInstancegroup(f, out))

	return cmd
}
