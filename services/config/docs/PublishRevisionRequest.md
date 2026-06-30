# PublishRevisionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revision** | **int32** |  | 
**ChangeSummary** | Pointer to **string** |  | [optional] 

## Methods

### NewPublishRevisionRequest

`func NewPublishRevisionRequest(revision int32, ) *PublishRevisionRequest`

NewPublishRevisionRequest instantiates a new PublishRevisionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublishRevisionRequestWithDefaults

`func NewPublishRevisionRequestWithDefaults() *PublishRevisionRequest`

NewPublishRevisionRequestWithDefaults instantiates a new PublishRevisionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevision

`func (o *PublishRevisionRequest) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *PublishRevisionRequest) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *PublishRevisionRequest) SetRevision(v int32)`

SetRevision sets Revision field to given value.


### GetChangeSummary

`func (o *PublishRevisionRequest) GetChangeSummary() string`

GetChangeSummary returns the ChangeSummary field if non-nil, zero value otherwise.

### GetChangeSummaryOk

`func (o *PublishRevisionRequest) GetChangeSummaryOk() (*string, bool)`

GetChangeSummaryOk returns a tuple with the ChangeSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeSummary

`func (o *PublishRevisionRequest) SetChangeSummary(v string)`

SetChangeSummary sets ChangeSummary field to given value.

### HasChangeSummary

`func (o *PublishRevisionRequest) HasChangeSummary() bool`

HasChangeSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


