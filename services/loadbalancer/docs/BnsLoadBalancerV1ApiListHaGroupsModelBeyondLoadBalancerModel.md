# BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 고가용성 그룹 ID | 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Provider** | Pointer to **NullableString** |  | [optional] 
**Scheme** | Pointer to [**NullableScheme**](Scheme.md) |  | [optional] 
**ProjectId** | Pointer to **NullableString** |  | [optional] 
**DnsName** | Pointer to **NullableString** |  | [optional] 
**TypeId** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 
**ProvisioningStatus** | Pointer to [**NullableProvisioningStatus**](ProvisioningStatus.md) |  | [optional] 
**OperatingStatus** | Pointer to [**NullableLoadBalancerOperatingStatus**](LoadBalancerOperatingStatus.md) |  | [optional] 
**VpcId** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to [**NullableLoadBalancerType**](LoadBalancerType.md) |  | [optional] 
**VpcName** | Pointer to **NullableString** |  | [optional] 
**VpcCidrBlock** | Pointer to **NullableString** |  | [optional] 
**AvailabilityZones** | Pointer to [**[]AvailabilityZone**](AvailabilityZone.md) |  | [optional] 
**LoadBalancers** | Pointer to [**[]BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel**](BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel.md) |  | [optional] 

## Methods

### NewBnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel

`func NewBnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel(id string, ) *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel`

NewBnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel instantiates a new BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModelWithDefaults

`func NewBnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModelWithDefaults() *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel`

NewBnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModelWithDefaults instantiates a new BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetProvider

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### SetProviderNil

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) SetProviderNil(b bool)`

 SetProviderNil sets the value for Provider to be an explicit nil

### UnsetProvider
`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) UnsetProvider()`

UnsetProvider ensures that no value is present for Provider, not even an explicit nil
### GetScheme

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) GetScheme() Scheme`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) GetSchemeOk() (*Scheme, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) SetScheme(v Scheme)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### SetSchemeNil

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) SetSchemeNil(b bool)`

 SetSchemeNil sets the value for Scheme to be an explicit nil

### UnsetScheme
`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) UnsetScheme()`

UnsetScheme ensures that no value is present for Scheme, not even an explicit nil
### GetProjectId

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetDnsName

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### SetDnsNameNil

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) SetDnsNameNil(b bool)`

 SetDnsNameNil sets the value for DnsName to be an explicit nil

### UnsetDnsName
`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) UnsetDnsName()`

UnsetDnsName ensures that no value is present for DnsName, not even an explicit nil
### GetTypeId

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) GetTypeId() string`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) GetTypeIdOk() (*string, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) SetTypeId(v string)`

SetTypeId sets TypeId field to given value.

### HasTypeId

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) HasTypeId() bool`

HasTypeId returns a boolean if a field has been set.

### SetTypeIdNil

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) SetTypeIdNil(b bool)`

 SetTypeIdNil sets the value for TypeId to be an explicit nil

### UnsetTypeId
`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) UnsetTypeId()`

UnsetTypeId ensures that no value is present for TypeId, not even an explicit nil
### GetCreatedAt

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetProvisioningStatus

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) GetProvisioningStatus() ProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) GetProvisioningStatusOk() (*ProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) SetProvisioningStatus(v ProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.

### HasProvisioningStatus

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) HasProvisioningStatus() bool`

HasProvisioningStatus returns a boolean if a field has been set.

### SetProvisioningStatusNil

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) SetProvisioningStatusNil(b bool)`

 SetProvisioningStatusNil sets the value for ProvisioningStatus to be an explicit nil

### UnsetProvisioningStatus
`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) UnsetProvisioningStatus()`

UnsetProvisioningStatus ensures that no value is present for ProvisioningStatus, not even an explicit nil
### GetOperatingStatus

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) GetOperatingStatus() LoadBalancerOperatingStatus`

GetOperatingStatus returns the OperatingStatus field if non-nil, zero value otherwise.

### GetOperatingStatusOk

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) GetOperatingStatusOk() (*LoadBalancerOperatingStatus, bool)`

GetOperatingStatusOk returns a tuple with the OperatingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingStatus

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) SetOperatingStatus(v LoadBalancerOperatingStatus)`

SetOperatingStatus sets OperatingStatus field to given value.

### HasOperatingStatus

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) HasOperatingStatus() bool`

HasOperatingStatus returns a boolean if a field has been set.

### SetOperatingStatusNil

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) SetOperatingStatusNil(b bool)`

 SetOperatingStatusNil sets the value for OperatingStatus to be an explicit nil

### UnsetOperatingStatus
`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) UnsetOperatingStatus()`

UnsetOperatingStatus ensures that no value is present for OperatingStatus, not even an explicit nil
### GetVpcId

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### SetVpcIdNil

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) SetVpcIdNil(b bool)`

 SetVpcIdNil sets the value for VpcId to be an explicit nil

### UnsetVpcId
`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) UnsetVpcId()`

UnsetVpcId ensures that no value is present for VpcId, not even an explicit nil
### GetType

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) GetType() LoadBalancerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) GetTypeOk() (*LoadBalancerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) SetType(v LoadBalancerType)`

SetType sets Type field to given value.

### HasType

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetVpcName

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) GetVpcName() string`

GetVpcName returns the VpcName field if non-nil, zero value otherwise.

### GetVpcNameOk

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) GetVpcNameOk() (*string, bool)`

GetVpcNameOk returns a tuple with the VpcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcName

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) SetVpcName(v string)`

SetVpcName sets VpcName field to given value.

### HasVpcName

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) HasVpcName() bool`

HasVpcName returns a boolean if a field has been set.

### SetVpcNameNil

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) SetVpcNameNil(b bool)`

 SetVpcNameNil sets the value for VpcName to be an explicit nil

### UnsetVpcName
`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) UnsetVpcName()`

UnsetVpcName ensures that no value is present for VpcName, not even an explicit nil
### GetVpcCidrBlock

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) GetVpcCidrBlock() string`

GetVpcCidrBlock returns the VpcCidrBlock field if non-nil, zero value otherwise.

### GetVpcCidrBlockOk

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) GetVpcCidrBlockOk() (*string, bool)`

GetVpcCidrBlockOk returns a tuple with the VpcCidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcCidrBlock

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) SetVpcCidrBlock(v string)`

SetVpcCidrBlock sets VpcCidrBlock field to given value.

### HasVpcCidrBlock

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) HasVpcCidrBlock() bool`

HasVpcCidrBlock returns a boolean if a field has been set.

### SetVpcCidrBlockNil

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) SetVpcCidrBlockNil(b bool)`

 SetVpcCidrBlockNil sets the value for VpcCidrBlock to be an explicit nil

### UnsetVpcCidrBlock
`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) UnsetVpcCidrBlock()`

UnsetVpcCidrBlock ensures that no value is present for VpcCidrBlock, not even an explicit nil
### GetAvailabilityZones

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) GetAvailabilityZones() []AvailabilityZone`

GetAvailabilityZones returns the AvailabilityZones field if non-nil, zero value otherwise.

### GetAvailabilityZonesOk

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) GetAvailabilityZonesOk() (*[]AvailabilityZone, bool)`

GetAvailabilityZonesOk returns a tuple with the AvailabilityZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZones

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) SetAvailabilityZones(v []AvailabilityZone)`

SetAvailabilityZones sets AvailabilityZones field to given value.

### HasAvailabilityZones

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) HasAvailabilityZones() bool`

HasAvailabilityZones returns a boolean if a field has been set.

### SetAvailabilityZonesNil

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) SetAvailabilityZonesNil(b bool)`

 SetAvailabilityZonesNil sets the value for AvailabilityZones to be an explicit nil

### UnsetAvailabilityZones
`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) UnsetAvailabilityZones()`

UnsetAvailabilityZones ensures that no value is present for AvailabilityZones, not even an explicit nil
### GetLoadBalancers

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) GetLoadBalancers() []BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel`

GetLoadBalancers returns the LoadBalancers field if non-nil, zero value otherwise.

### GetLoadBalancersOk

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) GetLoadBalancersOk() (*[]BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel, bool)`

GetLoadBalancersOk returns a tuple with the LoadBalancers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancers

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) SetLoadBalancers(v []BnsLoadBalancerV1ApiListHaGroupsModelLoadBalancerModel)`

SetLoadBalancers sets LoadBalancers field to given value.

### HasLoadBalancers

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) HasLoadBalancers() bool`

HasLoadBalancers returns a boolean if a field has been set.

### SetLoadBalancersNil

`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) SetLoadBalancersNil(b bool)`

 SetLoadBalancersNil sets the value for LoadBalancers to be an explicit nil

### UnsetLoadBalancers
`func (o *BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel) UnsetLoadBalancers()`

UnsetLoadBalancers ensures that no value is present for LoadBalancers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


