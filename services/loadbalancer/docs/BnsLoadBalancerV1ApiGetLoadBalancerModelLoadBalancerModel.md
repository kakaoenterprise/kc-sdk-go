# BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 로드 밸런서 ID | 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to [**NullableLoadBalancerType**](LoadBalancerType.md) |  | [optional] 
**ListenerIds** | Pointer to **[]string** |  | [optional] 
**ProjectId** | Pointer to **NullableString** |  | [optional] 
**ProvisioningStatus** | Pointer to [**NullableProvisioningStatus**](ProvisioningStatus.md) |  | [optional] 
**OperatingStatus** | Pointer to [**NullableLoadBalancerOperatingStatus**](LoadBalancerOperatingStatus.md) |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 
**AvailabilityZone** | Pointer to [**NullableAvailabilityZone**](AvailabilityZone.md) |  | [optional] 
**AccessLogs** | Pointer to **NullableString** |  | [optional] 
**BeyondLoadBalancerId** | Pointer to **NullableString** |  | [optional] 
**BeyondLoadBalancerName** | Pointer to **NullableString** |  | [optional] 
**BeyondLoadBalancerDnsName** | Pointer to **NullableString** |  | [optional] 
**TargetGroupCount** | Pointer to **NullableInt64** |  | [optional] 
**ListenerCount** | Pointer to **NullableInt64** |  | [optional] 
**PrivateVip** | Pointer to **NullableString** |  | [optional] 
**PublicVip** | Pointer to **NullableString** |  | [optional] 
**SubnetName** | Pointer to **NullableString** |  | [optional] 
**SubnetCidrBlock** | Pointer to **NullableString** |  | [optional] 
**VpcId** | Pointer to **NullableString** |  | [optional] 
**VpcName** | Pointer to **NullableString** |  | [optional] 
**SubnetId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel

`func NewBnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel(id string, ) *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel`

NewBnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel instantiates a new BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModelWithDefaults

`func NewBnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModelWithDefaults() *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel`

NewBnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModelWithDefaults instantiates a new BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetType

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) GetType() LoadBalancerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) GetTypeOk() (*LoadBalancerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) SetType(v LoadBalancerType)`

SetType sets Type field to given value.

### HasType

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetListenerIds

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) GetListenerIds() []string`

GetListenerIds returns the ListenerIds field if non-nil, zero value otherwise.

### GetListenerIdsOk

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) GetListenerIdsOk() (*[]string, bool)`

GetListenerIdsOk returns a tuple with the ListenerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenerIds

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) SetListenerIds(v []string)`

SetListenerIds sets ListenerIds field to given value.

### HasListenerIds

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) HasListenerIds() bool`

HasListenerIds returns a boolean if a field has been set.

### SetListenerIdsNil

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) SetListenerIdsNil(b bool)`

 SetListenerIdsNil sets the value for ListenerIds to be an explicit nil

### UnsetListenerIds
`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) UnsetListenerIds()`

UnsetListenerIds ensures that no value is present for ListenerIds, not even an explicit nil
### GetProjectId

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetProvisioningStatus

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) GetProvisioningStatus() ProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) GetProvisioningStatusOk() (*ProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) SetProvisioningStatus(v ProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.

### HasProvisioningStatus

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) HasProvisioningStatus() bool`

HasProvisioningStatus returns a boolean if a field has been set.

### SetProvisioningStatusNil

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) SetProvisioningStatusNil(b bool)`

 SetProvisioningStatusNil sets the value for ProvisioningStatus to be an explicit nil

### UnsetProvisioningStatus
`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) UnsetProvisioningStatus()`

UnsetProvisioningStatus ensures that no value is present for ProvisioningStatus, not even an explicit nil
### GetOperatingStatus

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) GetOperatingStatus() LoadBalancerOperatingStatus`

GetOperatingStatus returns the OperatingStatus field if non-nil, zero value otherwise.

### GetOperatingStatusOk

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) GetOperatingStatusOk() (*LoadBalancerOperatingStatus, bool)`

GetOperatingStatusOk returns a tuple with the OperatingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingStatus

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) SetOperatingStatus(v LoadBalancerOperatingStatus)`

SetOperatingStatus sets OperatingStatus field to given value.

### HasOperatingStatus

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) HasOperatingStatus() bool`

HasOperatingStatus returns a boolean if a field has been set.

### SetOperatingStatusNil

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) SetOperatingStatusNil(b bool)`

 SetOperatingStatusNil sets the value for OperatingStatus to be an explicit nil

### UnsetOperatingStatus
`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) UnsetOperatingStatus()`

UnsetOperatingStatus ensures that no value is present for OperatingStatus, not even an explicit nil
### GetCreatedAt

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetAvailabilityZone

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) GetAvailabilityZone() AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) GetAvailabilityZoneOk() (*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) SetAvailabilityZone(v AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### SetAvailabilityZoneNil

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) SetAvailabilityZoneNil(b bool)`

 SetAvailabilityZoneNil sets the value for AvailabilityZone to be an explicit nil

### UnsetAvailabilityZone
`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) UnsetAvailabilityZone()`

UnsetAvailabilityZone ensures that no value is present for AvailabilityZone, not even an explicit nil
### GetAccessLogs

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) GetAccessLogs() string`

GetAccessLogs returns the AccessLogs field if non-nil, zero value otherwise.

### GetAccessLogsOk

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) GetAccessLogsOk() (*string, bool)`

GetAccessLogsOk returns a tuple with the AccessLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLogs

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) SetAccessLogs(v string)`

SetAccessLogs sets AccessLogs field to given value.

### HasAccessLogs

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) HasAccessLogs() bool`

HasAccessLogs returns a boolean if a field has been set.

### SetAccessLogsNil

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) SetAccessLogsNil(b bool)`

 SetAccessLogsNil sets the value for AccessLogs to be an explicit nil

### UnsetAccessLogs
`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) UnsetAccessLogs()`

UnsetAccessLogs ensures that no value is present for AccessLogs, not even an explicit nil
### GetBeyondLoadBalancerId

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) GetBeyondLoadBalancerId() string`

GetBeyondLoadBalancerId returns the BeyondLoadBalancerId field if non-nil, zero value otherwise.

### GetBeyondLoadBalancerIdOk

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) GetBeyondLoadBalancerIdOk() (*string, bool)`

GetBeyondLoadBalancerIdOk returns a tuple with the BeyondLoadBalancerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeyondLoadBalancerId

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) SetBeyondLoadBalancerId(v string)`

SetBeyondLoadBalancerId sets BeyondLoadBalancerId field to given value.

### HasBeyondLoadBalancerId

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) HasBeyondLoadBalancerId() bool`

HasBeyondLoadBalancerId returns a boolean if a field has been set.

### SetBeyondLoadBalancerIdNil

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) SetBeyondLoadBalancerIdNil(b bool)`

 SetBeyondLoadBalancerIdNil sets the value for BeyondLoadBalancerId to be an explicit nil

### UnsetBeyondLoadBalancerId
`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) UnsetBeyondLoadBalancerId()`

UnsetBeyondLoadBalancerId ensures that no value is present for BeyondLoadBalancerId, not even an explicit nil
### GetBeyondLoadBalancerName

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) GetBeyondLoadBalancerName() string`

GetBeyondLoadBalancerName returns the BeyondLoadBalancerName field if non-nil, zero value otherwise.

### GetBeyondLoadBalancerNameOk

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) GetBeyondLoadBalancerNameOk() (*string, bool)`

GetBeyondLoadBalancerNameOk returns a tuple with the BeyondLoadBalancerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeyondLoadBalancerName

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) SetBeyondLoadBalancerName(v string)`

SetBeyondLoadBalancerName sets BeyondLoadBalancerName field to given value.

### HasBeyondLoadBalancerName

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) HasBeyondLoadBalancerName() bool`

HasBeyondLoadBalancerName returns a boolean if a field has been set.

### SetBeyondLoadBalancerNameNil

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) SetBeyondLoadBalancerNameNil(b bool)`

 SetBeyondLoadBalancerNameNil sets the value for BeyondLoadBalancerName to be an explicit nil

### UnsetBeyondLoadBalancerName
`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) UnsetBeyondLoadBalancerName()`

UnsetBeyondLoadBalancerName ensures that no value is present for BeyondLoadBalancerName, not even an explicit nil
### GetBeyondLoadBalancerDnsName

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) GetBeyondLoadBalancerDnsName() string`

GetBeyondLoadBalancerDnsName returns the BeyondLoadBalancerDnsName field if non-nil, zero value otherwise.

### GetBeyondLoadBalancerDnsNameOk

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) GetBeyondLoadBalancerDnsNameOk() (*string, bool)`

GetBeyondLoadBalancerDnsNameOk returns a tuple with the BeyondLoadBalancerDnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeyondLoadBalancerDnsName

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) SetBeyondLoadBalancerDnsName(v string)`

SetBeyondLoadBalancerDnsName sets BeyondLoadBalancerDnsName field to given value.

### HasBeyondLoadBalancerDnsName

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) HasBeyondLoadBalancerDnsName() bool`

HasBeyondLoadBalancerDnsName returns a boolean if a field has been set.

### SetBeyondLoadBalancerDnsNameNil

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) SetBeyondLoadBalancerDnsNameNil(b bool)`

 SetBeyondLoadBalancerDnsNameNil sets the value for BeyondLoadBalancerDnsName to be an explicit nil

### UnsetBeyondLoadBalancerDnsName
`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) UnsetBeyondLoadBalancerDnsName()`

UnsetBeyondLoadBalancerDnsName ensures that no value is present for BeyondLoadBalancerDnsName, not even an explicit nil
### GetTargetGroupCount

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) GetTargetGroupCount() int64`

GetTargetGroupCount returns the TargetGroupCount field if non-nil, zero value otherwise.

### GetTargetGroupCountOk

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) GetTargetGroupCountOk() (*int64, bool)`

GetTargetGroupCountOk returns a tuple with the TargetGroupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroupCount

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) SetTargetGroupCount(v int64)`

SetTargetGroupCount sets TargetGroupCount field to given value.

### HasTargetGroupCount

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) HasTargetGroupCount() bool`

HasTargetGroupCount returns a boolean if a field has been set.

### SetTargetGroupCountNil

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) SetTargetGroupCountNil(b bool)`

 SetTargetGroupCountNil sets the value for TargetGroupCount to be an explicit nil

### UnsetTargetGroupCount
`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) UnsetTargetGroupCount()`

UnsetTargetGroupCount ensures that no value is present for TargetGroupCount, not even an explicit nil
### GetListenerCount

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) GetListenerCount() int64`

GetListenerCount returns the ListenerCount field if non-nil, zero value otherwise.

### GetListenerCountOk

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) GetListenerCountOk() (*int64, bool)`

GetListenerCountOk returns a tuple with the ListenerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenerCount

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) SetListenerCount(v int64)`

SetListenerCount sets ListenerCount field to given value.

### HasListenerCount

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) HasListenerCount() bool`

HasListenerCount returns a boolean if a field has been set.

### SetListenerCountNil

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) SetListenerCountNil(b bool)`

 SetListenerCountNil sets the value for ListenerCount to be an explicit nil

### UnsetListenerCount
`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) UnsetListenerCount()`

UnsetListenerCount ensures that no value is present for ListenerCount, not even an explicit nil
### GetPrivateVip

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) GetPrivateVip() string`

GetPrivateVip returns the PrivateVip field if non-nil, zero value otherwise.

### GetPrivateVipOk

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) GetPrivateVipOk() (*string, bool)`

GetPrivateVipOk returns a tuple with the PrivateVip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateVip

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) SetPrivateVip(v string)`

SetPrivateVip sets PrivateVip field to given value.

### HasPrivateVip

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) HasPrivateVip() bool`

HasPrivateVip returns a boolean if a field has been set.

### SetPrivateVipNil

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) SetPrivateVipNil(b bool)`

 SetPrivateVipNil sets the value for PrivateVip to be an explicit nil

### UnsetPrivateVip
`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) UnsetPrivateVip()`

UnsetPrivateVip ensures that no value is present for PrivateVip, not even an explicit nil
### GetPublicVip

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) GetPublicVip() string`

GetPublicVip returns the PublicVip field if non-nil, zero value otherwise.

### GetPublicVipOk

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) GetPublicVipOk() (*string, bool)`

GetPublicVipOk returns a tuple with the PublicVip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicVip

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) SetPublicVip(v string)`

SetPublicVip sets PublicVip field to given value.

### HasPublicVip

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) HasPublicVip() bool`

HasPublicVip returns a boolean if a field has been set.

### SetPublicVipNil

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) SetPublicVipNil(b bool)`

 SetPublicVipNil sets the value for PublicVip to be an explicit nil

### UnsetPublicVip
`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) UnsetPublicVip()`

UnsetPublicVip ensures that no value is present for PublicVip, not even an explicit nil
### GetSubnetName

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) GetSubnetName() string`

GetSubnetName returns the SubnetName field if non-nil, zero value otherwise.

### GetSubnetNameOk

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) GetSubnetNameOk() (*string, bool)`

GetSubnetNameOk returns a tuple with the SubnetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetName

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) SetSubnetName(v string)`

SetSubnetName sets SubnetName field to given value.

### HasSubnetName

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) HasSubnetName() bool`

HasSubnetName returns a boolean if a field has been set.

### SetSubnetNameNil

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) SetSubnetNameNil(b bool)`

 SetSubnetNameNil sets the value for SubnetName to be an explicit nil

### UnsetSubnetName
`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) UnsetSubnetName()`

UnsetSubnetName ensures that no value is present for SubnetName, not even an explicit nil
### GetSubnetCidrBlock

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) GetSubnetCidrBlock() string`

GetSubnetCidrBlock returns the SubnetCidrBlock field if non-nil, zero value otherwise.

### GetSubnetCidrBlockOk

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) GetSubnetCidrBlockOk() (*string, bool)`

GetSubnetCidrBlockOk returns a tuple with the SubnetCidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetCidrBlock

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) SetSubnetCidrBlock(v string)`

SetSubnetCidrBlock sets SubnetCidrBlock field to given value.

### HasSubnetCidrBlock

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) HasSubnetCidrBlock() bool`

HasSubnetCidrBlock returns a boolean if a field has been set.

### SetSubnetCidrBlockNil

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) SetSubnetCidrBlockNil(b bool)`

 SetSubnetCidrBlockNil sets the value for SubnetCidrBlock to be an explicit nil

### UnsetSubnetCidrBlock
`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) UnsetSubnetCidrBlock()`

UnsetSubnetCidrBlock ensures that no value is present for SubnetCidrBlock, not even an explicit nil
### GetVpcId

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### SetVpcIdNil

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) SetVpcIdNil(b bool)`

 SetVpcIdNil sets the value for VpcId to be an explicit nil

### UnsetVpcId
`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) UnsetVpcId()`

UnsetVpcId ensures that no value is present for VpcId, not even an explicit nil
### GetVpcName

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) GetVpcName() string`

GetVpcName returns the VpcName field if non-nil, zero value otherwise.

### GetVpcNameOk

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) GetVpcNameOk() (*string, bool)`

GetVpcNameOk returns a tuple with the VpcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcName

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) SetVpcName(v string)`

SetVpcName sets VpcName field to given value.

### HasVpcName

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) HasVpcName() bool`

HasVpcName returns a boolean if a field has been set.

### SetVpcNameNil

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) SetVpcNameNil(b bool)`

 SetVpcNameNil sets the value for VpcName to be an explicit nil

### UnsetVpcName
`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) UnsetVpcName()`

UnsetVpcName ensures that no value is present for VpcName, not even an explicit nil
### GetSubnetId

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.

### HasSubnetId

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.

### SetSubnetIdNil

`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) SetSubnetIdNil(b bool)`

 SetSubnetIdNil sets the value for SubnetId to be an explicit nil

### UnsetSubnetId
`func (o *BnsLoadBalancerV1ApiGetLoadBalancerModelLoadBalancerModel) UnsetSubnetId()`

UnsetSubnetId ensures that no value is present for SubnetId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


