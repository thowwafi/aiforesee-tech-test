import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppComponent } from './app.component';
import { ItemComponent } from './item/item.component';
import { ModalModule } from 'ng-modal-lib';
import { HttpClientModule } from '@angular/common/http';
import { PriceService } from './price.service';
import { FormsModule } from '@angular/forms'; 

@NgModule({
  declarations: [
    AppComponent,
    ItemComponent
  ],
  imports: [
    BrowserModule,
    ModalModule,
    HttpClientModule,
    FormsModule
  ],
  providers: [PriceService ],
  bootstrap: [AppComponent]
})
export class AppModule { }
