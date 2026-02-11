# CreateTargetGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** |  | [optional] 
**LoadBalancerAlgorithm** | [**TargetGroupAlgorithm**](TargetGroupAlgorithm.md) | 로드 밸런싱 알고리즘 &lt;br/&gt; - &#x60;ROUND_ROBIN&#x60;: 라운드 로빈 방식 &lt;br/&gt; - &#x60;LEAST_CONNECTIONS&#x60;: 최소 연결 방식 &lt;br/&gt; - &#x60;SOURCE_IP&#x60;: 소스 IP 기반 방식 | 
**ListenerId** | Pointer to **NullableString** |  | [optional] 
**LoadBalancerId** | **string** | 연결할 로드 밸런서 ID | 
**Name** | **string** | 대상 그룹의 이름 | 
**Protocol** | [**TargetGroupProtocol**](TargetGroupProtocol.md) | 대상 그룹의 백엔드 통신에 사용할 프로토콜 &lt;br/&gt; - &#x60;HTTP&#x60;: HTTP 프로토콜 &lt;br/&gt; - &#x60;HTTPS&#x60;: HTTPS 프로토콜 &lt;br/&gt; - &#x60;TCP&#x60;: TCP 프로토콜 &lt;br/&gt; - &#x60;UDP&#x60;: UDP 프로토콜 &lt;br/&gt; - &#x60;PROXY&#x60;: 프록시 프로토콜 | 
**AlpnProtocols** | Pointer to [**NullableAlpnProtocol**](AlpnProtocol.md) |  | [optional] 
**SessionPersistence** | Pointer to [**NullableSessionPersistenceModel**](SessionPersistenceModel.md) |  | [optional] 

## Methods

### NewCreateTargetGroup

`func NewCreateTargetGroup(loadBalancerAlgorithm TargetGroupAlgorithm, loadBalancerId string, name string, protocol TargetGroupProtocol, ) *CreateTargetGroup`

NewCreateTargetGroup instantiates a new CreateTargetGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTargetGroupWithDefaults

`func NewCreateTargetGroupWithDefaults() *CreateTargetGroup`

NewCreateTargetGroupWithDefaults instantiates a new CreateTargetGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateTargetGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateTargetGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateTargetGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateTargetGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateTargetGroup) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateTargetGroup) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLoadBalancerAlgorithm

`func (o *CreateTargetGroup) GetLoadBalancerAlgorithm() TargetGroupAlgorithm`

GetLoadBalancerAlgorithm returns the LoadBalancerAlgorithm field if non-nil, zero value otherwise.

### GetLoadBalancerAlgorithmOk

`func (o *CreateTargetGroup) GetLoadBalancerAlgorithmOk() (*TargetGroupAlgorithm, bool)`

GetLoadBalancerAlgorithmOk returns a tuple with the LoadBalancerAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerAlgorithm

`func (o *CreateTargetGroup) SetLoadBalancerAlgorithm(v TargetGroupAlgorithm)`

SetLoadBalancerAlgorithm sets LoadBalancerAlgorithm field to given value.


### GetListenerId

`func (o *CreateTargetGroup) GetListenerId() string`

GetListenerId returns the ListenerId field if non-nil, zero value otherwise.

### GetListenerIdOk

`func (o *CreateTargetGroup) GetListenerIdOk() (*string, bool)`

GetListenerIdOk returns a tuple with the ListenerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenerId

`func (o *CreateTargetGroup) SetListenerId(v string)`

SetListenerId sets ListenerId field to given value.

### HasListenerId

`func (o *CreateTargetGroup) HasListenerId() bool`

HasListenerId returns a boolean if a field has been set.

### SetListenerIdNil

`func (o *CreateTargetGroup) SetListenerIdNil(b bool)`

 SetListenerIdNil sets the value for ListenerId to be an explicit nil

### UnsetListenerId
`func (o *CreateTargetGroup) UnsetListenerId()`

UnsetListenerId ensures that no value is present for ListenerId, not even an explicit nil
### GetLoadBalancerId

`func (o *CreateTargetGroup) GetLoadBalancerId() string`

GetLoadBalancerId returns the LoadBalancerId field if non-nil, zero value otherwise.

### GetLoadBalancerIdOk

`func (o *CreateTargetGroup) GetLoadBalancerIdOk() (*string, bool)`

GetLoadBalancerIdOk returns a tuple with the LoadBalancerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerId

`func (o *CreateTargetGroup) SetLoadBalancerId(v string)`

SetLoadBalancerId sets LoadBalancerId field to given value.


### GetName

`func (o *CreateTargetGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTargetGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTargetGroup) SetName(v string)`

SetName sets Name field to given value.


### GetProtocol

`func (o *CreateTargetGroup) GetProtocol() TargetGroupProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *CreateTargetGroup) GetProtocolOk() (*TargetGroupProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *CreateTargetGroup) SetProtocol(v TargetGroupProtocol)`

SetProtocol sets Protocol field to given value.


### GetAlpnProtocols

`func (o *CreateTargetGroup) GetAlpnProtocols() AlpnProtocol`

GetAlpnProtocols returns the AlpnProtocols field if non-nil, zero value otherwise.

### GetAlpnProtocolsOk

`func (o *CreateTargetGroup) GetAlpnProtocolsOk() (*AlpnProtocol, bool)`

GetAlpnProtocolsOk returns a tuple with the AlpnProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlpnProtocols

`func (o *CreateTargetGroup) SetAlpnProtocols(v AlpnProtocol)`

SetAlpnProtocols sets AlpnProtocols field to given value.

### HasAlpnProtocols

`func (o *CreateTargetGroup) HasAlpnProtocols() bool`

HasAlpnProtocols returns a boolean if a field has been set.

### SetAlpnProtocolsNil

`func (o *CreateTargetGroup) SetAlpnProtocolsNil(b bool)`

 SetAlpnProtocolsNil sets the value for AlpnProtocols to be an explicit nil

### UnsetAlpnProtocols
`func (o *CreateTargetGroup) UnsetAlpnProtocols()`

UnsetAlpnProtocols ensures that no value is present for AlpnProtocols, not even an explicit nil
### GetSessionPersistence

`func (o *CreateTargetGroup) GetSessionPersistence() SessionPersistenceModel`

GetSessionPersistence returns the SessionPersistence field if non-nil, zero value otherwise.

### GetSessionPersistenceOk

`func (o *CreateTargetGroup) GetSessionPersistenceOk() (*SessionPersistenceModel, bool)`

GetSessionPersistenceOk returns a tuple with the SessionPersistence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionPersistence

`func (o *CreateTargetGroup) SetSessionPersistence(v SessionPersistenceModel)`

SetSessionPersistence sets SessionPersistence field to given value.

### HasSessionPersistence

`func (o *CreateTargetGroup) HasSessionPersistence() bool`

HasSessionPersistence returns a boolean if a field has been set.

### SetSessionPersistenceNil

`func (o *CreateTargetGroup) SetSessionPersistenceNil(b bool)`

 SetSessionPersistenceNil sets the value for SessionPersistence to be an explicit nil

### UnsetSessionPersistence
`func (o *CreateTargetGroup) UnsetSessionPersistence()`

UnsetSessionPersistence ensures that no value is present for SessionPersistence, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


