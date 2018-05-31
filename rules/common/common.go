package common

import "github.com/adhiravishankar/when/rules"

var All = []rules.Rule{
	SlashDMY(rules.Override),
}
