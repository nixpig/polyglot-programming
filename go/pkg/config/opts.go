package projector

import "github.com/hellflame/argparse"

type Opts struct {
	Args   []string
	Config string
	Pwd    string
}

func GetOpts() (*Opts, error) {
	parser := argparse.NewParser("projector", "gets all the values", &argparse.ParserConfig{
		DisableDefaultShowHelp: true,
	})

	args := parser.Strings("a", "args", &argparse.Option{
		Positional: true,
		Required:   false,
		Default:    "",
	})

	config := parser.String("c", "config", &argparse.Option{
		Required: false,
		Default:  "",
	})

	pwd := parser.String("p", "pwd", &argparse.Option{
		Positional: false,
		Default:    "",
	})

	err := parser.Parse(nil)
	if err != nil {
		parser.PrintHelp()
		return nil, err
	}

	return &Opts{
		Args:   *args,
		Config: *config,
		Pwd:    *pwd,
	}, nil
}
