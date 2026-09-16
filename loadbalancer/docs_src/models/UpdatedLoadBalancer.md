# UpdatedLoadBalancer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 로드 밸런서의 ID | 
**Name** | **string** | 로드 밸런서의 이름 | 
**Description** | Pointer to **NullableString** | 로드 밸런서에 대한 설명 | [optional] 
**ProvisioningStatus** | Pointer to [**NullableProvisioningStatus**](ProvisioningStatus.md) | 프로비저닝 상태 | [optional] 
**OperatingStatus** | Pointer to [**NullableLoadBalancerOperatingStatus**](LoadBalancerOperatingStatus.md) | 운영 상태 | [optional] 
**AccessLogs** | Pointer to [**NullableAccessLog**](AccessLog.md) | 로드 밸런서 액세스 로그 설정 정보 | [optional] 
**ProjectId** | **string** | 해당 리소스가 속한 프로젝트의 ID | 
**CreatedAt** | Pointer to **NullableTime** | 리소스가 생성된 시간 - ISO 8601 형식 - UTC 기준 | [optional] 
**UpdatedAt** | Pointer to **NullableTime** | 리소스가 마지막으로 수정된 시간 - ISO 8601 형식 - UTC 기준 | [optional] 
**SubnetId** | Pointer to **NullableString** | 연결된 서브넷 ID | [optional] 
**Listeners** | Pointer to [**[]UpdatedListener**](UpdatedListener.md) | 연결된 리스너 객체 목록 | [optional] 
**TargetGroupCount** | Pointer to **NullableInt32** | 연결된 대상 그룹 수 | [optional] 
**Provider** | Pointer to **NullableString** | 로드 밸런서를 제공하는 백엔드 공급자 | [optional] 
**FlavorId** | Pointer to **NullableString** | 로드 밸런서의 인스턴스 유형 ID | [optional] 
**AvailabilityZone** | Pointer to [**NullableAvailabilityZone**](AvailabilityZone.md) | 리소스가 위치한 가용 영역 | [optional] 

## Methods

### NewUpdatedLoadBalancer

`func NewUpdatedLoadBalancer(id string, name string, projectId string, ) *UpdatedLoadBalancer`

NewUpdatedLoadBalancer instantiates a new UpdatedLoadBalancer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatedLoadBalancerWithDefaults

`func NewUpdatedLoadBalancerWithDefaults() *UpdatedLoadBalancer`

NewUpdatedLoadBalancerWithDefaults instantiates a new UpdatedLoadBalancer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdatedLoadBalancer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdatedLoadBalancer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdatedLoadBalancer) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *UpdatedLoadBalancer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdatedLoadBalancer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdatedLoadBalancer) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *UpdatedLoadBalancer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdatedLoadBalancer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdatedLoadBalancer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdatedLoadBalancer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdatedLoadBalancer) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdatedLoadBalancer) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetProvisioningStatus

`func (o *UpdatedLoadBalancer) GetProvisioningStatus() ProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *UpdatedLoadBalancer) GetProvisioningStatusOk() (*ProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *UpdatedLoadBalancer) SetProvisioningStatus(v ProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.

### HasProvisioningStatus

`func (o *UpdatedLoadBalancer) HasProvisioningStatus() bool`

HasProvisioningStatus returns a boolean if a field has been set.

### SetProvisioningStatusNil

`func (o *UpdatedLoadBalancer) SetProvisioningStatusNil(b bool)`

 SetProvisioningStatusNil sets the value for ProvisioningStatus to be an explicit nil

### UnsetProvisioningStatus
`func (o *UpdatedLoadBalancer) UnsetProvisioningStatus()`

UnsetProvisioningStatus ensures that no value is present for ProvisioningStatus, not even an explicit nil
### GetOperatingStatus

`func (o *UpdatedLoadBalancer) GetOperatingStatus() LoadBalancerOperatingStatus`

GetOperatingStatus returns the OperatingStatus field if non-nil, zero value otherwise.

### GetOperatingStatusOk

`func (o *UpdatedLoadBalancer) GetOperatingStatusOk() (*LoadBalancerOperatingStatus, bool)`

GetOperatingStatusOk returns a tuple with the OperatingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingStatus

`func (o *UpdatedLoadBalancer) SetOperatingStatus(v LoadBalancerOperatingStatus)`

SetOperatingStatus sets OperatingStatus field to given value.

### HasOperatingStatus

`func (o *UpdatedLoadBalancer) HasOperatingStatus() bool`

HasOperatingStatus returns a boolean if a field has been set.

### SetOperatingStatusNil

`func (o *UpdatedLoadBalancer) SetOperatingStatusNil(b bool)`

 SetOperatingStatusNil sets the value for OperatingStatus to be an explicit nil

### UnsetOperatingStatus
`func (o *UpdatedLoadBalancer) UnsetOperatingStatus()`

UnsetOperatingStatus ensures that no value is present for OperatingStatus, not even an explicit nil
### GetAccessLogs

`func (o *UpdatedLoadBalancer) GetAccessLogs() AccessLog`

GetAccessLogs returns the AccessLogs field if non-nil, zero value otherwise.

### GetAccessLogsOk

`func (o *UpdatedLoadBalancer) GetAccessLogsOk() (*AccessLog, bool)`

GetAccessLogsOk returns a tuple with the AccessLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLogs

`func (o *UpdatedLoadBalancer) SetAccessLogs(v AccessLog)`

SetAccessLogs sets AccessLogs field to given value.

### HasAccessLogs

`func (o *UpdatedLoadBalancer) HasAccessLogs() bool`

HasAccessLogs returns a boolean if a field has been set.

### SetAccessLogsNil

`func (o *UpdatedLoadBalancer) SetAccessLogsNil(b bool)`

 SetAccessLogsNil sets the value for AccessLogs to be an explicit nil

### UnsetAccessLogs
`func (o *UpdatedLoadBalancer) UnsetAccessLogs()`

UnsetAccessLogs ensures that no value is present for AccessLogs, not even an explicit nil
### GetProjectId

`func (o *UpdatedLoadBalancer) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *UpdatedLoadBalancer) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *UpdatedLoadBalancer) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetCreatedAt

`func (o *UpdatedLoadBalancer) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UpdatedLoadBalancer) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UpdatedLoadBalancer) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *UpdatedLoadBalancer) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *UpdatedLoadBalancer) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *UpdatedLoadBalancer) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *UpdatedLoadBalancer) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UpdatedLoadBalancer) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UpdatedLoadBalancer) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *UpdatedLoadBalancer) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *UpdatedLoadBalancer) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *UpdatedLoadBalancer) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetSubnetId

`func (o *UpdatedLoadBalancer) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *UpdatedLoadBalancer) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *UpdatedLoadBalancer) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.

### HasSubnetId

`func (o *UpdatedLoadBalancer) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.

### SetSubnetIdNil

`func (o *UpdatedLoadBalancer) SetSubnetIdNil(b bool)`

 SetSubnetIdNil sets the value for SubnetId to be an explicit nil

### UnsetSubnetId
`func (o *UpdatedLoadBalancer) UnsetSubnetId()`

UnsetSubnetId ensures that no value is present for SubnetId, not even an explicit nil
### GetListeners

`func (o *UpdatedLoadBalancer) GetListeners() []UpdatedListener`

GetListeners returns the Listeners field if non-nil, zero value otherwise.

### GetListenersOk

`func (o *UpdatedLoadBalancer) GetListenersOk() (*[]UpdatedListener, bool)`

GetListenersOk returns a tuple with the Listeners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListeners

`func (o *UpdatedLoadBalancer) SetListeners(v []UpdatedListener)`

SetListeners sets Listeners field to given value.

### HasListeners

`func (o *UpdatedLoadBalancer) HasListeners() bool`

HasListeners returns a boolean if a field has been set.

### SetListenersNil

`func (o *UpdatedLoadBalancer) SetListenersNil(b bool)`

 SetListenersNil sets the value for Listeners to be an explicit nil

### UnsetListeners
`func (o *UpdatedLoadBalancer) UnsetListeners()`

UnsetListeners ensures that no value is present for Listeners, not even an explicit nil
### GetTargetGroupCount

`func (o *UpdatedLoadBalancer) GetTargetGroupCount() int32`

GetTargetGroupCount returns the TargetGroupCount field if non-nil, zero value otherwise.

### GetTargetGroupCountOk

`func (o *UpdatedLoadBalancer) GetTargetGroupCountOk() (*int32, bool)`

GetTargetGroupCountOk returns a tuple with the TargetGroupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroupCount

`func (o *UpdatedLoadBalancer) SetTargetGroupCount(v int32)`

SetTargetGroupCount sets TargetGroupCount field to given value.

### HasTargetGroupCount

`func (o *UpdatedLoadBalancer) HasTargetGroupCount() bool`

HasTargetGroupCount returns a boolean if a field has been set.

### SetTargetGroupCountNil

`func (o *UpdatedLoadBalancer) SetTargetGroupCountNil(b bool)`

 SetTargetGroupCountNil sets the value for TargetGroupCount to be an explicit nil

### UnsetTargetGroupCount
`func (o *UpdatedLoadBalancer) UnsetTargetGroupCount()`

UnsetTargetGroupCount ensures that no value is present for TargetGroupCount, not even an explicit nil
### GetProvider

`func (o *UpdatedLoadBalancer) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *UpdatedLoadBalancer) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *UpdatedLoadBalancer) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *UpdatedLoadBalancer) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### SetProviderNil

`func (o *UpdatedLoadBalancer) SetProviderNil(b bool)`

 SetProviderNil sets the value for Provider to be an explicit nil

### UnsetProvider
`func (o *UpdatedLoadBalancer) UnsetProvider()`

UnsetProvider ensures that no value is present for Provider, not even an explicit nil
### GetFlavorId

`func (o *UpdatedLoadBalancer) GetFlavorId() string`

GetFlavorId returns the FlavorId field if non-nil, zero value otherwise.

### GetFlavorIdOk

`func (o *UpdatedLoadBalancer) GetFlavorIdOk() (*string, bool)`

GetFlavorIdOk returns a tuple with the FlavorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavorId

`func (o *UpdatedLoadBalancer) SetFlavorId(v string)`

SetFlavorId sets FlavorId field to given value.

### HasFlavorId

`func (o *UpdatedLoadBalancer) HasFlavorId() bool`

HasFlavorId returns a boolean if a field has been set.

### SetFlavorIdNil

`func (o *UpdatedLoadBalancer) SetFlavorIdNil(b bool)`

 SetFlavorIdNil sets the value for FlavorId to be an explicit nil

### UnsetFlavorId
`func (o *UpdatedLoadBalancer) UnsetFlavorId()`

UnsetFlavorId ensures that no value is present for FlavorId, not even an explicit nil
### GetAvailabilityZone

`func (o *UpdatedLoadBalancer) GetAvailabilityZone() AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *UpdatedLoadBalancer) GetAvailabilityZoneOk() (*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *UpdatedLoadBalancer) SetAvailabilityZone(v AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *UpdatedLoadBalancer) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### SetAvailabilityZoneNil

`func (o *UpdatedLoadBalancer) SetAvailabilityZoneNil(b bool)`

 SetAvailabilityZoneNil sets the value for AvailabilityZone to be an explicit nil

### UnsetAvailabilityZone
`func (o *UpdatedLoadBalancer) UnsetAvailabilityZone()`

UnsetAvailabilityZone ensures that no value is present for AvailabilityZone, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


