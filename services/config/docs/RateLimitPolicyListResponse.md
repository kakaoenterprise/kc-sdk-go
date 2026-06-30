# RateLimitPolicyListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Data** | Pointer to [**[]RateLimitPolicy**](RateLimitPolicy.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewRateLimitPolicyListResponse

`func NewRateLimitPolicyListResponse(success bool, ) *RateLimitPolicyListResponse`

NewRateLimitPolicyListResponse instantiates a new RateLimitPolicyListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateLimitPolicyListResponseWithDefaults

`func NewRateLimitPolicyListResponseWithDefaults() *RateLimitPolicyListResponse`

NewRateLimitPolicyListResponseWithDefaults instantiates a new RateLimitPolicyListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *RateLimitPolicyListResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *RateLimitPolicyListResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *RateLimitPolicyListResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *RateLimitPolicyListResponse) GetData() []RateLimitPolicy`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RateLimitPolicyListResponse) GetDataOk() (*[]RateLimitPolicy, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RateLimitPolicyListResponse) SetData(v []RateLimitPolicy)`

SetData sets Data field to given value.

### HasData

`func (o *RateLimitPolicyListResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *RateLimitPolicyListResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RateLimitPolicyListResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RateLimitPolicyListResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *RateLimitPolicyListResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


