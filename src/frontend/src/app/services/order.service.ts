import { Injectable, OnDestroy } from '@angular/core';
import { Order } from '../models/order';
import { Cart } from '../models/cart';
import { Observable, Subscription } from 'rxjs';
import { Product } from '../models/product';
import { OrderProduct } from '../models/orderProduct';
import { of } from 'rxjs';
import { Store } from '@ngrx/store';
import { LoginState } from '../stores/login/login.state';
import { getUser } from '../stores/login/login.selectors';
import { User } from '../models/user';

@Injectable({
  providedIn: 'root'
})
export class OrderService implements OnDestroy {
  private orders: Order[] = [];
  private user: User | null = null;
  private getUserSubscription$ = new Subscription();

  constructor(private store: Store<LoginState>) {
    this.getUserSubscription$ = this.store.select(getUser)
      .subscribe(u => {
        this.user = u
      });
  }

  public ngOnDestroy(): void {
    this.getUserSubscription$.unsubscribe();
  }

  public getAll(): Observable<Order[]> {
    return of(this.orders);
  }

  public getMyOrders() {
    return of(this.orders.filter(o => o.userId === this.user?.id ?? 0));
  }

  public get(id: number): Observable<Order | undefined> {
    return of(this.orders.find(o => o.id === id));
  }

  public add(carts: Cart[]): Observable<number> {
    const id = this.orders.length > 0 ? this.orders[this.orders.length - 1].id + 1 : 1;
    this.orders.push({
      id: id,
      created: new Date(),
      orderNumber: new Date().toISOString(),
      price: carts.reduce((total, cart) => total + (cart.product?.price ?? 0), 0),
      modified: undefined,
      orderProducts: this.addOrderProducts(carts),
      userId: this.user?.id ?? 0
    });
    
    return new Observable((ob) => { ob.next(id); ob.complete(); });
  }

  private addOrderProducts(carts: Cart[]): OrderProduct[] {
    let id = this.getLastIdFromOrderProducts();
    const orderProducts: OrderProduct[] = [];
    
    for (let cart of carts) {
      id++;
      orderProducts.push({
        id,
        name: cart.product?.name ?? '',
        price: cart.product?.price ?? 0,
        productId: cart.product?.id ?? 0
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
