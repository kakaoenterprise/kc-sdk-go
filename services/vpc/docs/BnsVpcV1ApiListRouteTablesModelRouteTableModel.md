# BnsVpcV1ApiListRouteTablesModelRouteTableModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 라우팅 테이블의 ID | 
**Name** | Pointer to **NullableString** |  | [optional] 
**Associations** | Pointer to [**[]BnsVpcV1ApiListRouteTablesModelAssociationModel**](BnsVpcV1ApiListRouteTablesModelAssociationModel.md) |  | [optional] 
**Routes** | Pointer to [**[]BnsVpcV1ApiListRouteTablesModelRouteModel**](BnsVpcV1ApiListRouteTablesModelRouteModel.md) |  | [optional] 
**VpcId** | Pointer to **NullableString** |  | [optional] 
**ProvisioningStatus** | Pointer to [**NullableProvisioningStatus**](ProvisioningStatus.md) |  | [optional] 
**VpcName** | Pointer to **NullableString** |  | [optional] 
**VpcProvisioningStatus** | Pointer to [**NullableProvisioningStatus**](ProvisioningStatus.md) |  | [optional] 
**ProjectId** | Pointer to **NullableString** |  | [optional] 
**ProjectName** | Pointer to **NullableString** |  | [optional] 
**IsMain** | Pointer to **NullableBool** |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewBnsVpcV1ApiListRouteTablesModelRouteTableModel

`func NewBnsVpcV1ApiListRouteTablesModelRouteTableModel(id string, ) *BnsVpcV1ApiListRouteTablesModelRouteTableModel`

NewBnsVpcV1ApiListRouteTablesModelRouteTableModel instantiates a new BnsVpcV1ApiListRouteTablesModelRouteTableModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsVpcV1ApiListRouteTablesModelRouteTableModelWithDefaults

`func NewBnsVpcV1ApiListRouteTablesModelRouteTableModelWithDefaults() *BnsVpcV1ApiListRouteTablesModelRouteTableModel`

NewBnsVpcV1ApiListRouteTablesModelRouteTableModelWithDefaults instantiates a new BnsVpcV1ApiListRouteTablesModelRouteTableModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAssociations

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) GetAssociations() []BnsVpcV1ApiListRouteTablesModelAssociationModel`

GetAssociations returns the Associations field if non-nil, zero value otherwise.

### GetAssociationsOk

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) GetAssociationsOk() (*[]BnsVpcV1ApiListRouteTablesModelAssociationModel, bool)`

GetAssociationsOk returns a tuple with the Associations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociations

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) SetAssociations(v []BnsVpcV1ApiListRouteTablesModelAssociationModel)`

SetAssociations sets Associations field to given value.

### HasAssociations

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) HasAssociations() bool`

HasAssociations returns a boolean if a field has been set.

### SetAssociationsNil

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) SetAssociationsNil(b bool)`

 SetAssociationsNil sets the value for Associations to be an explicit nil

### UnsetAssociations
`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) UnsetAssociations()`

UnsetAssociations ensures that no value is present for Associations, not even an explicit nil
### GetRoutes

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) GetRoutes() []BnsVpcV1ApiListRouteTablesModelRouteModel`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) GetRoutesOk() (*[]BnsVpcV1ApiListRouteTablesModelRouteModel, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) SetRoutes(v []BnsVpcV1ApiListRouteTablesModelRouteModel)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.

### SetRoutesNil

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) SetRoutesNil(b bool)`

 SetRoutesNil sets the value for Routes to be an explicit nil

### UnsetRoutes
`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) UnsetRoutes()`

UnsetRoutes ensures that no value is present for Routes, not even an explicit nil
### GetVpcId

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### SetVpcIdNil

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) SetVpcIdNil(b bool)`

 SetVpcIdNil sets the value for VpcId to be an explicit nil

### UnsetVpcId
`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) UnsetVpcId()`

UnsetVpcId ensures that no value is present for VpcId, not even an explicit nil
### GetProvisioningStatus

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) GetProvisioningStatus() ProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) GetProvisioningStatusOk() (*ProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) SetProvisioningStatus(v ProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.

### HasProvisioningStatus

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) HasProvisioningStatus() bool`

HasProvisioningStatus returns a boolean if a field has been set.

### SetProvisioningStatusNil

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) SetProvisioningStatusNil(b bool)`

 SetProvisioningStatusNil sets the value for ProvisioningStatus to be an explicit nil

### UnsetProvisioningStatus
`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) UnsetProvisioningStatus()`

UnsetProvisioningStatus ensures that no value is present for ProvisioningStatus, not even an explicit nil
### GetVpcName

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) GetVpcName() string`

GetVpcName returns the VpcName field if non-nil, zero value otherwise.

### GetVpcNameOk

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) GetVpcNameOk() (*string, bool)`

GetVpcNameOk returns a tuple with the VpcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcName

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) SetVpcName(v string)`

SetVpcName sets VpcName field to given value.

### HasVpcName

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) HasVpcName() bool`

HasVpcName returns a boolean if a field has been set.

### SetVpcNameNil

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) SetVpcNameNil(b bool)`

 SetVpcNameNil sets the value for VpcName to be an explicit nil

### UnsetVpcName
`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) UnsetVpcName()`

UnsetVpcName ensures that no value is present for VpcName, not even an explicit nil
### GetVpcProvisioningStatus

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) GetVpcProvisioningStatus() ProvisioningStatus`

GetVpcProvisioningStatus returns the VpcProvisioningStatus field if non-nil, zero value otherwise.

### GetVpcProvisioningStatusOk

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) GetVpcProvisioningStatusOk() (*ProvisioningStatus, bool)`

GetVpcProvisioningStatusOk returns a tuple with the VpcProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcProvisioningStatus

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) SetVpcProvisioningStatus(v ProvisioningStatus)`

SetVpcProvisioningStatus sets VpcProvisioningStatus field to given value.

### HasVpcProvisioningStatus

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) HasVpcProvisioningStatus() bool`

HasVpcProvisioningStatus returns a boolean if a field has been set.

### SetVpcProvisioningStatusNil

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) SetVpcProvisioningStatusNil(b bool)`

 SetVpcProvisioningStatusNil sets the value for VpcProvisioningStatus to be an explicit nil

### UnsetVpcProvisioningStatus
`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) UnsetVpcProvisioningStatus()`

UnsetVpcProvisioningStatus ensures that no value is present for VpcProvisioningStatus, not even an explicit nil
### GetProjectId

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetProjectName

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### SetProjectNameNil

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) SetProjectNameNil(b bool)`

 SetProjectNameNil sets the value for ProjectName to be an explicit nil

### UnsetProjectName
`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) UnsetProjectName()`

UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil
### GetIsMain

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) GetIsMain() bool`

GetIsMain returns the IsMain field if non-nil, zero value otherwise.

### GetIsMainOk

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) GetIsMainOk() (*bool, bool)`

GetIsMainOk returns a tuple with the IsMain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMain

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) SetIsMain(v bool)`

SetIsMain sets IsMain field to given value.

### HasIsMain

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) HasIsMain() bool`

HasIsMain returns a boolean if a field has been set.

### SetIsMainNil

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) SetIsMainNil(b bool)`

 SetIsMainNil sets the value for IsMain to be an explicit nil

### UnsetIsMain
`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) UnsetIsMain()`

UnsetIsMain ensures that no value is present for IsMain, not even an explicit nil
### GetCreatedAt

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *BnsVpcV1ApiListRouteTablesModelRouteTableModel) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


