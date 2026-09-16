# VerifySignatureRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Verify** | [**VerifySignature**](VerifySignature.md) | 전자 서명 또는 HMAC 검증 요청 또는 결과 정보 | 

## Methods

### NewVerifySignatureRequest

`func NewVerifySignatureRequest(verify VerifySignature, ) *VerifySignatureRequest`

NewVerifySignatureRequest instantiates a new VerifySignatureRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifySignatureRequestWithDefaults

`func NewVerifySignatureRequestWithDefaults() *VerifySignatureRequest`

NewVerifySignatureRequestWithDefaults instantiates a new VerifySignatureRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerify

`func (o *VerifySignatureRequest) GetVerify() VerifySignature`

GetVerify returns the Verify field if non-nil, zero value otherwise.

### GetVerifyOk

`func (o *VerifySignatureRequest) GetVerifyOk() (*VerifySignature, bool)`

GetVerifyOk returns a tuple with the Verify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerify

`func (o *VerifySignatureRequest) SetVerify(v VerifySignature)`

SetVerify sets Verify field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


