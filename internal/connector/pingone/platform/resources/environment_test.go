// Copyright © 2025 Ping Identity Corporation
// Code generated by ping-cli-generator

package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource/pingone_platform_testable_resources"
)

func Test_Environment(t *testing.T) {
	clientInfo := testutils.GetClientInfo(t)

	tr := pingone_platform_testable_resources.Environment(t, clientInfo)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: tr.ExportableResource.ResourceType(),
			ResourceName: tr.ExportableResource.ResourceType(),
			ResourceID:   clientInfo.PingOneExportEnvironmentID,
		},
	}

	testutils.ValidateImportBlocks(t, tr.ExportableResource, &expectedImportBlocks)
}
