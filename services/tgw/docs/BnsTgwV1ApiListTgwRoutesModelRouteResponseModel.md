# BnsTgwV1ApiListTgwRoutesModelRouteResponseModel

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
**Resource** | [**BnsTgwV1ApiListTgwRoutesModelResourceResponseModel**](BnsTgwV1ApiListTgwRoutesModelResourceResponseModel.md) | 연결된 리소스의 요약 정보 객체 | 

## Methods

### NewBnsTgwV1ApiListTgwRoutesModelRouteResponseModel

`func NewBnsTgwV1ApiListTgwRoutesModelRouteResponseModel(resource BnsTgwV1ApiListTgwRoutesModelResourceResponseModel, ) *BnsTgwV1ApiListTgwRoutesModelRouteResponseModel`

NewBnsTgwV1ApiListTgwRoutesModelRouteResponseModel instantiates a new BnsTgwV1ApiListTgwRoutesModelRouteResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsTgwV1ApiListTgwRoutesModelRouteResponseModelWithDefaults

`func NewBnsTgwV1ApiListTgwRoutesModelRouteResponseModelWithDefaults() *BnsTgwV1ApiListTgwRoutesModelRouteResponseModel`

NewBnsTgwV1ApiListTgwRoutesModelRouteResponseModelWithDefaults instantiates a new BnsTgwV1ApiListTgwRoutesModelRouteResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsTgwV1ApiListTgwRoutesModelRouteResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsTgwV1ApiListTgwRoutesModelRouteResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsTgwV1ApiListTgwRoutesModelRouteResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BnsTgwV1ApiListTgwRoutesModelRouteResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *BnsTgwV1ApiListTgwRoutesModelRouteResponseModel) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *BnsTgwV1ApiListTgwRoutesModelRouteResponseModel) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetRouteType

`func (o *BnsTgwV1ApiListTgwRoutesModelRouteResponseModel) GetRouteType() string`

GetRouteType returns the RouteType field if non-nil, zero value otherwise.

### GetRouteTypeOk

`func (o *BnsTgwV1ApiListTgwRoutesModelRouteResponseModel) GetRouteTypeOk() (*string, bool)`

GetRouteTypeOk returns a tuple with the RouteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteType

`func (o *BnsTgwV1ApiListTgwRoutesModelRouteResponseModel) SetRouteType(v string)`

SetRouteType sets RouteType field to given value.

### HasRouteType

`func (o *BnsTgwV1ApiListTgwRoutesModelRouteResponseModel) HasRouteType() bool`

HasRouteType returns a boolean if a field has been set.

### SetRouteTypeNil

`func (o *BnsTgwV1ApiListTgwRoutesModelRouteResponseModel) SetRouteTypeNil(b bool)`

 SetRouteTypeNil sets the value for RouteType to be an explicit nil

### UnsetRouteType
`func (o *BnsTgwV1ApiListTgwRoutesModelRouteResponseModel) UnsetRouteType()`

UnsetRouteType ensures that no value is present for RouteType, not even an explicit nil
### GetDestinationCidrBlock

`func (o *BnsTgwV1ApiListTgwRoutesModelRouteResponseModel) GetDestinationCidrBlock() string`

GetDestinationCidrBlock returns the DestinationCidrBlock field if non-nil, zero value otherwise.

### GetDestinationCidrBlockOk

`func (o *BnsTgwV1ApiListTgwRoutesModelRouteResponseModel) GetDestinationCidrBlockOk() (*string, bool)`

GetDestinationCidrBlockOk returns a tuple with the DestinationCidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationCidrBlock

`func (o *BnsTgwV1ApiListTgwRoutesModelRouteResponseModel) SetDestinationCidrBlock(v string)`

SetDestinationCidrBlock sets DestinationCidrBlock field to given value.

### HasDestinationCidrBlock

`func (o *BnsTgwV1ApiListTgwRoutesModelRouteResponseModel) HasDestinationCidrBlock() bool`

HasDestinationCidrBlock returns a boolean if a field has been set.

### SetDestinationCidrBlockNil

`func (o *BnsTgwV1ApiListTgwRoutesModelRouteResponseModel) SetDestinationCidrBlockNil(b bool)`

 SetDestinationCidrBlockNil sets the value for DestinationCidrBlock to be an explicit nil

### UnsetDestinationCidrBlock
`func (o *BnsTgwV1ApiListTgwRoutesModelRouteResponseModel) UnsetDestinationCidrBlock()`

UnsetDestinationCidrBlock ensures that no value is present for DestinationCidrBlock, not even an explicit nil
### GetResourceAttachmentId

`func (o *BnsTgwV1ApiListTgwRoutesModelRouteResponseModel) GetResourceAttachmentId() string`

GetResourceAttachmentId returns the ResourceAttachmentId field if non-nil, zero value otherwise.

### GetResourceAttachmentIdOk

`func (o *BnsTgwV1ApiListTgwRoutesModelRouteResponseModel) GetResourceAttachmentIdOk() (*string, bool)`

GetResourceAttachmentIdOk returns a tuple with the ResourceAttachmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceAttachmentId

`func (o *BnsTgwV1ApiListTgwRoutesModelRouteResponseModel) SetResourceAttachmentId(v string)`

SetResourceAttachmentId sets ResourceAttachmentId field to given value.

### HasResourceAttachmentId

`func (o *BnsTgwV1ApiListTgwRoutesModelRouteResponseModel) HasResourceAttachmentId() bool`

HasResourceAttachmentId returns a boolean if a field has been set.

### SetResourceAttachmentIdNil

`func (o *BnsTgwV1ApiListTgwRoutesModelRouteResponseModel) SetResourceAttachmentIdNil(b bool)`

 SetResourceAttachmentIdNil sets the value for ResourceAttachmentId to be an explicit nil

### UnsetResourceAttachmentId
`func (o *BnsTgwV1ApiListTgwRoutesModelRouteResponseModel) UnsetResourceAttachmentId()`

UnsetResourceAttachmentId ensures that no value is present for ResourceAttachmentId, not even an explicit nil
### GetResourceId

`func (o *BnsTgwV1ApiListTgwRoutesModelRouteResponseModel) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *BnsTgwV1ApiListTgwRoutesModelRouteResponseModel) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *BnsTgwV1ApiListTgwRoutesModelRouteResponseModel) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *BnsTgwV1ApiListTgwRoutesModelRouteResponseModel) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### SetResourceIdNil

`func (o *BnsTgwV1ApiListTgwRoutesModelRouteResponseModel) SetResourceIdNil(b bool)`

 SetResourceIdNil sets the value for ResourceId to be an explicit nil

### UnsetResourceId
`func (o *BnsTgwV1ApiListTgwRoutesModelRouteResponseModel) UnsetResourceId()`

UnsetResourceId ensures that no value is present for ResourceId, not even an explicit nil
### GetResourceType

`func (o *BnsTgwV1ApiListTgwRoutesModelRouteResponseModel) GetResourceType() ResourceType`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *BnsTgwV1ApiListTgwRoutesModelRouteResponseModel) GetResourceTypeOk() (*ResourceType, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *BnsTgwV1ApiListTgwRoutesModelRouteResponseModel) SetResourceType(v ResourceType)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *BnsTgwV1ApiListTgwRoutesModelRouteResponseModel) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### SetResourceTypeNil

`func (o *BnsTgwV1ApiListTgwRoutesModelRouteResponseModel) SetResourceTypeNil(b bool)`

 SetResourceTypeNil sets the value for ResourceType to be an explicit nil

### UnsetResourceType
`func (o *BnsTgwV1ApiListTgwRoutesModelRouteResponseModel) UnsetResourceType()`

UnsetResourceType ensures that no value is present for ResourceType, not even an explicit nil
### GetTgwRouteTableId

`func (o *BnsTgwV1ApiListTgwRoutesModelRouteResponseModel) GetTgwRouteTableId() string`

GetTgwRouteTableId returns the TgwRouteTableId field if non-nil, zero value otherwise.

### GetTgwRouteTableIdOk

`func (o *BnsTgwV1ApiListTgwRoutesModelRouteResponseModel) GetTgwRouteTableIdOk() (*string, bool)`

GetTgwRouteTableIdOk returns a tuple with the TgwRouteTableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgwRouteTableId

`func (o *BnsTgwV1ApiListTgwRoutesModelRouteResponseModel) SetTgwRouteTableId(v string)`

SetTgwRouteTableId sets TgwRouteTableId field to given value.

### HasTgwRouteTableId

`func (o *BnsTgwV1ApiListTgwRoutesModelRouteResponseModel) HasTgwRouteTableId() bool`

HasTgwRouteTableId returns a boolean if a field has been set.

### SetTgwRouteTableIdNil

`func (o *BnsTgwV1ApiListTgwRoutesModelRouteResponseModel) SetTgwRouteTableIdNil(b bool)`

 SetTgwRouteTableIdNil sets the value for TgwRouteTableId to be an explicit nil

### UnsetTgwRouteTableId
`func (o *BnsTgwV1ApiListTgwRoutesModelRouteResponseModel) UnsetTgwRouteTableId()`

UnsetTgwRouteTableId ensures that no value is present for TgwRouteTableId, not even an explicit nil
### GetProvisioningStatus

`func (o *BnsTgwV1ApiListTgwRoutesModelRouteResponseModel) GetProvisioningStatus() TGWProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *BnsTgwV1ApiListTgwRoutesModelRouteResponseModel) GetProvisioningStatusOk() (*TGWProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *BnsTgwV1ApiListTgwRoutesModelRouteResponseModel) SetProvisioningStatus(v TGWProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.

### HasProvisioningStatus

`func (o *BnsTgwV1ApiListTgwRoutesModelRouteResponseModel) HasProvisioningStatus() bool`

HasProvisioningStatus returns a boolean if a field has been set.

### SetProvisioningStatusNil

`func (o *BnsTgwV1ApiListTgwRoutesModelRouteResponseModel) SetProvisioningStatusNil(b bool)`

 SetProvisioningStatusNil sets the value for ProvisioningStatus to be an explicit nil

### UnsetProvisioningStatus
`func (o *BnsTgwV1ApiListTgwRoutesModelRouteResponseModel) UnsetProvisioningStatus()`

UnsetProvisioningStatus ensures that no value is present for ProvisioningStatus, not even an explicit nil
### GetResource

`func (o *BnsTgwV1ApiListTgwRoutesModelRouteResponseModel) GetResource() BnsTgwV1ApiListTgwRoutesModelResourceResponseModel`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *BnsTgwV1ApiListTgwRoutesModelRouteResponseModel) GetResourceOk() (*BnsTgwV1ApiListTgwRoutesModelResourceResponseModel, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *BnsTgwV1ApiListTgwRoutesModelRouteResponseModel) SetResource(v BnsTgwV1ApiListTgwRoutesModelResourceResponseModel)`

SetResource sets Resource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


