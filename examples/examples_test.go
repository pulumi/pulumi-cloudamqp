// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

package examples

import (
	"github.com/pulumi/pulumi/pkg/testing/integration"
	"os"
	"path"
	"testing"
)

func TestAccInstance(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "instance"),
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
	return integration.ProgramTestOptions{}
}

func getJSBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions()
	baseJS := base.With(integration.ProgramTestOptions{
		ExpectRefreshChanges:     true,
		AllowEmptyPreviewChanges: true, //this is a temporary thing right now because we are getting weird behaviour!
		Config: map[string]string{
			"cloudamqp:apikey": getApiKey(t),
		},
		Dependencies: []string{
			"@pulumi/cloudamqp",
		},
	})

	return baseJS
}
