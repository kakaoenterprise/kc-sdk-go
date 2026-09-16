# UpdateSecurityGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecurityGroup** | [**UpdateSecurityGroup**](UpdateSecurityGroup.md) | 수정할 보안 그룹 정보 | 

## Methods

### NewUpdateSecurityGroupRequest

`func NewUpdateSecurityGroupRequest(securityGroup UpdateSecurityGroup, ) *UpdateSecurityGroupRequest`

NewUpdateSecurityGroupRequest instantiates a new UpdateSecurityGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSecurityGroupRequestWithDefaults

`func NewUpdateSecurityGroupRequestWithDefaults() *UpdateSecurityGroupRequest`

NewUpdateSecurityGroupRequestWithDefaults instantiates a new UpdateSecurityGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecurityGroup

`func (o *UpdateSecurityGroupRequest) GetSecurityGroup() UpdateSecurityGroup`

GetSecurityGroup returns the SecurityGroup field if non-nil, zero value otherwise.

### GetSecurityGroupOk

`func (o *UpdateSecurityGroupRequest) GetSecurityGroupOk() (*UpdateSecurityGroup, bool)`

GetSecurityGroupOk returns a tuple with the SecurityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroup

`func (o *UpdateSecurityGroupRequest) SetSecurityGroup(v UpdateSecurityGroup)`

SetSecurityGroup sets SecurityGroup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


