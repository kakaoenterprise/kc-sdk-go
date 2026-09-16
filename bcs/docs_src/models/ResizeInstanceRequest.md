# ResizeInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Flavor** | [**ResizeInstance**](ResizeInstance.md) | 변경할 인스턴스 유형 정보 | 

## Methods

### NewResizeInstanceRequest

`func NewResizeInstanceRequest(flavor ResizeInstance, ) *ResizeInstanceRequest`

NewResizeInstanceRequest instantiates a new ResizeInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResizeInstanceRequestWithDefaults

`func NewResizeInstanceRequestWithDefaults() *ResizeInstanceRequest`

NewResizeInstanceRequestWithDefaults instantiates a new ResizeInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlavor

`func (o *ResizeInstanceRequest) GetFlavor() ResizeInstance`

GetFlavor returns the Flavor field if non-nil, zero value otherwise.

### GetFlavorOk

`func (o *ResizeInstanceRequest) GetFlavorOk() (*ResizeInstance, bool)`

GetFlavorOk returns a tuple with the Flavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavor

`func (o *ResizeInstanceRequest) SetFlavor(v ResizeInstance)`

SetFlavor sets Flavor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


