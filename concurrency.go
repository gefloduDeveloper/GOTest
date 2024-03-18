package main
import(
	"fmt"
	"time"
	"sync"
)

var dbData = []string{"id1","id2","id3","id4","id5"}
var wg = sync.WaitGroup{}
var results = []string{}
var m = sync.RWMutex{}

func main(){
	t0:= time.Now()
	for i:=0; i<1000; i++{
		wg.Add(1)
		go dbCall2(1)
	}
	wg.Wait()
	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
}

//since the function barely does anything, is not importatn how many times it's calles the time remains the same
func dbCall2(i int){
	var delay float32 = 2000
	time.Sleep(time.Duration(delay)*time.Millisecond)
	wg.Done()
}


func dbCall(i int){
	var delay float32 = 2000
	time.Sleep(time.Duration(delay)*time.Millisecond)
	save(dbData[i])
	log()
	wg.Done()
}

func save(result string){
	m.Lock()
	results = append(results, result)
	m.Unlock()
}

func log(){
	m.RLock()
	fmt.Printf("\nThe results are %v", results)
	m.RUnlock()
}

