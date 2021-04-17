// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.
// +build dotnet all

package examples

import (
	"path"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func getCsharpBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions()
	baseCsharp := base.With(integration.ProgramTestOptions{
		Config: map[string]string{
			"cloudamqp:apikey": getApiKey(t),
		},
		Dependencies: []string{
			"Pulumi.CloudAmqp",
		},
	})

	return baseCsharp
}

func TestAccInstanceCsharp(t *testing.T) {
	test := getCsharpBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "instance", "csharp"),
		})

	integration.ProgramTest(t, &test)
}
