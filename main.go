package main

import (
	"database/sql"
	"fmt"
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
	id                 int    `json:"id"`
	productID          string `json:"product_id"`
	productCost        int    `json:"product_cost"`
	productAvailabilty int    `json:"product_availabilty"`
	productSubcat      string `json:"product_subcat"`
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
	rows, err := db.Query("SELECT id,product_id as productID,product_cost as productCost,product_availabilty as productAvailabilty,product_subcat as productSubcat FROM inventory;")
	for rows.Next() {
		err = rows.Scan(&invt.id, &invt.productID, &invt.productCost, &invt.productAvailabilty, &invt.productSubcat)
		checkErr(err)
		inventory = append(inventory, invt)
	}

	checkErr(err)
	defer rows.Close()
	c.JSON(200, gin.H{
		"result": fmt.Sprintf("%+v", inventory),
		"count":  len(inventory),
	})
}
