import { Category } from "./category"

export class Product {
    id: string = '';
	name: string = '';
	description: string = '';
	price: number = 0;
	category: Category | null = null;
    deleted: boolean = false;
}
