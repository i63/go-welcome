package main

import (
    "github.com/gin-gonic/gin"
    "os"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func index (c *gin.Context){
    hostname,err := os.Hostname()
    checkErr(err)
    c.String(200,"v3 "+ hostname)
}

func healthz (c *gin.Context){
    c.String(200,"OK")
}



type Inventory struct {
    id int 
    product_id string
    product_cost int
    product_availabilty int
    product_subcat string
}


func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

/*******************  MAIN Function **************/
func main(){
  app := gin.Default()
  app.GET("/", index)
  app.GET("/healthz", healthz)
  app.GET("/dbtest",fetch)
  app.Run(":8000")
}
/******************* End MAIN Function **************/




func fetch (c *gin.Context){
    connStr := os.Getenv("sql_user")+":"+os.Getenv("sql_password")+"@tcp("+os.Getenv("sql_host")+":3306)/"+os.Getenv("sql_db")
    db, err := sql.Open("mysql",connStr)
    checkErr(err)
    defer db.Close()
    invt := new(Inventory)
    db.QueryRow("SELECT * FROM inventory where product_subcat=1").Scan(&invt.id,&invt.product_id,&invt.product_cost,&invt.product_availabilty,&invt.product_subcat)
    checkErr(err)
    c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}

