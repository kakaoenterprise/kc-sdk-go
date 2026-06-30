# AzServiceListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Data** | Pointer to **[]string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewAzServiceListResponse

`func NewAzServiceListResponse(success bool, ) *AzServiceListResponse`

NewAzServiceListResponse instantiates a new AzServiceListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzServiceListResponseWithDefaults

`func NewAzServiceListResponseWithDefaults() *AzServiceListResponse`

NewAzServiceListResponseWithDefaults instantiates a new AzServiceListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *AzServiceListResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AzServiceListResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AzServiceListResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *AzServiceListResponse) GetData() []string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AzServiceListResponse) GetDataOk() (*[]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AzServiceListResponse) SetData(v []string)`

SetData sets Data field to given value.

### HasData

`func (o *AzServiceListResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *AzServiceListResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AzServiceListResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AzServiceListResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AzServiceListResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


