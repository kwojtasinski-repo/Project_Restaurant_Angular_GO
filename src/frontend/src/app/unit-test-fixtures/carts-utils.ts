import { Cart } from '../models/cart';
import { Product } from '../models/product';

export function createCartObject(id?: string, product?: Product, userId?: number): Cart {
  return {
    id: id ?? '1',
    product: product ?? {
      id: '1',
      name: 'name#1',
      price: 100,
      description: 'desc'
    },
    userId: userId ?? 1
  } as Cart
}

export function createCart(): Cart[] {
  return [
    createCartObject('1'),
    createCartObject('2'),
    createCartObject('3')
  ] as Cart[];
}
