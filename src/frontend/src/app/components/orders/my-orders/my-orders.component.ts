import { Component, OnInit } from '@angular/core';
import { NgxSpinnerService } from 'ngx-spinner';
import { take } from 'rxjs';
import { Order } from 'src/app/models/order';
import { OrderService } from 'src/app/services/order.service';

@Component({
  selector: 'app-my-orders',
  templateUrl: './my-orders.component.html',
  styleUrls: ['./my-orders.component.scss']
})
export class MyOrdersComponent implements OnInit {
  public orders: Order[] = [];
  public isLoading: boolean = true;
  public ordersToShow: Order[] = [];
  public term: string = '';
  
  constructor(private orderService: OrderService, private spinnerService: NgxSpinnerService) { }
  
  public ngOnInit(): void {
    this.spinnerService.show();
    this.orderService.getMyOrders()
      .pipe(take(1))
      .subscribe(o => {
        this.isLoading = false;
        this.orders = o;
        this.ordersToShow = o;
        this.spinnerService.hide();
      });
  }

  public search(term: string): void {
    this.ordersToShow = this.orders.filter(p => p.orderNumber.toLocaleLowerCase().startsWith(term.toLocaleLowerCase()));
  }
}
