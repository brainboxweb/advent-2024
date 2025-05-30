package branding

// // Move to external
// func TestNew(t *testing.T) {
// 	towels := NewTowels([]string{"r", "wr", "b", "g", "bwu", "rb", "gb", "br"})
// 	tests := []struct {
// 		pattern  string
// 		expected bool
// 	}{
// 		{
// 			"brwrr",
// 			true,
// 		},
// 		{
// 			"bggr",
// 			true,
// 		},
// 		{
// 			"gbbr",
// 			true,
// 		},
// 		{
// 			"rrbgbr",
// 			true,
// 		},
// 		{
// 			"ubwu",
// 			false,
// 		},
// 		{
// 			"bwurrg",
// 			true,
// 		},
// 		{
// 			"brgr",
// 			true,
// 		},
// 		{
// 			"bbrgwb",
// 			false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run("test", func(t *testing.T) {
// 			result := towels.match(tt.pattern)
// 			assert.Equal(t, tt.expected, result)
// 		})
// 	}
// }
