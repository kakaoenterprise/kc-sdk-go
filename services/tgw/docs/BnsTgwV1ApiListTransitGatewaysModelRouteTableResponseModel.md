# BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Region** | Pointer to [**NullableRegion**](Region.md) |  | [optional] 
**ProjectId** | Pointer to **NullableString** |  | [optional] 
**TgwId** | Pointer to **NullableString** |  | [optional] 
**IsDefaultAssociationRouteTable** | Pointer to **NullableBool** |  | [optional] 
**IsDefaultPropagationRouteTable** | Pointer to **NullableBool** |  | [optional] 
**ProvisioningStatus** | Pointer to [**NullableTGWProvisioningStatus**](TGWProvisioningStatus.md) |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewBnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel

`func NewBnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel() *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel`

NewBnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel instantiates a new BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModelWithDefaults

`func NewBnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModelWithDefaults() *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel`

NewBnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModelWithDefaults instantiates a new BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRegion

`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) GetRegion() Region`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) GetRegionOk() (*Region, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) SetRegion(v Region)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetProjectId

`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetTgwId

`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) GetTgwId() string`

GetTgwId returns the TgwId field if non-nil, zero value otherwise.

### GetTgwIdOk

`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) GetTgwIdOk() (*string, bool)`

GetTgwIdOk returns a tuple with the TgwId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgwId

`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) SetTgwId(v string)`

SetTgwId sets TgwId field to given value.

### HasTgwId

`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) HasTgwId() bool`

HasTgwId returns a boolean if a field has been set.

### SetTgwIdNil

`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) SetTgwIdNil(b bool)`

 SetTgwIdNil sets the value for TgwId to be an explicit nil

### UnsetTgwId
`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) UnsetTgwId()`

UnsetTgwId ensures that no value is present for TgwId, not even an explicit nil
### GetIsDefaultAssociationRouteTable

`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) GetIsDefaultAssociationRouteTable() bool`

GetIsDefaultAssociationRouteTable returns the IsDefaultAssociationRouteTable field if non-nil, zero value otherwise.

### GetIsDefaultAssociationRouteTableOk

`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) GetIsDefaultAssociationRouteTableOk() (*bool, bool)`

GetIsDefaultAssociationRouteTableOk returns a tuple with the IsDefaultAssociationRouteTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultAssociationRouteTable

`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) SetIsDefaultAssociationRouteTable(v bool)`

SetIsDefaultAssociationRouteTable sets IsDefaultAssociationRouteTable field to given value.

### HasIsDefaultAssociationRouteTable

`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) HasIsDefaultAssociationRouteTable() bool`

HasIsDefaultAssociationRouteTable returns a boolean if a field has been set.

### SetIsDefaultAssociationRouteTableNil

`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) SetIsDefaultAssociationRouteTableNil(b bool)`

 SetIsDefaultAssociationRouteTableNil sets the value for IsDefaultAssociationRouteTable to be an explicit nil

### UnsetIsDefaultAssociationRouteTable
`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) UnsetIsDefaultAssociationRouteTable()`

UnsetIsDefaultAssociationRouteTable ensures that no value is present for IsDefaultAssociationRouteTable, not even an explicit nil
### GetIsDefaultPropagationRouteTable

`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) GetIsDefaultPropagationRouteTable() bool`

GetIsDefaultPropagationRouteTable returns the IsDefaultPropagationRouteTable field if non-nil, zero value otherwise.

### GetIsDefaultPropagationRouteTableOk

`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) GetIsDefaultPropagationRouteTableOk() (*bool, bool)`

GetIsDefaultPropagationRouteTableOk returns a tuple with the IsDefaultPropagationRouteTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultPropagationRouteTable

`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) SetIsDefaultPropagationRouteTable(v bool)`

SetIsDefaultPropagationRouteTable sets IsDefaultPropagationRouteTable field to given value.

### HasIsDefaultPropagationRouteTable

`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) HasIsDefaultPropagationRouteTable() bool`

HasIsDefaultPropagationRouteTable returns a boolean if a field has been set.

### SetIsDefaultPropagationRouteTableNil

`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) SetIsDefaultPropagationRouteTableNil(b bool)`

 SetIsDefaultPropagationRouteTableNil sets the value for IsDefaultPropagationRouteTable to be an explicit nil

### UnsetIsDefaultPropagationRouteTable
`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) UnsetIsDefaultPropagationRouteTable()`

UnsetIsDefaultPropagationRouteTable ensures that no value is present for IsDefaultPropagationRouteTable, not even an explicit nil
### GetProvisioningStatus

`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) GetProvisioningStatus() TGWProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) GetProvisioningStatusOk() (*TGWProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) SetProvisioningStatus(v TGWProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.

### HasProvisioningStatus

`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) HasProvisioningStatus() bool`

HasProvisioningStatus returns a boolean if a field has been set.

### SetProvisioningStatusNil

`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) SetProvisioningStatusNil(b bool)`

 SetProvisioningStatusNil sets the value for ProvisioningStatus to be an explicit nil

### UnsetProvisioningStatus
`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) UnsetProvisioningStatus()`

UnsetProvisioningStatus ensures that no value is present for ProvisioningStatus, not even an explicit nil
### GetCreatedAt

`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *BnsTgwV1ApiListTransitGatewaysModelRouteTableResponseModel) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


