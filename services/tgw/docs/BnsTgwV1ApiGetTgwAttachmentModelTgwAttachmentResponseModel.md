# BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Attachment ID | 
**ProvisioningStatus** | Pointer to [**NullableTGWProvisioningStatus**](TGWProvisioningStatus.md) |  | [optional] 
**ProjectId** | Pointer to **NullableString** |  | [optional] 
**ProjectName** | Pointer to **NullableString** |  | [optional] 
**ResourceType** | Pointer to [**NullableResourceType**](ResourceType.md) |  | [optional] 
**ResourceId** | Pointer to **NullableString** |  | [optional] 
**ResourceName** | Pointer to **NullableString** |  | [optional] 
**ResourceCidrBlock** | Pointer to **NullableString** |  | [optional] 
**Tgw** | [**TrasitGatewayResponseModel**](TrasitGatewayResponseModel.md) | 연결된 Transit Gateway 정보 | 
**Resources** | Pointer to [**[]BnsTgwV1ApiGetTgwAttachmentModelResourceResponseModel**](BnsTgwV1ApiGetTgwAttachmentModelResourceResponseModel.md) |  | [optional] 
**RouteTable** | Pointer to [**NullableBnsTgwV1ApiGetTgwAttachmentModelRouteTableResponseModel**](BnsTgwV1ApiGetTgwAttachmentModelRouteTableResponseModel.md) |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewBnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel

`func NewBnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel(id string, tgw TrasitGatewayResponseModel, ) *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel`

NewBnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel instantiates a new BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModelWithDefaults

`func NewBnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModelWithDefaults() *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel`

NewBnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModelWithDefaults instantiates a new BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetProvisioningStatus

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) GetProvisioningStatus() TGWProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) GetProvisioningStatusOk() (*TGWProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) SetProvisioningStatus(v TGWProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.

### HasProvisioningStatus

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) HasProvisioningStatus() bool`

HasProvisioningStatus returns a boolean if a field has been set.

### SetProvisioningStatusNil

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) SetProvisioningStatusNil(b bool)`

 SetProvisioningStatusNil sets the value for ProvisioningStatus to be an explicit nil

### UnsetProvisioningStatus
`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) UnsetProvisioningStatus()`

UnsetProvisioningStatus ensures that no value is present for ProvisioningStatus, not even an explicit nil
### GetProjectId

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetProjectName

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### SetProjectNameNil

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) SetProjectNameNil(b bool)`

 SetProjectNameNil sets the value for ProjectName to be an explicit nil

### UnsetProjectName
`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) UnsetProjectName()`

UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil
### GetResourceType

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) GetResourceType() ResourceType`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) GetResourceTypeOk() (*ResourceType, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) SetResourceType(v ResourceType)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### SetResourceTypeNil

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) SetResourceTypeNil(b bool)`

 SetResourceTypeNil sets the value for ResourceType to be an explicit nil

### UnsetResourceType
`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) UnsetResourceType()`

UnsetResourceType ensures that no value is present for ResourceType, not even an explicit nil
### GetResourceId

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### SetResourceIdNil

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) SetResourceIdNil(b bool)`

 SetResourceIdNil sets the value for ResourceId to be an explicit nil

### UnsetResourceId
`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) UnsetResourceId()`

UnsetResourceId ensures that no value is present for ResourceId, not even an explicit nil
### GetResourceName

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### SetResourceNameNil

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) SetResourceNameNil(b bool)`

 SetResourceNameNil sets the value for ResourceName to be an explicit nil

### UnsetResourceName
`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) UnsetResourceName()`

UnsetResourceName ensures that no value is present for ResourceName, not even an explicit nil
### GetResourceCidrBlock

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) GetResourceCidrBlock() string`

GetResourceCidrBlock returns the ResourceCidrBlock field if non-nil, zero value otherwise.

### GetResourceCidrBlockOk

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) GetResourceCidrBlockOk() (*string, bool)`

GetResourceCidrBlockOk returns a tuple with the ResourceCidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceCidrBlock

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) SetResourceCidrBlock(v string)`

SetResourceCidrBlock sets ResourceCidrBlock field to given value.

### HasResourceCidrBlock

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) HasResourceCidrBlock() bool`

HasResourceCidrBlock returns a boolean if a field has been set.

### SetResourceCidrBlockNil

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) SetResourceCidrBlockNil(b bool)`

 SetResourceCidrBlockNil sets the value for ResourceCidrBlock to be an explicit nil

### UnsetResourceCidrBlock
`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) UnsetResourceCidrBlock()`

UnsetResourceCidrBlock ensures that no value is present for ResourceCidrBlock, not even an explicit nil
### GetTgw

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) GetTgw() TrasitGatewayResponseModel`

GetTgw returns the Tgw field if non-nil, zero value otherwise.

### GetTgwOk

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) GetTgwOk() (*TrasitGatewayResponseModel, bool)`

GetTgwOk returns a tuple with the Tgw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgw

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) SetTgw(v TrasitGatewayResponseModel)`

SetTgw sets Tgw field to given value.


### GetResources

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) GetResources() []BnsTgwV1ApiGetTgwAttachmentModelResourceResponseModel`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) GetResourcesOk() (*[]BnsTgwV1ApiGetTgwAttachmentModelResourceResponseModel, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) SetResources(v []BnsTgwV1ApiGetTgwAttachmentModelResourceResponseModel)`

SetResources sets Resources field to given value.

### HasResources

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) HasResources() bool`

HasResources returns a boolean if a field has been set.

### SetResourcesNil

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) SetResourcesNil(b bool)`

 SetResourcesNil sets the value for Resources to be an explicit nil

### UnsetResources
`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) UnsetResources()`

UnsetResources ensures that no value is present for Resources, not even an explicit nil
### GetRouteTable

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) GetRouteTable() BnsTgwV1ApiGetTgwAttachmentModelRouteTableResponseModel`

GetRouteTable returns the RouteTable field if non-nil, zero value otherwise.

### GetRouteTableOk

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) GetRouteTableOk() (*BnsTgwV1ApiGetTgwAttachmentModelRouteTableResponseModel, bool)`

GetRouteTableOk returns a tuple with the RouteTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteTable

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) SetRouteTable(v BnsTgwV1ApiGetTgwAttachmentModelRouteTableResponseModel)`

SetRouteTable sets RouteTable field to given value.

### HasRouteTable

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) HasRouteTable() bool`

HasRouteTable returns a boolean if a field has been set.

### SetRouteTableNil

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) SetRouteTableNil(b bool)`

 SetRouteTableNil sets the value for RouteTable to be an explicit nil

### UnsetRouteTable
`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) UnsetRouteTable()`

UnsetRouteTable ensures that no value is present for RouteTable, not even an explicit nil
### GetCreatedAt

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *BnsTgwV1ApiGetTgwAttachmentModelTgwAttachmentResponseModel) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


