# Fuelprices
```
CREATE TABLE fuel_prices (
    id SERIAL PRIMARY KEY,
    qty INT NOT NULL,
    premium_price INT NOT NULL,
    pertalite_price INT NOT NULL
);
```
This project was generated with [Angular CLI](https://github.com/angular/angular-cli) version 12.2.7.

## Development steps
1. Clone this repository
2. Update `.env` file with your configuration
    ```
    DB_USER = "your_user_name"
    DB_PASSWORD = "your_password"
    DB_NAME = "your_db_name"
    ```
3. Run `generate_data.go` file
    ```
    cd server/database/
    go run generate_initial_data.go
    ```
4. Run GO server
    ```
    cd ../
    go run server.go
    ```
5. Run `ng serve` for a dev server. Navigate to `http://localhost:4200/`. The app will automatically reload if you change any of the source files.
    ```
    cd ..
    ng serve
    ```
