# UpdateTargetsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Members** | [**[]UpdateTargets**](UpdateTargets.md) | 수정할 대상 그룹 멤버 목록 | 

## Methods

### NewUpdateTargetsRequest

`func NewUpdateTargetsRequest(members []UpdateTargets, ) *UpdateTargetsRequest`

NewUpdateTargetsRequest instantiates a new UpdateTargetsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTargetsRequestWithDefaults

`func NewUpdateTargetsRequestWithDefaults() *UpdateTargetsRequest`

NewUpdateTargetsRequestWithDefaults instantiates a new UpdateTargetsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMembers

`func (o *UpdateTargetsRequest) GetMembers() []UpdateTargets`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *UpdateTargetsRequest) GetMembersOk() (*[]UpdateTargets, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *UpdateTargetsRequest) SetMembers(v []UpdateTargets)`

SetMembers sets Members field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


