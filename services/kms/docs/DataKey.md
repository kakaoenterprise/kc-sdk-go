# DataKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlainText** | Pointer to **NullableString** |  | [optional] 
**CipherText** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDataKey

`func NewDataKey() *DataKey`

NewDataKey instantiates a new DataKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataKeyWithDefaults

`func NewDataKeyWithDefaults() *DataKey`

NewDataKeyWithDefaults instantiates a new DataKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlainText

`func (o *DataKey) GetPlainText() string`

GetPlainText returns the PlainText field if non-nil, zero value otherwise.

### GetPlainTextOk

`func (o *DataKey) GetPlainTextOk() (*string, bool)`

GetPlainTextOk returns a tuple with the PlainText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlainText

`func (o *DataKey) SetPlainText(v string)`

SetPlainText sets PlainText field to given value.

### HasPlainText

`func (o *DataKey) HasPlainText() bool`

HasPlainText returns a boolean if a field has been set.

### SetPlainTextNil

`func (o *DataKey) SetPlainTextNil(b bool)`

 SetPlainTextNil sets the value for PlainText to be an explicit nil

### UnsetPlainText
`func (o *DataKey) UnsetPlainText()`

UnsetPlainText ensures that no value is present for PlainText, not even an explicit nil
### GetCipherText

`func (o *DataKey) GetCipherText() string`

GetCipherText returns the CipherText field if non-nil, zero value otherwise.

### GetCipherTextOk

`func (o *DataKey) GetCipherTextOk() (*string, bool)`

GetCipherTextOk returns a tuple with the CipherText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCipherText

`func (o *DataKey) SetCipherText(v string)`

SetCipherText sets CipherText field to given value.

### HasCipherText

`func (o *DataKey) HasCipherText() bool`

HasCipherText returns a boolean if a field has been set.

### SetCipherTextNil

`func (o *DataKey) SetCipherTextNil(b bool)`

 SetCipherTextNil sets the value for CipherText to be an explicit nil

### UnsetCipherText
`func (o *DataKey) UnsetCipherText()`

UnsetCipherText ensures that no value is present for CipherText, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


