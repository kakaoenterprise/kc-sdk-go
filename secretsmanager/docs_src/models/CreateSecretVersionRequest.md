# CreateSecretVersionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | [**CreateSecretVersion**](CreateSecretVersion.md) | 시크릿 버전 | 

## Methods

### NewCreateSecretVersionRequest

`func NewCreateSecretVersionRequest(version CreateSecretVersion, ) *CreateSecretVersionRequest`

NewCreateSecretVersionRequest instantiates a new CreateSecretVersionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSecretVersionRequestWithDefaults

`func NewCreateSecretVersionRequestWithDefaults() *CreateSecretVersionRequest`

NewCreateSecretVersionRequestWithDefaults instantiates a new CreateSecretVersionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *CreateSecretVersionRequest) GetVersion() CreateSecretVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CreateSecretVersionRequest) GetVersionOk() (*CreateSecretVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CreateSecretVersionRequest) SetVersion(v CreateSecretVersion)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


