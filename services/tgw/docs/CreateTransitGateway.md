# CreateTransitGateway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Transit Gateway 이름 | 
**Options** | [**CreateTransitGatewayTgwOptionRequest**](CreateTransitGatewayTgwOptionRequest.md) | Transit Gateway 옵션 설정 | 

## Methods

### NewCreateTransitGateway

`func NewCreateTransitGateway(name string, options CreateTransitGatewayTgwOptionRequest, ) *CreateTransitGateway`

NewCreateTransitGateway instantiates a new CreateTransitGateway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTransitGatewayWithDefaults

`func NewCreateTransitGatewayWithDefaults() *CreateTransitGateway`

NewCreateTransitGatewayWithDefaults instantiates a new CreateTransitGateway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateTransitGateway) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTransitGateway) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTransitGateway) SetName(v string)`

SetName sets Name field to given value.


### GetOptions

`func (o *CreateTransitGateway) GetOptions() CreateTransitGatewayTgwOptionRequest`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *CreateTransitGateway) GetOptionsOk() (*CreateTransitGatewayTgwOptionRequest, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *CreateTransitGateway) SetOptions(v CreateTransitGatewayTgwOptionRequest)`

SetOptions sets Options field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


