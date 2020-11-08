package ioc

import (
	"testing"
)
func TestInject(t *testing.T) {

	ioc := New("test");
	type A struct {
		Name string
		Age int
	}
	ioc.Provide(A{"testa",20});

	type B struct {
		X string `inject:"A.Name"`
		Y int
	}
	ioc.Provide(B{"",100});

	type TestA struct {
		Age int `inject:"A.Age"`
		Name string `inject:"B.X"`
		Y int `inject:"B.Y"`
	}
	a, _ := ioc.Inject(TestA{});

	testA := a.(TestA);

	if testA.Name != "testa" {
		t.Errorf(`testA.Name(%q) = "testa"`, testA.Name)
	}

	if testA.Age != 20 {
		t.Errorf(`testA.Age(%q) = 20`, testA.Age)
	}

	if testA.Y != 100 {
		t.Errorf(`testA.Y(%q) = 100`, testA.Y)
	}
}