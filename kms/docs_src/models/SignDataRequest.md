# SignDataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sign** | [**SignData**](SignData.md) | 데이터 서명 요청 또는 결과 정보 | 

## Methods

### NewSignDataRequest

`func NewSignDataRequest(sign SignData, ) *SignDataRequest`

NewSignDataRequest instantiates a new SignDataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignDataRequestWithDefaults

`func NewSignDataRequestWithDefaults() *SignDataRequest`

NewSignDataRequestWithDefaults instantiates a new SignDataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSign

`func (o *SignDataRequest) GetSign() SignData`

GetSign returns the Sign field if non-nil, zero value otherwise.

### GetSignOk

`func (o *SignDataRequest) GetSignOk() (*SignData, bool)`

GetSignOk returns a tuple with the Sign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSign

`func (o *SignDataRequest) SetSign(v SignData)`

SetSign sets Sign field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


