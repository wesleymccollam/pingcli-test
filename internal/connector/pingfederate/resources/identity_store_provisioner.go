// Copyright © 2025 Ping Identity Corporation
// Code generated by ping-cli-generator

package resources

import (
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateIdentityStoreProvisionerResource{}
)

type PingFederateIdentityStoreProvisionerResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingFederateIdentityStoreProvisionerResource
func IdentityStoreProvisioner(clientInfo *connector.ClientInfo) *PingFederateIdentityStoreProvisionerResource {
	return &PingFederateIdentityStoreProvisionerResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateIdentityStoreProvisionerResource) ResourceType() string {
	return "pingfederate_identity_store_provisioner"
}

func (r *PingFederateIdentityStoreProvisionerResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	identityStoreProvisionerData, err := r.getIdentityStoreProvisionerData()
	if err != nil {
		return nil, err
	}

	for identityStoreProvisionerId, identityStoreProvisionerName := range identityStoreProvisionerData {
		commentData := map[string]string{
			"Identity Store Provisioner ID":   identityStoreProvisionerId,
			"Identity Store Provisioner Name": identityStoreProvisionerName,
			"Resource Type":                   r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       identityStoreProvisionerName,
			ResourceID:         identityStoreProvisionerId,
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingFederateIdentityStoreProvisionerResource) getIdentityStoreProvisionerData() (map[string]string, error) {
	identityStoreProvisionerData := make(map[string]string)

	apiObj, response, err := r.clientInfo.PingFederateApiClient.IdentityStoreProvisionersAPI.GetIdentityStoreProvisioners(r.clientInfo.PingFederateContext).Execute()
	ok, err := common.HandleClientResponse(response, err, "GetIdentityStoreProvisioners", r.ResourceType())
	if err != nil {
		return nil, err
	}
	if !ok {
		return identityStoreProvisionerData, nil
	}

	if apiObj == nil {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	items, itemsOk := apiObj.GetItemsOk()
	if !itemsOk {
		return nil, common.DataNilError(r.ResourceType(), response)
	}

	for _, identityStoreProvisioner := range items {
		identityStoreProvisionerId, identityStoreProvisionerIdOk := identityStoreProvisioner.GetIdOk()
		identityStoreProvisionerName, identityStoreProvisionerNameOk := identityStoreProvisioner.GetNameOk()

		if identityStoreProvisionerIdOk && identityStoreProvisionerNameOk {
			identityStoreProvisionerData[*identityStoreProvisionerId] = *identityStoreProvisionerName
		}
	}

	return identityStoreProvisionerData, nil
}
