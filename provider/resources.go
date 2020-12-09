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
	"unicode"

	"github.com/cloudamqp/terraform-provider-cloudamqp/cloudamqp"
	"github.com/pulumi/pulumi-terraform-bridge/v2/pkg/tfbridge"
	shimv1 "github.com/pulumi/pulumi-terraform-bridge/v2/pkg/tfshim/sdk-v1"
	"github.com/pulumi/pulumi/sdk/v2/go/common/tokens"
)

// all of the token components used below.
const (
	// packages:
	mainPkg = "cloudamqp"
	// modules:
	mainMod = "index" // the y module
)

// makeMember manufactures a type token for the package and the given module and type.
func makeMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(mainPkg + ":" + mod + ":" + mem)
}

// makeType manufactures a type token for the package and the given module and type.
func makeType(mod string, typ string) tokens.Type {
	return tokens.Type(makeMember(mod, typ))
}

// makeDataSource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the data source's
// first character.
func makeDataSource(mod string, res string) tokens.ModuleMember {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeMember(mod+"/"+fn, res)
}

// makeResource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the resource's
// first character.
func makeResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeType(mod+"/"+fn, res)
}

func refProviderLicense(license tfbridge.TFProviderLicense) *tfbridge.TFProviderLicense {
	return &license
}

func Provider() tfbridge.ProviderInfo {
	p := shimv1.NewProvider(cloudamqp.Provider())
	prov := tfbridge.ProviderInfo{
		P:                 p,
		Name:              "cloudamqp",
		GitHubOrg:         "cloudamqp",
		Description:       "A Pulumi package for creating and managing CloudAMQP resources.",
		Keywords:          []string{"pulumi", "cloudamqp"},
		License:           "Apache-2.0",
		TFProviderLicense: refProviderLicense(tfbridge.MITLicenseType),
		Homepage:          "https://pulumi.io",
		Repository:        "https://github.com/pulumi/pulumi-cloudamqp",
		Config: map[string]*tfbridge.SchemaInfo{
			"apikey": {
				Type: "string",
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"CLOUDAMQP_APIKEY"},
				},
			},
		},
		Resources: map[string]*tfbridge.ResourceInfo{
			"cloudamqp_alarm":              {Tok: makeResource(mainMod, "Alarm")},
			"cloudamqp_instance":           {Tok: makeResource(mainMod, "Instance")},
			"cloudamqp_notification":       {Tok: makeResource(mainMod, "Notification")},
			"cloudamqp_plugin":             {Tok: makeResource(mainMod, "Plugin")},
			"cloudamqp_plugin_community":   {Tok: makeResource(mainMod, "PluginCommunity")},
			"cloudamqp_security_firewall":  {Tok: makeResource(mainMod, "SecurityFirewall")},
			"cloudamqp_vpc_peering":        {Tok: makeResource(mainMod, "VpcPeering")},
			"cloudamqp_integration_log":    {Tok: makeResource(mainMod, "IntegrationLog")},
			"cloudamqp_integration_metric": {Tok: makeResource(mainMod, "IntegrationMetric")},
			"cloudamqp_webhook":            {Tok: makeResource(mainMod, "Webhook")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"cloudamqp_credentials":       {Tok: makeDataSource(mainMod, "getCredentials")},
			"cloudamqp_plugins":           {Tok: makeDataSource(mainMod, "getPlugins")},
			"cloudamqp_plugins_community": {Tok: makeDataSource(mainMod, "getPluginsCommunity")},
			"cloudamqp_vpc_info":          {Tok: makeDataSource(mainMod, "getVpcInfo")},
			"cloudamqp_instance":          {Tok: makeDataSource(mainMod, "getInstance")},
			"cloudamqp_notification":      {Tok: makeDataSource(mainMod, "getNotification")},
			"cloudamqp_alarm":             {Tok: makeDataSource(mainMod, "getAlarm")},
			"cloudamqp_nodes":             {Tok: makeDataSource(mainMod, "getNodes")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^2.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^8.0.25",
				"@types/mime": "^2.0.0",
			},
		},
		Python: &tfbridge.PythonInfo{
			Requires: map[string]string{
				"pulumi": ">=2.9.0,<3.0.0",
			},
			UsesIOClasses: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi":                       "2.*",
				"System.Collections.Immutable": "1.6.0",
			},
			Namespaces: map[string]string{
				mainPkg: "CloudAmqp",
			},
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
