package invoke

import (
	"fmt"
)

func (srv *InvokeSrv) GetLogFilename(id string) string {
	return fmt.Sprintf("%s.json", id)
}
