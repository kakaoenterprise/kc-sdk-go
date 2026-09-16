# ListTgwAttachmentsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attachments** | [**[]TgwAttachment**](TgwAttachment.md) | 조회된 TGW Attachment 정보 | 
**Pagination** | [**Pagination**](Pagination.md) | 페이지네이션 메타데이터 | 

## Methods

### NewListTgwAttachmentsResponse

`func NewListTgwAttachmentsResponse(attachments []TgwAttachment, pagination Pagination, ) *ListTgwAttachmentsResponse`

NewListTgwAttachmentsResponse instantiates a new ListTgwAttachmentsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTgwAttachmentsResponseWithDefaults

`func NewListTgwAttachmentsResponseWithDefaults() *ListTgwAttachmentsResponse`

NewListTgwAttachmentsResponseWithDefaults instantiates a new ListTgwAttachmentsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachments

`func (o *ListTgwAttachmentsResponse) GetAttachments() []TgwAttachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *ListTgwAttachmentsResponse) GetAttachmentsOk() (*[]TgwAttachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *ListTgwAttachmentsResponse) SetAttachments(v []TgwAttachment)`

SetAttachments sets Attachments field to given value.


### GetPagination

`func (o *ListTgwAttachmentsResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListTgwAttachmentsResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListTgwAttachmentsResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


