# VpcRouteTable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | 라우팅 테이블의 ID | [optional] 
**Name** | Pointer to **NullableString** | 라우팅 테이블 이름 | [optional] 
**Description** | Pointer to **NullableString** | 라우팅 테이블에 대한 설명 | [optional] 
**Associations** | Pointer to [**[]VpcAssociation**](VpcAssociation.md) | 해당 라우팅 테이블에 연결된 서브넷 연결 정보 목록 | [optional] 
**Routes** | Pointer to [**[]VpcRoute**](VpcRoute.md) | 라우팅 테이블에 정의된 라우팅 경로 목록 | [optional] 
**VpcId** | Pointer to **NullableString** | 라우팅 테이블이 연결된 VPC의 ID | [optional] 
**ProvisioningStatus** | Pointer to [**NullableProvisioningStatus**](ProvisioningStatus.md) | 프로비저닝 상태 | [optional] 
**VpcName** | Pointer to **NullableString** | 라우팅 테이블이 연결된 VPC의 이름 | [optional] 
**VpcProvisioningStatus** | Pointer to [**NullableProvisioningStatus**](ProvisioningStatus.md) | 연결된 VPC의 상태 | [optional] 
**ProjectId** | Pointer to **NullableString** | 라우팅 테이블이 속한 프로젝트의 ID | [optional] 
**ProjectName** | Pointer to **NullableString** | 프로젝트 이름 | [optional] 
**IsMain** | Pointer to **NullableBool** | 기본(main) 라우팅 테이블인지 여부 | [optional] 
**CreatedAt** | Pointer to **NullableTime** | 리소스가 생성된 시간 - ISO 8601 형식 - UTC 기준 | [optional] 
**UpdatedAt** | Pointer to **NullableTime** | 리소스가 마지막으로 수정된 시간 - ISO 8601 형식 - UTC 기준 | [optional] 

## Methods

### NewVpcRouteTable

`func NewVpcRouteTable() *VpcRouteTable`

NewVpcRouteTable instantiates a new VpcRouteTable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpcRouteTableWithDefaults

`func NewVpcRouteTableWithDefaults() *VpcRouteTable`

NewVpcRouteTableWithDefaults instantiates a new VpcRouteTable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VpcRouteTable) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VpcRouteTable) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VpcRouteTable) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VpcRouteTable) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *VpcRouteTable) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *VpcRouteTable) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *VpcRouteTable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VpcRouteTable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VpcRouteTable) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VpcRouteTable) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *VpcRouteTable) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *VpcRouteTable) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *VpcRouteTable) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VpcRouteTable) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VpcRouteTable) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VpcRouteTable) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *VpcRouteTable) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *VpcRouteTable) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetAssociations

`func (o *VpcRouteTable) GetAssociations() []VpcAssociation`

GetAssociations returns the Associations field if non-nil, zero value otherwise.

### GetAssociationsOk

`func (o *VpcRouteTable) GetAssociationsOk() (*[]VpcAssociation, bool)`

GetAssociationsOk returns a tuple with the Associations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociations

`func (o *VpcRouteTable) SetAssociations(v []VpcAssociation)`

SetAssociations sets Associations field to given value.

### HasAssociations

`func (o *VpcRouteTable) HasAssociations() bool`

HasAssociations returns a boolean if a field has been set.

### SetAssociationsNil

`func (o *VpcRouteTable) SetAssociationsNil(b bool)`

 SetAssociationsNil sets the value for Associations to be an explicit nil

### UnsetAssociations
`func (o *VpcRouteTable) UnsetAssociations()`

UnsetAssociations ensures that no value is present for Associations, not even an explicit nil
### GetRoutes

`func (o *VpcRouteTable) GetRoutes() []VpcRoute`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *VpcRouteTable) GetRoutesOk() (*[]VpcRoute, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *VpcRouteTable) SetRoutes(v []VpcRoute)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *VpcRouteTable) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.

### SetRoutesNil

`func (o *VpcRouteTable) SetRoutesNil(b bool)`

 SetRoutesNil sets the value for Routes to be an explicit nil

### UnsetRoutes
`func (o *VpcRouteTable) UnsetRoutes()`

UnsetRoutes ensures that no value is present for Routes, not even an explicit nil
### GetVpcId

`func (o *VpcRouteTable) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *VpcRouteTable) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *VpcRouteTable) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *VpcRouteTable) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### SetVpcIdNil

`func (o *VpcRouteTable) SetVpcIdNil(b bool)`

 SetVpcIdNil sets the value for VpcId to be an explicit nil

### UnsetVpcId
`func (o *VpcRouteTable) UnsetVpcId()`

UnsetVpcId ensures that no value is present for VpcId, not even an explicit nil
### GetProvisioningStatus

`func (o *VpcRouteTable) GetProvisioningStatus() ProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *VpcRouteTable) GetProvisioningStatusOk() (*ProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *VpcRouteTable) SetProvisioningStatus(v ProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.

### HasProvisioningStatus

`func (o *VpcRouteTable) HasProvisioningStatus() bool`

HasProvisioningStatus returns a boolean if a field has been set.

### SetProvisioningStatusNil

`func (o *VpcRouteTable) SetProvisioningStatusNil(b bool)`

 SetProvisioningStatusNil sets the value for ProvisioningStatus to be an explicit nil

### UnsetProvisioningStatus
`func (o *VpcRouteTable) UnsetProvisioningStatus()`

UnsetProvisioningStatus ensures that no value is present for ProvisioningStatus, not even an explicit nil
### GetVpcName

`func (o *VpcRouteTable) GetVpcName() string`

GetVpcName returns the VpcName field if non-nil, zero value otherwise.

### GetVpcNameOk

`func (o *VpcRouteTable) GetVpcNameOk() (*string, bool)`

GetVpcNameOk returns a tuple with the VpcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcName

`func (o *VpcRouteTable) SetVpcName(v string)`

SetVpcName sets VpcName field to given value.

### HasVpcName

`func (o *VpcRouteTable) HasVpcName() bool`

HasVpcName returns a boolean if a field has been set.

### SetVpcNameNil

`func (o *VpcRouteTable) SetVpcNameNil(b bool)`

 SetVpcNameNil sets the value for VpcName to be an explicit nil

### UnsetVpcName
`func (o *VpcRouteTable) UnsetVpcName()`

UnsetVpcName ensures that no value is present for VpcName, not even an explicit nil
### GetVpcProvisioningStatus

`func (o *VpcRouteTable) GetVpcProvisioningStatus() ProvisioningStatus`

GetVpcProvisioningStatus returns the VpcProvisioningStatus field if non-nil, zero value otherwise.

### GetVpcProvisioningStatusOk

`func (o *VpcRouteTable) GetVpcProvisioningStatusOk() (*ProvisioningStatus, bool)`

GetVpcProvisioningStatusOk returns a tuple with the VpcProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcProvisioningStatus

`func (o *VpcRouteTable) SetVpcProvisioningStatus(v ProvisioningStatus)`

SetVpcProvisioningStatus sets VpcProvisioningStatus field to given value.

### HasVpcProvisioningStatus

`func (o *VpcRouteTable) HasVpcProvisioningStatus() bool`

HasVpcProvisioningStatus returns a boolean if a field has been set.

### SetVpcProvisioningStatusNil

`func (o *VpcRouteTable) SetVpcProvisioningStatusNil(b bool)`

 SetVpcProvisioningStatusNil sets the value for VpcProvisioningStatus to be an explicit nil

### UnsetVpcProvisioningStatus
`func (o *VpcRouteTable) UnsetVpcProvisioningStatus()`

UnsetVpcProvisioningStatus ensures that no value is present for VpcProvisioningStatus, not even an explicit nil
### GetProjectId

`func (o *VpcRouteTable) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *VpcRouteTable) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *VpcRouteTable) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *VpcRouteTable) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *VpcRouteTable) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *VpcRouteTable) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetProjectName

`func (o *VpcRouteTable) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *VpcRouteTable) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *VpcRouteTable) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *VpcRouteTable) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### SetProjectNameNil

`func (o *VpcRouteTable) SetProjectNameNil(b bool)`

 SetProjectNameNil sets the value for ProjectName to be an explicit nil

### UnsetProjectName
`func (o *VpcRouteTable) UnsetProjectName()`

UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil
### GetIsMain

`func (o *VpcRouteTable) GetIsMain() bool`

GetIsMain returns the IsMain field if non-nil, zero value otherwise.

### GetIsMainOk

`func (o *VpcRouteTable) GetIsMainOk() (*bool, bool)`

GetIsMainOk returns a tuple with the IsMain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMain

`func (o *VpcRouteTable) SetIsMain(v bool)`

SetIsMain sets IsMain field to given value.

### HasIsMain

`func (o *VpcRouteTable) HasIsMain() bool`

HasIsMain returns a boolean if a field has been set.

### SetIsMainNil

`func (o *VpcRouteTable) SetIsMainNil(b bool)`

 SetIsMainNil sets the value for IsMain to be an explicit nil

### UnsetIsMain
`func (o *VpcRouteTable) UnsetIsMain()`

UnsetIsMain ensures that no value is present for IsMain, not even an explicit nil
### GetCreatedAt

`func (o *VpcRouteTable) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VpcRouteTable) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VpcRouteTable) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VpcRouteTable) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *VpcRouteTable) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *VpcRouteTable) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *VpcRouteTable) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VpcRouteTable) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VpcRouteTable) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *VpcRouteTable) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *VpcRouteTable) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *VpcRouteTable) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


