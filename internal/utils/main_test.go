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

func BenchmarkToOptimisticIntSlice(b *testing.B) {
	input := "0,9 -> 5,9\n8,0 -> 0,8\n9,4 -> 3,4\n2,2 -> 2,1\n7,0 -> 7,4\n6,4 -> 2,0\n0,9 -> 2,9\n3,4 -> 1,4\n0,0 -> 8,8\n5,5 -> 8,2"
	for i := 0; i < b.N; i++ {
		ToOptimisticIntSlice(input, false)
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

func TestToOptimisticIntSlice(t *testing.T) {
	input := "0,9 -> 5,9\n8,0 -> 0,8\n9,4 -> 3,4\n2,2 -> 2,1\n7,0 -> 7,4\n6,4 -> 2,0\n0,9 -> 2,9\n3,4 -> 1,4\n0,0 -> 8,8\n5,5 -> 8,2"

	assert.Equal(t, []int{0, 9, 5, 9, 8, 0, 0, 8, 9, 4, 3, 4, 2, 2, 2, 1, 7, 0, 7, 4, 6, 4, 2, 0, 0, 9, 2, 9, 3, 4, 1, 4, 0, 0, 8, 8, 5, 5, 8, 2}, ToOptimisticIntSlice(input, false))
}

func TestToRuneSlice(t *testing.T) {
	result := ToRuneSlice("abc,def,ghi,jkl,mno,pqr,stu,vwx,yz", ",")

	assert.Equal(t, [][]rune{{'a', 'b', 'c'}, {'d', 'e', 'f'}, {'g', 'h', 'i'}, {'j', 'k', 'l'}, {'m', 'n', 'o'}, {'p', 'q', 'r'}, {'s', 't', 'u'}, {'v', 'w', 'x'}, {'y', 'z'}}, result)
}

func TestToStringSlice(t *testing.T) {
	result := ToStringSlice("a,b,c,d,e,f,g,h,i,j,k,l,m,n,o,p,q,r,s,t,u,v,w,x,y,z", ",")

	assert.Equal(t, []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}, result)
}
