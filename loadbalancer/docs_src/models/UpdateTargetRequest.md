# UpdateTargetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Member** | [**UpdateTarget**](UpdateTarget.md) | 수정할 대상 그룹 멤버 정보 | 

## Methods

### NewUpdateTargetRequest

`func NewUpdateTargetRequest(member UpdateTarget, ) *UpdateTargetRequest`

NewUpdateTargetRequest instantiates a new UpdateTargetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTargetRequestWithDefaults

`func NewUpdateTargetRequestWithDefaults() *UpdateTargetRequest`

NewUpdateTargetRequestWithDefaults instantiates a new UpdateTargetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMember

`func (o *UpdateTargetRequest) GetMember() UpdateTarget`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *UpdateTargetRequest) GetMemberOk() (*UpdateTarget, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *UpdateTargetRequest) SetMember(v UpdateTarget)`

SetMember sets Member field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


