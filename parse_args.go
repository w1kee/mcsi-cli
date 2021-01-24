package main

type cli struct {
	Url string
	All bool
}

func parseArgs(rawArgs []string) *cli {
	if len(rawArgs) == 1 {
		return nil
	}

	switch rawArgs[1] {
	case "-a":
		if len(rawArgs) > 2 {
			return &cli{rawArgs[2], true}
		}
	default:
		return &cli{rawArgs[1], false}
	}

	return nil
}
