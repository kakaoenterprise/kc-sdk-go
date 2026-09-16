# CreateTransitGatewayRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tgw** | [**CreateTransitGateway**](CreateTransitGateway.md) | 생성할 Transit Gateway 정보 | 

## Methods

### NewCreateTransitGatewayRequest

`func NewCreateTransitGatewayRequest(tgw CreateTransitGateway, ) *CreateTransitGatewayRequest`

NewCreateTransitGatewayRequest instantiates a new CreateTransitGatewayRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTransitGatewayRequestWithDefaults

`func NewCreateTransitGatewayRequestWithDefaults() *CreateTransitGatewayRequest`

NewCreateTransitGatewayRequestWithDefaults instantiates a new CreateTransitGatewayRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTgw

`func (o *CreateTransitGatewayRequest) GetTgw() CreateTransitGateway`

GetTgw returns the Tgw field if non-nil, zero value otherwise.

### GetTgwOk

`func (o *CreateTransitGatewayRequest) GetTgwOk() (*CreateTransitGateway, bool)`

GetTgwOk returns a tuple with the Tgw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgw

`func (o *CreateTransitGatewayRequest) SetTgw(v CreateTransitGateway)`

SetTgw sets Tgw field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


