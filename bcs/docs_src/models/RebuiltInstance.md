# RebuiltInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 인스턴스 ID | 
**Name** | **string** | 인스턴스의 이름 | 
**ProjectId** | **string** | 인스턴스가 속한 프로젝트의 ID | 
**UserId** | **string** | 인스턴스를 생성한 사용자의 ID | 
**Metadata** | Pointer to **map[string]string** | 인스턴스에 설정된 사용자 정의 메타데이터(Key-Value 쌍) | [optional] 
**Flavor** | [**RebuildInstanceFlavor**](RebuildInstanceFlavor.md) | 인스턴스의 가상 하드웨어 프로파일 정보 | 
**CreatedAt** | **time.Time** | 리소스가 생성된 시간 - ISO 8601 형식 - UTC 기준 | 
**UpdatedAt** | **time.Time** | 리소스가 마지막으로 수정된 시간 - ISO 8601 형식 - UTC 기준 | 
**AvailabilityZone** | Pointer to [**NullableAvailabilityZone**](AvailabilityZone.md) | 인스턴스가 위치한 가용 영역 | [optional] 
**KeyName** | Pointer to **NullableString** | 재설정할 키페어 이름 | [optional] 
**TerminatedAt** | Pointer to **NullableTime** | 인스턴스가 종료된 시점 | [optional] 
**SecurityGroups** | Pointer to [**[]SecurityGroup**](SecurityGroup.md) | 인스턴스에 적용된 보안 그룹 목록 | [optional] 
**TaskState** | Pointer to **NullableString** | 인스턴스가 현재 수행 중인 작업 상태 | [optional] 
**VmState** | Pointer to **NullableString** | 인스턴스의 가상머신 상태 | [optional] 
**PowerState** | Pointer to [**NullablePowerState**](PowerState.md) | 인스턴스의 전원 상태를 숫자로 표현한 값 | [optional] 
**AttachedVolumes** | Pointer to [**[]RebuildInstanceAttachedVolume**](RebuildInstanceAttachedVolume.md) | 인스턴스에 연결된 볼륨 정보 | [optional] 
**Description** | Pointer to **NullableString** | 인스턴스에 대한 설명 | [optional] 
**Hostname** | Pointer to **NullableString** | 인스턴스의 호스트 이름 | [optional] 
**Addresses** | Pointer to [**[]InstanceAddress**](InstanceAddress.md) | 인스턴스에 할당된 네트워크 주소 목록 | [optional] 

## Methods

### NewRebuiltInstance

`func NewRebuiltInstance(id string, name string, projectId string, userId string, flavor RebuildInstanceFlavor, createdAt time.Time, updatedAt time.Time, ) *RebuiltInstance`

NewRebuiltInstance instantiates a new RebuiltInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRebuiltInstanceWithDefaults

`func NewRebuiltInstanceWithDefaults() *RebuiltInstance`

NewRebuiltInstanceWithDefaults instantiates a new RebuiltInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RebuiltInstance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RebuiltInstance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RebuiltInstance) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *RebuiltInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RebuiltInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RebuiltInstance) SetName(v string)`

SetName sets Name field to given value.


### GetProjectId

`func (o *RebuiltInstance) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *RebuiltInstance) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *RebuiltInstance) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetUserId

`func (o *RebuiltInstance) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *RebuiltInstance) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *RebuiltInstance) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetMetadata

`func (o *RebuiltInstance) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RebuiltInstance) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RebuiltInstance) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RebuiltInstance) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *RebuiltInstance) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *RebuiltInstance) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetFlavor

`func (o *RebuiltInstance) GetFlavor() RebuildInstanceFlavor`

GetFlavor returns the Flavor field if non-nil, zero value otherwise.

### GetFlavorOk

`func (o *RebuiltInstance) GetFlavorOk() (*RebuildInstanceFlavor, bool)`

GetFlavorOk returns a tuple with the Flavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavor

`func (o *RebuiltInstance) SetFlavor(v RebuildInstanceFlavor)`

SetFlavor sets Flavor field to given value.


### GetCreatedAt

`func (o *RebuiltInstance) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RebuiltInstance) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RebuiltInstance) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *RebuiltInstance) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RebuiltInstance) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RebuiltInstance) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetAvailabilityZone

`func (o *RebuiltInstance) GetAvailabilityZone() AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *RebuiltInstance) GetAvailabilityZoneOk() (*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *RebuiltInstance) SetAvailabilityZone(v AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *RebuiltInstance) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### SetAvailabilityZoneNil

`func (o *RebuiltInstance) SetAvailabilityZoneNil(b bool)`

 SetAvailabilityZoneNil sets the value for AvailabilityZone to be an explicit nil

### UnsetAvailabilityZone
`func (o *RebuiltInstance) UnsetAvailabilityZone()`

UnsetAvailabilityZone ensures that no value is present for AvailabilityZone, not even an explicit nil
### GetKeyName

`func (o *RebuiltInstance) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *RebuiltInstance) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *RebuiltInstance) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.

### HasKeyName

`func (o *RebuiltInstance) HasKeyName() bool`

HasKeyName returns a boolean if a field has been set.

### SetKeyNameNil

`func (o *RebuiltInstance) SetKeyNameNil(b bool)`

 SetKeyNameNil sets the value for KeyName to be an explicit nil

### UnsetKeyName
`func (o *RebuiltInstance) UnsetKeyName()`

UnsetKeyName ensures that no value is present for KeyName, not even an explicit nil
### GetTerminatedAt

`func (o *RebuiltInstance) GetTerminatedAt() time.Time`

GetTerminatedAt returns the TerminatedAt field if non-nil, zero value otherwise.

### GetTerminatedAtOk

`func (o *RebuiltInstance) GetTerminatedAtOk() (*time.Time, bool)`

GetTerminatedAtOk returns a tuple with the TerminatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminatedAt

`func (o *RebuiltInstance) SetTerminatedAt(v time.Time)`

SetTerminatedAt sets TerminatedAt field to given value.

### HasTerminatedAt

`func (o *RebuiltInstance) HasTerminatedAt() bool`

HasTerminatedAt returns a boolean if a field has been set.

### SetTerminatedAtNil

`func (o *RebuiltInstance) SetTerminatedAtNil(b bool)`

 SetTerminatedAtNil sets the value for TerminatedAt to be an explicit nil

### UnsetTerminatedAt
`func (o *RebuiltInstance) UnsetTerminatedAt()`

UnsetTerminatedAt ensures that no value is present for TerminatedAt, not even an explicit nil
### GetSecurityGroups

`func (o *RebuiltInstance) GetSecurityGroups() []SecurityGroup`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *RebuiltInstance) GetSecurityGroupsOk() (*[]SecurityGroup, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *RebuiltInstance) SetSecurityGroups(v []SecurityGroup)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *RebuiltInstance) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### SetSecurityGroupsNil

`func (o *RebuiltInstance) SetSecurityGroupsNil(b bool)`

 SetSecurityGroupsNil sets the value for SecurityGroups to be an explicit nil

### UnsetSecurityGroups
`func (o *RebuiltInstance) UnsetSecurityGroups()`

UnsetSecurityGroups ensures that no value is present for SecurityGroups, not even an explicit nil
### GetTaskState

`func (o *RebuiltInstance) GetTaskState() string`

GetTaskState returns the TaskState field if non-nil, zero value otherwise.

### GetTaskStateOk

`func (o *RebuiltInstance) GetTaskStateOk() (*string, bool)`

GetTaskStateOk returns a tuple with the TaskState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskState

`func (o *RebuiltInstance) SetTaskState(v string)`

SetTaskState sets TaskState field to given value.

### HasTaskState

`func (o *RebuiltInstance) HasTaskState() bool`

HasTaskState returns a boolean if a field has been set.

### SetTaskStateNil

`func (o *RebuiltInstance) SetTaskStateNil(b bool)`

 SetTaskStateNil sets the value for TaskState to be an explicit nil

### UnsetTaskState
`func (o *RebuiltInstance) UnsetTaskState()`

UnsetTaskState ensures that no value is present for TaskState, not even an explicit nil
### GetVmState

`func (o *RebuiltInstance) GetVmState() string`

GetVmState returns the VmState field if non-nil, zero value otherwise.

### GetVmStateOk

`func (o *RebuiltInstance) GetVmStateOk() (*string, bool)`

GetVmStateOk returns a tuple with the VmState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmState

`func (o *RebuiltInstance) SetVmState(v string)`

SetVmState sets VmState field to given value.

### HasVmState

`func (o *RebuiltInstance) HasVmState() bool`

HasVmState returns a boolean if a field has been set.

### SetVmStateNil

`func (o *RebuiltInstance) SetVmStateNil(b bool)`

 SetVmStateNil sets the value for VmState to be an explicit nil

### UnsetVmState
`func (o *RebuiltInstance) UnsetVmState()`

UnsetVmState ensures that no value is present for VmState, not even an explicit nil
### GetPowerState

`func (o *RebuiltInstance) GetPowerState() PowerState`

GetPowerState returns the PowerState field if non-nil, zero value otherwise.

### GetPowerStateOk

`func (o *RebuiltInstance) GetPowerStateOk() (*PowerState, bool)`

GetPowerStateOk returns a tuple with the PowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerState

`func (o *RebuiltInstance) SetPowerState(v PowerState)`

SetPowerState sets PowerState field to given value.

### HasPowerState

`func (o *RebuiltInstance) HasPowerState() bool`

HasPowerState returns a boolean if a field has been set.

### SetPowerStateNil

`func (o *RebuiltInstance) SetPowerStateNil(b bool)`

 SetPowerStateNil sets the value for PowerState to be an explicit nil

### UnsetPowerState
`func (o *RebuiltInstance) UnsetPowerState()`

UnsetPowerState ensures that no value is present for PowerState, not even an explicit nil
### GetAttachedVolumes

`func (o *RebuiltInstance) GetAttachedVolumes() []RebuildInstanceAttachedVolume`

GetAttachedVolumes returns the AttachedVolumes field if non-nil, zero value otherwise.

### GetAttachedVolumesOk

`func (o *RebuiltInstance) GetAttachedVolumesOk() (*[]RebuildInstanceAttachedVolume, bool)`

GetAttachedVolumesOk returns a tuple with the AttachedVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedVolumes

`func (o *RebuiltInstance) SetAttachedVolumes(v []RebuildInstanceAttachedVolume)`

SetAttachedVolumes sets AttachedVolumes field to given value.

### HasAttachedVolumes

`func (o *RebuiltInstance) HasAttachedVolumes() bool`

HasAttachedVolumes returns a boolean if a field has been set.

### SetAttachedVolumesNil

`func (o *RebuiltInstance) SetAttachedVolumesNil(b bool)`

 SetAttachedVolumesNil sets the value for AttachedVolumes to be an explicit nil

### UnsetAttachedVolumes
`func (o *RebuiltInstance) UnsetAttachedVolumes()`

UnsetAttachedVolumes ensures that no value is present for AttachedVolumes, not even an explicit nil
### GetDescription

`func (o *RebuiltInstance) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RebuiltInstance) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RebuiltInstance) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RebuiltInstance) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *RebuiltInstance) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *RebuiltInstance) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetHostname

`func (o *RebuiltInstance) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *RebuiltInstance) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *RebuiltInstance) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *RebuiltInstance) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### SetHostnameNil

`func (o *RebuiltInstance) SetHostnameNil(b bool)`

 SetHostnameNil sets the value for Hostname to be an explicit nil

### UnsetHostname
`func (o *RebuiltInstance) UnsetHostname()`

UnsetHostname ensures that no value is present for Hostname, not even an explicit nil
### GetAddresses

`func (o *RebuiltInstance) GetAddresses() []InstanceAddress`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *RebuiltInstance) GetAddressesOk() (*[]InstanceAddress, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *RebuiltInstance) SetAddresses(v []InstanceAddress)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *RebuiltInstance) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### SetAddressesNil

`func (o *RebuiltInstance) SetAddressesNil(b bool)`

 SetAddressesNil sets the value for Addresses to be an explicit nil

### UnsetAddresses
`func (o *RebuiltInstance) UnsetAddresses()`

UnsetAddresses ensures that no value is present for Addresses, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


