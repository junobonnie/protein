package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	e0 := 0
	minE := 0
	//protein := randomProtein(111)
	answer := "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAALLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLL"
	count := 0
	T := 2.0

	var err error
	if e0, err = calculator(answer, false); err != nil {
		fmt.Println("Error in calculator:", err)
		return
	}
	runtime.GOMAXPROCS(5) //runtime.NumCPU()/3 + 1)
	fmt.Println(runtime.GOMAXPROCS(0))

	for minE > -403 {
		des := []int{0}
		candidates := []string{answer}
		count++
		e1 := e0

		var wg sync.WaitGroup
		var mu sync.Mutex

		// Launch 50 goroutines to execute boltzmanFold concurrently
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				result, eResult := boltzmanFold(answer, 3.0, 0.05, false, false, false)
				fmt.Println(result, eResult)

				mu.Lock()
				des = append(des, eResult-e0)
				candidates = append(candidates, result)
				if eResult < minE {
					minE = eResult
					fmt.Println("New minimum found:", result, minE)
					// 파일을 추가 모드로 열기
					file, err := os.OpenFile("result.txt", os.O_APPEND|os.O_WRONLY, 0644)
					if err != nil {
						fmt.Println("파일 열기 중 오류:", err)
						return
					}
					defer file.Close()

					// 파일에 문자열 추가
					memo := fmt.Sprintf("\nNew minimum found: %s %d", result, minE)
					_, err = file.WriteString(memo + "\n")
					if err != nil {
						fmt.Println("파일 쓰기 중 오류:", err)
						return
					}
				}
				mu.Unlock()
			}()
		}

		// Wait for all goroutines to finish
		wg.Wait()

		answer = dice(candidates, des, T)
		e0, err = calculator(answer, false)
		if err != nil {
			fmt.Println("Error in calculator:", err)
			return
		}
		if e0 == e1 {
			T += 1.0
		} else if e0 > e1 {
			T -= 0.5
		} else {
			T -= 0.2
		}
		if T < 0.5 {
			T = 0.5
		}
		fmt.Println(T)
		fmt.Printf("%d >> %s %d\n", count, answer, e0)
	}
}

// Generates a random protein sequence of length n using characters "A", "B", and "C"
func randomProtein(n int) string {
	var sb strings.Builder
	for i := 0; i < n; i++ {
		sb.WriteByte("ABC"[rand.Intn(3)])
	}
	return sb.String()
}
