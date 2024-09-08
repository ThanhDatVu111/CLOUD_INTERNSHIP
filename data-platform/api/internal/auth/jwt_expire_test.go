package auth

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetJwtDuration(t *testing.T) {
	tests := []struct {
		name             string
		jwtExpiryStr     string
		expectedDuration time.Duration
		expectError      bool
	}{
		{"Valid Duration", "5", 5 * time.Minute, false},
		{"Empty String", "", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			duration, err := GetJwtDuration(tt.jwtExpiryStr, time.Minute)
			if tt.expectError {
				assert.Error(t, err, "GetJwtDuration(%v) expected error but got none", tt.jwtExpiryStr)
			} else {
				assert.NoError(t, err, "GetJwtDuration(%v) unexpected error", tt.jwtExpiryStr)
				assert.Equal(t, tt.expectedDuration, *duration, "GetJwtDuration(%v) = %v, want %v", tt.jwtExpiryStr, *duration, tt.expectedDuration)
			}
		})
	}
}

func TestVerifyJwtExpiry(t *testing.T) {
	tests := []struct {
		name           string
		issuedTime     time.Time
		expiryDuration time.Duration
		expectValid    bool
		expectError    bool
	}{
		{"Valid Duration", time.Now().Add(2 * time.Minute), 10 * time.Minute, true, false},
		{"Expired Duration", time.Now().Add(-30 * time.Minute), -10 * time.Minute, false, true},
		{"Invalid Duration", time.Now(), -1 * time.Minute, false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenString, err := createTestToken(tt.expiryDuration)
			assert.NoError(t, err, "createTestToken() unexpected error")

			valid, err := VerifyJwtExpiry(tokenString, tt.issuedTime)

			if tt.expectError {
				assert.Error(t, err, "[%s] expected error but got none", tt.name)
			} else {
				assert.NoError(t, err, "[%s] unexpected error", tt.name)
				assert.Equal(t, tt.expectValid, valid, "[%s] Validity = %v, want %v", tt.name, valid, tt.expectValid)
			}
		})
	}
}
