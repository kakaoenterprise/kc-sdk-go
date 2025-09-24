# BnsLoadBalancerV1ApiCreateListenerModelListenerModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 리스너 ID | 
**Name** | **string** | 리스너 이름 | 
**Description** | **string** | 리스너에 대한 설명 | 
**ProvisioningStatus** | [**ProvisioningStatus**](ProvisioningStatus.md) | 프로비저닝 상태 | 
**OperatingStatus** | [**LoadBalancerOperatingStatus**](LoadBalancerOperatingStatus.md) | 운영 상태 | 
**Protocol** | [**Protocol**](Protocol.md) | 수신 프로토콜 | 
**ProtocolPort** | **int32** | 수신 포트 번호 | 
**ConnectionLimit** | **int32** | 최대 동시 연결 수 | 
**DefaultTlsContainerRef** | Pointer to **NullableString** |  | [optional] 
**SniContainerRefs** | **[]string** | SNI 인증서 참조 목록 | 
**ProjectId** | **string** | 프로젝트 ID | 
**DefaultTargetGroupId** | Pointer to **NullableString** |  | [optional] 
**L7Policies** | [**[]BnsLoadBalancerV1ApiCreateListenerModelL7PolicyModel**](BnsLoadBalancerV1ApiCreateListenerModelL7PolicyModel.md) | 연결된 L7 정책 목록 | 
**InsertHeaders** | **map[string]interface{}** | 삽입할 HTTP 헤더 설정 | 
**CreatedAt** | **time.Time** | 리소스가 생성된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 
**LoadBalancers** | [**[]BnsLoadBalancerV1ApiCreateListenerModelLoadBalancerModel**](BnsLoadBalancerV1ApiCreateListenerModelLoadBalancerModel.md) | 연결된 로드 밸런서 목록 | 
**TimeoutClientData** | **int32** | 클라이언트 데이터 수신 타임아웃 (초) | 
**AllowedCidrs** | Pointer to **[]string** |  | [optional] 
**TlsCiphers** | Pointer to **NullableString** |  | [optional] 
**TlsVersions** | Pointer to **[]string** |  | [optional] 
**AlpnProtocols** | Pointer to **[]string** |  | [optional] 

## Methods

### NewBnsLoadBalancerV1ApiCreateListenerModelListenerModel

`func NewBnsLoadBalancerV1ApiCreateListenerModelListenerModel(id string, name string, description string, provisioningStatus ProvisioningStatus, operatingStatus LoadBalancerOperatingStatus, protocol Protocol, protocolPort int32, connectionLimit int32, sniContainerRefs []string, projectId string, l7Policies []BnsLoadBalancerV1ApiCreateListenerModelL7PolicyModel, insertHeaders map[string]interface{}, createdAt time.Time, loadBalancers []BnsLoadBalancerV1ApiCreateListenerModelLoadBalancerModel, timeoutClientData int32, ) *BnsLoadBalancerV1ApiCreateListenerModelListenerModel`

NewBnsLoadBalancerV1ApiCreateListenerModelListenerModel instantiates a new BnsLoadBalancerV1ApiCreateListenerModelListenerModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsLoadBalancerV1ApiCreateListenerModelListenerModelWithDefaults

`func NewBnsLoadBalancerV1ApiCreateListenerModelListenerModelWithDefaults() *BnsLoadBalancerV1ApiCreateListenerModelListenerModel`

NewBnsLoadBalancerV1ApiCreateListenerModelListenerModelWithDefaults instantiates a new BnsLoadBalancerV1ApiCreateListenerModelListenerModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetProvisioningStatus

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) GetProvisioningStatus() ProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) GetProvisioningStatusOk() (*ProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) SetProvisioningStatus(v ProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.


### GetOperatingStatus

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) GetOperatingStatus() LoadBalancerOperatingStatus`

GetOperatingStatus returns the OperatingStatus field if non-nil, zero value otherwise.

### GetOperatingStatusOk

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) GetOperatingStatusOk() (*LoadBalancerOperatingStatus, bool)`

GetOperatingStatusOk returns a tuple with the OperatingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingStatus

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) SetOperatingStatus(v LoadBalancerOperatingStatus)`

SetOperatingStatus sets OperatingStatus field to given value.


### GetProtocol

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) GetProtocol() Protocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) GetProtocolOk() (*Protocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) SetProtocol(v Protocol)`

SetProtocol sets Protocol field to given value.


### GetProtocolPort

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) GetProtocolPort() int32`

GetProtocolPort returns the ProtocolPort field if non-nil, zero value otherwise.

### GetProtocolPortOk

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) GetProtocolPortOk() (*int32, bool)`

GetProtocolPortOk returns a tuple with the ProtocolPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolPort

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) SetProtocolPort(v int32)`

SetProtocolPort sets ProtocolPort field to given value.


### GetConnectionLimit

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) GetConnectionLimit() int32`

GetConnectionLimit returns the ConnectionLimit field if non-nil, zero value otherwise.

### GetConnectionLimitOk

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) GetConnectionLimitOk() (*int32, bool)`

GetConnectionLimitOk returns a tuple with the ConnectionLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionLimit

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) SetConnectionLimit(v int32)`

SetConnectionLimit sets ConnectionLimit field to given value.


### GetDefaultTlsContainerRef

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) GetDefaultTlsContainerRef() string`

GetDefaultTlsContainerRef returns the DefaultTlsContainerRef field if non-nil, zero value otherwise.

### GetDefaultTlsContainerRefOk

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) GetDefaultTlsContainerRefOk() (*string, bool)`

GetDefaultTlsContainerRefOk returns a tuple with the DefaultTlsContainerRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTlsContainerRef

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) SetDefaultTlsContainerRef(v string)`

SetDefaultTlsContainerRef sets DefaultTlsContainerRef field to given value.

### HasDefaultTlsContainerRef

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) HasDefaultTlsContainerRef() bool`

HasDefaultTlsContainerRef returns a boolean if a field has been set.

### SetDefaultTlsContainerRefNil

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) SetDefaultTlsContainerRefNil(b bool)`

 SetDefaultTlsContainerRefNil sets the value for DefaultTlsContainerRef to be an explicit nil

### UnsetDefaultTlsContainerRef
`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) UnsetDefaultTlsContainerRef()`

UnsetDefaultTlsContainerRef ensures that no value is present for DefaultTlsContainerRef, not even an explicit nil
### GetSniContainerRefs

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) GetSniContainerRefs() []string`

GetSniContainerRefs returns the SniContainerRefs field if non-nil, zero value otherwise.

### GetSniContainerRefsOk

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) GetSniContainerRefsOk() (*[]string, bool)`

GetSniContainerRefsOk returns a tuple with the SniContainerRefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSniContainerRefs

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) SetSniContainerRefs(v []string)`

SetSniContainerRefs sets SniContainerRefs field to given value.


### GetProjectId

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetDefaultTargetGroupId

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) GetDefaultTargetGroupId() string`

GetDefaultTargetGroupId returns the DefaultTargetGroupId field if non-nil, zero value otherwise.

### GetDefaultTargetGroupIdOk

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) GetDefaultTargetGroupIdOk() (*string, bool)`

GetDefaultTargetGroupIdOk returns a tuple with the DefaultTargetGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTargetGroupId

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) SetDefaultTargetGroupId(v string)`

SetDefaultTargetGroupId sets DefaultTargetGroupId field to given value.

### HasDefaultTargetGroupId

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) HasDefaultTargetGroupId() bool`

HasDefaultTargetGroupId returns a boolean if a field has been set.

### SetDefaultTargetGroupIdNil

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) SetDefaultTargetGroupIdNil(b bool)`

 SetDefaultTargetGroupIdNil sets the value for DefaultTargetGroupId to be an explicit nil

### UnsetDefaultTargetGroupId
`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) UnsetDefaultTargetGroupId()`

UnsetDefaultTargetGroupId ensures that no value is present for DefaultTargetGroupId, not even an explicit nil
### GetL7Policies

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) GetL7Policies() []BnsLoadBalancerV1ApiCreateListenerModelL7PolicyModel`

GetL7Policies returns the L7Policies field if non-nil, zero value otherwise.

### GetL7PoliciesOk

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) GetL7PoliciesOk() (*[]BnsLoadBalancerV1ApiCreateListenerModelL7PolicyModel, bool)`

GetL7PoliciesOk returns a tuple with the L7Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL7Policies

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) SetL7Policies(v []BnsLoadBalancerV1ApiCreateListenerModelL7PolicyModel)`

SetL7Policies sets L7Policies field to given value.


### GetInsertHeaders

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) GetInsertHeaders() map[string]interface{}`

GetInsertHeaders returns the InsertHeaders field if non-nil, zero value otherwise.

### GetInsertHeadersOk

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) GetInsertHeadersOk() (*map[string]interface{}, bool)`

GetInsertHeadersOk returns a tuple with the InsertHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertHeaders

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) SetInsertHeaders(v map[string]interface{})`

SetInsertHeaders sets InsertHeaders field to given value.


### GetCreatedAt

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetLoadBalancers

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) GetLoadBalancers() []BnsLoadBalancerV1ApiCreateListenerModelLoadBalancerModel`

GetLoadBalancers returns the LoadBalancers field if non-nil, zero value otherwise.

### GetLoadBalancersOk

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) GetLoadBalancersOk() (*[]BnsLoadBalancerV1ApiCreateListenerModelLoadBalancerModel, bool)`

GetLoadBalancersOk returns a tuple with the LoadBalancers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancers

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) SetLoadBalancers(v []BnsLoadBalancerV1ApiCreateListenerModelLoadBalancerModel)`

SetLoadBalancers sets LoadBalancers field to given value.


### GetTimeoutClientData

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) GetTimeoutClientData() int32`

GetTimeoutClientData returns the TimeoutClientData field if non-nil, zero value otherwise.

### GetTimeoutClientDataOk

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) GetTimeoutClientDataOk() (*int32, bool)`

GetTimeoutClientDataOk returns a tuple with the TimeoutClientData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutClientData

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) SetTimeoutClientData(v int32)`

SetTimeoutClientData sets TimeoutClientData field to given value.


### GetAllowedCidrs

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) GetAllowedCidrs() []string`

GetAllowedCidrs returns the AllowedCidrs field if non-nil, zero value otherwise.

### GetAllowedCidrsOk

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) GetAllowedCidrsOk() (*[]string, bool)`

GetAllowedCidrsOk returns a tuple with the AllowedCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedCidrs

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) SetAllowedCidrs(v []string)`

SetAllowedCidrs sets AllowedCidrs field to given value.

### HasAllowedCidrs

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) HasAllowedCidrs() bool`

HasAllowedCidrs returns a boolean if a field has been set.

### SetAllowedCidrsNil

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) SetAllowedCidrsNil(b bool)`

 SetAllowedCidrsNil sets the value for AllowedCidrs to be an explicit nil

### UnsetAllowedCidrs
`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) UnsetAllowedCidrs()`

UnsetAllowedCidrs ensures that no value is present for AllowedCidrs, not even an explicit nil
### GetTlsCiphers

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) GetTlsCiphers() string`

GetTlsCiphers returns the TlsCiphers field if non-nil, zero value otherwise.

### GetTlsCiphersOk

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) GetTlsCiphersOk() (*string, bool)`

GetTlsCiphersOk returns a tuple with the TlsCiphers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCiphers

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) SetTlsCiphers(v string)`

SetTlsCiphers sets TlsCiphers field to given value.

### HasTlsCiphers

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) HasTlsCiphers() bool`

HasTlsCiphers returns a boolean if a field has been set.

### SetTlsCiphersNil

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) SetTlsCiphersNil(b bool)`

 SetTlsCiphersNil sets the value for TlsCiphers to be an explicit nil

### UnsetTlsCiphers
`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) UnsetTlsCiphers()`

UnsetTlsCiphers ensures that no value is present for TlsCiphers, not even an explicit nil
### GetTlsVersions

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) GetTlsVersions() []string`

GetTlsVersions returns the TlsVersions field if non-nil, zero value otherwise.

### GetTlsVersionsOk

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) GetTlsVersionsOk() (*[]string, bool)`

GetTlsVersionsOk returns a tuple with the TlsVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsVersions

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) SetTlsVersions(v []string)`

SetTlsVersions sets TlsVersions field to given value.

### HasTlsVersions

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) HasTlsVersions() bool`

HasTlsVersions returns a boolean if a field has been set.

### SetTlsVersionsNil

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) SetTlsVersionsNil(b bool)`

 SetTlsVersionsNil sets the value for TlsVersions to be an explicit nil

### UnsetTlsVersions
`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) UnsetTlsVersions()`

UnsetTlsVersions ensures that no value is present for TlsVersions, not even an explicit nil
### GetAlpnProtocols

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) GetAlpnProtocols() []string`

GetAlpnProtocols returns the AlpnProtocols field if non-nil, zero value otherwise.

### GetAlpnProtocolsOk

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) GetAlpnProtocolsOk() (*[]string, bool)`

GetAlpnProtocolsOk returns a tuple with the AlpnProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlpnProtocols

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) SetAlpnProtocols(v []string)`

SetAlpnProtocols sets AlpnProtocols field to given value.

### HasAlpnProtocols

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) HasAlpnProtocols() bool`

HasAlpnProtocols returns a boolean if a field has been set.

### SetAlpnProtocolsNil

`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) SetAlpnProtocolsNil(b bool)`

 SetAlpnProtocolsNil sets the value for AlpnProtocols to be an explicit nil

### UnsetAlpnProtocols
`func (o *BnsLoadBalancerV1ApiCreateListenerModelListenerModel) UnsetAlpnProtocols()`

UnsetAlpnProtocols ensures that no value is present for AlpnProtocols, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


