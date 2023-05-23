import { Component } from '@angular/core';
import { Order } from 'src/app/models/order';

@Component({
  selector: 'app-order-view',
  templateUrl: './order-view.component.html',
  styleUrls: ['./order-view.component.scss']
})
export class OrderViewComponent {
  public order: Order = { 
    id: 1,
    orderNumber: new Date().toString(),
    price: 100.50,
    userId: 1,
    created: new Date(),
    orderProducts: [
      {
        id: 1,
        name: 'abc#1',
        price: 50.25,
        productId: 1
      },
      {
        id: 2,
        name: 'abc#1',
        price: 50.25,
        productId: 1
      },
    ],
    modified: new Date()
  };
}
