# ParameterResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | MySQL 파라미터 이름 | 
**Value** | Pointer to **NullableString** |  | [optional] 
**ParameterType** | [**ParameterType**](ParameterType.md) | 파라미터 적용 유형 (예시: &#x60;STATIC&#x60;, &#x60;DYNAMIC&#x60;, &#x60;REPLICA&#x60; 등) | 
**DataType** | [**DataType**](DataType.md) | 파라미터 값의 데이터 형식 (예: &#x60;STRING&#x60;, &#x60;INTEGER&#x60; 등 ) | 
**ValidationValueFormat** | Pointer to **NullableString** |  | [optional] 
**IsEditable** | **bool** | 사용자 수정 가능 여부 | 
**IsRequired** | **bool** | 필수 파라미터 여부 | 

## Methods

### NewParameterResponseModel

`func NewParameterResponseModel(key string, parameterType ParameterType, dataType DataType, isEditable bool, isRequired bool, ) *ParameterResponseModel`

NewParameterResponseModel instantiates a new ParameterResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParameterResponseModelWithDefaults

`func NewParameterResponseModelWithDefaults() *ParameterResponseModel`

NewParameterResponseModelWithDefaults instantiates a new ParameterResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ParameterResponseModel) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ParameterResponseModel) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ParameterResponseModel) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *ParameterResponseModel) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ParameterResponseModel) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ParameterResponseModel) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ParameterResponseModel) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ParameterResponseModel) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ParameterResponseModel) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetParameterType

`func (o *ParameterResponseModel) GetParameterType() ParameterType`

GetParameterType returns the ParameterType field if non-nil, zero value otherwise.

### GetParameterTypeOk

`func (o *ParameterResponseModel) GetParameterTypeOk() (*ParameterType, bool)`

GetParameterTypeOk returns a tuple with the ParameterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterType

`func (o *ParameterResponseModel) SetParameterType(v ParameterType)`

SetParameterType sets ParameterType field to given value.


### GetDataType

`func (o *ParameterResponseModel) GetDataType() DataType`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *ParameterResponseModel) GetDataTypeOk() (*DataType, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *ParameterResponseModel) SetDataType(v DataType)`

SetDataType sets DataType field to given value.


### GetValidationValueFormat

`func (o *ParameterResponseModel) GetValidationValueFormat() string`

GetValidationValueFormat returns the ValidationValueFormat field if non-nil, zero value otherwise.

### GetValidationValueFormatOk

`func (o *ParameterResponseModel) GetValidationValueFormatOk() (*string, bool)`

GetValidationValueFormatOk returns a tuple with the ValidationValueFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationValueFormat

`func (o *ParameterResponseModel) SetValidationValueFormat(v string)`

SetValidationValueFormat sets ValidationValueFormat field to given value.

### HasValidationValueFormat

`func (o *ParameterResponseModel) HasValidationValueFormat() bool`

HasValidationValueFormat returns a boolean if a field has been set.

### SetValidationValueFormatNil

`func (o *ParameterResponseModel) SetValidationValueFormatNil(b bool)`

 SetValidationValueFormatNil sets the value for ValidationValueFormat to be an explicit nil

### UnsetValidationValueFormat
`func (o *ParameterResponseModel) UnsetValidationValueFormat()`

UnsetValidationValueFormat ensures that no value is present for ValidationValueFormat, not even an explicit nil
### GetIsEditable

`func (o *ParameterResponseModel) GetIsEditable() bool`

GetIsEditable returns the IsEditable field if non-nil, zero value otherwise.

### GetIsEditableOk

`func (o *ParameterResponseModel) GetIsEditableOk() (*bool, bool)`

GetIsEditableOk returns a tuple with the IsEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEditable

`func (o *ParameterResponseModel) SetIsEditable(v bool)`

SetIsEditable sets IsEditable field to given value.


### GetIsRequired

`func (o *ParameterResponseModel) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *ParameterResponseModel) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *ParameterResponseModel) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


