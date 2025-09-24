# KubernetesEngineV1ApiGetClusterNodeDetailsModelStatusResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | [**[]AddressResponseModel**](AddressResponseModel.md) | 노드 주소 목록 | 
**Allocatable** | [**AllocatableResourcesResponseModel**](AllocatableResourcesResponseModel.md) | 노드에서 사용 가능한 리소스 | 
**Conditions** | [**[]ConditionResponseModel**](ConditionResponseModel.md) | 노드 status 조건 목록(예: Ready, DiskPressure)  &lt;!---status 영문 유지---&gt; | 

## Methods

### NewKubernetesEngineV1ApiGetClusterNodeDetailsModelStatusResponseModel

`func NewKubernetesEngineV1ApiGetClusterNodeDetailsModelStatusResponseModel(addresses []AddressResponseModel, allocatable AllocatableResourcesResponseModel, conditions []ConditionResponseModel, ) *KubernetesEngineV1ApiGetClusterNodeDetailsModelStatusResponseModel`

NewKubernetesEngineV1ApiGetClusterNodeDetailsModelStatusResponseModel instantiates a new KubernetesEngineV1ApiGetClusterNodeDetailsModelStatusResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesEngineV1ApiGetClusterNodeDetailsModelStatusResponseModelWithDefaults

`func NewKubernetesEngineV1ApiGetClusterNodeDetailsModelStatusResponseModelWithDefaults() *KubernetesEngineV1ApiGetClusterNodeDetailsModelStatusResponseModel`

NewKubernetesEngineV1ApiGetClusterNodeDetailsModelStatusResponseModelWithDefaults instantiates a new KubernetesEngineV1ApiGetClusterNodeDetailsModelStatusResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *KubernetesEngineV1ApiGetClusterNodeDetailsModelStatusResponseModel) GetAddresses() []AddressResponseModel`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *KubernetesEngineV1ApiGetClusterNodeDetailsModelStatusResponseModel) GetAddressesOk() (*[]AddressResponseModel, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *KubernetesEngineV1ApiGetClusterNodeDetailsModelStatusResponseModel) SetAddresses(v []AddressResponseModel)`

SetAddresses sets Addresses field to given value.


### GetAllocatable

`func (o *KubernetesEngineV1ApiGetClusterNodeDetailsModelStatusResponseModel) GetAllocatable() AllocatableResourcesResponseModel`

GetAllocatable returns the Allocatable field if non-nil, zero value otherwise.

### GetAllocatableOk

`func (o *KubernetesEngineV1ApiGetClusterNodeDetailsModelStatusResponseModel) GetAllocatableOk() (*AllocatableResourcesResponseModel, bool)`

GetAllocatableOk returns a tuple with the Allocatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatable

`func (o *KubernetesEngineV1ApiGetClusterNodeDetailsModelStatusResponseModel) SetAllocatable(v AllocatableResourcesResponseModel)`

SetAllocatable sets Allocatable field to given value.


### GetConditions

`func (o *KubernetesEngineV1ApiGetClusterNodeDetailsModelStatusResponseModel) GetConditions() []ConditionResponseModel`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *KubernetesEngineV1ApiGetClusterNodeDetailsModelStatusResponseModel) GetConditionsOk() (*[]ConditionResponseModel, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *KubernetesEngineV1ApiGetClusterNodeDetailsModelStatusResponseModel) SetConditions(v []ConditionResponseModel)`

SetConditions sets Conditions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


