# BcsInstanceV1ApiGetInstanceTypeModelFlavorModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 인스턴스 유형 ID | 
**Name** | Pointer to **NullableString** |  | [optional] 
**Vcpus** | Pointer to **NullableInt32** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**IsBurstable** | Pointer to **NullableBool** |  | [optional] 
**Architecture** | Pointer to **NullableString** |  | [optional] 
**Manufacturer** | Pointer to **NullableString** |  | [optional] 
**Group** | Pointer to **NullableString** |  | [optional] 
**InstanceType** | Pointer to [**NullableInstanceType**](InstanceType.md) |  | [optional] 
**Processor** | Pointer to **NullableString** |  | [optional] 
**MemoryMb** | Pointer to **NullableInt64** |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 
**AvailabilityZone** | Pointer to [**[]AvailabilityZone**](AvailabilityZone.md) |  | [optional] 
**Available** | Pointer to **map[string]int32** |  | [optional] 
**InstanceFamily** | Pointer to **NullableString** |  | [optional] 
**InstanceSize** | Pointer to **NullableString** |  | [optional] 
**DiskType** | Pointer to **NullableString** |  | [optional] 
**RootGb** | Pointer to **NullableInt32** |  | [optional] 
**OsDistro** | Pointer to **NullableString** |  | [optional] 
**HwCount** | Pointer to **NullableInt32** |  | [optional] 
**HwType** | Pointer to **NullableString** |  | [optional] 
**HwName** | Pointer to **NullableString** |  | [optional] 
**MaximumNetworkInterfaces** | Pointer to **NullableInt32** |  | [optional] 
**IsHyperThreadingDisabled** | Pointer to **NullableBool** |  | [optional] 
**IsHyperThreadingSupported** | Pointer to **NullableBool** |  | [optional] 
**IsHyperThreadingDisableSupported** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewBcsInstanceV1ApiGetInstanceTypeModelFlavorModel

`func NewBcsInstanceV1ApiGetInstanceTypeModelFlavorModel(id string, ) *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel`

NewBcsInstanceV1ApiGetInstanceTypeModelFlavorModel instantiates a new BcsInstanceV1ApiGetInstanceTypeModelFlavorModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBcsInstanceV1ApiGetInstanceTypeModelFlavorModelWithDefaults

`func NewBcsInstanceV1ApiGetInstanceTypeModelFlavorModelWithDefaults() *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel`

NewBcsInstanceV1ApiGetInstanceTypeModelFlavorModelWithDefaults instantiates a new BcsInstanceV1ApiGetInstanceTypeModelFlavorModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetVcpus

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) GetVcpus() int32`

GetVcpus returns the Vcpus field if non-nil, zero value otherwise.

### GetVcpusOk

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) GetVcpusOk() (*int32, bool)`

GetVcpusOk returns a tuple with the Vcpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpus

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) SetVcpus(v int32)`

SetVcpus sets Vcpus field to given value.

### HasVcpus

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) HasVcpus() bool`

HasVcpus returns a boolean if a field has been set.

### SetVcpusNil

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) SetVcpusNil(b bool)`

 SetVcpusNil sets the value for Vcpus to be an explicit nil

### UnsetVcpus
`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) UnsetVcpus()`

UnsetVcpus ensures that no value is present for Vcpus, not even an explicit nil
### GetDescription

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIsBurstable

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) GetIsBurstable() bool`

GetIsBurstable returns the IsBurstable field if non-nil, zero value otherwise.

### GetIsBurstableOk

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) GetIsBurstableOk() (*bool, bool)`

GetIsBurstableOk returns a tuple with the IsBurstable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBurstable

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) SetIsBurstable(v bool)`

SetIsBurstable sets IsBurstable field to given value.

### HasIsBurstable

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) HasIsBurstable() bool`

HasIsBurstable returns a boolean if a field has been set.

### SetIsBurstableNil

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) SetIsBurstableNil(b bool)`

 SetIsBurstableNil sets the value for IsBurstable to be an explicit nil

### UnsetIsBurstable
`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) UnsetIsBurstable()`

UnsetIsBurstable ensures that no value is present for IsBurstable, not even an explicit nil
### GetArchitecture

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.

### HasArchitecture

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) HasArchitecture() bool`

HasArchitecture returns a boolean if a field has been set.

### SetArchitectureNil

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) SetArchitectureNil(b bool)`

 SetArchitectureNil sets the value for Architecture to be an explicit nil

### UnsetArchitecture
`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) UnsetArchitecture()`

UnsetArchitecture ensures that no value is present for Architecture, not even an explicit nil
### GetManufacturer

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### SetManufacturerNil

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) SetManufacturerNil(b bool)`

 SetManufacturerNil sets the value for Manufacturer to be an explicit nil

### UnsetManufacturer
`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) UnsetManufacturer()`

UnsetManufacturer ensures that no value is present for Manufacturer, not even an explicit nil
### GetGroup

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetInstanceType

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) GetInstanceType() InstanceType`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) GetInstanceTypeOk() (*InstanceType, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) SetInstanceType(v InstanceType)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### SetInstanceTypeNil

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) SetInstanceTypeNil(b bool)`

 SetInstanceTypeNil sets the value for InstanceType to be an explicit nil

### UnsetInstanceType
`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) UnsetInstanceType()`

UnsetInstanceType ensures that no value is present for InstanceType, not even an explicit nil
### GetProcessor

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) GetProcessor() string`

GetProcessor returns the Processor field if non-nil, zero value otherwise.

### GetProcessorOk

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) GetProcessorOk() (*string, bool)`

GetProcessorOk returns a tuple with the Processor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessor

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) SetProcessor(v string)`

SetProcessor sets Processor field to given value.

### HasProcessor

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) HasProcessor() bool`

HasProcessor returns a boolean if a field has been set.

### SetProcessorNil

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) SetProcessorNil(b bool)`

 SetProcessorNil sets the value for Processor to be an explicit nil

### UnsetProcessor
`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) UnsetProcessor()`

UnsetProcessor ensures that no value is present for Processor, not even an explicit nil
### GetMemoryMb

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) GetMemoryMb() int64`

GetMemoryMb returns the MemoryMb field if non-nil, zero value otherwise.

### GetMemoryMbOk

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) GetMemoryMbOk() (*int64, bool)`

GetMemoryMbOk returns a tuple with the MemoryMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMb

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) SetMemoryMb(v int64)`

SetMemoryMb sets MemoryMb field to given value.

### HasMemoryMb

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) HasMemoryMb() bool`

HasMemoryMb returns a boolean if a field has been set.

### SetMemoryMbNil

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) SetMemoryMbNil(b bool)`

 SetMemoryMbNil sets the value for MemoryMb to be an explicit nil

### UnsetMemoryMb
`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) UnsetMemoryMb()`

UnsetMemoryMb ensures that no value is present for MemoryMb, not even an explicit nil
### GetCreatedAt

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetAvailabilityZone

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) GetAvailabilityZone() []AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) GetAvailabilityZoneOk() (*[]AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) SetAvailabilityZone(v []AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### SetAvailabilityZoneNil

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) SetAvailabilityZoneNil(b bool)`

 SetAvailabilityZoneNil sets the value for AvailabilityZone to be an explicit nil

### UnsetAvailabilityZone
`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) UnsetAvailabilityZone()`

UnsetAvailabilityZone ensures that no value is present for AvailabilityZone, not even an explicit nil
### GetAvailable

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) GetAvailable() map[string]int32`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) GetAvailableOk() (*map[string]int32, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) SetAvailable(v map[string]int32)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### SetAvailableNil

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) SetAvailableNil(b bool)`

 SetAvailableNil sets the value for Available to be an explicit nil

### UnsetAvailable
`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) UnsetAvailable()`

UnsetAvailable ensures that no value is present for Available, not even an explicit nil
### GetInstanceFamily

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) GetInstanceFamily() string`

GetInstanceFamily returns the InstanceFamily field if non-nil, zero value otherwise.

### GetInstanceFamilyOk

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) GetInstanceFamilyOk() (*string, bool)`

GetInstanceFamilyOk returns a tuple with the InstanceFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceFamily

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) SetInstanceFamily(v string)`

SetInstanceFamily sets InstanceFamily field to given value.

### HasInstanceFamily

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) HasInstanceFamily() bool`

HasInstanceFamily returns a boolean if a field has been set.

### SetInstanceFamilyNil

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) SetInstanceFamilyNil(b bool)`

 SetInstanceFamilyNil sets the value for InstanceFamily to be an explicit nil

### UnsetInstanceFamily
`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) UnsetInstanceFamily()`

UnsetInstanceFamily ensures that no value is present for InstanceFamily, not even an explicit nil
### GetInstanceSize

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) GetInstanceSize() string`

GetInstanceSize returns the InstanceSize field if non-nil, zero value otherwise.

### GetInstanceSizeOk

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) GetInstanceSizeOk() (*string, bool)`

GetInstanceSizeOk returns a tuple with the InstanceSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSize

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) SetInstanceSize(v string)`

SetInstanceSize sets InstanceSize field to given value.

### HasInstanceSize

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) HasInstanceSize() bool`

HasInstanceSize returns a boolean if a field has been set.

### SetInstanceSizeNil

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) SetInstanceSizeNil(b bool)`

 SetInstanceSizeNil sets the value for InstanceSize to be an explicit nil

### UnsetInstanceSize
`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) UnsetInstanceSize()`

UnsetInstanceSize ensures that no value is present for InstanceSize, not even an explicit nil
### GetDiskType

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) GetDiskType() string`

GetDiskType returns the DiskType field if non-nil, zero value otherwise.

### GetDiskTypeOk

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) GetDiskTypeOk() (*string, bool)`

GetDiskTypeOk returns a tuple with the DiskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskType

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) SetDiskType(v string)`

SetDiskType sets DiskType field to given value.

### HasDiskType

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) HasDiskType() bool`

HasDiskType returns a boolean if a field has been set.

### SetDiskTypeNil

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) SetDiskTypeNil(b bool)`

 SetDiskTypeNil sets the value for DiskType to be an explicit nil

### UnsetDiskType
`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) UnsetDiskType()`

UnsetDiskType ensures that no value is present for DiskType, not even an explicit nil
### GetRootGb

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) GetRootGb() int32`

GetRootGb returns the RootGb field if non-nil, zero value otherwise.

### GetRootGbOk

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) GetRootGbOk() (*int32, bool)`

GetRootGbOk returns a tuple with the RootGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootGb

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) SetRootGb(v int32)`

SetRootGb sets RootGb field to given value.

### HasRootGb

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) HasRootGb() bool`

HasRootGb returns a boolean if a field has been set.

### SetRootGbNil

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) SetRootGbNil(b bool)`

 SetRootGbNil sets the value for RootGb to be an explicit nil

### UnsetRootGb
`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) UnsetRootGb()`

UnsetRootGb ensures that no value is present for RootGb, not even an explicit nil
### GetOsDistro

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) GetOsDistro() string`

GetOsDistro returns the OsDistro field if non-nil, zero value otherwise.

### GetOsDistroOk

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) GetOsDistroOk() (*string, bool)`

GetOsDistroOk returns a tuple with the OsDistro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsDistro

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) SetOsDistro(v string)`

SetOsDistro sets OsDistro field to given value.

### HasOsDistro

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) HasOsDistro() bool`

HasOsDistro returns a boolean if a field has been set.

### SetOsDistroNil

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) SetOsDistroNil(b bool)`

 SetOsDistroNil sets the value for OsDistro to be an explicit nil

### UnsetOsDistro
`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) UnsetOsDistro()`

UnsetOsDistro ensures that no value is present for OsDistro, not even an explicit nil
### GetHwCount

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) GetHwCount() int32`

GetHwCount returns the HwCount field if non-nil, zero value otherwise.

### GetHwCountOk

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) GetHwCountOk() (*int32, bool)`

GetHwCountOk returns a tuple with the HwCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwCount

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) SetHwCount(v int32)`

SetHwCount sets HwCount field to given value.

### HasHwCount

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) HasHwCount() bool`

HasHwCount returns a boolean if a field has been set.

### SetHwCountNil

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) SetHwCountNil(b bool)`

 SetHwCountNil sets the value for HwCount to be an explicit nil

### UnsetHwCount
`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) UnsetHwCount()`

UnsetHwCount ensures that no value is present for HwCount, not even an explicit nil
### GetHwType

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) GetHwType() string`

GetHwType returns the HwType field if non-nil, zero value otherwise.

### GetHwTypeOk

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) GetHwTypeOk() (*string, bool)`

GetHwTypeOk returns a tuple with the HwType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwType

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) SetHwType(v string)`

SetHwType sets HwType field to given value.

### HasHwType

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) HasHwType() bool`

HasHwType returns a boolean if a field has been set.

### SetHwTypeNil

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) SetHwTypeNil(b bool)`

 SetHwTypeNil sets the value for HwType to be an explicit nil

### UnsetHwType
`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) UnsetHwType()`

UnsetHwType ensures that no value is present for HwType, not even an explicit nil
### GetHwName

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) GetHwName() string`

GetHwName returns the HwName field if non-nil, zero value otherwise.

### GetHwNameOk

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) GetHwNameOk() (*string, bool)`

GetHwNameOk returns a tuple with the HwName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwName

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) SetHwName(v string)`

SetHwName sets HwName field to given value.

### HasHwName

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) HasHwName() bool`

HasHwName returns a boolean if a field has been set.

### SetHwNameNil

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) SetHwNameNil(b bool)`

 SetHwNameNil sets the value for HwName to be an explicit nil

### UnsetHwName
`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) UnsetHwName()`

UnsetHwName ensures that no value is present for HwName, not even an explicit nil
### GetMaximumNetworkInterfaces

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) GetMaximumNetworkInterfaces() int32`

GetMaximumNetworkInterfaces returns the MaximumNetworkInterfaces field if non-nil, zero value otherwise.

### GetMaximumNetworkInterfacesOk

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) GetMaximumNetworkInterfacesOk() (*int32, bool)`

GetMaximumNetworkInterfacesOk returns a tuple with the MaximumNetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumNetworkInterfaces

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) SetMaximumNetworkInterfaces(v int32)`

SetMaximumNetworkInterfaces sets MaximumNetworkInterfaces field to given value.

### HasMaximumNetworkInterfaces

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) HasMaximumNetworkInterfaces() bool`

HasMaximumNetworkInterfaces returns a boolean if a field has been set.

### SetMaximumNetworkInterfacesNil

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) SetMaximumNetworkInterfacesNil(b bool)`

 SetMaximumNetworkInterfacesNil sets the value for MaximumNetworkInterfaces to be an explicit nil

### UnsetMaximumNetworkInterfaces
`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) UnsetMaximumNetworkInterfaces()`

UnsetMaximumNetworkInterfaces ensures that no value is present for MaximumNetworkInterfaces, not even an explicit nil
### GetIsHyperThreadingDisabled

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) GetIsHyperThreadingDisabled() bool`

GetIsHyperThreadingDisabled returns the IsHyperThreadingDisabled field if non-nil, zero value otherwise.

### GetIsHyperThreadingDisabledOk

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) GetIsHyperThreadingDisabledOk() (*bool, bool)`

GetIsHyperThreadingDisabledOk returns a tuple with the IsHyperThreadingDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHyperThreadingDisabled

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) SetIsHyperThreadingDisabled(v bool)`

SetIsHyperThreadingDisabled sets IsHyperThreadingDisabled field to given value.

### HasIsHyperThreadingDisabled

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) HasIsHyperThreadingDisabled() bool`

HasIsHyperThreadingDisabled returns a boolean if a field has been set.

### SetIsHyperThreadingDisabledNil

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) SetIsHyperThreadingDisabledNil(b bool)`

 SetIsHyperThreadingDisabledNil sets the value for IsHyperThreadingDisabled to be an explicit nil

### UnsetIsHyperThreadingDisabled
`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) UnsetIsHyperThreadingDisabled()`

UnsetIsHyperThreadingDisabled ensures that no value is present for IsHyperThreadingDisabled, not even an explicit nil
### GetIsHyperThreadingSupported

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) GetIsHyperThreadingSupported() bool`

GetIsHyperThreadingSupported returns the IsHyperThreadingSupported field if non-nil, zero value otherwise.

### GetIsHyperThreadingSupportedOk

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) GetIsHyperThreadingSupportedOk() (*bool, bool)`

GetIsHyperThreadingSupportedOk returns a tuple with the IsHyperThreadingSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHyperThreadingSupported

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) SetIsHyperThreadingSupported(v bool)`

SetIsHyperThreadingSupported sets IsHyperThreadingSupported field to given value.

### HasIsHyperThreadingSupported

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) HasIsHyperThreadingSupported() bool`

HasIsHyperThreadingSupported returns a boolean if a field has been set.

### SetIsHyperThreadingSupportedNil

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) SetIsHyperThreadingSupportedNil(b bool)`

 SetIsHyperThreadingSupportedNil sets the value for IsHyperThreadingSupported to be an explicit nil

### UnsetIsHyperThreadingSupported
`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) UnsetIsHyperThreadingSupported()`

UnsetIsHyperThreadingSupported ensures that no value is present for IsHyperThreadingSupported, not even an explicit nil
### GetIsHyperThreadingDisableSupported

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) GetIsHyperThreadingDisableSupported() bool`

GetIsHyperThreadingDisableSupported returns the IsHyperThreadingDisableSupported field if non-nil, zero value otherwise.

### GetIsHyperThreadingDisableSupportedOk

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) GetIsHyperThreadingDisableSupportedOk() (*bool, bool)`

GetIsHyperThreadingDisableSupportedOk returns a tuple with the IsHyperThreadingDisableSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHyperThreadingDisableSupported

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) SetIsHyperThreadingDisableSupported(v bool)`

SetIsHyperThreadingDisableSupported sets IsHyperThreadingDisableSupported field to given value.

### HasIsHyperThreadingDisableSupported

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) HasIsHyperThreadingDisableSupported() bool`

HasIsHyperThreadingDisableSupported returns a boolean if a field has been set.

### SetIsHyperThreadingDisableSupportedNil

`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) SetIsHyperThreadingDisableSupportedNil(b bool)`

 SetIsHyperThreadingDisableSupportedNil sets the value for IsHyperThreadingDisableSupported to be an explicit nil

### UnsetIsHyperThreadingDisableSupported
`func (o *BcsInstanceV1ApiGetInstanceTypeModelFlavorModel) UnsetIsHyperThreadingDisableSupported()`

UnsetIsHyperThreadingDisableSupported ensures that no value is present for IsHyperThreadingDisableSupported, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


