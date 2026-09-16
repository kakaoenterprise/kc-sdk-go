# kc-sdk-go/v2/image

KakaoCloud SDK 서비스 진입점(스펙 기반 생성 인터페이스).

## 사용 예

```go
cfg, err := config.LoadDefaultConfig(ctx)
client := image.NewFromConfig(cfg)
```

저수준 빌더 인터페이스(Execute 체인)가 필요하면 `NewAPIClient`를 직접 사용합니다.

!!! warning "임시(POC) 문서"
    이 문서는 public pkg.go.dev 배포 전까지 제공되는 임시 문서입니다. 정식 배포 이후에는
    [pkg.go.dev](https://pkg.go.dev/github.com/kakaoenterprise/kc-sdk-go/v2/image)의
    문서가 공식 레퍼런스가 됩니다.
