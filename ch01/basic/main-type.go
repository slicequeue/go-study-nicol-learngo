package main // 패키지를 선언 기능 묶음
import (
	"fmt"
	"strings"
)

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func lenAndUpper2(name string) (length int, uppercase string) { // naked return - 이 기능을 통해 무언가를 또 할 수 있는 것을 만들 수 있음 매우 섹시!
	defer fmt.Println("I'm done") // function  수행하고 나서 무언가를 수행 시킬 수 있다. 이미지를 닫거나 파일을 닫거나 디비연결을 끊거나 등등
	// 이는 효과가 매우 큼! 아주 좋은 기능! - 끝나고 난 뒤에 실행한다!!
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
	// return length, uppercase // 이것도 가능!
}

func superAdd(numbers ...int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}

func canIDrink(age int) bool {
	if age < 18 {
		return false
	}
	return true
}

func canIDrink2(age int) bool {
	if koreanAge := age + 2 ; koreanAge < 18 { // 조건에서만 사용하는 것을 명확하게 할 수 있다! variable expression
		return false
	}
	return true
}

func canIDrink3(age int) bool {
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	
	}
	return false
}

func main1() { // main 은 컴파일을 위해서 필요함
	// fmt.Println("Hello world")
	name := "nico"
	name = "lynn"
	fmt.Println(name)

	fmt.Println(multiply(1,2))

	totalLength1, upperName := lenAndUpper("nico")
	fmt.Println(totalLength1, upperName)

	totalLength2, _ := lenAndUpper("nico")
	fmt.Println(totalLength2)

	repeatMe("nico", "lynn", "dal", "marl", "flynn")

	totalLength3, upperName3 := lenAndUpper2("nico")
	fmt.Println(totalLength3, upperName3)

	total := superAdd(1,2,3,4,5,6)
	fmt.Println(total)

	fmt.Println(canIDrink(16))
	fmt.Println(canIDrink2(16))
	fmt.Println(canIDrink3(18))
}
