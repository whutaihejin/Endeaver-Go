package word

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func TestPalindrome(t *testing.T) {
	assert.True(t, IsPalindrome("detartrated"))
	assert.True(t, IsPalindrome("kayak"))
}

func TestNonPalindrome(t *testing.T) {
	assert.False(t, IsPalindrome("palindrome"))
}

func TestFrenchPalindrome(t *testing.T) {
	assert.True(t, IsPalindrome("été"))
}

func TestCanalPalindrome(t *testing.T) {
	assert.True(t, IsPalindrome("A man, a plan, a canal: Panama"))
}

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"detartrated", true},
		{"A man, a plan, a canal: Panama", true},
		{"Evil I did dwell; lewd did I live", true},
		{"Able was I ere I saw Elba", true},
		{"été", true},
		{"Et se resservir, ivresse reste.", true},
		{"Palindrome", false},
		{"desserts", false},
	}
	for _, test := range tests {
		assert.Equal(t, test.want, IsPalindrome(test.input))
	}
}

func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

func TestRandomPalindrome(t *testing.T) {
	// Initialize a pseudo-random number generator
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		assert.True(t, IsPalindrome(p))
	}
}
