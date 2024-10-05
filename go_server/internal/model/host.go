package model

import (
	"context"
	"encoding/json"
	"log"
	"themynet/internal/db"
	"themynet/internal/db/sqls/sqls"
	"time"
)

const (
	TABLE_NAME = "hosts"
)

// Creates a host and returns the hostID
func Createhost(createHostParams sqls.CreateHostParams) (host sqls.Host, err error) {
	// TODO: make createdAt a DEFAULT CURRENT_DATE
	createHostParams.Createdat.Time = time.Now()
	createHostParams.Recordedat.Time = time.Now()
	createHostParams.Createdby.String = "HamnYce"

	ctx := context.Background()
	queries := sqls.New(db.DBSingleton())

	host, err = queries.CreateHost(ctx, createHostParams)

	return
}

func Retrievehost(hostID int64) (host *sqls.Host, err error) {
	ctx := context.Background()
	queries := sqls.New(db.DBSingleton())

	retrievedhost, err := queries.GetHost(ctx, hostID)
	if err != nil {
		return
	}

	host = &retrievedhost

	return
}

func Retrievehosts(amount, offset uint) (hosts []sqls.Host, err error) {
	ctx := context.Background()
	queries := sqls.New(db.DBSingleton())

	hosts, err = queries.ListHosts(ctx, sqls.ListHostsParams{Limit: int64(amount), Offset: int64(offset)})

	return
}

func Updatehosts(updateHostParamsList []sqls.UpdateHostParams) (hosts []sqls.Host, err error) {
	ctx := context.Background()
	queries := sqls.New(db.DBSingleton())

	for _, updateHostParams := range updateHostParamsList {
		host, err := queries.UpdateHost(ctx, updateHostParams)
		if err != nil {
			log.Fatal(err)
		}
		hosts = append(hosts, host)
	}

	return
}

func Deletehost(hostID int64) (err error) {
	ctx := context.Background()
	queries := sqls.New(db.DBSingleton())

	_, err = queries.DeleteHost(ctx, hostID)

	return
}

func HostToMap(h sqls.Host) (m map[string]any, err error) {
	hostJSON, err := json.Marshal(h)
	if err != nil {
		return
	}
	err = json.Unmarshal(hostJSON, &m)
	return
}
