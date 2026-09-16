# CreatedListener

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 리스너 ID | 
**Name** | Pointer to **NullableString** | 리스너 이름 | [optional] 
**Description** | Pointer to **NullableString** | 리스너 설명 | [optional] 
**ProvisioningStatus** | Pointer to [**NullableProvisioningStatus**](ProvisioningStatus.md) | 리스너 구성 상태 | [optional] 
**OperatingStatus** | Pointer to [**NullableLoadBalancerOperatingStatus**](LoadBalancerOperatingStatus.md) | 리스너 운영 상태 | [optional] 
**Protocol** | [**Protocol**](Protocol.md) | 리스너 프로토콜 | 
**ProtocolPort** | **int32** | 리스너 포트 | 
**DefaultTlsContainerRef** | Pointer to **NullableString** | 기본 TLS 인증서 참조 ID | [optional] 
**SniContainerRefs** | **[]string** | SNI 인증서 목록 | 
**DefaultTargetGroupId** | Pointer to **NullableString** | 기본 대상 그룹 ID | [optional] 
**InsertHeaders** | **map[string]interface{}** | 요청 전달 시 삽입할 HTTP 헤더 목록 | 
**TlsVersions** | Pointer to [**[]TLSVersion**](TLSVersion.md) | 허용 TLS 버전 목록 | [optional] 
**AlpnProtocols** | Pointer to [**[]LoadBalancerAlpnProtocol**](LoadBalancerAlpnProtocol.md) | ALPN 프로토콜 목록 | [optional] 
**L7policies** | [**[]CreatedL7Policy2**](CreatedL7Policy2.md) | L7 정책 목록 | 

## Methods

### NewCreatedListener

`func NewCreatedListener(id string, protocol Protocol, protocolPort int32, sniContainerRefs []string, insertHeaders map[string]interface{}, l7policies []CreatedL7Policy2, ) *CreatedListener`

NewCreatedListener instantiates a new CreatedListener object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatedListenerWithDefaults

`func NewCreatedListenerWithDefaults() *CreatedListener`

NewCreatedListenerWithDefaults instantiates a new CreatedListener object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreatedListener) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreatedListener) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreatedListener) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *CreatedListener) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatedListener) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatedListener) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreatedListener) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreatedListener) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreatedListener) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *CreatedListener) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreatedListener) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreatedListener) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreatedListener) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreatedListener) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreatedListener) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetProvisioningStatus

`func (o *CreatedListener) GetProvisioningStatus() ProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *CreatedListener) GetProvisioningStatusOk() (*ProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *CreatedListener) SetProvisioningStatus(v ProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.

### HasProvisioningStatus

`func (o *CreatedListener) HasProvisioningStatus() bool`

HasProvisioningStatus returns a boolean if a field has been set.

### SetProvisioningStatusNil

`func (o *CreatedListener) SetProvisioningStatusNil(b bool)`

 SetProvisioningStatusNil sets the value for ProvisioningStatus to be an explicit nil

### UnsetProvisioningStatus
`func (o *CreatedListener) UnsetProvisioningStatus()`

UnsetProvisioningStatus ensures that no value is present for ProvisioningStatus, not even an explicit nil
### GetOperatingStatus

`func (o *CreatedListener) GetOperatingStatus() LoadBalancerOperatingStatus`

GetOperatingStatus returns the OperatingStatus field if non-nil, zero value otherwise.

### GetOperatingStatusOk

`func (o *CreatedListener) GetOperatingStatusOk() (*LoadBalancerOperatingStatus, bool)`

GetOperatingStatusOk returns a tuple with the OperatingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingStatus

`func (o *CreatedListener) SetOperatingStatus(v LoadBalancerOperatingStatus)`

SetOperatingStatus sets OperatingStatus field to given value.

### HasOperatingStatus

`func (o *CreatedListener) HasOperatingStatus() bool`

HasOperatingStatus returns a boolean if a field has been set.

### SetOperatingStatusNil

`func (o *CreatedListener) SetOperatingStatusNil(b bool)`

 SetOperatingStatusNil sets the value for OperatingStatus to be an explicit nil

### UnsetOperatingStatus
`func (o *CreatedListener) UnsetOperatingStatus()`

UnsetOperatingStatus ensures that no value is present for OperatingStatus, not even an explicit nil
### GetProtocol

`func (o *CreatedListener) GetProtocol() Protocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *CreatedListener) GetProtocolOk() (*Protocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *CreatedListener) SetProtocol(v Protocol)`

SetProtocol sets Protocol field to given value.


### GetProtocolPort

`func (o *CreatedListener) GetProtocolPort() int32`

GetProtocolPort returns the ProtocolPort field if non-nil, zero value otherwise.

### GetProtocolPortOk

`func (o *CreatedListener) GetProtocolPortOk() (*int32, bool)`

GetProtocolPortOk returns a tuple with the ProtocolPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolPort

`func (o *CreatedListener) SetProtocolPort(v int32)`

SetProtocolPort sets ProtocolPort field to given value.


### GetDefaultTlsContainerRef

`func (o *CreatedListener) GetDefaultTlsContainerRef() string`

GetDefaultTlsContainerRef returns the DefaultTlsContainerRef field if non-nil, zero value otherwise.

### GetDefaultTlsContainerRefOk

`func (o *CreatedListener) GetDefaultTlsContainerRefOk() (*string, bool)`

GetDefaultTlsContainerRefOk returns a tuple with the DefaultTlsContainerRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTlsContainerRef

`func (o *CreatedListener) SetDefaultTlsContainerRef(v string)`

SetDefaultTlsContainerRef sets DefaultTlsContainerRef field to given value.

### HasDefaultTlsContainerRef

`func (o *CreatedListener) HasDefaultTlsContainerRef() bool`

HasDefaultTlsContainerRef returns a boolean if a field has been set.

### SetDefaultTlsContainerRefNil

`func (o *CreatedListener) SetDefaultTlsContainerRefNil(b bool)`

 SetDefaultTlsContainerRefNil sets the value for DefaultTlsContainerRef to be an explicit nil

### UnsetDefaultTlsContainerRef
`func (o *CreatedListener) UnsetDefaultTlsContainerRef()`

UnsetDefaultTlsContainerRef ensures that no value is present for DefaultTlsContainerRef, not even an explicit nil
### GetSniContainerRefs

`func (o *CreatedListener) GetSniContainerRefs() []string`

GetSniContainerRefs returns the SniContainerRefs field if non-nil, zero value otherwise.

### GetSniContainerRefsOk

`func (o *CreatedListener) GetSniContainerRefsOk() (*[]string, bool)`

GetSniContainerRefsOk returns a tuple with the SniContainerRefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSniContainerRefs

`func (o *CreatedListener) SetSniContainerRefs(v []string)`

SetSniContainerRefs sets SniContainerRefs field to given value.


### GetDefaultTargetGroupId

`func (o *CreatedListener) GetDefaultTargetGroupId() string`

GetDefaultTargetGroupId returns the DefaultTargetGroupId field if non-nil, zero value otherwise.

### GetDefaultTargetGroupIdOk

`func (o *CreatedListener) GetDefaultTargetGroupIdOk() (*string, bool)`

GetDefaultTargetGroupIdOk returns a tuple with the DefaultTargetGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTargetGroupId

`func (o *CreatedListener) SetDefaultTargetGroupId(v string)`

SetDefaultTargetGroupId sets DefaultTargetGroupId field to given value.

### HasDefaultTargetGroupId

`func (o *CreatedListener) HasDefaultTargetGroupId() bool`

HasDefaultTargetGroupId returns a boolean if a field has been set.

### SetDefaultTargetGroupIdNil

`func (o *CreatedListener) SetDefaultTargetGroupIdNil(b bool)`

 SetDefaultTargetGroupIdNil sets the value for DefaultTargetGroupId to be an explicit nil

### UnsetDefaultTargetGroupId
`func (o *CreatedListener) UnsetDefaultTargetGroupId()`

UnsetDefaultTargetGroupId ensures that no value is present for DefaultTargetGroupId, not even an explicit nil
### GetInsertHeaders

`func (o *CreatedListener) GetInsertHeaders() map[string]interface{}`

GetInsertHeaders returns the InsertHeaders field if non-nil, zero value otherwise.

### GetInsertHeadersOk

`func (o *CreatedListener) GetInsertHeadersOk() (*map[string]interface{}, bool)`

GetInsertHeadersOk returns a tuple with the InsertHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertHeaders

`func (o *CreatedListener) SetInsertHeaders(v map[string]interface{})`

SetInsertHeaders sets InsertHeaders field to given value.


### GetTlsVersions

`func (o *CreatedListener) GetTlsVersions() []*TLSVersion`

GetTlsVersions returns the TlsVersions field if non-nil, zero value otherwise.

### GetTlsVersionsOk

`func (o *CreatedListener) GetTlsVersionsOk() (*[]*TLSVersion, bool)`

GetTlsVersionsOk returns a tuple with the TlsVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsVersions

`func (o *CreatedListener) SetTlsVersions(v []*TLSVersion)`

SetTlsVersions sets TlsVersions field to given value.

### HasTlsVersions

`func (o *CreatedListener) HasTlsVersions() bool`

HasTlsVersions returns a boolean if a field has been set.

### SetTlsVersionsNil

`func (o *CreatedListener) SetTlsVersionsNil(b bool)`

 SetTlsVersionsNil sets the value for TlsVersions to be an explicit nil

### UnsetTlsVersions
`func (o *CreatedListener) UnsetTlsVersions()`

UnsetTlsVersions ensures that no value is present for TlsVersions, not even an explicit nil
### GetAlpnProtocols

`func (o *CreatedListener) GetAlpnProtocols() []*LoadBalancerAlpnProtocol`

GetAlpnProtocols returns the AlpnProtocols field if non-nil, zero value otherwise.

### GetAlpnProtocolsOk

`func (o *CreatedListener) GetAlpnProtocolsOk() (*[]*LoadBalancerAlpnProtocol, bool)`

GetAlpnProtocolsOk returns a tuple with the AlpnProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlpnProtocols

`func (o *CreatedListener) SetAlpnProtocols(v []*LoadBalancerAlpnProtocol)`

SetAlpnProtocols sets AlpnProtocols field to given value.

### HasAlpnProtocols

`func (o *CreatedListener) HasAlpnProtocols() bool`

HasAlpnProtocols returns a boolean if a field has been set.

### SetAlpnProtocolsNil

`func (o *CreatedListener) SetAlpnProtocolsNil(b bool)`

 SetAlpnProtocolsNil sets the value for AlpnProtocols to be an explicit nil

### UnsetAlpnProtocols
`func (o *CreatedListener) UnsetAlpnProtocols()`

UnsetAlpnProtocols ensures that no value is present for AlpnProtocols, not even an explicit nil
### GetL7policies

`func (o *CreatedListener) GetL7policies() []CreatedL7Policy2`

GetL7policies returns the L7policies field if non-nil, zero value otherwise.

### GetL7policiesOk

`func (o *CreatedListener) GetL7policiesOk() (*[]CreatedL7Policy2, bool)`

GetL7policiesOk returns a tuple with the L7policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL7policies

`func (o *CreatedListener) SetL7policies(v []CreatedL7Policy2)`

SetL7policies sets L7policies field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


