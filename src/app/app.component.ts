import { Component, OnInit } from "@angular/core";
import { Observable } from "rxjs";
import { PriceService } from "./price.service";
import { HttpClient } from "@angular/common/http";
import { Price } from "./price";

@Component({
    selector: "app-root",
    templateUrl: "./app.component.html",
    styleUrls: ["./app.component.css"],
})
export class AppComponent {
    items = [{ "1": "2" }];
    all_prices!: Price[];
    price = new Price();
    resp: any;
    error: any;
    constructor(private priceService: PriceService) {}
    ngOnInit(): void {
        this.refreshPrice();
    }

    refreshPrice() {
        this.priceService.fetchPrice().subscribe((res) => {
            console.log(res);
            this.resp = res;
            this.all_prices = this.resp.data;
        });
    }

    addPrice() {
        console.log("this.price", this.price);
        this.priceService.addPrice(this.price).subscribe(
            (data) => {
                console.log(data);
                if (data.type == "success") {
                  this.refreshPrice();
                  this.toggleModal();
                } else {
                  console.log("oops", data.message);
                  this.error = data.message;
                }
            },
            (error) => {
                console.log("oops", error);
                this.error = error;
            }
        );
    }

    showModal = false;
    toggleModal() {
        this.showModal = !this.showModal;
    }
}
