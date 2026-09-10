# UpdateListenerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Listener** | [**UpdateListener**](UpdateListener.md) | 수정할 리스너 정보 | 

## Methods

### NewUpdateListenerRequest

`func NewUpdateListenerRequest(listener UpdateListener, ) *UpdateListenerRequest`

NewUpdateListenerRequest instantiates a new UpdateListenerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateListenerRequestWithDefaults

`func NewUpdateListenerRequestWithDefaults() *UpdateListenerRequest`

NewUpdateListenerRequestWithDefaults instantiates a new UpdateListenerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetListener

`func (o *UpdateListenerRequest) GetListener() UpdateListener`

GetListener returns the Listener field if non-nil, zero value otherwise.

### GetListenerOk

`func (o *UpdateListenerRequest) GetListenerOk() (*UpdateListener, bool)`

GetListenerOk returns a tuple with the Listener field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListener

`func (o *UpdateListenerRequest) SetListener(v UpdateListener)`

SetListener sets Listener field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


