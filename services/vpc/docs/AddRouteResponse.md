# AddRouteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VpcRoute** | [**Route**](Route.md) | 추가된 라우팅 정보 | 

## Methods

### NewAddRouteResponse

`func NewAddRouteResponse(vpcRoute Route, ) *AddRouteResponse`

NewAddRouteResponse instantiates a new AddRouteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddRouteResponseWithDefaults

`func NewAddRouteResponseWithDefaults() *AddRouteResponse`

NewAddRouteResponseWithDefaults instantiates a new AddRouteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVpcRoute

`func (o *AddRouteResponse) GetVpcRoute() Route`

GetVpcRoute returns the VpcRoute field if non-nil, zero value otherwise.

### GetVpcRouteOk

`func (o *AddRouteResponse) GetVpcRouteOk() (*Route, bool)`

GetVpcRouteOk returns a tuple with the VpcRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcRoute

`func (o *AddRouteResponse) SetVpcRoute(v Route)`

SetVpcRoute sets VpcRoute field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


