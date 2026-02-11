# BnsTgwV1ApiListTgwRouteTablesModelAssociationResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**ResourceAttachmentId** | Pointer to **NullableString** |  | [optional] 
**ResourceId** | Pointer to **NullableString** |  | [optional] 
**ResourceType** | Pointer to [**NullableResourceType**](ResourceType.md) |  | [optional] 
**TgwRouteTableId** | Pointer to **NullableString** |  | [optional] 
**ProvisioningStatus** | Pointer to [**NullableTGWProvisioningStatus**](TGWProvisioningStatus.md) |  | [optional] 
**Resource** | [**BnsTgwV1ApiListTgwRouteTablesModelResourceResponseModel**](BnsTgwV1ApiListTgwRouteTablesModelResourceResponseModel.md) | 연결된 리소스의 기본 정보 객체 | 

## Methods

### NewBnsTgwV1ApiListTgwRouteTablesModelAssociationResponseModel

`func NewBnsTgwV1ApiListTgwRouteTablesModelAssociationResponseModel(resource BnsTgwV1ApiListTgwRouteTablesModelResourceResponseModel, ) *BnsTgwV1ApiListTgwRouteTablesModelAssociationResponseModel`

NewBnsTgwV1ApiListTgwRouteTablesModelAssociationResponseModel instantiates a new BnsTgwV1ApiListTgwRouteTablesModelAssociationResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsTgwV1ApiListTgwRouteTablesModelAssociationResponseModelWithDefaults

`func NewBnsTgwV1ApiListTgwRouteTablesModelAssociationResponseModelWithDefaults() *BnsTgwV1ApiListTgwRouteTablesModelAssociationResponseModel`

NewBnsTgwV1ApiListTgwRouteTablesModelAssociationResponseModelWithDefaults instantiates a new BnsTgwV1ApiListTgwRouteTablesModelAssociationResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsTgwV1ApiListTgwRouteTablesModelAssociationResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsTgwV1ApiListTgwRouteTablesModelAssociationResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsTgwV1ApiListTgwRouteTablesModelAssociationResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BnsTgwV1ApiListTgwRouteTablesModelAssociationResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *BnsTgwV1ApiListTgwRouteTablesModelAssociationResponseModel) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *BnsTgwV1ApiListTgwRouteTablesModelAssociationResponseModel) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetResourceAttachmentId

`func (o *BnsTgwV1ApiListTgwRouteTablesModelAssociationResponseModel) GetResourceAttachmentId() string`

GetResourceAttachmentId returns the ResourceAttachmentId field if non-nil, zero value otherwise.

### GetResourceAttachmentIdOk

`func (o *BnsTgwV1ApiListTgwRouteTablesModelAssociationResponseModel) GetResourceAttachmentIdOk() (*string, bool)`

GetResourceAttachmentIdOk returns a tuple with the ResourceAttachmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceAttachmentId

`func (o *BnsTgwV1ApiListTgwRouteTablesModelAssociationResponseModel) SetResourceAttachmentId(v string)`

SetResourceAttachmentId sets ResourceAttachmentId field to given value.

### HasResourceAttachmentId

`func (o *BnsTgwV1ApiListTgwRouteTablesModelAssociationResponseModel) HasResourceAttachmentId() bool`

HasResourceAttachmentId returns a boolean if a field has been set.

### SetResourceAttachmentIdNil

`func (o *BnsTgwV1ApiListTgwRouteTablesModelAssociationResponseModel) SetResourceAttachmentIdNil(b bool)`

 SetResourceAttachmentIdNil sets the value for ResourceAttachmentId to be an explicit nil

### UnsetResourceAttachmentId
`func (o *BnsTgwV1ApiListTgwRouteTablesModelAssociationResponseModel) UnsetResourceAttachmentId()`

UnsetResourceAttachmentId ensures that no value is present for ResourceAttachmentId, not even an explicit nil
### GetResourceId

`func (o *BnsTgwV1ApiListTgwRouteTablesModelAssociationResponseModel) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *BnsTgwV1ApiListTgwRouteTablesModelAssociationResponseModel) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *BnsTgwV1ApiListTgwRouteTablesModelAssociationResponseModel) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *BnsTgwV1ApiListTgwRouteTablesModelAssociationResponseModel) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### SetResourceIdNil

`func (o *BnsTgwV1ApiListTgwRouteTablesModelAssociationResponseModel) SetResourceIdNil(b bool)`

 SetResourceIdNil sets the value for ResourceId to be an explicit nil

### UnsetResourceId
`func (o *BnsTgwV1ApiListTgwRouteTablesModelAssociationResponseModel) UnsetResourceId()`

UnsetResourceId ensures that no value is present for ResourceId, not even an explicit nil
### GetResourceType

`func (o *BnsTgwV1ApiListTgwRouteTablesModelAssociationResponseModel) GetResourceType() ResourceType`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *BnsTgwV1ApiListTgwRouteTablesModelAssociationResponseModel) GetResourceTypeOk() (*ResourceType, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *BnsTgwV1ApiListTgwRouteTablesModelAssociationResponseModel) SetResourceType(v ResourceType)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *BnsTgwV1ApiListTgwRouteTablesModelAssociationResponseModel) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### SetResourceTypeNil

`func (o *BnsTgwV1ApiListTgwRouteTablesModelAssociationResponseModel) SetResourceTypeNil(b bool)`

 SetResourceTypeNil sets the value for ResourceType to be an explicit nil

### UnsetResourceType
`func (o *BnsTgwV1ApiListTgwRouteTablesModelAssociationResponseModel) UnsetResourceType()`

UnsetResourceType ensures that no value is present for ResourceType, not even an explicit nil
### GetTgwRouteTableId

`func (o *BnsTgwV1ApiListTgwRouteTablesModelAssociationResponseModel) GetTgwRouteTableId() string`

GetTgwRouteTableId returns the TgwRouteTableId field if non-nil, zero value otherwise.

### GetTgwRouteTableIdOk

`func (o *BnsTgwV1ApiListTgwRouteTablesModelAssociationResponseModel) GetTgwRouteTableIdOk() (*string, bool)`

GetTgwRouteTableIdOk returns a tuple with the TgwRouteTableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgwRouteTableId

`func (o *BnsTgwV1ApiListTgwRouteTablesModelAssociationResponseModel) SetTgwRouteTableId(v string)`

SetTgwRouteTableId sets TgwRouteTableId field to given value.

### HasTgwRouteTableId

`func (o *BnsTgwV1ApiListTgwRouteTablesModelAssociationResponseModel) HasTgwRouteTableId() bool`

HasTgwRouteTableId returns a boolean if a field has been set.

### SetTgwRouteTableIdNil

`func (o *BnsTgwV1ApiListTgwRouteTablesModelAssociationResponseModel) SetTgwRouteTableIdNil(b bool)`

 SetTgwRouteTableIdNil sets the value for TgwRouteTableId to be an explicit nil

### UnsetTgwRouteTableId
`func (o *BnsTgwV1ApiListTgwRouteTablesModelAssociationResponseModel) UnsetTgwRouteTableId()`

UnsetTgwRouteTableId ensures that no value is present for TgwRouteTableId, not even an explicit nil
### GetProvisioningStatus

`func (o *BnsTgwV1ApiListTgwRouteTablesModelAssociationResponseModel) GetProvisioningStatus() TGWProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *BnsTgwV1ApiListTgwRouteTablesModelAssociationResponseModel) GetProvisioningStatusOk() (*TGWProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *BnsTgwV1ApiListTgwRouteTablesModelAssociationResponseModel) SetProvisioningStatus(v TGWProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.

### HasProvisioningStatus

`func (o *BnsTgwV1ApiListTgwRouteTablesModelAssociationResponseModel) HasProvisioningStatus() bool`

HasProvisioningStatus returns a boolean if a field has been set.

### SetProvisioningStatusNil

`func (o *BnsTgwV1ApiListTgwRouteTablesModelAssociationResponseModel) SetProvisioningStatusNil(b bool)`

 SetProvisioningStatusNil sets the value for ProvisioningStatus to be an explicit nil

### UnsetProvisioningStatus
`func (o *BnsTgwV1ApiListTgwRouteTablesModelAssociationResponseModel) UnsetProvisioningStatus()`

UnsetProvisioningStatus ensures that no value is present for ProvisioningStatus, not even an explicit nil
### GetResource

`func (o *BnsTgwV1ApiListTgwRouteTablesModelAssociationResponseModel) GetResource() BnsTgwV1ApiListTgwRouteTablesModelResourceResponseModel`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *BnsTgwV1ApiListTgwRouteTablesModelAssociationResponseModel) GetResourceOk() (*BnsTgwV1ApiListTgwRouteTablesModelResourceResponseModel, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *BnsTgwV1ApiListTgwRouteTablesModelAssociationResponseModel) SetResource(v BnsTgwV1ApiListTgwRouteTablesModelResourceResponseModel)`

SetResource sets Resource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


