# BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Region** | Pointer to [**NullableRegion**](Region.md) |  | [optional] 
**IsShared** | Pointer to **NullableBool** |  | [optional] 
**Attachments** | Pointer to [**[]BnsTgwV1ApiGetTransitGatewayModelAttachmentResponseModel**](BnsTgwV1ApiGetTransitGatewayModelAttachmentResponseModel.md) |  | [optional] 
**Options** | Pointer to [**NullableOptionResponseModel**](OptionResponseModel.md) |  | [optional] 
**ProvisioningStatus** | Pointer to [**NullableTGWProvisioningStatus**](TGWProvisioningStatus.md) |  | [optional] 
**ProjectId** | Pointer to **NullableString** |  | [optional] 
**ProjectName** | Pointer to **NullableString** |  | [optional] 
**OwnerProjectId** | Pointer to **NullableString** |  | [optional] 
**OwnerProjectName** | Pointer to **NullableString** |  | [optional] 
**RouteTables** | Pointer to [**[]BnsTgwV1ApiGetTransitGatewayModelRouteTableResponseModel**](BnsTgwV1ApiGetTransitGatewayModelRouteTableResponseModel.md) |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewBnsTgwV1ApiGetTransitGatewayModelTgwResponseModel

`func NewBnsTgwV1ApiGetTransitGatewayModelTgwResponseModel() *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel`

NewBnsTgwV1ApiGetTransitGatewayModelTgwResponseModel instantiates a new BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsTgwV1ApiGetTransitGatewayModelTgwResponseModelWithDefaults

`func NewBnsTgwV1ApiGetTransitGatewayModelTgwResponseModelWithDefaults() *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel`

NewBnsTgwV1ApiGetTransitGatewayModelTgwResponseModelWithDefaults instantiates a new BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRegion

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) GetRegion() Region`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) GetRegionOk() (*Region, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) SetRegion(v Region)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetIsShared

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) GetIsShared() bool`

GetIsShared returns the IsShared field if non-nil, zero value otherwise.

### GetIsSharedOk

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) GetIsSharedOk() (*bool, bool)`

GetIsSharedOk returns a tuple with the IsShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShared

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) SetIsShared(v bool)`

SetIsShared sets IsShared field to given value.

### HasIsShared

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) HasIsShared() bool`

HasIsShared returns a boolean if a field has been set.

### SetIsSharedNil

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) SetIsSharedNil(b bool)`

 SetIsSharedNil sets the value for IsShared to be an explicit nil

### UnsetIsShared
`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) UnsetIsShared()`

UnsetIsShared ensures that no value is present for IsShared, not even an explicit nil
### GetAttachments

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) GetAttachments() []BnsTgwV1ApiGetTransitGatewayModelAttachmentResponseModel`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) GetAttachmentsOk() (*[]BnsTgwV1ApiGetTransitGatewayModelAttachmentResponseModel, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) SetAttachments(v []BnsTgwV1ApiGetTransitGatewayModelAttachmentResponseModel)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### SetAttachmentsNil

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) SetAttachmentsNil(b bool)`

 SetAttachmentsNil sets the value for Attachments to be an explicit nil

### UnsetAttachments
`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) UnsetAttachments()`

UnsetAttachments ensures that no value is present for Attachments, not even an explicit nil
### GetOptions

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) GetOptions() OptionResponseModel`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) GetOptionsOk() (*OptionResponseModel, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) SetOptions(v OptionResponseModel)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetProvisioningStatus

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) GetProvisioningStatus() TGWProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) GetProvisioningStatusOk() (*TGWProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) SetProvisioningStatus(v TGWProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.

### HasProvisioningStatus

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) HasProvisioningStatus() bool`

HasProvisioningStatus returns a boolean if a field has been set.

### SetProvisioningStatusNil

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) SetProvisioningStatusNil(b bool)`

 SetProvisioningStatusNil sets the value for ProvisioningStatus to be an explicit nil

### UnsetProvisioningStatus
`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) UnsetProvisioningStatus()`

UnsetProvisioningStatus ensures that no value is present for ProvisioningStatus, not even an explicit nil
### GetProjectId

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetProjectName

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### SetProjectNameNil

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) SetProjectNameNil(b bool)`

 SetProjectNameNil sets the value for ProjectName to be an explicit nil

### UnsetProjectName
`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) UnsetProjectName()`

UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil
### GetOwnerProjectId

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) GetOwnerProjectId() string`

GetOwnerProjectId returns the OwnerProjectId field if non-nil, zero value otherwise.

### GetOwnerProjectIdOk

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) GetOwnerProjectIdOk() (*string, bool)`

GetOwnerProjectIdOk returns a tuple with the OwnerProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerProjectId

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) SetOwnerProjectId(v string)`

SetOwnerProjectId sets OwnerProjectId field to given value.

### HasOwnerProjectId

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) HasOwnerProjectId() bool`

HasOwnerProjectId returns a boolean if a field has been set.

### SetOwnerProjectIdNil

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) SetOwnerProjectIdNil(b bool)`

 SetOwnerProjectIdNil sets the value for OwnerProjectId to be an explicit nil

### UnsetOwnerProjectId
`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) UnsetOwnerProjectId()`

UnsetOwnerProjectId ensures that no value is present for OwnerProjectId, not even an explicit nil
### GetOwnerProjectName

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) GetOwnerProjectName() string`

GetOwnerProjectName returns the OwnerProjectName field if non-nil, zero value otherwise.

### GetOwnerProjectNameOk

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) GetOwnerProjectNameOk() (*string, bool)`

GetOwnerProjectNameOk returns a tuple with the OwnerProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerProjectName

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) SetOwnerProjectName(v string)`

SetOwnerProjectName sets OwnerProjectName field to given value.

### HasOwnerProjectName

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) HasOwnerProjectName() bool`

HasOwnerProjectName returns a boolean if a field has been set.

### SetOwnerProjectNameNil

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) SetOwnerProjectNameNil(b bool)`

 SetOwnerProjectNameNil sets the value for OwnerProjectName to be an explicit nil

### UnsetOwnerProjectName
`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) UnsetOwnerProjectName()`

UnsetOwnerProjectName ensures that no value is present for OwnerProjectName, not even an explicit nil
### GetRouteTables

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) GetRouteTables() []BnsTgwV1ApiGetTransitGatewayModelRouteTableResponseModel`

GetRouteTables returns the RouteTables field if non-nil, zero value otherwise.

### GetRouteTablesOk

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) GetRouteTablesOk() (*[]BnsTgwV1ApiGetTransitGatewayModelRouteTableResponseModel, bool)`

GetRouteTablesOk returns a tuple with the RouteTables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteTables

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) SetRouteTables(v []BnsTgwV1ApiGetTransitGatewayModelRouteTableResponseModel)`

SetRouteTables sets RouteTables field to given value.

### HasRouteTables

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) HasRouteTables() bool`

HasRouteTables returns a boolean if a field has been set.

### SetRouteTablesNil

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) SetRouteTablesNil(b bool)`

 SetRouteTablesNil sets the value for RouteTables to be an explicit nil

### UnsetRouteTables
`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) UnsetRouteTables()`

UnsetRouteTables ensures that no value is present for RouteTables, not even an explicit nil
### GetCreatedAt

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *BnsTgwV1ApiGetTransitGatewayModelTgwResponseModel) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


