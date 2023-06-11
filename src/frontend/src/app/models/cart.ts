export class Cart {
    id: number = 0;
    product: CartProduct | undefined;
    userId: number = 0;
}

interface CartProduct {
    id: number;
	name: string;
	description: string;
	price: number;
}
