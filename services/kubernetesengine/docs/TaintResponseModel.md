# TaintResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | 테인트 키 | 
**Value** | **string** | 테인트 값 | 
**Effect** | [**NodePoolTaintEffect**](NodePoolTaintEffect.md) | 테인트(해당 노드에 파드가 스케줄되지 않도록 하는 정책) 효과 &lt;br/&gt; - &#x60;NoExecute&#x60;: 새로운 파드는 스케줄되지 않고, 기존에 실행 중인 파드도 축출(evict)됨  &lt;br/&gt; - &#x60;NoSchedule&#x60;: 새로운 파드는 해당 노드에 스케줄되지 않음 (이미 실행 중인 파드는 영향 없음)  &lt;br/&gt; - &#x60;PreferNoSchedule&#x60;: 가능한 경우 파드를 스케줄하지 않지만, 꼭 필요하면 스케줄될 수 있음(소프트 제약) | 

## Methods

### NewTaintResponseModel

`func NewTaintResponseModel(key string, value string, effect NodePoolTaintEffect, ) *TaintResponseModel`

NewTaintResponseModel instantiates a new TaintResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaintResponseModelWithDefaults

`func NewTaintResponseModelWithDefaults() *TaintResponseModel`

NewTaintResponseModelWithDefaults instantiates a new TaintResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *TaintResponseModel) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TaintResponseModel) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TaintResponseModel) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *TaintResponseModel) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TaintResponseModel) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TaintResponseModel) SetValue(v string)`

SetValue sets Value field to given value.


### GetEffect

`func (o *TaintResponseModel) GetEffect() NodePoolTaintEffect`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *TaintResponseModel) GetEffectOk() (*NodePoolTaintEffect, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *TaintResponseModel) SetEffect(v NodePoolTaintEffect)`

SetEffect sets Effect field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


