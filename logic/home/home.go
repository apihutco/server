package home

import (
	"fmt"

	"github.com/apihutco/server/config"
)

func Page() string {
	return fmt.Sprintf("APIHut.\nDocs: %s", config.Conf.Site.DocsUrl)
}
