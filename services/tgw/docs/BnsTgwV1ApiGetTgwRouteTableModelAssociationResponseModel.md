# BnsTgwV1ApiGetTgwRouteTableModelAssociationResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**ResourceAttachmentId** | Pointer to **NullableString** |  | [optional] 
**ResourceId** | Pointer to **NullableString** |  | [optional] 
**ResourceType** | Pointer to [**NullableResourceType**](ResourceType.md) |  | [optional] 
**TgwRouteTableId** | Pointer to **NullableString** |  | [optional] 
**ProvisioningStatus** | Pointer to [**NullableTGWProvisioningStatus**](TGWProvisioningStatus.md) |  | [optional] 
**Resource** | [**BnsTgwV1ApiGetTgwRouteTableModelResourceResponseModel**](BnsTgwV1ApiGetTgwRouteTableModelResourceResponseModel.md) | 연결된 리소스 정보 | 

## Methods

### NewBnsTgwV1ApiGetTgwRouteTableModelAssociationResponseModel

`func NewBnsTgwV1ApiGetTgwRouteTableModelAssociationResponseModel(resource BnsTgwV1ApiGetTgwRouteTableModelResourceResponseModel, ) *BnsTgwV1ApiGetTgwRouteTableModelAssociationResponseModel`

NewBnsTgwV1ApiGetTgwRouteTableModelAssociationResponseModel instantiates a new BnsTgwV1ApiGetTgwRouteTableModelAssociationResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsTgwV1ApiGetTgwRouteTableModelAssociationResponseModelWithDefaults

`func NewBnsTgwV1ApiGetTgwRouteTableModelAssociationResponseModelWithDefaults() *BnsTgwV1ApiGetTgwRouteTableModelAssociationResponseModel`

NewBnsTgwV1ApiGetTgwRouteTableModelAssociationResponseModelWithDefaults instantiates a new BnsTgwV1ApiGetTgwRouteTableModelAssociationResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsTgwV1ApiGetTgwRouteTableModelAssociationResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsTgwV1ApiGetTgwRouteTableModelAssociationResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsTgwV1ApiGetTgwRouteTableModelAssociationResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BnsTgwV1ApiGetTgwRouteTableModelAssociationResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *BnsTgwV1ApiGetTgwRouteTableModelAssociationResponseModel) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *BnsTgwV1ApiGetTgwRouteTableModelAssociationResponseModel) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetResourceAttachmentId

`func (o *BnsTgwV1ApiGetTgwRouteTableModelAssociationResponseModel) GetResourceAttachmentId() string`

GetResourceAttachmentId returns the ResourceAttachmentId field if non-nil, zero value otherwise.

### GetResourceAttachmentIdOk

`func (o *BnsTgwV1ApiGetTgwRouteTableModelAssociationResponseModel) GetResourceAttachmentIdOk() (*string, bool)`

GetResourceAttachmentIdOk returns a tuple with the ResourceAttachmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceAttachmentId

`func (o *BnsTgwV1ApiGetTgwRouteTableModelAssociationResponseModel) SetResourceAttachmentId(v string)`

SetResourceAttachmentId sets ResourceAttachmentId field to given value.

### HasResourceAttachmentId

`func (o *BnsTgwV1ApiGetTgwRouteTableModelAssociationResponseModel) HasResourceAttachmentId() bool`

HasResourceAttachmentId returns a boolean if a field has been set.

### SetResourceAttachmentIdNil

`func (o *BnsTgwV1ApiGetTgwRouteTableModelAssociationResponseModel) SetResourceAttachmentIdNil(b bool)`

 SetResourceAttachmentIdNil sets the value for ResourceAttachmentId to be an explicit nil

### UnsetResourceAttachmentId
`func (o *BnsTgwV1ApiGetTgwRouteTableModelAssociationResponseModel) UnsetResourceAttachmentId()`

UnsetResourceAttachmentId ensures that no value is present for ResourceAttachmentId, not even an explicit nil
### GetResourceId

`func (o *BnsTgwV1ApiGetTgwRouteTableModelAssociationResponseModel) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *BnsTgwV1ApiGetTgwRouteTableModelAssociationResponseModel) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *BnsTgwV1ApiGetTgwRouteTableModelAssociationResponseModel) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *BnsTgwV1ApiGetTgwRouteTableModelAssociationResponseModel) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### SetResourceIdNil

`func (o *BnsTgwV1ApiGetTgwRouteTableModelAssociationResponseModel) SetResourceIdNil(b bool)`

 SetResourceIdNil sets the value for ResourceId to be an explicit nil

### UnsetResourceId
`func (o *BnsTgwV1ApiGetTgwRouteTableModelAssociationResponseModel) UnsetResourceId()`

UnsetResourceId ensures that no value is present for ResourceId, not even an explicit nil
### GetResourceType

`func (o *BnsTgwV1ApiGetTgwRouteTableModelAssociationResponseModel) GetResourceType() ResourceType`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *BnsTgwV1ApiGetTgwRouteTableModelAssociationResponseModel) GetResourceTypeOk() (*ResourceType, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *BnsTgwV1ApiGetTgwRouteTableModelAssociationResponseModel) SetResourceType(v ResourceType)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *BnsTgwV1ApiGetTgwRouteTableModelAssociationResponseModel) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### SetResourceTypeNil

`func (o *BnsTgwV1ApiGetTgwRouteTableModelAssociationResponseModel) SetResourceTypeNil(b bool)`

 SetResourceTypeNil sets the value for ResourceType to be an explicit nil

### UnsetResourceType
`func (o *BnsTgwV1ApiGetTgwRouteTableModelAssociationResponseModel) UnsetResourceType()`

UnsetResourceType ensures that no value is present for ResourceType, not even an explicit nil
### GetTgwRouteTableId

`func (o *BnsTgwV1ApiGetTgwRouteTableModelAssociationResponseModel) GetTgwRouteTableId() string`

GetTgwRouteTableId returns the TgwRouteTableId field if non-nil, zero value otherwise.

### GetTgwRouteTableIdOk

`func (o *BnsTgwV1ApiGetTgwRouteTableModelAssociationResponseModel) GetTgwRouteTableIdOk() (*string, bool)`

GetTgwRouteTableIdOk returns a tuple with the TgwRouteTableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgwRouteTableId

`func (o *BnsTgwV1ApiGetTgwRouteTableModelAssociationResponseModel) SetTgwRouteTableId(v string)`

SetTgwRouteTableId sets TgwRouteTableId field to given value.

### HasTgwRouteTableId

`func (o *BnsTgwV1ApiGetTgwRouteTableModelAssociationResponseModel) HasTgwRouteTableId() bool`

HasTgwRouteTableId returns a boolean if a field has been set.

### SetTgwRouteTableIdNil

`func (o *BnsTgwV1ApiGetTgwRouteTableModelAssociationResponseModel) SetTgwRouteTableIdNil(b bool)`

 SetTgwRouteTableIdNil sets the value for TgwRouteTableId to be an explicit nil

### UnsetTgwRouteTableId
`func (o *BnsTgwV1ApiGetTgwRouteTableModelAssociationResponseModel) UnsetTgwRouteTableId()`

UnsetTgwRouteTableId ensures that no value is present for TgwRouteTableId, not even an explicit nil
### GetProvisioningStatus

`func (o *BnsTgwV1ApiGetTgwRouteTableModelAssociationResponseModel) GetProvisioningStatus() TGWProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *BnsTgwV1ApiGetTgwRouteTableModelAssociationResponseModel) GetProvisioningStatusOk() (*TGWProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *BnsTgwV1ApiGetTgwRouteTableModelAssociationResponseModel) SetProvisioningStatus(v TGWProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.

### HasProvisioningStatus

`func (o *BnsTgwV1ApiGetTgwRouteTableModelAssociationResponseModel) HasProvisioningStatus() bool`

HasProvisioningStatus returns a boolean if a field has been set.

### SetProvisioningStatusNil

`func (o *BnsTgwV1ApiGetTgwRouteTableModelAssociationResponseModel) SetProvisioningStatusNil(b bool)`

 SetProvisioningStatusNil sets the value for ProvisioningStatus to be an explicit nil

### UnsetProvisioningStatus
`func (o *BnsTgwV1ApiGetTgwRouteTableModelAssociationResponseModel) UnsetProvisioningStatus()`

UnsetProvisioningStatus ensures that no value is present for ProvisioningStatus, not even an explicit nil
### GetResource

`func (o *BnsTgwV1ApiGetTgwRouteTableModelAssociationResponseModel) GetResource() BnsTgwV1ApiGetTgwRouteTableModelResourceResponseModel`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *BnsTgwV1ApiGetTgwRouteTableModelAssociationResponseModel) GetResourceOk() (*BnsTgwV1ApiGetTgwRouteTableModelResourceResponseModel, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *BnsTgwV1ApiGetTgwRouteTableModelAssociationResponseModel) SetResource(v BnsTgwV1ApiGetTgwRouteTableModelResourceResponseModel)`

SetResource sets Resource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


