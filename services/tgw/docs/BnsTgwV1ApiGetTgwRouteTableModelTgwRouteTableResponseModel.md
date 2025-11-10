# BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**ProjectId** | Pointer to **NullableString** |  | [optional] 
**ProjectName** | Pointer to **NullableString** |  | [optional] 
**Region** | Pointer to [**NullableRegion**](Region.md) |  | [optional] 
**TgwId** | Pointer to **NullableString** |  | [optional] 
**TgwName** | Pointer to **NullableString** |  | [optional] 
**IsDefaultAssociationRouteTable** | Pointer to **NullableBool** |  | [optional] 
**IsDefaultPropagationRouteTable** | Pointer to **NullableBool** |  | [optional] 
**ProvisioningStatus** | Pointer to [**NullableTGWProvisioningStatus**](TGWProvisioningStatus.md) |  | [optional] 
**Routes** | Pointer to [**[]BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel**](BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel.md) |  | [optional] 
**Associations** | Pointer to [**[]BnsTgwV1ApiGetTgwRouteTableModelAssociationResponseModel**](BnsTgwV1ApiGetTgwRouteTableModelAssociationResponseModel.md) |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewBnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel

`func NewBnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel() *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel`

NewBnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel instantiates a new BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModelWithDefaults

`func NewBnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModelWithDefaults() *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel`

NewBnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModelWithDefaults instantiates a new BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetProjectId

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetProjectName

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### SetProjectNameNil

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) SetProjectNameNil(b bool)`

 SetProjectNameNil sets the value for ProjectName to be an explicit nil

### UnsetProjectName
`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) UnsetProjectName()`

UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil
### GetRegion

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) GetRegion() Region`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) GetRegionOk() (*Region, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) SetRegion(v Region)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetTgwId

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) GetTgwId() string`

GetTgwId returns the TgwId field if non-nil, zero value otherwise.

### GetTgwIdOk

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) GetTgwIdOk() (*string, bool)`

GetTgwIdOk returns a tuple with the TgwId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgwId

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) SetTgwId(v string)`

SetTgwId sets TgwId field to given value.

### HasTgwId

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) HasTgwId() bool`

HasTgwId returns a boolean if a field has been set.

### SetTgwIdNil

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) SetTgwIdNil(b bool)`

 SetTgwIdNil sets the value for TgwId to be an explicit nil

### UnsetTgwId
`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) UnsetTgwId()`

UnsetTgwId ensures that no value is present for TgwId, not even an explicit nil
### GetTgwName

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) GetTgwName() string`

GetTgwName returns the TgwName field if non-nil, zero value otherwise.

### GetTgwNameOk

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) GetTgwNameOk() (*string, bool)`

GetTgwNameOk returns a tuple with the TgwName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgwName

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) SetTgwName(v string)`

SetTgwName sets TgwName field to given value.

### HasTgwName

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) HasTgwName() bool`

HasTgwName returns a boolean if a field has been set.

### SetTgwNameNil

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) SetTgwNameNil(b bool)`

 SetTgwNameNil sets the value for TgwName to be an explicit nil

### UnsetTgwName
`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) UnsetTgwName()`

UnsetTgwName ensures that no value is present for TgwName, not even an explicit nil
### GetIsDefaultAssociationRouteTable

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) GetIsDefaultAssociationRouteTable() bool`

GetIsDefaultAssociationRouteTable returns the IsDefaultAssociationRouteTable field if non-nil, zero value otherwise.

### GetIsDefaultAssociationRouteTableOk

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) GetIsDefaultAssociationRouteTableOk() (*bool, bool)`

GetIsDefaultAssociationRouteTableOk returns a tuple with the IsDefaultAssociationRouteTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultAssociationRouteTable

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) SetIsDefaultAssociationRouteTable(v bool)`

SetIsDefaultAssociationRouteTable sets IsDefaultAssociationRouteTable field to given value.

### HasIsDefaultAssociationRouteTable

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) HasIsDefaultAssociationRouteTable() bool`

HasIsDefaultAssociationRouteTable returns a boolean if a field has been set.

### SetIsDefaultAssociationRouteTableNil

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) SetIsDefaultAssociationRouteTableNil(b bool)`

 SetIsDefaultAssociationRouteTableNil sets the value for IsDefaultAssociationRouteTable to be an explicit nil

### UnsetIsDefaultAssociationRouteTable
`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) UnsetIsDefaultAssociationRouteTable()`

UnsetIsDefaultAssociationRouteTable ensures that no value is present for IsDefaultAssociationRouteTable, not even an explicit nil
### GetIsDefaultPropagationRouteTable

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) GetIsDefaultPropagationRouteTable() bool`

GetIsDefaultPropagationRouteTable returns the IsDefaultPropagationRouteTable field if non-nil, zero value otherwise.

### GetIsDefaultPropagationRouteTableOk

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) GetIsDefaultPropagationRouteTableOk() (*bool, bool)`

GetIsDefaultPropagationRouteTableOk returns a tuple with the IsDefaultPropagationRouteTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultPropagationRouteTable

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) SetIsDefaultPropagationRouteTable(v bool)`

SetIsDefaultPropagationRouteTable sets IsDefaultPropagationRouteTable field to given value.

### HasIsDefaultPropagationRouteTable

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) HasIsDefaultPropagationRouteTable() bool`

HasIsDefaultPropagationRouteTable returns a boolean if a field has been set.

### SetIsDefaultPropagationRouteTableNil

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) SetIsDefaultPropagationRouteTableNil(b bool)`

 SetIsDefaultPropagationRouteTableNil sets the value for IsDefaultPropagationRouteTable to be an explicit nil

### UnsetIsDefaultPropagationRouteTable
`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) UnsetIsDefaultPropagationRouteTable()`

UnsetIsDefaultPropagationRouteTable ensures that no value is present for IsDefaultPropagationRouteTable, not even an explicit nil
### GetProvisioningStatus

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) GetProvisioningStatus() TGWProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) GetProvisioningStatusOk() (*TGWProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) SetProvisioningStatus(v TGWProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.

### HasProvisioningStatus

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) HasProvisioningStatus() bool`

HasProvisioningStatus returns a boolean if a field has been set.

### SetProvisioningStatusNil

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) SetProvisioningStatusNil(b bool)`

 SetProvisioningStatusNil sets the value for ProvisioningStatus to be an explicit nil

### UnsetProvisioningStatus
`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) UnsetProvisioningStatus()`

UnsetProvisioningStatus ensures that no value is present for ProvisioningStatus, not even an explicit nil
### GetRoutes

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) GetRoutes() []BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) GetRoutesOk() (*[]BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) SetRoutes(v []BnsTgwV1ApiGetTgwRouteTableModelRouteResponseModel)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.

### SetRoutesNil

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) SetRoutesNil(b bool)`

 SetRoutesNil sets the value for Routes to be an explicit nil

### UnsetRoutes
`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) UnsetRoutes()`

UnsetRoutes ensures that no value is present for Routes, not even an explicit nil
### GetAssociations

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) GetAssociations() []BnsTgwV1ApiGetTgwRouteTableModelAssociationResponseModel`

GetAssociations returns the Associations field if non-nil, zero value otherwise.

### GetAssociationsOk

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) GetAssociationsOk() (*[]BnsTgwV1ApiGetTgwRouteTableModelAssociationResponseModel, bool)`

GetAssociationsOk returns a tuple with the Associations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociations

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) SetAssociations(v []BnsTgwV1ApiGetTgwRouteTableModelAssociationResponseModel)`

SetAssociations sets Associations field to given value.

### HasAssociations

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) HasAssociations() bool`

HasAssociations returns a boolean if a field has been set.

### SetAssociationsNil

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) SetAssociationsNil(b bool)`

 SetAssociationsNil sets the value for Associations to be an explicit nil

### UnsetAssociations
`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) UnsetAssociations()`

UnsetAssociations ensures that no value is present for Associations, not even an explicit nil
### GetCreatedAt

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *BnsTgwV1ApiGetTgwRouteTableModelTgwRouteTableResponseModel) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


