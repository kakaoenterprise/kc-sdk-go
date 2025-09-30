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
	Endpoints  Endpoints
	HTTPClient *http.Client
	Token      string
	UserAgent  string
}
