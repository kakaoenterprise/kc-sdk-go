# OpenapiEnvConfigResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Data** | Pointer to [**OpenapiEnvConfig**](OpenapiEnvConfig.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewOpenapiEnvConfigResponse

`func NewOpenapiEnvConfigResponse(success bool, ) *OpenapiEnvConfigResponse`

NewOpenapiEnvConfigResponse instantiates a new OpenapiEnvConfigResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenapiEnvConfigResponseWithDefaults

`func NewOpenapiEnvConfigResponseWithDefaults() *OpenapiEnvConfigResponse`

NewOpenapiEnvConfigResponseWithDefaults instantiates a new OpenapiEnvConfigResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *OpenapiEnvConfigResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *OpenapiEnvConfigResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *OpenapiEnvConfigResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *OpenapiEnvConfigResponse) GetData() OpenapiEnvConfig`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OpenapiEnvConfigResponse) GetDataOk() (*OpenapiEnvConfig, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OpenapiEnvConfigResponse) SetData(v OpenapiEnvConfig)`

SetData sets Data field to given value.

### HasData

`func (o *OpenapiEnvConfigResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *OpenapiEnvConfigResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *OpenapiEnvConfigResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *OpenapiEnvConfigResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *OpenapiEnvConfigResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


