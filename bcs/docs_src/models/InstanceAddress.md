# InstanceAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkInterfaceId** | Pointer to **NullableString** | 네트워크 인터페이스 ID | [optional] 
**PrivateIp** | Pointer to **NullableString** | 프라이빗 IP 주소 | [optional] 
**PublicIp** | Pointer to **NullableString** | 퍼블릭 IP 주소 | [optional] 

## Methods

### NewInstanceAddress

`func NewInstanceAddress() *InstanceAddress`

NewInstanceAddress instantiates a new InstanceAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceAddressWithDefaults

`func NewInstanceAddressWithDefaults() *InstanceAddress`

NewInstanceAddressWithDefaults instantiates a new InstanceAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkInterfaceId

`func (o *InstanceAddress) GetNetworkInterfaceId() string`

GetNetworkInterfaceId returns the NetworkInterfaceId field if non-nil, zero value otherwise.

### GetNetworkInterfaceIdOk

`func (o *InstanceAddress) GetNetworkInterfaceIdOk() (*string, bool)`

GetNetworkInterfaceIdOk returns a tuple with the NetworkInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaceId

`func (o *InstanceAddress) SetNetworkInterfaceId(v string)`

SetNetworkInterfaceId sets NetworkInterfaceId field to given value.

### HasNetworkInterfaceId

`func (o *InstanceAddress) HasNetworkInterfaceId() bool`

HasNetworkInterfaceId returns a boolean if a field has been set.

### SetNetworkInterfaceIdNil

`func (o *InstanceAddress) SetNetworkInterfaceIdNil(b bool)`

 SetNetworkInterfaceIdNil sets the value for NetworkInterfaceId to be an explicit nil

### UnsetNetworkInterfaceId
`func (o *InstanceAddress) UnsetNetworkInterfaceId()`

UnsetNetworkInterfaceId ensures that no value is present for NetworkInterfaceId, not even an explicit nil
### GetPrivateIp

`func (o *InstanceAddress) GetPrivateIp() string`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *InstanceAddress) GetPrivateIpOk() (*string, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *InstanceAddress) SetPrivateIp(v string)`

SetPrivateIp sets PrivateIp field to given value.

### HasPrivateIp

`func (o *InstanceAddress) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.

### SetPrivateIpNil

`func (o *InstanceAddress) SetPrivateIpNil(b bool)`

 SetPrivateIpNil sets the value for PrivateIp to be an explicit nil

### UnsetPrivateIp
`func (o *InstanceAddress) UnsetPrivateIp()`

UnsetPrivateIp ensures that no value is present for PrivateIp, not even an explicit nil
### GetPublicIp

`func (o *InstanceAddress) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *InstanceAddress) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *InstanceAddress) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *InstanceAddress) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### SetPublicIpNil

`func (o *InstanceAddress) SetPublicIpNil(b bool)`

 SetPublicIpNil sets the value for PublicIp to be an explicit nil

### UnsetPublicIp
`func (o *InstanceAddress) UnsetPublicIp()`

UnsetPublicIp ensures that no value is present for PublicIp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


