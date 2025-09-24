# BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 로드 밸런서 ID | 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to [**NullableLoadBalancerType**](LoadBalancerType.md) |  | [optional] 
**ProvisioningStatus** | Pointer to [**NullableProvisioningStatus**](ProvisioningStatus.md) |  | [optional] 
**OperatingStatus** | Pointer to [**NullableLoadBalancerOperatingStatus**](LoadBalancerOperatingStatus.md) |  | [optional] 
**AvailabilityZone** | Pointer to [**NullableAvailabilityZone**](AvailabilityZone.md) |  | [optional] 
**TypeId** | Pointer to **NullableString** |  | [optional] 
**SubnetId** | Pointer to **NullableString** |  | [optional] 
**SubnetName** | Pointer to **NullableString** |  | [optional] 
**SubnetCidrBlock** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewBnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel

`func NewBnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel(id string, ) *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel`

NewBnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel instantiates a new BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModelWithDefaults

`func NewBnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModelWithDefaults() *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel`

NewBnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModelWithDefaults instantiates a new BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetType

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) GetType() LoadBalancerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) GetTypeOk() (*LoadBalancerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) SetType(v LoadBalancerType)`

SetType sets Type field to given value.

### HasType

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetProvisioningStatus

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) GetProvisioningStatus() ProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) GetProvisioningStatusOk() (*ProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) SetProvisioningStatus(v ProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.

### HasProvisioningStatus

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) HasProvisioningStatus() bool`

HasProvisioningStatus returns a boolean if a field has been set.

### SetProvisioningStatusNil

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) SetProvisioningStatusNil(b bool)`

 SetProvisioningStatusNil sets the value for ProvisioningStatus to be an explicit nil

### UnsetProvisioningStatus
`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) UnsetProvisioningStatus()`

UnsetProvisioningStatus ensures that no value is present for ProvisioningStatus, not even an explicit nil
### GetOperatingStatus

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) GetOperatingStatus() LoadBalancerOperatingStatus`

GetOperatingStatus returns the OperatingStatus field if non-nil, zero value otherwise.

### GetOperatingStatusOk

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) GetOperatingStatusOk() (*LoadBalancerOperatingStatus, bool)`

GetOperatingStatusOk returns a tuple with the OperatingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingStatus

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) SetOperatingStatus(v LoadBalancerOperatingStatus)`

SetOperatingStatus sets OperatingStatus field to given value.

### HasOperatingStatus

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) HasOperatingStatus() bool`

HasOperatingStatus returns a boolean if a field has been set.

### SetOperatingStatusNil

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) SetOperatingStatusNil(b bool)`

 SetOperatingStatusNil sets the value for OperatingStatus to be an explicit nil

### UnsetOperatingStatus
`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) UnsetOperatingStatus()`

UnsetOperatingStatus ensures that no value is present for OperatingStatus, not even an explicit nil
### GetAvailabilityZone

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) GetAvailabilityZone() AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) GetAvailabilityZoneOk() (*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) SetAvailabilityZone(v AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### SetAvailabilityZoneNil

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) SetAvailabilityZoneNil(b bool)`

 SetAvailabilityZoneNil sets the value for AvailabilityZone to be an explicit nil

### UnsetAvailabilityZone
`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) UnsetAvailabilityZone()`

UnsetAvailabilityZone ensures that no value is present for AvailabilityZone, not even an explicit nil
### GetTypeId

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) GetTypeId() string`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) GetTypeIdOk() (*string, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) SetTypeId(v string)`

SetTypeId sets TypeId field to given value.

### HasTypeId

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) HasTypeId() bool`

HasTypeId returns a boolean if a field has been set.

### SetTypeIdNil

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) SetTypeIdNil(b bool)`

 SetTypeIdNil sets the value for TypeId to be an explicit nil

### UnsetTypeId
`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) UnsetTypeId()`

UnsetTypeId ensures that no value is present for TypeId, not even an explicit nil
### GetSubnetId

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.

### HasSubnetId

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.

### SetSubnetIdNil

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) SetSubnetIdNil(b bool)`

 SetSubnetIdNil sets the value for SubnetId to be an explicit nil

### UnsetSubnetId
`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) UnsetSubnetId()`

UnsetSubnetId ensures that no value is present for SubnetId, not even an explicit nil
### GetSubnetName

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) GetSubnetName() string`

GetSubnetName returns the SubnetName field if non-nil, zero value otherwise.

### GetSubnetNameOk

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) GetSubnetNameOk() (*string, bool)`

GetSubnetNameOk returns a tuple with the SubnetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetName

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) SetSubnetName(v string)`

SetSubnetName sets SubnetName field to given value.

### HasSubnetName

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) HasSubnetName() bool`

HasSubnetName returns a boolean if a field has been set.

### SetSubnetNameNil

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) SetSubnetNameNil(b bool)`

 SetSubnetNameNil sets the value for SubnetName to be an explicit nil

### UnsetSubnetName
`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) UnsetSubnetName()`

UnsetSubnetName ensures that no value is present for SubnetName, not even an explicit nil
### GetSubnetCidrBlock

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) GetSubnetCidrBlock() string`

GetSubnetCidrBlock returns the SubnetCidrBlock field if non-nil, zero value otherwise.

### GetSubnetCidrBlockOk

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) GetSubnetCidrBlockOk() (*string, bool)`

GetSubnetCidrBlockOk returns a tuple with the SubnetCidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetCidrBlock

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) SetSubnetCidrBlock(v string)`

SetSubnetCidrBlock sets SubnetCidrBlock field to given value.

### HasSubnetCidrBlock

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) HasSubnetCidrBlock() bool`

HasSubnetCidrBlock returns a boolean if a field has been set.

### SetSubnetCidrBlockNil

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) SetSubnetCidrBlockNil(b bool)`

 SetSubnetCidrBlockNil sets the value for SubnetCidrBlock to be an explicit nil

### UnsetSubnetCidrBlock
`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) UnsetSubnetCidrBlock()`

UnsetSubnetCidrBlock ensures that no value is present for SubnetCidrBlock, not even an explicit nil
### GetCreatedAt

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


