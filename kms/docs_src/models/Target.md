# Target

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignmentType** | Pointer to **NullableString** | 접근 제어 대상의 IAM 할당 유형 | [optional] 
**Id** | Pointer to **NullableString** | 접근 제어 대상의 고유 ID | [optional] 
**Name** | Pointer to **NullableString** | 접근 제어 대상 이름 | [optional] 
**IsExists** | Pointer to **NullableBool** | 접근 제어 대상이 IAM에 존재하는지 여부 | [optional] 
**CreatedBy** | Pointer to **NullableString** | 접근 제어 대상을 추가한 사용자 또는 주체 | [optional] 
**CreatedAt** | Pointer to **NullableString** | 접근 제어 대상이 추가된 일시 | [optional] 

## Methods

### NewTarget

`func NewTarget() *Target`

NewTarget instantiates a new Target object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetWithDefaults

`func NewTargetWithDefaults() *Target`

NewTargetWithDefaults instantiates a new Target object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignmentType

`func (o *Target) GetAssignmentType() string`

GetAssignmentType returns the AssignmentType field if non-nil, zero value otherwise.

### GetAssignmentTypeOk

`func (o *Target) GetAssignmentTypeOk() (*string, bool)`

GetAssignmentTypeOk returns a tuple with the AssignmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentType

`func (o *Target) SetAssignmentType(v string)`

SetAssignmentType sets AssignmentType field to given value.

### HasAssignmentType

`func (o *Target) HasAssignmentType() bool`

HasAssignmentType returns a boolean if a field has been set.

### SetAssignmentTypeNil

`func (o *Target) SetAssignmentTypeNil(b bool)`

 SetAssignmentTypeNil sets the value for AssignmentType to be an explicit nil

### UnsetAssignmentType
`func (o *Target) UnsetAssignmentType()`

UnsetAssignmentType ensures that no value is present for AssignmentType, not even an explicit nil
### GetId

`func (o *Target) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Target) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Target) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Target) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Target) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Target) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *Target) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Target) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Target) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Target) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Target) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Target) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIsExists

`func (o *Target) GetIsExists() bool`

GetIsExists returns the IsExists field if non-nil, zero value otherwise.

### GetIsExistsOk

`func (o *Target) GetIsExistsOk() (*bool, bool)`

GetIsExistsOk returns a tuple with the IsExists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExists

`func (o *Target) SetIsExists(v bool)`

SetIsExists sets IsExists field to given value.

### HasIsExists

`func (o *Target) HasIsExists() bool`

HasIsExists returns a boolean if a field has been set.

### SetIsExistsNil

`func (o *Target) SetIsExistsNil(b bool)`

 SetIsExistsNil sets the value for IsExists to be an explicit nil

### UnsetIsExists
`func (o *Target) UnsetIsExists()`

UnsetIsExists ensures that no value is present for IsExists, not even an explicit nil
### GetCreatedBy

`func (o *Target) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Target) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Target) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Target) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *Target) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *Target) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetCreatedAt

`func (o *Target) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Target) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Target) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Target) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *Target) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *Target) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


