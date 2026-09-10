# UpdateRouteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VpcRoute** | [**UpdateRoute**](UpdateRoute.md) | 수정할 라우팅 정보 | 

## Methods

### NewUpdateRouteRequest

`func NewUpdateRouteRequest(vpcRoute UpdateRoute, ) *UpdateRouteRequest`

NewUpdateRouteRequest instantiates a new UpdateRouteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRouteRequestWithDefaults

`func NewUpdateRouteRequestWithDefaults() *UpdateRouteRequest`

NewUpdateRouteRequestWithDefaults instantiates a new UpdateRouteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVpcRoute

`func (o *UpdateRouteRequest) GetVpcRoute() UpdateRoute`

GetVpcRoute returns the VpcRoute field if non-nil, zero value otherwise.

### GetVpcRouteOk

`func (o *UpdateRouteRequest) GetVpcRouteOk() (*UpdateRoute, bool)`

GetVpcRouteOk returns a tuple with the VpcRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcRoute

`func (o *UpdateRouteRequest) SetVpcRoute(v UpdateRoute)`

SetVpcRoute sets VpcRoute field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


