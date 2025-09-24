# KubernetesEngineV1ApiListNodePoolsModelTaintInfoResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Effect** | **string** | 테인트(해당 노드에 파드가 스케줄되지 않도록 하는 정책) 효과 | 
**Key** | **string** | 테인트의 고유 키 | 
**Value** | **string** | 테인트 값 | 

## Methods

### NewKubernetesEngineV1ApiListNodePoolsModelTaintInfoResponseModel

`func NewKubernetesEngineV1ApiListNodePoolsModelTaintInfoResponseModel(effect string, key string, value string, ) *KubernetesEngineV1ApiListNodePoolsModelTaintInfoResponseModel`

NewKubernetesEngineV1ApiListNodePoolsModelTaintInfoResponseModel instantiates a new KubernetesEngineV1ApiListNodePoolsModelTaintInfoResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesEngineV1ApiListNodePoolsModelTaintInfoResponseModelWithDefaults

`func NewKubernetesEngineV1ApiListNodePoolsModelTaintInfoResponseModelWithDefaults() *KubernetesEngineV1ApiListNodePoolsModelTaintInfoResponseModel`

NewKubernetesEngineV1ApiListNodePoolsModelTaintInfoResponseModelWithDefaults instantiates a new KubernetesEngineV1ApiListNodePoolsModelTaintInfoResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEffect

`func (o *KubernetesEngineV1ApiListNodePoolsModelTaintInfoResponseModel) GetEffect() string`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *KubernetesEngineV1ApiListNodePoolsModelTaintInfoResponseModel) GetEffectOk() (*string, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *KubernetesEngineV1ApiListNodePoolsModelTaintInfoResponseModel) SetEffect(v string)`

SetEffect sets Effect field to given value.


### GetKey

`func (o *KubernetesEngineV1ApiListNodePoolsModelTaintInfoResponseModel) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *KubernetesEngineV1ApiListNodePoolsModelTaintInfoResponseModel) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *KubernetesEngineV1ApiListNodePoolsModelTaintInfoResponseModel) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *KubernetesEngineV1ApiListNodePoolsModelTaintInfoResponseModel) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *KubernetesEngineV1ApiListNodePoolsModelTaintInfoResponseModel) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *KubernetesEngineV1ApiListNodePoolsModelTaintInfoResponseModel) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


