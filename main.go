package main
import (
		"fmt"
		"net/http"
)

func ThisWebTest(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Hello World"));
}

func ShowThisSlice (sln []int) int {

		var ThisTmp int;
		for i := 0; i < len(sln); i++ { 
		if(sln[i] == 13){ ThisTmp = 42 } 
		}

		return ThisTmp;
}

func main(){
		var number = 42;
		var num int;
		num = 43;
		fmt.Println(number, " ", num);

var nums [3]int;
nums[0] = 42;

for i := 1; i < 10; i++ { print(nums[0], " "); }

// Slice
slinum := []int{4, 4, 2, 2};
println(slinum[1]);

slinum = append(slinum, 13);
 println("size:", cap(slinum));

// #######

var SomeFrom int = ShowThisSlice(slinum);
println("FROM FUNC: ", SomeFrom);

// #####
// web data

http.HandleFunc("/", ThisWebTest);
err := http.ListenAndServe("localhost:9999", nil)
		if err != nil { 
				println("Listen: ", err);
		}
// ####

}
