# CreateListenerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Listener** | [**CreateListener**](CreateListener.md) | 생성할 리스너 정보 | 

## Methods

### NewCreateListenerRequest

`func NewCreateListenerRequest(listener CreateListener, ) *CreateListenerRequest`

NewCreateListenerRequest instantiates a new CreateListenerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateListenerRequestWithDefaults

`func NewCreateListenerRequestWithDefaults() *CreateListenerRequest`

NewCreateListenerRequestWithDefaults instantiates a new CreateListenerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetListener

`func (o *CreateListenerRequest) GetListener() CreateListener`

GetListener returns the Listener field if non-nil, zero value otherwise.

### GetListenerOk

`func (o *CreateListenerRequest) GetListenerOk() (*CreateListener, bool)`

GetListenerOk returns a tuple with the Listener field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListener

`func (o *CreateListenerRequest) SetListener(v CreateListener)`

SetListener sets Listener field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


