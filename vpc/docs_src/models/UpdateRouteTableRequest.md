# UpdateRouteTableRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VpcRouteTable** | [**UpdateRouteTable**](UpdateRouteTable.md) | 라우팅 테이블의 메타데이터 | 

## Methods

### NewUpdateRouteTableRequest

`func NewUpdateRouteTableRequest(vpcRouteTable UpdateRouteTable, ) *UpdateRouteTableRequest`

NewUpdateRouteTableRequest instantiates a new UpdateRouteTableRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRouteTableRequestWithDefaults

`func NewUpdateRouteTableRequestWithDefaults() *UpdateRouteTableRequest`

NewUpdateRouteTableRequestWithDefaults instantiates a new UpdateRouteTableRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVpcRouteTable

`func (o *UpdateRouteTableRequest) GetVpcRouteTable() UpdateRouteTable`

GetVpcRouteTable returns the VpcRouteTable field if non-nil, zero value otherwise.

### GetVpcRouteTableOk

`func (o *UpdateRouteTableRequest) GetVpcRouteTableOk() (*UpdateRouteTable, bool)`

GetVpcRouteTableOk returns a tuple with the VpcRouteTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcRouteTable

`func (o *UpdateRouteTableRequest) SetVpcRouteTable(v UpdateRouteTable)`

SetVpcRouteTable sets VpcRouteTable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


