# CreateListener

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadBalancerId** | **string** | 리스너를 생성할 대상 로드 밸런서 ID | 
**Protocol** | [**Protocol**](Protocol.md) | 리스너가 수신할 트래픽의 프로토콜 &lt;br/&gt; - &#x60;HTTP&#x60;: HTTP 프로토콜 &lt;br/&gt; - &#x60;TCP&#x60;: TCP 프로토콜 &lt;br/&gt; - &#x60;UDP&#x60;: UDP 프로토콜 &lt;br/&gt; - &#x60;TERMINATED_HTTPS&#x60;: SSL 종료된 HTTPS 프로토콜 | 
**ProtocolPort** | **int32** | 수신 포트 번호 | 
**TargetGroupId** | Pointer to **NullableString** |  | [optional] 
**DefaultTlsContainerRef** | Pointer to **NullableString** |  | [optional] 
**TlsMinVersion** | Pointer to [**NullableTLSVersion**](TLSVersion.md) |  | [optional] 

## Methods

### NewCreateListener

`func NewCreateListener(loadBalancerId string, protocol Protocol, protocolPort int32, ) *CreateListener`

NewCreateListener instantiates a new CreateListener object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateListenerWithDefaults

`func NewCreateListenerWithDefaults() *CreateListener`

NewCreateListenerWithDefaults instantiates a new CreateListener object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoadBalancerId

`func (o *CreateListener) GetLoadBalancerId() string`

GetLoadBalancerId returns the LoadBalancerId field if non-nil, zero value otherwise.

### GetLoadBalancerIdOk

`func (o *CreateListener) GetLoadBalancerIdOk() (*string, bool)`

GetLoadBalancerIdOk returns a tuple with the LoadBalancerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerId

`func (o *CreateListener) SetLoadBalancerId(v string)`

SetLoadBalancerId sets LoadBalancerId field to given value.


### GetProtocol

`func (o *CreateListener) GetProtocol() Protocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *CreateListener) GetProtocolOk() (*Protocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *CreateListener) SetProtocol(v Protocol)`

SetProtocol sets Protocol field to given value.


### GetProtocolPort

`func (o *CreateListener) GetProtocolPort() int32`

GetProtocolPort returns the ProtocolPort field if non-nil, zero value otherwise.

### GetProtocolPortOk

`func (o *CreateListener) GetProtocolPortOk() (*int32, bool)`

GetProtocolPortOk returns a tuple with the ProtocolPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolPort

`func (o *CreateListener) SetProtocolPort(v int32)`

SetProtocolPort sets ProtocolPort field to given value.


### GetTargetGroupId

`func (o *CreateListener) GetTargetGroupId() string`

GetTargetGroupId returns the TargetGroupId field if non-nil, zero value otherwise.

### GetTargetGroupIdOk

`func (o *CreateListener) GetTargetGroupIdOk() (*string, bool)`

GetTargetGroupIdOk returns a tuple with the TargetGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroupId

`func (o *CreateListener) SetTargetGroupId(v string)`

SetTargetGroupId sets TargetGroupId field to given value.

### HasTargetGroupId

`func (o *CreateListener) HasTargetGroupId() bool`

HasTargetGroupId returns a boolean if a field has been set.

### SetTargetGroupIdNil

`func (o *CreateListener) SetTargetGroupIdNil(b bool)`

 SetTargetGroupIdNil sets the value for TargetGroupId to be an explicit nil

### UnsetTargetGroupId
`func (o *CreateListener) UnsetTargetGroupId()`

UnsetTargetGroupId ensures that no value is present for TargetGroupId, not even an explicit nil
### GetDefaultTlsContainerRef

`func (o *CreateListener) GetDefaultTlsContainerRef() string`

GetDefaultTlsContainerRef returns the DefaultTlsContainerRef field if non-nil, zero value otherwise.

### GetDefaultTlsContainerRefOk

`func (o *CreateListener) GetDefaultTlsContainerRefOk() (*string, bool)`

GetDefaultTlsContainerRefOk returns a tuple with the DefaultTlsContainerRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTlsContainerRef

`func (o *CreateListener) SetDefaultTlsContainerRef(v string)`

SetDefaultTlsContainerRef sets DefaultTlsContainerRef field to given value.

### HasDefaultTlsContainerRef

`func (o *CreateListener) HasDefaultTlsContainerRef() bool`

HasDefaultTlsContainerRef returns a boolean if a field has been set.

### SetDefaultTlsContainerRefNil

`func (o *CreateListener) SetDefaultTlsContainerRefNil(b bool)`

 SetDefaultTlsContainerRefNil sets the value for DefaultTlsContainerRef to be an explicit nil

### UnsetDefaultTlsContainerRef
`func (o *CreateListener) UnsetDefaultTlsContainerRef()`

UnsetDefaultTlsContainerRef ensures that no value is present for DefaultTlsContainerRef, not even an explicit nil
### GetTlsMinVersion

`func (o *CreateListener) GetTlsMinVersion() TLSVersion`

GetTlsMinVersion returns the TlsMinVersion field if non-nil, zero value otherwise.

### GetTlsMinVersionOk

`func (o *CreateListener) GetTlsMinVersionOk() (*TLSVersion, bool)`

GetTlsMinVersionOk returns a tuple with the TlsMinVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsMinVersion

`func (o *CreateListener) SetTlsMinVersion(v TLSVersion)`

SetTlsMinVersion sets TlsMinVersion field to given value.

### HasTlsMinVersion

`func (o *CreateListener) HasTlsMinVersion() bool`

HasTlsMinVersion returns a boolean if a field has been set.

### SetTlsMinVersionNil

`func (o *CreateListener) SetTlsMinVersionNil(b bool)`

 SetTlsMinVersionNil sets the value for TlsMinVersion to be an explicit nil

### UnsetTlsMinVersion
`func (o *CreateListener) UnsetTlsMinVersion()`

UnsetTlsMinVersion ensures that no value is present for TlsMinVersion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


