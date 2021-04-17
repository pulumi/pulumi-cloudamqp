// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.
// +build python all

package examples

import (
	"path"
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

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

func TestAccInstancePython(t *testing.T) {
	test := getPythonBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "instance", "python"),
		})

	integration.ProgramTest(t, &test)
}
