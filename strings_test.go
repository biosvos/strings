package strings

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestRegex(t *testing.T) {
	contents, err := os.ReadFile("file.txt")
	require.NoError(t, err)
	_, err = Strings(contents).Regex(`<button type="button" class="(.+?)" on`)
	require.NoError(t, err)
}

func TestJoin(t *testing.T) {
	for _, tt := range []struct {
		args     []string
		expected string
	}{
		{
			args:     []string{},
			expected: "",
		}, {
			args:     []string{"a"},
			expected: "a",
		}, {
			args:     []string{"a", "b"},
			expected: "a,b",
		},
	} {
		t.Run("", func(t *testing.T) {
			join := Join(tt.args, ",")
			require.Equal(t, tt.expected, join)
		})
	}
}

func TestReplace(t *testing.T) {
	body := `* <span class="colour" style="color:rgb(38, 55, 71)">Python과 Python Framework (Django / Flask 등)를 이용한 백엔드 서버 개발 경력 3년 이상</span>`
	replace, err := Strings(body).Replace("(<.+?>)", "")
	require.NoError(t, err)
	require.Equal(t, "* Python과 Python Framework (Django / Flask 등)를 이용한 백엔드 서버 개발 경력 3년 이상", replace)
}
