# CreateTgwRouteTable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TgwId** | **string** | 라우팅 테이블을 생성할 Transit Gateway ID | 
**Name** | **string** | 라우팅 테이블 이름 | 

## Methods

### NewCreateTgwRouteTable

`func NewCreateTgwRouteTable(tgwId string, name string, ) *CreateTgwRouteTable`

NewCreateTgwRouteTable instantiates a new CreateTgwRouteTable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTgwRouteTableWithDefaults

`func NewCreateTgwRouteTableWithDefaults() *CreateTgwRouteTable`

NewCreateTgwRouteTableWithDefaults instantiates a new CreateTgwRouteTable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTgwId

`func (o *CreateTgwRouteTable) GetTgwId() string`

GetTgwId returns the TgwId field if non-nil, zero value otherwise.

### GetTgwIdOk

`func (o *CreateTgwRouteTable) GetTgwIdOk() (*string, bool)`

GetTgwIdOk returns a tuple with the TgwId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgwId

`func (o *CreateTgwRouteTable) SetTgwId(v string)`

SetTgwId sets TgwId field to given value.


### GetName

`func (o *CreateTgwRouteTable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTgwRouteTable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTgwRouteTable) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


