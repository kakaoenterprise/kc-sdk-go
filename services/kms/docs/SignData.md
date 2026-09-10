# SignData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlainText** | **string** | Base64로 인코딩된 평문 데이터 | 

## Methods

### NewSignData

`func NewSignData(plainText string, ) *SignData`

NewSignData instantiates a new SignData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignDataWithDefaults

`func NewSignDataWithDefaults() *SignData`

NewSignDataWithDefaults instantiates a new SignData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlainText

`func (o *SignData) GetPlainText() string`

GetPlainText returns the PlainText field if non-nil, zero value otherwise.

### GetPlainTextOk

`func (o *SignData) GetPlainTextOk() (*string, bool)`

GetPlainTextOk returns a tuple with the PlainText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlainText

`func (o *SignData) SetPlainText(v string)`

SetPlainText sets PlainText field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


