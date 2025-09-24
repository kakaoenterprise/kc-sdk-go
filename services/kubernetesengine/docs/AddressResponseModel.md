# AddressResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | 노드의 IP 주소 | 
**Type** | **string** | 노드 주소 유형 | 

## Methods

### NewAddressResponseModel

`func NewAddressResponseModel(address string, type_ string, ) *AddressResponseModel`

NewAddressResponseModel instantiates a new AddressResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressResponseModelWithDefaults

`func NewAddressResponseModelWithDefaults() *AddressResponseModel`

NewAddressResponseModelWithDefaults instantiates a new AddressResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *AddressResponseModel) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AddressResponseModel) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AddressResponseModel) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetType

`func (o *AddressResponseModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddressResponseModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddressResponseModel) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


