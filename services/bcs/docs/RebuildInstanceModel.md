# RebuildInstanceModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 인스턴스 ID | 
**Name** | **string** | 인스턴스의 이름 | 
**ProjectId** | **string** | 인스턴스가 속한 프로젝트의 ID | 
**UserId** | **string** | 인스턴스를 생성한 사용자의 ID | 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Flavor** | [**BcsInstanceV1ApiRebuildInstanceModelInstanceFlavorModel**](BcsInstanceV1ApiRebuildInstanceModelInstanceFlavorModel.md) | 인스턴스의 가상 하드웨어 프로파일 정보 | 
**CreatedAt** | **time.Time** | 리소스가 생성된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
**UpdatedAt** | **time.Time** | 리소스가 마지막으로 수정된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
**Addresses** | Pointer to **[]map[string]string** |  | [optional] 
**AvailabilityZone** | Pointer to [**NullableAvailabilityZone**](AvailabilityZone.md) |  | [optional] 
**KeyName** | Pointer to **NullableString** |  | [optional] 
**TerminatedAt** | Pointer to **NullableString** |  | [optional] 
**SecurityGroups** | Pointer to [**[]BcsInstanceV1ApiRebuildInstanceModelInstanceSecurityGroupModel**](BcsInstanceV1ApiRebuildInstanceModelInstanceSecurityGroupModel.md) |  | [optional] 
**TaskState** | Pointer to **NullableString** |  | [optional] 
**VmState** | Pointer to **NullableString** |  | [optional] 
**PowerState** | Pointer to [**NullablePowerState**](PowerState.md) |  | [optional] 
**AttachedVolumes** | Pointer to [**[]BcsInstanceV1ApiRebuildInstanceModelInstanceAttachedVolumeModel**](BcsInstanceV1ApiRebuildInstanceModelInstanceAttachedVolumeModel.md) |  | [optional] 
**Description** | **string** | 인스턴스에 대한 설명 | 
**Hostname** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewRebuildInstanceModel

`func NewRebuildInstanceModel(id string, name string, projectId string, userId string, flavor BcsInstanceV1ApiRebuildInstanceModelInstanceFlavorModel, createdAt time.Time, updatedAt time.Time, description string, ) *RebuildInstanceModel`

NewRebuildInstanceModel instantiates a new RebuildInstanceModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRebuildInstanceModelWithDefaults

`func NewRebuildInstanceModelWithDefaults() *RebuildInstanceModel`

NewRebuildInstanceModelWithDefaults instantiates a new RebuildInstanceModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RebuildInstanceModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RebuildInstanceModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RebuildInstanceModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *RebuildInstanceModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RebuildInstanceModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RebuildInstanceModel) SetName(v string)`

SetName sets Name field to given value.


### GetProjectId

`func (o *RebuildInstanceModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *RebuildInstanceModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *RebuildInstanceModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetUserId

`func (o *RebuildInstanceModel) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *RebuildInstanceModel) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *RebuildInstanceModel) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetMetadata

`func (o *RebuildInstanceModel) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RebuildInstanceModel) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RebuildInstanceModel) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RebuildInstanceModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *RebuildInstanceModel) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *RebuildInstanceModel) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetFlavor

`func (o *RebuildInstanceModel) GetFlavor() BcsInstanceV1ApiRebuildInstanceModelInstanceFlavorModel`

GetFlavor returns the Flavor field if non-nil, zero value otherwise.

### GetFlavorOk

`func (o *RebuildInstanceModel) GetFlavorOk() (*BcsInstanceV1ApiRebuildInstanceModelInstanceFlavorModel, bool)`

GetFlavorOk returns a tuple with the Flavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavor

`func (o *RebuildInstanceModel) SetFlavor(v BcsInstanceV1ApiRebuildInstanceModelInstanceFlavorModel)`

SetFlavor sets Flavor field to given value.


### GetCreatedAt

`func (o *RebuildInstanceModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RebuildInstanceModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RebuildInstanceModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *RebuildInstanceModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RebuildInstanceModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RebuildInstanceModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetAddresses

`func (o *RebuildInstanceModel) GetAddresses() []map[string]string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *RebuildInstanceModel) GetAddressesOk() (*[]map[string]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *RebuildInstanceModel) SetAddresses(v []map[string]string)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *RebuildInstanceModel) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### SetAddressesNil

`func (o *RebuildInstanceModel) SetAddressesNil(b bool)`

 SetAddressesNil sets the value for Addresses to be an explicit nil

### UnsetAddresses
`func (o *RebuildInstanceModel) UnsetAddresses()`

UnsetAddresses ensures that no value is present for Addresses, not even an explicit nil
### GetAvailabilityZone

`func (o *RebuildInstanceModel) GetAvailabilityZone() AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *RebuildInstanceModel) GetAvailabilityZoneOk() (*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *RebuildInstanceModel) SetAvailabilityZone(v AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *RebuildInstanceModel) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### SetAvailabilityZoneNil

`func (o *RebuildInstanceModel) SetAvailabilityZoneNil(b bool)`

 SetAvailabilityZoneNil sets the value for AvailabilityZone to be an explicit nil

### UnsetAvailabilityZone
`func (o *RebuildInstanceModel) UnsetAvailabilityZone()`

UnsetAvailabilityZone ensures that no value is present for AvailabilityZone, not even an explicit nil
### GetKeyName

`func (o *RebuildInstanceModel) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *RebuildInstanceModel) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *RebuildInstanceModel) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.

### HasKeyName

`func (o *RebuildInstanceModel) HasKeyName() bool`

HasKeyName returns a boolean if a field has been set.

### SetKeyNameNil

`func (o *RebuildInstanceModel) SetKeyNameNil(b bool)`

 SetKeyNameNil sets the value for KeyName to be an explicit nil

### UnsetKeyName
`func (o *RebuildInstanceModel) UnsetKeyName()`

UnsetKeyName ensures that no value is present for KeyName, not even an explicit nil
### GetTerminatedAt

`func (o *RebuildInstanceModel) GetTerminatedAt() string`

GetTerminatedAt returns the TerminatedAt field if non-nil, zero value otherwise.

### GetTerminatedAtOk

`func (o *RebuildInstanceModel) GetTerminatedAtOk() (*string, bool)`

GetTerminatedAtOk returns a tuple with the TerminatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminatedAt

`func (o *RebuildInstanceModel) SetTerminatedAt(v string)`

SetTerminatedAt sets TerminatedAt field to given value.

### HasTerminatedAt

`func (o *RebuildInstanceModel) HasTerminatedAt() bool`

HasTerminatedAt returns a boolean if a field has been set.

### SetTerminatedAtNil

`func (o *RebuildInstanceModel) SetTerminatedAtNil(b bool)`

 SetTerminatedAtNil sets the value for TerminatedAt to be an explicit nil

### UnsetTerminatedAt
`func (o *RebuildInstanceModel) UnsetTerminatedAt()`

UnsetTerminatedAt ensures that no value is present for TerminatedAt, not even an explicit nil
### GetSecurityGroups

`func (o *RebuildInstanceModel) GetSecurityGroups() []BcsInstanceV1ApiRebuildInstanceModelInstanceSecurityGroupModel`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *RebuildInstanceModel) GetSecurityGroupsOk() (*[]BcsInstanceV1ApiRebuildInstanceModelInstanceSecurityGroupModel, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *RebuildInstanceModel) SetSecurityGroups(v []BcsInstanceV1ApiRebuildInstanceModelInstanceSecurityGroupModel)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *RebuildInstanceModel) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### SetSecurityGroupsNil

`func (o *RebuildInstanceModel) SetSecurityGroupsNil(b bool)`

 SetSecurityGroupsNil sets the value for SecurityGroups to be an explicit nil

### UnsetSecurityGroups
`func (o *RebuildInstanceModel) UnsetSecurityGroups()`

UnsetSecurityGroups ensures that no value is present for SecurityGroups, not even an explicit nil
### GetTaskState

`func (o *RebuildInstanceModel) GetTaskState() string`

GetTaskState returns the TaskState field if non-nil, zero value otherwise.

### GetTaskStateOk

`func (o *RebuildInstanceModel) GetTaskStateOk() (*string, bool)`

GetTaskStateOk returns a tuple with the TaskState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskState

`func (o *RebuildInstanceModel) SetTaskState(v string)`

SetTaskState sets TaskState field to given value.

### HasTaskState

`func (o *RebuildInstanceModel) HasTaskState() bool`

HasTaskState returns a boolean if a field has been set.

### SetTaskStateNil

`func (o *RebuildInstanceModel) SetTaskStateNil(b bool)`

 SetTaskStateNil sets the value for TaskState to be an explicit nil

### UnsetTaskState
`func (o *RebuildInstanceModel) UnsetTaskState()`

UnsetTaskState ensures that no value is present for TaskState, not even an explicit nil
### GetVmState

`func (o *RebuildInstanceModel) GetVmState() string`

GetVmState returns the VmState field if non-nil, zero value otherwise.

### GetVmStateOk

`func (o *RebuildInstanceModel) GetVmStateOk() (*string, bool)`

GetVmStateOk returns a tuple with the VmState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmState

`func (o *RebuildInstanceModel) SetVmState(v string)`

SetVmState sets VmState field to given value.

### HasVmState

`func (o *RebuildInstanceModel) HasVmState() bool`

HasVmState returns a boolean if a field has been set.

### SetVmStateNil

`func (o *RebuildInstanceModel) SetVmStateNil(b bool)`

 SetVmStateNil sets the value for VmState to be an explicit nil

### UnsetVmState
`func (o *RebuildInstanceModel) UnsetVmState()`

UnsetVmState ensures that no value is present for VmState, not even an explicit nil
### GetPowerState

`func (o *RebuildInstanceModel) GetPowerState() PowerState`

GetPowerState returns the PowerState field if non-nil, zero value otherwise.

### GetPowerStateOk

`func (o *RebuildInstanceModel) GetPowerStateOk() (*PowerState, bool)`

GetPowerStateOk returns a tuple with the PowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerState

`func (o *RebuildInstanceModel) SetPowerState(v PowerState)`

SetPowerState sets PowerState field to given value.

### HasPowerState

`func (o *RebuildInstanceModel) HasPowerState() bool`

HasPowerState returns a boolean if a field has been set.

### SetPowerStateNil

`func (o *RebuildInstanceModel) SetPowerStateNil(b bool)`

 SetPowerStateNil sets the value for PowerState to be an explicit nil

### UnsetPowerState
`func (o *RebuildInstanceModel) UnsetPowerState()`

UnsetPowerState ensures that no value is present for PowerState, not even an explicit nil
### GetAttachedVolumes

`func (o *RebuildInstanceModel) GetAttachedVolumes() []BcsInstanceV1ApiRebuildInstanceModelInstanceAttachedVolumeModel`

GetAttachedVolumes returns the AttachedVolumes field if non-nil, zero value otherwise.

### GetAttachedVolumesOk

`func (o *RebuildInstanceModel) GetAttachedVolumesOk() (*[]BcsInstanceV1ApiRebuildInstanceModelInstanceAttachedVolumeModel, bool)`

GetAttachedVolumesOk returns a tuple with the AttachedVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedVolumes

`func (o *RebuildInstanceModel) SetAttachedVolumes(v []BcsInstanceV1ApiRebuildInstanceModelInstanceAttachedVolumeModel)`

SetAttachedVolumes sets AttachedVolumes field to given value.

### HasAttachedVolumes

`func (o *RebuildInstanceModel) HasAttachedVolumes() bool`

HasAttachedVolumes returns a boolean if a field has been set.

### SetAttachedVolumesNil

`func (o *RebuildInstanceModel) SetAttachedVolumesNil(b bool)`

 SetAttachedVolumesNil sets the value for AttachedVolumes to be an explicit nil

### UnsetAttachedVolumes
`func (o *RebuildInstanceModel) UnsetAttachedVolumes()`

UnsetAttachedVolumes ensures that no value is present for AttachedVolumes, not even an explicit nil
### GetDescription

`func (o *RebuildInstanceModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RebuildInstanceModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RebuildInstanceModel) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetHostname

`func (o *RebuildInstanceModel) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *RebuildInstanceModel) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *RebuildInstanceModel) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *RebuildInstanceModel) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### SetHostnameNil

`func (o *RebuildInstanceModel) SetHostnameNil(b bool)`

 SetHostnameNil sets the value for Hostname to be an explicit nil

### UnsetHostname
`func (o *RebuildInstanceModel) UnsetHostname()`

UnsetHostname ensures that no value is present for Hostname, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


