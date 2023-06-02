import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';
import { Category } from '../models/category';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class CategoryService {
  private backendUrl = (window as any )['__env']['backendUrl'];

  constructor(private httpClient: HttpClient) { }

  public add(category: Category): Observable<void> {
    return this.httpClient.post<void>(`${this.backendUrl}/api/categories`, category, { withCredentials: true });
  }

  public update(category: Category): Observable<void> {
    return this.httpClient.put<void>(`${this.backendUrl}/api/categories/${category.id}`, category, { withCredentials: true });
  }

  public getAll(): Observable<Category[]> {
    return this.httpClient.get<Category[]>(`${this.backendUrl}/api/categories`, { withCredentials: true });
  }

  public get(id: number): Observable<Category | undefined> {
    return this.httpClient.get<Category | undefined>(`${this.backendUrl}/api/categories/${id}`, { withCredentials: true });
  }
}
