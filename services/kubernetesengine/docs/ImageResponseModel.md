# ImageResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Architecture** | **string** | 이미지 아키텍처 | 
**IsGpuType** | **bool** | GPU 지원 이미지 여부 | 
**Id** | **string** | 이미지의 고유 ID | 
**InstanceType** | **string** | 이미지가 지원하는 인스턴스 유형 | 
**K8sVersion** | **string** | 대상 Kubernetes 버전 | 
**KernelVersion** | **string** | 커널 버전 | 
**KeyPackage** | **string** | 이미지에 포함된 키 패키지 정보 | 
**Name** | **string** | 이미지 이름 | 
**OsDistro** | **string** | 운영체제 배포판 | 
**OsType** | **string** | 운영체제 유형 | 
**OsVersion** | **string** | 운영체제 버전 | 

## Methods

### NewImageResponseModel

`func NewImageResponseModel(architecture string, isGpuType bool, id string, instanceType string, k8sVersion string, kernelVersion string, keyPackage string, name string, osDistro string, osType string, osVersion string, ) *ImageResponseModel`

NewImageResponseModel instantiates a new ImageResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageResponseModelWithDefaults

`func NewImageResponseModelWithDefaults() *ImageResponseModel`

NewImageResponseModelWithDefaults instantiates a new ImageResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchitecture

`func (o *ImageResponseModel) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *ImageResponseModel) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *ImageResponseModel) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.


### GetIsGpuType

`func (o *ImageResponseModel) GetIsGpuType() bool`

GetIsGpuType returns the IsGpuType field if non-nil, zero value otherwise.

### GetIsGpuTypeOk

`func (o *ImageResponseModel) GetIsGpuTypeOk() (*bool, bool)`

GetIsGpuTypeOk returns a tuple with the IsGpuType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGpuType

`func (o *ImageResponseModel) SetIsGpuType(v bool)`

SetIsGpuType sets IsGpuType field to given value.


### GetId

`func (o *ImageResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImageResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImageResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetInstanceType

`func (o *ImageResponseModel) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *ImageResponseModel) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *ImageResponseModel) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.


### GetK8sVersion

`func (o *ImageResponseModel) GetK8sVersion() string`

GetK8sVersion returns the K8sVersion field if non-nil, zero value otherwise.

### GetK8sVersionOk

`func (o *ImageResponseModel) GetK8sVersionOk() (*string, bool)`

GetK8sVersionOk returns a tuple with the K8sVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sVersion

`func (o *ImageResponseModel) SetK8sVersion(v string)`

SetK8sVersion sets K8sVersion field to given value.


### GetKernelVersion

`func (o *ImageResponseModel) GetKernelVersion() string`

GetKernelVersion returns the KernelVersion field if non-nil, zero value otherwise.

### GetKernelVersionOk

`func (o *ImageResponseModel) GetKernelVersionOk() (*string, bool)`

GetKernelVersionOk returns a tuple with the KernelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKernelVersion

`func (o *ImageResponseModel) SetKernelVersion(v string)`

SetKernelVersion sets KernelVersion field to given value.


### GetKeyPackage

`func (o *ImageResponseModel) GetKeyPackage() string`

GetKeyPackage returns the KeyPackage field if non-nil, zero value otherwise.

### GetKeyPackageOk

`func (o *ImageResponseModel) GetKeyPackageOk() (*string, bool)`

GetKeyPackageOk returns a tuple with the KeyPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPackage

`func (o *ImageResponseModel) SetKeyPackage(v string)`

SetKeyPackage sets KeyPackage field to given value.


### GetName

`func (o *ImageResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImageResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImageResponseModel) SetName(v string)`

SetName sets Name field to given value.


### GetOsDistro

`func (o *ImageResponseModel) GetOsDistro() string`

GetOsDistro returns the OsDistro field if non-nil, zero value otherwise.

### GetOsDistroOk

`func (o *ImageResponseModel) GetOsDistroOk() (*string, bool)`

GetOsDistroOk returns a tuple with the OsDistro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsDistro

`func (o *ImageResponseModel) SetOsDistro(v string)`

SetOsDistro sets OsDistro field to given value.


### GetOsType

`func (o *ImageResponseModel) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *ImageResponseModel) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *ImageResponseModel) SetOsType(v string)`

SetOsType sets OsType field to given value.


### GetOsVersion

`func (o *ImageResponseModel) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *ImageResponseModel) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *ImageResponseModel) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


