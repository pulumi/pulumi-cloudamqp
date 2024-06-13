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
	"fmt"
	"net/http"
	"path"
	"unicode"

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

func makeToken(mod, name string) string {
	fulllMod := string(unicode.ToLower(rune(name[0]))) + name[1:]
	return fmt.Sprintf("%s:%s/%s:%s", mainPkg, mod, fulllMod, name)
}

// makeDataSource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the data source's
// first character.
func makeDataSource(mod string, dataSource string) tokens.ModuleMember {
	return tokens.ModuleMember(makeToken(mod, dataSource))
}

// makeResource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the resource's
// first character.
func makeResource(mod string, res string) tokens.Type {
	return tokens.Type(makeToken(mod, res))
}

func ref[T any](t T) *T { return &t }

func Provider() tfbridge.ProviderInfo {
	p := shimv2.NewProvider(cloudamqp.Provider("", http.DefaultClient))
	prov := tfbridge.ProviderInfo{
		P:                 p,
		Name:              "cloudamqp",
		GitHubOrg:         "cloudamqp",
		Description:       "A Pulumi package for creating and managing CloudAMQP resources.",
		Keywords:          []string{"pulumi", "cloudamqp"},
		License:           "Apache-2.0",
		TFProviderLicense: ref(tfbridge.MITLicenseType),
		Homepage:          "https://pulumi.io",
		Repository:        "https://github.com/pulumi/pulumi-cloudamqp",
		Config:            map[string]*tfbridge.SchemaInfo{},
		Resources: map[string]*tfbridge.ResourceInfo{
			"cloudamqp_alarm":             {Tok: makeResource(mainMod, "Alarm")},
			"cloudamqp_instance":          {Tok: makeResource(mainMod, "Instance")},
			"cloudamqp_notification":      {Tok: makeResource(mainMod, "Notification")},
			"cloudamqp_plugin":            {Tok: makeResource(mainMod, "Plugin")},
			"cloudamqp_plugin_community":  {Tok: makeResource(mainMod, "PluginCommunity")},
			"cloudamqp_security_firewall": {Tok: makeResource(mainMod, "SecurityFirewall")},
			"cloudamqp_vpc_peering":       {Tok: makeResource(mainMod, "VpcPeering")},
			"cloudamqp_integration_log":   {Tok: makeResource(mainMod, "IntegrationLog")},
			"cloudamqp_integration_metric": {
				Tok:  makeResource(mainMod, "IntegrationMetric"),
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"cloudamqp_webhook":                {Tok: makeResource(mainMod, "Webhook")},
			"cloudamqp_custom_domain":          {Tok: makeResource(mainMod, "CustomDomain")},
			"cloudamqp_vpc_gcp_peering":        {Tok: makeResource(mainMod, "VpcGcpPeering")},
			"cloudamqp_upgrade_rabbitmq":       {Tok: makeResource(mainMod, "UpgradeRabbitmq")},
			"cloudamqp_vpc":                    {Tok: makeResource(mainMod, "Vpc")},
			"cloudamqp_node_actions":           {Tok: makeResource(mainMod, "NodeActions")},
			"cloudamqp_rabbitmq_configuration": {Tok: makeResource(mainMod, "RabbitConfiguration")},
			"cloudamqp_extra_disk_size": {
				Tok: makeResource(mainMod, "ExtraDiskSize"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"extra_disk_size": {
						CSharpName: "ExtraDiskSizeGb",
					},
				},
			},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"cloudamqp_credentials":         {Tok: makeDataSource(mainMod, "getCredentials")},
			"cloudamqp_plugins":             {Tok: makeDataSource(mainMod, "getPlugins")},
			"cloudamqp_plugins_community":   {Tok: makeDataSource(mainMod, "getPluginsCommunity")},
			"cloudamqp_vpc_info":            {Tok: makeDataSource(mainMod, "getVpcInfo")},
			"cloudamqp_instance":            {Tok: makeDataSource(mainMod, "getInstance")},
			"cloudamqp_notification":        {Tok: makeDataSource(mainMod, "getNotification")},
			"cloudamqp_alarm":               {Tok: makeDataSource(mainMod, "getAlarm")},
			"cloudamqp_nodes":               {Tok: makeDataSource(mainMod, "getNodes")},
			"cloudamqp_account":             {Tok: makeDataSource(mainMod, "getAccount")},
			"cloudamqp_vpc_gcp_info":        {Tok: makeDataSource(mainMod, "getVpcGcpInfo")},
			"cloudamqp_account_vpcs":        {Tok: makeDataSource(mainMod, "getAccountVpcs")},
			"cloudamqp_upgradable_versions": {Tok: makeDataSource(mainMod, "getUpgradableVersions")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0",
				"@types/mime": "^2.0.0",
			},
			RespectSchemaVersion: true,
		},
		Python: &tfbridge.PythonInfo{
			RespectSchemaVersion: true,
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
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
