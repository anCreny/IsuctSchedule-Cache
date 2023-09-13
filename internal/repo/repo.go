package repo

import (
	"fmt"
	"github.com/restream/reindexer/v3"
	_ "github.com/restream/reindexer/v3/bindings/cproto"
	"main/config"
	"main/internal/repo/structs"
)

var r *Repo = nil

type Repo struct {
	Rx *reindexer.Reindexer
}

func Init() error {
	var userPath string
	var cfg = config.Cfg.RxCfg
	r = &Repo{}

	if len(cfg.Username) > 0 {
		userPath = fmt.Sprintf("%v:%v@", cfg.Username, cfg.Password)
	}

	connectionPath := fmt.Sprintf("cproto://%v%v:%v/%v", userPath, cfg.Host, cfg.Port, cfg.Database)
	r.Rx = reindexer.NewReindex(connectionPath, reindexer.WithCreateDBIfMissing())

	if err := r.Rx.OpenNamespace(cfg.Namespaces.Teachers, reindexer.DefaultNamespaceOptions(), structs.Timetable{}); err != nil {
		return fmt.Errorf("error occurred while openning %s namespace: %s", cfg.Namespaces.Teachers, err.Error())
	}

	if err := r.Rx.OpenNamespace(cfg.Namespaces.Groups, reindexer.DefaultNamespaceOptions(), structs.Timetable{}); err != nil {
		return fmt.Errorf("error occured while openning %s namespace:, %s", cfg.Namespaces.Groups, err.Error())
	}

	if err := r.Rx.OpenNamespace(cfg.Namespaces.Names, reindexer.DefaultNamespaceOptions(), structs.TeachersNames{}); err != nil {
		return fmt.Errorf("error occurred while openning %s namespace:, %s", cfg.Namespaces.Names, err.Error())
	}

	return nil

}
