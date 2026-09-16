# AddRouteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VpcRoute** | [**AddRoute**](AddRoute.md) | 추가할 라우팅 정보 | 

## Methods

### NewAddRouteRequest

`func NewAddRouteRequest(vpcRoute AddRoute, ) *AddRouteRequest`

NewAddRouteRequest instantiates a new AddRouteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddRouteRequestWithDefaults

`func NewAddRouteRequestWithDefaults() *AddRouteRequest`

NewAddRouteRequestWithDefaults instantiates a new AddRouteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVpcRoute

`func (o *AddRouteRequest) GetVpcRoute() AddRoute`

GetVpcRoute returns the VpcRoute field if non-nil, zero value otherwise.

### GetVpcRouteOk

`func (o *AddRouteRequest) GetVpcRouteOk() (*AddRoute, bool)`

GetVpcRouteOk returns a tuple with the VpcRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcRoute

`func (o *AddRouteRequest) SetVpcRoute(v AddRoute)`

SetVpcRoute sets VpcRoute field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


