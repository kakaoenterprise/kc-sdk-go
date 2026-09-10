# AddSecretsAccessTargetsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Targets** | [**AddSecretsAccessTargets**](AddSecretsAccessTargets.md) | 접근 제어 허용 대상 정보 | 

## Methods

### NewAddSecretsAccessTargetsRequest

`func NewAddSecretsAccessTargetsRequest(targets AddSecretsAccessTargets, ) *AddSecretsAccessTargetsRequest`

NewAddSecretsAccessTargetsRequest instantiates a new AddSecretsAccessTargetsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddSecretsAccessTargetsRequestWithDefaults

`func NewAddSecretsAccessTargetsRequestWithDefaults() *AddSecretsAccessTargetsRequest`

NewAddSecretsAccessTargetsRequestWithDefaults instantiates a new AddSecretsAccessTargetsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargets

`func (o *AddSecretsAccessTargetsRequest) GetTargets() AddSecretsAccessTargets`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *AddSecretsAccessTargetsRequest) GetTargetsOk() (*AddSecretsAccessTargets, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *AddSecretsAccessTargetsRequest) SetTargets(v AddSecretsAccessTargets)`

SetTargets sets Targets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


