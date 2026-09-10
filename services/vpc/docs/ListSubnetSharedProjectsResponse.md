# ListSubnetSharedProjectsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Projects** | [**[]Project**](Project.md) | 서브넷을 공유받은 프로젝트 목록 | 

## Methods

### NewListSubnetSharedProjectsResponse

`func NewListSubnetSharedProjectsResponse(projects []Project, ) *ListSubnetSharedProjectsResponse`

NewListSubnetSharedProjectsResponse instantiates a new ListSubnetSharedProjectsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSubnetSharedProjectsResponseWithDefaults

`func NewListSubnetSharedProjectsResponseWithDefaults() *ListSubnetSharedProjectsResponse`

NewListSubnetSharedProjectsResponseWithDefaults instantiates a new ListSubnetSharedProjectsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjects

`func (o *ListSubnetSharedProjectsResponse) GetProjects() []Project`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *ListSubnetSharedProjectsResponse) GetProjectsOk() (*[]Project, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *ListSubnetSharedProjectsResponse) SetProjects(v []Project)`

SetProjects sets Projects field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


