package main
import(
	"fmt"
	"time"
)

func main (){
	go  Greeter ("hello")
	 Greeter ("world")
}

func Greeter(s string){
	for i := 0; i < 6; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(s);

	}
}