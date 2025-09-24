# VpcInfoRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | VPC의 고유 ID | 
**Subnets** | **[]string** | 연결할 서브넷 ID 목록 | 

## Methods

### NewVpcInfoRequestModel

`func NewVpcInfoRequestModel(id string, subnets []string, ) *VpcInfoRequestModel`

NewVpcInfoRequestModel instantiates a new VpcInfoRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpcInfoRequestModelWithDefaults

`func NewVpcInfoRequestModelWithDefaults() *VpcInfoRequestModel`

NewVpcInfoRequestModelWithDefaults instantiates a new VpcInfoRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VpcInfoRequestModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VpcInfoRequestModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VpcInfoRequestModel) SetId(v string)`

SetId sets Id field to given value.


### GetSubnets

`func (o *VpcInfoRequestModel) GetSubnets() []string`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *VpcInfoRequestModel) GetSubnetsOk() (*[]string, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *VpcInfoRequestModel) SetSubnets(v []string)`

SetSubnets sets Subnets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


