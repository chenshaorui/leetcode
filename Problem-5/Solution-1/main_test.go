package Solution_1

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		wants []string
	}{
		{
			name: "Test",
			args: args{
				s: "babad",
			},
			wants: []string{
				"bab",
				"aba",
			},
		},
		{
			name: "Test",
			args: args{
				s: "cbbd",
			},
			wants: []string{
				"bb",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestPalindrome(tt.args.s)

			isNotIncluded := true
			for _, want := range tt.wants {
				if got == want {
					isNotIncluded = false
					break
				}
			}

			if isNotIncluded {
				t.Errorf("longestPalindrome() = %v, wants %v", got, tt.wants)
			}
		})
	}
}
