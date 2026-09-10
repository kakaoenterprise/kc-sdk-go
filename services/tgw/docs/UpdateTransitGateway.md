# UpdateTransitGateway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Options** | Pointer to [**NullableUpdateTransitGatewayTgwOptionRequest**](UpdateTransitGatewayTgwOptionRequest.md) |  | [optional] 

## Methods

### NewUpdateTransitGateway

`func NewUpdateTransitGateway() *UpdateTransitGateway`

NewUpdateTransitGateway instantiates a new UpdateTransitGateway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTransitGatewayWithDefaults

`func NewUpdateTransitGatewayWithDefaults() *UpdateTransitGateway`

NewUpdateTransitGatewayWithDefaults instantiates a new UpdateTransitGateway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateTransitGateway) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateTransitGateway) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateTransitGateway) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateTransitGateway) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateTransitGateway) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateTransitGateway) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOptions

`func (o *UpdateTransitGateway) GetOptions() UpdateTransitGatewayTgwOptionRequest`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *UpdateTransitGateway) GetOptionsOk() (*UpdateTransitGatewayTgwOptionRequest, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *UpdateTransitGateway) SetOptions(v UpdateTransitGatewayTgwOptionRequest)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *UpdateTransitGateway) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *UpdateTransitGateway) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *UpdateTransitGateway) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


