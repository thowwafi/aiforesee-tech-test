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
    items = [{"1": "2"}];
    all_prices!: Price[];
    resp:any;
    constructor(private priceService: PriceService) {}
    ngOnInit(): void {
      this.priceService.fetchPrice().subscribe(
        (res) => {
          console.log(res)
          this.resp = res;
          this.all_prices = this.resp.data;
        },
        (err) => console.log(err),
        () => console.log('done!')
      );
    }
    // title = 'todo';

    // filter: 'all' | 'active' | 'done' = 'all';

    // allItems = [
    //   { description: 'eat', done: true },
    //   { description: 'sleep', done: false },
    //   { description: 'play', done: false },
    //   { description: 'laugh', done: false },
    // ];

    // get items() {
    //   if (this.filter === 'all') {
    //     return this.allItems;
    //   }
    //   return this.allItems.filter((item) =>
    //     this.filter === 'done' ? item.done : !item.done
    //   );
    // }
    addItem(qty: number, premium_price: number, pertalite_price: number) {
      console.log(qty)
      console.log(premium_price)
      console.log(pertalite_price)
    }
    // remove(item: any) {
    //   this.allItems.splice(this.allItems.indexOf(item), 1);
    // }
}
