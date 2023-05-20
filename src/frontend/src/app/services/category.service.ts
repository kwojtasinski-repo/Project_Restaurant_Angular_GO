import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';
import { Category } from '../models/category';

@Injectable({
  providedIn: 'root'
})
export class CategoryService {

  constructor() { }

  public add(category: Category): Observable<void> {
    this.categories.push(category);
    return new Observable((ob) => { ob.next(); ob.complete(); });
  }

  public update(category: Category): Observable<void> {
    const response: Observable<void> = new Observable((ob) => { ob.next(); ob.complete(); });
    const index = this.categories.findIndex(c => c.id === category.id);
    if (index < 0) {
      return response;
    }
    this.categories[index] = category;
    return response;
  }

  public getAll(): Observable<Category[]> {
    return of(this.categories);
  }

  public get(id: number): Observable<Category | undefined> {
    return of(this.categories.find(c => c.id === id));
  }

  private categories: Category[] = [
    {
      id: 1,
      name: 'Category #1',
      deleted: false
    },
    {
      id: 2,
      name: 'Category #2',
      deleted: false
    },
    {
      id: 3,
      name: 'Category #3',
      deleted: true
    },
  ];
}
