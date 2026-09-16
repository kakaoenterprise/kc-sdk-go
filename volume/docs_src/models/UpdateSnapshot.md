# UpdateSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | 변경할 스냅샷의 이름 | [optional] 
**Description** | Pointer to **NullableString** | 스냅샷에 대한 설명 - 필드를 사용하지 않거나 &#x60;null&#x60;이면 기존값 유지 | [optional] 

## Methods

### NewUpdateSnapshot

`func NewUpdateSnapshot() *UpdateSnapshot`

NewUpdateSnapshot instantiates a new UpdateSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSnapshotWithDefaults

`func NewUpdateSnapshotWithDefaults() *UpdateSnapshot`

NewUpdateSnapshotWithDefaults instantiates a new UpdateSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateSnapshot) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateSnapshot) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateSnapshot) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateSnapshot) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateSnapshot) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateSnapshot) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *UpdateSnapshot) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateSnapshot) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateSnapshot) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateSnapshot) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateSnapshot) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateSnapshot) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


