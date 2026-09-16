# GetKeyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | [**Key**](Key.md) | KMS 키 정보 | 

## Methods

### NewGetKeyResponse

`func NewGetKeyResponse(key Key, ) *GetKeyResponse`

NewGetKeyResponse instantiates a new GetKeyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetKeyResponseWithDefaults

`func NewGetKeyResponseWithDefaults() *GetKeyResponse`

NewGetKeyResponseWithDefaults instantiates a new GetKeyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *GetKeyResponse) GetKey() Key`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *GetKeyResponse) GetKeyOk() (*Key, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *GetKeyResponse) SetKey(v Key)`

SetKey sets Key field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


