package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombinations(t *testing.T) {
	assert.Equal(t, [][]rune{
		{'a', 'a', 'a'},
		{'a', 'a', 'b'},
		{'a', 'b', 'a'},
		{'a', 'b', 'b'},
		{'b', 'a', 'a'},
		{'b', 'a', 'b'},
		{'b', 'b', 'a'},
		{'b', 'b', 'b'},
	}, Combinations([]rune{'a', 'b'}, 3))
}

func BenchmarkToIntegerSlice(b *testing.B) {
	input := "1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20"
	for i := 0; i < b.N; i++ {
		ToIntegerSlice(input, ",")
	}
}

func BenchmarkToRuneSlice(b *testing.B) {
	input := "abc,def,ghi,jkl,mno,pqr,stu,vwx,yz"
	for i := 0; i < b.N; i++ {
		ToRuneSlice(input, ",")
	}
}

func BenchmarkToStringSlice(b *testing.B) {
	input := "a,b,c,d,e,f,g,h,i,j,k,l,m,n,o,p,q,r,s,t,u,v,w,x,y,z"
	for i := 0; i < b.N; i++ {
		ToStringSlice(input, ",")
	}
}

func TestToIntegerSlice(t *testing.T) {
	result, _ := ToIntegerSlice("1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20", ",")

	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, result)
}

func TestToRuneSlice(t *testing.T) {
	result := ToRuneSlice("abc,def,ghi,jkl,mno,pqr,stu,vwx,yz", ",")

	assert.Equal(t, [][]rune{{'a', 'b', 'c'}, {'d', 'e', 'f'}, {'g', 'h', 'i'}, {'j', 'k', 'l'}, {'m', 'n', 'o'}, {'p', 'q', 'r'}, {'s', 't', 'u'}, {'v', 'w', 'x'}, {'y', 'z'}}, result)
}

func TestToStringSlice(t *testing.T) {
	result := ToStringSlice("a,b,c,d,e,f,g,h,i,j,k,l,m,n,o,p,q,r,s,t,u,v,w,x,y,z", ",")

	assert.Equal(t, []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}, result)
}
