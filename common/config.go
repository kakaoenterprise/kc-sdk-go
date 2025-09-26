package common

import "net/http"

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

type Config struct {
	Endpoints  Endpoints
	HTTPClient *http.Client
	Token      string
	UserAgent  string
}
