# UpdateTgwRouteTableResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RouteTable** | [**TgwRouteTableResult**](TgwRouteTableResult.md) | 수정된 Transit Gateway 라우팅 테이블 정보 | 

## Methods

### NewUpdateTgwRouteTableResponse

`func NewUpdateTgwRouteTableResponse(routeTable TgwRouteTableResult, ) *UpdateTgwRouteTableResponse`

NewUpdateTgwRouteTableResponse instantiates a new UpdateTgwRouteTableResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTgwRouteTableResponseWithDefaults

`func NewUpdateTgwRouteTableResponseWithDefaults() *UpdateTgwRouteTableResponse`

NewUpdateTgwRouteTableResponseWithDefaults instantiates a new UpdateTgwRouteTableResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRouteTable

`func (o *UpdateTgwRouteTableResponse) GetRouteTable() TgwRouteTableResult`

GetRouteTable returns the RouteTable field if non-nil, zero value otherwise.

### GetRouteTableOk

`func (o *UpdateTgwRouteTableResponse) GetRouteTableOk() (*TgwRouteTableResult, bool)`

GetRouteTableOk returns a tuple with the RouteTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteTable

`func (o *UpdateTgwRouteTableResponse) SetRouteTable(v TgwRouteTableResult)`

SetRouteTable sets RouteTable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


