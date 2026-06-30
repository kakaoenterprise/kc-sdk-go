# RateLimitRevision

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revision** | **int32** |  | 
**PublishedBy** | **string** |  | 
**PublishedAt** | Pointer to **time.Time** |  | [optional] 
**ChangeSummary** | Pointer to **string** |  | [optional] 

## Methods

### NewRateLimitRevision

`func NewRateLimitRevision(revision int32, publishedBy string, ) *RateLimitRevision`

NewRateLimitRevision instantiates a new RateLimitRevision object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateLimitRevisionWithDefaults

`func NewRateLimitRevisionWithDefaults() *RateLimitRevision`

NewRateLimitRevisionWithDefaults instantiates a new RateLimitRevision object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevision

`func (o *RateLimitRevision) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *RateLimitRevision) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *RateLimitRevision) SetRevision(v int32)`

SetRevision sets Revision field to given value.


### GetPublishedBy

`func (o *RateLimitRevision) GetPublishedBy() string`

GetPublishedBy returns the PublishedBy field if non-nil, zero value otherwise.

### GetPublishedByOk

`func (o *RateLimitRevision) GetPublishedByOk() (*string, bool)`

GetPublishedByOk returns a tuple with the PublishedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedBy

`func (o *RateLimitRevision) SetPublishedBy(v string)`

SetPublishedBy sets PublishedBy field to given value.


### GetPublishedAt

`func (o *RateLimitRevision) GetPublishedAt() time.Time`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *RateLimitRevision) GetPublishedAtOk() (*time.Time, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *RateLimitRevision) SetPublishedAt(v time.Time)`

SetPublishedAt sets PublishedAt field to given value.

### HasPublishedAt

`func (o *RateLimitRevision) HasPublishedAt() bool`

HasPublishedAt returns a boolean if a field has been set.

### GetChangeSummary

`func (o *RateLimitRevision) GetChangeSummary() string`

GetChangeSummary returns the ChangeSummary field if non-nil, zero value otherwise.

### GetChangeSummaryOk

`func (o *RateLimitRevision) GetChangeSummaryOk() (*string, bool)`

GetChangeSummaryOk returns a tuple with the ChangeSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeSummary

`func (o *RateLimitRevision) SetChangeSummary(v string)`

SetChangeSummary sets ChangeSummary field to given value.

### HasChangeSummary

`func (o *RateLimitRevision) HasChangeSummary() bool`

HasChangeSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


