package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
    DB_USER     = "alfi"
    DB_PASSWORD = "1234567890"
    DB_NAME     = "aiforesee_application"
)

// DB set up
func setupDB() *sql.DB {
    dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	// dbURL := "postgres://iwggzgabufxdyc:b81c780e6921121c9d37cf23ef612ba2d276b81b7ae15de4019df264ab50dfa0@ec2-107-22-245-82.compute-1.amazonaws.com:5432/d7rc2qfknu4r74"
    db, err := sql.Open("postgres", dbinfo)

    checkErr(err)

    return db
}


type FuelPrice struct {
    Id int `json:"id"`
    Qty int `json:"qty"`
    PremiumPrice int `json:"premium_price"`
    PertalitePrice int `json:"pertalite_price"`
}

type JsonResponse struct {
    Type    string `json:"type"`
    Data    []FuelPrice `json:"data"`
    Message string `json:"message"`
}
type JsonResponseRetrieve struct {
    Type    string `json:"type"`
    Data    FuelPrice `json:"data"`
    Message string `json:"message"`
}

// Go main function
func main() {
    // mux := http.NewServeMux()
    // mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    //     w.Header().Set("Content-Type", "application/json")
    //     w.Write([]byte("{\"hello\": \"world\"}"))
    // })

    // // cors.Default() setup the middleware with default options being
    // // all origins accepted with simple methods (GET, POST). See
    // // documentation below for more options.
    // handler := cors.Default().Handler(mux)
    // http.ListenAndServe(":8080", handler)

    // Init the mux router
    router := mux.NewRouter()

    // Get all fuel prices
    router.HandleFunc("/fuel_prices/", GetFuelPrices).Methods("GET", "OPTIONS")

    // Create a FuelPrice
    router.HandleFunc("/fuel_prices/", CreateFuelPrice).Methods("POST", "OPTIONS")

	// Read a fuel price by id
    router.HandleFunc("/fuel_prices/{fuelpriceid}/", ReadFuelPrice).Methods("GET", "OPTIONS")

	// Update a fuel price by id
    router.HandleFunc("/fuel_prices/{fuelpriceid}/", UpdateFuelPrice).Methods("PUT", "OPTIONS")

    // Delete a specific FuelPrice by the FuelPriceID
    router.HandleFunc("/fuel_prices/{fuelpriceid}/", DeleteFuelPrice).Methods("DELETE", "OPTIONS")

    // serve the app
    fmt.Println("Server at 8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}

func printMessage(message string) {
    fmt.Println("")
    fmt.Println(message)
    fmt.Println("")
}

// Function for handling errors
func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

func GetFuelPrices(w http.ResponseWriter, r *http.Request) {
    db := setupDB()
	w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "*")
    w.Header().Set("Access-Control-Allow-Headers", "*")

    if r.Method == "OPTIONS" {
        w.Write([]byte("allowed"))
        return
    }

    printMessage("Getting fuel prices...")

    rows, err := db.Query("SELECT * FROM fuel_prices")

    // check errors
    checkErr(err)

    // var response []JsonResponse
    var fuel_prices []FuelPrice

    for rows.Next() {
        var id int
        var qty int
        var premium_price int
        var pertalite_price int

        err = rows.Scan(&id, &qty, &premium_price, &pertalite_price)

        // check errors
        checkErr(err)

        fuel_prices = append(fuel_prices, FuelPrice{Id: id, Qty: qty, PremiumPrice: premium_price, PertalitePrice: pertalite_price})
    }

    var response = JsonResponse{Type: "success", Data: fuel_prices}

    json.NewEncoder(w).Encode(response)
}

func CreateFuelPrice(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "*")
    w.Header().Set("Access-Control-Allow-Headers", "*")

    decoder := json.NewDecoder(r.Body)
    var price FuelPrice
    err := decoder.Decode(&price)
    if err != nil {
        panic(err)
    }
    log.Println(price)
	qty := price.Qty
    premium_price := price.PremiumPrice
    pertalite_price := price.PertalitePrice

    log.Println(qty)
    log.Println(premium_price)
    log.Println(pertalite_price)

    var response = JsonResponseRetrieve{}

    if qty == 0 || premium_price == 0 || pertalite_price == 0 {
        response = JsonResponseRetrieve{Type: "error", Message: "You are missing Qty or premium_price or pertalite_price parameter."}
    } else {
        db := setupDB()

        printMessage("Inserting fuel_price into DB")

        var lastInsertID int
		err := db.QueryRow("INSERT INTO fuel_prices(qty, premium_price, pertalite_price) VALUES($1, $2, $3) returning id;", qty, premium_price, pertalite_price).Scan(&lastInsertID)
		checkErr(err)
		row := db.QueryRow("SELECT * FROM fuel_prices WHERE id=$1;", lastInsertID)
		err_ := row.Scan(&lastInsertID, &qty, &premium_price,&pertalite_price)
		checkErr(err_)

		fuelprice := FuelPrice{Id: lastInsertID, Qty: qty, PremiumPrice: premium_price, PertalitePrice: pertalite_price}

		response = JsonResponseRetrieve{Type: "success", Message: "The full price has been inserted successfully!", Data: fuelprice}
    }

    json.NewEncoder(w).Encode(response)
}


func ReadFuelPrice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)

    fuelpriceid := params["fuelpriceid"]

    var response = JsonResponseRetrieve{}

    if fuelpriceid == "" {
        response = JsonResponseRetrieve{Type: "error", Message: "You are missing fuelpriceid parameter."}
    } else {
        db := setupDB()

        printMessage("Read a fuel price from DB")
		var id int
		var qty int
		var premium_price int
		var pertalite_price int

        row := db.QueryRow("SELECT * FROM fuel_prices WHERE id=$1;", fuelpriceid)
		err := row.Scan(&id, &qty, &premium_price,&pertalite_price)

        checkErr(err)

		fuelprice := FuelPrice{Id: id, Qty: qty, PremiumPrice: premium_price, PertalitePrice: pertalite_price}

		response = JsonResponseRetrieve{Type: "success", Data: fuelprice}
    }

    json.NewEncoder(w).Encode(response)
}


func UpdateFuelPrice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)

    fuelpriceid := params["fuelpriceid"]
	qty := r.FormValue("qty")
    premium_price := r.FormValue("premium_price")
    pertalite_price := r.FormValue("pertalite_price")

    var response = JsonResponseRetrieve{}

    if fuelpriceid == "" {
        response = JsonResponseRetrieve{Type: "error", Message: "You are missing fuelpriceid parameter."}
    } else {
        db := setupDB()

        printMessage("Read a fuel price from DB")

		sqlStatement := `
			UPDATE fuel_prices
			SET qty = $2, premium_price = $3, pertalite_price=$4
			WHERE id = $1;`
		_, err := db.Exec(sqlStatement, fuelpriceid, qty, premium_price, pertalite_price)

		checkErr(err)

		fuelpriceid_int, _ := strconv.Atoi(fuelpriceid)
		qty_int, _ := strconv.Atoi(qty)
		premium_price_int, _ := strconv.Atoi(premium_price)
		pertalite_price_int, _ := strconv.Atoi(pertalite_price)

		fuelprice := FuelPrice{Id: fuelpriceid_int, Qty: qty_int, PremiumPrice: premium_price_int, PertalitePrice: pertalite_price_int}

        checkErr(err)

		response = JsonResponseRetrieve{Type: "success", Message: "Successfully updated fuel price.", Data: fuelprice}
    }

    json.NewEncoder(w).Encode(response)
}


func DeleteFuelPrice(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)

    fuelpriceid := params["fuelpriceid"]

    var response = JsonResponse{}

    if fuelpriceid == "" {
        response = JsonResponse{Type: "error", Message: "You are missing fuelpriceid parameter."}
    } else {
        db := setupDB()

        printMessage("Deleting fuel price from DB")

        _, err := db.Exec("DELETE FROM fuel_prices where id = $1", fuelpriceid)

        // check errors
        checkErr(err)

        response = JsonResponse{Type: "success", Message: "The fuel price has been deleted successfully!"}
    }

    json.NewEncoder(w).Encode(response)
}