# GetListenerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Listener** | [**LoadBalancerListener**](LoadBalancerListener.md) | 리스너 ID | 

## Methods

### NewGetListenerResponse

`func NewGetListenerResponse(listener LoadBalancerListener, ) *GetListenerResponse`

NewGetListenerResponse instantiates a new GetListenerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetListenerResponseWithDefaults

`func NewGetListenerResponseWithDefaults() *GetListenerResponse`

NewGetListenerResponseWithDefaults instantiates a new GetListenerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetListener

`func (o *GetListenerResponse) GetListener() LoadBalancerListener`

GetListener returns the Listener field if non-nil, zero value otherwise.

### GetListenerOk

`func (o *GetListenerResponse) GetListenerOk() (*LoadBalancerListener, bool)`

GetListenerOk returns a tuple with the Listener field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListener

`func (o *GetListenerResponse) SetListener(v LoadBalancerListener)`

SetListener sets Listener field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


