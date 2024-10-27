package input

import "scibat/src/shared/domain/event"

type Adapter interface {
	InitCommand() event.Command
}
