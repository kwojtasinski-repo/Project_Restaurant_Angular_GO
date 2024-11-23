import { Observable } from 'rxjs';
import { completeObservable, errorObservable } from './test-utils';
import { CategoryService } from '../services/category.service';
import { Category } from '../models/category';
import { Injectable } from '@angular/core';

@Injectable()
export class InMemoryCategoryService extends CategoryService {
    private categories: Category[] = [];

    public override add(category: Category): Observable<void> {
        category.id = this.categories.length > 0 ? (new Number(this.categories[this.categories.length - 1].id).valueOf() + 1).toString() : '1';
        this.categories.push(category)
        return completeObservable<void>();
    }
    
    public override update(product: Category): Observable<void> {
        const productIndex = this.categories.findIndex(p => p.id === product.id);
        if (productIndex === -1) {
            return errorObservable(`Category with id ${product.id} not found`);
        }
        
        return completeObservable<void>();
    }
    
    public override getAll(): Observable<Category[]> {
        return completeObservable<Category[]>(this.categories);
    }
    
    public override get(id: string): Observable<Category | undefined> {
        const product = this.categories.find(p => p.id === id);
        return completeObservable<Category | undefined>(product);
    }
}
