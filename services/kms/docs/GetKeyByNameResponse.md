# GetKeyByNameResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | [**GetKeyByNameKey**](GetKeyByNameKey.md) | KMS 키 정보 | 

## Methods

### NewGetKeyByNameResponse

`func NewGetKeyByNameResponse(key GetKeyByNameKey, ) *GetKeyByNameResponse`

NewGetKeyByNameResponse instantiates a new GetKeyByNameResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetKeyByNameResponseWithDefaults

`func NewGetKeyByNameResponseWithDefaults() *GetKeyByNameResponse`

NewGetKeyByNameResponseWithDefaults instantiates a new GetKeyByNameResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *GetKeyByNameResponse) GetKey() GetKeyByNameKey`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *GetKeyByNameResponse) GetKeyOk() (*GetKeyByNameKey, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *GetKeyByNameResponse) SetKey(v GetKeyByNameKey)`

SetKey sets Key field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


