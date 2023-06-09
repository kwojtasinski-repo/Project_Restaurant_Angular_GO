import { Product } from "../models/product";
import { ProductService } from '../services/product.service';
import { Observable } from 'rxjs';
import { completeObservable, errorObservable } from "./test-utils";
import { HttpClient } from "@angular/common/http";
import { HttpXhrBackend } from "@angular/common/http";

class InMemoryProductService extends ProductService {
    private products: Product[] = [];

    public override add(product: Product): Observable<void> {
        product.id = this.products.length > 0 ? this.products[this.products.length - 1].id + 1 : 1;
        this.products.push(product)
        return completeObservable<void>();
    }
    
    public override update(product: Product): Observable<void> {
        const productIndex = this.products.findIndex(p => p.id === product.id);
        if (productIndex === -1) {
            return errorObservable(`Product with id ${product.id} not found`);
        }
        
        return completeObservable<void>();
    }
    
    public override getAll(): Observable<Product[]> {
        return completeObservable<Product[]>(this.products);
    }
    
    public override get(id: number): Observable<Product | undefined> {
        const product = this.products.find(p => p.id === id);
        return completeObservable<Product | undefined>(product);
    }
}

const productService = new InMemoryProductService(new HttpClient(new HttpXhrBackend({
    build: () => new XMLHttpRequest()
})), "");
export default productService;