export class Cart {
    id: string = '';
    product: CartProduct | undefined;
    userId: number = 0;
}

interface CartProduct {
    id: string;
	name: string;
	description: string;
	price: number;
}
