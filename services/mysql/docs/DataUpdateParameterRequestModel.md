# DataUpdateParameterRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | 수정할 MySQL 파라미터 이름 | 
**Value** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDataUpdateParameterRequestModel

`func NewDataUpdateParameterRequestModel(key string, ) *DataUpdateParameterRequestModel`

NewDataUpdateParameterRequestModel instantiates a new DataUpdateParameterRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataUpdateParameterRequestModelWithDefaults

`func NewDataUpdateParameterRequestModelWithDefaults() *DataUpdateParameterRequestModel`

NewDataUpdateParameterRequestModelWithDefaults instantiates a new DataUpdateParameterRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *DataUpdateParameterRequestModel) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *DataUpdateParameterRequestModel) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *DataUpdateParameterRequestModel) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *DataUpdateParameterRequestModel) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DataUpdateParameterRequestModel) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DataUpdateParameterRequestModel) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *DataUpdateParameterRequestModel) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *DataUpdateParameterRequestModel) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *DataUpdateParameterRequestModel) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


