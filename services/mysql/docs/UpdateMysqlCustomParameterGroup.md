# UpdateMysqlCustomParameterGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplyMode** | Pointer to [**NullableApplyMode**](ApplyMode.md) |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Parameters** | Pointer to [**[]DataUpdateParameterRequest**](DataUpdateParameterRequest.md) |  | [optional] 

## Methods

### NewUpdateMysqlCustomParameterGroup

`func NewUpdateMysqlCustomParameterGroup() *UpdateMysqlCustomParameterGroup`

NewUpdateMysqlCustomParameterGroup instantiates a new UpdateMysqlCustomParameterGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMysqlCustomParameterGroupWithDefaults

`func NewUpdateMysqlCustomParameterGroupWithDefaults() *UpdateMysqlCustomParameterGroup`

NewUpdateMysqlCustomParameterGroupWithDefaults instantiates a new UpdateMysqlCustomParameterGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplyMode

`func (o *UpdateMysqlCustomParameterGroup) GetApplyMode() ApplyMode`

GetApplyMode returns the ApplyMode field if non-nil, zero value otherwise.

### GetApplyModeOk

`func (o *UpdateMysqlCustomParameterGroup) GetApplyModeOk() (*ApplyMode, bool)`

GetApplyModeOk returns a tuple with the ApplyMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyMode

`func (o *UpdateMysqlCustomParameterGroup) SetApplyMode(v ApplyMode)`

SetApplyMode sets ApplyMode field to given value.

### HasApplyMode

`func (o *UpdateMysqlCustomParameterGroup) HasApplyMode() bool`

HasApplyMode returns a boolean if a field has been set.

### SetApplyModeNil

`func (o *UpdateMysqlCustomParameterGroup) SetApplyModeNil(b bool)`

 SetApplyModeNil sets the value for ApplyMode to be an explicit nil

### UnsetApplyMode
`func (o *UpdateMysqlCustomParameterGroup) UnsetApplyMode()`

UnsetApplyMode ensures that no value is present for ApplyMode, not even an explicit nil
### GetDescription

`func (o *UpdateMysqlCustomParameterGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateMysqlCustomParameterGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateMysqlCustomParameterGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateMysqlCustomParameterGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateMysqlCustomParameterGroup) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateMysqlCustomParameterGroup) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetParameters

`func (o *UpdateMysqlCustomParameterGroup) GetParameters() []DataUpdateParameterRequest`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *UpdateMysqlCustomParameterGroup) GetParametersOk() (*[]DataUpdateParameterRequest, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *UpdateMysqlCustomParameterGroup) SetParameters(v []DataUpdateParameterRequest)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *UpdateMysqlCustomParameterGroup) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *UpdateMysqlCustomParameterGroup) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *UpdateMysqlCustomParameterGroup) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


