# BnsVpcV1ApiListRouteTableAssociationsModelAssociationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 라우팅 테이블 연결 리소스의 ID | 
**ProvisioningStatus** | Pointer to [**NullableProvisioningStatus**](ProvisioningStatus.md) |  | [optional] 
**VpcId** | Pointer to **NullableString** |  | [optional] 
**VpcName** | Pointer to **NullableString** |  | [optional] 
**SubnetId** | Pointer to **NullableString** |  | [optional] 
**SubnetName** | Pointer to **NullableString** |  | [optional] 
**SubnetCidrBlock** | Pointer to **NullableString** |  | [optional] 
**AvailabilityZone** | Pointer to [**NullableAvailabilityZone**](AvailabilityZone.md) |  | [optional] 

## Methods

### NewBnsVpcV1ApiListRouteTableAssociationsModelAssociationModel

`func NewBnsVpcV1ApiListRouteTableAssociationsModelAssociationModel(id string, ) *BnsVpcV1ApiListRouteTableAssociationsModelAssociationModel`

NewBnsVpcV1ApiListRouteTableAssociationsModelAssociationModel instantiates a new BnsVpcV1ApiListRouteTableAssociationsModelAssociationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsVpcV1ApiListRouteTableAssociationsModelAssociationModelWithDefaults

`func NewBnsVpcV1ApiListRouteTableAssociationsModelAssociationModelWithDefaults() *BnsVpcV1ApiListRouteTableAssociationsModelAssociationModel`

NewBnsVpcV1ApiListRouteTableAssociationsModelAssociationModelWithDefaults instantiates a new BnsVpcV1ApiListRouteTableAssociationsModelAssociationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsVpcV1ApiListRouteTableAssociationsModelAssociationModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsVpcV1ApiListRouteTableAssociationsModelAssociationModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsVpcV1ApiListRouteTableAssociationsModelAssociationModel) SetId(v string)`

SetId sets Id field to given value.


### GetProvisioningStatus

`func (o *BnsVpcV1ApiListRouteTableAssociationsModelAssociationModel) GetProvisioningStatus() ProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *BnsVpcV1ApiListRouteTableAssociationsModelAssociationModel) GetProvisioningStatusOk() (*ProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *BnsVpcV1ApiListRouteTableAssociationsModelAssociationModel) SetProvisioningStatus(v ProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.

### HasProvisioningStatus

`func (o *BnsVpcV1ApiListRouteTableAssociationsModelAssociationModel) HasProvisioningStatus() bool`

HasProvisioningStatus returns a boolean if a field has been set.

### SetProvisioningStatusNil

`func (o *BnsVpcV1ApiListRouteTableAssociationsModelAssociationModel) SetProvisioningStatusNil(b bool)`

 SetProvisioningStatusNil sets the value for ProvisioningStatus to be an explicit nil

### UnsetProvisioningStatus
`func (o *BnsVpcV1ApiListRouteTableAssociationsModelAssociationModel) UnsetProvisioningStatus()`

UnsetProvisioningStatus ensures that no value is present for ProvisioningStatus, not even an explicit nil
### GetVpcId

`func (o *BnsVpcV1ApiListRouteTableAssociationsModelAssociationModel) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *BnsVpcV1ApiListRouteTableAssociationsModelAssociationModel) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *BnsVpcV1ApiListRouteTableAssociationsModelAssociationModel) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *BnsVpcV1ApiListRouteTableAssociationsModelAssociationModel) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### SetVpcIdNil

`func (o *BnsVpcV1ApiListRouteTableAssociationsModelAssociationModel) SetVpcIdNil(b bool)`

 SetVpcIdNil sets the value for VpcId to be an explicit nil

### UnsetVpcId
`func (o *BnsVpcV1ApiListRouteTableAssociationsModelAssociationModel) UnsetVpcId()`

UnsetVpcId ensures that no value is present for VpcId, not even an explicit nil
### GetVpcName

`func (o *BnsVpcV1ApiListRouteTableAssociationsModelAssociationModel) GetVpcName() string`

GetVpcName returns the VpcName field if non-nil, zero value otherwise.

### GetVpcNameOk

`func (o *BnsVpcV1ApiListRouteTableAssociationsModelAssociationModel) GetVpcNameOk() (*string, bool)`

GetVpcNameOk returns a tuple with the VpcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcName

`func (o *BnsVpcV1ApiListRouteTableAssociationsModelAssociationModel) SetVpcName(v string)`

SetVpcName sets VpcName field to given value.

### HasVpcName

`func (o *BnsVpcV1ApiListRouteTableAssociationsModelAssociationModel) HasVpcName() bool`

HasVpcName returns a boolean if a field has been set.

### SetVpcNameNil

`func (o *BnsVpcV1ApiListRouteTableAssociationsModelAssociationModel) SetVpcNameNil(b bool)`

 SetVpcNameNil sets the value for VpcName to be an explicit nil

### UnsetVpcName
`func (o *BnsVpcV1ApiListRouteTableAssociationsModelAssociationModel) UnsetVpcName()`

UnsetVpcName ensures that no value is present for VpcName, not even an explicit nil
### GetSubnetId

`func (o *BnsVpcV1ApiListRouteTableAssociationsModelAssociationModel) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *BnsVpcV1ApiListRouteTableAssociationsModelAssociationModel) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *BnsVpcV1ApiListRouteTableAssociationsModelAssociationModel) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.

### HasSubnetId

`func (o *BnsVpcV1ApiListRouteTableAssociationsModelAssociationModel) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.

### SetSubnetIdNil

`func (o *BnsVpcV1ApiListRouteTableAssociationsModelAssociationModel) SetSubnetIdNil(b bool)`

 SetSubnetIdNil sets the value for SubnetId to be an explicit nil

### UnsetSubnetId
`func (o *BnsVpcV1ApiListRouteTableAssociationsModelAssociationModel) UnsetSubnetId()`

UnsetSubnetId ensures that no value is present for SubnetId, not even an explicit nil
### GetSubnetName

`func (o *BnsVpcV1ApiListRouteTableAssociationsModelAssociationModel) GetSubnetName() string`

GetSubnetName returns the SubnetName field if non-nil, zero value otherwise.

### GetSubnetNameOk

`func (o *BnsVpcV1ApiListRouteTableAssociationsModelAssociationModel) GetSubnetNameOk() (*string, bool)`

GetSubnetNameOk returns a tuple with the SubnetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetName

`func (o *BnsVpcV1ApiListRouteTableAssociationsModelAssociationModel) SetSubnetName(v string)`

SetSubnetName sets SubnetName field to given value.

### HasSubnetName

`func (o *BnsVpcV1ApiListRouteTableAssociationsModelAssociationModel) HasSubnetName() bool`

HasSubnetName returns a boolean if a field has been set.

### SetSubnetNameNil

`func (o *BnsVpcV1ApiListRouteTableAssociationsModelAssociationModel) SetSubnetNameNil(b bool)`

 SetSubnetNameNil sets the value for SubnetName to be an explicit nil

### UnsetSubnetName
`func (o *BnsVpcV1ApiListRouteTableAssociationsModelAssociationModel) UnsetSubnetName()`

UnsetSubnetName ensures that no value is present for SubnetName, not even an explicit nil
### GetSubnetCidrBlock

`func (o *BnsVpcV1ApiListRouteTableAssociationsModelAssociationModel) GetSubnetCidrBlock() string`

GetSubnetCidrBlock returns the SubnetCidrBlock field if non-nil, zero value otherwise.

### GetSubnetCidrBlockOk

`func (o *BnsVpcV1ApiListRouteTableAssociationsModelAssociationModel) GetSubnetCidrBlockOk() (*string, bool)`

GetSubnetCidrBlockOk returns a tuple with the SubnetCidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetCidrBlock

`func (o *BnsVpcV1ApiListRouteTableAssociationsModelAssociationModel) SetSubnetCidrBlock(v string)`

SetSubnetCidrBlock sets SubnetCidrBlock field to given value.

### HasSubnetCidrBlock

`func (o *BnsVpcV1ApiListRouteTableAssociationsModelAssociationModel) HasSubnetCidrBlock() bool`

HasSubnetCidrBlock returns a boolean if a field has been set.

### SetSubnetCidrBlockNil

`func (o *BnsVpcV1ApiListRouteTableAssociationsModelAssociationModel) SetSubnetCidrBlockNil(b bool)`

 SetSubnetCidrBlockNil sets the value for SubnetCidrBlock to be an explicit nil

### UnsetSubnetCidrBlock
`func (o *BnsVpcV1ApiListRouteTableAssociationsModelAssociationModel) UnsetSubnetCidrBlock()`

UnsetSubnetCidrBlock ensures that no value is present for SubnetCidrBlock, not even an explicit nil
### GetAvailabilityZone

`func (o *BnsVpcV1ApiListRouteTableAssociationsModelAssociationModel) GetAvailabilityZone() AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *BnsVpcV1ApiListRouteTableAssociationsModelAssociationModel) GetAvailabilityZoneOk() (*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *BnsVpcV1ApiListRouteTableAssociationsModelAssociationModel) SetAvailabilityZone(v AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *BnsVpcV1ApiListRouteTableAssociationsModelAssociationModel) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### SetAvailabilityZoneNil

`func (o *BnsVpcV1ApiListRouteTableAssociationsModelAssociationModel) SetAvailabilityZoneNil(b bool)`

 SetAvailabilityZoneNil sets the value for AvailabilityZone to be an explicit nil

### UnsetAvailabilityZone
`func (o *BnsVpcV1ApiListRouteTableAssociationsModelAssociationModel) UnsetAvailabilityZone()`

UnsetAvailabilityZone ensures that no value is present for AvailabilityZone, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


