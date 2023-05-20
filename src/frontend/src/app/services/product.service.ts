import { Injectable } from '@angular/core';
import { Product } from '../models/product';
import { Observable, of } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ProductService {
  constructor() { }

  public add(product: Product): Observable<void> {
    this.products.push(product);
    return new Observable((ob) => { ob.next(); ob.complete(); });
  }

  public update(product: Product): Observable<void> {
    const response: Observable<void> = new Observable((ob) => { ob.next(); ob.complete(); });
    const index = this.products.findIndex(p => p.id === product.id);
    if (index < 0) {
      return response;
    }
    this.products[index] = product;
    return response;
  }

  public getAll(): Observable<Product[]> {
    return of(this.products.filter(p => p.deleted !== true));
  }

  public get(id: number): Observable<Product | undefined> {
    return of(this.products.find(p => p.id === id));
  }

  private products: Product[] = [
    {
      id: 1,
      name: 'Product #1',
      price: 10.25,
      description: 'Description #1',
      category: {
        id: 1,
        name: 'Category #1',
        deleted: false
      },
      deleted: false
    },
    {
      id: 2,
      name: 'Product #2',
      price: 525.25,
      description: 'Description #2',
      category: {
        id: 2,
        name: 'Category #2',
        deleted: false
      },
      deleted: false
    },
    {
      id: 3,
      name: 'Product #3',
      price: 150.25,
      description: '',
      category: {
        id: 2,
        name: 'Category #2',
        deleted: false
      },
      deleted: false
    },
    {
      id: 4,
      name: 'Product #4',
      price: 35.25,
      description: 'Description #4',
      category: {
        id: 3,
        name: 'Category #3',
        deleted: true
      },
      deleted: false
    },
    {
      id: 5,
      name: 'Product #5',
      price: 15.25,
      description: 'Description #5',
      category: {
        id: 3,
        name: 'Category #3',
        deleted: true
      },
      deleted: true
    }
  ];
}
