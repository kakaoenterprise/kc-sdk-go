# AddTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Address** | **string** | 대상 인스턴스의 IP 주소 | 
**ProtocolPort** | **int32** | 트래픽 수신에 사용할 포트 번호 | 
**SubnetId** | **string** | 대상 인스턴스가 위치한 서브넷 ID &lt;br/&gt;- [List subnets](https://docs.kakaocloud.com/openapi/networking/vpc/list-subnets)에서 확인 | 
**Weight** | Pointer to **NullableInt32** |  | [optional] 
**MonitorPort** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewAddTarget

`func NewAddTarget(address string, protocolPort int32, subnetId string, ) *AddTarget`

NewAddTarget instantiates a new AddTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddTargetWithDefaults

`func NewAddTargetWithDefaults() *AddTarget`

NewAddTargetWithDefaults instantiates a new AddTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddTarget) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddTarget) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AddTarget) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AddTarget) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAddress

`func (o *AddTarget) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AddTarget) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AddTarget) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetProtocolPort

`func (o *AddTarget) GetProtocolPort() int32`

GetProtocolPort returns the ProtocolPort field if non-nil, zero value otherwise.

### GetProtocolPortOk

`func (o *AddTarget) GetProtocolPortOk() (*int32, bool)`

GetProtocolPortOk returns a tuple with the ProtocolPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolPort

`func (o *AddTarget) SetProtocolPort(v int32)`

SetProtocolPort sets ProtocolPort field to given value.


### GetSubnetId

`func (o *AddTarget) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *AddTarget) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *AddTarget) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.


### GetWeight

`func (o *AddTarget) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *AddTarget) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *AddTarget) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *AddTarget) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### SetWeightNil

`func (o *AddTarget) SetWeightNil(b bool)`

 SetWeightNil sets the value for Weight to be an explicit nil

### UnsetWeight
`func (o *AddTarget) UnsetWeight()`

UnsetWeight ensures that no value is present for Weight, not even an explicit nil
### GetMonitorPort

`func (o *AddTarget) GetMonitorPort() int32`

GetMonitorPort returns the MonitorPort field if non-nil, zero value otherwise.

### GetMonitorPortOk

`func (o *AddTarget) GetMonitorPortOk() (*int32, bool)`

GetMonitorPortOk returns a tuple with the MonitorPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorPort

`func (o *AddTarget) SetMonitorPort(v int32)`

SetMonitorPort sets MonitorPort field to given value.

### HasMonitorPort

`func (o *AddTarget) HasMonitorPort() bool`

HasMonitorPort returns a boolean if a field has been set.

### SetMonitorPortNil

`func (o *AddTarget) SetMonitorPortNil(b bool)`

 SetMonitorPortNil sets the value for MonitorPort to be an explicit nil

### UnsetMonitorPort
`func (o *AddTarget) UnsetMonitorPort()`

UnsetMonitorPort ensures that no value is present for MonitorPort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


