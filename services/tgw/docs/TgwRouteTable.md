# TgwRouteTable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**ProjectId** | Pointer to **NullableString** |  | [optional] 
**ProjectName** | Pointer to **NullableString** |  | [optional] 
**Region** | Pointer to **NullableString** |  | [optional] 
**TgwId** | Pointer to **NullableString** |  | [optional] 
**TgwName** | Pointer to **NullableString** |  | [optional] 
**IsDefaultAssociationRouteTable** | Pointer to **NullableBool** |  | [optional] 
**IsDefaultPropagationRouteTable** | Pointer to **NullableBool** |  | [optional] 
**ProvisioningStatus** | Pointer to [**NullableTGWRouteTableProvisioningStatus**](TGWRouteTableProvisioningStatus.md) |  | [optional] 
**Routes** | Pointer to [**[]Route**](Route.md) |  | [optional] 
**Associations** | Pointer to [**[]Association**](Association.md) |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewTgwRouteTable

`func NewTgwRouteTable() *TgwRouteTable`

NewTgwRouteTable instantiates a new TgwRouteTable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTgwRouteTableWithDefaults

`func NewTgwRouteTableWithDefaults() *TgwRouteTable`

NewTgwRouteTableWithDefaults instantiates a new TgwRouteTable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TgwRouteTable) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TgwRouteTable) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TgwRouteTable) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TgwRouteTable) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *TgwRouteTable) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TgwRouteTable) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *TgwRouteTable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TgwRouteTable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TgwRouteTable) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TgwRouteTable) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TgwRouteTable) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TgwRouteTable) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetProjectId

`func (o *TgwRouteTable) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *TgwRouteTable) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *TgwRouteTable) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *TgwRouteTable) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *TgwRouteTable) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *TgwRouteTable) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetProjectName

`func (o *TgwRouteTable) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *TgwRouteTable) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *TgwRouteTable) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *TgwRouteTable) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### SetProjectNameNil

`func (o *TgwRouteTable) SetProjectNameNil(b bool)`

 SetProjectNameNil sets the value for ProjectName to be an explicit nil

### UnsetProjectName
`func (o *TgwRouteTable) UnsetProjectName()`

UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil
### GetRegion

`func (o *TgwRouteTable) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *TgwRouteTable) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *TgwRouteTable) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *TgwRouteTable) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *TgwRouteTable) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *TgwRouteTable) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetTgwId

`func (o *TgwRouteTable) GetTgwId() string`

GetTgwId returns the TgwId field if non-nil, zero value otherwise.

### GetTgwIdOk

`func (o *TgwRouteTable) GetTgwIdOk() (*string, bool)`

GetTgwIdOk returns a tuple with the TgwId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgwId

`func (o *TgwRouteTable) SetTgwId(v string)`

SetTgwId sets TgwId field to given value.

### HasTgwId

`func (o *TgwRouteTable) HasTgwId() bool`

HasTgwId returns a boolean if a field has been set.

### SetTgwIdNil

`func (o *TgwRouteTable) SetTgwIdNil(b bool)`

 SetTgwIdNil sets the value for TgwId to be an explicit nil

### UnsetTgwId
`func (o *TgwRouteTable) UnsetTgwId()`

UnsetTgwId ensures that no value is present for TgwId, not even an explicit nil
### GetTgwName

`func (o *TgwRouteTable) GetTgwName() string`

GetTgwName returns the TgwName field if non-nil, zero value otherwise.

### GetTgwNameOk

`func (o *TgwRouteTable) GetTgwNameOk() (*string, bool)`

GetTgwNameOk returns a tuple with the TgwName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgwName

`func (o *TgwRouteTable) SetTgwName(v string)`

SetTgwName sets TgwName field to given value.

### HasTgwName

`func (o *TgwRouteTable) HasTgwName() bool`

HasTgwName returns a boolean if a field has been set.

### SetTgwNameNil

`func (o *TgwRouteTable) SetTgwNameNil(b bool)`

 SetTgwNameNil sets the value for TgwName to be an explicit nil

### UnsetTgwName
`func (o *TgwRouteTable) UnsetTgwName()`

UnsetTgwName ensures that no value is present for TgwName, not even an explicit nil
### GetIsDefaultAssociationRouteTable

`func (o *TgwRouteTable) GetIsDefaultAssociationRouteTable() bool`

GetIsDefaultAssociationRouteTable returns the IsDefaultAssociationRouteTable field if non-nil, zero value otherwise.

### GetIsDefaultAssociationRouteTableOk

`func (o *TgwRouteTable) GetIsDefaultAssociationRouteTableOk() (*bool, bool)`

GetIsDefaultAssociationRouteTableOk returns a tuple with the IsDefaultAssociationRouteTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultAssociationRouteTable

`func (o *TgwRouteTable) SetIsDefaultAssociationRouteTable(v bool)`

SetIsDefaultAssociationRouteTable sets IsDefaultAssociationRouteTable field to given value.

### HasIsDefaultAssociationRouteTable

`func (o *TgwRouteTable) HasIsDefaultAssociationRouteTable() bool`

HasIsDefaultAssociationRouteTable returns a boolean if a field has been set.

### SetIsDefaultAssociationRouteTableNil

`func (o *TgwRouteTable) SetIsDefaultAssociationRouteTableNil(b bool)`

 SetIsDefaultAssociationRouteTableNil sets the value for IsDefaultAssociationRouteTable to be an explicit nil

### UnsetIsDefaultAssociationRouteTable
`func (o *TgwRouteTable) UnsetIsDefaultAssociationRouteTable()`

UnsetIsDefaultAssociationRouteTable ensures that no value is present for IsDefaultAssociationRouteTable, not even an explicit nil
### GetIsDefaultPropagationRouteTable

`func (o *TgwRouteTable) GetIsDefaultPropagationRouteTable() bool`

GetIsDefaultPropagationRouteTable returns the IsDefaultPropagationRouteTable field if non-nil, zero value otherwise.

### GetIsDefaultPropagationRouteTableOk

`func (o *TgwRouteTable) GetIsDefaultPropagationRouteTableOk() (*bool, bool)`

GetIsDefaultPropagationRouteTableOk returns a tuple with the IsDefaultPropagationRouteTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultPropagationRouteTable

`func (o *TgwRouteTable) SetIsDefaultPropagationRouteTable(v bool)`

SetIsDefaultPropagationRouteTable sets IsDefaultPropagationRouteTable field to given value.

### HasIsDefaultPropagationRouteTable

`func (o *TgwRouteTable) HasIsDefaultPropagationRouteTable() bool`

HasIsDefaultPropagationRouteTable returns a boolean if a field has been set.

### SetIsDefaultPropagationRouteTableNil

`func (o *TgwRouteTable) SetIsDefaultPropagationRouteTableNil(b bool)`

 SetIsDefaultPropagationRouteTableNil sets the value for IsDefaultPropagationRouteTable to be an explicit nil

### UnsetIsDefaultPropagationRouteTable
`func (o *TgwRouteTable) UnsetIsDefaultPropagationRouteTable()`

UnsetIsDefaultPropagationRouteTable ensures that no value is present for IsDefaultPropagationRouteTable, not even an explicit nil
### GetProvisioningStatus

`func (o *TgwRouteTable) GetProvisioningStatus() TGWRouteTableProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *TgwRouteTable) GetProvisioningStatusOk() (*TGWRouteTableProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *TgwRouteTable) SetProvisioningStatus(v TGWRouteTableProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.

### HasProvisioningStatus

`func (o *TgwRouteTable) HasProvisioningStatus() bool`

HasProvisioningStatus returns a boolean if a field has been set.

### SetProvisioningStatusNil

`func (o *TgwRouteTable) SetProvisioningStatusNil(b bool)`

 SetProvisioningStatusNil sets the value for ProvisioningStatus to be an explicit nil

### UnsetProvisioningStatus
`func (o *TgwRouteTable) UnsetProvisioningStatus()`

UnsetProvisioningStatus ensures that no value is present for ProvisioningStatus, not even an explicit nil
### GetRoutes

`func (o *TgwRouteTable) GetRoutes() []Route`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *TgwRouteTable) GetRoutesOk() (*[]Route, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *TgwRouteTable) SetRoutes(v []Route)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *TgwRouteTable) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.

### SetRoutesNil

`func (o *TgwRouteTable) SetRoutesNil(b bool)`

 SetRoutesNil sets the value for Routes to be an explicit nil

### UnsetRoutes
`func (o *TgwRouteTable) UnsetRoutes()`

UnsetRoutes ensures that no value is present for Routes, not even an explicit nil
### GetAssociations

`func (o *TgwRouteTable) GetAssociations() []Association`

GetAssociations returns the Associations field if non-nil, zero value otherwise.

### GetAssociationsOk

`func (o *TgwRouteTable) GetAssociationsOk() (*[]Association, bool)`

GetAssociationsOk returns a tuple with the Associations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociations

`func (o *TgwRouteTable) SetAssociations(v []Association)`

SetAssociations sets Associations field to given value.

### HasAssociations

`func (o *TgwRouteTable) HasAssociations() bool`

HasAssociations returns a boolean if a field has been set.

### SetAssociationsNil

`func (o *TgwRouteTable) SetAssociationsNil(b bool)`

 SetAssociationsNil sets the value for Associations to be an explicit nil

### UnsetAssociations
`func (o *TgwRouteTable) UnsetAssociations()`

UnsetAssociations ensures that no value is present for Associations, not even an explicit nil
### GetCreatedAt

`func (o *TgwRouteTable) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TgwRouteTable) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TgwRouteTable) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TgwRouteTable) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *TgwRouteTable) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *TgwRouteTable) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *TgwRouteTable) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TgwRouteTable) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TgwRouteTable) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TgwRouteTable) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *TgwRouteTable) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *TgwRouteTable) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


