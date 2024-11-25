import { Injectable, inject } from '@angular/core';
import { Cart } from '../models/cart';
import { Observable } from 'rxjs';
import { Product } from '../models/product';
import { HttpClient } from '@angular/common/http';
import { API_URL } from '../providers/api-url-provider';

@Injectable({
  providedIn: 'root'
})
export class CartService {
  private httpClient = inject(HttpClient);
  private backendUrl = inject<string>(API_URL);

  private cartPath = 'api/carts';

  public add(product: Product): Observable<void> {
    return this.httpClient.post<void>(`${this.backendUrl}/${this.cartPath}`, { productId: product.id }, { withCredentials: true });
  }

  public delete(id: string): Observable<void> {
    return this.httpClient.delete<void>(`${this.backendUrl}/${this.cartPath}/${id}`, { withCredentials: true });
  }

  public getCart(): Observable<Cart[]> {
    return this.httpClient.get<Cart[]>(`${this.backendUrl}/${this.cartPath}/my`, { withCredentials: true });
  }
}
