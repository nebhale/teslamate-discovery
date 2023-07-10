package units

import (
	"fmt"

	"github.com/spf13/cobra"
)

type RangeType string

const (
	Estimated RangeType = "estimated"
	Ideal     RangeType = "ideal"
	Rated     RangeType = "rated"
)

func (r RangeType) Prefix() string {
	if r == Estimated {
		return "est"
	}

	if r == Ideal {
		return "ideal"
	}

	return "rated"
}

func (r *RangeType) Set(v string) error {
	switch v {
	case "estimated":
		*r = Estimated
	case "ideal":
		*r = Ideal
	case "rated":
		*r = Rated
	default:
		return fmt.Errorf("must be one of estimated, ideal, rated")
	}
	return nil
}

func (r RangeType) String() string {
	return string(r)
}

func (r RangeType) Type() string {
	return "string"
}

func RangeTypeCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return []string{string(Estimated), string(Ideal), string(Rated)}, cobra.ShellCompDirectiveDefault
}
