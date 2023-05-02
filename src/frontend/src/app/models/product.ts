import { Category } from "./category"

export class Product {
    id: number = 0;
	name: string = '';
	description: string = '';
	price: number = 0;
	category: Category | null = null;
    deleted: boolean = false;
}
