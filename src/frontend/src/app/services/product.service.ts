import { Injectable, inject } from '@angular/core';
import { Product } from '../models/product';
import { Observable, catchError, of } from 'rxjs';
import { HttpClient } from '@angular/common/http';
import { ProductSendDto } from '../models/product-send-dto';
import { API_URL } from '../providers/api-url-provider';

@Injectable({
  providedIn: 'root'
})
export class ProductService {
  private httpClient = inject(HttpClient);
  private backendUrl = inject<string>(API_URL);

  private productPath = 'api/products';

  public add(product: ProductSendDto): Observable<void> {
    return this.httpClient.post<void>(`${this.backendUrl}/${this.productPath}`, product, { withCredentials: true });
  }

  public update(product: ProductSendDto): Observable<void> {
    return this.httpClient.put<void>(`${this.backendUrl}/${this.productPath}/${product.id}`, product, { withCredentials: true });
  }

  public getAll(): Observable<Product[]> {
    return this.httpClient.get<Product[]>(`${this.backendUrl}/${this.productPath}`, { withCredentials: true });
  }

  public get(id: string): Observable<Product | undefined> {
    return this.httpClient.get<Product | undefined>(`${this.backendUrl}/${this.productPath}/${id}`, { withCredentials: true })
      .pipe(
        catchError(err => {
          if (err.status === 404) {
            return of(undefined)
          }

          throw err;
        })
      );
  }
}
