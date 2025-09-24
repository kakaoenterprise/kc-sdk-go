# BcsInstanceV1ApiListInstancesModelInstanceModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 인스턴스의 ID | 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Flavor** | Pointer to [**NullableBcsInstanceV1ApiListInstancesModelInstanceFlavorModel**](BcsInstanceV1ApiListInstancesModelInstanceFlavorModel.md) |  | [optional] 
**Addresses** | Pointer to [**[]BcsInstanceV1ApiListInstancesModelInstanceAddressModel**](BcsInstanceV1ApiListInstancesModelInstanceAddressModel.md) |  | [optional] 
**IsHyperThreading** | Pointer to **NullableBool** |  | [optional] 
**IsHadoop** | Pointer to **NullableBool** |  | [optional] 
**IsK8se** | Pointer to **NullableBool** |  | [optional] 
**Image** | Pointer to [**NullableBcsInstanceV1ApiListInstancesModelInstanceImageModel**](BcsInstanceV1ApiListInstancesModelInstanceImageModel.md) |  | [optional] 
**VmState** | Pointer to **NullableString** |  | [optional] 
**TaskState** | Pointer to **NullableString** |  | [optional] 
**PowerState** | Pointer to [**NullablePowerState**](PowerState.md) |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**UserId** | Pointer to **NullableString** |  | [optional] 
**ProjectId** | Pointer to **NullableString** |  | [optional] 
**KeyName** | Pointer to **NullableString** |  | [optional] 
**Hostname** | Pointer to **NullableString** |  | [optional] 
**AvailabilityZone** | Pointer to [**NullableAvailabilityZone**](AvailabilityZone.md) |  | [optional] 
**AttachedVolumes** | Pointer to [**[]BcsInstanceV1ApiListInstancesModelInstanceAttachedVolumeModel**](BcsInstanceV1ApiListInstancesModelInstanceAttachedVolumeModel.md) |  | [optional] 
**AttachedVolumeCount** | Pointer to **NullableInt64** |  | [optional] 
**SecurityGroups** | Pointer to [**[]BcsInstanceV1ApiListInstancesModelInstanceSecurityGroupModel**](BcsInstanceV1ApiListInstancesModelInstanceSecurityGroupModel.md) |  | [optional] 
**SecurityGroupCount** | Pointer to **NullableInt64** |  | [optional] 
**InstanceType** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewBcsInstanceV1ApiListInstancesModelInstanceModel

`func NewBcsInstanceV1ApiListInstancesModelInstanceModel(id string, ) *BcsInstanceV1ApiListInstancesModelInstanceModel`

NewBcsInstanceV1ApiListInstancesModelInstanceModel instantiates a new BcsInstanceV1ApiListInstancesModelInstanceModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBcsInstanceV1ApiListInstancesModelInstanceModelWithDefaults

`func NewBcsInstanceV1ApiListInstancesModelInstanceModelWithDefaults() *BcsInstanceV1ApiListInstancesModelInstanceModel`

NewBcsInstanceV1ApiListInstancesModelInstanceModelWithDefaults instantiates a new BcsInstanceV1ApiListInstancesModelInstanceModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetMetadata

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetFlavor

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) GetFlavor() BcsInstanceV1ApiListInstancesModelInstanceFlavorModel`

GetFlavor returns the Flavor field if non-nil, zero value otherwise.

### GetFlavorOk

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) GetFlavorOk() (*BcsInstanceV1ApiListInstancesModelInstanceFlavorModel, bool)`

GetFlavorOk returns a tuple with the Flavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavor

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) SetFlavor(v BcsInstanceV1ApiListInstancesModelInstanceFlavorModel)`

SetFlavor sets Flavor field to given value.

### HasFlavor

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) HasFlavor() bool`

HasFlavor returns a boolean if a field has been set.

### SetFlavorNil

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) SetFlavorNil(b bool)`

 SetFlavorNil sets the value for Flavor to be an explicit nil

### UnsetFlavor
`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) UnsetFlavor()`

UnsetFlavor ensures that no value is present for Flavor, not even an explicit nil
### GetAddresses

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) GetAddresses() []BcsInstanceV1ApiListInstancesModelInstanceAddressModel`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) GetAddressesOk() (*[]BcsInstanceV1ApiListInstancesModelInstanceAddressModel, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) SetAddresses(v []BcsInstanceV1ApiListInstancesModelInstanceAddressModel)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### SetAddressesNil

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) SetAddressesNil(b bool)`

 SetAddressesNil sets the value for Addresses to be an explicit nil

### UnsetAddresses
`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) UnsetAddresses()`

UnsetAddresses ensures that no value is present for Addresses, not even an explicit nil
### GetIsHyperThreading

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) GetIsHyperThreading() bool`

GetIsHyperThreading returns the IsHyperThreading field if non-nil, zero value otherwise.

### GetIsHyperThreadingOk

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) GetIsHyperThreadingOk() (*bool, bool)`

GetIsHyperThreadingOk returns a tuple with the IsHyperThreading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHyperThreading

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) SetIsHyperThreading(v bool)`

SetIsHyperThreading sets IsHyperThreading field to given value.

### HasIsHyperThreading

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) HasIsHyperThreading() bool`

HasIsHyperThreading returns a boolean if a field has been set.

### SetIsHyperThreadingNil

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) SetIsHyperThreadingNil(b bool)`

 SetIsHyperThreadingNil sets the value for IsHyperThreading to be an explicit nil

### UnsetIsHyperThreading
`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) UnsetIsHyperThreading()`

UnsetIsHyperThreading ensures that no value is present for IsHyperThreading, not even an explicit nil
### GetIsHadoop

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) GetIsHadoop() bool`

GetIsHadoop returns the IsHadoop field if non-nil, zero value otherwise.

### GetIsHadoopOk

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) GetIsHadoopOk() (*bool, bool)`

GetIsHadoopOk returns a tuple with the IsHadoop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHadoop

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) SetIsHadoop(v bool)`

SetIsHadoop sets IsHadoop field to given value.

### HasIsHadoop

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) HasIsHadoop() bool`

HasIsHadoop returns a boolean if a field has been set.

### SetIsHadoopNil

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) SetIsHadoopNil(b bool)`

 SetIsHadoopNil sets the value for IsHadoop to be an explicit nil

### UnsetIsHadoop
`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) UnsetIsHadoop()`

UnsetIsHadoop ensures that no value is present for IsHadoop, not even an explicit nil
### GetIsK8se

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) GetIsK8se() bool`

GetIsK8se returns the IsK8se field if non-nil, zero value otherwise.

### GetIsK8seOk

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) GetIsK8seOk() (*bool, bool)`

GetIsK8seOk returns a tuple with the IsK8se field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsK8se

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) SetIsK8se(v bool)`

SetIsK8se sets IsK8se field to given value.

### HasIsK8se

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) HasIsK8se() bool`

HasIsK8se returns a boolean if a field has been set.

### SetIsK8seNil

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) SetIsK8seNil(b bool)`

 SetIsK8seNil sets the value for IsK8se to be an explicit nil

### UnsetIsK8se
`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) UnsetIsK8se()`

UnsetIsK8se ensures that no value is present for IsK8se, not even an explicit nil
### GetImage

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) GetImage() BcsInstanceV1ApiListInstancesModelInstanceImageModel`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) GetImageOk() (*BcsInstanceV1ApiListInstancesModelInstanceImageModel, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) SetImage(v BcsInstanceV1ApiListInstancesModelInstanceImageModel)`

SetImage sets Image field to given value.

### HasImage

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) HasImage() bool`

HasImage returns a boolean if a field has been set.

### SetImageNil

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetVmState

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) GetVmState() string`

GetVmState returns the VmState field if non-nil, zero value otherwise.

### GetVmStateOk

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) GetVmStateOk() (*string, bool)`

GetVmStateOk returns a tuple with the VmState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmState

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) SetVmState(v string)`

SetVmState sets VmState field to given value.

### HasVmState

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) HasVmState() bool`

HasVmState returns a boolean if a field has been set.

### SetVmStateNil

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) SetVmStateNil(b bool)`

 SetVmStateNil sets the value for VmState to be an explicit nil

### UnsetVmState
`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) UnsetVmState()`

UnsetVmState ensures that no value is present for VmState, not even an explicit nil
### GetTaskState

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) GetTaskState() string`

GetTaskState returns the TaskState field if non-nil, zero value otherwise.

### GetTaskStateOk

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) GetTaskStateOk() (*string, bool)`

GetTaskStateOk returns a tuple with the TaskState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskState

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) SetTaskState(v string)`

SetTaskState sets TaskState field to given value.

### HasTaskState

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) HasTaskState() bool`

HasTaskState returns a boolean if a field has been set.

### SetTaskStateNil

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) SetTaskStateNil(b bool)`

 SetTaskStateNil sets the value for TaskState to be an explicit nil

### UnsetTaskState
`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) UnsetTaskState()`

UnsetTaskState ensures that no value is present for TaskState, not even an explicit nil
### GetPowerState

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) GetPowerState() PowerState`

GetPowerState returns the PowerState field if non-nil, zero value otherwise.

### GetPowerStateOk

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) GetPowerStateOk() (*PowerState, bool)`

GetPowerStateOk returns a tuple with the PowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerState

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) SetPowerState(v PowerState)`

SetPowerState sets PowerState field to given value.

### HasPowerState

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) HasPowerState() bool`

HasPowerState returns a boolean if a field has been set.

### SetPowerStateNil

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) SetPowerStateNil(b bool)`

 SetPowerStateNil sets the value for PowerState to be an explicit nil

### UnsetPowerState
`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) UnsetPowerState()`

UnsetPowerState ensures that no value is present for PowerState, not even an explicit nil
### GetStatus

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUserId

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetProjectId

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetKeyName

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.

### HasKeyName

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) HasKeyName() bool`

HasKeyName returns a boolean if a field has been set.

### SetKeyNameNil

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) SetKeyNameNil(b bool)`

 SetKeyNameNil sets the value for KeyName to be an explicit nil

### UnsetKeyName
`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) UnsetKeyName()`

UnsetKeyName ensures that no value is present for KeyName, not even an explicit nil
### GetHostname

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### SetHostnameNil

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) SetHostnameNil(b bool)`

 SetHostnameNil sets the value for Hostname to be an explicit nil

### UnsetHostname
`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) UnsetHostname()`

UnsetHostname ensures that no value is present for Hostname, not even an explicit nil
### GetAvailabilityZone

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) GetAvailabilityZone() AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) GetAvailabilityZoneOk() (*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) SetAvailabilityZone(v AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### SetAvailabilityZoneNil

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) SetAvailabilityZoneNil(b bool)`

 SetAvailabilityZoneNil sets the value for AvailabilityZone to be an explicit nil

### UnsetAvailabilityZone
`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) UnsetAvailabilityZone()`

UnsetAvailabilityZone ensures that no value is present for AvailabilityZone, not even an explicit nil
### GetAttachedVolumes

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) GetAttachedVolumes() []BcsInstanceV1ApiListInstancesModelInstanceAttachedVolumeModel`

GetAttachedVolumes returns the AttachedVolumes field if non-nil, zero value otherwise.

### GetAttachedVolumesOk

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) GetAttachedVolumesOk() (*[]BcsInstanceV1ApiListInstancesModelInstanceAttachedVolumeModel, bool)`

GetAttachedVolumesOk returns a tuple with the AttachedVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedVolumes

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) SetAttachedVolumes(v []BcsInstanceV1ApiListInstancesModelInstanceAttachedVolumeModel)`

SetAttachedVolumes sets AttachedVolumes field to given value.

### HasAttachedVolumes

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) HasAttachedVolumes() bool`

HasAttachedVolumes returns a boolean if a field has been set.

### SetAttachedVolumesNil

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) SetAttachedVolumesNil(b bool)`

 SetAttachedVolumesNil sets the value for AttachedVolumes to be an explicit nil

### UnsetAttachedVolumes
`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) UnsetAttachedVolumes()`

UnsetAttachedVolumes ensures that no value is present for AttachedVolumes, not even an explicit nil
### GetAttachedVolumeCount

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) GetAttachedVolumeCount() int64`

GetAttachedVolumeCount returns the AttachedVolumeCount field if non-nil, zero value otherwise.

### GetAttachedVolumeCountOk

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) GetAttachedVolumeCountOk() (*int64, bool)`

GetAttachedVolumeCountOk returns a tuple with the AttachedVolumeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedVolumeCount

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) SetAttachedVolumeCount(v int64)`

SetAttachedVolumeCount sets AttachedVolumeCount field to given value.

### HasAttachedVolumeCount

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) HasAttachedVolumeCount() bool`

HasAttachedVolumeCount returns a boolean if a field has been set.

### SetAttachedVolumeCountNil

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) SetAttachedVolumeCountNil(b bool)`

 SetAttachedVolumeCountNil sets the value for AttachedVolumeCount to be an explicit nil

### UnsetAttachedVolumeCount
`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) UnsetAttachedVolumeCount()`

UnsetAttachedVolumeCount ensures that no value is present for AttachedVolumeCount, not even an explicit nil
### GetSecurityGroups

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) GetSecurityGroups() []BcsInstanceV1ApiListInstancesModelInstanceSecurityGroupModel`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) GetSecurityGroupsOk() (*[]BcsInstanceV1ApiListInstancesModelInstanceSecurityGroupModel, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) SetSecurityGroups(v []BcsInstanceV1ApiListInstancesModelInstanceSecurityGroupModel)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### SetSecurityGroupsNil

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) SetSecurityGroupsNil(b bool)`

 SetSecurityGroupsNil sets the value for SecurityGroups to be an explicit nil

### UnsetSecurityGroups
`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) UnsetSecurityGroups()`

UnsetSecurityGroups ensures that no value is present for SecurityGroups, not even an explicit nil
### GetSecurityGroupCount

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) GetSecurityGroupCount() int64`

GetSecurityGroupCount returns the SecurityGroupCount field if non-nil, zero value otherwise.

### GetSecurityGroupCountOk

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) GetSecurityGroupCountOk() (*int64, bool)`

GetSecurityGroupCountOk returns a tuple with the SecurityGroupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupCount

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) SetSecurityGroupCount(v int64)`

SetSecurityGroupCount sets SecurityGroupCount field to given value.

### HasSecurityGroupCount

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) HasSecurityGroupCount() bool`

HasSecurityGroupCount returns a boolean if a field has been set.

### SetSecurityGroupCountNil

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) SetSecurityGroupCountNil(b bool)`

 SetSecurityGroupCountNil sets the value for SecurityGroupCount to be an explicit nil

### UnsetSecurityGroupCount
`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) UnsetSecurityGroupCount()`

UnsetSecurityGroupCount ensures that no value is present for SecurityGroupCount, not even an explicit nil
### GetInstanceType

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### SetInstanceTypeNil

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) SetInstanceTypeNil(b bool)`

 SetInstanceTypeNil sets the value for InstanceType to be an explicit nil

### UnsetInstanceType
`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) UnsetInstanceType()`

UnsetInstanceType ensures that no value is present for InstanceType, not even an explicit nil
### GetCreatedAt

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *BcsInstanceV1ApiListInstancesModelInstanceModel) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


