# VerifyHmac

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Input** | **string** | 서명을 검증할 원본 데이터 | 
**Output** | **string** | 검증할 HMAC  - [Generate HMAC](/openapi/security/kms/generate-hmac) 응답에서 확인 | 

## Methods

### NewVerifyHmac

`func NewVerifyHmac(input string, output string, ) *VerifyHmac`

NewVerifyHmac instantiates a new VerifyHmac object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyHmacWithDefaults

`func NewVerifyHmacWithDefaults() *VerifyHmac`

NewVerifyHmacWithDefaults instantiates a new VerifyHmac object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInput

`func (o *VerifyHmac) GetInput() string`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *VerifyHmac) GetInputOk() (*string, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *VerifyHmac) SetInput(v string)`

SetInput sets Input field to given value.


### GetOutput

`func (o *VerifyHmac) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *VerifyHmac) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *VerifyHmac) SetOutput(v string)`

SetOutput sets Output field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


