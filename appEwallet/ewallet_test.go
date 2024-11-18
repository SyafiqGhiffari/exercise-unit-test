package ewallet

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	pengujian = true
}

func TestJalankanPerintah(t *testing.T) {
	type args struct {
		daftarPerintah []string
	}
	tests := []struct {
		name     string
		args     args
		expected float64
		err      error
	}{
		// error case 1
		{
			name: "tarik tanpa setor",
			args: args{
				daftarPerintah: []string{"withdraw"},
			},
			expected: 0,
			err:      errors.New("saldo anda tidak mencukupi"),
		},
		// test case 1
		{
			name: "setor sekali",
			args: args{
				daftarPerintah: []string{"deposit"},
			},
			expected: 50000,
			err:      nil,
		},
		// test case 2
		{
			name: "setor 2x",
			args: args{
				daftarPerintah: []string{"deposit", "deposit"},
			},
			expected: 150000,
			err:      nil,
		},
		// test case 3
		{
			name: "setor 3x",
			args: args{
				daftarPerintah: []string{"deposit", "deposit", "deposit"},
			},
			expected: 300000,
			err:      nil,
		},
		// test case 4
		{
			name: "setor 2x dan tarik 5x",
			args: args{
				daftarPerintah: []string{"deposit", "deposit", "withdraw", "withdraw", "withdraw", "withdraw", "withdraw"},
			},
			expected: 275000,
			err:      nil,
		},
		// test case 5
		{
			name: "setor 1x dan tarik 9x",
			args: args{
				daftarPerintah: []string{"deposit", "withdraw", "withdraw", "withdraw", "withdraw", "withdraw", "withdraw", "withdraw", "withdraw", "withdraw"},
			},
			expected: 100000,
			err:      nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := JalankanPerintah(tt.args.daftarPerintah)
			assert.Equal(t, res, tt.expected)
			if err != nil {
				assert.Equal(t, err, tt.err)
			}
		})
	}
}
