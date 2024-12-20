import { Injectable, inject } from '@angular/core';
import { Order } from '../models/order';
import { Cart } from '../models/cart';
import { Observable, map } from 'rxjs';
import { HttpClient } from '@angular/common/http';
import { API_URL } from '../providers/api-url-provider';

@Injectable({
  providedIn: 'root'
})
export class OrderService {
  private httpClient = inject(HttpClient);
  private backendUrl = inject<string>(API_URL);

  private orderPath = 'api/orders';

  public getAll(): Observable<Order[]> {
    return this.httpClient.get<Order[]>(`${this.backendUrl}/${this.orderPath}`, { withCredentials: true })
  }

  public getMyOrders(): Observable<Order[]> {
    return this.httpClient.get<Order[]>(`${this.backendUrl}/${this.orderPath}/my`, { withCredentials: true })
  }

  public get(id: string): Observable<Order | undefined> {
    return this.httpClient.get<Order>(`${this.backendUrl}/${this.orderPath}/${id}`, { withCredentials: true })
  }

  public addFromCart(): Observable<string> {
    return this.httpClient.post<Order>(`${this.backendUrl}/${this.orderPath}/from-cart`, {}, { withCredentials: true })
      .pipe(map(order => order.id));
  }

  public add(carts: Cart[]): Observable<string> {
    return this.httpClient.post<Order>(`${this.backendUrl}/${this.orderPath}/from-cart`, {
        productIds: carts.map(c => c.product?.id)
      }, { withCredentials: true })
        .pipe(map(order => order.id));
  }
}
