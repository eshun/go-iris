package models

import (
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"errors"
	"strconv"
	"github.com/kataras/iris"
	"github.com/go-xorm/core"
)

var Orm *xorm.Engine

//数据库连接
func InitDB(app *iris.Application) {
	var dataSourceName string
	var err error
	driverName := AppConfig.DriverName
	if driverName == "sqlite3" {

	} else if driverName == "mysql" {
		dataSourceName = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", AppConfig.DBUser, AppConfig.DBPassword, AppConfig.DBServer, strconv.Itoa(AppConfig.DBPort), AppConfig.DBName, AppConfig.DBCharset)
	}
	Orm, err = xorm.NewEngine(AppConfig.DriverName, dataSourceName)
	if err != nil {
		app.Logger().Fatalf("orm failed to initialized: %v", err)
		return
	}
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "tb_")
	Orm.SetTableMapper(tbMapper)

	Orm.ShowSQL(AppConfig.ShowSql)
	Orm.Logger().SetLevel(core.LOG_INFO)
	//Orm.SetMaxIdleConns(setting.MaxIdle)
	//Orm.SetMaxOpenConns(setting.MaxOpen)

	Syncdb(app)
}

func Syncdb(app *iris.Application) {
	err := Orm.Sync2(new(Address), new(Ad), new(AdPosition), new(Admin), new(Attribute), new(AttributeCategory), new(Brand),
		new(Cart), new(Category), new(Channel), new(Collect), new(Comment), new(CommentPicture), new(Coupon), new(Feedback),
		new(Footprint), new(Goods), new(GoodsAttribute), new(GoodsGallery), new(GoodsIssue), new(GoodsSpecification),
		new(Keywords), new(Order), new(OrderGoods), new(Product), new(Region), new(RelatedGoods), new(SearchHistory),
		new(Specification), new(Topic), new(TopicCategory), new(User), new(UserCoupon), new(UserLevel))
	if err != nil {
		app.Logger().Fatalf("orm failed to initialized table: %v", err)
		return
	}

	//Orm.CreateTables()
	//Orm.IsTableExist();DropTables(),IsTableEmpty()
}

var (
	ErrNotExist = errors.New("not exist")
)