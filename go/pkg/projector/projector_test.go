package projector_test

import (
	"testing"

	projector "github.com/nixpig/polyglot-programming/pkg/projector"
)

func getData() *projector.Data {
	data := &projector.Data{
		Projector: map[string]map[string]string{
			"/home/username/": {
				"key1": "value1",
				"key2": "value2",
			},
			"/home/username/project1": {
				"key1": "value3",
				"key2": "value4",
			},
			"/home/username/project1/subproject1": {
				"key1": "value5",
				"key2": "value6",
			},
			"/home/username/project2": {
				"key1": "value7",
				"key2": "value8",
			},
			"/home/username/project2/subproject2": {
				"key1": "value9",
				"key2": "value10",
			},
		},
	}

	return data
}

func getProjector(pwd string, data *projector.Data) *projector.Projector {
	return projector.CreateProjector(&projector.Config{
		Args:      []string{},
		Operation: projector.Print,
		Pwd:       pwd,
		Config:    "Hello, world!",
	}, data)
}

func test(t *testing.T, proj projector.Projector, key string, expectedValue string) {

	value, _ := proj.GetValue(key)
	if value != expectedValue {
		t.Errorf("expected to find value of '%v' for '%v', but found '%v'", expectedValue, key, value)
	}

}

func TestGetValue(t *testing.T) {
	data := getData()
	proj1 := getProjector("/home/username/", data)
	proj2 := getProjector("/home/username/project1/subproject1", data)

	test(t, *proj1, "key2", "value2")
	test(t, *proj2, "key1", "value5")
}

func TestSetValue(t *testing.T) {
	data := getData()
	proj := getProjector("/home/username/", data)

	proj.SetValue("foo", "bar")

	test(t, *proj, "foo", "bar")
}

func TestRemoveValue(t *testing.T) {
	data := getData()
	proj := getProjector("/home/username/", data)

	// ensure it actually exists before trying to delete so we can verify something has actually changed!
	test(t, *proj, "key2", "value2")

	proj.RemoveValue("key2")

	test(t, *proj, "key2", "")

}
