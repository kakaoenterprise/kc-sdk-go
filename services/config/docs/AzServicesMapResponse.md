# AzServicesMapResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Data** | Pointer to **map[string][]string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewAzServicesMapResponse

`func NewAzServicesMapResponse(success bool, ) *AzServicesMapResponse`

NewAzServicesMapResponse instantiates a new AzServicesMapResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzServicesMapResponseWithDefaults

`func NewAzServicesMapResponseWithDefaults() *AzServicesMapResponse`

NewAzServicesMapResponseWithDefaults instantiates a new AzServicesMapResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *AzServicesMapResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AzServicesMapResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AzServicesMapResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *AzServicesMapResponse) GetData() map[string][]string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AzServicesMapResponse) GetDataOk() (*map[string][]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AzServicesMapResponse) SetData(v map[string][]string)`

SetData sets Data field to given value.

### HasData

`func (o *AzServicesMapResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *AzServicesMapResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AzServicesMapResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AzServicesMapResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AzServicesMapResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


