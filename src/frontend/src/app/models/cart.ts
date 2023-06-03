import { Product } from "./product";

export class Cart {
    id: number = 0;
    product: Product | undefined;
    userId: number = 0;
}