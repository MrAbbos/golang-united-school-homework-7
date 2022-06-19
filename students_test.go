package coverage

import (
	"os"
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

const (
	stringErrorFormat  = "Expected: %s, got %s"
	intErrorFormat     = "Expected: %d, got %d"
	valueErrorFormat = "Expected: %v, got %v"
)

func prepareTestData() People {
	var peopleCollection People
	var timeStamp = time.UTC
	peopleCollection = append(peopleCollection, Person{
		"Abbos",
		"Abbosov",
		time.Date(2002, time.July, 21, 5, 8, 0, 0, timeStamp),
	})

	peopleCollection = append(peopleCollection, Person{
		"Ivan",
		"Ivanov",
		time.Date(2004, time.March, 28, 12, 21, 12, 21, timeStamp),
	})
	peopleCollection = append(peopleCollection, Person{
		"Alijon",
		"Alijonov",
		time.Date(2000, time.March, 2, 2, 0, 0, 0, timeStamp),
	})
	peopleCollection = append(peopleCollection, Person{
		"Kimdir",
		"Kimdirovich",
		time.Date(2000, time.January, 1, 1, 1, 1, 1, timeStamp),
	})
	peopleCollection = append(peopleCollection, Person{
		"Usha",
		"Ushov",
		time.Date(2022, time.February, 1, 2, 3, 4, 5, timeStamp),
	})
	peopleCollection = append(peopleCollection, Person{
		"Karochi",
		"Karochov",
		time.Date(2002, time.January, 2, 0, 0, 2, 0, timeStamp),
	})
	return peopleCollection
}

func TestLenFunctionIsEmpty(t *testing.T) {
	var people People
	var expected = 0
	var actual = people.Len()
	if expected != actual {
		t.Errorf(intErrorFormat, expected, actual)
	}
}

func TestLenFunctionIsNotEmpty(t *testing.T) {
	var people = prepareTestData()
	var expected = 7
	var actual = people.Len()
	if expected != actual {
		t.Errorf(intErrorFormat, expected, actual)
	}
}

func TestLessFunction(t *testing.T) {
	tData := []struct {
		I        int
		J        int
		Expected bool
	}{
		{I: 1, J: 2, Expected: true},  
		{I: 1, J: 0, Expected: true},  
		{I: 4, J: 3, Expected: false}, 
		{I: 3, J: 4, Expected: true},  
		{I: 4, J: 5, Expected: false}, 
		{I: 5, J: 4, Expected: true},  
		{I: 5, J: 6, Expected: false}, 
		{I: 5, J: 5, Expected: false}, 
	}
	var people = prepareTestData()
	for _, v := range tData {
		var actual = people.Less(v.I, v.J)
		if v.Expected != actual {
			t.Errorf(valueErrorFormat, v.Expected, actual)
		}
	}
}

func TestSwapFunction(t *testing.T) {
	tData := []struct {
		FirstIndex  int
		SecondIndex int
	}{
		{FirstIndex: 0, SecondIndex: 3},
		{FirstIndex: 3, SecondIndex: 3},
		{FirstIndex: 5, SecondIndex: 0},
	}
	for _, v := range tData {
		var people = prepareTestData()
		var originalCollection = prepareTestData()

		people.Swap(v.FirstIndex, v.SecondIndex)

		var expected = originalCollection[v.SecondIndex].firstName
		var actual = people[v.FirstIndex].firstName

		if actual != expected {
			t.Errorf(stringErrorFormat, expected, actual)
		}
		expected = originalCollection[v.FirstIndex].firstName
		actual = people[v.SecondIndex].firstName
		if actual != expected {
			t.Errorf(stringErrorFormat, expected, actual)
		}
	}
}
