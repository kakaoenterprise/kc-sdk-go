# CreateVpcRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vpc** | [**CreateVpc**](CreateVpc.md) | 생성할 VPC 정보 | 

## Methods

### NewCreateVpcRequest

`func NewCreateVpcRequest(vpc CreateVpc, ) *CreateVpcRequest`

NewCreateVpcRequest instantiates a new CreateVpcRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVpcRequestWithDefaults

`func NewCreateVpcRequestWithDefaults() *CreateVpcRequest`

NewCreateVpcRequestWithDefaults instantiates a new CreateVpcRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVpc

`func (o *CreateVpcRequest) GetVpc() CreateVpc`

GetVpc returns the Vpc field if non-nil, zero value otherwise.

### GetVpcOk

`func (o *CreateVpcRequest) GetVpcOk() (*CreateVpc, bool)`

GetVpcOk returns a tuple with the Vpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpc

`func (o *CreateVpcRequest) SetVpc(v CreateVpc)`

SetVpc sets Vpc field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


