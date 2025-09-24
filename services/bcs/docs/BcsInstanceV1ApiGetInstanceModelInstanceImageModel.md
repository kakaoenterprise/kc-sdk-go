# BcsInstanceV1ApiGetInstanceModelInstanceImageModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 이미지의 고유 ID | 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Owner** | Pointer to **NullableString** |  | [optional] 
**IsWindows** | Pointer to **NullableBool** |  | [optional] 
**Size** | Pointer to **NullableInt64** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**ImageType** | Pointer to **NullableString** |  | [optional] 
**DiskFormat** | Pointer to **NullableString** |  | [optional] 
**InstanceType** | Pointer to [**NullableInstanceType**](InstanceType.md) |  | [optional] 
**MemberStatus** | Pointer to **NullableString** |  | [optional] 
**MinDisk** | Pointer to **NullableInt32** |  | [optional] 
**MinMemory** | Pointer to **NullableInt32** |  | [optional] 
**OsAdmin** | Pointer to **NullableString** |  | [optional] 
**OsDistro** | Pointer to **NullableString** |  | [optional] 
**OsType** | Pointer to **NullableString** |  | [optional] 
**OsArchitecture** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewBcsInstanceV1ApiGetInstanceModelInstanceImageModel

`func NewBcsInstanceV1ApiGetInstanceModelInstanceImageModel(id string, ) *BcsInstanceV1ApiGetInstanceModelInstanceImageModel`

NewBcsInstanceV1ApiGetInstanceModelInstanceImageModel instantiates a new BcsInstanceV1ApiGetInstanceModelInstanceImageModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBcsInstanceV1ApiGetInstanceModelInstanceImageModelWithDefaults

`func NewBcsInstanceV1ApiGetInstanceModelInstanceImageModelWithDefaults() *BcsInstanceV1ApiGetInstanceModelInstanceImageModel`

NewBcsInstanceV1ApiGetInstanceModelInstanceImageModelWithDefaults instantiates a new BcsInstanceV1ApiGetInstanceModelInstanceImageModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetOwner

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetIsWindows

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) GetIsWindows() bool`

GetIsWindows returns the IsWindows field if non-nil, zero value otherwise.

### GetIsWindowsOk

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) GetIsWindowsOk() (*bool, bool)`

GetIsWindowsOk returns a tuple with the IsWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWindows

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) SetIsWindows(v bool)`

SetIsWindows sets IsWindows field to given value.

### HasIsWindows

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) HasIsWindows() bool`

HasIsWindows returns a boolean if a field has been set.

### SetIsWindowsNil

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) SetIsWindowsNil(b bool)`

 SetIsWindowsNil sets the value for IsWindows to be an explicit nil

### UnsetIsWindows
`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) UnsetIsWindows()`

UnsetIsWindows ensures that no value is present for IsWindows, not even an explicit nil
### GetSize

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetStatus

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetImageType

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) GetImageType() string`

GetImageType returns the ImageType field if non-nil, zero value otherwise.

### GetImageTypeOk

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) GetImageTypeOk() (*string, bool)`

GetImageTypeOk returns a tuple with the ImageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageType

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) SetImageType(v string)`

SetImageType sets ImageType field to given value.

### HasImageType

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) HasImageType() bool`

HasImageType returns a boolean if a field has been set.

### SetImageTypeNil

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) SetImageTypeNil(b bool)`

 SetImageTypeNil sets the value for ImageType to be an explicit nil

### UnsetImageType
`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) UnsetImageType()`

UnsetImageType ensures that no value is present for ImageType, not even an explicit nil
### GetDiskFormat

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) GetDiskFormat() string`

GetDiskFormat returns the DiskFormat field if non-nil, zero value otherwise.

### GetDiskFormatOk

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) GetDiskFormatOk() (*string, bool)`

GetDiskFormatOk returns a tuple with the DiskFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskFormat

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) SetDiskFormat(v string)`

SetDiskFormat sets DiskFormat field to given value.

### HasDiskFormat

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) HasDiskFormat() bool`

HasDiskFormat returns a boolean if a field has been set.

### SetDiskFormatNil

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) SetDiskFormatNil(b bool)`

 SetDiskFormatNil sets the value for DiskFormat to be an explicit nil

### UnsetDiskFormat
`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) UnsetDiskFormat()`

UnsetDiskFormat ensures that no value is present for DiskFormat, not even an explicit nil
### GetInstanceType

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) GetInstanceType() InstanceType`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) GetInstanceTypeOk() (*InstanceType, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) SetInstanceType(v InstanceType)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### SetInstanceTypeNil

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) SetInstanceTypeNil(b bool)`

 SetInstanceTypeNil sets the value for InstanceType to be an explicit nil

### UnsetInstanceType
`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) UnsetInstanceType()`

UnsetInstanceType ensures that no value is present for InstanceType, not even an explicit nil
### GetMemberStatus

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) GetMemberStatus() string`

GetMemberStatus returns the MemberStatus field if non-nil, zero value otherwise.

### GetMemberStatusOk

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) GetMemberStatusOk() (*string, bool)`

GetMemberStatusOk returns a tuple with the MemberStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberStatus

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) SetMemberStatus(v string)`

SetMemberStatus sets MemberStatus field to given value.

### HasMemberStatus

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) HasMemberStatus() bool`

HasMemberStatus returns a boolean if a field has been set.

### SetMemberStatusNil

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) SetMemberStatusNil(b bool)`

 SetMemberStatusNil sets the value for MemberStatus to be an explicit nil

### UnsetMemberStatus
`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) UnsetMemberStatus()`

UnsetMemberStatus ensures that no value is present for MemberStatus, not even an explicit nil
### GetMinDisk

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) GetMinDisk() int32`

GetMinDisk returns the MinDisk field if non-nil, zero value otherwise.

### GetMinDiskOk

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) GetMinDiskOk() (*int32, bool)`

GetMinDiskOk returns a tuple with the MinDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDisk

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) SetMinDisk(v int32)`

SetMinDisk sets MinDisk field to given value.

### HasMinDisk

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) HasMinDisk() bool`

HasMinDisk returns a boolean if a field has been set.

### SetMinDiskNil

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) SetMinDiskNil(b bool)`

 SetMinDiskNil sets the value for MinDisk to be an explicit nil

### UnsetMinDisk
`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) UnsetMinDisk()`

UnsetMinDisk ensures that no value is present for MinDisk, not even an explicit nil
### GetMinMemory

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) GetMinMemory() int32`

GetMinMemory returns the MinMemory field if non-nil, zero value otherwise.

### GetMinMemoryOk

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) GetMinMemoryOk() (*int32, bool)`

GetMinMemoryOk returns a tuple with the MinMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinMemory

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) SetMinMemory(v int32)`

SetMinMemory sets MinMemory field to given value.

### HasMinMemory

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) HasMinMemory() bool`

HasMinMemory returns a boolean if a field has been set.

### SetMinMemoryNil

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) SetMinMemoryNil(b bool)`

 SetMinMemoryNil sets the value for MinMemory to be an explicit nil

### UnsetMinMemory
`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) UnsetMinMemory()`

UnsetMinMemory ensures that no value is present for MinMemory, not even an explicit nil
### GetOsAdmin

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) GetOsAdmin() string`

GetOsAdmin returns the OsAdmin field if non-nil, zero value otherwise.

### GetOsAdminOk

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) GetOsAdminOk() (*string, bool)`

GetOsAdminOk returns a tuple with the OsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsAdmin

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) SetOsAdmin(v string)`

SetOsAdmin sets OsAdmin field to given value.

### HasOsAdmin

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) HasOsAdmin() bool`

HasOsAdmin returns a boolean if a field has been set.

### SetOsAdminNil

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) SetOsAdminNil(b bool)`

 SetOsAdminNil sets the value for OsAdmin to be an explicit nil

### UnsetOsAdmin
`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) UnsetOsAdmin()`

UnsetOsAdmin ensures that no value is present for OsAdmin, not even an explicit nil
### GetOsDistro

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) GetOsDistro() string`

GetOsDistro returns the OsDistro field if non-nil, zero value otherwise.

### GetOsDistroOk

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) GetOsDistroOk() (*string, bool)`

GetOsDistroOk returns a tuple with the OsDistro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsDistro

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) SetOsDistro(v string)`

SetOsDistro sets OsDistro field to given value.

### HasOsDistro

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) HasOsDistro() bool`

HasOsDistro returns a boolean if a field has been set.

### SetOsDistroNil

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) SetOsDistroNil(b bool)`

 SetOsDistroNil sets the value for OsDistro to be an explicit nil

### UnsetOsDistro
`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) UnsetOsDistro()`

UnsetOsDistro ensures that no value is present for OsDistro, not even an explicit nil
### GetOsType

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) SetOsType(v string)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### SetOsTypeNil

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) SetOsTypeNil(b bool)`

 SetOsTypeNil sets the value for OsType to be an explicit nil

### UnsetOsType
`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) UnsetOsType()`

UnsetOsType ensures that no value is present for OsType, not even an explicit nil
### GetOsArchitecture

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) GetOsArchitecture() string`

GetOsArchitecture returns the OsArchitecture field if non-nil, zero value otherwise.

### GetOsArchitectureOk

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) GetOsArchitectureOk() (*string, bool)`

GetOsArchitectureOk returns a tuple with the OsArchitecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsArchitecture

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) SetOsArchitecture(v string)`

SetOsArchitecture sets OsArchitecture field to given value.

### HasOsArchitecture

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) HasOsArchitecture() bool`

HasOsArchitecture returns a boolean if a field has been set.

### SetOsArchitectureNil

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) SetOsArchitectureNil(b bool)`

 SetOsArchitectureNil sets the value for OsArchitecture to be an explicit nil

### UnsetOsArchitecture
`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) UnsetOsArchitecture()`

UnsetOsArchitecture ensures that no value is present for OsArchitecture, not even an explicit nil
### GetCreatedAt

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *BcsInstanceV1ApiGetInstanceModelInstanceImageModel) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


