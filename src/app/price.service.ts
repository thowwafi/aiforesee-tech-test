import { Injectable } from "@angular/core";
import { HttpClient } from "@angular/common/http";
import { Observable } from "rxjs";

@Injectable()
export class PriceService {
    constructor(private http: HttpClient) {}
    fetchPrice(): Observable<any> {
        return this.http.get("http://127.0.0.1:8080/fuel_prices/")
    }
}