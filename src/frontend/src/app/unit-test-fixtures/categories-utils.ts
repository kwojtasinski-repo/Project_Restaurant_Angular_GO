import { Category } from '../models/category';
import { createProduct } from './test-utils';

export const createCategory = (id: number | undefined = undefined, 
    name: string | undefined = undefined) => {
  return { 
    id: id ?? '0',
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
