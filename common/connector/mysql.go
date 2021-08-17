package connector

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

var DB *MysqlConnector

type DBConfig struct {
	Username string
	Password string
	Hostname string
	DBName   string
}

func NewMysql(cfg *DBConfig) *MysqlConnector {
	m := &MysqlConnector{
		username: cfg.Username,
		password: cfg.Password,
		hostname: cfg.Hostname,
	}
	err := m.connect(cfg.DBName)
	if err != nil {
		logrus.Errorf("fail to connect to database: %s, error:%v", cfg.DBName, err)
	}

	return m
}

type MysqlConnector struct {
	username string
	password string
	hostname string
	db       *sql.DB
}

func (c *MysqlConnector) dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", c.username, c.password, c.hostname, dbName)
}

func (c *MysqlConnector) connect(dbName string) error {
	logrus.Info("connecting to mysql ...")
	db, err := sql.Open("mysql", c.dsn(dbName))
	c.db = db
	if err != nil {
		logrus.Error("fail to connect to mysql!")
		return err
	}
	db.SetConnMaxLifetime(time.Minute * 30)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	// res, err := db.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS "+dbName)
	// if err != nil {
	// 	logrus.Errorf("Error %s when creating DB\n", err)
	// 	return err
	// }
	// no, err := res.RowsAffected()
	// if err != nil {
	// 	logrus.Printf("Error %s when fetching rows", err)
	// 	return err
	// }
	// logrus.Infof("rows affected %d\n", no)

	err = db.PingContext(ctx)
	if err != nil {
		logrus.Errorf("fail to connect to db: %s, error:%v", dbName, err)
		return err
	}
	logrus.Info("connected to ", dbName)

	// // create table
	// query := `CREATE TABLE IF NOT EXISTS product(product_id int primary key auto_increment, product_name text, product_price int, created_at datetime default CURRENT_TIMESTAMP, updated_at datetime default CURRENT_TIMESTAMP)`
	// res, err := db.ExecContext(ctx, query)
	// if err != nil {
	// 	logrus.Panicf("Error %s when creating product table, %v", err, res)
	// 	return err
	// }
	return nil
}

func (c *MysqlConnector) Close() error {
	return c.db.Close()
}
