# CreateRouteModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RouteType** | [**RouteTableRouteType**](RouteTableRouteType.md) | 라우팅 대상 유형 &lt;br/&gt; - &#x60;igw&#x60;: 인터넷 게이트웨이 &lt;br/&gt; - &#x60;instance&#x60;: 인스턴스 &lt;br/&gt; - &#x60;tgw&#x60;: Transit Gateway | 
**TargetId** | **string** | 지정 대상의 리소스 ID | 
**Destination** | **string** | 목적지 네트워크 주소 (CIDR 형식) | 

## Methods

### NewCreateRouteModel

`func NewCreateRouteModel(routeType RouteTableRouteType, targetId string, destination string, ) *CreateRouteModel`

NewCreateRouteModel instantiates a new CreateRouteModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRouteModelWithDefaults

`func NewCreateRouteModelWithDefaults() *CreateRouteModel`

NewCreateRouteModelWithDefaults instantiates a new CreateRouteModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRouteType

`func (o *CreateRouteModel) GetRouteType() RouteTableRouteType`

GetRouteType returns the RouteType field if non-nil, zero value otherwise.

### GetRouteTypeOk

`func (o *CreateRouteModel) GetRouteTypeOk() (*RouteTableRouteType, bool)`

GetRouteTypeOk returns a tuple with the RouteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteType

`func (o *CreateRouteModel) SetRouteType(v RouteTableRouteType)`

SetRouteType sets RouteType field to given value.


### GetTargetId

`func (o *CreateRouteModel) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *CreateRouteModel) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *CreateRouteModel) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.


### GetDestination

`func (o *CreateRouteModel) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *CreateRouteModel) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *CreateRouteModel) SetDestination(v string)`

SetDestination sets Destination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


