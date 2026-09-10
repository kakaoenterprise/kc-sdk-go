# CreateRouteTableRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VpcRouteTable** | [**CreateRouteTable**](CreateRouteTable.md) | 생성할 라우팅 테이블 정보 | 

## Methods

### NewCreateRouteTableRequest

`func NewCreateRouteTableRequest(vpcRouteTable CreateRouteTable, ) *CreateRouteTableRequest`

NewCreateRouteTableRequest instantiates a new CreateRouteTableRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRouteTableRequestWithDefaults

`func NewCreateRouteTableRequestWithDefaults() *CreateRouteTableRequest`

NewCreateRouteTableRequestWithDefaults instantiates a new CreateRouteTableRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVpcRouteTable

`func (o *CreateRouteTableRequest) GetVpcRouteTable() CreateRouteTable`

GetVpcRouteTable returns the VpcRouteTable field if non-nil, zero value otherwise.

### GetVpcRouteTableOk

`func (o *CreateRouteTableRequest) GetVpcRouteTableOk() (*CreateRouteTable, bool)`

GetVpcRouteTableOk returns a tuple with the VpcRouteTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcRouteTable

`func (o *CreateRouteTableRequest) SetVpcRouteTable(v CreateRouteTable)`

SetVpcRouteTable sets VpcRouteTable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


