package main

import (
	"database/sql"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func index(c *gin.Context) {
	hostname, err := os.Hostname()
	checkErr(err)
	c.String(200, "v3 "+hostname)
}

func healthz(c *gin.Context) {
	c.String(200, "OK")
}

type InventoryItem struct {
	ID                 int    `json:"id"`
	ProductID          string `json:"product_id"`
	ProductCost        int    `json:"product_cost"`
	ProductAvailabilty int    `json:"product_availabilty"`
	ProductSubcat      string `json:"product_subcat"`
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

/*******************  MAIN Function **************/
func main() {
	app := gin.Default()
	app.GET("/", index)
	app.GET("/healthz", healthz)
	app.GET("/inventory", fetch)
	app.Run(":8000")
}

/******************* End MAIN Function **************/

func fetch(c *gin.Context) {
	var (
		invt      InventoryItem
		inventory []InventoryItem
	)
	connStr := os.Getenv("sql_user") + ":" + os.Getenv("sql_password") + "@tcp(" + os.Getenv("sql_host") + ":3306)/" + os.Getenv("sql_db")
	db, err := sql.Open("mysql", connStr)
	checkErr(err)
	defer db.Close()
	rows, err := db.Query("SELECT * FROM inventory;")
	for rows.Next() {
		err = rows.Scan(&invt.ID, &invt.ProductID, &invt.ProductCost, &invt.ProductAvailabilty, &invt.ProductSubcat)
		checkErr(err)
		inventory = append(inventory, invt)
	}

	checkErr(err)
	defer rows.Close()
	c.JSON(200, gin.H{
		"result": inventory,
		"count":  len(inventory),
	})
}
