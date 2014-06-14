package problems

import "testing"

func TestMultiplesOf3And5(t *testing.T) {
	const in, out = 10, 23
	if x := MultiplesOf3And5(in); x != out {
		t.Errorf("MultiplesOf3And5(%v) = %v, want %v", in, x, out)
	}
}

func TestLargestPalindromeProduct(t *testing.T) {
	const in1, in2, out = 1, 99, 9009

	if x := LargestPalindromeProduct(in1, in2); x != out {
		t.Errorf("LargestPalindromeProduct(%v, %v) = %v, want %v", in1, in2, x, out)
	}
}

func TestReverseString(t *testing.T) {
	var in, out = []string{"Test", "Fuenf", "54"}, []string{"tseT", "fneuF", "45"}

	for i := 0; i < len(in); i++ {
		if x := reverseString(in[i]); x != out[i] {
			t.Errorf("TestReverseString(%v) = %v, want %v", in[i], x, out[i])
		}
	}
}

func TestIsPalindrome(t *testing.T) {
	var in, out = []int{9009, 200, 4004, 54}, []bool{true, false, true, false}

	for i := 0; i < len(in); i++ {
		if x := isPalindrome(in[i]); x != out[i] {
			t.Errorf("TestIsPalindrome(%v) = %v, want %v", in[i], x, out[i])
		}
	}
}
