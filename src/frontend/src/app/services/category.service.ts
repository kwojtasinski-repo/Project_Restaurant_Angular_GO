import { Injectable, inject } from '@angular/core';
import { Observable } from 'rxjs';
import { Category } from '../models/category';
import { HttpClient } from '@angular/common/http';
import { API_URL } from '../providers/api-url-provider';

@Injectable({
  providedIn: 'root'
})
export class CategoryService {
  private httpClient = inject(HttpClient);
  private backendUrl = inject<string>(API_URL);

  private categoryPath = 'api/categories';

  public add(category: Category): Observable<void> {
    return this.httpClient.post<void>(`${this.backendUrl}/${this.categoryPath}`, category, { withCredentials: true });
  }

  public update(category: Category): Observable<void> {
    return this.httpClient.put<void>(`${this.backendUrl}/${this.categoryPath}/${category.id}`, category, { withCredentials: true });
  }

  public getAll(): Observable<Category[]> {
    return this.httpClient.get<Category[]>(`${this.backendUrl}/${this.categoryPath}`, { withCredentials: true });
  }

  public get(id: string): Observable<Category | undefined> {
    return this.httpClient.get<Category | undefined>(`${this.backendUrl}/${this.categoryPath}/${id}`, { withCredentials: true });
  }
}
