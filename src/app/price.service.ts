import { Injectable } from "@angular/core";
import { HttpClient, HttpParams } from "@angular/common/http";
import { Observable } from "rxjs";
import { Price } from "./price";

@Injectable()
export class PriceService {
    constructor(private http: HttpClient) {}
    baseURL: string = "http://127.0.0.1:8080/";
    fetchPrices(): Observable<any> {
        return this.http.get(this.baseURL + "fuel_prices/")
    }
    showPrice(price_id: number): Observable<any> {
        return this.http.get(`${this.baseURL}fuel_prices/${price_id}/`);
    }
    addPrice(price: Price): Observable<any> {
        const headers = { 'content-type': 'application/json'}  
        const body = JSON.stringify(price);
        console.log("body", body)
        return this.http.post(this.baseURL + 'fuel_prices/', body, {'headers':headers})
    }
    updatePrice(price: Price): Observable<any> {
        const headers = { 'content-type': 'application/json'}  
        const body = JSON.stringify(price);
        console.log("body", body)
        return this.http.put(`${this.baseURL}fuel_prices/${price.id}/`, body, {'headers':headers})
    }
    deletePrice(price_id: number) {
        return this.http.delete(`${this.baseURL}fuel_prices/${price_id}/`);
    }
}