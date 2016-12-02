package writer

import (
	"github.com/tonglil/labeler/types"

	"github.com/google/go-github/github"
)

var steps = `
1. Rename: intersection of local (l.From) + remote (l.Name)
2. Update: intersection of local (l.Name) + remote (l.Name)
3. Create: l.Name compliment of remote
4. Delete: l.Name compliment of locale
`

func Run(client *github.Client, file string, opt *types.Options) error {
	// open file
	// unmarshal file contents
	// configure the right repo
	// ep = [https://api.github.com/] repos/ [tonglil/labeler] /labels
	// get all remote labels from repo
	// for l in local labels with "from:"
	// if l doesn't exist
	//   if l.From exists
	//     rename from to it
	//       PATCH ep + /:name <- l.from
	//         {
	//           "name": "l.name",
	//           "color": "l.color"
	//         }
	//   else create it
	//     POST ep
	//       {
	//         "name": "l.name",
	//         "color": "l.color"
	//       }
	// do create non-existing local labels
	// do update existing local labels
	// do delete remaining remote labels

	return
}
