package services

import (
	"github.com/fadyat/avito-internship-2022/internal/persistence/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	mockedHealthRepo = &mock.HealthRepo{}
)

func TestHealthService_Ping(t *testing.T) {
	h := NewHealthService(mockedHealthRepo)
	require.NoError(t, h.Ping())
}
