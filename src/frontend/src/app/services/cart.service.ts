import { Injectable } from '@angular/core';
import { Cart } from '../models/cart';
import { Observable, of } from 'rxjs';
import { Product } from '../models/product';

@Injectable({
  providedIn: 'root'
})
export class CartService {
  private cart: Cart = {
    products: []
  };

  constructor() { }

  public add(product: Product): Observable<void> {
    this.cart = { products: [...this.cart.products, product] };
    return new Observable((ob) => { ob.next(); ob.complete(); });
  }

  public delete(product: Product): Observable<void> {
    this.cart = { products: this.cart.products.filter(p => p.id !== product.id) };
    return new Observable((ob) => { ob.next(); ob.complete(); });
  }

  public getCart(): Observable<Cart> {
    return of(this.cart);
  }

  public finalizeCart(): void {
    this.cart = {
      products: []
    };
  }
}
