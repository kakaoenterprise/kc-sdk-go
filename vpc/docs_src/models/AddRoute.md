# AddRoute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RouteType** | [**RouteType**](RouteType.md) | 라우팅 대상 유형 | 
**TargetId** | **string** | &#x60;route_type&#x60;에 해당하는 라우팅 대상 리소스 ID | 
**Destination** | **string** | 목적지 네트워크 주소 (CIDR 형식) | 

## Methods

### NewAddRoute

`func NewAddRoute(routeType RouteType, targetId string, destination string, ) *AddRoute`

NewAddRoute instantiates a new AddRoute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddRouteWithDefaults

`func NewAddRouteWithDefaults() *AddRoute`

NewAddRouteWithDefaults instantiates a new AddRoute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRouteType

`func (o *AddRoute) GetRouteType() RouteType`

GetRouteType returns the RouteType field if non-nil, zero value otherwise.

### GetRouteTypeOk

`func (o *AddRoute) GetRouteTypeOk() (*RouteType, bool)`

GetRouteTypeOk returns a tuple with the RouteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteType

`func (o *AddRoute) SetRouteType(v RouteType)`

SetRouteType sets RouteType field to given value.


### GetTargetId

`func (o *AddRoute) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *AddRoute) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *AddRoute) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.


### GetDestination

`func (o *AddRoute) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *AddRoute) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *AddRoute) SetDestination(v string)`

SetDestination sets Destination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


