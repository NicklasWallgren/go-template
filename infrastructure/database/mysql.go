package database

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	sqltrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/database/sql"
	gormtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorm.io/gorm.v1"
	mysqlgorm "gorm.io/driver/mysql"
	"log"
	"time"

	"github.com/NicklasWallgren/go-template/config"
	"github.com/NicklasWallgren/go-template/infrastructure/logger"
	"github.com/avast/retry-go"
	"gorm.io/gorm"
)

// Database modal
type Database struct {
	*gorm.DB
}

// NewDatabase creates a new database instance
func NewDatabase(config *config.AppConfig, logger logger.Logger) Database {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&multiStatements=true", config.Database.User, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Name)

	// TODO, lazy load database connection pool, or fail fast?
	db, err := connect(url, logger)
	if err != nil {
		logger.Info("Url: ", url)
		logger.Panic(err)
	}

	return Database{DB: db}
}

func connect(url string, logger logger.Logger) (db *gorm.DB, err error) {
	// The database might not have been initialized yet, retry
	err = retry.Do(func() (err error) {

		// Register augments the provided driver with tracing, enabling it to be loaded by gormtrace.Open.
		sqltrace.Register("mysql", &mysql.MySQLDriver{}, sqltrace.WithServiceName("go-template"))
		sqlDb, err := sqltrace.Open("mysql", url)
		if err != nil {
			log.Fatal(err)
		}

		db, err = gormtrace.Open(mysqlgorm.New(mysqlgorm.Config{Conn: sqlDb}), &gorm.Config{Logger: logger.GetGormLogger()}, gormtrace.WithServiceName("go-template"))
		//db, err = gormtrace.Open(mysqlgorm.New(mysqlgorm.Config{DSN: url}), &gorm.Config{Logger: logger.GetGormLogger()}, gormtrace.WithServiceName("go-template"))

		//// Create a root span, giving name, server and resource.
		//span, _ := tracer.StartSpanFromContext(context.Background(), "my-query",
		//	tracer.SpanType(ext.SpanTypeSQL),
		//	tracer.ServiceName("my-db"),
		//	tracer.ResourceName("initial-access"),
		//)
		//defer span.Finish()

		return err
	}, retry.Delay(200*time.Millisecond), retry.Attempts(10))

	return db, err
}
