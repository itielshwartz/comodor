package db

import (
	"awesomeProject/models"
	"context"
	"database/sql"
	log "github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/boil"
)

var idx = []string{models.ReleaseColumns.Namespace, models.ReleaseColumns.Name, models.ReleaseColumns.Cluster}

func SaveReleases(db *sql.DB, ctx context.Context, rlsRaw []*models.Release) {
	for _, item := range rlsRaw {
		err := item.Upsert(ctx, db, true, idx, boil.Infer(), boil.Infer())
		if err != nil {
			log.Error(err)
		}
	}
}
