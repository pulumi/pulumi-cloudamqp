// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cloudamqp

import (
	"bytes"
	"fmt"
	"net/http"
	"path"

	"github.com/cloudamqp/terraform-provider-cloudamqp/cloudamqp"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	tfbridgetokens "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"

	"github.com/pulumi/pulumi-cloudamqp/provider/v3/pkg/version"
)

// all of the token components used below.
const (
	// packages:
	mainPkg = "cloudamqp"
	// modules:
	mainMod = "index" // the y module
)

func makeResource(res string) tokens.Type { return tfbridge.MakeResource(mainPkg, mainMod, res) }

func ref[T any](t T) *T { return &t }

func Provider() tfbridge.ProviderInfo {
	prov := tfbridge.ProviderInfo{
		P:                 shimv2.NewProvider(cloudamqp.Provider("", http.DefaultClient)),
		Name:              "cloudamqp",
		DisplayName:       "CloudAMQP",
		GitHubOrg:         "cloudamqp",
		Description:       "A Pulumi package for creating and managing CloudAMQP resources.",
		Keywords:          []string{"pulumi", "cloudamqp"},
		License:           "Apache-2.0",
		TFProviderLicense: ref(tfbridge.MITLicenseType),
		Homepage:          "https://pulumi.io",
		Repository:        "https://github.com/pulumi/pulumi-cloudamqp",
		Config:            map[string]*tfbridge.SchemaInfo{},
		DocRules:          &tfbridge.DocRuleInfo{EditRules: docEditRules},
		// Resources: map[string]*tfbridge.ResourceInfo{
		// 	"cloudamqp_rabbitmq_configuration": {Tok: makeResource("RabbitConfiguration")},
		// 	"cloudamqp_integration_metric": {
		// 		Docs: &tfbridge.DocInfo{AllowMissing: true},
		// 	},
		// 	"cloudamqp_extra_disk_size": {
		// 		Fields: map[string]*tfbridge.SchemaInfo{
		// 			"extra_disk_size": {
		// 				CSharpName: "ExtraDiskSizeGb",
		// 			},
		// 		},
		// 	},
		// },
		JavaScript: &tfbridge.JavaScriptInfo{
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0",
				"@types/mime": "^2.0.0",
			},
			RespectSchemaVersion: true,
		},
		Python: &tfbridge.PythonInfo{
			RespectSchemaVersion: true,

			PyProject: struct{ Enabled bool }{true},
		},

		Golang: &tfbridge.GolangInfo{
			ImportBasePath: path.Join(
				fmt.Sprintf("github.com/pulumi/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
			RespectSchemaVersion:           true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RespectSchemaVersion: true,
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
			Namespaces: map[string]string{
				mainPkg: "CloudAmqp",
			},
		},
	}

	prov.MustComputeTokens(tfbridgetokens.SingleModule("cloudamqp_", mainMod,
		tfbridgetokens.MakeStandard(mainPkg)))

	prov.SetAutonaming(255, "-")

	return prov
}

func docEditRules(defaults []tfbridge.DocsEdit) []tfbridge.DocsEdit {
	return append(
		defaults,
		removeBlockReference,
	)
}

// Removes a block reference
var removeBlockReference = tfbridge.DocsEdit{
	Path: "index.md",
	Edit: func(_ string, content []byte) ([]byte, error) {
		content = bytes.ReplaceAll(content, []byte(" in the `provider` block"), nil)
		return content, nil
	},
}
