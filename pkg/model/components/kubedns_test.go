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
	// "k8s.io/kops/upup/pkg/fi"
	// "k8s.io/kops/pkg/apis/kops/util"
	"k8s.io/kops/pkg/assets"
)

func buildKubeDNSCluster() *kopsapi.Cluster {
	return &kopsapi.Cluster{
		Spec: kopsapi.ClusterSpec{
			CloudProvider:     "aws",
			KubernetesVersion: "1.20.5",
			// NonMasqueradeCIDR: "100.64.0.0/10",
			ServiceClusterIPRange: "10.135.128.0/17",
			// ServiceClusterIPRange: &kopsapi.ServiceClusterIPRangeSpec{},
			// Networking: &kopsapi.NetworkingSpec{
			// 	Kubenet: &kopsapi.KubenetNetworkingSpec{},
			// },
			// KubeDNS: &kopsapi.KubeDNSConfig{
			// 	// NodeLocalDNS: &kopsapi.NodeLocalDNSConfig{},
			// },
			//    Provider: "CoreDNS",
			//    N
			// },
		},
	}
}

func Test_Forward_To_KubeDNS(t *testing.T) {
		c := buildKubeDNSCluster()
		// c.Spec.KubeDNS.enabled = true.ContainerRuntime = "containerd"
		// c.Spec.KubeDNS.NodeLocalDNS.Enabled = fi.Bool(true)
		// c.Spec.KubeDNS.NodeLocalDNS.Enabled = fi.Bool(true)
		b := assets.NewAssetBuilder(c, "")

		ob := &KubeDnsOptionsBuilder{
			&OptionsContext{
				AssetBuilder: b,
			},
		}

		err := ob.BuildOptions(&c.Spec)
		if err != nil {
			t.Fatalf("some meaningful error message here: %v", err)
		}

		if !*c.Spec.KubeDNS.NodeLocalDNS.Enabled && c.Spec.KubeDNS.NodeLocalDNS.ForwardToKubeDNS != nil {
		  t.Fatalf("Node local not enabled but forwardToKubeDNS is set to %s", strconv.FormatBool(*c.Spec.KubeDNS.NodeLocalDNS.ForwardToKubeDNS))
		}
		// t.Fatalf("Forward to KubeDNS?: %s", strconv.FormatBool(*c.Spec.KubeDNS.NodeLocalDNS.Enabled)) //NodeLocalDNS.ForwardToKubeDNS))

		// t.Fatalf("Node local not enabled: %s", strconv.FormatBool(*c.Spec.KubeDNS.NodeLocalDNS.Enabled))

		//t.Fatalf("Forward to KubeDNS?: %s", strconv.FormatBool(*c.Spec.KubeDNS.NodeLocalDNS.ForwardToKubeDNS))
		//t.Fatalf("Node local not enabled: %s", strconv.FormatBool(*c.Spec.KubeDNS.NodeLocalDNS.Enabled))
		// version, err := util.ParseKubernetesVersion(v)
		// if err != nil {
		// 	t.Fatalf("unexpected error from ParseKubernetesVersion %s: %v", v, err)
		// }

		// ob := &ContainerdOptionsBuilder{
		// 	&OptionsContext{
		// 		AssetBuilder:      b,
		// 		KubernetesVersion: *version,
		// 	},
		// }

		// err = ob.BuildOptions(&c.Spec)
		// if err != nil {
		// 	t.Fatalf("unexpected error from BuildOptions: %v", err)
		// }

		// if c.Spec.Containerd.SkipInstall == true {
		// 	t.Fatalf("expecting install when Kubernetes version >= 1.11: %s", v)
		// }
}

// func Test_Build_Containerd_Unneeded_Runtime(t *testing.T) {
// 	dockerVersions := []string{"1.13.1", "17.03.2", "18.06.3"}

// 	for _, v := range dockerVersions {

// 		c := buildContainerdCluster("1.11.0")
// 		c.Spec.ContainerRuntime = "docker"
// 		c.Spec.Docker = &kopsapi.DockerConfig{
// 			Version: &v,
// 		}
// 		b := assets.NewAssetBuilder(c, "")

// 		ob := &ContainerdOptionsBuilder{
// 			&OptionsContext{
// 				AssetBuilder: b,
// 			},
// 		}

// 		err := ob.BuildOptions(&c.Spec)
// 		if err != nil {
// 			t.Fatalf("unexpected error from BuildOptions: %v", err)
// 		}

// 		if c.Spec.Containerd.SkipInstall != true {
// 			t.Fatalf("unexpected install when Docker version < 19.09: %s", v)
// 		}
// 	}
// }

// func Test_Build_Containerd_Needed_Runtime(t *testing.T) {
// 	dockerVersions := []string{"18.09.3", "18.09.9", "19.03.4"}

// 	for _, v := range dockerVersions {

// 		c := buildContainerdCluster("1.11.0")
// 		c.Spec.ContainerRuntime = "docker"
// 		c.Spec.Docker = &kopsapi.DockerConfig{
// 			Version: &v,
// 		}
// 		b := assets.NewAssetBuilder(c, "")

// 		ob := &ContainerdOptionsBuilder{
// 			&OptionsContext{
// 				AssetBuilder: b,
// 			},
// 		}

// 		err := ob.BuildOptions(&c.Spec)
// 		if err != nil {
// 			t.Fatalf("unexpected error from BuildOptions: %v", err)
// 		}

// 		if c.Spec.Containerd.SkipInstall == true {
// 			t.Fatalf("expected install when Docker version >= 19.09: %s", v)
// 		}
// 	}
// }
