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

package openstacktasks

import (
	"io/ioutil"
	"os"
	"testing"

	"k8s.io/kops/upup/pkg/fi/cloudup/openstack"
	"k8s.io/kops/upup/pkg/fi/cloudup/terraform"
)

type renderTest struct {
	Resource interface{}
	Expected string
}

func doRenderTests(t *testing.T, method string, cases []*renderTest) {
	outdir, err := ioutil.TempDir("", "kops-render-")
	if err != nil {
		t.Errorf("failed to create local render directory: %s", err)
		t.FailNow()
	}
	defer func () {
		err := os.RemoveAll(outdir)
		if err != nil {
			t.Errorf("failed to remove temp dir %q: %v", outdir, err)
		}
	}()

	for i, c := range cases {
		var filename string
		var target interface{}

		cloud := openstack.BuildMockOpenstackCloud("nova")

		switch method {
		case "RenderTerraform":
			target = terraform.NewTerraformTarget(cloud, "test", outdir, nil)
			filename = "kubernetes.tf"
		default:
			t.Errorf("unknown render method: %s", method)
			t.FailNow()
		}

		t.Logf("do something with %s and %s", filename, target)

		t.Logf("case %d, expected: %s", i, c.Expected)
	}
}