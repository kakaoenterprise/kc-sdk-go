# ResizeInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 변경할 인스턴스 유형 ID - [List instance types](/openapi/bcs/list-instance-types)에서 확인 - [인스턴스 유형별 사양](https://docs.kakaocloud.com/service/bcs/bcs-specifications) 참고 | 
**IsDisableHyperThreading** | Pointer to **NullableBool** | 하이퍼스레딩 비활성화 여부 - &#x60;true&#x60;: 하이퍼스레딩 비활성화 - &#x60;false&#x60;: 하이퍼스레딩 활성화 | [optional] 

## Methods

### NewResizeInstance

`func NewResizeInstance(id string, ) *ResizeInstance`

NewResizeInstance instantiates a new ResizeInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResizeInstanceWithDefaults

`func NewResizeInstanceWithDefaults() *ResizeInstance`

NewResizeInstanceWithDefaults instantiates a new ResizeInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResizeInstance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResizeInstance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResizeInstance) SetId(v string)`

SetId sets Id field to given value.


### GetIsDisableHyperThreading

`func (o *ResizeInstance) GetIsDisableHyperThreading() bool`

GetIsDisableHyperThreading returns the IsDisableHyperThreading field if non-nil, zero value otherwise.

### GetIsDisableHyperThreadingOk

`func (o *ResizeInstance) GetIsDisableHyperThreadingOk() (*bool, bool)`

GetIsDisableHyperThreadingOk returns a tuple with the IsDisableHyperThreading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisableHyperThreading

`func (o *ResizeInstance) SetIsDisableHyperThreading(v bool)`

SetIsDisableHyperThreading sets IsDisableHyperThreading field to given value.

### HasIsDisableHyperThreading

`func (o *ResizeInstance) HasIsDisableHyperThreading() bool`

HasIsDisableHyperThreading returns a boolean if a field has been set.

### SetIsDisableHyperThreadingNil

`func (o *ResizeInstance) SetIsDisableHyperThreadingNil(b bool)`

 SetIsDisableHyperThreadingNil sets the value for IsDisableHyperThreading to be an explicit nil

### UnsetIsDisableHyperThreading
`func (o *ResizeInstance) UnsetIsDisableHyperThreading()`

UnsetIsDisableHyperThreading ensures that no value is present for IsDisableHyperThreading, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


