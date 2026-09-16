# GetClusterNodeDetailsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | [**NodeDetail**](NodeDetail.md) | 상세 Kubernetes 노드 정보 | 

## Methods

### NewGetClusterNodeDetailsResponse

`func NewGetClusterNodeDetailsResponse(details NodeDetail, ) *GetClusterNodeDetailsResponse`

NewGetClusterNodeDetailsResponse instantiates a new GetClusterNodeDetailsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetClusterNodeDetailsResponseWithDefaults

`func NewGetClusterNodeDetailsResponseWithDefaults() *GetClusterNodeDetailsResponse`

NewGetClusterNodeDetailsResponseWithDefaults instantiates a new GetClusterNodeDetailsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *GetClusterNodeDetailsResponse) GetDetails() NodeDetail`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *GetClusterNodeDetailsResponse) GetDetailsOk() (*NodeDetail, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *GetClusterNodeDetailsResponse) SetDetails(v NodeDetail)`

SetDetails sets Details field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


