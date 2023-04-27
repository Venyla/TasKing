package data

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/google/uuid"
	"os"
	dataModels "webservice/data/dataModels"
)

var db *pg.DB

func Init() error {
	opt, err := pg.ParseURL(os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	db = pg.Connect(opt)

	err = createSchema()
	if err != nil {
		panic(err)
	}

	return nil
}

func GetDbConnection() *pg.Conn {
	return db.Conn()
}

func createSchema() error {
	fmt.Print("Creating DB...")
	db := GetDbConnection()
	defer db.Close()

	transaction, transactionError := db.Begin()
	if transactionError != nil {
		panic(transactionError)
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Something went wrong during creation: %s\n", r)
			transaction.Rollback()
			panic(r)
		}
	}()

	models := []interface{}{
		(*dataModels.Task)(nil),
		(*dataModels.TaskHistory)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			fmt.Println("Error occurred! Rolling back...")
			transaction.Rollback()
			return err
		}
	}

	db.Model(&dataModels.Task{
		TaskId:       uuid.New(),
		Title:        "Drink Beer",
		IconUrl:      "https://imageresizer.static9.net.au/mAbtmTO6BX05IdEILplNAgXv_Wc=/1200x675/https://prod.static9.net.au/fs/ff4238d6-65e7-4f73-afa2-537d3f64378e",
		XCoordinates: 20,
		YCoordinates: 35,
	}).Insert()

	db.Model(&dataModels.Task{
		TaskId:       uuid.New(),
		Title:        "Pet Elvis",
		IconUrl:      "https://scontent.fzrh3-1.fna.fbcdn.net/v/t39.30808-6/309430860_474589338041453_5993993981776996852_n.jpg?_nc_cat=108&ccb=1-7&_nc_sid=09cbfe&_nc_ohc=7Tkot0fTAsoAX-FYOaX&_nc_ht=scontent.fzrh3-1.fna&oh=00_AfAsy-DH7-AJZmeKHnIMzUDN24XHC4BrAXHZ-SXf_ys_Kg&oe=643AC639",
		XCoordinates: 43,
		YCoordinates: 51,
	}).Insert()

	db.Model(&dataModels.Task{
		TaskId: uuid.New(),
		Title:  "Visit Erika", IconUrl: "https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Ftse3.mm.bing.net%2Fth%3Fid%3DOIP.9YTVRgjKM5V4xxRDVrMBcwHaJ4%26pid%3DApi&f=1&ipt=d917c56f3d343346bbafd198f4fb180c0dcecfd6ab290b62d2a57902489c2fcb&ipo=images",
		XCoordinates: 40,
		YCoordinates: 45,
	}).Insert()

	db.Model(&dataModels.Task{
		TaskId:       uuid.New(),
		Title:        "Get Coffee",
		IconUrl:      "https://external-content.duckduckgo.com/iu/?u=http%3A%2F%2Fpngimg.com%2Fuploads%2Fmug_coffee%2Fmug_coffee_PNG16824.png&f=1&nofb=1&ipt=b69c4e760e4328689109ac952ad7db0a8a02b86678be54356fea6609aa07b699&ipo=images",
		XCoordinates: 50, YCoordinates: 50,
	}).Insert()

	transaction.Commit()
	fmt.Println("DB created!")

	// TODO maybe do a health check

	return nil
}
