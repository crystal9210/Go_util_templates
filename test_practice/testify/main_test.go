package main

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type mockProbe struct {
	mock.Mock
}

func (m *mockProbe) Probe(url string) bool {
	args := m.Called(url)
	return args.Bool(0)
}

type ProbeSuite struct {
	suite.Suite
	m *mockProbe
}

func (s *ProbeSuite) SetupTest() {
	s.m = new(mockProbe)
	defaultProbe = s.m
}

func (s *ProbeSuite) TestSuccess() {
	s.m.On("Probe", "http://example.com").Return(true)

	s.HTTPStatusCode(handleProbe, "GET", "/probe", nil, http.StatusOK)
	s.HTTPBodyContains(handleProbe, "GET", "/probe", nil, "probe_success 1\n")
}

func (s *ProbeSuite) TestFail() {
	s.m.On("Probe", "http://example.com").Return(false)

	s.HTTPStatusCode(handleProbe, "GET", "/probe", nil, http.StatusOK)
	s.HTTPBodyContains(handleProbe, "GET", "/probe", nil, "probe_success 0\n")
}

func TestProbeSuite(t *testing.T) {
	suite.Run(t, new(ProbeSuite))
}
