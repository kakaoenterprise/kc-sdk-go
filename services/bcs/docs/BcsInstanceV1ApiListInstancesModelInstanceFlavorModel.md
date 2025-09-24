# BcsInstanceV1ApiListInstancesModelInstanceFlavorModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 인스턴스 유형 ID | 
**Name** | Pointer to **NullableString** |  | [optional] 
**Group** | Pointer to **NullableString** |  | [optional] 
**Vcpus** | Pointer to **NullableInt32** |  | [optional] 
**IsBurstable** | Pointer to **NullableBool** |  | [optional] 
**Manufacturer** | Pointer to **NullableString** |  | [optional] 
**MemoryMb** | Pointer to **NullableInt32** |  | [optional] 
**RootGb** | Pointer to **NullableInt32** |  | [optional] 
**DiskType** | Pointer to **NullableString** |  | [optional] 
**InstanceFamily** | Pointer to **NullableString** |  | [optional] 
**OsDistro** | Pointer to **[]string** |  | [optional] 
**MaximumNetworkInterfaces** | Pointer to **NullableInt32** |  | [optional] 
**HwType** | Pointer to **NullableString** |  | [optional] 
**HwCount** | Pointer to **NullableInt32** |  | [optional] 
**IsHyperThreadingSupported** | Pointer to **NullableBool** |  | [optional] 
**RealVcpus** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewBcsInstanceV1ApiListInstancesModelInstanceFlavorModel

`func NewBcsInstanceV1ApiListInstancesModelInstanceFlavorModel(id string, ) *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel`

NewBcsInstanceV1ApiListInstancesModelInstanceFlavorModel instantiates a new BcsInstanceV1ApiListInstancesModelInstanceFlavorModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBcsInstanceV1ApiListInstancesModelInstanceFlavorModelWithDefaults

`func NewBcsInstanceV1ApiListInstancesModelInstanceFlavorModelWithDefaults() *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel`

NewBcsInstanceV1ApiListInstancesModelInstanceFlavorModelWithDefaults instantiates a new BcsInstanceV1ApiListInstancesModelInstanceFlavorModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetGroup

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetVcpus

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) GetVcpus() int32`

GetVcpus returns the Vcpus field if non-nil, zero value otherwise.

### GetVcpusOk

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) GetVcpusOk() (*int32, bool)`

GetVcpusOk returns a tuple with the Vcpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpus

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) SetVcpus(v int32)`

SetVcpus sets Vcpus field to given value.

### HasVcpus

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) HasVcpus() bool`

HasVcpus returns a boolean if a field has been set.

### SetVcpusNil

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) SetVcpusNil(b bool)`

 SetVcpusNil sets the value for Vcpus to be an explicit nil

### UnsetVcpus
`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) UnsetVcpus()`

UnsetVcpus ensures that no value is present for Vcpus, not even an explicit nil
### GetIsBurstable

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) GetIsBurstable() bool`

GetIsBurstable returns the IsBurstable field if non-nil, zero value otherwise.

### GetIsBurstableOk

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) GetIsBurstableOk() (*bool, bool)`

GetIsBurstableOk returns a tuple with the IsBurstable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBurstable

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) SetIsBurstable(v bool)`

SetIsBurstable sets IsBurstable field to given value.

### HasIsBurstable

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) HasIsBurstable() bool`

HasIsBurstable returns a boolean if a field has been set.

### SetIsBurstableNil

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) SetIsBurstableNil(b bool)`

 SetIsBurstableNil sets the value for IsBurstable to be an explicit nil

### UnsetIsBurstable
`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) UnsetIsBurstable()`

UnsetIsBurstable ensures that no value is present for IsBurstable, not even an explicit nil
### GetManufacturer

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### SetManufacturerNil

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) SetManufacturerNil(b bool)`

 SetManufacturerNil sets the value for Manufacturer to be an explicit nil

### UnsetManufacturer
`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) UnsetManufacturer()`

UnsetManufacturer ensures that no value is present for Manufacturer, not even an explicit nil
### GetMemoryMb

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) GetMemoryMb() int32`

GetMemoryMb returns the MemoryMb field if non-nil, zero value otherwise.

### GetMemoryMbOk

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) GetMemoryMbOk() (*int32, bool)`

GetMemoryMbOk returns a tuple with the MemoryMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMb

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) SetMemoryMb(v int32)`

SetMemoryMb sets MemoryMb field to given value.

### HasMemoryMb

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) HasMemoryMb() bool`

HasMemoryMb returns a boolean if a field has been set.

### SetMemoryMbNil

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) SetMemoryMbNil(b bool)`

 SetMemoryMbNil sets the value for MemoryMb to be an explicit nil

### UnsetMemoryMb
`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) UnsetMemoryMb()`

UnsetMemoryMb ensures that no value is present for MemoryMb, not even an explicit nil
### GetRootGb

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) GetRootGb() int32`

GetRootGb returns the RootGb field if non-nil, zero value otherwise.

### GetRootGbOk

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) GetRootGbOk() (*int32, bool)`

GetRootGbOk returns a tuple with the RootGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootGb

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) SetRootGb(v int32)`

SetRootGb sets RootGb field to given value.

### HasRootGb

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) HasRootGb() bool`

HasRootGb returns a boolean if a field has been set.

### SetRootGbNil

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) SetRootGbNil(b bool)`

 SetRootGbNil sets the value for RootGb to be an explicit nil

### UnsetRootGb
`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) UnsetRootGb()`

UnsetRootGb ensures that no value is present for RootGb, not even an explicit nil
### GetDiskType

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) GetDiskType() string`

GetDiskType returns the DiskType field if non-nil, zero value otherwise.

### GetDiskTypeOk

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) GetDiskTypeOk() (*string, bool)`

GetDiskTypeOk returns a tuple with the DiskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskType

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) SetDiskType(v string)`

SetDiskType sets DiskType field to given value.

### HasDiskType

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) HasDiskType() bool`

HasDiskType returns a boolean if a field has been set.

### SetDiskTypeNil

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) SetDiskTypeNil(b bool)`

 SetDiskTypeNil sets the value for DiskType to be an explicit nil

### UnsetDiskType
`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) UnsetDiskType()`

UnsetDiskType ensures that no value is present for DiskType, not even an explicit nil
### GetInstanceFamily

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) GetInstanceFamily() string`

GetInstanceFamily returns the InstanceFamily field if non-nil, zero value otherwise.

### GetInstanceFamilyOk

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) GetInstanceFamilyOk() (*string, bool)`

GetInstanceFamilyOk returns a tuple with the InstanceFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceFamily

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) SetInstanceFamily(v string)`

SetInstanceFamily sets InstanceFamily field to given value.

### HasInstanceFamily

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) HasInstanceFamily() bool`

HasInstanceFamily returns a boolean if a field has been set.

### SetInstanceFamilyNil

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) SetInstanceFamilyNil(b bool)`

 SetInstanceFamilyNil sets the value for InstanceFamily to be an explicit nil

### UnsetInstanceFamily
`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) UnsetInstanceFamily()`

UnsetInstanceFamily ensures that no value is present for InstanceFamily, not even an explicit nil
### GetOsDistro

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) GetOsDistro() []string`

GetOsDistro returns the OsDistro field if non-nil, zero value otherwise.

### GetOsDistroOk

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) GetOsDistroOk() (*[]string, bool)`

GetOsDistroOk returns a tuple with the OsDistro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsDistro

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) SetOsDistro(v []string)`

SetOsDistro sets OsDistro field to given value.

### HasOsDistro

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) HasOsDistro() bool`

HasOsDistro returns a boolean if a field has been set.

### SetOsDistroNil

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) SetOsDistroNil(b bool)`

 SetOsDistroNil sets the value for OsDistro to be an explicit nil

### UnsetOsDistro
`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) UnsetOsDistro()`

UnsetOsDistro ensures that no value is present for OsDistro, not even an explicit nil
### GetMaximumNetworkInterfaces

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) GetMaximumNetworkInterfaces() int32`

GetMaximumNetworkInterfaces returns the MaximumNetworkInterfaces field if non-nil, zero value otherwise.

### GetMaximumNetworkInterfacesOk

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) GetMaximumNetworkInterfacesOk() (*int32, bool)`

GetMaximumNetworkInterfacesOk returns a tuple with the MaximumNetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumNetworkInterfaces

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) SetMaximumNetworkInterfaces(v int32)`

SetMaximumNetworkInterfaces sets MaximumNetworkInterfaces field to given value.

### HasMaximumNetworkInterfaces

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) HasMaximumNetworkInterfaces() bool`

HasMaximumNetworkInterfaces returns a boolean if a field has been set.

### SetMaximumNetworkInterfacesNil

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) SetMaximumNetworkInterfacesNil(b bool)`

 SetMaximumNetworkInterfacesNil sets the value for MaximumNetworkInterfaces to be an explicit nil

### UnsetMaximumNetworkInterfaces
`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) UnsetMaximumNetworkInterfaces()`

UnsetMaximumNetworkInterfaces ensures that no value is present for MaximumNetworkInterfaces, not even an explicit nil
### GetHwType

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) GetHwType() string`

GetHwType returns the HwType field if non-nil, zero value otherwise.

### GetHwTypeOk

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) GetHwTypeOk() (*string, bool)`

GetHwTypeOk returns a tuple with the HwType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwType

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) SetHwType(v string)`

SetHwType sets HwType field to given value.

### HasHwType

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) HasHwType() bool`

HasHwType returns a boolean if a field has been set.

### SetHwTypeNil

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) SetHwTypeNil(b bool)`

 SetHwTypeNil sets the value for HwType to be an explicit nil

### UnsetHwType
`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) UnsetHwType()`

UnsetHwType ensures that no value is present for HwType, not even an explicit nil
### GetHwCount

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) GetHwCount() int32`

GetHwCount returns the HwCount field if non-nil, zero value otherwise.

### GetHwCountOk

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) GetHwCountOk() (*int32, bool)`

GetHwCountOk returns a tuple with the HwCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwCount

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) SetHwCount(v int32)`

SetHwCount sets HwCount field to given value.

### HasHwCount

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) HasHwCount() bool`

HasHwCount returns a boolean if a field has been set.

### SetHwCountNil

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) SetHwCountNil(b bool)`

 SetHwCountNil sets the value for HwCount to be an explicit nil

### UnsetHwCount
`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) UnsetHwCount()`

UnsetHwCount ensures that no value is present for HwCount, not even an explicit nil
### GetIsHyperThreadingSupported

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) GetIsHyperThreadingSupported() bool`

GetIsHyperThreadingSupported returns the IsHyperThreadingSupported field if non-nil, zero value otherwise.

### GetIsHyperThreadingSupportedOk

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) GetIsHyperThreadingSupportedOk() (*bool, bool)`

GetIsHyperThreadingSupportedOk returns a tuple with the IsHyperThreadingSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHyperThreadingSupported

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) SetIsHyperThreadingSupported(v bool)`

SetIsHyperThreadingSupported sets IsHyperThreadingSupported field to given value.

### HasIsHyperThreadingSupported

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) HasIsHyperThreadingSupported() bool`

HasIsHyperThreadingSupported returns a boolean if a field has been set.

### SetIsHyperThreadingSupportedNil

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) SetIsHyperThreadingSupportedNil(b bool)`

 SetIsHyperThreadingSupportedNil sets the value for IsHyperThreadingSupported to be an explicit nil

### UnsetIsHyperThreadingSupported
`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) UnsetIsHyperThreadingSupported()`

UnsetIsHyperThreadingSupported ensures that no value is present for IsHyperThreadingSupported, not even an explicit nil
### GetRealVcpus

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) GetRealVcpus() int32`

GetRealVcpus returns the RealVcpus field if non-nil, zero value otherwise.

### GetRealVcpusOk

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) GetRealVcpusOk() (*int32, bool)`

GetRealVcpusOk returns a tuple with the RealVcpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealVcpus

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) SetRealVcpus(v int32)`

SetRealVcpus sets RealVcpus field to given value.

### HasRealVcpus

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) HasRealVcpus() bool`

HasRealVcpus returns a boolean if a field has been set.

### SetRealVcpusNil

`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) SetRealVcpusNil(b bool)`

 SetRealVcpusNil sets the value for RealVcpus to be an explicit nil

### UnsetRealVcpus
`func (o *BcsInstanceV1ApiListInstancesModelInstanceFlavorModel) UnsetRealVcpus()`

UnsetRealVcpus ensures that no value is present for RealVcpus, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


