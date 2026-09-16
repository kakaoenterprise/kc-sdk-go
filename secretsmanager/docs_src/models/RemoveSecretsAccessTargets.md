# RemoveSecretsAccessTargets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | **[]string** | 접근 제어 대상의 고유 ID 목록 - [List secrets access targets](/openapi/security/secrets-manager/list-secrets-access-targets)에서 확인 | 

## Methods

### NewRemoveSecretsAccessTargets

`func NewRemoveSecretsAccessTargets(ids []string, ) *RemoveSecretsAccessTargets`

NewRemoveSecretsAccessTargets instantiates a new RemoveSecretsAccessTargets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoveSecretsAccessTargetsWithDefaults

`func NewRemoveSecretsAccessTargetsWithDefaults() *RemoveSecretsAccessTargets`

NewRemoveSecretsAccessTargetsWithDefaults instantiates a new RemoveSecretsAccessTargets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *RemoveSecretsAccessTargets) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *RemoveSecretsAccessTargets) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *RemoveSecretsAccessTargets) SetIds(v []string)`

SetIds sets Ids field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


