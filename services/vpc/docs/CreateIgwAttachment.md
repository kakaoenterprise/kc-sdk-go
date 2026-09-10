# CreateIgwAttachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VpcId** | **string** | 인터넷 게이트웨이를 연결할 VPC의 고유 ID &lt;br/&gt;- [List VPCs](https://docs.kakaocloud.com/openapi/networking/vpc/list-vpcs)에서 확인 | 

## Methods

### NewCreateIgwAttachment

`func NewCreateIgwAttachment(vpcId string, ) *CreateIgwAttachment`

NewCreateIgwAttachment instantiates a new CreateIgwAttachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIgwAttachmentWithDefaults

`func NewCreateIgwAttachmentWithDefaults() *CreateIgwAttachment`

NewCreateIgwAttachmentWithDefaults instantiates a new CreateIgwAttachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVpcId

`func (o *CreateIgwAttachment) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *CreateIgwAttachment) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *CreateIgwAttachment) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


