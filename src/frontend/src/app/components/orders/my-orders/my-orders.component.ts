import { Component, OnInit, WritableSignal, computed, inject, signal } from '@angular/core';
import { NgxSpinnerService } from 'ngx-spinner';
import { EMPTY, catchError, finalize, take, tap } from 'rxjs';
import { Order } from 'src/app/models/order';
import { OrderService } from 'src/app/services/order.service';
import { MoneyPipe } from '../../../pipes/money-pipe';
import { RouterLink } from '@angular/router';
import { SearchBarComponent } from '../../search-bar/search-bar.component';

@Component({
    selector: 'app-my-orders',
    templateUrl: './my-orders.component.html',
    styleUrls: ['./my-orders.component.scss'],
    standalone: true,
    imports: [SearchBarComponent, RouterLink, MoneyPipe]
})
export class MyOrdersComponent implements OnInit {
  private orderService = inject(OrderService);
  private spinnerService = inject(NgxSpinnerService);

  public orders: WritableSignal<Order[]> = signal([]);
  public term: WritableSignal<string> = signal('');
  public error: WritableSignal<string | undefined> = signal(undefined);
  
  public ordersToShow = computed(() => {
    const term = this.term();
    return this.orders().filter(c =>
      c.orderNumber.toLocaleLowerCase().startsWith(term.toLocaleLowerCase())
    );
  });

  public ngOnInit(): void {
    this.orderService.getMyOrders()
      .pipe(
        take(1),
        tap(() => this.spinnerService.show()),
        finalize(() => this.spinnerService.hide()),
        catchError((error) => {
          if (error.status === 0) {
            this.error.set('Sprawdź połączenie z internetem');
          } else if (error.status === 500) {
            this.error.set('Coś poszło nie tak, spróbuj ponownie później');
          }
          console.error(error);
          return EMPTY;
        })
      ).subscribe(orders => this.orders.set(orders));
  }
}
