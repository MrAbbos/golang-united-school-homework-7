package coverage

import (
	"os"
	"reflect"
	"time"
	"testing"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW

const typeError = "Expected %t , got %t"


func TestLen(t *testing.T){
	var people = People{
		{firstName:"Mirabbos",lastName:"Botirjonov",birthDay: time.Now() },
		{firstName:"Abbos",lastName:"Abbosov",birthDay: time.Now() },
	}
	result :=people.Len()
	expected := len(people)

	if result!=expected{
		t.Errorf("Expected %d , got %d",expected,result)
	}
}

func TestLessBirthDay(t *testing.T){
	var people = People{
		{firstName:"Mirabbos",lastName:"Botirjonov",birthDay: time.Now() },
		{firstName:"Abbos",lastName:"Abbosov",birthDay: time.Now().AddDate(1,1,1)},
	}
	result :=people.Less(0,1)
	expected := false
	if result!=expected{
		t.Errorf(typeError,expected,result)
	}
}

func TestLessLastName(t *testing.T){
	var people = People{
		{firstName:"Mirabbos",lastName:"Botirjonov",birthDay: time.Now() },
		{firstName:"Mirabbos",lastName:"dAbbosov",birthDay: time.Now()},
	}
	result :=people.Less(0,1)
	expected := true
	if result!=expected{
		t.Errorf(typeError,expected,result)
	}
}

func TestLessFirstName(t *testing.T){
	var people = People{
		{firstName:"Mirabbos",lastName:"Botirjonov",birthDay: time.Now() },
		{firstName:"Abbos",lastName:"Abbosov",birthDay: time.Now()},
	}
	result :=people.Less(0,1)
	expected := false
	if result!=expected{
		t.Errorf(typeError,expected,result)
	}
}

func TestSwap(t *testing.T){
	var people = People{
		{firstName:"Mirabbos",lastName:"Botirjonov",birthDay: time.Now() },
		{firstName:"Abbos",lastName:"Botirjonov",birthDay: time.Now()},
	}
	people.Swap(0,1)
	
	if people[0].firstName!="Abbos" || people[1].firstName!="Mirabbos"{
		t.Errorf("Swap error %s",people[0].firstName)
	}
}


func TestNew(t *testing.T) {

	matrix,err := New("1\n2\n3\n4\n5\n6")
	expected:=Matrix{6 ,1 ,[]int{1 ,2, 3, 4, 5, 6}}


	if err!=nil{
		t.Errorf("Error in New procedure %s",err.Error())
	}

	if !reflect.DeepEqual(expected,*matrix){
		t.Error("Matrix not equals ")
	}

}

func TestRows(t *testing.T) {
	m:= Matrix{3 ,2 ,[]int{1 ,2, 3, 4, 5, 6}}
	rows:= m.Rows()
	expected:=[][]int{{1 ,2},{3,4}, {5, 6}}

	if !reflect.DeepEqual(expected,rows){
		t.Error("Rows not equals")	
	}
}
func TestCols(t *testing.T) {

	m:= Matrix{3 ,2 ,[]int{1 ,2, 3, 4, 5, 6}}
	cols:= m.Cols()
	expected:=[][]int{{1 ,3, 5}, {2 ,4, 6}}

	if !reflect.DeepEqual(expected,cols){
		t.Error("Colls not equals")	
	}


}


func TestSet(t *testing.T) {
	m:= Matrix{3 ,2 ,[]int{1 ,2, 3, 4, 5, 6}}	
	set_result:= m.Set(1,1,9)
	expected:=Matrix{3 ,2 ,[]int{1 ,2, 3, 9, 5, 6}}

	if (!reflect.DeepEqual(expected,m))||set_result==false{
		t.Error("Values not equals")	
	}
}

func TestSetOverflow(t *testing.T) {
	m:= Matrix{3 ,2 ,[]int{1 ,2, 3, 4, 5, 6}}	
	set_result:= m.Set(1,2,9)
	
	if set_result==true{
		t.Error("Set result must be false")	
	}
}