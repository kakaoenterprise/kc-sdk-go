# CreateTgwAttachmentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attachment** | [**TgwAttachmentResult**](TgwAttachmentResult.md) | Transit Gateway Attachment 정보 | 

## Methods

### NewCreateTgwAttachmentResponse

`func NewCreateTgwAttachmentResponse(attachment TgwAttachmentResult, ) *CreateTgwAttachmentResponse`

NewCreateTgwAttachmentResponse instantiates a new CreateTgwAttachmentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTgwAttachmentResponseWithDefaults

`func NewCreateTgwAttachmentResponseWithDefaults() *CreateTgwAttachmentResponse`

NewCreateTgwAttachmentResponseWithDefaults instantiates a new CreateTgwAttachmentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachment

`func (o *CreateTgwAttachmentResponse) GetAttachment() TgwAttachmentResult`

GetAttachment returns the Attachment field if non-nil, zero value otherwise.

### GetAttachmentOk

`func (o *CreateTgwAttachmentResponse) GetAttachmentOk() (*TgwAttachmentResult, bool)`

GetAttachmentOk returns a tuple with the Attachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachment

`func (o *CreateTgwAttachmentResponse) SetAttachment(v TgwAttachmentResult)`

SetAttachment sets Attachment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


