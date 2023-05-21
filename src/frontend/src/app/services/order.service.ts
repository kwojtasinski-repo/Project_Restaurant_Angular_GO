import { Injectable } from '@angular/core';
import { Order } from '../models/order';
import { Cart } from '../models/cart';
import { Observable } from 'rxjs';
import { Product } from '../models/product';
import { OrderProduct } from '../models/orderProduct';
import { of } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class OrderService {
  private orders: Order[] = [];

  constructor() { }

  public getAll(): Observable<Order[]> {
    return of(this.orders);
  }

  public add(cart: Cart): Observable<void> {
    this.orders.push({
      id: this.orders.length > 0 ? this.orders[this.orders.length - 1].id + 1 : 1,
      created: new Date(),
      orderNumber: new Date().toISOString(),
      price: cart.products.reduce((total, product) => total + product.price, 0),
      modified: undefined,
      orderProducts: this.addOrderProducts(cart.products),
      userId: 0
    });
    return new Observable((ob) => { ob.next(); ob.complete(); });
  }

  private addOrderProducts(products: Product[]): OrderProduct[] {
    let id = this.getLastIdFromOrderProducts();
    const orderProducts: OrderProduct[] = [];
    
    for (let product of products) {
      id++;
      orderProducts.push({
        id,
        name: product.name,
        price: product.price,
        productId: product.id
      });
    }

    return orderProducts;
  }

  private getLastIdFromOrderProducts(): number {
    const id = 1;
    if (this.orders.length === 0) {
      return id;
    }

    for (let i = this.orders.length - 1; i >= 0; i--) {
      if (this.orders[i].orderProducts.length === 0) {
        continue;
      }

      return this.orders[i].orderProducts[this.orders[i].orderProducts.length - 1].id
    }

    return id;
  }
}
