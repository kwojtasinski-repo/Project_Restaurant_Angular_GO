import { Component, OnInit } from '@angular/core';
import { NgxSpinnerService } from 'ngx-spinner';
import { take } from 'rxjs';
import { Cart } from 'src/app/models/cart';
import { CartService } from 'src/app/services/cart.service';

@Component({
  selector: 'app-carts',
  templateUrl: './carts.component.html',
  styleUrls: ['./carts.component.scss']
})
export class CartsComponent implements OnInit {
  public cart: Cart = { products: [] };
  public isLoading: boolean = true;

  constructor(private cartService: CartService, private spinnerService: NgxSpinnerService) { }

  public ngOnInit(): void {
    this.spinnerService.show();
    this.cartService.getCart()
      .pipe(take(1))
      .subscribe(c => {
        setTimeout(() => {
        this.cart = c;
        this.isLoading = false;
        this.spinnerService.hide();});
      });
  }
}
