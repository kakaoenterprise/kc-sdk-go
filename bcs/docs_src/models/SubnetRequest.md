# SubnetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | 인스턴스를 연결할 서브넷의 ID - [List subnets](/openapi/networking/vpc/list-subnets)에서 확인 | [optional] 
**NetworkInterfaceId** | Pointer to **NullableString** | 사용할 네트워크 인터페이스의 ID - &#x60;private_ip&#x60;와 동시에 사용할 수 없음 - [List network interfaces](/openapi/networking/vpc/list-network-interfaces)에서 확인 | [optional] 
**PrivateIp** | Pointer to **NullableString** | 필요 시, 인스턴스에 수동으로 프라이빗 IP를 할당 - 자동 할당이 아닌 고정 IP 구성이 필요할 경우 사용 - &#x60;network_interface_id&#x60;와 동시에 사용할 수 없음 | [optional] 

## Methods

### NewSubnetRequest

`func NewSubnetRequest() *SubnetRequest`

NewSubnetRequest instantiates a new SubnetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubnetRequestWithDefaults

`func NewSubnetRequestWithDefaults() *SubnetRequest`

NewSubnetRequestWithDefaults instantiates a new SubnetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubnetRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubnetRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubnetRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SubnetRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SubnetRequest) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SubnetRequest) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetNetworkInterfaceId

`func (o *SubnetRequest) GetNetworkInterfaceId() string`

GetNetworkInterfaceId returns the NetworkInterfaceId field if non-nil, zero value otherwise.

### GetNetworkInterfaceIdOk

`func (o *SubnetRequest) GetNetworkInterfaceIdOk() (*string, bool)`

GetNetworkInterfaceIdOk returns a tuple with the NetworkInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaceId

`func (o *SubnetRequest) SetNetworkInterfaceId(v string)`

SetNetworkInterfaceId sets NetworkInterfaceId field to given value.

### HasNetworkInterfaceId

`func (o *SubnetRequest) HasNetworkInterfaceId() bool`

HasNetworkInterfaceId returns a boolean if a field has been set.

### SetNetworkInterfaceIdNil

`func (o *SubnetRequest) SetNetworkInterfaceIdNil(b bool)`

 SetNetworkInterfaceIdNil sets the value for NetworkInterfaceId to be an explicit nil

### UnsetNetworkInterfaceId
`func (o *SubnetRequest) UnsetNetworkInterfaceId()`

UnsetNetworkInterfaceId ensures that no value is present for NetworkInterfaceId, not even an explicit nil
### GetPrivateIp

`func (o *SubnetRequest) GetPrivateIp() string`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *SubnetRequest) GetPrivateIpOk() (*string, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *SubnetRequest) SetPrivateIp(v string)`

SetPrivateIp sets PrivateIp field to given value.

### HasPrivateIp

`func (o *SubnetRequest) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.

### SetPrivateIpNil

`func (o *SubnetRequest) SetPrivateIpNil(b bool)`

 SetPrivateIpNil sets the value for PrivateIp to be an explicit nil

### UnsetPrivateIp
`func (o *SubnetRequest) UnsetPrivateIp()`

UnsetPrivateIp ensures that no value is present for PrivateIp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


