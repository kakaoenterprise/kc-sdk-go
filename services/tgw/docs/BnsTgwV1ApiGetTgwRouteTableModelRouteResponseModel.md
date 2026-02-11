# BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**RouteType** | Pointer to **NullableString** |  | [optional] 
**DestinationCidrBlock** | Pointer to **NullableString** |  | [optional] 
**ResourceAttachmentId** | Pointer to **NullableString** |  | [optional] 
**ResourceId** | Pointer to **NullableString** |  | [optional] 
**ResourceType** | Pointer to [**NullableResourceType**](ResourceType.md) |  | [optional] 
**TgwRouteTableId** | Pointer to **NullableString** |  | [optional] 
**ProvisioningStatus** | Pointer to [**NullableTGWProvisioningStatus**](TGWProvisioningStatus.md) |  | [optional] 
**Resource** | [**BnsTgwV1ApiGetTgwRouteTableModelResourceResponseModel**](BnsTgwV1ApiGetTgwRouteTableModelResourceResponseModel.md) | 연결된 리소스 정보 | 

## Methods

### NewBnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel

`func NewBnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel(resource BnsTgwV1ApiGetTgwRouteTableModelResourceResponseModel, ) *BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel`

NewBnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel instantiates a new BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsTgwV1ApiGetTgwRouteTableModelRouteResponseModelWithDefaults

`func NewBnsTgwV1ApiGetTgwRouteTableModelRouteResponseModelWithDefaults() *BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel`

NewBnsTgwV1ApiGetTgwRouteTableModelRouteResponseModelWithDefaults instantiates a new BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetRouteType

`func (o *BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel) GetRouteType() string`

GetRouteType returns the RouteType field if non-nil, zero value otherwise.

### GetRouteTypeOk

`func (o *BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel) GetRouteTypeOk() (*string, bool)`

GetRouteTypeOk returns a tuple with the RouteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteType

`func (o *BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel) SetRouteType(v string)`

SetRouteType sets RouteType field to given value.

### HasRouteType

`func (o *BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel) HasRouteType() bool`

HasRouteType returns a boolean if a field has been set.

### SetRouteTypeNil

`func (o *BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel) SetRouteTypeNil(b bool)`

 SetRouteTypeNil sets the value for RouteType to be an explicit nil

### UnsetRouteType
`func (o *BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel) UnsetRouteType()`

UnsetRouteType ensures that no value is present for RouteType, not even an explicit nil
### GetDestinationCidrBlock

`func (o *BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel) GetDestinationCidrBlock() string`

GetDestinationCidrBlock returns the DestinationCidrBlock field if non-nil, zero value otherwise.

### GetDestinationCidrBlockOk

`func (o *BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel) GetDestinationCidrBlockOk() (*string, bool)`

GetDestinationCidrBlockOk returns a tuple with the DestinationCidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationCidrBlock

`func (o *BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel) SetDestinationCidrBlock(v string)`

SetDestinationCidrBlock sets DestinationCidrBlock field to given value.

### HasDestinationCidrBlock

`func (o *BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel) HasDestinationCidrBlock() bool`

HasDestinationCidrBlock returns a boolean if a field has been set.

### SetDestinationCidrBlockNil

`func (o *BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel) SetDestinationCidrBlockNil(b bool)`

 SetDestinationCidrBlockNil sets the value for DestinationCidrBlock to be an explicit nil

### UnsetDestinationCidrBlock
`func (o *BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel) UnsetDestinationCidrBlock()`

UnsetDestinationCidrBlock ensures that no value is present for DestinationCidrBlock, not even an explicit nil
### GetResourceAttachmentId

`func (o *BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel) GetResourceAttachmentId() string`

GetResourceAttachmentId returns the ResourceAttachmentId field if non-nil, zero value otherwise.

### GetResourceAttachmentIdOk

`func (o *BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel) GetResourceAttachmentIdOk() (*string, bool)`

GetResourceAttachmentIdOk returns a tuple with the ResourceAttachmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceAttachmentId

`func (o *BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel) SetResourceAttachmentId(v string)`

SetResourceAttachmentId sets ResourceAttachmentId field to given value.

### HasResourceAttachmentId

`func (o *BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel) HasResourceAttachmentId() bool`

HasResourceAttachmentId returns a boolean if a field has been set.

### SetResourceAttachmentIdNil

`func (o *BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel) SetResourceAttachmentIdNil(b bool)`

 SetResourceAttachmentIdNil sets the value for ResourceAttachmentId to be an explicit nil

### UnsetResourceAttachmentId
`func (o *BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel) UnsetResourceAttachmentId()`

UnsetResourceAttachmentId ensures that no value is present for ResourceAttachmentId, not even an explicit nil
### GetResourceId

`func (o *BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### SetResourceIdNil

`func (o *BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel) SetResourceIdNil(b bool)`

 SetResourceIdNil sets the value for ResourceId to be an explicit nil

### UnsetResourceId
`func (o *BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel) UnsetResourceId()`

UnsetResourceId ensures that no value is present for ResourceId, not even an explicit nil
### GetResourceType

`func (o *BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel) GetResourceType() ResourceType`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel) GetResourceTypeOk() (*ResourceType, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel) SetResourceType(v ResourceType)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### SetResourceTypeNil

`func (o *BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel) SetResourceTypeNil(b bool)`

 SetResourceTypeNil sets the value for ResourceType to be an explicit nil

### UnsetResourceType
`func (o *BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel) UnsetResourceType()`

UnsetResourceType ensures that no value is present for ResourceType, not even an explicit nil
### GetTgwRouteTableId

`func (o *BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel) GetTgwRouteTableId() string`

GetTgwRouteTableId returns the TgwRouteTableId field if non-nil, zero value otherwise.

### GetTgwRouteTableIdOk

`func (o *BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel) GetTgwRouteTableIdOk() (*string, bool)`

GetTgwRouteTableIdOk returns a tuple with the TgwRouteTableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgwRouteTableId

`func (o *BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel) SetTgwRouteTableId(v string)`

SetTgwRouteTableId sets TgwRouteTableId field to given value.

### HasTgwRouteTableId

`func (o *BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel) HasTgwRouteTableId() bool`

HasTgwRouteTableId returns a boolean if a field has been set.

### SetTgwRouteTableIdNil

`func (o *BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel) SetTgwRouteTableIdNil(b bool)`

 SetTgwRouteTableIdNil sets the value for TgwRouteTableId to be an explicit nil

### UnsetTgwRouteTableId
`func (o *BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel) UnsetTgwRouteTableId()`

UnsetTgwRouteTableId ensures that no value is present for TgwRouteTableId, not even an explicit nil
### GetProvisioningStatus

`func (o *BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel) GetProvisioningStatus() TGWProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel) GetProvisioningStatusOk() (*TGWProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel) SetProvisioningStatus(v TGWProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.

### HasProvisioningStatus

`func (o *BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel) HasProvisioningStatus() bool`

HasProvisioningStatus returns a boolean if a field has been set.

### SetProvisioningStatusNil

`func (o *BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel) SetProvisioningStatusNil(b bool)`

 SetProvisioningStatusNil sets the value for ProvisioningStatus to be an explicit nil

### UnsetProvisioningStatus
`func (o *BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel) UnsetProvisioningStatus()`

UnsetProvisioningStatus ensures that no value is present for ProvisioningStatus, not even an explicit nil
### GetResource

`func (o *BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel) GetResource() BnsTgwV1ApiGetTgwRouteTableModelResourceResponseModel`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel) GetResourceOk() (*BnsTgwV1ApiGetTgwRouteTableModelResourceResponseModel, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel) SetResource(v BnsTgwV1ApiGetTgwRouteTableModelResourceResponseModel)`

SetResource sets Resource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


