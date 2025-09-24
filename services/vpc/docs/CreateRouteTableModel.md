# CreateRouteTableModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | 생성할 라우팅 테이블의 이름 | 
**VpcId** | **string** | 라우팅 테이블을 생성할 대상 VPC의 ID | 

## Methods

### NewCreateRouteTableModel

`func NewCreateRouteTableModel(name string, vpcId string, ) *CreateRouteTableModel`

NewCreateRouteTableModel instantiates a new CreateRouteTableModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRouteTableModelWithDefaults

`func NewCreateRouteTableModelWithDefaults() *CreateRouteTableModel`

NewCreateRouteTableModelWithDefaults instantiates a new CreateRouteTableModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateRouteTableModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateRouteTableModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateRouteTableModel) SetName(v string)`

SetName sets Name field to given value.


### GetVpcId

`func (o *CreateRouteTableModel) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *CreateRouteTableModel) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *CreateRouteTableModel) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


