import { Product } from '../models/product';
import { ProductService } from '../services/product.service';
import { Observable } from 'rxjs';
import { completeObservable, errorObservable } from './test-utils';
import { ProductSendDto } from '../models/product-send-dto';
import { Injectable } from '@angular/core';

@Injectable()
export class InMemoryProductService extends ProductService {
    private products: Product[] = [];

    public override add(product: ProductSendDto): Observable<void> {
        product.id = this.products.length > 0 ? (new Number(this.products[this.products.length - 1].id).valueOf() + 1).toString() : '1';
        this.products.push({
            id: product.id,
            name: product.name,
            price: product.price,
            description: product.description ?? '',
            category: {
                id: product.categoryId,
                name: '',
                deleted: false
            },
            deleted: false,
        })
        return completeObservable<void>();
    }
    
    public override update(product: ProductSendDto): Observable<void> {
        const productIndex = this.products.findIndex(p => p.id === product.id);
        if (productIndex === -1) {
            return errorObservable(`Product with id ${product.id} not found`);
        }
        
        return completeObservable<void>();
    }
    
    public override getAll(): Observable<Product[]> {
        return completeObservable<Product[]>(this.products);
    }
    
    public override get(id: string): Observable<Product | undefined> {
        const product = this.products.find(p => p.id === id);
        return completeObservable<Product | undefined>(product);
    }
}
