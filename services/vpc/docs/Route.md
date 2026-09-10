# Route

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 라우팅의 고유 ID | 
**ProvisioningStatus** | [**ProvisioningStatus**](ProvisioningStatus.md) | 프로비저닝 상태 | 
**Destination** | **string** | 목적지 네트워크 주소(CIDR 형식) | 
**IsLocalRoute** | **bool** | 로컬 경로 여부&lt;br/&gt;  - &#x60;true&#x60;인 경우 동일 VPC 내부 통신용 경로 | 

## Methods

### NewRoute

`func NewRoute(id string, provisioningStatus ProvisioningStatus, destination string, isLocalRoute bool, ) *Route`

NewRoute instantiates a new Route object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteWithDefaults

`func NewRouteWithDefaults() *Route`

NewRouteWithDefaults instantiates a new Route object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Route) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Route) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Route) SetId(v string)`

SetId sets Id field to given value.


### GetProvisioningStatus

`func (o *Route) GetProvisioningStatus() ProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *Route) GetProvisioningStatusOk() (*ProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *Route) SetProvisioningStatus(v ProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.


### GetDestination

`func (o *Route) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *Route) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *Route) SetDestination(v string)`

SetDestination sets Destination field to given value.


### GetIsLocalRoute

`func (o *Route) GetIsLocalRoute() bool`

GetIsLocalRoute returns the IsLocalRoute field if non-nil, zero value otherwise.

### GetIsLocalRouteOk

`func (o *Route) GetIsLocalRouteOk() (*bool, bool)`

GetIsLocalRouteOk returns a tuple with the IsLocalRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocalRoute

`func (o *Route) SetIsLocalRoute(v bool)`

SetIsLocalRoute sets IsLocalRoute field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


