# DefaultParameterGroupEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **NullableString** | 이벤트가 발생한 시각 - ISO 8601 형식 - UTC 기준 | [optional] 
**Description** | Pointer to **NullableString** | 이벤트에 대한 상세 설명 | [optional] 
**Name** | Pointer to **NullableString** | 이벤트 이름 | [optional] 

## Methods

### NewDefaultParameterGroupEvent

`func NewDefaultParameterGroupEvent() *DefaultParameterGroupEvent`

NewDefaultParameterGroupEvent instantiates a new DefaultParameterGroupEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefaultParameterGroupEventWithDefaults

`func NewDefaultParameterGroupEventWithDefaults() *DefaultParameterGroupEvent`

NewDefaultParameterGroupEventWithDefaults instantiates a new DefaultParameterGroupEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *DefaultParameterGroupEvent) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DefaultParameterGroupEvent) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DefaultParameterGroupEvent) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DefaultParameterGroupEvent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *DefaultParameterGroupEvent) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *DefaultParameterGroupEvent) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDescription

`func (o *DefaultParameterGroupEvent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DefaultParameterGroupEvent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DefaultParameterGroupEvent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DefaultParameterGroupEvent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *DefaultParameterGroupEvent) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DefaultParameterGroupEvent) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetName

`func (o *DefaultParameterGroupEvent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DefaultParameterGroupEvent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DefaultParameterGroupEvent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DefaultParameterGroupEvent) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *DefaultParameterGroupEvent) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DefaultParameterGroupEvent) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


