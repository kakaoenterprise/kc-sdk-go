# BnsTgwV1ApiListTgwRouteTableAssociationsModelAssociationResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**ResourceAttachmentId** | Pointer to **NullableString** |  | [optional] 
**ResourceId** | Pointer to **NullableString** |  | [optional] 
**ResourceType** | Pointer to [**NullableResourceType**](ResourceType.md) |  | [optional] 
**TgwRouteTableId** | Pointer to **NullableString** |  | [optional] 
**ProvisioningStatus** | Pointer to [**NullableTGWProvisioningStatus**](TGWProvisioningStatus.md) |  | [optional] 
**Resource** | [**BnsTgwV1ApiListTgwRouteTableAssociationsModelResourceResponseModel**](BnsTgwV1ApiListTgwRouteTableAssociationsModelResourceResponseModel.md) | Association에 연결된 리소스 정보 객체 | 

## Methods

### NewBnsTgwV1ApiListTgwRouteTableAssociationsModelAssociationResponseModel

`func NewBnsTgwV1ApiListTgwRouteTableAssociationsModelAssociationResponseModel(resource BnsTgwV1ApiListTgwRouteTableAssociationsModelResourceResponseModel, ) *BnsTgwV1ApiListTgwRouteTableAssociationsModelAssociationResponseModel`

NewBnsTgwV1ApiListTgwRouteTableAssociationsModelAssociationResponseModel instantiates a new BnsTgwV1ApiListTgwRouteTableAssociationsModelAssociationResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsTgwV1ApiListTgwRouteTableAssociationsModelAssociationResponseModelWithDefaults

`func NewBnsTgwV1ApiListTgwRouteTableAssociationsModelAssociationResponseModelWithDefaults() *BnsTgwV1ApiListTgwRouteTableAssociationsModelAssociationResponseModel`

NewBnsTgwV1ApiListTgwRouteTableAssociationsModelAssociationResponseModelWithDefaults instantiates a new BnsTgwV1ApiListTgwRouteTableAssociationsModelAssociationResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsTgwV1ApiListTgwRouteTableAssociationsModelAssociationResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsTgwV1ApiListTgwRouteTableAssociationsModelAssociationResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsTgwV1ApiListTgwRouteTableAssociationsModelAssociationResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BnsTgwV1ApiListTgwRouteTableAssociationsModelAssociationResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *BnsTgwV1ApiListTgwRouteTableAssociationsModelAssociationResponseModel) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *BnsTgwV1ApiListTgwRouteTableAssociationsModelAssociationResponseModel) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetResourceAttachmentId

`func (o *BnsTgwV1ApiListTgwRouteTableAssociationsModelAssociationResponseModel) GetResourceAttachmentId() string`

GetResourceAttachmentId returns the ResourceAttachmentId field if non-nil, zero value otherwise.

### GetResourceAttachmentIdOk

`func (o *BnsTgwV1ApiListTgwRouteTableAssociationsModelAssociationResponseModel) GetResourceAttachmentIdOk() (*string, bool)`

GetResourceAttachmentIdOk returns a tuple with the ResourceAttachmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceAttachmentId

`func (o *BnsTgwV1ApiListTgwRouteTableAssociationsModelAssociationResponseModel) SetResourceAttachmentId(v string)`

SetResourceAttachmentId sets ResourceAttachmentId field to given value.

### HasResourceAttachmentId

`func (o *BnsTgwV1ApiListTgwRouteTableAssociationsModelAssociationResponseModel) HasResourceAttachmentId() bool`

HasResourceAttachmentId returns a boolean if a field has been set.

### SetResourceAttachmentIdNil

`func (o *BnsTgwV1ApiListTgwRouteTableAssociationsModelAssociationResponseModel) SetResourceAttachmentIdNil(b bool)`

 SetResourceAttachmentIdNil sets the value for ResourceAttachmentId to be an explicit nil

### UnsetResourceAttachmentId
`func (o *BnsTgwV1ApiListTgwRouteTableAssociationsModelAssociationResponseModel) UnsetResourceAttachmentId()`

UnsetResourceAttachmentId ensures that no value is present for ResourceAttachmentId, not even an explicit nil
### GetResourceId

`func (o *BnsTgwV1ApiListTgwRouteTableAssociationsModelAssociationResponseModel) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *BnsTgwV1ApiListTgwRouteTableAssociationsModelAssociationResponseModel) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *BnsTgwV1ApiListTgwRouteTableAssociationsModelAssociationResponseModel) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *BnsTgwV1ApiListTgwRouteTableAssociationsModelAssociationResponseModel) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### SetResourceIdNil

`func (o *BnsTgwV1ApiListTgwRouteTableAssociationsModelAssociationResponseModel) SetResourceIdNil(b bool)`

 SetResourceIdNil sets the value for ResourceId to be an explicit nil

### UnsetResourceId
`func (o *BnsTgwV1ApiListTgwRouteTableAssociationsModelAssociationResponseModel) UnsetResourceId()`

UnsetResourceId ensures that no value is present for ResourceId, not even an explicit nil
### GetResourceType

`func (o *BnsTgwV1ApiListTgwRouteTableAssociationsModelAssociationResponseModel) GetResourceType() ResourceType`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *BnsTgwV1ApiListTgwRouteTableAssociationsModelAssociationResponseModel) GetResourceTypeOk() (*ResourceType, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *BnsTgwV1ApiListTgwRouteTableAssociationsModelAssociationResponseModel) SetResourceType(v ResourceType)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *BnsTgwV1ApiListTgwRouteTableAssociationsModelAssociationResponseModel) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### SetResourceTypeNil

`func (o *BnsTgwV1ApiListTgwRouteTableAssociationsModelAssociationResponseModel) SetResourceTypeNil(b bool)`

 SetResourceTypeNil sets the value for ResourceType to be an explicit nil

### UnsetResourceType
`func (o *BnsTgwV1ApiListTgwRouteTableAssociationsModelAssociationResponseModel) UnsetResourceType()`

UnsetResourceType ensures that no value is present for ResourceType, not even an explicit nil
### GetTgwRouteTableId

`func (o *BnsTgwV1ApiListTgwRouteTableAssociationsModelAssociationResponseModel) GetTgwRouteTableId() string`

GetTgwRouteTableId returns the TgwRouteTableId field if non-nil, zero value otherwise.

### GetTgwRouteTableIdOk

`func (o *BnsTgwV1ApiListTgwRouteTableAssociationsModelAssociationResponseModel) GetTgwRouteTableIdOk() (*string, bool)`

GetTgwRouteTableIdOk returns a tuple with the TgwRouteTableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgwRouteTableId

`func (o *BnsTgwV1ApiListTgwRouteTableAssociationsModelAssociationResponseModel) SetTgwRouteTableId(v string)`

SetTgwRouteTableId sets TgwRouteTableId field to given value.

### HasTgwRouteTableId

`func (o *BnsTgwV1ApiListTgwRouteTableAssociationsModelAssociationResponseModel) HasTgwRouteTableId() bool`

HasTgwRouteTableId returns a boolean if a field has been set.

### SetTgwRouteTableIdNil

`func (o *BnsTgwV1ApiListTgwRouteTableAssociationsModelAssociationResponseModel) SetTgwRouteTableIdNil(b bool)`

 SetTgwRouteTableIdNil sets the value for TgwRouteTableId to be an explicit nil

### UnsetTgwRouteTableId
`func (o *BnsTgwV1ApiListTgwRouteTableAssociationsModelAssociationResponseModel) UnsetTgwRouteTableId()`

UnsetTgwRouteTableId ensures that no value is present for TgwRouteTableId, not even an explicit nil
### GetProvisioningStatus

`func (o *BnsTgwV1ApiListTgwRouteTableAssociationsModelAssociationResponseModel) GetProvisioningStatus() TGWProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *BnsTgwV1ApiListTgwRouteTableAssociationsModelAssociationResponseModel) GetProvisioningStatusOk() (*TGWProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *BnsTgwV1ApiListTgwRouteTableAssociationsModelAssociationResponseModel) SetProvisioningStatus(v TGWProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.

### HasProvisioningStatus

`func (o *BnsTgwV1ApiListTgwRouteTableAssociationsModelAssociationResponseModel) HasProvisioningStatus() bool`

HasProvisioningStatus returns a boolean if a field has been set.

### SetProvisioningStatusNil

`func (o *BnsTgwV1ApiListTgwRouteTableAssociationsModelAssociationResponseModel) SetProvisioningStatusNil(b bool)`

 SetProvisioningStatusNil sets the value for ProvisioningStatus to be an explicit nil

### UnsetProvisioningStatus
`func (o *BnsTgwV1ApiListTgwRouteTableAssociationsModelAssociationResponseModel) UnsetProvisioningStatus()`

UnsetProvisioningStatus ensures that no value is present for ProvisioningStatus, not even an explicit nil
### GetResource

`func (o *BnsTgwV1ApiListTgwRouteTableAssociationsModelAssociationResponseModel) GetResource() BnsTgwV1ApiListTgwRouteTableAssociationsModelResourceResponseModel`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *BnsTgwV1ApiListTgwRouteTableAssociationsModelAssociationResponseModel) GetResourceOk() (*BnsTgwV1ApiListTgwRouteTableAssociationsModelResourceResponseModel, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *BnsTgwV1ApiListTgwRouteTableAssociationsModelAssociationResponseModel) SetResource(v BnsTgwV1ApiListTgwRouteTableAssociationsModelResourceResponseModel)`

SetResource sets Resource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


