// Copyright © 2025 Ping Identity Corporation
// Code generated by ping-cli-generator

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
	_ connector.ExportableResource = &PingOneResourceScopeResource{}
)

type PingOneResourceScopeResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingOneResourceScopeResource
func ResourceScope(clientInfo *connector.ClientInfo) *PingOneResourceScopeResource {
	return &PingOneResourceScopeResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneResourceScopeResource) ResourceType() string {
	return "pingone_resource_scope"
}

func (r *PingOneResourceScopeResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	resourceData, err := r.getResourceData()
	if err != nil {
		return nil, err
	}

	for resourceId, resourceName := range resourceData {
		resourceScopeData, err := r.getResourceScopeData(resourceId)
		if err != nil {
			return nil, err
		}

		for resourceScopeId, resourceScopeName := range resourceScopeData {
			commentData := map[string]string{
				"Resource ID":           resourceId,
				"Resource Name":         resourceName,
				"Resource Scope ID":     resourceScopeId,
				"Resource Scope Name":   resourceScopeName,
				"Export Environment ID": r.clientInfo.PingOneExportEnvironmentID,
				"Resource Type":         r.ResourceType(),
			}

			importBlock := connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       fmt.Sprintf("%s_%s", resourceName, resourceScopeName),
				ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.PingOneExportEnvironmentID, resourceId, resourceScopeId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			}

			importBlocks = append(importBlocks, importBlock)
		}
	}

	return &importBlocks, nil
}

func (r *PingOneResourceScopeResource) getResourceData() (map[string]string, error) {
	resourceData := make(map[string]string)

	iter := r.clientInfo.PingOneApiClient.ManagementAPIClient.ResourcesApi.ReadAllResources(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID).Execute()
	apiObjs, err := pingone.GetManagementAPIObjectsFromIterator[management.EntityArrayEmbeddedResourcesInner](iter, "ReadAllResources", "GetResources", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, innerObj := range apiObjs {
		if innerObj.Resource != nil {
			resourceId, resourceIdOk := innerObj.Resource.GetIdOk()
			resourceName, resourceNameOk := innerObj.Resource.GetNameOk()
			resourceType, resourceTypeOk := innerObj.Resource.GetTypeOk()

			if resourceIdOk && resourceNameOk && resourceTypeOk && *resourceType == management.ENUMRESOURCETYPE_CUSTOM {
				resourceData[*resourceId] = *resourceName
			}
		}
	}

	return resourceData, nil
}

func (r *PingOneResourceScopeResource) getResourceScopeData(resourceId string) (map[string]string, error) {
	resourceScopeData := make(map[string]string)

	iter := r.clientInfo.PingOneApiClient.ManagementAPIClient.ResourceScopesApi.ReadAllResourceScopes(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID, resourceId).Execute()
	apiObjs, err := pingone.GetManagementAPIObjectsFromIterator[management.ResourceScope](iter, "ReadAllResourceScopes", "GetScopes", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, resourceScope := range apiObjs {
		resourceScopeId, resourceScopeIdOk := resourceScope.GetIdOk()
		resourceScopeName, resourceScopeNameOk := resourceScope.GetNameOk()

		if resourceScopeIdOk && resourceScopeNameOk {
			resourceScopeData[*resourceScopeId] = *resourceScopeName
		}
	}

	return resourceScopeData, nil
}
