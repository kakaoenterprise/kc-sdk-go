# ControlPlaneEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** | 컨트롤 플레인 접속 주소 | 
**Port** | **int32** | 컨트롤 플레인 접속 포트 | 

## Methods

### NewControlPlaneEndpoint

`func NewControlPlaneEndpoint(host string, port int32, ) *ControlPlaneEndpoint`

NewControlPlaneEndpoint instantiates a new ControlPlaneEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControlPlaneEndpointWithDefaults

`func NewControlPlaneEndpointWithDefaults() *ControlPlaneEndpoint`

NewControlPlaneEndpointWithDefaults instantiates a new ControlPlaneEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *ControlPlaneEndpoint) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ControlPlaneEndpoint) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ControlPlaneEndpoint) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *ControlPlaneEndpoint) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ControlPlaneEndpoint) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ControlPlaneEndpoint) SetPort(v int32)`

SetPort sets Port field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


