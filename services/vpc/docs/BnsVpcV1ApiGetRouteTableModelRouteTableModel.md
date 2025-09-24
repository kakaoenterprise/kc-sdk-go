# BnsVpcV1ApiGetRouteTableModelRouteTableModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 라우팅 테이블의 ID | 
**Name** | Pointer to **NullableString** |  | [optional] 
**Associations** | Pointer to [**[]BnsVpcV1ApiGetRouteTableModelAssociationModel**](BnsVpcV1ApiGetRouteTableModelAssociationModel.md) |  | [optional] 
**Routes** | Pointer to [**[]BnsVpcV1ApiGetRouteTableModelRouteModel**](BnsVpcV1ApiGetRouteTableModelRouteModel.md) |  | [optional] 
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

### NewBnsVpcV1ApiGetRouteTableModelRouteTableModel

`func NewBnsVpcV1ApiGetRouteTableModelRouteTableModel(id string, ) *BnsVpcV1ApiGetRouteTableModelRouteTableModel`

NewBnsVpcV1ApiGetRouteTableModelRouteTableModel instantiates a new BnsVpcV1ApiGetRouteTableModelRouteTableModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsVpcV1ApiGetRouteTableModelRouteTableModelWithDefaults

`func NewBnsVpcV1ApiGetRouteTableModelRouteTableModelWithDefaults() *BnsVpcV1ApiGetRouteTableModelRouteTableModel`

NewBnsVpcV1ApiGetRouteTableModelRouteTableModelWithDefaults instantiates a new BnsVpcV1ApiGetRouteTableModelRouteTableModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAssociations

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) GetAssociations() []BnsVpcV1ApiGetRouteTableModelAssociationModel`

GetAssociations returns the Associations field if non-nil, zero value otherwise.

### GetAssociationsOk

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) GetAssociationsOk() (*[]BnsVpcV1ApiGetRouteTableModelAssociationModel, bool)`

GetAssociationsOk returns a tuple with the Associations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociations

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) SetAssociations(v []BnsVpcV1ApiGetRouteTableModelAssociationModel)`

SetAssociations sets Associations field to given value.

### HasAssociations

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) HasAssociations() bool`

HasAssociations returns a boolean if a field has been set.

### SetAssociationsNil

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) SetAssociationsNil(b bool)`

 SetAssociationsNil sets the value for Associations to be an explicit nil

### UnsetAssociations
`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) UnsetAssociations()`

UnsetAssociations ensures that no value is present for Associations, not even an explicit nil
### GetRoutes

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) GetRoutes() []BnsVpcV1ApiGetRouteTableModelRouteModel`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) GetRoutesOk() (*[]BnsVpcV1ApiGetRouteTableModelRouteModel, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) SetRoutes(v []BnsVpcV1ApiGetRouteTableModelRouteModel)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.

### SetRoutesNil

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) SetRoutesNil(b bool)`

 SetRoutesNil sets the value for Routes to be an explicit nil

### UnsetRoutes
`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) UnsetRoutes()`

UnsetRoutes ensures that no value is present for Routes, not even an explicit nil
### GetVpcId

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### SetVpcIdNil

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) SetVpcIdNil(b bool)`

 SetVpcIdNil sets the value for VpcId to be an explicit nil

### UnsetVpcId
`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) UnsetVpcId()`

UnsetVpcId ensures that no value is present for VpcId, not even an explicit nil
### GetProvisioningStatus

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) GetProvisioningStatus() ProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) GetProvisioningStatusOk() (*ProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) SetProvisioningStatus(v ProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.

### HasProvisioningStatus

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) HasProvisioningStatus() bool`

HasProvisioningStatus returns a boolean if a field has been set.

### SetProvisioningStatusNil

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) SetProvisioningStatusNil(b bool)`

 SetProvisioningStatusNil sets the value for ProvisioningStatus to be an explicit nil

### UnsetProvisioningStatus
`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) UnsetProvisioningStatus()`

UnsetProvisioningStatus ensures that no value is present for ProvisioningStatus, not even an explicit nil
### GetVpcName

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) GetVpcName() string`

GetVpcName returns the VpcName field if non-nil, zero value otherwise.

### GetVpcNameOk

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) GetVpcNameOk() (*string, bool)`

GetVpcNameOk returns a tuple with the VpcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcName

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) SetVpcName(v string)`

SetVpcName sets VpcName field to given value.

### HasVpcName

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) HasVpcName() bool`

HasVpcName returns a boolean if a field has been set.

### SetVpcNameNil

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) SetVpcNameNil(b bool)`

 SetVpcNameNil sets the value for VpcName to be an explicit nil

### UnsetVpcName
`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) UnsetVpcName()`

UnsetVpcName ensures that no value is present for VpcName, not even an explicit nil
### GetVpcProvisioningStatus

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) GetVpcProvisioningStatus() ProvisioningStatus`

GetVpcProvisioningStatus returns the VpcProvisioningStatus field if non-nil, zero value otherwise.

### GetVpcProvisioningStatusOk

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) GetVpcProvisioningStatusOk() (*ProvisioningStatus, bool)`

GetVpcProvisioningStatusOk returns a tuple with the VpcProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcProvisioningStatus

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) SetVpcProvisioningStatus(v ProvisioningStatus)`

SetVpcProvisioningStatus sets VpcProvisioningStatus field to given value.

### HasVpcProvisioningStatus

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) HasVpcProvisioningStatus() bool`

HasVpcProvisioningStatus returns a boolean if a field has been set.

### SetVpcProvisioningStatusNil

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) SetVpcProvisioningStatusNil(b bool)`

 SetVpcProvisioningStatusNil sets the value for VpcProvisioningStatus to be an explicit nil

### UnsetVpcProvisioningStatus
`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) UnsetVpcProvisioningStatus()`

UnsetVpcProvisioningStatus ensures that no value is present for VpcProvisioningStatus, not even an explicit nil
### GetProjectId

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetProjectName

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### SetProjectNameNil

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) SetProjectNameNil(b bool)`

 SetProjectNameNil sets the value for ProjectName to be an explicit nil

### UnsetProjectName
`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) UnsetProjectName()`

UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil
### GetIsMain

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) GetIsMain() bool`

GetIsMain returns the IsMain field if non-nil, zero value otherwise.

### GetIsMainOk

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) GetIsMainOk() (*bool, bool)`

GetIsMainOk returns a tuple with the IsMain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMain

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) SetIsMain(v bool)`

SetIsMain sets IsMain field to given value.

### HasIsMain

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) HasIsMain() bool`

HasIsMain returns a boolean if a field has been set.

### SetIsMainNil

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) SetIsMainNil(b bool)`

 SetIsMainNil sets the value for IsMain to be an explicit nil

### UnsetIsMain
`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) UnsetIsMain()`

UnsetIsMain ensures that no value is present for IsMain, not even an explicit nil
### GetCreatedAt

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *BnsVpcV1ApiGetRouteTableModelRouteTableModel) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


