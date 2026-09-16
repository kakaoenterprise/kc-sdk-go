# GetSecretResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Secret** | [**Secret**](Secret.md) | 시크릿 정보 | 

## Methods

### NewGetSecretResponse

`func NewGetSecretResponse(secret Secret, ) *GetSecretResponse`

NewGetSecretResponse instantiates a new GetSecretResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSecretResponseWithDefaults

`func NewGetSecretResponseWithDefaults() *GetSecretResponse`

NewGetSecretResponseWithDefaults instantiates a new GetSecretResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecret

`func (o *GetSecretResponse) GetSecret() Secret`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *GetSecretResponse) GetSecretOk() (*Secret, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *GetSecretResponse) SetSecret(v Secret)`

SetSecret sets Secret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


