package main
import "fmt"

func main(){
	name := "Sarin";
	condition := true;
	i:=0
	for condition {
		fmt.Println("Hello",name);
		i++;
		
		if(i>5){
			condition = false;
		}
	}
}