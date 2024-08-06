package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func slowWebsiteChecker(_ string) bool{
	time.Sleep(20*time.Millisecond)
	return true
}
func BenchmarkCheckWebsites(b * testing.B){
	urls := make([]string,100)
	for i:=0 ; i<len(urls);i++{
		urls[i] = fmt.Sprintf("urlnumber%v.com",i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowWebsiteChecker,urls)
	}
}

