# ListenerModelInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | [**Protocol**](Protocol.md) | 리스너가 수신할 프로토콜 &lt;br/&gt; - &#x60;HTTP&#x60;: HTTP 프로토콜 &lt;br/&gt; - &#x60;TCP&#x60;: TCP 프로토콜 &lt;br/&gt; - &#x60;UDP&#x60;: UDP 프로토콜 &lt;br/&gt; - &#x60;TERMINATED_HTTPS&#x60;: SSL 종료된 HTTPS 프로토콜 | 
**ProtocolPort** | **int32** | 리스너가 수신할 포트 번호 | 
**DefaultTlsContainerRef** | Pointer to **NullableString** |  | [optional] 
**TlsMinVersion** | Pointer to [**NullableTLSVersion**](TLSVersion.md) |  | [optional] 

## Methods

### NewListenerModelInput

`func NewListenerModelInput(protocol Protocol, protocolPort int32, ) *ListenerModelInput`

NewListenerModelInput instantiates a new ListenerModelInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListenerModelInputWithDefaults

`func NewListenerModelInputWithDefaults() *ListenerModelInput`

NewListenerModelInputWithDefaults instantiates a new ListenerModelInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *ListenerModelInput) GetProtocol() Protocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ListenerModelInput) GetProtocolOk() (*Protocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ListenerModelInput) SetProtocol(v Protocol)`

SetProtocol sets Protocol field to given value.


### GetProtocolPort

`func (o *ListenerModelInput) GetProtocolPort() int32`

GetProtocolPort returns the ProtocolPort field if non-nil, zero value otherwise.

### GetProtocolPortOk

`func (o *ListenerModelInput) GetProtocolPortOk() (*int32, bool)`

GetProtocolPortOk returns a tuple with the ProtocolPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolPort

`func (o *ListenerModelInput) SetProtocolPort(v int32)`

SetProtocolPort sets ProtocolPort field to given value.


### GetDefaultTlsContainerRef

`func (o *ListenerModelInput) GetDefaultTlsContainerRef() string`

GetDefaultTlsContainerRef returns the DefaultTlsContainerRef field if non-nil, zero value otherwise.

### GetDefaultTlsContainerRefOk

`func (o *ListenerModelInput) GetDefaultTlsContainerRefOk() (*string, bool)`

GetDefaultTlsContainerRefOk returns a tuple with the DefaultTlsContainerRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTlsContainerRef

`func (o *ListenerModelInput) SetDefaultTlsContainerRef(v string)`

SetDefaultTlsContainerRef sets DefaultTlsContainerRef field to given value.

### HasDefaultTlsContainerRef

`func (o *ListenerModelInput) HasDefaultTlsContainerRef() bool`

HasDefaultTlsContainerRef returns a boolean if a field has been set.

### SetDefaultTlsContainerRefNil

`func (o *ListenerModelInput) SetDefaultTlsContainerRefNil(b bool)`

 SetDefaultTlsContainerRefNil sets the value for DefaultTlsContainerRef to be an explicit nil

### UnsetDefaultTlsContainerRef
`func (o *ListenerModelInput) UnsetDefaultTlsContainerRef()`

UnsetDefaultTlsContainerRef ensures that no value is present for DefaultTlsContainerRef, not even an explicit nil
### GetTlsMinVersion

`func (o *ListenerModelInput) GetTlsMinVersion() TLSVersion`

GetTlsMinVersion returns the TlsMinVersion field if non-nil, zero value otherwise.

### GetTlsMinVersionOk

`func (o *ListenerModelInput) GetTlsMinVersionOk() (*TLSVersion, bool)`

GetTlsMinVersionOk returns a tuple with the TlsMinVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsMinVersion

`func (o *ListenerModelInput) SetTlsMinVersion(v TLSVersion)`

SetTlsMinVersion sets TlsMinVersion field to given value.

### HasTlsMinVersion

`func (o *ListenerModelInput) HasTlsMinVersion() bool`

HasTlsMinVersion returns a boolean if a field has been set.

### SetTlsMinVersionNil

`func (o *ListenerModelInput) SetTlsMinVersionNil(b bool)`

 SetTlsMinVersionNil sets the value for TlsMinVersion to be an explicit nil

### UnsetTlsMinVersion
`func (o *ListenerModelInput) UnsetTlsMinVersion()`

UnsetTlsMinVersion ensures that no value is present for TlsMinVersion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


