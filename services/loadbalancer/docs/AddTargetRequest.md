# AddTargetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Member** | [**AddTarget**](AddTarget.md) | 추가할 대상 그룹 멤버 정보 | 

## Methods

### NewAddTargetRequest

`func NewAddTargetRequest(member AddTarget, ) *AddTargetRequest`

NewAddTargetRequest instantiates a new AddTargetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddTargetRequestWithDefaults

`func NewAddTargetRequestWithDefaults() *AddTargetRequest`

NewAddTargetRequestWithDefaults instantiates a new AddTargetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMember

`func (o *AddTargetRequest) GetMember() AddTarget`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *AddTargetRequest) GetMemberOk() (*AddTarget, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *AddTargetRequest) SetMember(v AddTarget)`

SetMember sets Member field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


