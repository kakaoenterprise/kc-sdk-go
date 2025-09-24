# PodResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationTimestamp** | **string** | 파드 생성 시간(ISO 8601, UTC) | 
**Ip** | **string** | 파드의 IP 주소 | 
**Name** | **string** | 파드 이름 | 
**Namespace** | **string** | 파드의 네임스페이스 | 
**Status** | **string** | 파드 상태 | 

## Methods

### NewPodResponseModel

`func NewPodResponseModel(creationTimestamp string, ip string, name string, namespace string, status string, ) *PodResponseModel`

NewPodResponseModel instantiates a new PodResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPodResponseModelWithDefaults

`func NewPodResponseModelWithDefaults() *PodResponseModel`

NewPodResponseModelWithDefaults instantiates a new PodResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationTimestamp

`func (o *PodResponseModel) GetCreationTimestamp() string`

GetCreationTimestamp returns the CreationTimestamp field if non-nil, zero value otherwise.

### GetCreationTimestampOk

`func (o *PodResponseModel) GetCreationTimestampOk() (*string, bool)`

GetCreationTimestampOk returns a tuple with the CreationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTimestamp

`func (o *PodResponseModel) SetCreationTimestamp(v string)`

SetCreationTimestamp sets CreationTimestamp field to given value.


### GetIp

`func (o *PodResponseModel) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *PodResponseModel) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *PodResponseModel) SetIp(v string)`

SetIp sets Ip field to given value.


### GetName

`func (o *PodResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PodResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PodResponseModel) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *PodResponseModel) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *PodResponseModel) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *PodResponseModel) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetStatus

`func (o *PodResponseModel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PodResponseModel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PodResponseModel) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


