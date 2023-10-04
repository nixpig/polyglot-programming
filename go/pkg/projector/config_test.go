package projector_test

import (
	"reflect"
	"testing"

	projector "github.com/nixpig/polyglot-programming/pkg/projector"
)

func getOpts(args []string) *projector.Opts {
	opts := &projector.Opts{
		Args:   args,
		Config: "",
		Pwd:    "",
	}

	return opts
}

func testConfig(t *testing.T, args []string, expectedArgs []string, operation projector.Operation) {
	opts := getOpts(args)

	config, err := projector.NewConfig(opts)
	if err != nil {
		t.Errorf("expected to receive no error, but received: %v", err)
	}

	if config.Operation != operation {
		t.Errorf("expected operation was %v, but actual was %v", operation, config.Operation)
	}

	if !reflect.DeepEqual(expectedArgs, config.Args) {
		t.Errorf("expected args to be %+v, but actual was %+v", expectedArgs, config.Args)
	}

}

func TestConfigPrint(t *testing.T) {
	testConfig(t, []string{}, []string{}, projector.Print)
}

func TestConfigPrintKey(t *testing.T) {
	testConfig(t, []string{"foo"}, []string{"foo"}, projector.Print)
}

func TestConfigAddKeyValue(t *testing.T) {
	testConfig(t, []string{"add", "foo", "bar"}, []string{"foo", "bar"}, projector.Add)
}

func TestConfigRemoveKey(t *testing.T) {
	testConfig(t, []string{"remove", "foo"}, []string{"foo"}, projector.Remove)
}
