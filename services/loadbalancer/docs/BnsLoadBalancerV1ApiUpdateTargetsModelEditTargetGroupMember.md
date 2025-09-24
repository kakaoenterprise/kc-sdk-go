# BnsLoadBalancerV1ApiUpdateTargetsModelEditTargetGroupMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Address** | **string** | 대상 인스턴스의 IP 주소 | 
**ProtocolPort** | **int32** | 대상 인스턴스의 수신 포트 번호 | 
**SubnetId** | **string** | 대상 인스턴스가 속한 서브넷의 ID | 
**Weight** | Pointer to **NullableInt32** |  | [optional] 
**MonitorPort** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewBnsLoadBalancerV1ApiUpdateTargetsModelEditTargetGroupMember

`func NewBnsLoadBalancerV1ApiUpdateTargetsModelEditTargetGroupMember(address string, protocolPort int32, subnetId string, ) *BnsLoadBalancerV1ApiUpdateTargetsModelEditTargetGroupMember`

NewBnsLoadBalancerV1ApiUpdateTargetsModelEditTargetGroupMember instantiates a new BnsLoadBalancerV1ApiUpdateTargetsModelEditTargetGroupMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsLoadBalancerV1ApiUpdateTargetsModelEditTargetGroupMemberWithDefaults

`func NewBnsLoadBalancerV1ApiUpdateTargetsModelEditTargetGroupMemberWithDefaults() *BnsLoadBalancerV1ApiUpdateTargetsModelEditTargetGroupMember`

NewBnsLoadBalancerV1ApiUpdateTargetsModelEditTargetGroupMemberWithDefaults instantiates a new BnsLoadBalancerV1ApiUpdateTargetsModelEditTargetGroupMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BnsLoadBalancerV1ApiUpdateTargetsModelEditTargetGroupMember) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BnsLoadBalancerV1ApiUpdateTargetsModelEditTargetGroupMember) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BnsLoadBalancerV1ApiUpdateTargetsModelEditTargetGroupMember) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BnsLoadBalancerV1ApiUpdateTargetsModelEditTargetGroupMember) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BnsLoadBalancerV1ApiUpdateTargetsModelEditTargetGroupMember) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BnsLoadBalancerV1ApiUpdateTargetsModelEditTargetGroupMember) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAddress

`func (o *BnsLoadBalancerV1ApiUpdateTargetsModelEditTargetGroupMember) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *BnsLoadBalancerV1ApiUpdateTargetsModelEditTargetGroupMember) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *BnsLoadBalancerV1ApiUpdateTargetsModelEditTargetGroupMember) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetProtocolPort

`func (o *BnsLoadBalancerV1ApiUpdateTargetsModelEditTargetGroupMember) GetProtocolPort() int32`

GetProtocolPort returns the ProtocolPort field if non-nil, zero value otherwise.

### GetProtocolPortOk

`func (o *BnsLoadBalancerV1ApiUpdateTargetsModelEditTargetGroupMember) GetProtocolPortOk() (*int32, bool)`

GetProtocolPortOk returns a tuple with the ProtocolPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolPort

`func (o *BnsLoadBalancerV1ApiUpdateTargetsModelEditTargetGroupMember) SetProtocolPort(v int32)`

SetProtocolPort sets ProtocolPort field to given value.


### GetSubnetId

`func (o *BnsLoadBalancerV1ApiUpdateTargetsModelEditTargetGroupMember) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *BnsLoadBalancerV1ApiUpdateTargetsModelEditTargetGroupMember) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *BnsLoadBalancerV1ApiUpdateTargetsModelEditTargetGroupMember) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.


### GetWeight

`func (o *BnsLoadBalancerV1ApiUpdateTargetsModelEditTargetGroupMember) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *BnsLoadBalancerV1ApiUpdateTargetsModelEditTargetGroupMember) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *BnsLoadBalancerV1ApiUpdateTargetsModelEditTargetGroupMember) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *BnsLoadBalancerV1ApiUpdateTargetsModelEditTargetGroupMember) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### SetWeightNil

`func (o *BnsLoadBalancerV1ApiUpdateTargetsModelEditTargetGroupMember) SetWeightNil(b bool)`

 SetWeightNil sets the value for Weight to be an explicit nil

### UnsetWeight
`func (o *BnsLoadBalancerV1ApiUpdateTargetsModelEditTargetGroupMember) UnsetWeight()`

UnsetWeight ensures that no value is present for Weight, not even an explicit nil
### GetMonitorPort

`func (o *BnsLoadBalancerV1ApiUpdateTargetsModelEditTargetGroupMember) GetMonitorPort() int32`

GetMonitorPort returns the MonitorPort field if non-nil, zero value otherwise.

### GetMonitorPortOk

`func (o *BnsLoadBalancerV1ApiUpdateTargetsModelEditTargetGroupMember) GetMonitorPortOk() (*int32, bool)`

GetMonitorPortOk returns a tuple with the MonitorPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorPort

`func (o *BnsLoadBalancerV1ApiUpdateTargetsModelEditTargetGroupMember) SetMonitorPort(v int32)`

SetMonitorPort sets MonitorPort field to given value.

### HasMonitorPort

`func (o *BnsLoadBalancerV1ApiUpdateTargetsModelEditTargetGroupMember) HasMonitorPort() bool`

HasMonitorPort returns a boolean if a field has been set.

### SetMonitorPortNil

`func (o *BnsLoadBalancerV1ApiUpdateTargetsModelEditTargetGroupMember) SetMonitorPortNil(b bool)`

 SetMonitorPortNil sets the value for MonitorPort to be an explicit nil

### UnsetMonitorPort
`func (o *BnsLoadBalancerV1ApiUpdateTargetsModelEditTargetGroupMember) UnsetMonitorPort()`

UnsetMonitorPort ensures that no value is present for MonitorPort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


