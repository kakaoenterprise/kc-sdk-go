# ImageInfoResponseModel

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

### NewImageInfoResponseModel

`func NewImageInfoResponseModel(architecture string, isGpuType bool, id string, instanceType string, kernelVersion string, keyPackage string, name string, osDistro string, osType string, osVersion string, ) *ImageInfoResponseModel`

NewImageInfoResponseModel instantiates a new ImageInfoResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageInfoResponseModelWithDefaults

`func NewImageInfoResponseModelWithDefaults() *ImageInfoResponseModel`

NewImageInfoResponseModelWithDefaults instantiates a new ImageInfoResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchitecture

`func (o *ImageInfoResponseModel) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *ImageInfoResponseModel) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *ImageInfoResponseModel) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.


### GetIsGpuType

`func (o *ImageInfoResponseModel) GetIsGpuType() bool`

GetIsGpuType returns the IsGpuType field if non-nil, zero value otherwise.

### GetIsGpuTypeOk

`func (o *ImageInfoResponseModel) GetIsGpuTypeOk() (*bool, bool)`

GetIsGpuTypeOk returns a tuple with the IsGpuType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGpuType

`func (o *ImageInfoResponseModel) SetIsGpuType(v bool)`

SetIsGpuType sets IsGpuType field to given value.


### GetId

`func (o *ImageInfoResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImageInfoResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImageInfoResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetInstanceType

`func (o *ImageInfoResponseModel) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *ImageInfoResponseModel) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *ImageInfoResponseModel) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.


### GetKernelVersion

`func (o *ImageInfoResponseModel) GetKernelVersion() string`

GetKernelVersion returns the KernelVersion field if non-nil, zero value otherwise.

### GetKernelVersionOk

`func (o *ImageInfoResponseModel) GetKernelVersionOk() (*string, bool)`

GetKernelVersionOk returns a tuple with the KernelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKernelVersion

`func (o *ImageInfoResponseModel) SetKernelVersion(v string)`

SetKernelVersion sets KernelVersion field to given value.


### GetKeyPackage

`func (o *ImageInfoResponseModel) GetKeyPackage() string`

GetKeyPackage returns the KeyPackage field if non-nil, zero value otherwise.

### GetKeyPackageOk

`func (o *ImageInfoResponseModel) GetKeyPackageOk() (*string, bool)`

GetKeyPackageOk returns a tuple with the KeyPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPackage

`func (o *ImageInfoResponseModel) SetKeyPackage(v string)`

SetKeyPackage sets KeyPackage field to given value.


### GetName

`func (o *ImageInfoResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImageInfoResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImageInfoResponseModel) SetName(v string)`

SetName sets Name field to given value.


### GetOsDistro

`func (o *ImageInfoResponseModel) GetOsDistro() string`

GetOsDistro returns the OsDistro field if non-nil, zero value otherwise.

### GetOsDistroOk

`func (o *ImageInfoResponseModel) GetOsDistroOk() (*string, bool)`

GetOsDistroOk returns a tuple with the OsDistro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsDistro

`func (o *ImageInfoResponseModel) SetOsDistro(v string)`

SetOsDistro sets OsDistro field to given value.


### GetOsType

`func (o *ImageInfoResponseModel) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *ImageInfoResponseModel) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *ImageInfoResponseModel) SetOsType(v string)`

SetOsType sets OsType field to given value.


### GetOsVersion

`func (o *ImageInfoResponseModel) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *ImageInfoResponseModel) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *ImageInfoResponseModel) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


