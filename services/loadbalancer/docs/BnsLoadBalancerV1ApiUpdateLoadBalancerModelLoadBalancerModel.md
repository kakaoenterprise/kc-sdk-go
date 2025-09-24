# BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 로드 밸런서의 ID | 
**Name** | **string** | 로드 밸런서의 이름 | 
**Description** | Pointer to **NullableString** |  | [optional] 
**ProvisioningStatus** | Pointer to [**NullableProvisioningStatus**](ProvisioningStatus.md) |  | [optional] 
**OperatingStatus** | Pointer to [**NullableLoadBalancerOperatingStatus**](LoadBalancerOperatingStatus.md) |  | [optional] 
**AccessLogs** | Pointer to [**NullableAccessLogs**](AccessLogs.md) |  | [optional] 
**ProjectId** | **string** | 해당 리소스가 속한 프로젝트의 ID | 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 
**SubnetId** | Pointer to **NullableString** |  | [optional] 
**Listeners** | Pointer to [**[]BnsLoadBalancerV1ApiUpdateLoadBalancerModelListenerIdModel**](BnsLoadBalancerV1ApiUpdateLoadBalancerModelListenerIdModel.md) |  | [optional] 
**Provider** | Pointer to **NullableString** |  | [optional] 
**FlavorId** | Pointer to **NullableString** |  | [optional] 
**AvailabilityZone** | Pointer to [**NullableAvailabilityZone**](AvailabilityZone.md) |  | [optional] 
**TargetGroupCount** | **int32** | 연결된 대상 그룹 수 | 

## Methods

### NewBnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel

`func NewBnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel(id string, name string, projectId string, targetGroupCount int32, ) *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel`

NewBnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel instantiates a new BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModelWithDefaults

`func NewBnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModelWithDefaults() *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel`

NewBnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModelWithDefaults instantiates a new BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetProvisioningStatus

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) GetProvisioningStatus() ProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) GetProvisioningStatusOk() (*ProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) SetProvisioningStatus(v ProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.

### HasProvisioningStatus

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) HasProvisioningStatus() bool`

HasProvisioningStatus returns a boolean if a field has been set.

### SetProvisioningStatusNil

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) SetProvisioningStatusNil(b bool)`

 SetProvisioningStatusNil sets the value for ProvisioningStatus to be an explicit nil

### UnsetProvisioningStatus
`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) UnsetProvisioningStatus()`

UnsetProvisioningStatus ensures that no value is present for ProvisioningStatus, not even an explicit nil
### GetOperatingStatus

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) GetOperatingStatus() LoadBalancerOperatingStatus`

GetOperatingStatus returns the OperatingStatus field if non-nil, zero value otherwise.

### GetOperatingStatusOk

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) GetOperatingStatusOk() (*LoadBalancerOperatingStatus, bool)`

GetOperatingStatusOk returns a tuple with the OperatingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingStatus

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) SetOperatingStatus(v LoadBalancerOperatingStatus)`

SetOperatingStatus sets OperatingStatus field to given value.

### HasOperatingStatus

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) HasOperatingStatus() bool`

HasOperatingStatus returns a boolean if a field has been set.

### SetOperatingStatusNil

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) SetOperatingStatusNil(b bool)`

 SetOperatingStatusNil sets the value for OperatingStatus to be an explicit nil

### UnsetOperatingStatus
`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) UnsetOperatingStatus()`

UnsetOperatingStatus ensures that no value is present for OperatingStatus, not even an explicit nil
### GetAccessLogs

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) GetAccessLogs() AccessLogs`

GetAccessLogs returns the AccessLogs field if non-nil, zero value otherwise.

### GetAccessLogsOk

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) GetAccessLogsOk() (*AccessLogs, bool)`

GetAccessLogsOk returns a tuple with the AccessLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLogs

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) SetAccessLogs(v AccessLogs)`

SetAccessLogs sets AccessLogs field to given value.

### HasAccessLogs

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) HasAccessLogs() bool`

HasAccessLogs returns a boolean if a field has been set.

### SetAccessLogsNil

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) SetAccessLogsNil(b bool)`

 SetAccessLogsNil sets the value for AccessLogs to be an explicit nil

### UnsetAccessLogs
`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) UnsetAccessLogs()`

UnsetAccessLogs ensures that no value is present for AccessLogs, not even an explicit nil
### GetProjectId

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetCreatedAt

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetSubnetId

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.

### HasSubnetId

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.

### SetSubnetIdNil

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) SetSubnetIdNil(b bool)`

 SetSubnetIdNil sets the value for SubnetId to be an explicit nil

### UnsetSubnetId
`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) UnsetSubnetId()`

UnsetSubnetId ensures that no value is present for SubnetId, not even an explicit nil
### GetListeners

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) GetListeners() []BnsLoadBalancerV1ApiUpdateLoadBalancerModelListenerIdModel`

GetListeners returns the Listeners field if non-nil, zero value otherwise.

### GetListenersOk

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) GetListenersOk() (*[]BnsLoadBalancerV1ApiUpdateLoadBalancerModelListenerIdModel, bool)`

GetListenersOk returns a tuple with the Listeners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListeners

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) SetListeners(v []BnsLoadBalancerV1ApiUpdateLoadBalancerModelListenerIdModel)`

SetListeners sets Listeners field to given value.

### HasListeners

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) HasListeners() bool`

HasListeners returns a boolean if a field has been set.

### SetListenersNil

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) SetListenersNil(b bool)`

 SetListenersNil sets the value for Listeners to be an explicit nil

### UnsetListeners
`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) UnsetListeners()`

UnsetListeners ensures that no value is present for Listeners, not even an explicit nil
### GetProvider

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### SetProviderNil

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) SetProviderNil(b bool)`

 SetProviderNil sets the value for Provider to be an explicit nil

### UnsetProvider
`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) UnsetProvider()`

UnsetProvider ensures that no value is present for Provider, not even an explicit nil
### GetFlavorId

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) GetFlavorId() string`

GetFlavorId returns the FlavorId field if non-nil, zero value otherwise.

### GetFlavorIdOk

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) GetFlavorIdOk() (*string, bool)`

GetFlavorIdOk returns a tuple with the FlavorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavorId

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) SetFlavorId(v string)`

SetFlavorId sets FlavorId field to given value.

### HasFlavorId

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) HasFlavorId() bool`

HasFlavorId returns a boolean if a field has been set.

### SetFlavorIdNil

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) SetFlavorIdNil(b bool)`

 SetFlavorIdNil sets the value for FlavorId to be an explicit nil

### UnsetFlavorId
`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) UnsetFlavorId()`

UnsetFlavorId ensures that no value is present for FlavorId, not even an explicit nil
### GetAvailabilityZone

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) GetAvailabilityZone() AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) GetAvailabilityZoneOk() (*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) SetAvailabilityZone(v AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### SetAvailabilityZoneNil

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) SetAvailabilityZoneNil(b bool)`

 SetAvailabilityZoneNil sets the value for AvailabilityZone to be an explicit nil

### UnsetAvailabilityZone
`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) UnsetAvailabilityZone()`

UnsetAvailabilityZone ensures that no value is present for AvailabilityZone, not even an explicit nil
### GetTargetGroupCount

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) GetTargetGroupCount() int32`

GetTargetGroupCount returns the TargetGroupCount field if non-nil, zero value otherwise.

### GetTargetGroupCountOk

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) GetTargetGroupCountOk() (*int32, bool)`

GetTargetGroupCountOk returns a tuple with the TargetGroupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroupCount

`func (o *BnsLoadBalancerV1ApiUpdateLoadBalancerModelLoadBalancerModel) SetTargetGroupCount(v int32)`

SetTargetGroupCount sets TargetGroupCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


