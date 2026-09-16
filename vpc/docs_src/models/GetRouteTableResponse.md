# GetRouteTableResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VpcRouteTable** | [**VpcRouteTable**](VpcRouteTable.md) | 라우팅 테이블 상세 정보 | 

## Methods

### NewGetRouteTableResponse

`func NewGetRouteTableResponse(vpcRouteTable VpcRouteTable, ) *GetRouteTableResponse`

NewGetRouteTableResponse instantiates a new GetRouteTableResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRouteTableResponseWithDefaults

`func NewGetRouteTableResponseWithDefaults() *GetRouteTableResponse`

NewGetRouteTableResponseWithDefaults instantiates a new GetRouteTableResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVpcRouteTable

`func (o *GetRouteTableResponse) GetVpcRouteTable() VpcRouteTable`

GetVpcRouteTable returns the VpcRouteTable field if non-nil, zero value otherwise.

### GetVpcRouteTableOk

`func (o *GetRouteTableResponse) GetVpcRouteTableOk() (*VpcRouteTable, bool)`

GetVpcRouteTableOk returns a tuple with the VpcRouteTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcRouteTable

`func (o *GetRouteTableResponse) SetVpcRouteTable(v VpcRouteTable)`

SetVpcRouteTable sets VpcRouteTable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


