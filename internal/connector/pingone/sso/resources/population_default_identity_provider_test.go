// Copyright © 2025 Ping Identity Corporation
// Code generated by ping-cli-generator

package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource/pingone_sso_testable_resources"
)

func Test_PopulationDefaultIdentityProvider(t *testing.T) {
	clientInfo := testutils.GetClientInfo(t)

	tr := pingone_sso_testable_resources.PopulationDefaultIdentityProvider(t, clientInfo)

	tr.CreateResource(t)
	defer tr.DeleteResource(t)

	populationTr := tr.Dependencies[0]

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: tr.ExportableResource.ResourceType(),
			ResourceName: fmt.Sprintf("%s_default_identity_provider", populationTr.ResourceInfo.CreationInfo[testutils_resource.ENUM_NAME]),
			ResourceID:   fmt.Sprintf("%s/%s", clientInfo.PingOneExportEnvironmentID, populationTr.ResourceInfo.CreationInfo[testutils_resource.ENUM_ID]),
		},
	}

	// Existing populations are generated. Test subset.
	testutils.ValidateImportBlockSubset(t, tr.ExportableResource, &expectedImportBlocks)
}
