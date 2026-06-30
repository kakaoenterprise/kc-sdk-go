# MysqlV1ApiUpdateMysqlCustomParameterGroupModelCustomParameterGroupRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplyMode** | Pointer to [**NullableApplyMode**](ApplyMode.md) |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Parameters** | Pointer to [**[]DataUpdateParameterRequestModel**](DataUpdateParameterRequestModel.md) |  | [optional] 

## Methods

### NewMysqlV1ApiUpdateMysqlCustomParameterGroupModelCustomParameterGroupRequestModel

`func NewMysqlV1ApiUpdateMysqlCustomParameterGroupModelCustomParameterGroupRequestModel() *MysqlV1ApiUpdateMysqlCustomParameterGroupModelCustomParameterGroupRequestModel`

NewMysqlV1ApiUpdateMysqlCustomParameterGroupModelCustomParameterGroupRequestModel instantiates a new MysqlV1ApiUpdateMysqlCustomParameterGroupModelCustomParameterGroupRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMysqlV1ApiUpdateMysqlCustomParameterGroupModelCustomParameterGroupRequestModelWithDefaults

`func NewMysqlV1ApiUpdateMysqlCustomParameterGroupModelCustomParameterGroupRequestModelWithDefaults() *MysqlV1ApiUpdateMysqlCustomParameterGroupModelCustomParameterGroupRequestModel`

NewMysqlV1ApiUpdateMysqlCustomParameterGroupModelCustomParameterGroupRequestModelWithDefaults instantiates a new MysqlV1ApiUpdateMysqlCustomParameterGroupModelCustomParameterGroupRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplyMode

`func (o *MysqlV1ApiUpdateMysqlCustomParameterGroupModelCustomParameterGroupRequestModel) GetApplyMode() ApplyMode`

GetApplyMode returns the ApplyMode field if non-nil, zero value otherwise.

### GetApplyModeOk

`func (o *MysqlV1ApiUpdateMysqlCustomParameterGroupModelCustomParameterGroupRequestModel) GetApplyModeOk() (*ApplyMode, bool)`

GetApplyModeOk returns a tuple with the ApplyMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyMode

`func (o *MysqlV1ApiUpdateMysqlCustomParameterGroupModelCustomParameterGroupRequestModel) SetApplyMode(v ApplyMode)`

SetApplyMode sets ApplyMode field to given value.

### HasApplyMode

`func (o *MysqlV1ApiUpdateMysqlCustomParameterGroupModelCustomParameterGroupRequestModel) HasApplyMode() bool`

HasApplyMode returns a boolean if a field has been set.

### SetApplyModeNil

`func (o *MysqlV1ApiUpdateMysqlCustomParameterGroupModelCustomParameterGroupRequestModel) SetApplyModeNil(b bool)`

 SetApplyModeNil sets the value for ApplyMode to be an explicit nil

### UnsetApplyMode
`func (o *MysqlV1ApiUpdateMysqlCustomParameterGroupModelCustomParameterGroupRequestModel) UnsetApplyMode()`

UnsetApplyMode ensures that no value is present for ApplyMode, not even an explicit nil
### GetDescription

`func (o *MysqlV1ApiUpdateMysqlCustomParameterGroupModelCustomParameterGroupRequestModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MysqlV1ApiUpdateMysqlCustomParameterGroupModelCustomParameterGroupRequestModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MysqlV1ApiUpdateMysqlCustomParameterGroupModelCustomParameterGroupRequestModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MysqlV1ApiUpdateMysqlCustomParameterGroupModelCustomParameterGroupRequestModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MysqlV1ApiUpdateMysqlCustomParameterGroupModelCustomParameterGroupRequestModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MysqlV1ApiUpdateMysqlCustomParameterGroupModelCustomParameterGroupRequestModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetParameters

`func (o *MysqlV1ApiUpdateMysqlCustomParameterGroupModelCustomParameterGroupRequestModel) GetParameters() []DataUpdateParameterRequestModel`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *MysqlV1ApiUpdateMysqlCustomParameterGroupModelCustomParameterGroupRequestModel) GetParametersOk() (*[]DataUpdateParameterRequestModel, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *MysqlV1ApiUpdateMysqlCustomParameterGroupModelCustomParameterGroupRequestModel) SetParameters(v []DataUpdateParameterRequestModel)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *MysqlV1ApiUpdateMysqlCustomParameterGroupModelCustomParameterGroupRequestModel) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *MysqlV1ApiUpdateMysqlCustomParameterGroupModelCustomParameterGroupRequestModel) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *MysqlV1ApiUpdateMysqlCustomParameterGroupModelCustomParameterGroupRequestModel) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


