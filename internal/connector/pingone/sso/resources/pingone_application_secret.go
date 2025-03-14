package resources

import (
	"fmt"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingone"
	"github.com/pingidentity/pingcli/internal/logger"
	"github.com/pingidentity/pingcli/internal/output"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneApplicationSecretResource{}
)

type PingOneApplicationSecretResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingOneApplicationSecretResource
func ApplicationSecret(clientInfo *connector.PingOneClientInfo) *PingOneApplicationSecretResource {
	return &PingOneApplicationSecretResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneApplicationSecretResource) ResourceType() string {
	return "pingone_application_secret"
}

func (r *PingOneApplicationSecretResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	applicationData, err := r.getApplicationData()
	if err != nil {
		return nil, err
	}

	for appId, appName := range applicationData {
		ok, err := r.checkApplicationSecretData(appId)
		if err != nil {
			return nil, err
		}

		if !ok {
			continue
		}

		commentData := map[string]string{
			"Application ID":        appId,
			"Application Name":      appName,
			"Export Environment ID": r.clientInfo.ExportEnvironmentID,
			"Resource Type":         r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       fmt.Sprintf("%s_secret", appName),
			ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, appId),
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingOneApplicationSecretResource) getApplicationData() (map[string]string, error) {
	applicationData := make(map[string]string)

	iter := r.clientInfo.ApiClient.ManagementAPIClient.ApplicationsApi.ReadAllApplications(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()
	applications, err := pingone.GetManagementAPIObjectsFromIterator[management.ReadOneApplication200Response](iter, "ReadAllApplications", "GetApplications", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, app := range applications {
		var (
			appId     *string
			appIdOk   bool
			appName   *string
			appNameOk bool
		)

		switch {
		case app.ApplicationOIDC != nil:
			appId, appIdOk = app.ApplicationOIDC.GetIdOk()
			appName, appNameOk = app.ApplicationOIDC.GetNameOk()
		case app.ApplicationSAML != nil:
			appId, appIdOk = app.ApplicationSAML.GetIdOk()
			appName, appNameOk = app.ApplicationSAML.GetNameOk()
		case app.ApplicationExternalLink != nil:
			appId, appIdOk = app.ApplicationExternalLink.GetIdOk()
			appName, appNameOk = app.ApplicationExternalLink.GetNameOk()
		default:
			continue
		}

		if appIdOk && appNameOk {
			applicationData[*appId] = *appName
		}
	}

	return applicationData, nil
}

func (r *PingOneApplicationSecretResource) checkApplicationSecretData(appId string) (bool, error) {
	// The platform enforces that worker apps cannot read their own secret
	// Make sure we can read the secret before adding it to the import blocks
	_, response, err := r.clientInfo.ApiClient.ManagementAPIClient.ApplicationSecretApi.ReadApplicationSecret(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, appId).Execute()
	defer response.Body.Close()

	// Use output package to warn the user of any errors or non-200 response codes
	// Expected behavior in this case is to skip the resource, and continue exporting the other resources
	if err != nil || response.StatusCode >= 300 || response.StatusCode < 200 {
		output.Warn("Failed to read secret for application", map[string]interface{}{
			"Application ID":    appId,
			"API Function Name": "ReadApplicationSecret",
			"Response Code":     response.Status,
			"Response Body":     response.Body,
		})
		return false, nil
	}

	return true, nil

}
