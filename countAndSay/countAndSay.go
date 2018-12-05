package main

func countAndSay(n int) string {
	var res []byte
	var next []byte
	res = append(res, '1')
	for ; n-1 > 0; n-- {
		start := 0
		next = make([]byte, 0)
		var i int
		for i = 0; i < len(res); i++ {
			if res[i] != res[start] {
				next = append(next, byte(i-start)+'0', res[start])
				start = i
			}
		}
		next = append(next, byte(i-start)+'0', res[start])
		res = next
	}
	return string(res)
}
