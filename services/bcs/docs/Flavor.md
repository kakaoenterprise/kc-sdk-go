# Flavor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Vcpus** | Pointer to **NullableInt32** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**IsBurstable** | Pointer to **NullableBool** |  | [optional] 
**Architecture** | Pointer to **NullableString** |  | [optional] 
**Manufacturer** | Pointer to **NullableString** |  | [optional] 
**Group** | Pointer to **NullableString** |  | [optional] 
**InstanceType** | Pointer to **NullableString** |  | [optional] 
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

### NewFlavor

`func NewFlavor() *Flavor`

NewFlavor instantiates a new Flavor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlavorWithDefaults

`func NewFlavorWithDefaults() *Flavor`

NewFlavorWithDefaults instantiates a new Flavor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Flavor) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Flavor) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Flavor) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Flavor) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Flavor) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Flavor) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *Flavor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Flavor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Flavor) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Flavor) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Flavor) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Flavor) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetVcpus

`func (o *Flavor) GetVcpus() int32`

GetVcpus returns the Vcpus field if non-nil, zero value otherwise.

### GetVcpusOk

`func (o *Flavor) GetVcpusOk() (*int32, bool)`

GetVcpusOk returns a tuple with the Vcpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpus

`func (o *Flavor) SetVcpus(v int32)`

SetVcpus sets Vcpus field to given value.

### HasVcpus

`func (o *Flavor) HasVcpus() bool`

HasVcpus returns a boolean if a field has been set.

### SetVcpusNil

`func (o *Flavor) SetVcpusNil(b bool)`

 SetVcpusNil sets the value for Vcpus to be an explicit nil

### UnsetVcpus
`func (o *Flavor) UnsetVcpus()`

UnsetVcpus ensures that no value is present for Vcpus, not even an explicit nil
### GetDescription

`func (o *Flavor) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Flavor) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Flavor) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Flavor) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Flavor) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Flavor) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIsBurstable

`func (o *Flavor) GetIsBurstable() bool`

GetIsBurstable returns the IsBurstable field if non-nil, zero value otherwise.

### GetIsBurstableOk

`func (o *Flavor) GetIsBurstableOk() (*bool, bool)`

GetIsBurstableOk returns a tuple with the IsBurstable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBurstable

`func (o *Flavor) SetIsBurstable(v bool)`

SetIsBurstable sets IsBurstable field to given value.

### HasIsBurstable

`func (o *Flavor) HasIsBurstable() bool`

HasIsBurstable returns a boolean if a field has been set.

### SetIsBurstableNil

`func (o *Flavor) SetIsBurstableNil(b bool)`

 SetIsBurstableNil sets the value for IsBurstable to be an explicit nil

### UnsetIsBurstable
`func (o *Flavor) UnsetIsBurstable()`

UnsetIsBurstable ensures that no value is present for IsBurstable, not even an explicit nil
### GetArchitecture

`func (o *Flavor) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *Flavor) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *Flavor) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.

### HasArchitecture

`func (o *Flavor) HasArchitecture() bool`

HasArchitecture returns a boolean if a field has been set.

### SetArchitectureNil

`func (o *Flavor) SetArchitectureNil(b bool)`

 SetArchitectureNil sets the value for Architecture to be an explicit nil

### UnsetArchitecture
`func (o *Flavor) UnsetArchitecture()`

UnsetArchitecture ensures that no value is present for Architecture, not even an explicit nil
### GetManufacturer

`func (o *Flavor) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *Flavor) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *Flavor) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *Flavor) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### SetManufacturerNil

`func (o *Flavor) SetManufacturerNil(b bool)`

 SetManufacturerNil sets the value for Manufacturer to be an explicit nil

### UnsetManufacturer
`func (o *Flavor) UnsetManufacturer()`

UnsetManufacturer ensures that no value is present for Manufacturer, not even an explicit nil
### GetGroup

`func (o *Flavor) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *Flavor) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *Flavor) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *Flavor) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *Flavor) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *Flavor) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetInstanceType

`func (o *Flavor) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *Flavor) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *Flavor) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *Flavor) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### SetInstanceTypeNil

`func (o *Flavor) SetInstanceTypeNil(b bool)`

 SetInstanceTypeNil sets the value for InstanceType to be an explicit nil

### UnsetInstanceType
`func (o *Flavor) UnsetInstanceType()`

UnsetInstanceType ensures that no value is present for InstanceType, not even an explicit nil
### GetProcessor

`func (o *Flavor) GetProcessor() string`

GetProcessor returns the Processor field if non-nil, zero value otherwise.

### GetProcessorOk

`func (o *Flavor) GetProcessorOk() (*string, bool)`

GetProcessorOk returns a tuple with the Processor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessor

`func (o *Flavor) SetProcessor(v string)`

SetProcessor sets Processor field to given value.

### HasProcessor

`func (o *Flavor) HasProcessor() bool`

HasProcessor returns a boolean if a field has been set.

### SetProcessorNil

`func (o *Flavor) SetProcessorNil(b bool)`

 SetProcessorNil sets the value for Processor to be an explicit nil

### UnsetProcessor
`func (o *Flavor) UnsetProcessor()`

UnsetProcessor ensures that no value is present for Processor, not even an explicit nil
### GetMemoryMb

`func (o *Flavor) GetMemoryMb() int64`

GetMemoryMb returns the MemoryMb field if non-nil, zero value otherwise.

### GetMemoryMbOk

`func (o *Flavor) GetMemoryMbOk() (*int64, bool)`

GetMemoryMbOk returns a tuple with the MemoryMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMb

`func (o *Flavor) SetMemoryMb(v int64)`

SetMemoryMb sets MemoryMb field to given value.

### HasMemoryMb

`func (o *Flavor) HasMemoryMb() bool`

HasMemoryMb returns a boolean if a field has been set.

### SetMemoryMbNil

`func (o *Flavor) SetMemoryMbNil(b bool)`

 SetMemoryMbNil sets the value for MemoryMb to be an explicit nil

### UnsetMemoryMb
`func (o *Flavor) UnsetMemoryMb()`

UnsetMemoryMb ensures that no value is present for MemoryMb, not even an explicit nil
### GetCreatedAt

`func (o *Flavor) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Flavor) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Flavor) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Flavor) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *Flavor) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *Flavor) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *Flavor) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Flavor) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Flavor) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Flavor) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *Flavor) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *Flavor) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetAvailabilityZone

`func (o *Flavor) GetAvailabilityZone() []*AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *Flavor) GetAvailabilityZoneOk() (*[]*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *Flavor) SetAvailabilityZone(v []*AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *Flavor) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### SetAvailabilityZoneNil

`func (o *Flavor) SetAvailabilityZoneNil(b bool)`

 SetAvailabilityZoneNil sets the value for AvailabilityZone to be an explicit nil

### UnsetAvailabilityZone
`func (o *Flavor) UnsetAvailabilityZone()`

UnsetAvailabilityZone ensures that no value is present for AvailabilityZone, not even an explicit nil
### GetAvailable

`func (o *Flavor) GetAvailable() map[string]int32`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *Flavor) GetAvailableOk() (*map[string]int32, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *Flavor) SetAvailable(v map[string]int32)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *Flavor) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### SetAvailableNil

`func (o *Flavor) SetAvailableNil(b bool)`

 SetAvailableNil sets the value for Available to be an explicit nil

### UnsetAvailable
`func (o *Flavor) UnsetAvailable()`

UnsetAvailable ensures that no value is present for Available, not even an explicit nil
### GetInstanceFamily

`func (o *Flavor) GetInstanceFamily() string`

GetInstanceFamily returns the InstanceFamily field if non-nil, zero value otherwise.

### GetInstanceFamilyOk

`func (o *Flavor) GetInstanceFamilyOk() (*string, bool)`

GetInstanceFamilyOk returns a tuple with the InstanceFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceFamily

`func (o *Flavor) SetInstanceFamily(v string)`

SetInstanceFamily sets InstanceFamily field to given value.

### HasInstanceFamily

`func (o *Flavor) HasInstanceFamily() bool`

HasInstanceFamily returns a boolean if a field has been set.

### SetInstanceFamilyNil

`func (o *Flavor) SetInstanceFamilyNil(b bool)`

 SetInstanceFamilyNil sets the value for InstanceFamily to be an explicit nil

### UnsetInstanceFamily
`func (o *Flavor) UnsetInstanceFamily()`

UnsetInstanceFamily ensures that no value is present for InstanceFamily, not even an explicit nil
### GetInstanceSize

`func (o *Flavor) GetInstanceSize() string`

GetInstanceSize returns the InstanceSize field if non-nil, zero value otherwise.

### GetInstanceSizeOk

`func (o *Flavor) GetInstanceSizeOk() (*string, bool)`

GetInstanceSizeOk returns a tuple with the InstanceSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSize

`func (o *Flavor) SetInstanceSize(v string)`

SetInstanceSize sets InstanceSize field to given value.

### HasInstanceSize

`func (o *Flavor) HasInstanceSize() bool`

HasInstanceSize returns a boolean if a field has been set.

### SetInstanceSizeNil

`func (o *Flavor) SetInstanceSizeNil(b bool)`

 SetInstanceSizeNil sets the value for InstanceSize to be an explicit nil

### UnsetInstanceSize
`func (o *Flavor) UnsetInstanceSize()`

UnsetInstanceSize ensures that no value is present for InstanceSize, not even an explicit nil
### GetDiskType

`func (o *Flavor) GetDiskType() string`

GetDiskType returns the DiskType field if non-nil, zero value otherwise.

### GetDiskTypeOk

`func (o *Flavor) GetDiskTypeOk() (*string, bool)`

GetDiskTypeOk returns a tuple with the DiskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskType

`func (o *Flavor) SetDiskType(v string)`

SetDiskType sets DiskType field to given value.

### HasDiskType

`func (o *Flavor) HasDiskType() bool`

HasDiskType returns a boolean if a field has been set.

### SetDiskTypeNil

`func (o *Flavor) SetDiskTypeNil(b bool)`

 SetDiskTypeNil sets the value for DiskType to be an explicit nil

### UnsetDiskType
`func (o *Flavor) UnsetDiskType()`

UnsetDiskType ensures that no value is present for DiskType, not even an explicit nil
### GetRootGb

`func (o *Flavor) GetRootGb() int32`

GetRootGb returns the RootGb field if non-nil, zero value otherwise.

### GetRootGbOk

`func (o *Flavor) GetRootGbOk() (*int32, bool)`

GetRootGbOk returns a tuple with the RootGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootGb

`func (o *Flavor) SetRootGb(v int32)`

SetRootGb sets RootGb field to given value.

### HasRootGb

`func (o *Flavor) HasRootGb() bool`

HasRootGb returns a boolean if a field has been set.

### SetRootGbNil

`func (o *Flavor) SetRootGbNil(b bool)`

 SetRootGbNil sets the value for RootGb to be an explicit nil

### UnsetRootGb
`func (o *Flavor) UnsetRootGb()`

UnsetRootGb ensures that no value is present for RootGb, not even an explicit nil
### GetOsDistro

`func (o *Flavor) GetOsDistro() string`

GetOsDistro returns the OsDistro field if non-nil, zero value otherwise.

### GetOsDistroOk

`func (o *Flavor) GetOsDistroOk() (*string, bool)`

GetOsDistroOk returns a tuple with the OsDistro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsDistro

`func (o *Flavor) SetOsDistro(v string)`

SetOsDistro sets OsDistro field to given value.

### HasOsDistro

`func (o *Flavor) HasOsDistro() bool`

HasOsDistro returns a boolean if a field has been set.

### SetOsDistroNil

`func (o *Flavor) SetOsDistroNil(b bool)`

 SetOsDistroNil sets the value for OsDistro to be an explicit nil

### UnsetOsDistro
`func (o *Flavor) UnsetOsDistro()`

UnsetOsDistro ensures that no value is present for OsDistro, not even an explicit nil
### GetHwCount

`func (o *Flavor) GetHwCount() int32`

GetHwCount returns the HwCount field if non-nil, zero value otherwise.

### GetHwCountOk

`func (o *Flavor) GetHwCountOk() (*int32, bool)`

GetHwCountOk returns a tuple with the HwCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwCount

`func (o *Flavor) SetHwCount(v int32)`

SetHwCount sets HwCount field to given value.

### HasHwCount

`func (o *Flavor) HasHwCount() bool`

HasHwCount returns a boolean if a field has been set.

### SetHwCountNil

`func (o *Flavor) SetHwCountNil(b bool)`

 SetHwCountNil sets the value for HwCount to be an explicit nil

### UnsetHwCount
`func (o *Flavor) UnsetHwCount()`

UnsetHwCount ensures that no value is present for HwCount, not even an explicit nil
### GetHwType

`func (o *Flavor) GetHwType() string`

GetHwType returns the HwType field if non-nil, zero value otherwise.

### GetHwTypeOk

`func (o *Flavor) GetHwTypeOk() (*string, bool)`

GetHwTypeOk returns a tuple with the HwType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwType

`func (o *Flavor) SetHwType(v string)`

SetHwType sets HwType field to given value.

### HasHwType

`func (o *Flavor) HasHwType() bool`

HasHwType returns a boolean if a field has been set.

### SetHwTypeNil

`func (o *Flavor) SetHwTypeNil(b bool)`

 SetHwTypeNil sets the value for HwType to be an explicit nil

### UnsetHwType
`func (o *Flavor) UnsetHwType()`

UnsetHwType ensures that no value is present for HwType, not even an explicit nil
### GetHwName

`func (o *Flavor) GetHwName() string`

GetHwName returns the HwName field if non-nil, zero value otherwise.

### GetHwNameOk

`func (o *Flavor) GetHwNameOk() (*string, bool)`

GetHwNameOk returns a tuple with the HwName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwName

`func (o *Flavor) SetHwName(v string)`

SetHwName sets HwName field to given value.

### HasHwName

`func (o *Flavor) HasHwName() bool`

HasHwName returns a boolean if a field has been set.

### SetHwNameNil

`func (o *Flavor) SetHwNameNil(b bool)`

 SetHwNameNil sets the value for HwName to be an explicit nil

### UnsetHwName
`func (o *Flavor) UnsetHwName()`

UnsetHwName ensures that no value is present for HwName, not even an explicit nil
### GetMaximumNetworkInterfaces

`func (o *Flavor) GetMaximumNetworkInterfaces() int32`

GetMaximumNetworkInterfaces returns the MaximumNetworkInterfaces field if non-nil, zero value otherwise.

### GetMaximumNetworkInterfacesOk

`func (o *Flavor) GetMaximumNetworkInterfacesOk() (*int32, bool)`

GetMaximumNetworkInterfacesOk returns a tuple with the MaximumNetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumNetworkInterfaces

`func (o *Flavor) SetMaximumNetworkInterfaces(v int32)`

SetMaximumNetworkInterfaces sets MaximumNetworkInterfaces field to given value.

### HasMaximumNetworkInterfaces

`func (o *Flavor) HasMaximumNetworkInterfaces() bool`

HasMaximumNetworkInterfaces returns a boolean if a field has been set.

### SetMaximumNetworkInterfacesNil

`func (o *Flavor) SetMaximumNetworkInterfacesNil(b bool)`

 SetMaximumNetworkInterfacesNil sets the value for MaximumNetworkInterfaces to be an explicit nil

### UnsetMaximumNetworkInterfaces
`func (o *Flavor) UnsetMaximumNetworkInterfaces()`

UnsetMaximumNetworkInterfaces ensures that no value is present for MaximumNetworkInterfaces, not even an explicit nil
### GetIsHyperThreadingDisabled

`func (o *Flavor) GetIsHyperThreadingDisabled() bool`

GetIsHyperThreadingDisabled returns the IsHyperThreadingDisabled field if non-nil, zero value otherwise.

### GetIsHyperThreadingDisabledOk

`func (o *Flavor) GetIsHyperThreadingDisabledOk() (*bool, bool)`

GetIsHyperThreadingDisabledOk returns a tuple with the IsHyperThreadingDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHyperThreadingDisabled

`func (o *Flavor) SetIsHyperThreadingDisabled(v bool)`

SetIsHyperThreadingDisabled sets IsHyperThreadingDisabled field to given value.

### HasIsHyperThreadingDisabled

`func (o *Flavor) HasIsHyperThreadingDisabled() bool`

HasIsHyperThreadingDisabled returns a boolean if a field has been set.

### SetIsHyperThreadingDisabledNil

`func (o *Flavor) SetIsHyperThreadingDisabledNil(b bool)`

 SetIsHyperThreadingDisabledNil sets the value for IsHyperThreadingDisabled to be an explicit nil

### UnsetIsHyperThreadingDisabled
`func (o *Flavor) UnsetIsHyperThreadingDisabled()`

UnsetIsHyperThreadingDisabled ensures that no value is present for IsHyperThreadingDisabled, not even an explicit nil
### GetIsHyperThreadingSupported

`func (o *Flavor) GetIsHyperThreadingSupported() bool`

GetIsHyperThreadingSupported returns the IsHyperThreadingSupported field if non-nil, zero value otherwise.

### GetIsHyperThreadingSupportedOk

`func (o *Flavor) GetIsHyperThreadingSupportedOk() (*bool, bool)`

GetIsHyperThreadingSupportedOk returns a tuple with the IsHyperThreadingSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHyperThreadingSupported

`func (o *Flavor) SetIsHyperThreadingSupported(v bool)`

SetIsHyperThreadingSupported sets IsHyperThreadingSupported field to given value.

### HasIsHyperThreadingSupported

`func (o *Flavor) HasIsHyperThreadingSupported() bool`

HasIsHyperThreadingSupported returns a boolean if a field has been set.

### SetIsHyperThreadingSupportedNil

`func (o *Flavor) SetIsHyperThreadingSupportedNil(b bool)`

 SetIsHyperThreadingSupportedNil sets the value for IsHyperThreadingSupported to be an explicit nil

### UnsetIsHyperThreadingSupported
`func (o *Flavor) UnsetIsHyperThreadingSupported()`

UnsetIsHyperThreadingSupported ensures that no value is present for IsHyperThreadingSupported, not even an explicit nil
### GetIsHyperThreadingDisableSupported

`func (o *Flavor) GetIsHyperThreadingDisableSupported() bool`

GetIsHyperThreadingDisableSupported returns the IsHyperThreadingDisableSupported field if non-nil, zero value otherwise.

### GetIsHyperThreadingDisableSupportedOk

`func (o *Flavor) GetIsHyperThreadingDisableSupportedOk() (*bool, bool)`

GetIsHyperThreadingDisableSupportedOk returns a tuple with the IsHyperThreadingDisableSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHyperThreadingDisableSupported

`func (o *Flavor) SetIsHyperThreadingDisableSupported(v bool)`

SetIsHyperThreadingDisableSupported sets IsHyperThreadingDisableSupported field to given value.

### HasIsHyperThreadingDisableSupported

`func (o *Flavor) HasIsHyperThreadingDisableSupported() bool`

HasIsHyperThreadingDisableSupported returns a boolean if a field has been set.

### SetIsHyperThreadingDisableSupportedNil

`func (o *Flavor) SetIsHyperThreadingDisableSupportedNil(b bool)`

 SetIsHyperThreadingDisableSupportedNil sets the value for IsHyperThreadingDisableSupported to be an explicit nil

### UnsetIsHyperThreadingDisableSupported
`func (o *Flavor) UnsetIsHyperThreadingDisableSupported()`

UnsetIsHyperThreadingDisableSupported ensures that no value is present for IsHyperThreadingDisableSupported, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


