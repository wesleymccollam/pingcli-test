// Copyright © 2025 Ping Identity Corporation
// Code generated by ping-cli-generator

package pingone_mfa_testable_resources

import (
	"testing"

	"github.com/patrickcping/pingone-go-sdk-v2/mfa"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingone/mfa/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource/pingone_sso_testable_resources"
)

func MfaApplicationPushCredential(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo: clientInfo,
		CreateFunc: createMfaApplicationPushCredential,
		DeleteFunc: deleteMfaApplicationPushCredential,
		Dependencies: []*testutils_resource.TestableResource{
			pingone_sso_testable_resources.ApplicationNative(t, clientInfo),
		},
		ExportableResource: resources.MfaApplicationPushCredential(clientInfo),
	}
}

func createMfaApplicationPushCredential(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, strArgs ...string) testutils_resource.ResourceInfo {
	t.Helper()

	if len(strArgs) != 1 {
		t.Errorf("Unexpected number of arguments provided to createMfaApplicationPushCredential(): %v", strArgs)

		return testutils_resource.ResourceInfo{}
	}
	applicationId := strArgs[0]

	request := clientInfo.PingOneApiClient.MFAAPIClient.ApplicationsApplicationMFAPushCredentialsApi.CreateMFAPushCredential(clientInfo.PingOneContext, clientInfo.PingOneExportEnvironmentID, applicationId)
	clientStruct := mfa.MFAPushCredentialRequest{
		MFAPushCredentialHMS: &mfa.MFAPushCredentialHMS{
			Type:         mfa.ENUMMFAPUSHCREDENTIALATTRTYPE_HMS,
			ClientId:     "897389789432",
			ClientSecret: "B23897498",
		},
	}

	request = request.MFAPushCredentialRequest(clientStruct)

	resource, response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "CreateMFAPushCredential", resourceType)
	if err != nil {
		t.Errorf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)

		return testutils_resource.ResourceInfo{}
	}
	if !ok {
		t.Errorf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)

		return testutils_resource.ResourceInfo{}
	}

	return testutils_resource.ResourceInfo{
		DeletionIds: []string{
			applicationId,
			*resource.Id,
		},
		CreationInfo: map[testutils_resource.ResourceCreationInfoType]string{
			testutils_resource.ENUM_ID:   *resource.Id,
			testutils_resource.ENUM_TYPE: string(*resource.Type),
		},
	}
}

func deleteMfaApplicationPushCredential(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, ids ...string) {
	t.Helper()

	if len(ids) != 2 {
		t.Errorf("Unexpected number of arguments provided to deleteMfaApplicationPushCredential(): %v", ids)

		return
	}

	request := clientInfo.PingOneApiClient.MFAAPIClient.ApplicationsApplicationMFAPushCredentialsApi.DeleteMFAPushCredential(clientInfo.PingOneContext, clientInfo.PingOneExportEnvironmentID, ids[0], ids[1])

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteMFAPushCredential", resourceType)
	if err != nil {
		t.Errorf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)

		return
	}
	if !ok {
		t.Errorf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)

		return
	}
}
