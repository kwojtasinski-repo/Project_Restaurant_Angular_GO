import { Inject, Injectable } from '@angular/core';
import { Cart } from '../models/cart';
import { Observable } from 'rxjs';
import { Product } from '../models/product';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class CartService {
  private cartPath = 'api/carts';

  constructor(private httpClient: HttpClient, @Inject('API_URL') private backendUrl: string) { }

  public add(product: Product): Observable<void> {
    return this.httpClient.post<void>(`${this.backendUrl}/${this.cartPath}`, { productId: product.id }, { withCredentials: true });
  }

  public delete(id: number): Observable<void> {
    return this.httpClient.delete<void>(`${this.backendUrl}/${this.cartPath}/${id}`, { withCredentials: true });
  }

  public getCart(): Observable<Cart[]> {
    return this.httpClient.get<Cart[]>(`${this.backendUrl}/${this.cartPath}/my`, { withCredentials: true });
  }
}
