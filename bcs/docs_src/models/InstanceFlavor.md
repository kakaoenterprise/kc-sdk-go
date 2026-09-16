# InstanceFlavor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | 인스턴스 유형 ID | [optional] 
**Name** | Pointer to **NullableString** | 인스턴스 유형 이름 | [optional] 
**Group** | Pointer to **NullableString** | 인스턴스 그룹 (성능, 요금 기준 구분) - 예시:  &#x60;cpu_optimized&#x60;, &#x60;memory_optimized&#x60;, &#x60;gpu&#x60; 등 | [optional] 
**Vcpus** | Pointer to **NullableInt32** | 가상 CPU(Virtual CPU) 수 | [optional] 
**IsBurstable** | Pointer to **NullableBool** | [버스터블 인스턴스](https://docs.kakaocloud.com/service/bcs/bcs-specifications/general-purpose/burstable-main) 여부 | [optional] 
**Manufacturer** | Pointer to **NullableString** | 하드웨어 제조사 | [optional] 
**MemoryMb** | Pointer to **NullableInt32** | 메모리 크기 (MB 단위) | [optional] 
**RootGb** | Pointer to **NullableInt32** | 루트 디스크의 크기 (GB 단위) | [optional] 
**DiskType** | Pointer to **NullableString** | 디스크 유형 | [optional] 
**InstanceFamily** | Pointer to **NullableString** | 인스턴스 계열 | [optional] 
**OsDistro** | Pointer to **[]string** | 운영체제 배포판 | [optional] 
**MaximumNetworkInterfaces** | Pointer to **NullableInt32** | 연결 가능한 최대 네트워크 인터페이스 수 | [optional] 
**HwType** | Pointer to **NullableString** | 하드웨어 유형 - 예시: GPU | [optional] 
**HwCount** | Pointer to **NullableInt32** | 해당 인스턴스 유형에 포함된 하드웨어 개수 - 예시: GPU 개수 | [optional] 
**IsHyperThreadingSupported** | Pointer to **NullableBool** | 하이퍼스레딩 지원 여부 | [optional] 
**RealVcpus** | Pointer to **NullableInt32** | 실제 사용 가능한 가상 CPU 수 | [optional] 

## Methods

### NewInstanceFlavor

`func NewInstanceFlavor() *InstanceFlavor`

NewInstanceFlavor instantiates a new InstanceFlavor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceFlavorWithDefaults

`func NewInstanceFlavorWithDefaults() *InstanceFlavor`

NewInstanceFlavorWithDefaults instantiates a new InstanceFlavor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstanceFlavor) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceFlavor) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceFlavor) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InstanceFlavor) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *InstanceFlavor) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *InstanceFlavor) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *InstanceFlavor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceFlavor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceFlavor) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InstanceFlavor) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *InstanceFlavor) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *InstanceFlavor) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetGroup

`func (o *InstanceFlavor) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *InstanceFlavor) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *InstanceFlavor) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *InstanceFlavor) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *InstanceFlavor) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *InstanceFlavor) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetVcpus

`func (o *InstanceFlavor) GetVcpus() int32`

GetVcpus returns the Vcpus field if non-nil, zero value otherwise.

### GetVcpusOk

`func (o *InstanceFlavor) GetVcpusOk() (*int32, bool)`

GetVcpusOk returns a tuple with the Vcpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpus

`func (o *InstanceFlavor) SetVcpus(v int32)`

SetVcpus sets Vcpus field to given value.

### HasVcpus

`func (o *InstanceFlavor) HasVcpus() bool`

HasVcpus returns a boolean if a field has been set.

### SetVcpusNil

`func (o *InstanceFlavor) SetVcpusNil(b bool)`

 SetVcpusNil sets the value for Vcpus to be an explicit nil

### UnsetVcpus
`func (o *InstanceFlavor) UnsetVcpus()`

UnsetVcpus ensures that no value is present for Vcpus, not even an explicit nil
### GetIsBurstable

`func (o *InstanceFlavor) GetIsBurstable() bool`

GetIsBurstable returns the IsBurstable field if non-nil, zero value otherwise.

### GetIsBurstableOk

`func (o *InstanceFlavor) GetIsBurstableOk() (*bool, bool)`

GetIsBurstableOk returns a tuple with the IsBurstable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBurstable

`func (o *InstanceFlavor) SetIsBurstable(v bool)`

SetIsBurstable sets IsBurstable field to given value.

### HasIsBurstable

`func (o *InstanceFlavor) HasIsBurstable() bool`

HasIsBurstable returns a boolean if a field has been set.

### SetIsBurstableNil

`func (o *InstanceFlavor) SetIsBurstableNil(b bool)`

 SetIsBurstableNil sets the value for IsBurstable to be an explicit nil

### UnsetIsBurstable
`func (o *InstanceFlavor) UnsetIsBurstable()`

UnsetIsBurstable ensures that no value is present for IsBurstable, not even an explicit nil
### GetManufacturer

`func (o *InstanceFlavor) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *InstanceFlavor) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *InstanceFlavor) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *InstanceFlavor) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### SetManufacturerNil

`func (o *InstanceFlavor) SetManufacturerNil(b bool)`

 SetManufacturerNil sets the value for Manufacturer to be an explicit nil

### UnsetManufacturer
`func (o *InstanceFlavor) UnsetManufacturer()`

UnsetManufacturer ensures that no value is present for Manufacturer, not even an explicit nil
### GetMemoryMb

`func (o *InstanceFlavor) GetMemoryMb() int32`

GetMemoryMb returns the MemoryMb field if non-nil, zero value otherwise.

### GetMemoryMbOk

`func (o *InstanceFlavor) GetMemoryMbOk() (*int32, bool)`

GetMemoryMbOk returns a tuple with the MemoryMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMb

`func (o *InstanceFlavor) SetMemoryMb(v int32)`

SetMemoryMb sets MemoryMb field to given value.

### HasMemoryMb

`func (o *InstanceFlavor) HasMemoryMb() bool`

HasMemoryMb returns a boolean if a field has been set.

### SetMemoryMbNil

`func (o *InstanceFlavor) SetMemoryMbNil(b bool)`

 SetMemoryMbNil sets the value for MemoryMb to be an explicit nil

### UnsetMemoryMb
`func (o *InstanceFlavor) UnsetMemoryMb()`

UnsetMemoryMb ensures that no value is present for MemoryMb, not even an explicit nil
### GetRootGb

`func (o *InstanceFlavor) GetRootGb() int32`

GetRootGb returns the RootGb field if non-nil, zero value otherwise.

### GetRootGbOk

`func (o *InstanceFlavor) GetRootGbOk() (*int32, bool)`

GetRootGbOk returns a tuple with the RootGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootGb

`func (o *InstanceFlavor) SetRootGb(v int32)`

SetRootGb sets RootGb field to given value.

### HasRootGb

`func (o *InstanceFlavor) HasRootGb() bool`

HasRootGb returns a boolean if a field has been set.

### SetRootGbNil

`func (o *InstanceFlavor) SetRootGbNil(b bool)`

 SetRootGbNil sets the value for RootGb to be an explicit nil

### UnsetRootGb
`func (o *InstanceFlavor) UnsetRootGb()`

UnsetRootGb ensures that no value is present for RootGb, not even an explicit nil
### GetDiskType

`func (o *InstanceFlavor) GetDiskType() string`

GetDiskType returns the DiskType field if non-nil, zero value otherwise.

### GetDiskTypeOk

`func (o *InstanceFlavor) GetDiskTypeOk() (*string, bool)`

GetDiskTypeOk returns a tuple with the DiskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskType

`func (o *InstanceFlavor) SetDiskType(v string)`

SetDiskType sets DiskType field to given value.

### HasDiskType

`func (o *InstanceFlavor) HasDiskType() bool`

HasDiskType returns a boolean if a field has been set.

### SetDiskTypeNil

`func (o *InstanceFlavor) SetDiskTypeNil(b bool)`

 SetDiskTypeNil sets the value for DiskType to be an explicit nil

### UnsetDiskType
`func (o *InstanceFlavor) UnsetDiskType()`

UnsetDiskType ensures that no value is present for DiskType, not even an explicit nil
### GetInstanceFamily

`func (o *InstanceFlavor) GetInstanceFamily() string`

GetInstanceFamily returns the InstanceFamily field if non-nil, zero value otherwise.

### GetInstanceFamilyOk

`func (o *InstanceFlavor) GetInstanceFamilyOk() (*string, bool)`

GetInstanceFamilyOk returns a tuple with the InstanceFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceFamily

`func (o *InstanceFlavor) SetInstanceFamily(v string)`

SetInstanceFamily sets InstanceFamily field to given value.

### HasInstanceFamily

`func (o *InstanceFlavor) HasInstanceFamily() bool`

HasInstanceFamily returns a boolean if a field has been set.

### SetInstanceFamilyNil

`func (o *InstanceFlavor) SetInstanceFamilyNil(b bool)`

 SetInstanceFamilyNil sets the value for InstanceFamily to be an explicit nil

### UnsetInstanceFamily
`func (o *InstanceFlavor) UnsetInstanceFamily()`

UnsetInstanceFamily ensures that no value is present for InstanceFamily, not even an explicit nil
### GetOsDistro

`func (o *InstanceFlavor) GetOsDistro() []string`

GetOsDistro returns the OsDistro field if non-nil, zero value otherwise.

### GetOsDistroOk

`func (o *InstanceFlavor) GetOsDistroOk() (*[]string, bool)`

GetOsDistroOk returns a tuple with the OsDistro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsDistro

`func (o *InstanceFlavor) SetOsDistro(v []string)`

SetOsDistro sets OsDistro field to given value.

### HasOsDistro

`func (o *InstanceFlavor) HasOsDistro() bool`

HasOsDistro returns a boolean if a field has been set.

### SetOsDistroNil

`func (o *InstanceFlavor) SetOsDistroNil(b bool)`

 SetOsDistroNil sets the value for OsDistro to be an explicit nil

### UnsetOsDistro
`func (o *InstanceFlavor) UnsetOsDistro()`

UnsetOsDistro ensures that no value is present for OsDistro, not even an explicit nil
### GetMaximumNetworkInterfaces

`func (o *InstanceFlavor) GetMaximumNetworkInterfaces() int32`

GetMaximumNetworkInterfaces returns the MaximumNetworkInterfaces field if non-nil, zero value otherwise.

### GetMaximumNetworkInterfacesOk

`func (o *InstanceFlavor) GetMaximumNetworkInterfacesOk() (*int32, bool)`

GetMaximumNetworkInterfacesOk returns a tuple with the MaximumNetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumNetworkInterfaces

`func (o *InstanceFlavor) SetMaximumNetworkInterfaces(v int32)`

SetMaximumNetworkInterfaces sets MaximumNetworkInterfaces field to given value.

### HasMaximumNetworkInterfaces

`func (o *InstanceFlavor) HasMaximumNetworkInterfaces() bool`

HasMaximumNetworkInterfaces returns a boolean if a field has been set.

### SetMaximumNetworkInterfacesNil

`func (o *InstanceFlavor) SetMaximumNetworkInterfacesNil(b bool)`

 SetMaximumNetworkInterfacesNil sets the value for MaximumNetworkInterfaces to be an explicit nil

### UnsetMaximumNetworkInterfaces
`func (o *InstanceFlavor) UnsetMaximumNetworkInterfaces()`

UnsetMaximumNetworkInterfaces ensures that no value is present for MaximumNetworkInterfaces, not even an explicit nil
### GetHwType

`func (o *InstanceFlavor) GetHwType() string`

GetHwType returns the HwType field if non-nil, zero value otherwise.

### GetHwTypeOk

`func (o *InstanceFlavor) GetHwTypeOk() (*string, bool)`

GetHwTypeOk returns a tuple with the HwType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwType

`func (o *InstanceFlavor) SetHwType(v string)`

SetHwType sets HwType field to given value.

### HasHwType

`func (o *InstanceFlavor) HasHwType() bool`

HasHwType returns a boolean if a field has been set.

### SetHwTypeNil

`func (o *InstanceFlavor) SetHwTypeNil(b bool)`

 SetHwTypeNil sets the value for HwType to be an explicit nil

### UnsetHwType
`func (o *InstanceFlavor) UnsetHwType()`

UnsetHwType ensures that no value is present for HwType, not even an explicit nil
### GetHwCount

`func (o *InstanceFlavor) GetHwCount() int32`

GetHwCount returns the HwCount field if non-nil, zero value otherwise.

### GetHwCountOk

`func (o *InstanceFlavor) GetHwCountOk() (*int32, bool)`

GetHwCountOk returns a tuple with the HwCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwCount

`func (o *InstanceFlavor) SetHwCount(v int32)`

SetHwCount sets HwCount field to given value.

### HasHwCount

`func (o *InstanceFlavor) HasHwCount() bool`

HasHwCount returns a boolean if a field has been set.

### SetHwCountNil

`func (o *InstanceFlavor) SetHwCountNil(b bool)`

 SetHwCountNil sets the value for HwCount to be an explicit nil

### UnsetHwCount
`func (o *InstanceFlavor) UnsetHwCount()`

UnsetHwCount ensures that no value is present for HwCount, not even an explicit nil
### GetIsHyperThreadingSupported

`func (o *InstanceFlavor) GetIsHyperThreadingSupported() bool`

GetIsHyperThreadingSupported returns the IsHyperThreadingSupported field if non-nil, zero value otherwise.

### GetIsHyperThreadingSupportedOk

`func (o *InstanceFlavor) GetIsHyperThreadingSupportedOk() (*bool, bool)`

GetIsHyperThreadingSupportedOk returns a tuple with the IsHyperThreadingSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHyperThreadingSupported

`func (o *InstanceFlavor) SetIsHyperThreadingSupported(v bool)`

SetIsHyperThreadingSupported sets IsHyperThreadingSupported field to given value.

### HasIsHyperThreadingSupported

`func (o *InstanceFlavor) HasIsHyperThreadingSupported() bool`

HasIsHyperThreadingSupported returns a boolean if a field has been set.

### SetIsHyperThreadingSupportedNil

`func (o *InstanceFlavor) SetIsHyperThreadingSupportedNil(b bool)`

 SetIsHyperThreadingSupportedNil sets the value for IsHyperThreadingSupported to be an explicit nil

### UnsetIsHyperThreadingSupported
`func (o *InstanceFlavor) UnsetIsHyperThreadingSupported()`

UnsetIsHyperThreadingSupported ensures that no value is present for IsHyperThreadingSupported, not even an explicit nil
### GetRealVcpus

`func (o *InstanceFlavor) GetRealVcpus() int32`

GetRealVcpus returns the RealVcpus field if non-nil, zero value otherwise.

### GetRealVcpusOk

`func (o *InstanceFlavor) GetRealVcpusOk() (*int32, bool)`

GetRealVcpusOk returns a tuple with the RealVcpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealVcpus

`func (o *InstanceFlavor) SetRealVcpus(v int32)`

SetRealVcpus sets RealVcpus field to given value.

### HasRealVcpus

`func (o *InstanceFlavor) HasRealVcpus() bool`

HasRealVcpus returns a boolean if a field has been set.

### SetRealVcpusNil

`func (o *InstanceFlavor) SetRealVcpusNil(b bool)`

 SetRealVcpusNil sets the value for RealVcpus to be an explicit nil

### UnsetRealVcpus
`func (o *InstanceFlavor) UnsetRealVcpus()`

UnsetRealVcpus ensures that no value is present for RealVcpus, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


