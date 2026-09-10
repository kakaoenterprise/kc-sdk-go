# Vpc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Region** | Pointer to [**NullableRegion**](Region.md) |  | [optional] 
**Igw** | Pointer to [**NullableVpcIgw**](VpcIgw.md) |  | [optional] 
**DefaultRouteTable** | Pointer to [**NullableDefaultRouteTable**](DefaultRouteTable.md) |  | [optional] 
**ProjectId** | Pointer to **NullableString** |  | [optional] 
**ProjectName** | Pointer to **NullableString** |  | [optional] 
**CidrBlock** | Pointer to **NullableString** |  | [optional] 
**IsDefault** | Pointer to **NullableBool** |  | [optional] 
**ProvisioningStatus** | Pointer to [**NullableProvisioningStatus**](ProvisioningStatus.md) |  | [optional] 
**IsEnableDnsSupport** | Pointer to **NullableBool** |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewVpc

`func NewVpc() *Vpc`

NewVpc instantiates a new Vpc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpcWithDefaults

`func NewVpcWithDefaults() *Vpc`

NewVpcWithDefaults instantiates a new Vpc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Vpc) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Vpc) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Vpc) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Vpc) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Vpc) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Vpc) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *Vpc) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Vpc) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Vpc) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Vpc) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Vpc) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Vpc) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *Vpc) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Vpc) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Vpc) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Vpc) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Vpc) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Vpc) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetRegion

`func (o *Vpc) GetRegion() Region`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Vpc) GetRegionOk() (*Region, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Vpc) SetRegion(v Region)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Vpc) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *Vpc) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *Vpc) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetIgw

`func (o *Vpc) GetIgw() VpcIgw`

GetIgw returns the Igw field if non-nil, zero value otherwise.

### GetIgwOk

`func (o *Vpc) GetIgwOk() (*VpcIgw, bool)`

GetIgwOk returns a tuple with the Igw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgw

`func (o *Vpc) SetIgw(v VpcIgw)`

SetIgw sets Igw field to given value.

### HasIgw

`func (o *Vpc) HasIgw() bool`

HasIgw returns a boolean if a field has been set.

### SetIgwNil

`func (o *Vpc) SetIgwNil(b bool)`

 SetIgwNil sets the value for Igw to be an explicit nil

### UnsetIgw
`func (o *Vpc) UnsetIgw()`

UnsetIgw ensures that no value is present for Igw, not even an explicit nil
### GetDefaultRouteTable

`func (o *Vpc) GetDefaultRouteTable() DefaultRouteTable`

GetDefaultRouteTable returns the DefaultRouteTable field if non-nil, zero value otherwise.

### GetDefaultRouteTableOk

`func (o *Vpc) GetDefaultRouteTableOk() (*DefaultRouteTable, bool)`

GetDefaultRouteTableOk returns a tuple with the DefaultRouteTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRouteTable

`func (o *Vpc) SetDefaultRouteTable(v DefaultRouteTable)`

SetDefaultRouteTable sets DefaultRouteTable field to given value.

### HasDefaultRouteTable

`func (o *Vpc) HasDefaultRouteTable() bool`

HasDefaultRouteTable returns a boolean if a field has been set.

### SetDefaultRouteTableNil

`func (o *Vpc) SetDefaultRouteTableNil(b bool)`

 SetDefaultRouteTableNil sets the value for DefaultRouteTable to be an explicit nil

### UnsetDefaultRouteTable
`func (o *Vpc) UnsetDefaultRouteTable()`

UnsetDefaultRouteTable ensures that no value is present for DefaultRouteTable, not even an explicit nil
### GetProjectId

`func (o *Vpc) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *Vpc) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *Vpc) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *Vpc) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *Vpc) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *Vpc) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetProjectName

`func (o *Vpc) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *Vpc) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *Vpc) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *Vpc) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### SetProjectNameNil

`func (o *Vpc) SetProjectNameNil(b bool)`

 SetProjectNameNil sets the value for ProjectName to be an explicit nil

### UnsetProjectName
`func (o *Vpc) UnsetProjectName()`

UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil
### GetCidrBlock

`func (o *Vpc) GetCidrBlock() string`

GetCidrBlock returns the CidrBlock field if non-nil, zero value otherwise.

### GetCidrBlockOk

`func (o *Vpc) GetCidrBlockOk() (*string, bool)`

GetCidrBlockOk returns a tuple with the CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlock

`func (o *Vpc) SetCidrBlock(v string)`

SetCidrBlock sets CidrBlock field to given value.

### HasCidrBlock

`func (o *Vpc) HasCidrBlock() bool`

HasCidrBlock returns a boolean if a field has been set.

### SetCidrBlockNil

`func (o *Vpc) SetCidrBlockNil(b bool)`

 SetCidrBlockNil sets the value for CidrBlock to be an explicit nil

### UnsetCidrBlock
`func (o *Vpc) UnsetCidrBlock()`

UnsetCidrBlock ensures that no value is present for CidrBlock, not even an explicit nil
### GetIsDefault

`func (o *Vpc) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *Vpc) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *Vpc) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *Vpc) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### SetIsDefaultNil

`func (o *Vpc) SetIsDefaultNil(b bool)`

 SetIsDefaultNil sets the value for IsDefault to be an explicit nil

### UnsetIsDefault
`func (o *Vpc) UnsetIsDefault()`

UnsetIsDefault ensures that no value is present for IsDefault, not even an explicit nil
### GetProvisioningStatus

`func (o *Vpc) GetProvisioningStatus() ProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *Vpc) GetProvisioningStatusOk() (*ProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *Vpc) SetProvisioningStatus(v ProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.

### HasProvisioningStatus

`func (o *Vpc) HasProvisioningStatus() bool`

HasProvisioningStatus returns a boolean if a field has been set.

### SetProvisioningStatusNil

`func (o *Vpc) SetProvisioningStatusNil(b bool)`

 SetProvisioningStatusNil sets the value for ProvisioningStatus to be an explicit nil

### UnsetProvisioningStatus
`func (o *Vpc) UnsetProvisioningStatus()`

UnsetProvisioningStatus ensures that no value is present for ProvisioningStatus, not even an explicit nil
### GetIsEnableDnsSupport

`func (o *Vpc) GetIsEnableDnsSupport() bool`

GetIsEnableDnsSupport returns the IsEnableDnsSupport field if non-nil, zero value otherwise.

### GetIsEnableDnsSupportOk

`func (o *Vpc) GetIsEnableDnsSupportOk() (*bool, bool)`

GetIsEnableDnsSupportOk returns a tuple with the IsEnableDnsSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnableDnsSupport

`func (o *Vpc) SetIsEnableDnsSupport(v bool)`

SetIsEnableDnsSupport sets IsEnableDnsSupport field to given value.

### HasIsEnableDnsSupport

`func (o *Vpc) HasIsEnableDnsSupport() bool`

HasIsEnableDnsSupport returns a boolean if a field has been set.

### SetIsEnableDnsSupportNil

`func (o *Vpc) SetIsEnableDnsSupportNil(b bool)`

 SetIsEnableDnsSupportNil sets the value for IsEnableDnsSupport to be an explicit nil

### UnsetIsEnableDnsSupport
`func (o *Vpc) UnsetIsEnableDnsSupport()`

UnsetIsEnableDnsSupport ensures that no value is present for IsEnableDnsSupport, not even an explicit nil
### GetCreatedAt

`func (o *Vpc) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Vpc) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Vpc) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Vpc) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *Vpc) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *Vpc) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *Vpc) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Vpc) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Vpc) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Vpc) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *Vpc) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *Vpc) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


