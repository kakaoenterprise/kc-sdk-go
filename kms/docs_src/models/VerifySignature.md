# VerifySignature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Input** | **string** | 서명을 검증할 원본 데이터 | 
**Output** | **string** | 검증할 전자 서명  - [Sign data](/openapi/security/kms/sign-data) 응답에서 확인 | 

## Methods

### NewVerifySignature

`func NewVerifySignature(input string, output string, ) *VerifySignature`

NewVerifySignature instantiates a new VerifySignature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifySignatureWithDefaults

`func NewVerifySignatureWithDefaults() *VerifySignature`

NewVerifySignatureWithDefaults instantiates a new VerifySignature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInput

`func (o *VerifySignature) GetInput() string`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *VerifySignature) GetInputOk() (*string, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *VerifySignature) SetInput(v string)`

SetInput sets Input field to given value.


### GetOutput

`func (o *VerifySignature) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *VerifySignature) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *VerifySignature) SetOutput(v string)`

SetOutput sets Output field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


