# SetNodePoolSecurityGroupsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecurityGroups** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSetNodePoolSecurityGroupsRequest

`func NewSetNodePoolSecurityGroupsRequest() *SetNodePoolSecurityGroupsRequest`

NewSetNodePoolSecurityGroupsRequest instantiates a new SetNodePoolSecurityGroupsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetNodePoolSecurityGroupsRequestWithDefaults

`func NewSetNodePoolSecurityGroupsRequestWithDefaults() *SetNodePoolSecurityGroupsRequest`

NewSetNodePoolSecurityGroupsRequestWithDefaults instantiates a new SetNodePoolSecurityGroupsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecurityGroups

`func (o *SetNodePoolSecurityGroupsRequest) GetSecurityGroups() []string`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *SetNodePoolSecurityGroupsRequest) GetSecurityGroupsOk() (*[]string, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *SetNodePoolSecurityGroupsRequest) SetSecurityGroups(v []string)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *SetNodePoolSecurityGroupsRequest) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### SetSecurityGroupsNil

`func (o *SetNodePoolSecurityGroupsRequest) SetSecurityGroupsNil(b bool)`

 SetSecurityGroupsNil sets the value for SecurityGroups to be an explicit nil

### UnsetSecurityGroups
`func (o *SetNodePoolSecurityGroupsRequest) UnsetSecurityGroups()`

UnsetSecurityGroups ensures that no value is present for SecurityGroups, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


