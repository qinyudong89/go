package map_test

import "testing"

func TestInitMap(t *testing.T) {
	students := make(map[string]int)
	t.Log(len(students))
	students["morgan"] = 45
	t.Log(students)
}

func TestGetMap(t *testing.T) {
	students := map[string]int{"bob": 12, "tom": 34}
	t.Log("bob: ", students["bob"])
}

func TestDelMap(t *testing.T) {
	students := map[string]int{"bob": 12, "tom": 34}
	t.Log(students)
	delete(students, "tom")
	t.Log(students)
}

func TestRangMap(t *testing.T) {
	students := map[string]int{"bob": 12, "tom": 34}
	for k, v := range students {
		t.Log(k, v)
	}
}
