# UpdateSecurityGroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecurityGroup** | [**SecurityGroupResult**](SecurityGroupResult.md) | 수정된 보안 그룹 정보 | 

## Methods

### NewUpdateSecurityGroupResponse

`func NewUpdateSecurityGroupResponse(securityGroup SecurityGroupResult, ) *UpdateSecurityGroupResponse`

NewUpdateSecurityGroupResponse instantiates a new UpdateSecurityGroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSecurityGroupResponseWithDefaults

`func NewUpdateSecurityGroupResponseWithDefaults() *UpdateSecurityGroupResponse`

NewUpdateSecurityGroupResponseWithDefaults instantiates a new UpdateSecurityGroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecurityGroup

`func (o *UpdateSecurityGroupResponse) GetSecurityGroup() SecurityGroupResult`

GetSecurityGroup returns the SecurityGroup field if non-nil, zero value otherwise.

### GetSecurityGroupOk

`func (o *UpdateSecurityGroupResponse) GetSecurityGroupOk() (*SecurityGroupResult, bool)`

GetSecurityGroupOk returns a tuple with the SecurityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroup

`func (o *UpdateSecurityGroupResponse) SetSecurityGroup(v SecurityGroupResult)`

SetSecurityGroup sets SecurityGroup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


