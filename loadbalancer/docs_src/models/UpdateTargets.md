# UpdateTargets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | 대상 인스턴스의 이름 | [optional] 
**Address** | **string** | 대상 인스턴스의 IP 주소 | 
**ProtocolPort** | **int32** | 대상 인스턴스의 수신 포트 번호 | 
**SubnetId** | **string** | 대상 인스턴스가 속한 서브넷의 ID - [List subnets](/openapi/networking/vpc/list-subnets)에서 확인 | 
**Weight** | Pointer to **NullableInt32** | 트래픽 분산 가중치 | [optional] 
**MonitorPort** | Pointer to **NullableInt32** | 헬스 체크 요청을 보낼 포트 번호 | [optional] 

## Methods

### NewUpdateTargets

`func NewUpdateTargets(address string, protocolPort int32, subnetId string, ) *UpdateTargets`

NewUpdateTargets instantiates a new UpdateTargets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTargetsWithDefaults

`func NewUpdateTargetsWithDefaults() *UpdateTargets`

NewUpdateTargetsWithDefaults instantiates a new UpdateTargets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateTargets) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateTargets) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateTargets) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateTargets) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateTargets) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateTargets) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAddress

`func (o *UpdateTargets) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *UpdateTargets) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *UpdateTargets) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetProtocolPort

`func (o *UpdateTargets) GetProtocolPort() int32`

GetProtocolPort returns the ProtocolPort field if non-nil, zero value otherwise.

### GetProtocolPortOk

`func (o *UpdateTargets) GetProtocolPortOk() (*int32, bool)`

GetProtocolPortOk returns a tuple with the ProtocolPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolPort

`func (o *UpdateTargets) SetProtocolPort(v int32)`

SetProtocolPort sets ProtocolPort field to given value.


### GetSubnetId

`func (o *UpdateTargets) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *UpdateTargets) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *UpdateTargets) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.


### GetWeight

`func (o *UpdateTargets) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *UpdateTargets) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *UpdateTargets) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *UpdateTargets) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### SetWeightNil

`func (o *UpdateTargets) SetWeightNil(b bool)`

 SetWeightNil sets the value for Weight to be an explicit nil

### UnsetWeight
`func (o *UpdateTargets) UnsetWeight()`

UnsetWeight ensures that no value is present for Weight, not even an explicit nil
### GetMonitorPort

`func (o *UpdateTargets) GetMonitorPort() int32`

GetMonitorPort returns the MonitorPort field if non-nil, zero value otherwise.

### GetMonitorPortOk

`func (o *UpdateTargets) GetMonitorPortOk() (*int32, bool)`

GetMonitorPortOk returns a tuple with the MonitorPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorPort

`func (o *UpdateTargets) SetMonitorPort(v int32)`

SetMonitorPort sets MonitorPort field to given value.

### HasMonitorPort

`func (o *UpdateTargets) HasMonitorPort() bool`

HasMonitorPort returns a boolean if a field has been set.

### SetMonitorPortNil

`func (o *UpdateTargets) SetMonitorPortNil(b bool)`

 SetMonitorPortNil sets the value for MonitorPort to be an explicit nil

### UnsetMonitorPort
`func (o *UpdateTargets) UnsetMonitorPort()`

UnsetMonitorPort ensures that no value is present for MonitorPort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


