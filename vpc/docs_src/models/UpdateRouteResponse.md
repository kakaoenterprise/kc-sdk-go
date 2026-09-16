# UpdateRouteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VpcRoute** | [**Route**](Route.md) | 수정된 라우팅 정보 | 

## Methods

### NewUpdateRouteResponse

`func NewUpdateRouteResponse(vpcRoute Route, ) *UpdateRouteResponse`

NewUpdateRouteResponse instantiates a new UpdateRouteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRouteResponseWithDefaults

`func NewUpdateRouteResponseWithDefaults() *UpdateRouteResponse`

NewUpdateRouteResponseWithDefaults instantiates a new UpdateRouteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVpcRoute

`func (o *UpdateRouteResponse) GetVpcRoute() Route`

GetVpcRoute returns the VpcRoute field if non-nil, zero value otherwise.

### GetVpcRouteOk

`func (o *UpdateRouteResponse) GetVpcRouteOk() (*Route, bool)`

GetVpcRouteOk returns a tuple with the VpcRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcRoute

`func (o *UpdateRouteResponse) SetVpcRoute(v Route)`

SetVpcRoute sets VpcRoute field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


