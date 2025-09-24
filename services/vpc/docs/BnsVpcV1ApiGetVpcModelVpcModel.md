# BnsVpcV1ApiGetVpcModelVpcModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | VPC의 고유 ID | 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Region** | Pointer to [**NullableRegion**](Region.md) |  | [optional] 
**Igw** | Pointer to [**NullableBnsVpcV1ApiGetVpcModelIgwModel**](BnsVpcV1ApiGetVpcModelIgwModel.md) |  | [optional] 
**DefaultRouteTable** | Pointer to [**NullableBnsVpcV1ApiGetVpcModelRouteTableModel**](BnsVpcV1ApiGetVpcModelRouteTableModel.md) |  | [optional] 
**ProjectId** | Pointer to **NullableString** |  | [optional] 
**ProjectName** | Pointer to **NullableString** |  | [optional] 
**CidrBlock** | Pointer to **NullableString** |  | [optional] 
**IsDefault** | Pointer to **NullableBool** |  | [optional] 
**ProvisioningStatus** | Pointer to [**NullableProvisioningStatus**](ProvisioningStatus.md) |  | [optional] 
**IsEnableDnsSupport** | Pointer to **NullableBool** |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewBnsVpcV1ApiGetVpcModelVpcModel

`func NewBnsVpcV1ApiGetVpcModelVpcModel(id string, ) *BnsVpcV1ApiGetVpcModelVpcModel`

NewBnsVpcV1ApiGetVpcModelVpcModel instantiates a new BnsVpcV1ApiGetVpcModelVpcModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsVpcV1ApiGetVpcModelVpcModelWithDefaults

`func NewBnsVpcV1ApiGetVpcModelVpcModelWithDefaults() *BnsVpcV1ApiGetVpcModelVpcModel`

NewBnsVpcV1ApiGetVpcModelVpcModelWithDefaults instantiates a new BnsVpcV1ApiGetVpcModelVpcModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BnsVpcV1ApiGetVpcModelVpcModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BnsVpcV1ApiGetVpcModelVpcModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetRegion

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) GetRegion() Region`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) GetRegionOk() (*Region, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) SetRegion(v Region)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *BnsVpcV1ApiGetVpcModelVpcModel) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetIgw

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) GetIgw() BnsVpcV1ApiGetVpcModelIgwModel`

GetIgw returns the Igw field if non-nil, zero value otherwise.

### GetIgwOk

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) GetIgwOk() (*BnsVpcV1ApiGetVpcModelIgwModel, bool)`

GetIgwOk returns a tuple with the Igw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgw

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) SetIgw(v BnsVpcV1ApiGetVpcModelIgwModel)`

SetIgw sets Igw field to given value.

### HasIgw

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) HasIgw() bool`

HasIgw returns a boolean if a field has been set.

### SetIgwNil

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) SetIgwNil(b bool)`

 SetIgwNil sets the value for Igw to be an explicit nil

### UnsetIgw
`func (o *BnsVpcV1ApiGetVpcModelVpcModel) UnsetIgw()`

UnsetIgw ensures that no value is present for Igw, not even an explicit nil
### GetDefaultRouteTable

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) GetDefaultRouteTable() BnsVpcV1ApiGetVpcModelRouteTableModel`

GetDefaultRouteTable returns the DefaultRouteTable field if non-nil, zero value otherwise.

### GetDefaultRouteTableOk

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) GetDefaultRouteTableOk() (*BnsVpcV1ApiGetVpcModelRouteTableModel, bool)`

GetDefaultRouteTableOk returns a tuple with the DefaultRouteTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRouteTable

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) SetDefaultRouteTable(v BnsVpcV1ApiGetVpcModelRouteTableModel)`

SetDefaultRouteTable sets DefaultRouteTable field to given value.

### HasDefaultRouteTable

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) HasDefaultRouteTable() bool`

HasDefaultRouteTable returns a boolean if a field has been set.

### SetDefaultRouteTableNil

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) SetDefaultRouteTableNil(b bool)`

 SetDefaultRouteTableNil sets the value for DefaultRouteTable to be an explicit nil

### UnsetDefaultRouteTable
`func (o *BnsVpcV1ApiGetVpcModelVpcModel) UnsetDefaultRouteTable()`

UnsetDefaultRouteTable ensures that no value is present for DefaultRouteTable, not even an explicit nil
### GetProjectId

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *BnsVpcV1ApiGetVpcModelVpcModel) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetProjectName

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### SetProjectNameNil

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) SetProjectNameNil(b bool)`

 SetProjectNameNil sets the value for ProjectName to be an explicit nil

### UnsetProjectName
`func (o *BnsVpcV1ApiGetVpcModelVpcModel) UnsetProjectName()`

UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil
### GetCidrBlock

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) GetCidrBlock() string`

GetCidrBlock returns the CidrBlock field if non-nil, zero value otherwise.

### GetCidrBlockOk

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) GetCidrBlockOk() (*string, bool)`

GetCidrBlockOk returns a tuple with the CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlock

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) SetCidrBlock(v string)`

SetCidrBlock sets CidrBlock field to given value.

### HasCidrBlock

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) HasCidrBlock() bool`

HasCidrBlock returns a boolean if a field has been set.

### SetCidrBlockNil

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) SetCidrBlockNil(b bool)`

 SetCidrBlockNil sets the value for CidrBlock to be an explicit nil

### UnsetCidrBlock
`func (o *BnsVpcV1ApiGetVpcModelVpcModel) UnsetCidrBlock()`

UnsetCidrBlock ensures that no value is present for CidrBlock, not even an explicit nil
### GetIsDefault

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### SetIsDefaultNil

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) SetIsDefaultNil(b bool)`

 SetIsDefaultNil sets the value for IsDefault to be an explicit nil

### UnsetIsDefault
`func (o *BnsVpcV1ApiGetVpcModelVpcModel) UnsetIsDefault()`

UnsetIsDefault ensures that no value is present for IsDefault, not even an explicit nil
### GetProvisioningStatus

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) GetProvisioningStatus() ProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) GetProvisioningStatusOk() (*ProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) SetProvisioningStatus(v ProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.

### HasProvisioningStatus

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) HasProvisioningStatus() bool`

HasProvisioningStatus returns a boolean if a field has been set.

### SetProvisioningStatusNil

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) SetProvisioningStatusNil(b bool)`

 SetProvisioningStatusNil sets the value for ProvisioningStatus to be an explicit nil

### UnsetProvisioningStatus
`func (o *BnsVpcV1ApiGetVpcModelVpcModel) UnsetProvisioningStatus()`

UnsetProvisioningStatus ensures that no value is present for ProvisioningStatus, not even an explicit nil
### GetIsEnableDnsSupport

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) GetIsEnableDnsSupport() bool`

GetIsEnableDnsSupport returns the IsEnableDnsSupport field if non-nil, zero value otherwise.

### GetIsEnableDnsSupportOk

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) GetIsEnableDnsSupportOk() (*bool, bool)`

GetIsEnableDnsSupportOk returns a tuple with the IsEnableDnsSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnableDnsSupport

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) SetIsEnableDnsSupport(v bool)`

SetIsEnableDnsSupport sets IsEnableDnsSupport field to given value.

### HasIsEnableDnsSupport

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) HasIsEnableDnsSupport() bool`

HasIsEnableDnsSupport returns a boolean if a field has been set.

### SetIsEnableDnsSupportNil

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) SetIsEnableDnsSupportNil(b bool)`

 SetIsEnableDnsSupportNil sets the value for IsEnableDnsSupport to be an explicit nil

### UnsetIsEnableDnsSupport
`func (o *BnsVpcV1ApiGetVpcModelVpcModel) UnsetIsEnableDnsSupport()`

UnsetIsEnableDnsSupport ensures that no value is present for IsEnableDnsSupport, not even an explicit nil
### GetCreatedAt

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *BnsVpcV1ApiGetVpcModelVpcModel) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *BnsVpcV1ApiGetVpcModelVpcModel) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *BnsVpcV1ApiGetVpcModelVpcModel) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


