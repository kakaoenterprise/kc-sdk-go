# CreateSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsIncremental** | **bool** | 증분 스냅샷 여부 - 💡 스냅샷 일정을 통해 전체 스냅샷을 보유하고 있더라도, 수동으로 스냅샷을 생성할 경우에는 최초 1번의 전체 스냅샷 생성이 필요합니다. | 
**Name** | **string** | 생성할 스냅샷의 이름 | 
**Description** | Pointer to **NullableString** | 스냅샷에 대한 설명 | [optional] 

## Methods

### NewCreateSnapshot

`func NewCreateSnapshot(isIncremental bool, name string, ) *CreateSnapshot`

NewCreateSnapshot instantiates a new CreateSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSnapshotWithDefaults

`func NewCreateSnapshotWithDefaults() *CreateSnapshot`

NewCreateSnapshotWithDefaults instantiates a new CreateSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsIncremental

`func (o *CreateSnapshot) GetIsIncremental() bool`

GetIsIncremental returns the IsIncremental field if non-nil, zero value otherwise.

### GetIsIncrementalOk

`func (o *CreateSnapshot) GetIsIncrementalOk() (*bool, bool)`

GetIsIncrementalOk returns a tuple with the IsIncremental field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIncremental

`func (o *CreateSnapshot) SetIsIncremental(v bool)`

SetIsIncremental sets IsIncremental field to given value.


### GetName

`func (o *CreateSnapshot) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSnapshot) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSnapshot) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateSnapshot) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateSnapshot) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateSnapshot) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateSnapshot) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateSnapshot) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateSnapshot) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


