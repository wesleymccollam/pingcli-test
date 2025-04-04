// Copyright © 2025 Ping Identity Corporation

package resources

import (
	"fmt"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingone"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneLanguageUpdateResource{}
)

type PingOneLanguageUpdateResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingOneLanguageUpdateResource
func LanguageUpdate(clientInfo *connector.ClientInfo) *PingOneLanguageUpdateResource {
	return &PingOneLanguageUpdateResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneLanguageUpdateResource) ResourceType() string {
	return "pingone_language_update"
}

func (r *PingOneLanguageUpdateResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	languageUpdateData, err := r.getLanguageUpdateData()
	if err != nil {
		return nil, err
	}

	for languageId, languageName := range languageUpdateData {
		commentData := map[string]string{
			"Export Environment ID": r.clientInfo.PingOneExportEnvironmentID,
			"Language ID":           languageId,
			"Language Name":         languageName,
			"Resource Type":         r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       fmt.Sprintf("%s_update", languageName),
			ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.PingOneExportEnvironmentID, languageId),
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingOneLanguageUpdateResource) getLanguageUpdateData() (map[string]string, error) {
	languageUpdateData := make(map[string]string)

	iter := r.clientInfo.PingOneApiClient.ManagementAPIClient.LanguagesApi.ReadLanguages(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID).Execute()
	apiObjs, err := pingone.GetManagementAPIObjectsFromIterator[management.EntityArrayEmbeddedLanguagesInner](iter, "ReadLanguages", "GetLanguages", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, innerObj := range apiObjs {
		if innerObj.Language != nil {
			languageEnabled, languageEnabledOk := innerObj.Language.GetEnabledOk()
			languageLocale, languageLocaleOk := innerObj.Language.GetLocaleOk()
			languageDefault, languageDefaultOk := innerObj.Language.GetDefaultOk()

			if languageEnabledOk && languageLocaleOk && languageDefaultOk {
				// Export the language if it meets any of the criteria of the following 3 conditions:
				// 1) Any language enabled
				// 2) The 'en' language disabled
				// 3) If any language other than 'en' is the default

				if *languageEnabled || (*languageLocale == "en" && !*languageEnabled) || (*languageLocale != "en" && *languageDefault) {
					languageId, languageIdOk := innerObj.Language.GetIdOk()
					languageName, languageNameOk := innerObj.Language.GetNameOk()

					if languageIdOk && languageNameOk {
						languageUpdateData[*languageId] = *languageName
					}
				}
			}
		}
	}

	return languageUpdateData, nil
}
