# GetTgwAttachmentsResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attachments** | [**[]BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel**](BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel.md) | 조회된 TGW Attachment 정보 | 
**Pagination** | [**PaginationModel**](PaginationModel.md) | 페이지네이션 메타데이터 | 

## Methods

### NewGetTgwAttachmentsResponseModel

`func NewGetTgwAttachmentsResponseModel(attachments []BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel, pagination PaginationModel, ) *GetTgwAttachmentsResponseModel`

NewGetTgwAttachmentsResponseModel instantiates a new GetTgwAttachmentsResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTgwAttachmentsResponseModelWithDefaults

`func NewGetTgwAttachmentsResponseModelWithDefaults() *GetTgwAttachmentsResponseModel`

NewGetTgwAttachmentsResponseModelWithDefaults instantiates a new GetTgwAttachmentsResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachments

`func (o *GetTgwAttachmentsResponseModel) GetAttachments() []BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *GetTgwAttachmentsResponseModel) GetAttachmentsOk() (*[]BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *GetTgwAttachmentsResponseModel) SetAttachments(v []BnsTgwV1ApiListTgwAttachmentsModelTgwAttachmentResponseModel)`

SetAttachments sets Attachments field to given value.


### GetPagination

`func (o *GetTgwAttachmentsResponseModel) GetPagination() PaginationModel`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetTgwAttachmentsResponseModel) GetPaginationOk() (*PaginationModel, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetTgwAttachmentsResponseModel) SetPagination(v PaginationModel)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


