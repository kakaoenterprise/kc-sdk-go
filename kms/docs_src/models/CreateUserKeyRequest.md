# CreateUserKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | [**CreateUserKey**](CreateUserKey.md) | KMS 키 정보 | 

## Methods

### NewCreateUserKeyRequest

`func NewCreateUserKeyRequest(key CreateUserKey, ) *CreateUserKeyRequest`

NewCreateUserKeyRequest instantiates a new CreateUserKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserKeyRequestWithDefaults

`func NewCreateUserKeyRequestWithDefaults() *CreateUserKeyRequest`

NewCreateUserKeyRequestWithDefaults instantiates a new CreateUserKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *CreateUserKeyRequest) GetKey() CreateUserKey`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateUserKeyRequest) GetKeyOk() (*CreateUserKey, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateUserKeyRequest) SetKey(v CreateUserKey)`

SetKey sets Key field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


