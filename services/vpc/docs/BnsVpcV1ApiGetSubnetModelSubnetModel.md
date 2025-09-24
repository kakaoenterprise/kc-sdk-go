# BnsVpcV1ApiGetSubnetModelSubnetModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 서브넷 ID | 
**Name** | Pointer to **NullableString** |  | [optional] 
**IsShared** | Pointer to **NullableBool** |  | [optional] 
**AvailabilityZone** | Pointer to [**NullableAvailabilityZone**](AvailabilityZone.md) |  | [optional] 
**CidrBlock** | Pointer to **NullableString** |  | [optional] 
**ProjectId** | Pointer to **NullableString** |  | [optional] 
**OperatingStatus** | Pointer to [**NullableSubnetOperatingStatus**](SubnetOperatingStatus.md) |  | [optional] 
**ProvisioningStatus** | Pointer to [**NullableProvisioningStatus**](ProvisioningStatus.md) |  | [optional] 
**VpcId** | Pointer to **NullableString** |  | [optional] 
**VpcName** | Pointer to **NullableString** |  | [optional] 
**ProjectName** | Pointer to **NullableString** |  | [optional] 
**OwnerProjectId** | Pointer to **NullableString** |  | [optional] 
**RouteTableId** | Pointer to **NullableString** |  | [optional] 
**RouteTableName** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewBnsVpcV1ApiGetSubnetModelSubnetModel

`func NewBnsVpcV1ApiGetSubnetModelSubnetModel(id string, ) *BnsVpcV1ApiGetSubnetModelSubnetModel`

NewBnsVpcV1ApiGetSubnetModelSubnetModel instantiates a new BnsVpcV1ApiGetSubnetModelSubnetModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsVpcV1ApiGetSubnetModelSubnetModelWithDefaults

`func NewBnsVpcV1ApiGetSubnetModelSubnetModelWithDefaults() *BnsVpcV1ApiGetSubnetModelSubnetModel`

NewBnsVpcV1ApiGetSubnetModelSubnetModelWithDefaults instantiates a new BnsVpcV1ApiGetSubnetModelSubnetModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIsShared

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) GetIsShared() bool`

GetIsShared returns the IsShared field if non-nil, zero value otherwise.

### GetIsSharedOk

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) GetIsSharedOk() (*bool, bool)`

GetIsSharedOk returns a tuple with the IsShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShared

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) SetIsShared(v bool)`

SetIsShared sets IsShared field to given value.

### HasIsShared

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) HasIsShared() bool`

HasIsShared returns a boolean if a field has been set.

### SetIsSharedNil

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) SetIsSharedNil(b bool)`

 SetIsSharedNil sets the value for IsShared to be an explicit nil

### UnsetIsShared
`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) UnsetIsShared()`

UnsetIsShared ensures that no value is present for IsShared, not even an explicit nil
### GetAvailabilityZone

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) GetAvailabilityZone() AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) GetAvailabilityZoneOk() (*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) SetAvailabilityZone(v AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### SetAvailabilityZoneNil

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) SetAvailabilityZoneNil(b bool)`

 SetAvailabilityZoneNil sets the value for AvailabilityZone to be an explicit nil

### UnsetAvailabilityZone
`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) UnsetAvailabilityZone()`

UnsetAvailabilityZone ensures that no value is present for AvailabilityZone, not even an explicit nil
### GetCidrBlock

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) GetCidrBlock() string`

GetCidrBlock returns the CidrBlock field if non-nil, zero value otherwise.

### GetCidrBlockOk

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) GetCidrBlockOk() (*string, bool)`

GetCidrBlockOk returns a tuple with the CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlock

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) SetCidrBlock(v string)`

SetCidrBlock sets CidrBlock field to given value.

### HasCidrBlock

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) HasCidrBlock() bool`

HasCidrBlock returns a boolean if a field has been set.

### SetCidrBlockNil

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) SetCidrBlockNil(b bool)`

 SetCidrBlockNil sets the value for CidrBlock to be an explicit nil

### UnsetCidrBlock
`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) UnsetCidrBlock()`

UnsetCidrBlock ensures that no value is present for CidrBlock, not even an explicit nil
### GetProjectId

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetOperatingStatus

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) GetOperatingStatus() SubnetOperatingStatus`

GetOperatingStatus returns the OperatingStatus field if non-nil, zero value otherwise.

### GetOperatingStatusOk

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) GetOperatingStatusOk() (*SubnetOperatingStatus, bool)`

GetOperatingStatusOk returns a tuple with the OperatingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingStatus

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) SetOperatingStatus(v SubnetOperatingStatus)`

SetOperatingStatus sets OperatingStatus field to given value.

### HasOperatingStatus

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) HasOperatingStatus() bool`

HasOperatingStatus returns a boolean if a field has been set.

### SetOperatingStatusNil

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) SetOperatingStatusNil(b bool)`

 SetOperatingStatusNil sets the value for OperatingStatus to be an explicit nil

### UnsetOperatingStatus
`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) UnsetOperatingStatus()`

UnsetOperatingStatus ensures that no value is present for OperatingStatus, not even an explicit nil
### GetProvisioningStatus

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) GetProvisioningStatus() ProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) GetProvisioningStatusOk() (*ProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) SetProvisioningStatus(v ProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.

### HasProvisioningStatus

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) HasProvisioningStatus() bool`

HasProvisioningStatus returns a boolean if a field has been set.

### SetProvisioningStatusNil

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) SetProvisioningStatusNil(b bool)`

 SetProvisioningStatusNil sets the value for ProvisioningStatus to be an explicit nil

### UnsetProvisioningStatus
`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) UnsetProvisioningStatus()`

UnsetProvisioningStatus ensures that no value is present for ProvisioningStatus, not even an explicit nil
### GetVpcId

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### SetVpcIdNil

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) SetVpcIdNil(b bool)`

 SetVpcIdNil sets the value for VpcId to be an explicit nil

### UnsetVpcId
`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) UnsetVpcId()`

UnsetVpcId ensures that no value is present for VpcId, not even an explicit nil
### GetVpcName

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) GetVpcName() string`

GetVpcName returns the VpcName field if non-nil, zero value otherwise.

### GetVpcNameOk

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) GetVpcNameOk() (*string, bool)`

GetVpcNameOk returns a tuple with the VpcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcName

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) SetVpcName(v string)`

SetVpcName sets VpcName field to given value.

### HasVpcName

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) HasVpcName() bool`

HasVpcName returns a boolean if a field has been set.

### SetVpcNameNil

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) SetVpcNameNil(b bool)`

 SetVpcNameNil sets the value for VpcName to be an explicit nil

### UnsetVpcName
`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) UnsetVpcName()`

UnsetVpcName ensures that no value is present for VpcName, not even an explicit nil
### GetProjectName

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### SetProjectNameNil

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) SetProjectNameNil(b bool)`

 SetProjectNameNil sets the value for ProjectName to be an explicit nil

### UnsetProjectName
`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) UnsetProjectName()`

UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil
### GetOwnerProjectId

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) GetOwnerProjectId() string`

GetOwnerProjectId returns the OwnerProjectId field if non-nil, zero value otherwise.

### GetOwnerProjectIdOk

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) GetOwnerProjectIdOk() (*string, bool)`

GetOwnerProjectIdOk returns a tuple with the OwnerProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerProjectId

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) SetOwnerProjectId(v string)`

SetOwnerProjectId sets OwnerProjectId field to given value.

### HasOwnerProjectId

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) HasOwnerProjectId() bool`

HasOwnerProjectId returns a boolean if a field has been set.

### SetOwnerProjectIdNil

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) SetOwnerProjectIdNil(b bool)`

 SetOwnerProjectIdNil sets the value for OwnerProjectId to be an explicit nil

### UnsetOwnerProjectId
`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) UnsetOwnerProjectId()`

UnsetOwnerProjectId ensures that no value is present for OwnerProjectId, not even an explicit nil
### GetRouteTableId

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) GetRouteTableId() string`

GetRouteTableId returns the RouteTableId field if non-nil, zero value otherwise.

### GetRouteTableIdOk

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) GetRouteTableIdOk() (*string, bool)`

GetRouteTableIdOk returns a tuple with the RouteTableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteTableId

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) SetRouteTableId(v string)`

SetRouteTableId sets RouteTableId field to given value.

### HasRouteTableId

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) HasRouteTableId() bool`

HasRouteTableId returns a boolean if a field has been set.

### SetRouteTableIdNil

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) SetRouteTableIdNil(b bool)`

 SetRouteTableIdNil sets the value for RouteTableId to be an explicit nil

### UnsetRouteTableId
`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) UnsetRouteTableId()`

UnsetRouteTableId ensures that no value is present for RouteTableId, not even an explicit nil
### GetRouteTableName

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) GetRouteTableName() string`

GetRouteTableName returns the RouteTableName field if non-nil, zero value otherwise.

### GetRouteTableNameOk

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) GetRouteTableNameOk() (*string, bool)`

GetRouteTableNameOk returns a tuple with the RouteTableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteTableName

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) SetRouteTableName(v string)`

SetRouteTableName sets RouteTableName field to given value.

### HasRouteTableName

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) HasRouteTableName() bool`

HasRouteTableName returns a boolean if a field has been set.

### SetRouteTableNameNil

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) SetRouteTableNameNil(b bool)`

 SetRouteTableNameNil sets the value for RouteTableName to be an explicit nil

### UnsetRouteTableName
`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) UnsetRouteTableName()`

UnsetRouteTableName ensures that no value is present for RouteTableName, not even an explicit nil
### GetCreatedAt

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *BnsVpcV1ApiGetSubnetModelSubnetModel) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


