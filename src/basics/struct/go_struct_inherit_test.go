package go_struct

import "testing"

func TestStructA(t *testing.T) {
	person := Person{"Tom", 20}
	person.Eat()
	person.Walk()

	pPerson := &Person{"Tom", 20}
	pPerson.name = "Jerry"
	pPerson.age = 30

	student := Student{Person{"Jack", 19}, "MIT"}
	student.Eat()
	student.Study()

	teacher := Teacher{
		Person: Person{"Tom", 20},
		school: "MIT",
	}
	teacher.Eat()
	teacher.school = "Harvard"
}
