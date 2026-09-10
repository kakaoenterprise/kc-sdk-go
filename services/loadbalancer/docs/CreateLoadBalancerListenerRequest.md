# CreateLoadBalancerListenerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | [**Protocol**](Protocol.md) | 리스너가 수신할 프로토콜 | 
**ProtocolPort** | **int32** | 리스너가 수신할 포트 번호 | 
**DefaultTlsContainerRef** | Pointer to **NullableString** |  | [optional] 
**TlsMinVersion** | Pointer to [**NullableTLSVersion**](TLSVersion.md) |  | [optional] 
**DefaultTargetGroup** | Pointer to [**NullablePoolRefRequest**](PoolRefRequest.md) |  | [optional] 
**L7Policies** | Pointer to [**[]CreateLoadBalancerL7PolicyRequest**](CreateLoadBalancerL7PolicyRequest.md) |  | [optional] 

## Methods

### NewCreateLoadBalancerListenerRequest

`func NewCreateLoadBalancerListenerRequest(protocol Protocol, protocolPort int32, ) *CreateLoadBalancerListenerRequest`

NewCreateLoadBalancerListenerRequest instantiates a new CreateLoadBalancerListenerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLoadBalancerListenerRequestWithDefaults

`func NewCreateLoadBalancerListenerRequestWithDefaults() *CreateLoadBalancerListenerRequest`

NewCreateLoadBalancerListenerRequestWithDefaults instantiates a new CreateLoadBalancerListenerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *CreateLoadBalancerListenerRequest) GetProtocol() Protocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *CreateLoadBalancerListenerRequest) GetProtocolOk() (*Protocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *CreateLoadBalancerListenerRequest) SetProtocol(v Protocol)`

SetProtocol sets Protocol field to given value.


### GetProtocolPort

`func (o *CreateLoadBalancerListenerRequest) GetProtocolPort() int32`

GetProtocolPort returns the ProtocolPort field if non-nil, zero value otherwise.

### GetProtocolPortOk

`func (o *CreateLoadBalancerListenerRequest) GetProtocolPortOk() (*int32, bool)`

GetProtocolPortOk returns a tuple with the ProtocolPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolPort

`func (o *CreateLoadBalancerListenerRequest) SetProtocolPort(v int32)`

SetProtocolPort sets ProtocolPort field to given value.


### GetDefaultTlsContainerRef

`func (o *CreateLoadBalancerListenerRequest) GetDefaultTlsContainerRef() string`

GetDefaultTlsContainerRef returns the DefaultTlsContainerRef field if non-nil, zero value otherwise.

### GetDefaultTlsContainerRefOk

`func (o *CreateLoadBalancerListenerRequest) GetDefaultTlsContainerRefOk() (*string, bool)`

GetDefaultTlsContainerRefOk returns a tuple with the DefaultTlsContainerRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTlsContainerRef

`func (o *CreateLoadBalancerListenerRequest) SetDefaultTlsContainerRef(v string)`

SetDefaultTlsContainerRef sets DefaultTlsContainerRef field to given value.

### HasDefaultTlsContainerRef

`func (o *CreateLoadBalancerListenerRequest) HasDefaultTlsContainerRef() bool`

HasDefaultTlsContainerRef returns a boolean if a field has been set.

### SetDefaultTlsContainerRefNil

`func (o *CreateLoadBalancerListenerRequest) SetDefaultTlsContainerRefNil(b bool)`

 SetDefaultTlsContainerRefNil sets the value for DefaultTlsContainerRef to be an explicit nil

### UnsetDefaultTlsContainerRef
`func (o *CreateLoadBalancerListenerRequest) UnsetDefaultTlsContainerRef()`

UnsetDefaultTlsContainerRef ensures that no value is present for DefaultTlsContainerRef, not even an explicit nil
### GetTlsMinVersion

`func (o *CreateLoadBalancerListenerRequest) GetTlsMinVersion() TLSVersion`

GetTlsMinVersion returns the TlsMinVersion field if non-nil, zero value otherwise.

### GetTlsMinVersionOk

`func (o *CreateLoadBalancerListenerRequest) GetTlsMinVersionOk() (*TLSVersion, bool)`

GetTlsMinVersionOk returns a tuple with the TlsMinVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsMinVersion

`func (o *CreateLoadBalancerListenerRequest) SetTlsMinVersion(v TLSVersion)`

SetTlsMinVersion sets TlsMinVersion field to given value.

### HasTlsMinVersion

`func (o *CreateLoadBalancerListenerRequest) HasTlsMinVersion() bool`

HasTlsMinVersion returns a boolean if a field has been set.

### SetTlsMinVersionNil

`func (o *CreateLoadBalancerListenerRequest) SetTlsMinVersionNil(b bool)`

 SetTlsMinVersionNil sets the value for TlsMinVersion to be an explicit nil

### UnsetTlsMinVersion
`func (o *CreateLoadBalancerListenerRequest) UnsetTlsMinVersion()`

UnsetTlsMinVersion ensures that no value is present for TlsMinVersion, not even an explicit nil
### GetDefaultTargetGroup

`func (o *CreateLoadBalancerListenerRequest) GetDefaultTargetGroup() PoolRefRequest`

GetDefaultTargetGroup returns the DefaultTargetGroup field if non-nil, zero value otherwise.

### GetDefaultTargetGroupOk

`func (o *CreateLoadBalancerListenerRequest) GetDefaultTargetGroupOk() (*PoolRefRequest, bool)`

GetDefaultTargetGroupOk returns a tuple with the DefaultTargetGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTargetGroup

`func (o *CreateLoadBalancerListenerRequest) SetDefaultTargetGroup(v PoolRefRequest)`

SetDefaultTargetGroup sets DefaultTargetGroup field to given value.

### HasDefaultTargetGroup

`func (o *CreateLoadBalancerListenerRequest) HasDefaultTargetGroup() bool`

HasDefaultTargetGroup returns a boolean if a field has been set.

### SetDefaultTargetGroupNil

`func (o *CreateLoadBalancerListenerRequest) SetDefaultTargetGroupNil(b bool)`

 SetDefaultTargetGroupNil sets the value for DefaultTargetGroup to be an explicit nil

### UnsetDefaultTargetGroup
`func (o *CreateLoadBalancerListenerRequest) UnsetDefaultTargetGroup()`

UnsetDefaultTargetGroup ensures that no value is present for DefaultTargetGroup, not even an explicit nil
### GetL7Policies

`func (o *CreateLoadBalancerListenerRequest) GetL7Policies() []CreateLoadBalancerL7PolicyRequest`

GetL7Policies returns the L7Policies field if non-nil, zero value otherwise.

### GetL7PoliciesOk

`func (o *CreateLoadBalancerListenerRequest) GetL7PoliciesOk() (*[]CreateLoadBalancerL7PolicyRequest, bool)`

GetL7PoliciesOk returns a tuple with the L7Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL7Policies

`func (o *CreateLoadBalancerListenerRequest) SetL7Policies(v []CreateLoadBalancerL7PolicyRequest)`

SetL7Policies sets L7Policies field to given value.

### HasL7Policies

`func (o *CreateLoadBalancerListenerRequest) HasL7Policies() bool`

HasL7Policies returns a boolean if a field has been set.

### SetL7PoliciesNil

`func (o *CreateLoadBalancerListenerRequest) SetL7PoliciesNil(b bool)`

 SetL7PoliciesNil sets the value for L7Policies to be an explicit nil

### UnsetL7Policies
`func (o *CreateLoadBalancerListenerRequest) UnsetL7Policies()`

UnsetL7Policies ensures that no value is present for L7Policies, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


