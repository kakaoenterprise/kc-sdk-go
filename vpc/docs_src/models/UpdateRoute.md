# UpdateRoute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RouteType** | [**RouteType**](RouteType.md) | 라우팅 대상 유형 | 
**TargetId** | **string** | &#x60;route_type&#x60;에 해당하는 라우팅 대상 리소스 ID | 
**Destination** | **string** | 목적지 네트워크 주소 (CIDR 형식) | 

## Methods

### NewUpdateRoute

`func NewUpdateRoute(routeType RouteType, targetId string, destination string, ) *UpdateRoute`

NewUpdateRoute instantiates a new UpdateRoute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRouteWithDefaults

`func NewUpdateRouteWithDefaults() *UpdateRoute`

NewUpdateRouteWithDefaults instantiates a new UpdateRoute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRouteType

`func (o *UpdateRoute) GetRouteType() RouteType`

GetRouteType returns the RouteType field if non-nil, zero value otherwise.

### GetRouteTypeOk

`func (o *UpdateRoute) GetRouteTypeOk() (*RouteType, bool)`

GetRouteTypeOk returns a tuple with the RouteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteType

`func (o *UpdateRoute) SetRouteType(v RouteType)`

SetRouteType sets RouteType field to given value.


### GetTargetId

`func (o *UpdateRoute) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *UpdateRoute) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *UpdateRoute) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.


### GetDestination

`func (o *UpdateRoute) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *UpdateRoute) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *UpdateRoute) SetDestination(v string)`

SetDestination sets Destination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


