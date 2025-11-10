# BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 연결된 리소스 ID | 
**ProvisioningStatus** | Pointer to [**NullableTGWProvisioningStatus**](TGWProvisioningStatus.md) |  | [optional] 
**ProjectId** | Pointer to **NullableString** |  | [optional] 
**ProjectName** | Pointer to **NullableString** |  | [optional] 
**ResourceType** | Pointer to [**NullableResourceType**](ResourceType.md) |  | [optional] 
**ResourceId** | Pointer to **NullableString** |  | [optional] 
**ResourceName** | Pointer to **NullableString** |  | [optional] 
**ResourceCidrBlock** | Pointer to **NullableString** |  | [optional] 
**Tgw** | [**TrasitGatewayResponseModel**](TrasitGatewayResponseModel.md) | Transit Gateway 정보 | 
**Resources** | Pointer to [**[]BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel**](BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel.md) |  | [optional] 
**RouteTable** | Pointer to [**NullableBnsTgwV1ApiListTgwAttachmentsModelRouteTableResponseModel**](BnsTgwV1ApiListTgwAttachmentsModelRouteTableResponseModel.md) |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewBnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel

`func NewBnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel(id string, tgw TrasitGatewayResponseModel, ) *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel`

NewBnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel instantiates a new BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModelWithDefaults

`func NewBnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModelWithDefaults() *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel`

NewBnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModelWithDefaults instantiates a new BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetProvisioningStatus

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) GetProvisioningStatus() TGWProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) GetProvisioningStatusOk() (*TGWProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) SetProvisioningStatus(v TGWProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.

### HasProvisioningStatus

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) HasProvisioningStatus() bool`

HasProvisioningStatus returns a boolean if a field has been set.

### SetProvisioningStatusNil

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) SetProvisioningStatusNil(b bool)`

 SetProvisioningStatusNil sets the value for ProvisioningStatus to be an explicit nil

### UnsetProvisioningStatus
`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) UnsetProvisioningStatus()`

UnsetProvisioningStatus ensures that no value is present for ProvisioningStatus, not even an explicit nil
### GetProjectId

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetProjectName

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### SetProjectNameNil

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) SetProjectNameNil(b bool)`

 SetProjectNameNil sets the value for ProjectName to be an explicit nil

### UnsetProjectName
`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) UnsetProjectName()`

UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil
### GetResourceType

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) GetResourceType() ResourceType`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) GetResourceTypeOk() (*ResourceType, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) SetResourceType(v ResourceType)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### SetResourceTypeNil

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) SetResourceTypeNil(b bool)`

 SetResourceTypeNil sets the value for ResourceType to be an explicit nil

### UnsetResourceType
`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) UnsetResourceType()`

UnsetResourceType ensures that no value is present for ResourceType, not even an explicit nil
### GetResourceId

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### SetResourceIdNil

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) SetResourceIdNil(b bool)`

 SetResourceIdNil sets the value for ResourceId to be an explicit nil

### UnsetResourceId
`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) UnsetResourceId()`

UnsetResourceId ensures that no value is present for ResourceId, not even an explicit nil
### GetResourceName

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### SetResourceNameNil

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) SetResourceNameNil(b bool)`

 SetResourceNameNil sets the value for ResourceName to be an explicit nil

### UnsetResourceName
`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) UnsetResourceName()`

UnsetResourceName ensures that no value is present for ResourceName, not even an explicit nil
### GetResourceCidrBlock

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) GetResourceCidrBlock() string`

GetResourceCidrBlock returns the ResourceCidrBlock field if non-nil, zero value otherwise.

### GetResourceCidrBlockOk

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) GetResourceCidrBlockOk() (*string, bool)`

GetResourceCidrBlockOk returns a tuple with the ResourceCidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceCidrBlock

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) SetResourceCidrBlock(v string)`

SetResourceCidrBlock sets ResourceCidrBlock field to given value.

### HasResourceCidrBlock

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) HasResourceCidrBlock() bool`

HasResourceCidrBlock returns a boolean if a field has been set.

### SetResourceCidrBlockNil

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) SetResourceCidrBlockNil(b bool)`

 SetResourceCidrBlockNil sets the value for ResourceCidrBlock to be an explicit nil

### UnsetResourceCidrBlock
`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) UnsetResourceCidrBlock()`

UnsetResourceCidrBlock ensures that no value is present for ResourceCidrBlock, not even an explicit nil
### GetTgw

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) GetTgw() TrasitGatewayResponseModel`

GetTgw returns the Tgw field if non-nil, zero value otherwise.

### GetTgwOk

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) GetTgwOk() (*TrasitGatewayResponseModel, bool)`

GetTgwOk returns a tuple with the Tgw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgw

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) SetTgw(v TrasitGatewayResponseModel)`

SetTgw sets Tgw field to given value.


### GetResources

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) GetResources() []BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) GetResourcesOk() (*[]BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) SetResources(v []BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel)`

SetResources sets Resources field to given value.

### HasResources

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) HasResources() bool`

HasResources returns a boolean if a field has been set.

### SetResourcesNil

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) SetResourcesNil(b bool)`

 SetResourcesNil sets the value for Resources to be an explicit nil

### UnsetResources
`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) UnsetResources()`

UnsetResources ensures that no value is present for Resources, not even an explicit nil
### GetRouteTable

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) GetRouteTable() BnsTgwV1ApiListTgwAttachmentsModelRouteTableResponseModel`

GetRouteTable returns the RouteTable field if non-nil, zero value otherwise.

### GetRouteTableOk

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) GetRouteTableOk() (*BnsTgwV1ApiListTgwAttachmentsModelRouteTableResponseModel, bool)`

GetRouteTableOk returns a tuple with the RouteTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteTable

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) SetRouteTable(v BnsTgwV1ApiListTgwAttachmentsModelRouteTableResponseModel)`

SetRouteTable sets RouteTable field to given value.

### HasRouteTable

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) HasRouteTable() bool`

HasRouteTable returns a boolean if a field has been set.

### SetRouteTableNil

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) SetRouteTableNil(b bool)`

 SetRouteTableNil sets the value for RouteTable to be an explicit nil

### UnsetRouteTable
`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) UnsetRouteTable()`

UnsetRouteTable ensures that no value is present for RouteTable, not even an explicit nil
### GetCreatedAt

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


