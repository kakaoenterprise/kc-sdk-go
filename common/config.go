package common

import "net/http"

// Endpoints 각 서비스 엔드포인트
type Endpoints struct {
	Image            string
	Volume           string
	BCS              string
	VPC              string
	Network          string
	LoadBalancer     string
	IAM              string
	KubernetesEngine string
}

// Config SDK 전역 설정
type Config struct {
	Endpoints  Endpoints    // 필수: 서비스별 Base URL
	HTTPClient *http.Client // 선택: 기본값 nil이면 http.DefaultClient 사용
	Token      string       // 선택: 있으면 모든 요청에 자동 주입
	UserAgent  string       // 선택: 있으면 모든 요청에 자동 주입
}
