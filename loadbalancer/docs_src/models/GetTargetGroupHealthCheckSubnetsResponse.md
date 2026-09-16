# GetTargetGroupHealthCheckSubnetsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HealthCheckSubnets** | [**[]Subnet**](Subnet.md) | 대상 그룹 헬스 체크 서브넷 목록 | 

## Methods

### NewGetTargetGroupHealthCheckSubnetsResponse

`func NewGetTargetGroupHealthCheckSubnetsResponse(healthCheckSubnets []Subnet, ) *GetTargetGroupHealthCheckSubnetsResponse`

NewGetTargetGroupHealthCheckSubnetsResponse instantiates a new GetTargetGroupHealthCheckSubnetsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTargetGroupHealthCheckSubnetsResponseWithDefaults

`func NewGetTargetGroupHealthCheckSubnetsResponseWithDefaults() *GetTargetGroupHealthCheckSubnetsResponse`

NewGetTargetGroupHealthCheckSubnetsResponseWithDefaults instantiates a new GetTargetGroupHealthCheckSubnetsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHealthCheckSubnets

`func (o *GetTargetGroupHealthCheckSubnetsResponse) GetHealthCheckSubnets() []Subnet`

GetHealthCheckSubnets returns the HealthCheckSubnets field if non-nil, zero value otherwise.

### GetHealthCheckSubnetsOk

`func (o *GetTargetGroupHealthCheckSubnetsResponse) GetHealthCheckSubnetsOk() (*[]Subnet, bool)`

GetHealthCheckSubnetsOk returns a tuple with the HealthCheckSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckSubnets

`func (o *GetTargetGroupHealthCheckSubnetsResponse) SetHealthCheckSubnets(v []Subnet)`

SetHealthCheckSubnets sets HealthCheckSubnets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


