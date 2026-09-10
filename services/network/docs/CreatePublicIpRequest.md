# CreatePublicIpRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicIp** | [**CreatePublicIp**](CreatePublicIp.md) | 생성할 퍼블릭 IP 정보 | 

## Methods

### NewCreatePublicIpRequest

`func NewCreatePublicIpRequest(publicIp CreatePublicIp, ) *CreatePublicIpRequest`

NewCreatePublicIpRequest instantiates a new CreatePublicIpRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePublicIpRequestWithDefaults

`func NewCreatePublicIpRequestWithDefaults() *CreatePublicIpRequest`

NewCreatePublicIpRequestWithDefaults instantiates a new CreatePublicIpRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicIp

`func (o *CreatePublicIpRequest) GetPublicIp() CreatePublicIp`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *CreatePublicIpRequest) GetPublicIpOk() (*CreatePublicIp, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *CreatePublicIpRequest) SetPublicIp(v CreatePublicIp)`

SetPublicIp sets PublicIp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


