import { Injectable } from "@angular/core";
import { HttpClient } from "@angular/common/http";
import { Observable } from "rxjs";
import { Price } from "./price";

@Injectable()
export class PriceService {
    constructor(private http: HttpClient) {}
    baseURL: string = "http://127.0.0.1:8080/";
    fetchPrice(): Observable<any> {
        return this.http.get(this.baseURL + "fuel_prices/")
    }
    addPrice(price: Price): Observable<any> {
        const headers = { 'content-type': 'application/json'}  
        const body = JSON.stringify(price);
        console.log("body", body)
        return this.http.post(this.baseURL + 'fuel_prices/', body, {'headers':headers})
    }
}