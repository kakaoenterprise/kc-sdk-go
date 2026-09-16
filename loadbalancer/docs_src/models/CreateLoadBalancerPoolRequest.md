# CreateLoadBalancerPoolRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | 생성할 대상 그룹 이름 | 
**Description** | Pointer to **NullableString** | 대상 그룹 설명 | [optional] 
**Protocol** | [**LoadBalancerPoolProtocol**](LoadBalancerPoolProtocol.md) | 대상 그룹 통신 프로토콜 | 
**LoadBalancerAlgorithm** | [**LoadBalancerPoolAlgorithm**](LoadBalancerPoolAlgorithm.md) | 트래픽 분산 알고리즘 | 
**Healthmonitor** | Pointer to [**NullableCreateLoadBalancerHealthMonitorRequest**](CreateLoadBalancerHealthMonitorRequest.md) | 헬스 체크 설정 | [optional] 
**Members** | Pointer to [**[]CreateLoadBalancerMemberRequest**](CreateLoadBalancerMemberRequest.md) | 등록할 서버 목록 | [optional] 
**AlpnProtocols** | Pointer to [**NullableLoadBalancerAlpnProtocol**](LoadBalancerAlpnProtocol.md) | ALPN 프로토콜 설정 | [optional] 
**SessionPersistence** | Pointer to [**NullableSessionPersistenceRequest**](SessionPersistenceRequest.md) | 세션 유지 설정 | [optional] 

## Methods

### NewCreateLoadBalancerPoolRequest

`func NewCreateLoadBalancerPoolRequest(name string, protocol LoadBalancerPoolProtocol, loadBalancerAlgorithm LoadBalancerPoolAlgorithm, ) *CreateLoadBalancerPoolRequest`

NewCreateLoadBalancerPoolRequest instantiates a new CreateLoadBalancerPoolRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLoadBalancerPoolRequestWithDefaults

`func NewCreateLoadBalancerPoolRequestWithDefaults() *CreateLoadBalancerPoolRequest`

NewCreateLoadBalancerPoolRequestWithDefaults instantiates a new CreateLoadBalancerPoolRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateLoadBalancerPoolRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateLoadBalancerPoolRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateLoadBalancerPoolRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateLoadBalancerPoolRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateLoadBalancerPoolRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateLoadBalancerPoolRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateLoadBalancerPoolRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateLoadBalancerPoolRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateLoadBalancerPoolRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetProtocol

`func (o *CreateLoadBalancerPoolRequest) GetProtocol() LoadBalancerPoolProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *CreateLoadBalancerPoolRequest) GetProtocolOk() (*LoadBalancerPoolProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *CreateLoadBalancerPoolRequest) SetProtocol(v LoadBalancerPoolProtocol)`

SetProtocol sets Protocol field to given value.


### GetLoadBalancerAlgorithm

`func (o *CreateLoadBalancerPoolRequest) GetLoadBalancerAlgorithm() LoadBalancerPoolAlgorithm`

GetLoadBalancerAlgorithm returns the LoadBalancerAlgorithm field if non-nil, zero value otherwise.

### GetLoadBalancerAlgorithmOk

`func (o *CreateLoadBalancerPoolRequest) GetLoadBalancerAlgorithmOk() (*LoadBalancerPoolAlgorithm, bool)`

GetLoadBalancerAlgorithmOk returns a tuple with the LoadBalancerAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerAlgorithm

`func (o *CreateLoadBalancerPoolRequest) SetLoadBalancerAlgorithm(v LoadBalancerPoolAlgorithm)`

SetLoadBalancerAlgorithm sets LoadBalancerAlgorithm field to given value.


### GetHealthmonitor

`func (o *CreateLoadBalancerPoolRequest) GetHealthmonitor() CreateLoadBalancerHealthMonitorRequest`

GetHealthmonitor returns the Healthmonitor field if non-nil, zero value otherwise.

### GetHealthmonitorOk

`func (o *CreateLoadBalancerPoolRequest) GetHealthmonitorOk() (*CreateLoadBalancerHealthMonitorRequest, bool)`

GetHealthmonitorOk returns a tuple with the Healthmonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthmonitor

`func (o *CreateLoadBalancerPoolRequest) SetHealthmonitor(v CreateLoadBalancerHealthMonitorRequest)`

SetHealthmonitor sets Healthmonitor field to given value.

### HasHealthmonitor

`func (o *CreateLoadBalancerPoolRequest) HasHealthmonitor() bool`

HasHealthmonitor returns a boolean if a field has been set.

### SetHealthmonitorNil

`func (o *CreateLoadBalancerPoolRequest) SetHealthmonitorNil(b bool)`

 SetHealthmonitorNil sets the value for Healthmonitor to be an explicit nil

### UnsetHealthmonitor
`func (o *CreateLoadBalancerPoolRequest) UnsetHealthmonitor()`

UnsetHealthmonitor ensures that no value is present for Healthmonitor, not even an explicit nil
### GetMembers

`func (o *CreateLoadBalancerPoolRequest) GetMembers() []CreateLoadBalancerMemberRequest`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *CreateLoadBalancerPoolRequest) GetMembersOk() (*[]CreateLoadBalancerMemberRequest, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *CreateLoadBalancerPoolRequest) SetMembers(v []CreateLoadBalancerMemberRequest)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *CreateLoadBalancerPoolRequest) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### SetMembersNil

`func (o *CreateLoadBalancerPoolRequest) SetMembersNil(b bool)`

 SetMembersNil sets the value for Members to be an explicit nil

### UnsetMembers
`func (o *CreateLoadBalancerPoolRequest) UnsetMembers()`

UnsetMembers ensures that no value is present for Members, not even an explicit nil
### GetAlpnProtocols

`func (o *CreateLoadBalancerPoolRequest) GetAlpnProtocols() LoadBalancerAlpnProtocol`

GetAlpnProtocols returns the AlpnProtocols field if non-nil, zero value otherwise.

### GetAlpnProtocolsOk

`func (o *CreateLoadBalancerPoolRequest) GetAlpnProtocolsOk() (*LoadBalancerAlpnProtocol, bool)`

GetAlpnProtocolsOk returns a tuple with the AlpnProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlpnProtocols

`func (o *CreateLoadBalancerPoolRequest) SetAlpnProtocols(v LoadBalancerAlpnProtocol)`

SetAlpnProtocols sets AlpnProtocols field to given value.

### HasAlpnProtocols

`func (o *CreateLoadBalancerPoolRequest) HasAlpnProtocols() bool`

HasAlpnProtocols returns a boolean if a field has been set.

### SetAlpnProtocolsNil

`func (o *CreateLoadBalancerPoolRequest) SetAlpnProtocolsNil(b bool)`

 SetAlpnProtocolsNil sets the value for AlpnProtocols to be an explicit nil

### UnsetAlpnProtocols
`func (o *CreateLoadBalancerPoolRequest) UnsetAlpnProtocols()`

UnsetAlpnProtocols ensures that no value is present for AlpnProtocols, not even an explicit nil
### GetSessionPersistence

`func (o *CreateLoadBalancerPoolRequest) GetSessionPersistence() SessionPersistenceRequest`

GetSessionPersistence returns the SessionPersistence field if non-nil, zero value otherwise.

### GetSessionPersistenceOk

`func (o *CreateLoadBalancerPoolRequest) GetSessionPersistenceOk() (*SessionPersistenceRequest, bool)`

GetSessionPersistenceOk returns a tuple with the SessionPersistence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionPersistence

`func (o *CreateLoadBalancerPoolRequest) SetSessionPersistence(v SessionPersistenceRequest)`

SetSessionPersistence sets SessionPersistence field to given value.

### HasSessionPersistence

`func (o *CreateLoadBalancerPoolRequest) HasSessionPersistence() bool`

HasSessionPersistence returns a boolean if a field has been set.

### SetSessionPersistenceNil

`func (o *CreateLoadBalancerPoolRequest) SetSessionPersistenceNil(b bool)`

 SetSessionPersistenceNil sets the value for SessionPersistence to be an explicit nil

### UnsetSessionPersistence
`func (o *CreateLoadBalancerPoolRequest) UnsetSessionPersistence()`

UnsetSessionPersistence ensures that no value is present for SessionPersistence, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


