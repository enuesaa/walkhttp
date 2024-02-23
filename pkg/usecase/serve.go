package usecase

import (
	"github.com/enuesaa/walkin/pkg/servectl"
)

func Serve() error {
	ctl := servectl.New()
	
	if err := ctl.Run(); err != nil {
		return err
	}
	return nil
}
