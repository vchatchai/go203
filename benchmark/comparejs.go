package benchmark

import (
	"fmt"
	"time"
)

func Loop() {
	const cycles = 1000000000
	var result uint64 = 0
	start := time.Now()
	for i := 0; i < cycles; i++ {
		result = result + uint64(i)
	}
	end := time.Now()

	duration := end.Sub(start)
	fmt.Printf("Go looped %d times in %v seconds %v\n", cycles, time.Duration(duration), result)

}

/*
function main(){

    const cycles = 1000000000;
    let start = Date.now();
    for (let i =0;i < cycles; i++){

    }

    let end = Date.now();
    let duration = (end - start) / 1000;
    console.log("JavaScript looped %d times in %d seconds", cycles, duration)
}

main()

*/
