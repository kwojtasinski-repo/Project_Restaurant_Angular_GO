import { Component } from '@angular/core';
import { Order } from 'src/app/models/order';

@Component({
  selector: 'app-my-orders',
  templateUrl: './my-orders.component.html',
  styleUrls: ['./my-orders.component.scss']
})
export class MyOrdersComponent {
  public orders: Order[] = [
    { 
      id: 1,
      orderNumber: new Date().toString()+1,
      price: 100.50,
      userId: 1,
      created: new Date(),
      orderProducts: [],
      modified: new Date()
    },
    { 
      id: 2,
      orderNumber: new Date().toString()+2,
      price: 100.50,
      userId: 1,
      created: new Date(),
      orderProducts: [],
      modified: new Date()
    },
    { 
      id: 3,
      orderNumber: new Date().toString()+3,
      price: 100.50,
      userId: 1,
      created: new Date(),
      orderProducts: [],
      modified: new Date()
    }
  ];  
  public ordersToShow: Order[] = this.orders;
  public term: string = '';
  
  public search(term: string): void {
    this.ordersToShow = this.orders.filter(p => p.orderNumber.toLocaleLowerCase().startsWith(term.toLocaleLowerCase()));
  }
}
