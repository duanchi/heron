package abstract

import (
	"context"
	"database/sql"
	"github.com/xormplus/xorm"
	"github.com/xormplus/xorm/caches"
	"github.com/xormplus/xorm/contexts"
	"github.com/xormplus/xorm/core"
	"github.com/xormplus/xorm/dialects"
	"github.com/xormplus/xorm/log"
	"github.com/xormplus/xorm/names"
	"github.com/xormplus/xorm/schemas"
	"go.heurd.com/heron-go/heron/db"
	"io"
	"reflect"
	"strings"
	"time"
)

type Model struct {
	Mapper interface{}
	engine *xorm.Engine `json:"-" xorm:"-"`
	options map[string]interface{} `json:"-" xorm:"-"`
}

func (this *Model) Options (options map[string]interface{}) {
	if options["source"] != nil {
		this.SetSource(options["source"].(string))
	} else {
		this.SetSource("default")
	}
}

func (this *Model) Init() *Model {
	this.Options(map[string]interface{}{})
	return this
}

func (this *Model) GetEngine() *xorm.Engine {
	return this.engine
}

func (this *Model) SetEngine(db *xorm.Engine) {
	this.engine = db
}

func (this *Model) SetSource(name string) {
	this.engine = db.Engine(name)
}

// EnableSessionID if enable session id
func (this *Model) EnableSessionID(enable bool) {
	this.engine.EnableSessionID(enable)
}

// SetCacher sets cacher for the table
func (this *Model) SetCacher(tableName string, cacher caches.Cacher) {
	this.engine.SetCacher(tableName, cacher)
}

// GetCacher returns the cachher of the special table
func (this *Model) GetCacher(tableName string) caches.Cacher {
	return this.engine.GetCacher(tableName)
}

// SetQuotePolicy sets the special quote policy
func (this *Model) SetQuotePolicy(quotePolicy dialects.QuotePolicy) {
	this.engine.SetQuotePolicy(quotePolicy)
}

// BufferSize sets buffer size for iterate
func (this *Model) BufferSize(size int) *xorm.Session {
	return this.engine.BufferSize(size)
}

// ShowSQL show SQL statement or not on logger if log level is great than INFO
func (this *Model) ShowSQL(show ...bool) {
	this.engine.ShowSQL(show ...)
}

// Logger return the logger interface
func (this *Model) Logger() log.ContextLogger {
	return this.engine.Logger()
}

// SetLogger set the new logger
func (this *Model) SetLogger(logger interface{}) {
	this.engine.SetLogger(logger)
}

// SetLogLevel sets the logger level
func (this *Model) SetLogLevel(level log.LogLevel) {
	this.engine.SetLogLevel(level)
}

// SetDisableGlobalCache disable global cache or not
func (this *Model) SetDisableGlobalCache(disable bool) {
	this.engine.SetDisableGlobalCache(disable)
}

// DriverName return the current sql driver's name
func (this *Model) DriverName() string {
	return this.engine.DriverName()
}

// DataSourceName return the current connection string
func (this *Model) DataSourceName() string {
	return this.engine.DataSourceName()
}

// SetMapper set the name mapping rules
func (this *Model) SetMapper(mapper names.Mapper) {
	this.engine.SetMapper(mapper)
}

// SetTableMapper set the table name mapping rule
func (this *Model) SetTableMapper(mapper names.Mapper) {
	this.engine.SetTableMapper(mapper)
}

// SetColumnMapper set the column name mapping rule
func (this *Model) SetColumnMapper(mapper names.Mapper) {
	this.engine.SetColumnMapper(mapper)
}

// Quote Use QuoteStr quote the string sql
func (this *Model) Quote(value string) string {
	return this.engine.Quote(value)
}

// QuoteTo quotes string and writes into the buffer
func (this *Model) QuoteTo(buf *strings.Builder, value string) {
	this.engine.QuoteTo(buf, value)
}

// SQLType A simple wrapper to dialect's core.SqlType method
func (this *Model) SQLType(c *schemas.Column) string {
	return this.engine.SQLType(c)
}

// AutoIncrStr Database's autoincrement statement
func (this *Model) AutoIncrStr() string {
	return this.engine.AutoIncrStr()
}

// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
func (this *Model) SetConnMaxLifetime(d time.Duration) {
	this.engine.SetConnMaxLifetime(d)
}

// SetMaxOpenConns is only available for go 1.2+
func (this *Model) SetMaxOpenConns(conns int) {
	this.engine.SetMaxOpenConns(conns)
}

// SetMaxIdleConns set the max idle connections on pool, default is 2
func (this *Model) SetMaxIdleConns(conns int) {
	this.engine.SetMaxIdleConns(conns)
}

// SetDefaultCacher set the default cacher. Xorm's default not enable cacher.
func (this *Model) SetDefaultCacher(cacher caches.Cacher) {
	this.engine.SetDefaultCacher(cacher)
}

// GetDefaultCacher returns the default cacher
func (this *Model) GetDefaultCacher() caches.Cacher {
	return this.engine.GetDefaultCacher()
}

// NoCache If you has set default cacher, and you want temporilly stop use cache,
// you can use NoCache()
func (this *Model) NoCache() *xorm.Session {
	return this.engine.NoCache()
}

// NoCascade If you do not want to auto cascade load object
func (this *Model) NoCascade() *xorm.Session {
	return this.engine.NoCascade()
}

// MapCacher Set a table use a special cacher
func (this *Model) MapCacher(bean interface{}, cacher caches.Cacher) error {
	return this.engine.MapCacher(bean, cacher)
}

// NewDB provides an interface to operate database directly
func (this *Model) NewDB() (*core.DB, error) {
	return this.engine.NewDB()
}

// DB return the wrapper of sql.DB
func (this *Model) DB() *core.DB {
	return this.engine.DB()
}

// Dialect return database dialect
func (this *Model) Dialect() dialects.Dialect {
	return this.engine.Dialect()
}

// NewSession New a session
func (this *Model) NewSession() *xorm.Session {
	return this.engine.NewSession()
}

// Close the engine
func (this *Model) Close() error {
	return this.engine.Close()
}

// Ping tests if database is alive
func (this *Model) Ping() error {
	return this.engine.Ping()
}

// SQL method let's you manually write raw SQL and operate
// For example:
//
//         this.engine.SQL("select * from user").Find(&users)
//
// This    code will execute "select * from user" and set the records to users
func (this *Model) SQL(query interface{}, args ...interface{}) *xorm.Session {
	return this.engine.SQL(query, args...)
}

// NoAutoTime Default if your struct has "created" or "updated" filed tag, the fields
// will automatically be filled with current time when Insert or Update
// invoked. Call NoAutoTime if you dont' want to fill automatically.
func (this *Model) NoAutoTime() *xorm.Session {
	return this.engine.NoAutoTime()
}

// NoAutoCondition disable auto generate Where condition from bean or not
func (this *Model) NoAutoCondition(no ...bool) *xorm.Session {
	return this.engine.NoAutoCondition(no...)
}

// DBMetas Retrieve all tables, columns, indexes' informations from database.
func (this *Model) DBMetas() ([]*schemas.Table, error) {
	return this.engine.DBMetas()
}

// DumpAllToFile dump database all table structs and data to a file
func (this *Model) DumpAllToFile(fp string, tp ...schemas.DBType) error {
	return this.engine.DumpAllToFile(fp, tp...)
}

// DumpAll dump database all table structs and data to w
func (this *Model) DumpAll(w io.Writer, tp ...schemas.DBType) error {
	return this.engine.DumpAll(w, tp...)
}

// DumpTablesToFile dump specified tables to SQL file.
func (this *Model) DumpTablesToFile(tables []*schemas.Table, fp string, tp ...schemas.DBType) error {
	return this.engine.DumpTablesToFile(tables, fp, tp...)
}

// DumpTables dump specify tables to io.Writer
func (this *Model) DumpTables(tables []*schemas.Table, w io.Writer, tp ...schemas.DBType) error {
	return this.engine.DumpTables(tables, w, tp...)
}

// Cascade use cascade or not
func (this *Model) Cascade(trueOrFalse ...bool) *xorm.Session {
	return this.engine.Cascade(trueOrFalse...)
}

// Where method provide a condition query
func (this *Model) Where(query interface{}, args ...interface{}) *xorm.Session {
	return this.engine.Where(query, args...)
}

// ID method provoide a condition as (id) = ?
func (this *Model) ID(id interface{}) *xorm.Session {
	return this.engine.ID(id)
}

func (this *Model) Id(id interface{}) *xorm.Session {
	return this.engine.ID(id)
}

// Before apply before Processor, affected bean is passed to closure arg
func (this *Model) Before(closures func(interface{})) *xorm.Session {
	return this.engine.Before(closures)
}

// After apply after insert Processor, affected bean is passed to closure arg
func (this *Model) After(closures func(interface{})) *xorm.Session {
	return this.engine.After(closures)
}

// Charset set charset when create table, only support mysql now
func (this *Model) Charset(charset string) *xorm.Session {
	return this.engine.Charset(charset)
}

// StoreEngine set store engine when create table, only support mysql now
func (this *Model) StoreEngine(storeEngine string) *xorm.Session {
	return this.engine.StoreEngine(storeEngine)
}

// Distinct use for distinct columns. Caution: when you are using cache,
// distinct will not be cached because cache system need id,
// but distinct will not provide id
func (this *Model) Distinct(columns ...string) *xorm.Session {
	return this.engine.Distinct(columns...)
}

// Select customerize your select columns or contents
func (this *Model) Select(str string) *xorm.Session {
	return this.engine.Select(str)
}

// Cols only use the parameters as select or update columns
func (this *Model) Cols(columns ...string) *xorm.Session {
	return this.engine.Cols(columns...)
}

// AllCols indicates that all columns should be use
func (this *Model) AllCols() *xorm.Session {
	return this.engine.AllCols()
}

// MustCols specify some columns must use even if they are empty
func (this *Model) MustCols(columns ...string) *xorm.Session {
	return this.engine.MustCols(columns...)
}

// UseBool xorm automatically retrieve condition according struct, but
// if struct has bool field, it will ignore them. So use UseBool
// to tell system to do not ignore them.
// If no parameters, it will use all the bool field of struct, or
// it will use parameters's columns
func (this *Model) UseBool(columns ...string) *xorm.Session {
	return this.engine.UseBool(columns...)
}

// Omit only not use the parameters as select or update columns
func (this *Model) Omit(columns ...string) *xorm.Session {
	return this.engine.Omit(columns...)
}

// Nullable set null when column is zero-value and nullable for update
func (this *Model) Nullable(columns ...string) *xorm.Session {
	return this.engine.Nullable(columns...)
}

// In will generate "column IN (?, ?)"
func (this *Model) In(column string, args ...interface{}) *xorm.Session {
	return this.engine.In(column, args...)
}

// NotIn will generate "column NOT IN (?, ?)"
func (this *Model) NotIn(column string, args ...interface{}) *xorm.Session {
	return this.engine.NotIn(column, args...)
}

// Incr provides a update string like "column = column + ?"
func (this *Model) Incr(column string, args ...interface{}) *xorm.Session {
	return this.engine.Incr(column, args...)
}

// Decr provides a update string like "column = column - ?"
func (this *Model) Decr(column string, args ...interface{}) *xorm.Session {
	return this.engine.Decr(column, args...)
}

// SetExpr provides a update string like "column = {expression}"
func (this *Model) SetExpr(column string, expression interface{}) *xorm.Session {
	return this.engine.SetExpr(column, expression)
}

// Table temporarily change the Get, Find, Update's table
func (this *Model) Table(tableNameOrBean interface{}) *xorm.Session {
	return this.engine.Table(tableNameOrBean)
}

// Alias set the table alias
func (this *Model) Alias(alias string) *xorm.Session {
	return this.engine.Alias(alias)
}

// Limit will generate "LIMIT start, limit"
func (this *Model) Limit(limit int, start ...int) *xorm.Session {
	return this.engine.Limit(limit, start...)
}

// Desc will generate "ORDER BY column1 DESC, column2 DESC"
func (this *Model) Desc(colNames ...string) *xorm.Session {
	return this.engine.Desc(colNames...)
}

// Asc will generate "ORDER BY column1,column2 Asc"
// This method can chainable use.
//
//        this.engine.Desc("name").Asc("age").Find(&users)
//        // SELECT * FROM user ORDER BY name DESC, age ASC
//
func (this *Model) Asc(colNames ...string) *xorm.Session {
	return this.engine.Asc(colNames...)
}

// OrderBy will generate "ORDER BY order"
func (this *Model) OrderBy(order string) *xorm.Session {
	return this.engine.OrderBy(order)
}

// Prepare enables prepare statement
func (this *Model) Prepare() *xorm.Session {
	return this.engine.Prepare()
}

// Join the join_operator should be one of INNER, LEFT OUTER, CROSS etc - this will be prepended to JOIN
func (this *Model) Join(joinOperator string, tablename interface{}, condition string, args ...interface{}) *xorm.Session {
	return this.engine.Join(joinOperator, tablename, condition, args...)
}

// GroupBy generate group by statement
func (this *Model) GroupBy(keys string) *xorm.Session {
	return this.engine.GroupBy(keys)
}

// Having generate having statement
func (this *Model) Having(conditions string) *xorm.Session {
	return this.engine.Having(conditions)
}

// TableInfo get table info according to bean's content
func (this *Model) TableInfo(bean interface{}) (*schemas.Table, error) {
	return this.engine.TableInfo(bean)
}

// IsTableEmpty if a table has any reocrd
func (this *Model) IsTableEmpty(bean interface{}) (bool, error) {
	return this.engine.IsTableEmpty(bean)
}

// IsTableExist if a table is exist
func (this *Model) IsTableExist(beanOrTableName interface{}) (bool, error) {
	return this.engine.IsTableExist(beanOrTableName)
}

// IDOf get id from one struct
func (this *Model) IDOf(bean interface{}) (schemas.PK, error) {
	return this.engine.IDOf(bean)
}

// TableName returns table name with schema prefix if has
func (this *Model) TableName(bean interface{}, includeSchema ...bool) string {
	return this.engine.TableName(bean, includeSchema...)
}

// IDOfV get id from one value of struct
func (this *Model) IDOfV(rv reflect.Value) (schemas.PK, error) {
	return this.engine.IDOf(rv)
}

// CreateIndexes create indexes
func (this *Model) CreateIndexes(bean interface{}) error {
	return this.engine.CreateIndexes(bean)
}

// CreateUniques create uniques
func (this *Model) CreateUniques(bean interface{}) error {
	return this.engine.CreateUniques(bean)
}

// ClearCacheBean if enabled cache, clear the cache bean
func (this *Model) ClearCacheBean(bean interface{}, id string) error {
	return this.engine.ClearCacheBean(bean, id)
}

// ClearCache if enabled cache, clear some tables' cache
func (this *Model) ClearCache(beans ...interface{}) error {
	return this.engine.ClearCache(beans...)
}

// UnMapType remove table from tables cache
func (this *Model) UnMapType(t reflect.Type) {
	this.engine.UnMapType(t)
}

// Sync the new struct changes to database, this method will automatically add
// table, column, index, unique. but will not delete or change anything.
// If you change some field, you should change the database manually.
func (this *Model) Sync(beans ...interface{}) error {
	return this.engine.Sync(beans...)
}

// Sync2 synchronize structs to database tables
func (this *Model) Sync2(beans ...interface{}) error {
	return this.engine.Sync2(beans...)
}

// CreateTables create tabls according bean
func (this *Model) CreateTables(beans ...interface{}) error {
	return this.engine.CreateTables(beans...)
}

// DropTables drop specify tables
func (this *Model) DropTables(beans ...interface{}) error {
	return this.engine.DropTables(beans...)
}

// DropIndexes drop indexes of a table
func (this *Model) DropIndexes(bean interface{}) error {
	return this.engine.DropIndexes(bean)
}

// Exec raw sql
func (this *Model) Exec(sqlOrArgs ...interface{}) (sql.Result, error) {
	return this.engine.Exec(sqlOrArgs ...)
}

// Query a raw sql and return records as []map[string][]byte
func (this *Model) QueryBytes(sqlOrArgs ...interface{}) (resultsSlice []map[string][]byte, err error) {
	return this.engine.QueryBytes(sqlOrArgs...)
}

// Query a raw sql and return records as []map[string]Value
func (this *Model) QueryValue(sqlOrArgs ...interface{}) (resultsSlice []map[string]xorm.Value, err error) {
	return this.engine.QueryValue(sqlOrArgs...)
}

// Query a raw sql and return records as Result
func (this *Model) QueryResult(sqlOrArgs ...interface{}) (result *xorm.ResultValue) {
	return this.engine.QueryResult(sqlOrArgs...)
}

// QueryString runs a raw sql and return records as []map[string]string
func (this *Model) QueryString(sqlOrArgs ...interface{}) ([]map[string]string, error) {
	return this.engine.QueryString(sqlOrArgs...)
}

// QueryInterface runs a raw sql and return records as []map[string]interface{}
func (this *Model) QueryInterface(sqlOrArgs ...interface{}) ([]map[string]interface{}, error) {
	return this.engine.QueryInterface(sqlOrArgs...)
}

// Insert one or more records
func (this *Model) Insert(beans ...interface{}) (int64, error) {
	return this.engine.Insert(beans...)
}

// InsertOne insert only one record
func (this *Model) InsertOne() (int64, error) {
	return this.engine.InsertOne(this)
}

// Update records, bean's non-empty fields are updated contents,
// condiBean' non-empty filds are conditions
// CAUTION:
//        1.bool will defaultly be updated content nor conditions
//         You should call UseBool if you have bool to use.
//        2.float32 & float64 may be not inexact as conditions
func (this *Model) Update(condiBeans ...interface{}) (int64, error) {
	return this.engine.Update(this, condiBeans...)
}

// Delete records, bean's non-empty fields are conditions
func (this *Model) Delete() (int64, error) {
	return this.engine.Delete(this)
}

// Get retrieve one record from table, bean's non-empty fields
// are conditions
func (this *Model) Get() (bool, error) {
	return this.engine.Get(this)
}

// Exist returns true if the record exist otherwise return false
func (this *Model) Exist() (bool, error) {
	return this.engine.Exist(this)
}

// Find retrieve records from table, condiBeans's non-empty fields
// are conditions. beans could be []Struct, []*Struct, map[int64]Struct
// map[int64]*Struct
func (this *Model) Find(condiBeans ...interface{}) error {
	return this.engine.Find(this, condiBeans...)
}

// FindAndCount find the results and also return the counts
func (this *Model) FindAndCount(rowsSlicePtr interface{}, condiBean ...interface{}) (int64, error) {
	return this.engine.FindAndCount(rowsSlicePtr, condiBean...)
}

// Iterate record by record handle records from table, bean's non-empty fields
// are conditions.
func (this *Model) Iterate(bean interface{}, fun xorm.IterFunc) error {
	return this.engine.Iterate(bean, fun)
}

// Rows return sql.Rows compatible Rows obj, as a forward Iterator object for iterating record by record, bean's non-empty fields
// are conditions.
func (this *Model) Rows(bean interface{}) (*xorm.Rows, error) {
	return this.engine.Rows(bean)
}

// Count counts the records. bean's non-empty fields are conditions.
func (this *Model) Count(bean ...interface{}) (int64, error) {
	return this.engine.Count(bean...)
}

// Sum sum the records by some column. bean's non-empty fields are conditions.
func (this *Model) Sum(bean interface{}, colName string) (float64, error) {
	return this.engine.Sum(bean, colName)
}

// SumInt sum the records by some column. bean's non-empty fields are conditions.
func (this *Model) SumInt(bean interface{}, colName string) (int64, error) {
	return this.engine.SumInt(bean, colName)
}

// Sums sum the records by some columns. bean's non-empty fields are conditions.
func (this *Model) Sums(bean interface{}, colNames ...string) ([]float64, error) {
	return this.engine.Sums(bean, colNames...)
}

// SumsInt like Sums but return slice of int64 instead of float64.
func (this *Model) SumsInt(bean interface{}, colNames ...string) ([]int64, error) {
	return this.engine.SumsInt(bean, colNames...)
}

// ImportFile SQL DDL file
func (this *Model) ImportFile(ddlPath string) ([]sql.Result, error) {
	return this.engine.ImportFile(ddlPath)
}

// Import SQL DDL from io.Reader
func (this *Model) Import(r io.Reader) ([]sql.Result, error) {
	return this.engine.Import(r)
}

// GetColumnMapper returns the column name mapper
func (this *Model) GetColumnMapper() names.Mapper {
	return this.engine.GetColumnMapper()
}

// GetTableMapper returns the table name mapper
func (this *Model) GetTableMapper() names.Mapper {
	return this.engine.GetTableMapper()
}

// GetTZLocation returns time zone of the application
func (this *Model) GetTZLocation() *time.Location {
	return this.engine.GetTZLocation()
}

// SetTZLocation sets time zone of the application
func (this *Model) SetTZLocation(tz *time.Location) {
	this.engine.SetTZLocation(tz)
}

// GetTZDatabase returns time zone of the database
func (this *Model) GetTZDatabase() *time.Location {
	return this.engine.GetTZDatabase()
}

// SetTZDatabase sets time zone of the database
func (this *Model) SetTZDatabase(tz *time.Location) {
	this.engine.SetTZDatabase(tz)
}

// SetSchema sets the schema of database
func (this *Model) SetSchema(schema string) {
	this.engine.SetSchema(schema)
}

func (this *Model) AddHook(hook contexts.Hook) {
	this.engine.AddHook(hook)
}

// Unscoped always disable struct tag "deleted"
func (this *Model) Unscoped() *xorm.Session {
	return this.engine.Unscoped()
}

// ContextHook creates a session with the context
func (this *Model) Context(ctx context.Context) *xorm.Session {
	return this.engine.Context(ctx)
}

// SetDefaultContext set the default context
func (this *Model) SetDefaultContext(ctx context.Context) {
	this.engine.SetDefaultContext(ctx)
}

// PingContext tests if database is alive
func (this *Model) PingContext(ctx context.Context) error {
	return this.engine.PingContext(ctx)
}

// Transaction Execute sql wrapped in a transaction(abbr as tx), tx will automatic commit if no errors occurred
func (this *Model) Transaction(f func(*xorm.Session) (interface{}, error)) (interface{}, error) {
	return this.engine.Transaction(f)
}