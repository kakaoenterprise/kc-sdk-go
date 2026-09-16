# InsertHeaderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**XForwardedFor** | Pointer to [**NullableXForwardedFor**](XForwardedFor.md) | &#x60;X-Forwarded-For&#x60; 헤더 삽입 여부 | [optional] 
**XForwardedProto** | Pointer to [**NullableXForwardedProto**](XForwardedProto.md) | &#x60;X-Forwarded-Proto&#x60; 헤더 삽입 여부 | [optional] 
**XForwardedPort** | Pointer to [**NullableXForwardedProto**](XForwardedProto.md) | &#x60;X-Forwarded-Port&#x60; 헤더 삽입 여부 | [optional] 

## Methods

### NewInsertHeaderRequest

`func NewInsertHeaderRequest() *InsertHeaderRequest`

NewInsertHeaderRequest instantiates a new InsertHeaderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsertHeaderRequestWithDefaults

`func NewInsertHeaderRequestWithDefaults() *InsertHeaderRequest`

NewInsertHeaderRequestWithDefaults instantiates a new InsertHeaderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetXForwardedFor

`func (o *InsertHeaderRequest) GetXForwardedFor() XForwardedFor`

GetXForwardedFor returns the XForwardedFor field if non-nil, zero value otherwise.

### GetXForwardedForOk

`func (o *InsertHeaderRequest) GetXForwardedForOk() (*XForwardedFor, bool)`

GetXForwardedForOk returns a tuple with the XForwardedFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXForwardedFor

`func (o *InsertHeaderRequest) SetXForwardedFor(v XForwardedFor)`

SetXForwardedFor sets XForwardedFor field to given value.

### HasXForwardedFor

`func (o *InsertHeaderRequest) HasXForwardedFor() bool`

HasXForwardedFor returns a boolean if a field has been set.

### SetXForwardedForNil

`func (o *InsertHeaderRequest) SetXForwardedForNil(b bool)`

 SetXForwardedForNil sets the value for XForwardedFor to be an explicit nil

### UnsetXForwardedFor
`func (o *InsertHeaderRequest) UnsetXForwardedFor()`

UnsetXForwardedFor ensures that no value is present for XForwardedFor, not even an explicit nil
### GetXForwardedProto

`func (o *InsertHeaderRequest) GetXForwardedProto() XForwardedProto`

GetXForwardedProto returns the XForwardedProto field if non-nil, zero value otherwise.

### GetXForwardedProtoOk

`func (o *InsertHeaderRequest) GetXForwardedProtoOk() (*XForwardedProto, bool)`

GetXForwardedProtoOk returns a tuple with the XForwardedProto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXForwardedProto

`func (o *InsertHeaderRequest) SetXForwardedProto(v XForwardedProto)`

SetXForwardedProto sets XForwardedProto field to given value.

### HasXForwardedProto

`func (o *InsertHeaderRequest) HasXForwardedProto() bool`

HasXForwardedProto returns a boolean if a field has been set.

### SetXForwardedProtoNil

`func (o *InsertHeaderRequest) SetXForwardedProtoNil(b bool)`

 SetXForwardedProtoNil sets the value for XForwardedProto to be an explicit nil

### UnsetXForwardedProto
`func (o *InsertHeaderRequest) UnsetXForwardedProto()`

UnsetXForwardedProto ensures that no value is present for XForwardedProto, not even an explicit nil
### GetXForwardedPort

`func (o *InsertHeaderRequest) GetXForwardedPort() XForwardedProto`

GetXForwardedPort returns the XForwardedPort field if non-nil, zero value otherwise.

### GetXForwardedPortOk

`func (o *InsertHeaderRequest) GetXForwardedPortOk() (*XForwardedProto, bool)`

GetXForwardedPortOk returns a tuple with the XForwardedPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXForwardedPort

`func (o *InsertHeaderRequest) SetXForwardedPort(v XForwardedProto)`

SetXForwardedPort sets XForwardedPort field to given value.

### HasXForwardedPort

`func (o *InsertHeaderRequest) HasXForwardedPort() bool`

HasXForwardedPort returns a boolean if a field has been set.

### SetXForwardedPortNil

`func (o *InsertHeaderRequest) SetXForwardedPortNil(b bool)`

 SetXForwardedPortNil sets the value for XForwardedPort to be an explicit nil

### UnsetXForwardedPort
`func (o *InsertHeaderRequest) UnsetXForwardedPort()`

UnsetXForwardedPort ensures that no value is present for XForwardedPort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


