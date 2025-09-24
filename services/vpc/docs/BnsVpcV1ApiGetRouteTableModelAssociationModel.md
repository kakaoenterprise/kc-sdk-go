# BnsVpcV1ApiGetRouteTableModelAssociationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 연결 관계 ID | 
**ProvisioningStatus** | Pointer to [**NullableProvisioningStatus**](ProvisioningStatus.md) |  | [optional] 
**VpcId** | Pointer to **NullableString** |  | [optional] 
**VpcName** | Pointer to **NullableString** |  | [optional] 
**SubnetId** | Pointer to **NullableString** |  | [optional] 
**SubnetName** | Pointer to **NullableString** |  | [optional] 
**SubnetCidrBlock** | Pointer to **NullableString** |  | [optional] 
**AvailabilityZone** | Pointer to [**NullableAvailabilityZone**](AvailabilityZone.md) |  | [optional] 

## Methods

### NewBnsVpcV1ApiGetRouteTableModelAssociationModel

`func NewBnsVpcV1ApiGetRouteTableModelAssociationModel(id string, ) *BnsVpcV1ApiGetRouteTableModelAssociationModel`

NewBnsVpcV1ApiGetRouteTableModelAssociationModel instantiates a new BnsVpcV1ApiGetRouteTableModelAssociationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsVpcV1ApiGetRouteTableModelAssociationModelWithDefaults

`func NewBnsVpcV1ApiGetRouteTableModelAssociationModelWithDefaults() *BnsVpcV1ApiGetRouteTableModelAssociationModel`

NewBnsVpcV1ApiGetRouteTableModelAssociationModelWithDefaults instantiates a new BnsVpcV1ApiGetRouteTableModelAssociationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsVpcV1ApiGetRouteTableModelAssociationModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsVpcV1ApiGetRouteTableModelAssociationModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsVpcV1ApiGetRouteTableModelAssociationModel) SetId(v string)`

SetId sets Id field to given value.


### GetProvisioningStatus

`func (o *BnsVpcV1ApiGetRouteTableModelAssociationModel) GetProvisioningStatus() ProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *BnsVpcV1ApiGetRouteTableModelAssociationModel) GetProvisioningStatusOk() (*ProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *BnsVpcV1ApiGetRouteTableModelAssociationModel) SetProvisioningStatus(v ProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.

### HasProvisioningStatus

`func (o *BnsVpcV1ApiGetRouteTableModelAssociationModel) HasProvisioningStatus() bool`

HasProvisioningStatus returns a boolean if a field has been set.

### SetProvisioningStatusNil

`func (o *BnsVpcV1ApiGetRouteTableModelAssociationModel) SetProvisioningStatusNil(b bool)`

 SetProvisioningStatusNil sets the value for ProvisioningStatus to be an explicit nil

### UnsetProvisioningStatus
`func (o *BnsVpcV1ApiGetRouteTableModelAssociationModel) UnsetProvisioningStatus()`

UnsetProvisioningStatus ensures that no value is present for ProvisioningStatus, not even an explicit nil
### GetVpcId

`func (o *BnsVpcV1ApiGetRouteTableModelAssociationModel) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *BnsVpcV1ApiGetRouteTableModelAssociationModel) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *BnsVpcV1ApiGetRouteTableModelAssociationModel) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *BnsVpcV1ApiGetRouteTableModelAssociationModel) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### SetVpcIdNil

`func (o *BnsVpcV1ApiGetRouteTableModelAssociationModel) SetVpcIdNil(b bool)`

 SetVpcIdNil sets the value for VpcId to be an explicit nil

### UnsetVpcId
`func (o *BnsVpcV1ApiGetRouteTableModelAssociationModel) UnsetVpcId()`

UnsetVpcId ensures that no value is present for VpcId, not even an explicit nil
### GetVpcName

`func (o *BnsVpcV1ApiGetRouteTableModelAssociationModel) GetVpcName() string`

GetVpcName returns the VpcName field if non-nil, zero value otherwise.

### GetVpcNameOk

`func (o *BnsVpcV1ApiGetRouteTableModelAssociationModel) GetVpcNameOk() (*string, bool)`

GetVpcNameOk returns a tuple with the VpcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcName

`func (o *BnsVpcV1ApiGetRouteTableModelAssociationModel) SetVpcName(v string)`

SetVpcName sets VpcName field to given value.

### HasVpcName

`func (o *BnsVpcV1ApiGetRouteTableModelAssociationModel) HasVpcName() bool`

HasVpcName returns a boolean if a field has been set.

### SetVpcNameNil

`func (o *BnsVpcV1ApiGetRouteTableModelAssociationModel) SetVpcNameNil(b bool)`

 SetVpcNameNil sets the value for VpcName to be an explicit nil

### UnsetVpcName
`func (o *BnsVpcV1ApiGetRouteTableModelAssociationModel) UnsetVpcName()`

UnsetVpcName ensures that no value is present for VpcName, not even an explicit nil
### GetSubnetId

`func (o *BnsVpcV1ApiGetRouteTableModelAssociationModel) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *BnsVpcV1ApiGetRouteTableModelAssociationModel) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *BnsVpcV1ApiGetRouteTableModelAssociationModel) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.

### HasSubnetId

`func (o *BnsVpcV1ApiGetRouteTableModelAssociationModel) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.

### SetSubnetIdNil

`func (o *BnsVpcV1ApiGetRouteTableModelAssociationModel) SetSubnetIdNil(b bool)`

 SetSubnetIdNil sets the value for SubnetId to be an explicit nil

### UnsetSubnetId
`func (o *BnsVpcV1ApiGetRouteTableModelAssociationModel) UnsetSubnetId()`

UnsetSubnetId ensures that no value is present for SubnetId, not even an explicit nil
### GetSubnetName

`func (o *BnsVpcV1ApiGetRouteTableModelAssociationModel) GetSubnetName() string`

GetSubnetName returns the SubnetName field if non-nil, zero value otherwise.

### GetSubnetNameOk

`func (o *BnsVpcV1ApiGetRouteTableModelAssociationModel) GetSubnetNameOk() (*string, bool)`

GetSubnetNameOk returns a tuple with the SubnetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetName

`func (o *BnsVpcV1ApiGetRouteTableModelAssociationModel) SetSubnetName(v string)`

SetSubnetName sets SubnetName field to given value.

### HasSubnetName

`func (o *BnsVpcV1ApiGetRouteTableModelAssociationModel) HasSubnetName() bool`

HasSubnetName returns a boolean if a field has been set.

### SetSubnetNameNil

`func (o *BnsVpcV1ApiGetRouteTableModelAssociationModel) SetSubnetNameNil(b bool)`

 SetSubnetNameNil sets the value for SubnetName to be an explicit nil

### UnsetSubnetName
`func (o *BnsVpcV1ApiGetRouteTableModelAssociationModel) UnsetSubnetName()`

UnsetSubnetName ensures that no value is present for SubnetName, not even an explicit nil
### GetSubnetCidrBlock

`func (o *BnsVpcV1ApiGetRouteTableModelAssociationModel) GetSubnetCidrBlock() string`

GetSubnetCidrBlock returns the SubnetCidrBlock field if non-nil, zero value otherwise.

### GetSubnetCidrBlockOk

`func (o *BnsVpcV1ApiGetRouteTableModelAssociationModel) GetSubnetCidrBlockOk() (*string, bool)`

GetSubnetCidrBlockOk returns a tuple with the SubnetCidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetCidrBlock

`func (o *BnsVpcV1ApiGetRouteTableModelAssociationModel) SetSubnetCidrBlock(v string)`

SetSubnetCidrBlock sets SubnetCidrBlock field to given value.

### HasSubnetCidrBlock

`func (o *BnsVpcV1ApiGetRouteTableModelAssociationModel) HasSubnetCidrBlock() bool`

HasSubnetCidrBlock returns a boolean if a field has been set.

### SetSubnetCidrBlockNil

`func (o *BnsVpcV1ApiGetRouteTableModelAssociationModel) SetSubnetCidrBlockNil(b bool)`

 SetSubnetCidrBlockNil sets the value for SubnetCidrBlock to be an explicit nil

### UnsetSubnetCidrBlock
`func (o *BnsVpcV1ApiGetRouteTableModelAssociationModel) UnsetSubnetCidrBlock()`

UnsetSubnetCidrBlock ensures that no value is present for SubnetCidrBlock, not even an explicit nil
### GetAvailabilityZone

`func (o *BnsVpcV1ApiGetRouteTableModelAssociationModel) GetAvailabilityZone() AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *BnsVpcV1ApiGetRouteTableModelAssociationModel) GetAvailabilityZoneOk() (*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *BnsVpcV1ApiGetRouteTableModelAssociationModel) SetAvailabilityZone(v AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *BnsVpcV1ApiGetRouteTableModelAssociationModel) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### SetAvailabilityZoneNil

`func (o *BnsVpcV1ApiGetRouteTableModelAssociationModel) SetAvailabilityZoneNil(b bool)`

 SetAvailabilityZoneNil sets the value for AvailabilityZone to be an explicit nil

### UnsetAvailabilityZone
`func (o *BnsVpcV1ApiGetRouteTableModelAssociationModel) UnsetAvailabilityZone()`

UnsetAvailabilityZone ensures that no value is present for AvailabilityZone, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


