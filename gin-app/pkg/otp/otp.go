package otp

import "github.com/xlzd/gotp"

// One-time password

type Generator interface {
	RandomSecret(length int) string
}

type GOTPGenerator struct{}

func NewGOTPGenerator() *GOTPGenerator {
	return &GOTPGenerator{}
}

func (g *GOTPGenerator) RandomSecret(length int) string {
	return gotp.RandomSecret(length)
}
