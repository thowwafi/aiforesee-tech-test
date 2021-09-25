import { Component, OnInit } from "@angular/core";
import { PriceService } from "./price.service";
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
    error!: string;
    update = false
    constructor(private priceService: PriceService) {}
    ngOnInit(): void {
        this.refreshPrice();
    }

    refreshPrice() {
        this.priceService.fetchPrices().subscribe((res) => {
            console.log(res);
            this.resp = res;
            this.all_prices = this.resp.data;
        });
    }

    addPrice() {
        console.log("this.price", this.price);
        if (this.update) {
            this.updatePrice(this.price)
            return
        }
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
                this.error = error.message;
            }
        );
    }

    showPrice(price_id: number) {
        this.update = true
        this.priceService.showPrice(price_id)
            .subscribe((res) => {
                console.log(res);
                this.resp = res;
                this.price = this.resp.data;
                this.toggleModal(this.price.id)
            }
        );
    }

    updatePrice(price_id: Price) {
        console.log("update", price_id)
        this.priceService.updatePrice(this.price).subscribe(
            (data) => {
                console.log(data);
                if (data.type == "success") {
                  this.refreshPrice();
                  this.toggleModal();
                } else {
                  console.log("oops1", data.message);
                  this.error = data.message;
                }
            },
            (error) => {
                console.log("oops2", error);
                this.error = error.message;
            }
        );
    }

    deletePrice(price_id: number) {
        console.log(price_id);
        this.priceService.deletePrice(price_id)
            .subscribe(
                result => {
                    console.log(result)
                    this.refreshPrice();
                },
                err => console.error(err)
            );
    }

    showModal = false;
    toggleModal(price_id: number = 0) {
        if (!price_id) {
            this.price = new Price();
            this.update = false
        }
        this.error = ""
        this.showModal = !this.showModal;
    }
}
