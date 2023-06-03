import { Inject, Injectable } from '@angular/core';
import { Product } from '../models/product';
import { Observable, catchError, of } from 'rxjs';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class ProductService {
  private productPath = 'api/products';

  constructor(private httpClient: HttpClient, @Inject('API_URL') private backendUrl: string) { }

  public add(product: Product): Observable<void> {
    return this.httpClient.post<void>(`${this.backendUrl}/${this.productPath}`, product, { withCredentials: true });
  }

  public update(product: Product): Observable<void> {
    return this.httpClient.put<void>(`${this.backendUrl}/${this.productPath}/${product.id}`, product, { withCredentials: true });
  }

  public getAll(): Observable<Product[]> {
    return this.httpClient.get<Product[]>(`${this.backendUrl}/${this.productPath}`, { withCredentials: true });
  }

  public get(id: number): Observable<Product | undefined> {
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
