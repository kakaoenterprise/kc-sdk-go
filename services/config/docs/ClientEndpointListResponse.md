# ClientEndpointListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Data** | Pointer to [**[]ClientEndpoint**](ClientEndpoint.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewClientEndpointListResponse

`func NewClientEndpointListResponse(success bool, ) *ClientEndpointListResponse`

NewClientEndpointListResponse instantiates a new ClientEndpointListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientEndpointListResponseWithDefaults

`func NewClientEndpointListResponseWithDefaults() *ClientEndpointListResponse`

NewClientEndpointListResponseWithDefaults instantiates a new ClientEndpointListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ClientEndpointListResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ClientEndpointListResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ClientEndpointListResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *ClientEndpointListResponse) GetData() []ClientEndpoint`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ClientEndpointListResponse) GetDataOk() (*[]ClientEndpoint, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ClientEndpointListResponse) SetData(v []ClientEndpoint)`

SetData sets Data field to given value.

### HasData

`func (o *ClientEndpointListResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *ClientEndpointListResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ClientEndpointListResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ClientEndpointListResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ClientEndpointListResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


