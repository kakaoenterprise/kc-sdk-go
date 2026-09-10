# DataDetailParameters

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

### NewDataDetailParameters

`func NewDataDetailParameters(key string, dataType DataType, isEditable bool, isRequired bool, ) *DataDetailParameters`

NewDataDetailParameters instantiates a new DataDetailParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataDetailParametersWithDefaults

`func NewDataDetailParametersWithDefaults() *DataDetailParameters`

NewDataDetailParametersWithDefaults instantiates a new DataDetailParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *DataDetailParameters) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *DataDetailParameters) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *DataDetailParameters) SetKey(v string)`

SetKey sets Key field to given value.


### GetDataType

`func (o *DataDetailParameters) GetDataType() DataType`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *DataDetailParameters) GetDataTypeOk() (*DataType, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *DataDetailParameters) SetDataType(v DataType)`

SetDataType sets DataType field to given value.


### GetDefaultParameterValue

`func (o *DataDetailParameters) GetDefaultParameterValue() string`

GetDefaultParameterValue returns the DefaultParameterValue field if non-nil, zero value otherwise.

### GetDefaultParameterValueOk

`func (o *DataDetailParameters) GetDefaultParameterValueOk() (*string, bool)`

GetDefaultParameterValueOk returns a tuple with the DefaultParameterValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultParameterValue

`func (o *DataDetailParameters) SetDefaultParameterValue(v string)`

SetDefaultParameterValue sets DefaultParameterValue field to given value.

### HasDefaultParameterValue

`func (o *DataDetailParameters) HasDefaultParameterValue() bool`

HasDefaultParameterValue returns a boolean if a field has been set.

### SetDefaultParameterValueNil

`func (o *DataDetailParameters) SetDefaultParameterValueNil(b bool)`

 SetDefaultParameterValueNil sets the value for DefaultParameterValue to be an explicit nil

### UnsetDefaultParameterValue
`func (o *DataDetailParameters) UnsetDefaultParameterValue()`

UnsetDefaultParameterValue ensures that no value is present for DefaultParameterValue, not even an explicit nil
### GetIsEditable

`func (o *DataDetailParameters) GetIsEditable() bool`

GetIsEditable returns the IsEditable field if non-nil, zero value otherwise.

### GetIsEditableOk

`func (o *DataDetailParameters) GetIsEditableOk() (*bool, bool)`

GetIsEditableOk returns a tuple with the IsEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEditable

`func (o *DataDetailParameters) SetIsEditable(v bool)`

SetIsEditable sets IsEditable field to given value.


### GetParameterType

`func (o *DataDetailParameters) GetParameterType() ParameterType`

GetParameterType returns the ParameterType field if non-nil, zero value otherwise.

### GetParameterTypeOk

`func (o *DataDetailParameters) GetParameterTypeOk() (*ParameterType, bool)`

GetParameterTypeOk returns a tuple with the ParameterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterType

`func (o *DataDetailParameters) SetParameterType(v ParameterType)`

SetParameterType sets ParameterType field to given value.

### HasParameterType

`func (o *DataDetailParameters) HasParameterType() bool`

HasParameterType returns a boolean if a field has been set.

### SetParameterTypeNil

`func (o *DataDetailParameters) SetParameterTypeNil(b bool)`

 SetParameterTypeNil sets the value for ParameterType to be an explicit nil

### UnsetParameterType
`func (o *DataDetailParameters) UnsetParameterType()`

UnsetParameterType ensures that no value is present for ParameterType, not even an explicit nil
### GetValue

`func (o *DataDetailParameters) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DataDetailParameters) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DataDetailParameters) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *DataDetailParameters) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *DataDetailParameters) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *DataDetailParameters) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetIsRequired

`func (o *DataDetailParameters) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *DataDetailParameters) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *DataDetailParameters) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.


### GetValidationValueFormat

`func (o *DataDetailParameters) GetValidationValueFormat() string`

GetValidationValueFormat returns the ValidationValueFormat field if non-nil, zero value otherwise.

### GetValidationValueFormatOk

`func (o *DataDetailParameters) GetValidationValueFormatOk() (*string, bool)`

GetValidationValueFormatOk returns a tuple with the ValidationValueFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationValueFormat

`func (o *DataDetailParameters) SetValidationValueFormat(v string)`

SetValidationValueFormat sets ValidationValueFormat field to given value.

### HasValidationValueFormat

`func (o *DataDetailParameters) HasValidationValueFormat() bool`

HasValidationValueFormat returns a boolean if a field has been set.

### SetValidationValueFormatNil

`func (o *DataDetailParameters) SetValidationValueFormatNil(b bool)`

 SetValidationValueFormatNil sets the value for ValidationValueFormat to be an explicit nil

### UnsetValidationValueFormat
`func (o *DataDetailParameters) UnsetValidationValueFormat()`

UnsetValidationValueFormat ensures that no value is present for ValidationValueFormat, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


