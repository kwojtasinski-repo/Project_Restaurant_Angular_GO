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
  public ordersToShow: Order[] = [];
  public term: string = '';
  public error: string | undefined;
  
  constructor(private orderService: OrderService, private spinnerService: NgxSpinnerService) { }
  
  public ngOnInit(): void {
    this.spinnerService.show();
    this.orderService.getMyOrders()
      .pipe(take(1))
      .subscribe({ next: o => {
          this.orders = o;
          this.ordersToShow = o;
          this.spinnerService.hide();
        }, error: error => {
          if (error.status === 0) {
            this.error = 'Sprawdź połączenie z internetem';
          } else if (error.status === 500) {
            this.error = 'Coś poszło nie tak, spróbuj ponownie później';
          }
          this.spinnerService.hide();
          console.error(error);
        }
      });
  }

  public search(term: string): void {
    this.ordersToShow = this.orders.filter(p => p.orderNumber.toLocaleLowerCase().startsWith(term.toLocaleLowerCase()));
  }
}
