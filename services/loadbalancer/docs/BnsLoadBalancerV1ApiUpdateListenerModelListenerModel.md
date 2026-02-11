# BnsLoadBalancerV1ApiUpdateListenerModelListenerModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 리스너 고유 ID | 
**Name** | **string** | 리스너 이름 | 
**Description** | **string** | 리스너에 대한 설명 | 
**ProvisioningStatus** | [**ProvisioningStatus**](ProvisioningStatus.md) | 프로비저닝 상태 | 
**OperatingStatus** | [**LoadBalancerOperatingStatus**](LoadBalancerOperatingStatus.md) | 운영 상태 | 
**Protocol** | [**Protocol**](Protocol.md) | 리스너가 사용하는 프로토콜 | 
**ProtocolPort** | **int32** | 리스너가 사용하는 포트 | 
**DefaultTlsContainerRef** | Pointer to **NullableString** |  | [optional] 
**SniContainerRefs** | **[]string** | SNI(Server Name Indication)를 위한 추가 인증서 목록 | 
**ProjectId** | **string** | 리스너가 속한 프로젝트의 ID | 
**DefaultTargetGroupId** | Pointer to **NullableString** |  | [optional] 
**L7Policies** | [**[]BnsLoadBalancerV1ApiUpdateListenerModelL7PolicyModel**](BnsLoadBalancerV1ApiUpdateListenerModelL7PolicyModel.md) | 연결된 L7 정책 목록 | 
**InsertHeaders** | **map[string]interface{}** | 백엔드로 전달할 헤더 (예: X-Forwarded-For, X-Forwarded-Port) | 
**CreatedAt** | **time.Time** | 리소스가 생성된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 
**LoadBalancers** | [**[]BnsLoadBalancerV1ApiUpdateListenerModelLoadBalancerModel**](BnsLoadBalancerV1ApiUpdateListenerModelLoadBalancerModel.md) | 리스너가 연결된 로드 밸런서 목록 | 
**TimeoutClientData** | **int32** | 클라이언트에서 수신 데이터 대기 시간 (ms) | 
**TlsCiphers** | Pointer to **NullableString** |  | [optional] 
**TlsVersions** | Pointer to **[]string** |  | [optional] 

## Methods

### NewBnsLoadBalancerV1ApiUpdateListenerModelListenerModel

`func NewBnsLoadBalancerV1ApiUpdateListenerModelListenerModel(id string, name string, description string, provisioningStatus ProvisioningStatus, operatingStatus LoadBalancerOperatingStatus, protocol Protocol, protocolPort int32, sniContainerRefs []string, projectId string, l7Policies []BnsLoadBalancerV1ApiUpdateListenerModelL7PolicyModel, insertHeaders map[string]interface{}, createdAt time.Time, loadBalancers []BnsLoadBalancerV1ApiUpdateListenerModelLoadBalancerModel, timeoutClientData int32, ) *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel`

NewBnsLoadBalancerV1ApiUpdateListenerModelListenerModel instantiates a new BnsLoadBalancerV1ApiUpdateListenerModelListenerModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsLoadBalancerV1ApiUpdateListenerModelListenerModelWithDefaults

`func NewBnsLoadBalancerV1ApiUpdateListenerModelListenerModelWithDefaults() *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel`

NewBnsLoadBalancerV1ApiUpdateListenerModelListenerModelWithDefaults instantiates a new BnsLoadBalancerV1ApiUpdateListenerModelListenerModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetProvisioningStatus

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) GetProvisioningStatus() ProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) GetProvisioningStatusOk() (*ProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) SetProvisioningStatus(v ProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.


### GetOperatingStatus

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) GetOperatingStatus() LoadBalancerOperatingStatus`

GetOperatingStatus returns the OperatingStatus field if non-nil, zero value otherwise.

### GetOperatingStatusOk

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) GetOperatingStatusOk() (*LoadBalancerOperatingStatus, bool)`

GetOperatingStatusOk returns a tuple with the OperatingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingStatus

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) SetOperatingStatus(v LoadBalancerOperatingStatus)`

SetOperatingStatus sets OperatingStatus field to given value.


### GetProtocol

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) GetProtocol() Protocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) GetProtocolOk() (*Protocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) SetProtocol(v Protocol)`

SetProtocol sets Protocol field to given value.


### GetProtocolPort

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) GetProtocolPort() int32`

GetProtocolPort returns the ProtocolPort field if non-nil, zero value otherwise.

### GetProtocolPortOk

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) GetProtocolPortOk() (*int32, bool)`

GetProtocolPortOk returns a tuple with the ProtocolPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolPort

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) SetProtocolPort(v int32)`

SetProtocolPort sets ProtocolPort field to given value.


### GetDefaultTlsContainerRef

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) GetDefaultTlsContainerRef() string`

GetDefaultTlsContainerRef returns the DefaultTlsContainerRef field if non-nil, zero value otherwise.

### GetDefaultTlsContainerRefOk

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) GetDefaultTlsContainerRefOk() (*string, bool)`

GetDefaultTlsContainerRefOk returns a tuple with the DefaultTlsContainerRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTlsContainerRef

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) SetDefaultTlsContainerRef(v string)`

SetDefaultTlsContainerRef sets DefaultTlsContainerRef field to given value.

### HasDefaultTlsContainerRef

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) HasDefaultTlsContainerRef() bool`

HasDefaultTlsContainerRef returns a boolean if a field has been set.

### SetDefaultTlsContainerRefNil

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) SetDefaultTlsContainerRefNil(b bool)`

 SetDefaultTlsContainerRefNil sets the value for DefaultTlsContainerRef to be an explicit nil

### UnsetDefaultTlsContainerRef
`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) UnsetDefaultTlsContainerRef()`

UnsetDefaultTlsContainerRef ensures that no value is present for DefaultTlsContainerRef, not even an explicit nil
### GetSniContainerRefs

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) GetSniContainerRefs() []string`

GetSniContainerRefs returns the SniContainerRefs field if non-nil, zero value otherwise.

### GetSniContainerRefsOk

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) GetSniContainerRefsOk() (*[]string, bool)`

GetSniContainerRefsOk returns a tuple with the SniContainerRefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSniContainerRefs

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) SetSniContainerRefs(v []string)`

SetSniContainerRefs sets SniContainerRefs field to given value.


### GetProjectId

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetDefaultTargetGroupId

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) GetDefaultTargetGroupId() string`

GetDefaultTargetGroupId returns the DefaultTargetGroupId field if non-nil, zero value otherwise.

### GetDefaultTargetGroupIdOk

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) GetDefaultTargetGroupIdOk() (*string, bool)`

GetDefaultTargetGroupIdOk returns a tuple with the DefaultTargetGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTargetGroupId

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) SetDefaultTargetGroupId(v string)`

SetDefaultTargetGroupId sets DefaultTargetGroupId field to given value.

### HasDefaultTargetGroupId

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) HasDefaultTargetGroupId() bool`

HasDefaultTargetGroupId returns a boolean if a field has been set.

### SetDefaultTargetGroupIdNil

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) SetDefaultTargetGroupIdNil(b bool)`

 SetDefaultTargetGroupIdNil sets the value for DefaultTargetGroupId to be an explicit nil

### UnsetDefaultTargetGroupId
`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) UnsetDefaultTargetGroupId()`

UnsetDefaultTargetGroupId ensures that no value is present for DefaultTargetGroupId, not even an explicit nil
### GetL7Policies

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) GetL7Policies() []BnsLoadBalancerV1ApiUpdateListenerModelL7PolicyModel`

GetL7Policies returns the L7Policies field if non-nil, zero value otherwise.

### GetL7PoliciesOk

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) GetL7PoliciesOk() (*[]BnsLoadBalancerV1ApiUpdateListenerModelL7PolicyModel, bool)`

GetL7PoliciesOk returns a tuple with the L7Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL7Policies

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) SetL7Policies(v []BnsLoadBalancerV1ApiUpdateListenerModelL7PolicyModel)`

SetL7Policies sets L7Policies field to given value.


### GetInsertHeaders

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) GetInsertHeaders() map[string]interface{}`

GetInsertHeaders returns the InsertHeaders field if non-nil, zero value otherwise.

### GetInsertHeadersOk

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) GetInsertHeadersOk() (*map[string]interface{}, bool)`

GetInsertHeadersOk returns a tuple with the InsertHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertHeaders

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) SetInsertHeaders(v map[string]interface{})`

SetInsertHeaders sets InsertHeaders field to given value.


### GetCreatedAt

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetLoadBalancers

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) GetLoadBalancers() []BnsLoadBalancerV1ApiUpdateListenerModelLoadBalancerModel`

GetLoadBalancers returns the LoadBalancers field if non-nil, zero value otherwise.

### GetLoadBalancersOk

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) GetLoadBalancersOk() (*[]BnsLoadBalancerV1ApiUpdateListenerModelLoadBalancerModel, bool)`

GetLoadBalancersOk returns a tuple with the LoadBalancers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancers

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) SetLoadBalancers(v []BnsLoadBalancerV1ApiUpdateListenerModelLoadBalancerModel)`

SetLoadBalancers sets LoadBalancers field to given value.


### GetTimeoutClientData

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) GetTimeoutClientData() int32`

GetTimeoutClientData returns the TimeoutClientData field if non-nil, zero value otherwise.

### GetTimeoutClientDataOk

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) GetTimeoutClientDataOk() (*int32, bool)`

GetTimeoutClientDataOk returns a tuple with the TimeoutClientData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutClientData

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) SetTimeoutClientData(v int32)`

SetTimeoutClientData sets TimeoutClientData field to given value.


### GetTlsCiphers

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) GetTlsCiphers() string`

GetTlsCiphers returns the TlsCiphers field if non-nil, zero value otherwise.

### GetTlsCiphersOk

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) GetTlsCiphersOk() (*string, bool)`

GetTlsCiphersOk returns a tuple with the TlsCiphers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCiphers

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) SetTlsCiphers(v string)`

SetTlsCiphers sets TlsCiphers field to given value.

### HasTlsCiphers

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) HasTlsCiphers() bool`

HasTlsCiphers returns a boolean if a field has been set.

### SetTlsCiphersNil

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) SetTlsCiphersNil(b bool)`

 SetTlsCiphersNil sets the value for TlsCiphers to be an explicit nil

### UnsetTlsCiphers
`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) UnsetTlsCiphers()`

UnsetTlsCiphers ensures that no value is present for TlsCiphers, not even an explicit nil
### GetTlsVersions

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) GetTlsVersions() []string`

GetTlsVersions returns the TlsVersions field if non-nil, zero value otherwise.

### GetTlsVersionsOk

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) GetTlsVersionsOk() (*[]string, bool)`

GetTlsVersionsOk returns a tuple with the TlsVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsVersions

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) SetTlsVersions(v []string)`

SetTlsVersions sets TlsVersions field to given value.

### HasTlsVersions

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) HasTlsVersions() bool`

HasTlsVersions returns a boolean if a field has been set.

### SetTlsVersionsNil

`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) SetTlsVersionsNil(b bool)`

 SetTlsVersionsNil sets the value for TlsVersions to be an explicit nil

### UnsetTlsVersions
`func (o *BnsLoadBalancerV1ApiUpdateListenerModelListenerModel) UnsetTlsVersions()`

UnsetTlsVersions ensures that no value is present for TlsVersions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


