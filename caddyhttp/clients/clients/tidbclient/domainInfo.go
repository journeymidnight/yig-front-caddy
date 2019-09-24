package tidbclient

import (
	"database/sql"
	. "github.com/journeymidnight/yig-front-caddy/caddyerrors"
	"github.com/journeymidnight/yig-front-caddy/caddyhttp/clients/types"
)

func (DB *TidbClient) GetDomainOfBucketDomain(domainHost string) (info types.DomainInfo, err error) {
	var domainBucket string
	sql := info.GetDomainOfBucketDomain()
	args := []interface{}{domainHost}
	row := DB.ClientBusiness.QueryRow(sql, args...)
	err = row.Scan(&domainBucket)
	info.DomainBucket = domainBucket
	return
}

func (DB *TidbClient) GetBucket(bucket string) (uid string, err error) {
	sql := "select uid from buckets where bucketname=?"
	args := []interface{}{bucket}
	row := DB.ClientS3.QueryRow(sql, args...)
	err = row.Scan(&uid)
	if err != nil {
		return uid, ErrNoSuchKey
	}
	return uid, nil
}

func (DB *TidbClient) GetDomain(projectId string, domainHost string) (info types.DomainInfo, err error) {
	var pid, domainH, domainB string
	sql := info.GetDomain()
	args := []interface{}{projectId, domainHost}
	row := DB.ClientBusiness.QueryRow(sql, args...)
	err = row.Scan(&pid, &domainH, &domainB)
	if err != nil {
		return info, ErrNoSuchKey
	}
	info.ProjectId = pid
	info.DomainHost = domainH
	info.DomainBucket = domainB
	return
}

func (DB *TidbClient) GetDomainInfos(projectId string, bucketDomain string) (info []types.DomainInfo, err error) {
	sql := "select * from custom_domain where project_id=? and bucket_domain=?"
	args := []interface{}{projectId, bucketDomain}
	rows, err := DB.ClientBusiness.Query(sql, args...)
	if err != nil {
		return info, ErrInvalidSql
	}
	for rows.Next() {
		ICustomDomain := types.DomainInfo{}
		err = rows.Scan(&ICustomDomain.ProjectId, &ICustomDomain.DomainHost, &ICustomDomain.DomainBucket)
		info = append(info, ICustomDomain)
	}
	if err != nil {
		return info, ErrNoSuchKey
	}
	defer rows.Close()
	return
}

func (DB *TidbClient) GetAllDomainInfos(projectId string) (info []types.DomainInfo, err error) {
	sql := "select * from custom_domain where project_id=?"
	args := []interface{}{projectId}
	rows, err := DB.ClientBusiness.Query(sql, args...)
	if err != nil {
		return info, ErrInvalidSql
	}
	for rows.Next() {
		ICustomDomain := types.DomainInfo{}
		err = rows.Scan(&ICustomDomain.ProjectId, &ICustomDomain.DomainHost, &ICustomDomain.DomainBucket)
		info = append(info, ICustomDomain)
	}
	if err != nil {
		return info, ErrNoSuchKey
	}
	defer rows.Close()
	return
}

func (DB *TidbClient) InsertDomain(info types.DomainInfo) (err error) {
	var sqlTx *sql.Tx
	var tx interface{}
	tx, err = DB.ClientBusiness.Begin()
	if err != nil {
		return ErrSqlTransaction
	}
	defer func() {
		if err == nil {
			err = sqlTx.Commit()
		}
		if err != nil {
			sqlTx.Rollback()
		}
	}()
	sqlTx, _ = tx.(*sql.Tx)
	sql, args := info.InsertDomain()
	_, err = sqlTx.Exec(sql, args...)
	if err != nil {
		return ErrSqlInsert
	}
	return nil
}

func (DB *TidbClient) DelDomain(info types.DomainInfo) (err error) {
	var sqlTx *sql.Tx
	var tx interface{}
	tx, err = DB.ClientBusiness.Begin()
	if err != nil {
		return ErrSqlTransaction
	}
	defer func() {
		if err == nil {
			err = sqlTx.Commit()
		}
		if err != nil {
			sqlTx.Rollback()
		}
	}()
	sqlTx, _ = tx.(*sql.Tx)
	sql, args := info.DeleteDomain()
	_, err = sqlTx.Exec(sql, args...)
	if err != nil {
		return err
	}
	return nil
}