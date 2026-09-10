# GetSecretVersionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | [**GetSecretVersionVersion**](GetSecretVersionVersion.md) | 시크릿 버전 데이터 | 

## Methods

### NewGetSecretVersionResponse

`func NewGetSecretVersionResponse(version GetSecretVersionVersion, ) *GetSecretVersionResponse`

NewGetSecretVersionResponse instantiates a new GetSecretVersionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSecretVersionResponseWithDefaults

`func NewGetSecretVersionResponseWithDefaults() *GetSecretVersionResponse`

NewGetSecretVersionResponseWithDefaults instantiates a new GetSecretVersionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *GetSecretVersionResponse) GetVersion() GetSecretVersionVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetSecretVersionResponse) GetVersionOk() (*GetSecretVersionVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetSecretVersionResponse) SetVersion(v GetSecretVersionVersion)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


