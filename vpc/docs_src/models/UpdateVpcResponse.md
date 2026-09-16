# UpdateVpcResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vpc** | [**VpcResult**](VpcResult.md) | VPC 이름 | 

## Methods

### NewUpdateVpcResponse

`func NewUpdateVpcResponse(vpc VpcResult, ) *UpdateVpcResponse`

NewUpdateVpcResponse instantiates a new UpdateVpcResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateVpcResponseWithDefaults

`func NewUpdateVpcResponseWithDefaults() *UpdateVpcResponse`

NewUpdateVpcResponseWithDefaults instantiates a new UpdateVpcResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVpc

`func (o *UpdateVpcResponse) GetVpc() VpcResult`

GetVpc returns the Vpc field if non-nil, zero value otherwise.

### GetVpcOk

`func (o *UpdateVpcResponse) GetVpcOk() (*VpcResult, bool)`

GetVpcOk returns a tuple with the Vpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpc

`func (o *UpdateVpcResponse) SetVpc(v VpcResult)`

SetVpc sets Vpc field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


