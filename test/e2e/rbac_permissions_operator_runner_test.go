// THIS FILE IS GENERATED BY BOILERPLATE. DO NOT EDIT.
//go:build osde2e

package osde2etests

import (
	"os"
	"path/filepath"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const (
	testResultsDirectory = "/test-run-results"
	jUnitOutputFilename  = "junit-rbac-permissions-operator.xml"
)

// Test entrypoint. osde2e runs this as a test suite on test pod.
func TestRbacPermissionsOperator(t *testing.T) {
	RegisterFailHandler(Fail)
	suiteConfig, reporterConfig := GinkgoConfiguration()
	if _, ok := os.LookupEnv("DISABLE_JUNIT_REPORT"); !ok {
		reporterConfig.JUnitReport = filepath.Join(testResultsDirectory, jUnitOutputFilename)
	}
	RunSpecs(t, "Rbac Permissions Operator", suiteConfig, reporterConfig)
}
