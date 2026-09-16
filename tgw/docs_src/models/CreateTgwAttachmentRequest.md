# CreateTgwAttachmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attachment** | [**CreateTgwAttachment**](CreateTgwAttachment.md) | 생성할 Transit Gateway Attachment 정보 | 

## Methods

### NewCreateTgwAttachmentRequest

`func NewCreateTgwAttachmentRequest(attachment CreateTgwAttachment, ) *CreateTgwAttachmentRequest`

NewCreateTgwAttachmentRequest instantiates a new CreateTgwAttachmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTgwAttachmentRequestWithDefaults

`func NewCreateTgwAttachmentRequestWithDefaults() *CreateTgwAttachmentRequest`

NewCreateTgwAttachmentRequestWithDefaults instantiates a new CreateTgwAttachmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachment

`func (o *CreateTgwAttachmentRequest) GetAttachment() CreateTgwAttachment`

GetAttachment returns the Attachment field if non-nil, zero value otherwise.

### GetAttachmentOk

`func (o *CreateTgwAttachmentRequest) GetAttachmentOk() (*CreateTgwAttachment, bool)`

GetAttachmentOk returns a tuple with the Attachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachment

`func (o *CreateTgwAttachmentRequest) SetAttachment(v CreateTgwAttachment)`

SetAttachment sets Attachment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


