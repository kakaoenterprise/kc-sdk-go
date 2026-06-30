# AzPolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Data** | Pointer to [**AzPolicy**](AzPolicy.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewAzPolicyResponse

`func NewAzPolicyResponse(success bool, ) *AzPolicyResponse`

NewAzPolicyResponse instantiates a new AzPolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzPolicyResponseWithDefaults

`func NewAzPolicyResponseWithDefaults() *AzPolicyResponse`

NewAzPolicyResponseWithDefaults instantiates a new AzPolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *AzPolicyResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AzPolicyResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AzPolicyResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *AzPolicyResponse) GetData() AzPolicy`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AzPolicyResponse) GetDataOk() (*AzPolicy, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AzPolicyResponse) SetData(v AzPolicy)`

SetData sets Data field to given value.

### HasData

`func (o *AzPolicyResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *AzPolicyResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AzPolicyResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AzPolicyResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AzPolicyResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


