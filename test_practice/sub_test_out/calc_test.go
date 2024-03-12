package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	tests := []struct {
		name    string
		a       int
		b       int
		want    int
		wantErr bool // テストがFAILすることを期待する場合はtrueに設定
	}{
		{
			name:    "basic test",
			a:       3,
			b:       5,
			want:    8,
			wantErr: false,
		},
		{
			name:    "zero test",
			a:       0,
			b:       0,
			want:    0,
			wantErr: false,
		},
		// negative testをコメントアウトすると全体としてはokとなる
		{
			name:    "negative test",
			a:       -2,
			b:       3,
			want:    1,
			wantErr: true, // このテストはFAILすることを期待
		},
		{
			name:    "overflow test",
			a:       1000000,
			b:       1000000,
			want:    2000000,
			wantErr: false,
		},
		{
			name:    "fail test",
			a:       2,
			b:       2,
			want:    5, // 明らかに間違った期待値を設定してFAILを期待
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sum(tt.a, tt.b)
			if tt.wantErr {
				assert.NotEqual(t, tt.want, got, "Expected test '%s' to fail", tt.name)
			} else {
				assert.Equal(t, tt.want, got, "Expected test '%s' to pass", tt.name)
			}
		})
	}
}
