# RemoveUserKeyAccessTargets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | **[]string** | 접근 제어 대상의 고유 ID 목록 &lt;br/&gt;- [List user key access targets](https://docs.kakaocloud.com/openapi/security/kms/list-user-key-access-targets)에서 확인 | 

## Methods

### NewRemoveUserKeyAccessTargets

`func NewRemoveUserKeyAccessTargets(ids []string, ) *RemoveUserKeyAccessTargets`

NewRemoveUserKeyAccessTargets instantiates a new RemoveUserKeyAccessTargets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoveUserKeyAccessTargetsWithDefaults

`func NewRemoveUserKeyAccessTargetsWithDefaults() *RemoveUserKeyAccessTargets`

NewRemoveUserKeyAccessTargetsWithDefaults instantiates a new RemoveUserKeyAccessTargets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *RemoveUserKeyAccessTargets) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *RemoveUserKeyAccessTargets) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *RemoveUserKeyAccessTargets) SetIds(v []string)`

SetIds sets Ids field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


