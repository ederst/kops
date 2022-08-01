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

package components

import (
	"strconv"
	"testing"

	kopsapi "k8s.io/kops/pkg/apis/kops"
	"k8s.io/kops/upup/pkg/fi"
	"k8s.io/kops/pkg/assets"
)

func buildKubeDNSCluster() *kopsapi.Cluster {
	return &kopsapi.Cluster{
		Spec: kopsapi.ClusterSpec{
			CloudProvider:     "aws",
			KubernetesVersion: "1.20.5",
			ServiceClusterIPRange: "10.135.128.0/17",
			KubeDNS: &kopsapi.KubeDNSConfig{
				NodeLocalDNS: &kopsapi.NodeLocalDNSConfig{},
			},
		},
	}
}

func Test_ForwardToKubeDNS_Is_False_When_NodeLocalDNS_Is_Enabled(t *testing.T) {
	c := buildKubeDNSCluster()
	c.Spec.KubeDNS.NodeLocalDNS.Enabled = fi.Bool(true)

	b := assets.NewAssetBuilder(c, "")

	ob := &KubeDnsOptionsBuilder{
		&OptionsContext{
			AssetBuilder: b,
		},
	}

	err := ob.BuildOptions(&c.Spec)
	if err != nil {
		t.Fatalf("Error while building KubeDNS options: %v", err)
	}

	if !fi.BoolValue(c.Spec.KubeDNS.NodeLocalDNS.Enabled) {
		t.Fatalf("NodeLocalDNS is not enabled.")
	}

	if fi.BoolValue(c.Spec.KubeDNS.NodeLocalDNS.ForwardToKubeDNS) {
		t.Fatalf("ForwardToKubeDNS is enabled.")
	}

}

func Test_ForwardToKubeDNS_Is_Nil_When_NodeLocalDNS_Is_Disabled(t *testing.T) {
	c := buildKubeDNSCluster()

	b := assets.NewAssetBuilder(c, "")

	ob := &KubeDnsOptionsBuilder{
		&OptionsContext{
			AssetBuilder: b,
		},
	}

	err := ob.BuildOptions(&c.Spec)
	if err != nil {
		t.Fatalf("Error while building KubeDNS options: %v", err)
	}

	if fi.BoolValue(c.Spec.KubeDNS.NodeLocalDNS.Enabled) {
		t.Fatalf("NodeLocalDNS is enabled.")
	}

	if c.Spec.KubeDNS.NodeLocalDNS.ForwardToKubeDNS != nil {
		t.Fatalf("ForwardToKubeDNS is set to %s", strconv.FormatBool(fi.BoolValue(c.Spec.KubeDNS.NodeLocalDNS.ForwardToKubeDNS)))
	}
}
