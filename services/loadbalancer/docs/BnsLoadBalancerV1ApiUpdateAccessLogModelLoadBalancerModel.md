# BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 로드 밸런서의 ID | 
**Name** | **string** | 로드 밸런서의 이름 | 
**Description** | Pointer to **NullableString** |  | [optional] 
**ProvisioningStatus** | Pointer to [**NullableProvisioningStatus**](ProvisioningStatus.md) |  | [optional] 
**OperatingStatus** | Pointer to [**NullableLoadBalancerOperatingStatus**](LoadBalancerOperatingStatus.md) |  | [optional] 
**AccessLogs** | Pointer to [**NullableAccessLogsModel**](AccessLogsModel.md) |  | [optional] 
**ProjectId** | **string** | 로드 밸런서가 속한 프로젝트의 ID | 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 
**SubnetId** | Pointer to **NullableString** |  | [optional] 
**Listeners** | Pointer to [**[]BnsLoadBalancerV1ApiUpdateAccessLogModelListenerIdModel**](BnsLoadBalancerV1ApiUpdateAccessLogModelListenerIdModel.md) |  | [optional] 
**Provider** | Pointer to **NullableString** |  | [optional] 
**FlavorId** | Pointer to **NullableString** |  | [optional] 
**AvailabilityZone** | Pointer to [**NullableAvailabilityZone**](AvailabilityZone.md) |  | [optional] 
**TargetGroupCount** | **int32** | 로드 밸런서에 연결된 대상 그룹의 개수 | 

## Methods

### NewBnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel

`func NewBnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel(id string, name string, projectId string, targetGroupCount int32, ) *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel`

NewBnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel instantiates a new BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModelWithDefaults

`func NewBnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModelWithDefaults() *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel`

NewBnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModelWithDefaults instantiates a new BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetProvisioningStatus

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) GetProvisioningStatus() ProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) GetProvisioningStatusOk() (*ProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) SetProvisioningStatus(v ProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.

### HasProvisioningStatus

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) HasProvisioningStatus() bool`

HasProvisioningStatus returns a boolean if a field has been set.

### SetProvisioningStatusNil

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) SetProvisioningStatusNil(b bool)`

 SetProvisioningStatusNil sets the value for ProvisioningStatus to be an explicit nil

### UnsetProvisioningStatus
`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) UnsetProvisioningStatus()`

UnsetProvisioningStatus ensures that no value is present for ProvisioningStatus, not even an explicit nil
### GetOperatingStatus

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) GetOperatingStatus() LoadBalancerOperatingStatus`

GetOperatingStatus returns the OperatingStatus field if non-nil, zero value otherwise.

### GetOperatingStatusOk

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) GetOperatingStatusOk() (*LoadBalancerOperatingStatus, bool)`

GetOperatingStatusOk returns a tuple with the OperatingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingStatus

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) SetOperatingStatus(v LoadBalancerOperatingStatus)`

SetOperatingStatus sets OperatingStatus field to given value.

### HasOperatingStatus

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) HasOperatingStatus() bool`

HasOperatingStatus returns a boolean if a field has been set.

### SetOperatingStatusNil

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) SetOperatingStatusNil(b bool)`

 SetOperatingStatusNil sets the value for OperatingStatus to be an explicit nil

### UnsetOperatingStatus
`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) UnsetOperatingStatus()`

UnsetOperatingStatus ensures that no value is present for OperatingStatus, not even an explicit nil
### GetAccessLogs

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) GetAccessLogs() AccessLogsModel`

GetAccessLogs returns the AccessLogs field if non-nil, zero value otherwise.

### GetAccessLogsOk

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) GetAccessLogsOk() (*AccessLogsModel, bool)`

GetAccessLogsOk returns a tuple with the AccessLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLogs

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) SetAccessLogs(v AccessLogsModel)`

SetAccessLogs sets AccessLogs field to given value.

### HasAccessLogs

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) HasAccessLogs() bool`

HasAccessLogs returns a boolean if a field has been set.

### SetAccessLogsNil

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) SetAccessLogsNil(b bool)`

 SetAccessLogsNil sets the value for AccessLogs to be an explicit nil

### UnsetAccessLogs
`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) UnsetAccessLogs()`

UnsetAccessLogs ensures that no value is present for AccessLogs, not even an explicit nil
### GetProjectId

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetCreatedAt

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetSubnetId

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.

### HasSubnetId

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.

### SetSubnetIdNil

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) SetSubnetIdNil(b bool)`

 SetSubnetIdNil sets the value for SubnetId to be an explicit nil

### UnsetSubnetId
`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) UnsetSubnetId()`

UnsetSubnetId ensures that no value is present for SubnetId, not even an explicit nil
### GetListeners

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) GetListeners() []BnsLoadBalancerV1ApiUpdateAccessLogModelListenerIdModel`

GetListeners returns the Listeners field if non-nil, zero value otherwise.

### GetListenersOk

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) GetListenersOk() (*[]BnsLoadBalancerV1ApiUpdateAccessLogModelListenerIdModel, bool)`

GetListenersOk returns a tuple with the Listeners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListeners

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) SetListeners(v []BnsLoadBalancerV1ApiUpdateAccessLogModelListenerIdModel)`

SetListeners sets Listeners field to given value.

### HasListeners

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) HasListeners() bool`

HasListeners returns a boolean if a field has been set.

### SetListenersNil

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) SetListenersNil(b bool)`

 SetListenersNil sets the value for Listeners to be an explicit nil

### UnsetListeners
`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) UnsetListeners()`

UnsetListeners ensures that no value is present for Listeners, not even an explicit nil
### GetProvider

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### SetProviderNil

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) SetProviderNil(b bool)`

 SetProviderNil sets the value for Provider to be an explicit nil

### UnsetProvider
`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) UnsetProvider()`

UnsetProvider ensures that no value is present for Provider, not even an explicit nil
### GetFlavorId

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) GetFlavorId() string`

GetFlavorId returns the FlavorId field if non-nil, zero value otherwise.

### GetFlavorIdOk

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) GetFlavorIdOk() (*string, bool)`

GetFlavorIdOk returns a tuple with the FlavorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavorId

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) SetFlavorId(v string)`

SetFlavorId sets FlavorId field to given value.

### HasFlavorId

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) HasFlavorId() bool`

HasFlavorId returns a boolean if a field has been set.

### SetFlavorIdNil

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) SetFlavorIdNil(b bool)`

 SetFlavorIdNil sets the value for FlavorId to be an explicit nil

### UnsetFlavorId
`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) UnsetFlavorId()`

UnsetFlavorId ensures that no value is present for FlavorId, not even an explicit nil
### GetAvailabilityZone

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) GetAvailabilityZone() AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) GetAvailabilityZoneOk() (*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) SetAvailabilityZone(v AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### SetAvailabilityZoneNil

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) SetAvailabilityZoneNil(b bool)`

 SetAvailabilityZoneNil sets the value for AvailabilityZone to be an explicit nil

### UnsetAvailabilityZone
`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) UnsetAvailabilityZone()`

UnsetAvailabilityZone ensures that no value is present for AvailabilityZone, not even an explicit nil
### GetTargetGroupCount

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) GetTargetGroupCount() int32`

GetTargetGroupCount returns the TargetGroupCount field if non-nil, zero value otherwise.

### GetTargetGroupCountOk

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) GetTargetGroupCountOk() (*int32, bool)`

GetTargetGroupCountOk returns a tuple with the TargetGroupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroupCount

`func (o *BnsLoadBalancerV1ApiUpdateAccessLogModelLoadBalancerModel) SetTargetGroupCount(v int32)`

SetTargetGroupCount sets TargetGroupCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


