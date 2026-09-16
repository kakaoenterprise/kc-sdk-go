# CreatedLoadBalancer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Listeners** | [**[]CreatedListener**](CreatedListener.md) | 생성된 리스너 ID 목록 | 
**TargetGroups** | [**[]CreatedPool**](CreatedPool.md) | 대상 그룹 목록 | 
**TargetGroupCount** | Pointer to **NullableInt32** | 로드 밸런서에 연결된 대상 그룹 수 | [optional] 
**Id** | **string** | 생성된 로드 밸런서의 ID | 
**BeyondLoadBalancerId** | Pointer to **NullableString** | 연결된 고가용성 그룹의 ID (있을 경우) | [optional] 
**Name** | **string** | 로드 밸런서 이름 | 
**Description** | Pointer to **NullableString** | 로드 밸런서에 대한 설명 | [optional] 
**ProvisioningStatus** | [**ProvisioningStatus**](ProvisioningStatus.md) | 프로비저닝 상태 | 
**OperatingStatus** | [**LoadBalancerOperatingStatus**](LoadBalancerOperatingStatus.md) | 운영 상태 | 
**ProjectId** | **string** | 소속 프로젝트의 ID | 
**CreatedAt** | **time.Time** | 리소스가 생성된 시간 - ISO 8601 형식 - UTC 기준 | 
**UpdatedAt** | Pointer to **NullableTime** | 리소스가 마지막으로 수정된 시간 - ISO 8601 형식 - UTC 기준 | [optional] 
**SubnetId** | **string** | 로드 밸런서가 속한 서브넷 ID | 
**Provider** | **string** | 로드 밸런서 서비스 제공자 | 
**FlavorId** | **string** | 로드 밸런서 유형 | 
**AvailabilityZone** | [**AvailabilityZone**](AvailabilityZone.md) | 로드 밸런서가 위치한 가용 영역 | 
**ListenerCount** | Pointer to **NullableInt32** | 리스너 개수 | [optional] 

## Methods

### NewCreatedLoadBalancer

`func NewCreatedLoadBalancer(listeners []CreatedListener, targetGroups []CreatedPool, id string, name string, provisioningStatus ProvisioningStatus, operatingStatus LoadBalancerOperatingStatus, projectId string, createdAt time.Time, subnetId string, provider string, flavorId string, availabilityZone AvailabilityZone, ) *CreatedLoadBalancer`

NewCreatedLoadBalancer instantiates a new CreatedLoadBalancer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatedLoadBalancerWithDefaults

`func NewCreatedLoadBalancerWithDefaults() *CreatedLoadBalancer`

NewCreatedLoadBalancerWithDefaults instantiates a new CreatedLoadBalancer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetListeners

`func (o *CreatedLoadBalancer) GetListeners() []CreatedListener`

GetListeners returns the Listeners field if non-nil, zero value otherwise.

### GetListenersOk

`func (o *CreatedLoadBalancer) GetListenersOk() (*[]CreatedListener, bool)`

GetListenersOk returns a tuple with the Listeners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListeners

`func (o *CreatedLoadBalancer) SetListeners(v []CreatedListener)`

SetListeners sets Listeners field to given value.


### GetTargetGroups

`func (o *CreatedLoadBalancer) GetTargetGroups() []CreatedPool`

GetTargetGroups returns the TargetGroups field if non-nil, zero value otherwise.

### GetTargetGroupsOk

`func (o *CreatedLoadBalancer) GetTargetGroupsOk() (*[]CreatedPool, bool)`

GetTargetGroupsOk returns a tuple with the TargetGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroups

`func (o *CreatedLoadBalancer) SetTargetGroups(v []CreatedPool)`

SetTargetGroups sets TargetGroups field to given value.


### GetTargetGroupCount

`func (o *CreatedLoadBalancer) GetTargetGroupCount() int32`

GetTargetGroupCount returns the TargetGroupCount field if non-nil, zero value otherwise.

### GetTargetGroupCountOk

`func (o *CreatedLoadBalancer) GetTargetGroupCountOk() (*int32, bool)`

GetTargetGroupCountOk returns a tuple with the TargetGroupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroupCount

`func (o *CreatedLoadBalancer) SetTargetGroupCount(v int32)`

SetTargetGroupCount sets TargetGroupCount field to given value.

### HasTargetGroupCount

`func (o *CreatedLoadBalancer) HasTargetGroupCount() bool`

HasTargetGroupCount returns a boolean if a field has been set.

### SetTargetGroupCountNil

`func (o *CreatedLoadBalancer) SetTargetGroupCountNil(b bool)`

 SetTargetGroupCountNil sets the value for TargetGroupCount to be an explicit nil

### UnsetTargetGroupCount
`func (o *CreatedLoadBalancer) UnsetTargetGroupCount()`

UnsetTargetGroupCount ensures that no value is present for TargetGroupCount, not even an explicit nil
### GetId

`func (o *CreatedLoadBalancer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreatedLoadBalancer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreatedLoadBalancer) SetId(v string)`

SetId sets Id field to given value.


### GetBeyondLoadBalancerId

`func (o *CreatedLoadBalancer) GetBeyondLoadBalancerId() string`

GetBeyondLoadBalancerId returns the BeyondLoadBalancerId field if non-nil, zero value otherwise.

### GetBeyondLoadBalancerIdOk

`func (o *CreatedLoadBalancer) GetBeyondLoadBalancerIdOk() (*string, bool)`

GetBeyondLoadBalancerIdOk returns a tuple with the BeyondLoadBalancerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeyondLoadBalancerId

`func (o *CreatedLoadBalancer) SetBeyondLoadBalancerId(v string)`

SetBeyondLoadBalancerId sets BeyondLoadBalancerId field to given value.

### HasBeyondLoadBalancerId

`func (o *CreatedLoadBalancer) HasBeyondLoadBalancerId() bool`

HasBeyondLoadBalancerId returns a boolean if a field has been set.

### SetBeyondLoadBalancerIdNil

`func (o *CreatedLoadBalancer) SetBeyondLoadBalancerIdNil(b bool)`

 SetBeyondLoadBalancerIdNil sets the value for BeyondLoadBalancerId to be an explicit nil

### UnsetBeyondLoadBalancerId
`func (o *CreatedLoadBalancer) UnsetBeyondLoadBalancerId()`

UnsetBeyondLoadBalancerId ensures that no value is present for BeyondLoadBalancerId, not even an explicit nil
### GetName

`func (o *CreatedLoadBalancer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatedLoadBalancer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatedLoadBalancer) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreatedLoadBalancer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreatedLoadBalancer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreatedLoadBalancer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreatedLoadBalancer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreatedLoadBalancer) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreatedLoadBalancer) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetProvisioningStatus

`func (o *CreatedLoadBalancer) GetProvisioningStatus() ProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *CreatedLoadBalancer) GetProvisioningStatusOk() (*ProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *CreatedLoadBalancer) SetProvisioningStatus(v ProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.


### GetOperatingStatus

`func (o *CreatedLoadBalancer) GetOperatingStatus() LoadBalancerOperatingStatus`

GetOperatingStatus returns the OperatingStatus field if non-nil, zero value otherwise.

### GetOperatingStatusOk

`func (o *CreatedLoadBalancer) GetOperatingStatusOk() (*LoadBalancerOperatingStatus, bool)`

GetOperatingStatusOk returns a tuple with the OperatingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingStatus

`func (o *CreatedLoadBalancer) SetOperatingStatus(v LoadBalancerOperatingStatus)`

SetOperatingStatus sets OperatingStatus field to given value.


### GetProjectId

`func (o *CreatedLoadBalancer) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CreatedLoadBalancer) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CreatedLoadBalancer) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetCreatedAt

`func (o *CreatedLoadBalancer) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreatedLoadBalancer) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreatedLoadBalancer) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *CreatedLoadBalancer) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CreatedLoadBalancer) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CreatedLoadBalancer) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CreatedLoadBalancer) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *CreatedLoadBalancer) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *CreatedLoadBalancer) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetSubnetId

`func (o *CreatedLoadBalancer) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *CreatedLoadBalancer) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *CreatedLoadBalancer) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.


### GetProvider

`func (o *CreatedLoadBalancer) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CreatedLoadBalancer) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CreatedLoadBalancer) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetFlavorId

`func (o *CreatedLoadBalancer) GetFlavorId() string`

GetFlavorId returns the FlavorId field if non-nil, zero value otherwise.

### GetFlavorIdOk

`func (o *CreatedLoadBalancer) GetFlavorIdOk() (*string, bool)`

GetFlavorIdOk returns a tuple with the FlavorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavorId

`func (o *CreatedLoadBalancer) SetFlavorId(v string)`

SetFlavorId sets FlavorId field to given value.


### GetAvailabilityZone

`func (o *CreatedLoadBalancer) GetAvailabilityZone() AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *CreatedLoadBalancer) GetAvailabilityZoneOk() (*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *CreatedLoadBalancer) SetAvailabilityZone(v AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.


### GetListenerCount

`func (o *CreatedLoadBalancer) GetListenerCount() int32`

GetListenerCount returns the ListenerCount field if non-nil, zero value otherwise.

### GetListenerCountOk

`func (o *CreatedLoadBalancer) GetListenerCountOk() (*int32, bool)`

GetListenerCountOk returns a tuple with the ListenerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenerCount

`func (o *CreatedLoadBalancer) SetListenerCount(v int32)`

SetListenerCount sets ListenerCount field to given value.

### HasListenerCount

`func (o *CreatedLoadBalancer) HasListenerCount() bool`

HasListenerCount returns a boolean if a field has been set.

### SetListenerCountNil

`func (o *CreatedLoadBalancer) SetListenerCountNil(b bool)`

 SetListenerCountNil sets the value for ListenerCount to be an explicit nil

### UnsetListenerCount
`func (o *CreatedLoadBalancer) UnsetListenerCount()`

UnsetListenerCount ensures that no value is present for ListenerCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


