# CreateTgwRoute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationCidrBlock** | **string** | 대상 네트워크 CIDR 블록 | 
**TgwAttachmentId** | **string** | Transit Gateway와 연결된 Attachment의 연결 ID | 

## Methods

### NewCreateTgwRoute

`func NewCreateTgwRoute(destinationCidrBlock string, tgwAttachmentId string, ) *CreateTgwRoute`

NewCreateTgwRoute instantiates a new CreateTgwRoute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTgwRouteWithDefaults

`func NewCreateTgwRouteWithDefaults() *CreateTgwRoute`

NewCreateTgwRouteWithDefaults instantiates a new CreateTgwRoute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationCidrBlock

`func (o *CreateTgwRoute) GetDestinationCidrBlock() string`

GetDestinationCidrBlock returns the DestinationCidrBlock field if non-nil, zero value otherwise.

### GetDestinationCidrBlockOk

`func (o *CreateTgwRoute) GetDestinationCidrBlockOk() (*string, bool)`

GetDestinationCidrBlockOk returns a tuple with the DestinationCidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationCidrBlock

`func (o *CreateTgwRoute) SetDestinationCidrBlock(v string)`

SetDestinationCidrBlock sets DestinationCidrBlock field to given value.


### GetTgwAttachmentId

`func (o *CreateTgwRoute) GetTgwAttachmentId() string`

GetTgwAttachmentId returns the TgwAttachmentId field if non-nil, zero value otherwise.

### GetTgwAttachmentIdOk

`func (o *CreateTgwRoute) GetTgwAttachmentIdOk() (*string, bool)`

GetTgwAttachmentIdOk returns a tuple with the TgwAttachmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgwAttachmentId

`func (o *CreateTgwRoute) SetTgwAttachmentId(v string)`

SetTgwAttachmentId sets TgwAttachmentId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


