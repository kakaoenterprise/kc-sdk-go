# config

```go
package config // import "."

TYPES

type Config struct {
	HTTPClient     *http.Client
	Region         string
	ConfigEndpoint string
	UserAgent      string
}
    Config는 일반 SDK 사용자를 위해 조립된 service client 설정임.

func LoadDefaultConfig(ctx context.Context, optFns ...LoadOptionsFunc) (Config, error)
    LoadDefaultConfig는 인증과 기본 전송 정책이 조립된 SDK 설정을 반환함. 자격증명은 첫 API 요청에서 지연 해석함.

type Credentials struct {
	ApplicationCredentialID     string
	ApplicationCredentialSecret string
}
    Credentials는 KakaoCloud application credential 값임.

type CredentialsProvider interface {
	Resolve(context.Context) (Credentials, error)
}
    CredentialsProvider는 application credential을 지연 해석함.

func NewCredentialsChain(providers ...CredentialsProvider) CredentialsProvider
    NewCredentialsChain은 완전한 자격증명을 반환하는 첫 provider를 사용함.

func NewEnvironmentCredentialsProvider() CredentialsProvider
    NewEnvironmentCredentialsProvider는 KakaoCloud 환경변수에서 자격증명을 읽음.

func NewStaticCredentialsProvider(applicationCredentialID, applicationCredentialSecret string) CredentialsProvider
    NewStaticCredentialsProvider는 코드에서 받은 자격증명을 반환함.

type LoadOptions struct {
	HTTPClient     *http.Client
	Transport      http.RoundTripper
	Credentials    CredentialsProvider
	UserAgent      string
	Region         string
	AuthEndpoint   string
	ConfigEndpoint string
	MaxAttempts    int
	Timeout        time.Duration
	Token          string
	DebugLogging   bool
	// Has unexported fields.
}
    LoadOptions는 LoadDefaultConfig의 편의 client 조립 설정임.

type LoadOptionsFunc func(*LoadOptions) error
    LoadOptionsFunc는 LoadDefaultConfig 설정을 변경함.

func WithAuthEndpoint(authEndpoint string) LoadOptionsFunc
    WithAuthEndpoint는 IAM token 발급 endpoint를 교체함.

func WithConfigEndpoint(endpoint string) LoadOptionsFunc
    WithConfigEndpoint는 서비스 주소를 내려주는 Island(config 조회 서비스) endpoint를 교체함.
    stage 등 비공개 주소로 테스트할 때 씀 — 기본값은 Region으로 계산됨.

func WithCredentials(applicationCredentialID, applicationCredentialSecret string) LoadOptionsFunc
    WithCredentials는 명시 자격증명을 환경변수보다 먼저 해석함.

func WithCredentialsProvider(provider CredentialsProvider) LoadOptionsFunc
    WithCredentialsProvider는 자격증명 출처를 직접 지정함.

func WithDebugLogging(enabled bool) LoadOptionsFunc
    WithDebugLogging은 실제 요청·응답을 표준 에러로 출력함 (기본 꺼짐).

    메서드·URL·상태·소요 시간과 헤더를 찍음. 인증 헤더는 값을 가림. 본문은 크기가 커질 수 있어 찍지 않음.

func WithHTTPClient(client *http.Client) LoadOptionsFunc
    WithHTTPClient는 인증과 재시도 아래에서 사용할 HTTP client를 지정함.

func WithMaxAttempts(maxAttempts int) LoadOptionsFunc
    WithMaxAttempts는 최초 요청을 포함한 최대 시도 횟수를 지정함.

func WithRegion(region string) LoadOptionsFunc
    WithRegion은 인증과 서비스 endpoint를 결정할 KakaoCloud 리전을 지정함.

func WithTimeout(timeout time.Duration) LoadOptionsFunc
    WithTimeout은 시도 1회당 timeout을 지정함. 재시도를 포함한 호출 전체가 아님.

func WithToken(token string) LoadOptionsFunc
    WithToken은 이미 발급된 token을 사용하고 IAM token 발급과 갱신을 생략함.

func WithTransport(transport http.RoundTripper) LoadOptionsFunc
    WithTransport는 HTTP client의 나머지 정책을 유지하면서 기저 RoundTripper만 교체함.

func WithUserAgent(userAgent string) LoadOptionsFunc
    WithUserAgent는 SDK User-Agent를 지정함.
```
