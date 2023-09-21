// This folder contains information related to the app's context, particularly related to client connections within a multiverse

package common

import (
	"context"

	client "github.com/taubyte/dreamland/service"
)

type Context struct {
	Ctx        context.Context
	Multiverse *client.Client
}
