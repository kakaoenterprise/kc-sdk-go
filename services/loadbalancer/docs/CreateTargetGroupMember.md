# CreateTargetGroupMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Address** | **string** | 대상 인스턴스의 IP 주소 | 
**ProtocolPort** | **int32** | 트래픽 수신에 사용할 포트 번호 | 
**SubnetId** | **string** | 대상 인스턴스가 위치한 서브넷 ID | 
**Weight** | Pointer to **int32** | 트래픽 분산 가중치 | [optional] [default to 1]
**MonitorPort** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewCreateTargetGroupMember

`func NewCreateTargetGroupMember(address string, protocolPort int32, subnetId string, ) *CreateTargetGroupMember`

NewCreateTargetGroupMember instantiates a new CreateTargetGroupMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTargetGroupMemberWithDefaults

`func NewCreateTargetGroupMemberWithDefaults() *CreateTargetGroupMember`

NewCreateTargetGroupMemberWithDefaults instantiates a new CreateTargetGroupMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateTargetGroupMember) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTargetGroupMember) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTargetGroupMember) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateTargetGroupMember) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreateTargetGroupMember) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateTargetGroupMember) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAddress

`func (o *CreateTargetGroupMember) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CreateTargetGroupMember) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CreateTargetGroupMember) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetProtocolPort

`func (o *CreateTargetGroupMember) GetProtocolPort() int32`

GetProtocolPort returns the ProtocolPort field if non-nil, zero value otherwise.

### GetProtocolPortOk

`func (o *CreateTargetGroupMember) GetProtocolPortOk() (*int32, bool)`

GetProtocolPortOk returns a tuple with the ProtocolPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolPort

`func (o *CreateTargetGroupMember) SetProtocolPort(v int32)`

SetProtocolPort sets ProtocolPort field to given value.


### GetSubnetId

`func (o *CreateTargetGroupMember) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *CreateTargetGroupMember) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *CreateTargetGroupMember) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.


### GetWeight

`func (o *CreateTargetGroupMember) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *CreateTargetGroupMember) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *CreateTargetGroupMember) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *CreateTargetGroupMember) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetMonitorPort

`func (o *CreateTargetGroupMember) GetMonitorPort() int32`

GetMonitorPort returns the MonitorPort field if non-nil, zero value otherwise.

### GetMonitorPortOk

`func (o *CreateTargetGroupMember) GetMonitorPortOk() (*int32, bool)`

GetMonitorPortOk returns a tuple with the MonitorPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorPort

`func (o *CreateTargetGroupMember) SetMonitorPort(v int32)`

SetMonitorPort sets MonitorPort field to given value.

### HasMonitorPort

`func (o *CreateTargetGroupMember) HasMonitorPort() bool`

HasMonitorPort returns a boolean if a field has been set.

### SetMonitorPortNil

`func (o *CreateTargetGroupMember) SetMonitorPortNil(b bool)`

 SetMonitorPortNil sets the value for MonitorPort to be an explicit nil

### UnsetMonitorPort
`func (o *CreateTargetGroupMember) UnsetMonitorPort()`

UnsetMonitorPort ensures that no value is present for MonitorPort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


