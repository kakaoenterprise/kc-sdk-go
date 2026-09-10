# CreateLoadBalancerMemberRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Address** | **string** | 대상 서버 IP 주소 | 
**ProtocolPort** | **int32** | 대상 서버 포트 | 
**Weight** | Pointer to **NullableInt32** |  | [optional] 
**MonitorPort** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewCreateLoadBalancerMemberRequest

`func NewCreateLoadBalancerMemberRequest(address string, protocolPort int32, ) *CreateLoadBalancerMemberRequest`

NewCreateLoadBalancerMemberRequest instantiates a new CreateLoadBalancerMemberRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLoadBalancerMemberRequestWithDefaults

`func NewCreateLoadBalancerMemberRequestWithDefaults() *CreateLoadBalancerMemberRequest`

NewCreateLoadBalancerMemberRequestWithDefaults instantiates a new CreateLoadBalancerMemberRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateLoadBalancerMemberRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateLoadBalancerMemberRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateLoadBalancerMemberRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateLoadBalancerMemberRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreateLoadBalancerMemberRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateLoadBalancerMemberRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAddress

`func (o *CreateLoadBalancerMemberRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CreateLoadBalancerMemberRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CreateLoadBalancerMemberRequest) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetProtocolPort

`func (o *CreateLoadBalancerMemberRequest) GetProtocolPort() int32`

GetProtocolPort returns the ProtocolPort field if non-nil, zero value otherwise.

### GetProtocolPortOk

`func (o *CreateLoadBalancerMemberRequest) GetProtocolPortOk() (*int32, bool)`

GetProtocolPortOk returns a tuple with the ProtocolPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolPort

`func (o *CreateLoadBalancerMemberRequest) SetProtocolPort(v int32)`

SetProtocolPort sets ProtocolPort field to given value.


### GetWeight

`func (o *CreateLoadBalancerMemberRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *CreateLoadBalancerMemberRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *CreateLoadBalancerMemberRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *CreateLoadBalancerMemberRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### SetWeightNil

`func (o *CreateLoadBalancerMemberRequest) SetWeightNil(b bool)`

 SetWeightNil sets the value for Weight to be an explicit nil

### UnsetWeight
`func (o *CreateLoadBalancerMemberRequest) UnsetWeight()`

UnsetWeight ensures that no value is present for Weight, not even an explicit nil
### GetMonitorPort

`func (o *CreateLoadBalancerMemberRequest) GetMonitorPort() int32`

GetMonitorPort returns the MonitorPort field if non-nil, zero value otherwise.

### GetMonitorPortOk

`func (o *CreateLoadBalancerMemberRequest) GetMonitorPortOk() (*int32, bool)`

GetMonitorPortOk returns a tuple with the MonitorPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorPort

`func (o *CreateLoadBalancerMemberRequest) SetMonitorPort(v int32)`

SetMonitorPort sets MonitorPort field to given value.

### HasMonitorPort

`func (o *CreateLoadBalancerMemberRequest) HasMonitorPort() bool`

HasMonitorPort returns a boolean if a field has been set.

### SetMonitorPortNil

`func (o *CreateLoadBalancerMemberRequest) SetMonitorPortNil(b bool)`

 SetMonitorPortNil sets the value for MonitorPort to be an explicit nil

### UnsetMonitorPort
`func (o *CreateLoadBalancerMemberRequest) UnsetMonitorPort()`

UnsetMonitorPort ensures that no value is present for MonitorPort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


