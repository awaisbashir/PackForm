package middleware

import (
	"database/sql"
	"encoding/json" // Package to encode and decode the json into struct and vice versa
	"fmt"
	"go-post/models" // Models package where User schema is defined
	"log"
	"net/http" // Used to access the request and response object of the api
	"os"       // Used to read the environtment variable

	"github.com/joho/godotenv" // Package used to read the .env file
	_ "github.com/lib/pq"      // Postgres golang driver
)

// Response format
type response struct {
	ID      int64  `json:"id, omitempty"`
	Message string `json:"message, omitempty"`
}

// Create connection with postgres db
func createConnection() *sql.DB {
	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Open the connection
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		panic(err)
	}

	// Check the connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	// Return the connection
	return db

}

// Get all orders
func GetOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var qStr string;

	if searchStr := r.URL.Query().Get("filter"); searchStr != "" {
		qStr += ` AND ( OD.order_name ILIKE '%` + searchStr + `%'`;
		qStr += ` OR (
						SELECT COUNT(*)
						FROM order_items OIT
						WHERE OIT.product ILIKE '%` + searchStr + `%'
					) > 0
				)`
	}

	if start := r.URL.Query().Get("start"); start != "" {
		end := r.URL.Query().Get("end")

		qStr += ` AND OD.created_at BETWEEN '` + start + `'`;
		qStr += ` AND '` + end + `'`;
	}

	if sort := r.URL.Query().Get("sort"); sort != "" {
		qStr += ` ORDER BY ` + sort;
	} else {
		qStr += ` ORDER BY OD.created_at`
	}

	if order := r.URL.Query().Get("sort_order"); order == "true" {
		qStr += ` DESC`
	} else {
		qStr += ` ASC`;
	}

	if size := r.URL.Query().Get("size"); size != "" {
		qStr += ` LIMIT ` + size;
	}

	if page := r.URL.Query().Get("page"); page != "" {
		qStr += ` OFFSET ` + page;
	}

	// Get all the orders in the db
	orders, err := getOrders(qStr)

	if err != nil {
		log.Fatalf("Unable to get all order. %v", err)
	}

	// Send all the orders as response
	json.NewEncoder(w).Encode(orders)
}

// ------------------------------------------ Handler Functions ------------------------//

// Get all orders
func getOrders(q string) ([]models.OrderDetail, error) {

	// Create the postgres db connection
	db := createConnection()

	// Close the db connection
	defer db.Close()

	var orders = make([]models.OrderDetail,0)

	// Create the select sql query
	sqlStatement := `
		SELECT OD.order_name AS OrderName,
			TO_CHAR(OD.created_at , 'Mon FMDDth, HH:MI am') AS CreatedAt,
			CC.company_name AS CustomerCompany, 
			CS.name AS CustomerName, 
			(
				SELECT SUM(OI.price_per_unit * OI.quantity)
				FROM order_items OI
				WHERE OI.order_id = OD.id
				GROUP BY OI.order_id
			) AS TotalAmount,
			(
				SELECT COALESCE( SUM(items.price) ,0)
				FROM
					(
						SELECT SUM( DL.delivered_quantity * OIT.price_per_unit ) AS price
						FROM deliveries DL
						INNER JOIN order_items OIT ON OIT.id = DL.order_item_id
						WHERE OIT.order_id = OD.id
						GROUP BY DL.order_item_id
					) AS items
			) AS DeliveredAmount, 
			COUNT(*) OVER() AS Total
		FROM orders OD
		INNER JOIN customers CS ON CS.user_id = OD.customer_id
		INNER JOIN customer_companies CC ON CC.company_id = CS.company_id
		WHERE 1 = 1
	` + q;

	fmt.Println(sqlStatement)

	// Execute the sql statement
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// Close the statement
	defer rows.Close()

	// Iterate over the rows
	for rows.Next() {
		var order models.OrderDetail

		// unmarshal the row object the order
		err = rows.Scan(&order.OrderName, &order.CreatedAt, &order.CustomerCompany, &order.CustomerName, &order.TotalAmount, &order.DeliveredAmount, &order.Total)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// Append the order in the orders slice
		orders = append(orders, order)
	}

	// Return empty order on error
	return orders, err
}
