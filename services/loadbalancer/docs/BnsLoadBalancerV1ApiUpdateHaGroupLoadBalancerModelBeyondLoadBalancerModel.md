# BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 고가용성 로드 밸런서의 ID | 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ProvisioningStatus** | Pointer to [**NullableProvisioningStatus**](ProvisioningStatus.md) |  | [optional] 
**OperatingStatus** | Pointer to [**NullableLoadBalancerOperatingStatus**](LoadBalancerOperatingStatus.md) |  | [optional] 
**ProjectId** | Pointer to **NullableString** |  | [optional] 
**DnsName** | Pointer to **NullableString** |  | [optional] 
**VpcId** | Pointer to **NullableString** |  | [optional] 
**Scheme** | Pointer to [**NullableScheme**](Scheme.md) |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 
**TypeId** | Pointer to **NullableString** |  | [optional] 
**Subnets** | Pointer to [**[]BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelSubnetModel**](BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelSubnetModel.md) |  | [optional] 

## Methods

### NewBnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel

`func NewBnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel(id string, ) *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel`

NewBnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel instantiates a new BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModelWithDefaults

`func NewBnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModelWithDefaults() *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel`

NewBnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModelWithDefaults instantiates a new BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetProvisioningStatus

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) GetProvisioningStatus() ProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) GetProvisioningStatusOk() (*ProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) SetProvisioningStatus(v ProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.

### HasProvisioningStatus

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) HasProvisioningStatus() bool`

HasProvisioningStatus returns a boolean if a field has been set.

### SetProvisioningStatusNil

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) SetProvisioningStatusNil(b bool)`

 SetProvisioningStatusNil sets the value for ProvisioningStatus to be an explicit nil

### UnsetProvisioningStatus
`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) UnsetProvisioningStatus()`

UnsetProvisioningStatus ensures that no value is present for ProvisioningStatus, not even an explicit nil
### GetOperatingStatus

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) GetOperatingStatus() LoadBalancerOperatingStatus`

GetOperatingStatus returns the OperatingStatus field if non-nil, zero value otherwise.

### GetOperatingStatusOk

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) GetOperatingStatusOk() (*LoadBalancerOperatingStatus, bool)`

GetOperatingStatusOk returns a tuple with the OperatingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingStatus

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) SetOperatingStatus(v LoadBalancerOperatingStatus)`

SetOperatingStatus sets OperatingStatus field to given value.

### HasOperatingStatus

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) HasOperatingStatus() bool`

HasOperatingStatus returns a boolean if a field has been set.

### SetOperatingStatusNil

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) SetOperatingStatusNil(b bool)`

 SetOperatingStatusNil sets the value for OperatingStatus to be an explicit nil

### UnsetOperatingStatus
`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) UnsetOperatingStatus()`

UnsetOperatingStatus ensures that no value is present for OperatingStatus, not even an explicit nil
### GetProjectId

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetDnsName

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### SetDnsNameNil

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) SetDnsNameNil(b bool)`

 SetDnsNameNil sets the value for DnsName to be an explicit nil

### UnsetDnsName
`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) UnsetDnsName()`

UnsetDnsName ensures that no value is present for DnsName, not even an explicit nil
### GetVpcId

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### SetVpcIdNil

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) SetVpcIdNil(b bool)`

 SetVpcIdNil sets the value for VpcId to be an explicit nil

### UnsetVpcId
`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) UnsetVpcId()`

UnsetVpcId ensures that no value is present for VpcId, not even an explicit nil
### GetScheme

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) GetScheme() Scheme`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) GetSchemeOk() (*Scheme, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) SetScheme(v Scheme)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### SetSchemeNil

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) SetSchemeNil(b bool)`

 SetSchemeNil sets the value for Scheme to be an explicit nil

### UnsetScheme
`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) UnsetScheme()`

UnsetScheme ensures that no value is present for Scheme, not even an explicit nil
### GetCreatedAt

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetTypeId

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) GetTypeId() string`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) GetTypeIdOk() (*string, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) SetTypeId(v string)`

SetTypeId sets TypeId field to given value.

### HasTypeId

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) HasTypeId() bool`

HasTypeId returns a boolean if a field has been set.

### SetTypeIdNil

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) SetTypeIdNil(b bool)`

 SetTypeIdNil sets the value for TypeId to be an explicit nil

### UnsetTypeId
`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) UnsetTypeId()`

UnsetTypeId ensures that no value is present for TypeId, not even an explicit nil
### GetSubnets

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) GetSubnets() []BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelSubnetModel`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) GetSubnetsOk() (*[]BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelSubnetModel, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) SetSubnets(v []BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelSubnetModel)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.

### SetSubnetsNil

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) SetSubnetsNil(b bool)`

 SetSubnetsNil sets the value for Subnets to be an explicit nil

### UnsetSubnets
`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelBeyondLoadBalancerModel) UnsetSubnets()`

UnsetSubnets ensures that no value is present for Subnets, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


