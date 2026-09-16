# Parameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | MySQL 파라미터 이름 | 
**Value** | Pointer to **NullableString** | MySQL 파라미터에 설정된 값 | [optional] 
**ParameterType** | [**ParameterType**](ParameterType.md) | 파라미터 적용 유형 | 
**DataType** | [**DataType**](DataType.md) | 파라미터 값의 데이터 형식 | 
**ValidationValueFormat** | Pointer to **NullableString** | 파라미터 값 검증을 위한 형식 또는 범위 | [optional] 
**IsEditable** | **bool** | 사용자 수정 가능 여부 | 
**IsRequired** | **bool** | 필수 파라미터 여부 | 

## Methods

### NewParameter

`func NewParameter(key string, parameterType ParameterType, dataType DataType, isEditable bool, isRequired bool, ) *Parameter`

NewParameter instantiates a new Parameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParameterWithDefaults

`func NewParameterWithDefaults() *Parameter`

NewParameterWithDefaults instantiates a new Parameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *Parameter) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Parameter) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Parameter) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *Parameter) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Parameter) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Parameter) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *Parameter) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *Parameter) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *Parameter) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetParameterType

`func (o *Parameter) GetParameterType() ParameterType`

GetParameterType returns the ParameterType field if non-nil, zero value otherwise.

### GetParameterTypeOk

`func (o *Parameter) GetParameterTypeOk() (*ParameterType, bool)`

GetParameterTypeOk returns a tuple with the ParameterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterType

`func (o *Parameter) SetParameterType(v ParameterType)`

SetParameterType sets ParameterType field to given value.


### GetDataType

`func (o *Parameter) GetDataType() DataType`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *Parameter) GetDataTypeOk() (*DataType, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *Parameter) SetDataType(v DataType)`

SetDataType sets DataType field to given value.


### GetValidationValueFormat

`func (o *Parameter) GetValidationValueFormat() string`

GetValidationValueFormat returns the ValidationValueFormat field if non-nil, zero value otherwise.

### GetValidationValueFormatOk

`func (o *Parameter) GetValidationValueFormatOk() (*string, bool)`

GetValidationValueFormatOk returns a tuple with the ValidationValueFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationValueFormat

`func (o *Parameter) SetValidationValueFormat(v string)`

SetValidationValueFormat sets ValidationValueFormat field to given value.

### HasValidationValueFormat

`func (o *Parameter) HasValidationValueFormat() bool`

HasValidationValueFormat returns a boolean if a field has been set.

### SetValidationValueFormatNil

`func (o *Parameter) SetValidationValueFormatNil(b bool)`

 SetValidationValueFormatNil sets the value for ValidationValueFormat to be an explicit nil

### UnsetValidationValueFormat
`func (o *Parameter) UnsetValidationValueFormat()`

UnsetValidationValueFormat ensures that no value is present for ValidationValueFormat, not even an explicit nil
### GetIsEditable

`func (o *Parameter) GetIsEditable() bool`

GetIsEditable returns the IsEditable field if non-nil, zero value otherwise.

### GetIsEditableOk

`func (o *Parameter) GetIsEditableOk() (*bool, bool)`

GetIsEditableOk returns a tuple with the IsEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEditable

`func (o *Parameter) SetIsEditable(v bool)`

SetIsEditable sets IsEditable field to given value.


### GetIsRequired

`func (o *Parameter) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *Parameter) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *Parameter) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


