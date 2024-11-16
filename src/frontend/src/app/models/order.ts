import { OrderProduct } from './orderProduct';

export class Order {
    id: string = '';
    orderNumber: string = '';
    price: number = 0;
    created: Date = new Date();
    modified: Date | undefined;
    userId: number = 0;
    orderProducts: OrderProduct[] = [];
}
