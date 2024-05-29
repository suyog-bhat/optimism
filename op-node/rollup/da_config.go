package rollup

import (
	"cloud.google.com/go/bigtable"
	"context"
	"google.golang.org/api/option"
)

type DAConfig struct {
	GcpProjectId       string
	GcpAuthFileKeyPath string
}

func NewDAClient(daConfig DAConfig) *bigtable.Table {
	ctx := context.Background()
	client, err := bigtable.NewClient(ctx, daConfig.GcpProjectId, "op-data", option.WithCredentialsFile(daConfig.GcpAuthFileKeyPath))
	if err != nil {
		return nil
	}
	tbl := client.Open("tx-DA")
	return tbl
}
