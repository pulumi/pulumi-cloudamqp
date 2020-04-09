// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

package examples

import (
	"os"
	"path"
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

func TestAccInstanceTs(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "instance", "ts"),
		})

	integration.ProgramTest(t, &test)
}

func TestAccInstancePython(t *testing.T) {
	test := getPythonBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "instance", "python"),
		})

	integration.ProgramTest(t, &test)
}

func TestAccInstanceCsharp(t *testing.T) {
	test := getCsharpBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "instance", "csharp"),
		})

	integration.ProgramTest(t, &test)
}

func getApiKey(t *testing.T) string {
	name := os.Getenv("CLOUDAMQP_APIKEY")
	if name == "" {
		t.Skipf("Skipping test due to missing CLOUDAMQP_APIKEY environment variable")
	}

	return name
}

func getCwd(t *testing.T) string {
	cwd, err := os.Getwd()
	if err != nil {
		t.FailNow()
	}

	return cwd
}

func getBaseOptions() integration.ProgramTestOptions {
	return integration.ProgramTestOptions{
		Quick:                    true,
		ExpectRefreshChanges:     true,
		AllowEmptyPreviewChanges: true, //this is a temporary thing right now because we are getting weird behaviour!
	}
}

func getJSBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions()
	baseJS := base.With(integration.ProgramTestOptions{
		Config: map[string]string{
			"cloudamqp:apikey": getApiKey(t),
		},
		Dependencies: []string{
			"@pulumi/cloudamqp",
		},
	})

	return baseJS
}

func getPythonBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions()
	basePython := base.With(integration.ProgramTestOptions{
		Config: map[string]string{
			"cloudamqp:apikey": getApiKey(t),
		},
		Dependencies: []string{
			filepath.Join("..", "sdk", "python", "bin"),
		},
	})

	return basePython
}

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
