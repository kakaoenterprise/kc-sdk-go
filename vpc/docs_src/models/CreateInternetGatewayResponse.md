# CreateInternetGatewayResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Igw** | [**IGWResult**](IGWResult.md) | 생성된 인터넷 게이트웨이 정보 | 

## Methods

### NewCreateInternetGatewayResponse

`func NewCreateInternetGatewayResponse(igw IGWResult, ) *CreateInternetGatewayResponse`

NewCreateInternetGatewayResponse instantiates a new CreateInternetGatewayResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInternetGatewayResponseWithDefaults

`func NewCreateInternetGatewayResponseWithDefaults() *CreateInternetGatewayResponse`

NewCreateInternetGatewayResponseWithDefaults instantiates a new CreateInternetGatewayResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIgw

`func (o *CreateInternetGatewayResponse) GetIgw() IGWResult`

GetIgw returns the Igw field if non-nil, zero value otherwise.

### GetIgwOk

`func (o *CreateInternetGatewayResponse) GetIgwOk() (*IGWResult, bool)`

GetIgwOk returns a tuple with the Igw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgw

`func (o *CreateInternetGatewayResponse) SetIgw(v IGWResult)`

SetIgw sets Igw field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


