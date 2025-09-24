# BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Listeners** | [**[]BnsLoadBalancerV1ApiCreateLoadBalancerModelListenerIdModel**](BnsLoadBalancerV1ApiCreateLoadBalancerModelListenerIdModel.md) | 생성된 리스너 ID 목록 | 
**TargetGroupCount** | Pointer to **NullableInt32** |  | [optional] 
**Id** | **string** | 생성된 로드 밸런서의 ID | 
**BeyondLoadBalancerId** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** | 로드 밸런서 이름 | 
**Description** | Pointer to **NullableString** |  | [optional] 
**ProvisioningStatus** | [**ProvisioningStatus**](ProvisioningStatus.md) | 프로비저닝 상태 | 
**OperatingStatus** | [**LoadBalancerOperatingStatus**](LoadBalancerOperatingStatus.md) | 운영 상태 | 
**ProjectId** | **string** | 소속 프로젝트의 ID | 
**CreatedAt** | **time.Time** | 리소스가 생성된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 
**SubnetId** | Pointer to **NullableString** |  | [optional] 
**Provider** | **string** | 로드 밸런서 서비스 제공자 | 
**FlavorId** | **string** | 로드 밸런서 유형 | 
**AvailabilityZone** | [**AvailabilityZone**](AvailabilityZone.md) | 로드 밸런서가 위치한 가용 영역 | 

## Methods

### NewBnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel

`func NewBnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel(listeners []BnsLoadBalancerV1ApiCreateLoadBalancerModelListenerIdModel, id string, name string, provisioningStatus ProvisioningStatus, operatingStatus LoadBalancerOperatingStatus, projectId string, createdAt time.Time, provider string, flavorId string, availabilityZone AvailabilityZone, ) *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel`

NewBnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel instantiates a new BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModelWithDefaults

`func NewBnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModelWithDefaults() *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel`

NewBnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModelWithDefaults instantiates a new BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetListeners

`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) GetListeners() []BnsLoadBalancerV1ApiCreateLoadBalancerModelListenerIdModel`

GetListeners returns the Listeners field if non-nil, zero value otherwise.

### GetListenersOk

`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) GetListenersOk() (*[]BnsLoadBalancerV1ApiCreateLoadBalancerModelListenerIdModel, bool)`

GetListenersOk returns a tuple with the Listeners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListeners

`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) SetListeners(v []BnsLoadBalancerV1ApiCreateLoadBalancerModelListenerIdModel)`

SetListeners sets Listeners field to given value.


### GetTargetGroupCount

`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) GetTargetGroupCount() int32`

GetTargetGroupCount returns the TargetGroupCount field if non-nil, zero value otherwise.

### GetTargetGroupCountOk

`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) GetTargetGroupCountOk() (*int32, bool)`

GetTargetGroupCountOk returns a tuple with the TargetGroupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroupCount

`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) SetTargetGroupCount(v int32)`

SetTargetGroupCount sets TargetGroupCount field to given value.

### HasTargetGroupCount

`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) HasTargetGroupCount() bool`

HasTargetGroupCount returns a boolean if a field has been set.

### SetTargetGroupCountNil

`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) SetTargetGroupCountNil(b bool)`

 SetTargetGroupCountNil sets the value for TargetGroupCount to be an explicit nil

### UnsetTargetGroupCount
`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) UnsetTargetGroupCount()`

UnsetTargetGroupCount ensures that no value is present for TargetGroupCount, not even an explicit nil
### GetId

`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) SetId(v string)`

SetId sets Id field to given value.


### GetBeyondLoadBalancerId

`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) GetBeyondLoadBalancerId() string`

GetBeyondLoadBalancerId returns the BeyondLoadBalancerId field if non-nil, zero value otherwise.

### GetBeyondLoadBalancerIdOk

`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) GetBeyondLoadBalancerIdOk() (*string, bool)`

GetBeyondLoadBalancerIdOk returns a tuple with the BeyondLoadBalancerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeyondLoadBalancerId

`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) SetBeyondLoadBalancerId(v string)`

SetBeyondLoadBalancerId sets BeyondLoadBalancerId field to given value.

### HasBeyondLoadBalancerId

`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) HasBeyondLoadBalancerId() bool`

HasBeyondLoadBalancerId returns a boolean if a field has been set.

### SetBeyondLoadBalancerIdNil

`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) SetBeyondLoadBalancerIdNil(b bool)`

 SetBeyondLoadBalancerIdNil sets the value for BeyondLoadBalancerId to be an explicit nil

### UnsetBeyondLoadBalancerId
`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) UnsetBeyondLoadBalancerId()`

UnsetBeyondLoadBalancerId ensures that no value is present for BeyondLoadBalancerId, not even an explicit nil
### GetName

`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetProvisioningStatus

`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) GetProvisioningStatus() ProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) GetProvisioningStatusOk() (*ProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) SetProvisioningStatus(v ProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.


### GetOperatingStatus

`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) GetOperatingStatus() LoadBalancerOperatingStatus`

GetOperatingStatus returns the OperatingStatus field if non-nil, zero value otherwise.

### GetOperatingStatusOk

`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) GetOperatingStatusOk() (*LoadBalancerOperatingStatus, bool)`

GetOperatingStatusOk returns a tuple with the OperatingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingStatus

`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) SetOperatingStatus(v LoadBalancerOperatingStatus)`

SetOperatingStatus sets OperatingStatus field to given value.


### GetProjectId

`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetCreatedAt

`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetSubnetId

`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.

### HasSubnetId

`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.

### SetSubnetIdNil

`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) SetSubnetIdNil(b bool)`

 SetSubnetIdNil sets the value for SubnetId to be an explicit nil

### UnsetSubnetId
`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) UnsetSubnetId()`

UnsetSubnetId ensures that no value is present for SubnetId, not even an explicit nil
### GetProvider

`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetFlavorId

`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) GetFlavorId() string`

GetFlavorId returns the FlavorId field if non-nil, zero value otherwise.

### GetFlavorIdOk

`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) GetFlavorIdOk() (*string, bool)`

GetFlavorIdOk returns a tuple with the FlavorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavorId

`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) SetFlavorId(v string)`

SetFlavorId sets FlavorId field to given value.


### GetAvailabilityZone

`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) GetAvailabilityZone() AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) GetAvailabilityZoneOk() (*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *BnsLoadBalancerV1ApiCreateLoadBalancerModelLoadBalancerModel) SetAvailabilityZone(v AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


