# CreateSubnetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subnet** | [**SubnetResult**](SubnetResult.md) | 생성된 서브넷 정보 | 

## Methods

### NewCreateSubnetResponse

`func NewCreateSubnetResponse(subnet SubnetResult, ) *CreateSubnetResponse`

NewCreateSubnetResponse instantiates a new CreateSubnetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSubnetResponseWithDefaults

`func NewCreateSubnetResponseWithDefaults() *CreateSubnetResponse`

NewCreateSubnetResponseWithDefaults instantiates a new CreateSubnetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubnet

`func (o *CreateSubnetResponse) GetSubnet() SubnetResult`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *CreateSubnetResponse) GetSubnetOk() (*SubnetResult, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *CreateSubnetResponse) SetSubnet(v SubnetResult)`

SetSubnet sets Subnet field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


