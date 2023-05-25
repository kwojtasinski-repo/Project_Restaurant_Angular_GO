import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { take } from 'rxjs';
import { Product } from 'src/app/models/product';
import { ProductService } from 'src/app/services/product.service';
import { NgxSpinnerService } from 'ngx-spinner';
import { AuthStateService } from 'src/app/services/auth-state.service';
import { Store } from "@ngrx/store";
import { CartState } from 'src/app/stores/cart/cart.state';
import { addProductToCart } from 'src/app/stores/cart/cart.actions';

@Component({
  selector: 'app-view-products',
  templateUrl: './view-products.component.html',
  styleUrls: ['./view-products.component.scss']
})
export class ViewProductsComponent implements OnInit {
  public product: Product | undefined;
  public isLoading = true;
  public user$ = this.authService.getUser();

  constructor(private productService: ProductService, private route: ActivatedRoute, private spinnerService: NgxSpinnerService,
    private authService: AuthStateService, private cartStore: Store<CartState>) { }

  public ngOnInit(): void {
    this.spinnerService.show();
    const id = this.route.snapshot.paramMap.get('id') ? new Number(this.route.snapshot.paramMap.get('id')).valueOf() : 0;
    this.productService.get(id)
      .pipe(take(1))
      .subscribe(p => {
        this.product = p
        this.isLoading = false;
        this.spinnerService.hide();
      });
  }

  public addToCart(product: Product): void {
    this.cartStore.dispatch(addProductToCart({ product }));
  }
}
