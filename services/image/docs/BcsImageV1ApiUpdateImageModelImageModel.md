# BcsImageV1ApiUpdateImageModelImageModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 이미지의 고유 ID | 
**Name** | Pointer to **NullableString** |  | [optional] 
**Size** | Pointer to **NullableInt64** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**DiskFormat** | Pointer to **NullableString** |  | [optional] 
**ContainerFormat** | Pointer to **NullableString** |  | [optional] 
**Owner** | Pointer to **NullableString** |  | [optional] 
**MinDisk** | Pointer to **NullableInt32** |  | [optional] 
**MinRam** | Pointer to **NullableInt32** |  | [optional] 
**VirtualSize** | Pointer to **NullableInt64** |  | [optional] 
**IsShared** | Pointer to **NullableBool** |  | [optional] 
**Visibility** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**InstanceType** | Pointer to **NullableString** |  | [optional] 
**OsInfo** | Pointer to [**NullableBcsImageV1ApiUpdateImageModelOsInfoModel**](BcsImageV1ApiUpdateImageModelOsInfoModel.md) |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewBcsImageV1ApiUpdateImageModelImageModel

`func NewBcsImageV1ApiUpdateImageModelImageModel(id string, ) *BcsImageV1ApiUpdateImageModelImageModel`

NewBcsImageV1ApiUpdateImageModelImageModel instantiates a new BcsImageV1ApiUpdateImageModelImageModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBcsImageV1ApiUpdateImageModelImageModelWithDefaults

`func NewBcsImageV1ApiUpdateImageModelImageModelWithDefaults() *BcsImageV1ApiUpdateImageModelImageModel`

NewBcsImageV1ApiUpdateImageModelImageModelWithDefaults instantiates a new BcsImageV1ApiUpdateImageModelImageModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BcsImageV1ApiUpdateImageModelImageModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BcsImageV1ApiUpdateImageModelImageModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BcsImageV1ApiUpdateImageModelImageModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BcsImageV1ApiUpdateImageModelImageModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BcsImageV1ApiUpdateImageModelImageModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BcsImageV1ApiUpdateImageModelImageModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BcsImageV1ApiUpdateImageModelImageModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BcsImageV1ApiUpdateImageModelImageModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BcsImageV1ApiUpdateImageModelImageModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSize

`func (o *BcsImageV1ApiUpdateImageModelImageModel) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *BcsImageV1ApiUpdateImageModelImageModel) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *BcsImageV1ApiUpdateImageModelImageModel) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *BcsImageV1ApiUpdateImageModelImageModel) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *BcsImageV1ApiUpdateImageModelImageModel) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *BcsImageV1ApiUpdateImageModelImageModel) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetStatus

`func (o *BcsImageV1ApiUpdateImageModelImageModel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BcsImageV1ApiUpdateImageModelImageModel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BcsImageV1ApiUpdateImageModelImageModel) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BcsImageV1ApiUpdateImageModelImageModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *BcsImageV1ApiUpdateImageModelImageModel) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *BcsImageV1ApiUpdateImageModelImageModel) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetDiskFormat

`func (o *BcsImageV1ApiUpdateImageModelImageModel) GetDiskFormat() string`

GetDiskFormat returns the DiskFormat field if non-nil, zero value otherwise.

### GetDiskFormatOk

`func (o *BcsImageV1ApiUpdateImageModelImageModel) GetDiskFormatOk() (*string, bool)`

GetDiskFormatOk returns a tuple with the DiskFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskFormat

`func (o *BcsImageV1ApiUpdateImageModelImageModel) SetDiskFormat(v string)`

SetDiskFormat sets DiskFormat field to given value.

### HasDiskFormat

`func (o *BcsImageV1ApiUpdateImageModelImageModel) HasDiskFormat() bool`

HasDiskFormat returns a boolean if a field has been set.

### SetDiskFormatNil

`func (o *BcsImageV1ApiUpdateImageModelImageModel) SetDiskFormatNil(b bool)`

 SetDiskFormatNil sets the value for DiskFormat to be an explicit nil

### UnsetDiskFormat
`func (o *BcsImageV1ApiUpdateImageModelImageModel) UnsetDiskFormat()`

UnsetDiskFormat ensures that no value is present for DiskFormat, not even an explicit nil
### GetContainerFormat

`func (o *BcsImageV1ApiUpdateImageModelImageModel) GetContainerFormat() string`

GetContainerFormat returns the ContainerFormat field if non-nil, zero value otherwise.

### GetContainerFormatOk

`func (o *BcsImageV1ApiUpdateImageModelImageModel) GetContainerFormatOk() (*string, bool)`

GetContainerFormatOk returns a tuple with the ContainerFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerFormat

`func (o *BcsImageV1ApiUpdateImageModelImageModel) SetContainerFormat(v string)`

SetContainerFormat sets ContainerFormat field to given value.

### HasContainerFormat

`func (o *BcsImageV1ApiUpdateImageModelImageModel) HasContainerFormat() bool`

HasContainerFormat returns a boolean if a field has been set.

### SetContainerFormatNil

`func (o *BcsImageV1ApiUpdateImageModelImageModel) SetContainerFormatNil(b bool)`

 SetContainerFormatNil sets the value for ContainerFormat to be an explicit nil

### UnsetContainerFormat
`func (o *BcsImageV1ApiUpdateImageModelImageModel) UnsetContainerFormat()`

UnsetContainerFormat ensures that no value is present for ContainerFormat, not even an explicit nil
### GetOwner

`func (o *BcsImageV1ApiUpdateImageModelImageModel) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BcsImageV1ApiUpdateImageModelImageModel) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BcsImageV1ApiUpdateImageModelImageModel) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BcsImageV1ApiUpdateImageModelImageModel) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *BcsImageV1ApiUpdateImageModelImageModel) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *BcsImageV1ApiUpdateImageModelImageModel) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetMinDisk

`func (o *BcsImageV1ApiUpdateImageModelImageModel) GetMinDisk() int32`

GetMinDisk returns the MinDisk field if non-nil, zero value otherwise.

### GetMinDiskOk

`func (o *BcsImageV1ApiUpdateImageModelImageModel) GetMinDiskOk() (*int32, bool)`

GetMinDiskOk returns a tuple with the MinDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDisk

`func (o *BcsImageV1ApiUpdateImageModelImageModel) SetMinDisk(v int32)`

SetMinDisk sets MinDisk field to given value.

### HasMinDisk

`func (o *BcsImageV1ApiUpdateImageModelImageModel) HasMinDisk() bool`

HasMinDisk returns a boolean if a field has been set.

### SetMinDiskNil

`func (o *BcsImageV1ApiUpdateImageModelImageModel) SetMinDiskNil(b bool)`

 SetMinDiskNil sets the value for MinDisk to be an explicit nil

### UnsetMinDisk
`func (o *BcsImageV1ApiUpdateImageModelImageModel) UnsetMinDisk()`

UnsetMinDisk ensures that no value is present for MinDisk, not even an explicit nil
### GetMinRam

`func (o *BcsImageV1ApiUpdateImageModelImageModel) GetMinRam() int32`

GetMinRam returns the MinRam field if non-nil, zero value otherwise.

### GetMinRamOk

`func (o *BcsImageV1ApiUpdateImageModelImageModel) GetMinRamOk() (*int32, bool)`

GetMinRamOk returns a tuple with the MinRam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRam

`func (o *BcsImageV1ApiUpdateImageModelImageModel) SetMinRam(v int32)`

SetMinRam sets MinRam field to given value.

### HasMinRam

`func (o *BcsImageV1ApiUpdateImageModelImageModel) HasMinRam() bool`

HasMinRam returns a boolean if a field has been set.

### SetMinRamNil

`func (o *BcsImageV1ApiUpdateImageModelImageModel) SetMinRamNil(b bool)`

 SetMinRamNil sets the value for MinRam to be an explicit nil

### UnsetMinRam
`func (o *BcsImageV1ApiUpdateImageModelImageModel) UnsetMinRam()`

UnsetMinRam ensures that no value is present for MinRam, not even an explicit nil
### GetVirtualSize

`func (o *BcsImageV1ApiUpdateImageModelImageModel) GetVirtualSize() int64`

GetVirtualSize returns the VirtualSize field if non-nil, zero value otherwise.

### GetVirtualSizeOk

`func (o *BcsImageV1ApiUpdateImageModelImageModel) GetVirtualSizeOk() (*int64, bool)`

GetVirtualSizeOk returns a tuple with the VirtualSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualSize

`func (o *BcsImageV1ApiUpdateImageModelImageModel) SetVirtualSize(v int64)`

SetVirtualSize sets VirtualSize field to given value.

### HasVirtualSize

`func (o *BcsImageV1ApiUpdateImageModelImageModel) HasVirtualSize() bool`

HasVirtualSize returns a boolean if a field has been set.

### SetVirtualSizeNil

`func (o *BcsImageV1ApiUpdateImageModelImageModel) SetVirtualSizeNil(b bool)`

 SetVirtualSizeNil sets the value for VirtualSize to be an explicit nil

### UnsetVirtualSize
`func (o *BcsImageV1ApiUpdateImageModelImageModel) UnsetVirtualSize()`

UnsetVirtualSize ensures that no value is present for VirtualSize, not even an explicit nil
### GetIsShared

`func (o *BcsImageV1ApiUpdateImageModelImageModel) GetIsShared() bool`

GetIsShared returns the IsShared field if non-nil, zero value otherwise.

### GetIsSharedOk

`func (o *BcsImageV1ApiUpdateImageModelImageModel) GetIsSharedOk() (*bool, bool)`

GetIsSharedOk returns a tuple with the IsShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShared

`func (o *BcsImageV1ApiUpdateImageModelImageModel) SetIsShared(v bool)`

SetIsShared sets IsShared field to given value.

### HasIsShared

`func (o *BcsImageV1ApiUpdateImageModelImageModel) HasIsShared() bool`

HasIsShared returns a boolean if a field has been set.

### SetIsSharedNil

`func (o *BcsImageV1ApiUpdateImageModelImageModel) SetIsSharedNil(b bool)`

 SetIsSharedNil sets the value for IsShared to be an explicit nil

### UnsetIsShared
`func (o *BcsImageV1ApiUpdateImageModelImageModel) UnsetIsShared()`

UnsetIsShared ensures that no value is present for IsShared, not even an explicit nil
### GetVisibility

`func (o *BcsImageV1ApiUpdateImageModelImageModel) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *BcsImageV1ApiUpdateImageModelImageModel) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *BcsImageV1ApiUpdateImageModelImageModel) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *BcsImageV1ApiUpdateImageModelImageModel) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### SetVisibilityNil

`func (o *BcsImageV1ApiUpdateImageModelImageModel) SetVisibilityNil(b bool)`

 SetVisibilityNil sets the value for Visibility to be an explicit nil

### UnsetVisibility
`func (o *BcsImageV1ApiUpdateImageModelImageModel) UnsetVisibility()`

UnsetVisibility ensures that no value is present for Visibility, not even an explicit nil
### GetDescription

`func (o *BcsImageV1ApiUpdateImageModelImageModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BcsImageV1ApiUpdateImageModelImageModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BcsImageV1ApiUpdateImageModelImageModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BcsImageV1ApiUpdateImageModelImageModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BcsImageV1ApiUpdateImageModelImageModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BcsImageV1ApiUpdateImageModelImageModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetInstanceType

`func (o *BcsImageV1ApiUpdateImageModelImageModel) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *BcsImageV1ApiUpdateImageModelImageModel) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *BcsImageV1ApiUpdateImageModelImageModel) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *BcsImageV1ApiUpdateImageModelImageModel) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### SetInstanceTypeNil

`func (o *BcsImageV1ApiUpdateImageModelImageModel) SetInstanceTypeNil(b bool)`

 SetInstanceTypeNil sets the value for InstanceType to be an explicit nil

### UnsetInstanceType
`func (o *BcsImageV1ApiUpdateImageModelImageModel) UnsetInstanceType()`

UnsetInstanceType ensures that no value is present for InstanceType, not even an explicit nil
### GetOsInfo

`func (o *BcsImageV1ApiUpdateImageModelImageModel) GetOsInfo() BcsImageV1ApiUpdateImageModelOsInfoModel`

GetOsInfo returns the OsInfo field if non-nil, zero value otherwise.

### GetOsInfoOk

`func (o *BcsImageV1ApiUpdateImageModelImageModel) GetOsInfoOk() (*BcsImageV1ApiUpdateImageModelOsInfoModel, bool)`

GetOsInfoOk returns a tuple with the OsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsInfo

`func (o *BcsImageV1ApiUpdateImageModelImageModel) SetOsInfo(v BcsImageV1ApiUpdateImageModelOsInfoModel)`

SetOsInfo sets OsInfo field to given value.

### HasOsInfo

`func (o *BcsImageV1ApiUpdateImageModelImageModel) HasOsInfo() bool`

HasOsInfo returns a boolean if a field has been set.

### SetOsInfoNil

`func (o *BcsImageV1ApiUpdateImageModelImageModel) SetOsInfoNil(b bool)`

 SetOsInfoNil sets the value for OsInfo to be an explicit nil

### UnsetOsInfo
`func (o *BcsImageV1ApiUpdateImageModelImageModel) UnsetOsInfo()`

UnsetOsInfo ensures that no value is present for OsInfo, not even an explicit nil
### GetCreatedAt

`func (o *BcsImageV1ApiUpdateImageModelImageModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BcsImageV1ApiUpdateImageModelImageModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BcsImageV1ApiUpdateImageModelImageModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BcsImageV1ApiUpdateImageModelImageModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *BcsImageV1ApiUpdateImageModelImageModel) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *BcsImageV1ApiUpdateImageModelImageModel) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *BcsImageV1ApiUpdateImageModelImageModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BcsImageV1ApiUpdateImageModelImageModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BcsImageV1ApiUpdateImageModelImageModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BcsImageV1ApiUpdateImageModelImageModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *BcsImageV1ApiUpdateImageModelImageModel) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *BcsImageV1ApiUpdateImageModelImageModel) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


