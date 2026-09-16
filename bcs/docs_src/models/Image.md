# Image

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | 이미지의 고유 ID | [optional] 
**Name** | Pointer to **NullableString** | 이미지 이름 | [optional] 
**Description** | Pointer to **NullableString** | 이미지에 대한 설명 | [optional] 
**Owner** | Pointer to **NullableString** | 이미지 소유자(이미지를 소유한 프로젝트 ID) | [optional] 
**IsWindows** | Pointer to **NullableBool** | 이미지가 Windows 운영체제인지 여부 | [optional] 
**Size** | Pointer to **NullableInt64** | 이미지 크기 (bytes 단위) | [optional] 
**Status** | Pointer to **NullableString** | 이미지 상태 (예: active, queued, deleted 등) | [optional] 
**ImageType** | Pointer to **NullableString** | 이미지 유형 | [optional] 
**CreatedAt** | Pointer to **NullableTime** | 리소스가 생성된 시간 - ISO 8601 형식 - UTC 기준 | [optional] 
**UpdatedAt** | Pointer to **NullableTime** | 리소스가 마지막으로 수정된 시간 - ISO 8601 형식 - UTC 기준 | [optional] 
**DiskFormat** | Pointer to **NullableString** | 이미지의 디스크 포맷 (예: qcow2, raw 등) | [optional] 
**InstanceType** | Pointer to [**NullableInstanceType**](InstanceType.md) | 인스턴스 유형 | [optional] 
**MemberStatus** | Pointer to **NullableString** | 이미지가 공유된 프로젝트에서의 수락 상태 | [optional] 
**MinDisk** | Pointer to **NullableInt32** | 최소 디스크 요구 용량 (GB) | [optional] 
**MinMemory** | Pointer to **NullableInt32** | 최소 메모리 요구 용량 (MB) | [optional] 
**OsAdmin** | Pointer to **NullableString** | 운영체제 관리자 계정 이름 (예: root, Administrator 등) | [optional] 
**OsArchitecture** | Pointer to **NullableString** | 운영체제 아키텍처 (예: x86_64, arm64) | [optional] 
**OsDistro** | Pointer to **NullableString** | 운영체제 배포판 | [optional] 
**OsType** | Pointer to **NullableString** | 운영체제 유형 | [optional] 

## Methods

### NewImage

`func NewImage() *Image`

NewImage instantiates a new Image object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageWithDefaults

`func NewImageWithDefaults() *Image`

NewImageWithDefaults instantiates a new Image object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Image) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Image) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Image) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Image) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Image) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Image) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *Image) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Image) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Image) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Image) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Image) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Image) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *Image) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Image) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Image) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Image) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Image) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Image) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetOwner

`func (o *Image) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Image) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Image) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Image) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *Image) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *Image) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetIsWindows

`func (o *Image) GetIsWindows() bool`

GetIsWindows returns the IsWindows field if non-nil, zero value otherwise.

### GetIsWindowsOk

`func (o *Image) GetIsWindowsOk() (*bool, bool)`

GetIsWindowsOk returns a tuple with the IsWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWindows

`func (o *Image) SetIsWindows(v bool)`

SetIsWindows sets IsWindows field to given value.

### HasIsWindows

`func (o *Image) HasIsWindows() bool`

HasIsWindows returns a boolean if a field has been set.

### SetIsWindowsNil

`func (o *Image) SetIsWindowsNil(b bool)`

 SetIsWindowsNil sets the value for IsWindows to be an explicit nil

### UnsetIsWindows
`func (o *Image) UnsetIsWindows()`

UnsetIsWindows ensures that no value is present for IsWindows, not even an explicit nil
### GetSize

`func (o *Image) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Image) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Image) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *Image) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *Image) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *Image) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetStatus

`func (o *Image) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Image) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Image) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Image) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *Image) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *Image) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetImageType

`func (o *Image) GetImageType() string`

GetImageType returns the ImageType field if non-nil, zero value otherwise.

### GetImageTypeOk

`func (o *Image) GetImageTypeOk() (*string, bool)`

GetImageTypeOk returns a tuple with the ImageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageType

`func (o *Image) SetImageType(v string)`

SetImageType sets ImageType field to given value.

### HasImageType

`func (o *Image) HasImageType() bool`

HasImageType returns a boolean if a field has been set.

### SetImageTypeNil

`func (o *Image) SetImageTypeNil(b bool)`

 SetImageTypeNil sets the value for ImageType to be an explicit nil

### UnsetImageType
`func (o *Image) UnsetImageType()`

UnsetImageType ensures that no value is present for ImageType, not even an explicit nil
### GetCreatedAt

`func (o *Image) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Image) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Image) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Image) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *Image) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *Image) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *Image) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Image) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Image) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Image) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *Image) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *Image) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetDiskFormat

`func (o *Image) GetDiskFormat() string`

GetDiskFormat returns the DiskFormat field if non-nil, zero value otherwise.

### GetDiskFormatOk

`func (o *Image) GetDiskFormatOk() (*string, bool)`

GetDiskFormatOk returns a tuple with the DiskFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskFormat

`func (o *Image) SetDiskFormat(v string)`

SetDiskFormat sets DiskFormat field to given value.

### HasDiskFormat

`func (o *Image) HasDiskFormat() bool`

HasDiskFormat returns a boolean if a field has been set.

### SetDiskFormatNil

`func (o *Image) SetDiskFormatNil(b bool)`

 SetDiskFormatNil sets the value for DiskFormat to be an explicit nil

### UnsetDiskFormat
`func (o *Image) UnsetDiskFormat()`

UnsetDiskFormat ensures that no value is present for DiskFormat, not even an explicit nil
### GetInstanceType

`func (o *Image) GetInstanceType() InstanceType`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *Image) GetInstanceTypeOk() (*InstanceType, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *Image) SetInstanceType(v InstanceType)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *Image) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### SetInstanceTypeNil

`func (o *Image) SetInstanceTypeNil(b bool)`

 SetInstanceTypeNil sets the value for InstanceType to be an explicit nil

### UnsetInstanceType
`func (o *Image) UnsetInstanceType()`

UnsetInstanceType ensures that no value is present for InstanceType, not even an explicit nil
### GetMemberStatus

`func (o *Image) GetMemberStatus() string`

GetMemberStatus returns the MemberStatus field if non-nil, zero value otherwise.

### GetMemberStatusOk

`func (o *Image) GetMemberStatusOk() (*string, bool)`

GetMemberStatusOk returns a tuple with the MemberStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberStatus

`func (o *Image) SetMemberStatus(v string)`

SetMemberStatus sets MemberStatus field to given value.

### HasMemberStatus

`func (o *Image) HasMemberStatus() bool`

HasMemberStatus returns a boolean if a field has been set.

### SetMemberStatusNil

`func (o *Image) SetMemberStatusNil(b bool)`

 SetMemberStatusNil sets the value for MemberStatus to be an explicit nil

### UnsetMemberStatus
`func (o *Image) UnsetMemberStatus()`

UnsetMemberStatus ensures that no value is present for MemberStatus, not even an explicit nil
### GetMinDisk

`func (o *Image) GetMinDisk() int32`

GetMinDisk returns the MinDisk field if non-nil, zero value otherwise.

### GetMinDiskOk

`func (o *Image) GetMinDiskOk() (*int32, bool)`

GetMinDiskOk returns a tuple with the MinDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDisk

`func (o *Image) SetMinDisk(v int32)`

SetMinDisk sets MinDisk field to given value.

### HasMinDisk

`func (o *Image) HasMinDisk() bool`

HasMinDisk returns a boolean if a field has been set.

### SetMinDiskNil

`func (o *Image) SetMinDiskNil(b bool)`

 SetMinDiskNil sets the value for MinDisk to be an explicit nil

### UnsetMinDisk
`func (o *Image) UnsetMinDisk()`

UnsetMinDisk ensures that no value is present for MinDisk, not even an explicit nil
### GetMinMemory

`func (o *Image) GetMinMemory() int32`

GetMinMemory returns the MinMemory field if non-nil, zero value otherwise.

### GetMinMemoryOk

`func (o *Image) GetMinMemoryOk() (*int32, bool)`

GetMinMemoryOk returns a tuple with the MinMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinMemory

`func (o *Image) SetMinMemory(v int32)`

SetMinMemory sets MinMemory field to given value.

### HasMinMemory

`func (o *Image) HasMinMemory() bool`

HasMinMemory returns a boolean if a field has been set.

### SetMinMemoryNil

`func (o *Image) SetMinMemoryNil(b bool)`

 SetMinMemoryNil sets the value for MinMemory to be an explicit nil

### UnsetMinMemory
`func (o *Image) UnsetMinMemory()`

UnsetMinMemory ensures that no value is present for MinMemory, not even an explicit nil
### GetOsAdmin

`func (o *Image) GetOsAdmin() string`

GetOsAdmin returns the OsAdmin field if non-nil, zero value otherwise.

### GetOsAdminOk

`func (o *Image) GetOsAdminOk() (*string, bool)`

GetOsAdminOk returns a tuple with the OsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsAdmin

`func (o *Image) SetOsAdmin(v string)`

SetOsAdmin sets OsAdmin field to given value.

### HasOsAdmin

`func (o *Image) HasOsAdmin() bool`

HasOsAdmin returns a boolean if a field has been set.

### SetOsAdminNil

`func (o *Image) SetOsAdminNil(b bool)`

 SetOsAdminNil sets the value for OsAdmin to be an explicit nil

### UnsetOsAdmin
`func (o *Image) UnsetOsAdmin()`

UnsetOsAdmin ensures that no value is present for OsAdmin, not even an explicit nil
### GetOsArchitecture

`func (o *Image) GetOsArchitecture() string`

GetOsArchitecture returns the OsArchitecture field if non-nil, zero value otherwise.

### GetOsArchitectureOk

`func (o *Image) GetOsArchitectureOk() (*string, bool)`

GetOsArchitectureOk returns a tuple with the OsArchitecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsArchitecture

`func (o *Image) SetOsArchitecture(v string)`

SetOsArchitecture sets OsArchitecture field to given value.

### HasOsArchitecture

`func (o *Image) HasOsArchitecture() bool`

HasOsArchitecture returns a boolean if a field has been set.

### SetOsArchitectureNil

`func (o *Image) SetOsArchitectureNil(b bool)`

 SetOsArchitectureNil sets the value for OsArchitecture to be an explicit nil

### UnsetOsArchitecture
`func (o *Image) UnsetOsArchitecture()`

UnsetOsArchitecture ensures that no value is present for OsArchitecture, not even an explicit nil
### GetOsDistro

`func (o *Image) GetOsDistro() string`

GetOsDistro returns the OsDistro field if non-nil, zero value otherwise.

### GetOsDistroOk

`func (o *Image) GetOsDistroOk() (*string, bool)`

GetOsDistroOk returns a tuple with the OsDistro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsDistro

`func (o *Image) SetOsDistro(v string)`

SetOsDistro sets OsDistro field to given value.

### HasOsDistro

`func (o *Image) HasOsDistro() bool`

HasOsDistro returns a boolean if a field has been set.

### SetOsDistroNil

`func (o *Image) SetOsDistroNil(b bool)`

 SetOsDistroNil sets the value for OsDistro to be an explicit nil

### UnsetOsDistro
`func (o *Image) UnsetOsDistro()`

UnsetOsDistro ensures that no value is present for OsDistro, not even an explicit nil
### GetOsType

`func (o *Image) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *Image) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *Image) SetOsType(v string)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *Image) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### SetOsTypeNil

`func (o *Image) SetOsTypeNil(b bool)`

 SetOsTypeNil sets the value for OsType to be an explicit nil

### UnsetOsType
`func (o *Image) UnsetOsType()`

UnsetOsType ensures that no value is present for OsType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


