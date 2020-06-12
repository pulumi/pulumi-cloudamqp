// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

package examples

import (
	"os"
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

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
