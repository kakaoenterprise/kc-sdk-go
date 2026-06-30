# DataDetailParametersResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | 파라미터 이름 | 
**DataType** | [**DataType**](DataType.md) | 파라미터 데이터 타입 | 
**DefaultParameterValue** | Pointer to **NullableString** |  | [optional] 
**IsEditable** | **bool** | 파라미터 수정 가능 여부 | 
**ParameterType** | Pointer to [**NullableParameterType**](ParameterType.md) |  | [optional] 
**Value** | Pointer to **NullableString** |  | [optional] 
**IsRequired** | **bool** | 필수 파라미터 여부 | 
**ValidationValueFormat** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDataDetailParametersResponseModel

`func NewDataDetailParametersResponseModel(key string, dataType DataType, isEditable bool, isRequired bool, ) *DataDetailParametersResponseModel`

NewDataDetailParametersResponseModel instantiates a new DataDetailParametersResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataDetailParametersResponseModelWithDefaults

`func NewDataDetailParametersResponseModelWithDefaults() *DataDetailParametersResponseModel`

NewDataDetailParametersResponseModelWithDefaults instantiates a new DataDetailParametersResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *DataDetailParametersResponseModel) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *DataDetailParametersResponseModel) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *DataDetailParametersResponseModel) SetKey(v string)`

SetKey sets Key field to given value.


### GetDataType

`func (o *DataDetailParametersResponseModel) GetDataType() DataType`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *DataDetailParametersResponseModel) GetDataTypeOk() (*DataType, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *DataDetailParametersResponseModel) SetDataType(v DataType)`

SetDataType sets DataType field to given value.


### GetDefaultParameterValue

`func (o *DataDetailParametersResponseModel) GetDefaultParameterValue() string`

GetDefaultParameterValue returns the DefaultParameterValue field if non-nil, zero value otherwise.

### GetDefaultParameterValueOk

`func (o *DataDetailParametersResponseModel) GetDefaultParameterValueOk() (*string, bool)`

GetDefaultParameterValueOk returns a tuple with the DefaultParameterValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultParameterValue

`func (o *DataDetailParametersResponseModel) SetDefaultParameterValue(v string)`

SetDefaultParameterValue sets DefaultParameterValue field to given value.

### HasDefaultParameterValue

`func (o *DataDetailParametersResponseModel) HasDefaultParameterValue() bool`

HasDefaultParameterValue returns a boolean if a field has been set.

### SetDefaultParameterValueNil

`func (o *DataDetailParametersResponseModel) SetDefaultParameterValueNil(b bool)`

 SetDefaultParameterValueNil sets the value for DefaultParameterValue to be an explicit nil

### UnsetDefaultParameterValue
`func (o *DataDetailParametersResponseModel) UnsetDefaultParameterValue()`

UnsetDefaultParameterValue ensures that no value is present for DefaultParameterValue, not even an explicit nil
### GetIsEditable

`func (o *DataDetailParametersResponseModel) GetIsEditable() bool`

GetIsEditable returns the IsEditable field if non-nil, zero value otherwise.

### GetIsEditableOk

`func (o *DataDetailParametersResponseModel) GetIsEditableOk() (*bool, bool)`

GetIsEditableOk returns a tuple with the IsEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEditable

`func (o *DataDetailParametersResponseModel) SetIsEditable(v bool)`

SetIsEditable sets IsEditable field to given value.


### GetParameterType

`func (o *DataDetailParametersResponseModel) GetParameterType() ParameterType`

GetParameterType returns the ParameterType field if non-nil, zero value otherwise.

### GetParameterTypeOk

`func (o *DataDetailParametersResponseModel) GetParameterTypeOk() (*ParameterType, bool)`

GetParameterTypeOk returns a tuple with the ParameterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterType

`func (o *DataDetailParametersResponseModel) SetParameterType(v ParameterType)`

SetParameterType sets ParameterType field to given value.

### HasParameterType

`func (o *DataDetailParametersResponseModel) HasParameterType() bool`

HasParameterType returns a boolean if a field has been set.

### SetParameterTypeNil

`func (o *DataDetailParametersResponseModel) SetParameterTypeNil(b bool)`

 SetParameterTypeNil sets the value for ParameterType to be an explicit nil

### UnsetParameterType
`func (o *DataDetailParametersResponseModel) UnsetParameterType()`

UnsetParameterType ensures that no value is present for ParameterType, not even an explicit nil
### GetValue

`func (o *DataDetailParametersResponseModel) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DataDetailParametersResponseModel) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DataDetailParametersResponseModel) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *DataDetailParametersResponseModel) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *DataDetailParametersResponseModel) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *DataDetailParametersResponseModel) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetIsRequired

`func (o *DataDetailParametersResponseModel) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *DataDetailParametersResponseModel) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *DataDetailParametersResponseModel) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.


### GetValidationValueFormat

`func (o *DataDetailParametersResponseModel) GetValidationValueFormat() string`

GetValidationValueFormat returns the ValidationValueFormat field if non-nil, zero value otherwise.

### GetValidationValueFormatOk

`func (o *DataDetailParametersResponseModel) GetValidationValueFormatOk() (*string, bool)`

GetValidationValueFormatOk returns a tuple with the ValidationValueFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationValueFormat

`func (o *DataDetailParametersResponseModel) SetValidationValueFormat(v string)`

SetValidationValueFormat sets ValidationValueFormat field to given value.

### HasValidationValueFormat

`func (o *DataDetailParametersResponseModel) HasValidationValueFormat() bool`

HasValidationValueFormat returns a boolean if a field has been set.

### SetValidationValueFormatNil

`func (o *DataDetailParametersResponseModel) SetValidationValueFormatNil(b bool)`

 SetValidationValueFormatNil sets the value for ValidationValueFormat to be an explicit nil

### UnsetValidationValueFormat
`func (o *DataDetailParametersResponseModel) UnsetValidationValueFormat()`

UnsetValidationValueFormat ensures that no value is present for ValidationValueFormat, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


