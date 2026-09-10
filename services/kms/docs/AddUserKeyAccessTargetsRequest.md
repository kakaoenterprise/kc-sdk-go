# AddUserKeyAccessTargetsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Targets** | [**AddUserKeyAccessTargets**](AddUserKeyAccessTargets.md) | 접근 제어 허용 대상 목록 또는 처리 결과 | 

## Methods

### NewAddUserKeyAccessTargetsRequest

`func NewAddUserKeyAccessTargetsRequest(targets AddUserKeyAccessTargets, ) *AddUserKeyAccessTargetsRequest`

NewAddUserKeyAccessTargetsRequest instantiates a new AddUserKeyAccessTargetsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddUserKeyAccessTargetsRequestWithDefaults

`func NewAddUserKeyAccessTargetsRequestWithDefaults() *AddUserKeyAccessTargetsRequest`

NewAddUserKeyAccessTargetsRequestWithDefaults instantiates a new AddUserKeyAccessTargetsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargets

`func (o *AddUserKeyAccessTargetsRequest) GetTargets() AddUserKeyAccessTargets`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *AddUserKeyAccessTargetsRequest) GetTargetsOk() (*AddUserKeyAccessTargets, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *AddUserKeyAccessTargetsRequest) SetTargets(v AddUserKeyAccessTargets)`

SetTargets sets Targets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


