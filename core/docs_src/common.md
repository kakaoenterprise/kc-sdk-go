# common

```go
package common // import "."

CONSTANTS

const (
	Version = "2.0.0"
)

FUNCTIONS

func NewHTTPClient(config ClientRuntimeConfig) *http.Client
    NewHTTPClient는 주입된 client의 정책을 보존하고 SDK 공통 header만 추가함. 인증, retry, timeout,
    proxy와 TLS 설정은 호출자가 구성한 client가 소유함.

    AllowEndpointOverride가 true면 Island 조회 계층을 하나 더 끼워 넣고, false(기본값)면 설치하지 않아
    Island가 절대 안 건드림.

func Ptr[T any](value T) *T
    Ptr은 optional parameter를 구성하는 데 사용함.

func WrapError(err error, response *http.Response) error
    WrapError는 서비스별 generated error를 공통 오류 형태로 정규화함.

TYPES

type ClientRuntimeConfig struct {
	HTTPClient *http.Client
	BaseURL    string
	UserAgent  string
	APIVersion string

	// AllowEndpointOverride가 true면, BaseURL이 region 계산값이라는 뜻이고
	// 요청 시점에 Island(client-endpoints) 조회 결과로 대체될 수 있음.
	// false(기본값)면 BaseURL을 사용자가 명시한 것으로 보고 절대 건드리지 않음
	// — Island 조회 코드 자체가 설치되지 않음.
	AllowEndpointOverride bool

	// ConfigEndpoint는 Island(서비스 주소 조회 서버) 주소. 비어 있으면 Region으로
	// 계산함. AllowEndpointOverride가 false면 안 쓰임.
	ConfigEndpoint string

	// Region은 ConfigEndpoint 기본값 계산에만 씀. AllowEndpointOverride가
	// false면 안 쓰임.
	Region string
}
    ClientRuntimeConfig는 호출자가 조립한 http.Client를 그대로 주입해 서비스 SDK client를 구성하는 저수준
    입력임. 인증·재시도·timeout·proxy 정책은 HTTPClient를 구성한 호출자가 소유함 (CLI 전용 — 일반 SDK 사용자는
    config.Config·LoadDefaultConfig를 씀).

type Error struct {
	Status    int
	Code      string
	Message   string
	RequestID string
	Err       error
}
    Error는 KakaoCloud API 오류와 요청 추적 정보를 보존함.

func (e *Error) Error() string

func (e *Error) Unwrap() error

type NullablePagination struct {
	// Has unexported fields.
}

func NewNullablePagination(val *Pagination) *NullablePagination

func (v NullablePagination) Get() *Pagination

func (v NullablePagination) IsSet() bool

func (v NullablePagination) MarshalJSON() ([]byte, error)

func (v *NullablePagination) Set(val *Pagination)

func (v *NullablePagination) UnmarshalJSON(src []byte) error

func (v *NullablePagination) Unset()

type Pagination struct {
	// 조회 시작 위치
	Offset int64 `json:"offset"`
	// 페이지당 최대 반환 항목 수
	Limit int32 `json:"limit"`
	// 전체 항목 수
	Total int64 `json:"total"`
}
    Pagination struct for Pagination

func NewPagination(offset int64, limit int32, total int64) *Pagination
    NewPagination instantiates a new Pagination object This constructor will
    assign default values to properties that have it defined, and makes sure
    properties required by API are set, but the set of arguments will change
    when the set of required properties is changed

func NewPaginationWithDefaults() *Pagination
    NewPaginationWithDefaults instantiates a new Pagination object This
    constructor will only assign default values to properties that have it
    defined, but it doesn't guarantee that properties required by API are set

func (o *Pagination) GetLimit() int32
    GetLimit returns the Limit field value

func (o *Pagination) GetLimitOk() (*int32, bool)
    GetLimitOk returns a tuple with the Limit field value and a boolean to check
    if the value has been set.

func (o *Pagination) GetOffset() int64
    GetOffset returns the Offset field value

func (o *Pagination) GetOffsetOk() (*int64, bool)
    GetOffsetOk returns a tuple with the Offset field value and a boolean to
    check if the value has been set.

func (o *Pagination) GetTotal() int64
    GetTotal returns the Total field value

func (o *Pagination) GetTotalOk() (*int64, bool)
    GetTotalOk returns a tuple with the Total field value and a boolean to check
    if the value has been set.

func (o Pagination) MarshalJSON() ([]byte, error)

func (o *Pagination) SetLimit(v int32)
    SetLimit sets field value

func (o *Pagination) SetOffset(v int64)
    SetOffset sets field value

func (o *Pagination) SetTotal(v int64)
    SetTotal sets field value

func (o Pagination) ToMap() (map[string]interface{}, error)

func (o *Pagination) UnmarshalJSON(data []byte) (err error)
```
