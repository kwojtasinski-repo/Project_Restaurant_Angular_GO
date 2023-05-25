import { Component, OnInit } from '@angular/core';
import { Product } from 'src/app/models/product';
import { ProductService } from 'src/app/services/product.service';
import { take } from 'rxjs';
import { AuthStateService } from 'src/app/services/auth-state.service';
import { Store } from "@ngrx/store";
import { CartState } from 'src/app/stores/cart/cart.state';
import { addProductToCart } from 'src/app/stores/cart/cart.actions';

@Component({
  selector: 'app-menu',
  templateUrl: './menu.component.html',
  styleUrls: ['./menu.component.scss']
})
export class MenuComponent implements OnInit {
  public user$ = this.authService.getUser();
  public products: Product[] = [];
  public productsToShow: Product[] = [];
  public term: string = '';

  constructor(private productService: ProductService, private authService: AuthStateService, private cartStore: Store<CartState>) { }
  
  public ngOnInit(): void {
    this.productService.getAll()
      .pipe(take(1))
      .subscribe(p => {
        this.products = p;
        this.productsToShow = p;
      });
  }

  public search(term: string): void {
    this.productsToShow = this.products.filter(p => p.name.toLocaleLowerCase().startsWith(term.toLocaleLowerCase()));
  }

  public addToCart(product: Product): void {
    this.cartStore.dispatch(addProductToCart({ product }));
  }
}
