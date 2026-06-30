# ErrorResponseSubModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | 에러 코드 | 
**Message** | **string** | 상세 에러 메시지 | 

## Methods

### NewErrorResponseSubModel

`func NewErrorResponseSubModel(code string, message string, ) *ErrorResponseSubModel`

NewErrorResponseSubModel instantiates a new ErrorResponseSubModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorResponseSubModelWithDefaults

`func NewErrorResponseSubModelWithDefaults() *ErrorResponseSubModel`

NewErrorResponseSubModelWithDefaults instantiates a new ErrorResponseSubModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ErrorResponseSubModel) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorResponseSubModel) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorResponseSubModel) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *ErrorResponseSubModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorResponseSubModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorResponseSubModel) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


