# ImageInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Architecture** | **string** | 이미지 아키텍처 | 
**IsGpuType** | **bool** | GPU 지원 이미지 여부 | 
**Id** | **string** | 이미지의 고유 ID | 
**InstanceType** | **string** | 이미지가 지원하는 인스턴스 유형 | 
**KernelVersion** | **string** | 커널 버전 | 
**KeyPackage** | **string** | 이미지에 포함된 키 패키지 정보 | 
**Name** | **string** | 이미지 이름 | 
**OsDistro** | **string** | 운영체제 배포판 | 
**OsType** | **string** | 운영체제 유형 | 
**OsVersion** | **string** | 운영체제 버전 | 

## Methods

### NewImageInfo

`func NewImageInfo(architecture string, isGpuType bool, id string, instanceType string, kernelVersion string, keyPackage string, name string, osDistro string, osType string, osVersion string, ) *ImageInfo`

NewImageInfo instantiates a new ImageInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageInfoWithDefaults

`func NewImageInfoWithDefaults() *ImageInfo`

NewImageInfoWithDefaults instantiates a new ImageInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchitecture

`func (o *ImageInfo) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *ImageInfo) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *ImageInfo) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.


### GetIsGpuType

`func (o *ImageInfo) GetIsGpuType() bool`

GetIsGpuType returns the IsGpuType field if non-nil, zero value otherwise.

### GetIsGpuTypeOk

`func (o *ImageInfo) GetIsGpuTypeOk() (*bool, bool)`

GetIsGpuTypeOk returns a tuple with the IsGpuType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGpuType

`func (o *ImageInfo) SetIsGpuType(v bool)`

SetIsGpuType sets IsGpuType field to given value.


### GetId

`func (o *ImageInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImageInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImageInfo) SetId(v string)`

SetId sets Id field to given value.


### GetInstanceType

`func (o *ImageInfo) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *ImageInfo) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *ImageInfo) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.


### GetKernelVersion

`func (o *ImageInfo) GetKernelVersion() string`

GetKernelVersion returns the KernelVersion field if non-nil, zero value otherwise.

### GetKernelVersionOk

`func (o *ImageInfo) GetKernelVersionOk() (*string, bool)`

GetKernelVersionOk returns a tuple with the KernelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKernelVersion

`func (o *ImageInfo) SetKernelVersion(v string)`

SetKernelVersion sets KernelVersion field to given value.


### GetKeyPackage

`func (o *ImageInfo) GetKeyPackage() string`

GetKeyPackage returns the KeyPackage field if non-nil, zero value otherwise.

### GetKeyPackageOk

`func (o *ImageInfo) GetKeyPackageOk() (*string, bool)`

GetKeyPackageOk returns a tuple with the KeyPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPackage

`func (o *ImageInfo) SetKeyPackage(v string)`

SetKeyPackage sets KeyPackage field to given value.


### GetName

`func (o *ImageInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImageInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImageInfo) SetName(v string)`

SetName sets Name field to given value.


### GetOsDistro

`func (o *ImageInfo) GetOsDistro() string`

GetOsDistro returns the OsDistro field if non-nil, zero value otherwise.

### GetOsDistroOk

`func (o *ImageInfo) GetOsDistroOk() (*string, bool)`

GetOsDistroOk returns a tuple with the OsDistro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsDistro

`func (o *ImageInfo) SetOsDistro(v string)`

SetOsDistro sets OsDistro field to given value.


### GetOsType

`func (o *ImageInfo) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *ImageInfo) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *ImageInfo) SetOsType(v string)`

SetOsType sets OsType field to given value.


### GetOsVersion

`func (o *ImageInfo) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *ImageInfo) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *ImageInfo) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


