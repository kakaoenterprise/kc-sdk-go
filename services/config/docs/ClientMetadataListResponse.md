# ClientMetadataListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Data** | Pointer to [**[]ClientMetadata**](ClientMetadata.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewClientMetadataListResponse

`func NewClientMetadataListResponse(success bool, ) *ClientMetadataListResponse`

NewClientMetadataListResponse instantiates a new ClientMetadataListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientMetadataListResponseWithDefaults

`func NewClientMetadataListResponseWithDefaults() *ClientMetadataListResponse`

NewClientMetadataListResponseWithDefaults instantiates a new ClientMetadataListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ClientMetadataListResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ClientMetadataListResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ClientMetadataListResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *ClientMetadataListResponse) GetData() []ClientMetadata`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ClientMetadataListResponse) GetDataOk() (*[]ClientMetadata, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ClientMetadataListResponse) SetData(v []ClientMetadata)`

SetData sets Data field to given value.

### HasData

`func (o *ClientMetadataListResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *ClientMetadataListResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ClientMetadataListResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ClientMetadataListResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ClientMetadataListResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


