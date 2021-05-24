package service

import "context"

type TestService struct {
}

func NewTestService() *TestService {
	return &TestService{}
}

func (t *TestService) Hello(ctx context.Context) string {
	return "hello"
}
