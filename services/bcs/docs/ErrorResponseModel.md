# ErrorResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | [**ErrorResponseSubModel**](ErrorResponseSubModel.md) | 에러에 대한 정보 | 

## Methods

### NewErrorResponseModel

`func NewErrorResponseModel(error_ ErrorResponseSubModel, ) *ErrorResponseModel`

NewErrorResponseModel instantiates a new ErrorResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorResponseModelWithDefaults

`func NewErrorResponseModelWithDefaults() *ErrorResponseModel`

NewErrorResponseModelWithDefaults instantiates a new ErrorResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *ErrorResponseModel) GetError() ErrorResponseSubModel`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ErrorResponseModel) GetErrorOk() (*ErrorResponseSubModel, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ErrorResponseModel) SetError(v ErrorResponseSubModel)`

SetError sets Error field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


