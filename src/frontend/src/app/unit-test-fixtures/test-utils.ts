import { Observable } from 'rxjs';
import { Category } from '../models/category';
import { Product } from '../models/product';

export const changeInputValue = (htmlInputElement: any, value: any) => {
    htmlInputElement.value = value;
    htmlInputElement.dispatchEvent(new Event('input'));
};

export const changeSelectIndex = (htmlInputElement: any, selectedIndex: any) => {
    htmlInputElement.selectedIndex = selectedIndex;
    htmlInputElement.dispatchEvent(new Event('change'));
};

export function completeObservable<T>(value: T | undefined = undefined): Observable<T> {
    return new Observable<T>((o) => {
        if (value) {               
            o.next(value);
        } else {
            o.next();
        }
        o.complete();
    });
}

export function errorObservable<T>(error: any): Observable<T> {
    return new Observable<T>(o => o.error(error));
}

export const createProduct = (id: number | undefined = undefined, 
    name: string | undefined = undefined, 
    price: number | undefined = undefined,
    description: string | undefined = undefined,
    category: Category | undefined = undefined) => {
  return { 
    id: id ?? 0,
    name: name ?? 'product',
    category: category ?? {
      id: 1,
      name: 'category',
      deleted: false
    },
    price: price ?? 100,
    description: description ?? 'Desc1234',
    deleted: false
  } as Product
};


export const stubbedProducts = () => {
  return [
    createProduct(1, 'Product#1'),
    createProduct(2, 'Product#2'),
    createProduct(3, 'Product#3'),
    createProduct(4, 'Product#4'),
    createProduct(5, 'Product#5')
  ] as Product[];
};

export const createCategory = (id: number | undefined = undefined, 
    name: string | undefined = undefined) => {
  return { 
    id: id ?? 0,
    name: name ?? 'category',
    deleted: false
  } as Category
};

export const stubbedCategories = () => {
  return [
    createProduct(1, 'Category#1'),
    createProduct(2, 'Category#2'),
    createProduct(3, 'Category#3'),
    createProduct(4, 'Category#4'),
    createProduct(5, 'Category#5')
  ] as Category[];
}
