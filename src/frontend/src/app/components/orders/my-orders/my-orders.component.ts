import { Component, OnInit } from '@angular/core';
import { NgxSpinnerService } from 'ngx-spinner';
import { BehaviorSubject, EMPTY, Observable, catchError, finalize, map, shareReplay, take, tap } from 'rxjs';
import { Order } from 'src/app/models/order';
import { OrderService } from 'src/app/services/order.service';
import { MoneyPipe } from '../../../pipes/money-pipe';
import { RouterLink } from '@angular/router';
import { SearchBarComponent } from '../../search-bar/search-bar.component';
import { AsyncPipe } from '@angular/common';

@Component({
    selector: 'app-my-orders',
    templateUrl: './my-orders.component.html',
    styleUrls: ['./my-orders.component.scss'],
    standalone: true,
    imports: [SearchBarComponent, RouterLink, AsyncPipe, MoneyPipe]
})
export class MyOrdersComponent implements OnInit {
  public orders$: Observable<Order[]> = new BehaviorSubject([]);
  public ordersToShow$: Observable<Order[]> = new BehaviorSubject([]);
  public term: string = '';
  public error: string | undefined;
  
  constructor(private orderService: OrderService, private spinnerService: NgxSpinnerService) { }
  
  public ngOnInit(): void {
    this.orders$ = this.orderService.getMyOrders()
      .pipe(
        take(1),
        tap(() => this.spinnerService.show()),
        shareReplay(),
        finalize(() => this.spinnerService.hide()),
        catchError((error) => {
          if (error.status === 0) {
            this.error = 'Sprawdź połączenie z internetem';
          } else if (error.status === 500) {
            this.error = 'Coś poszło nie tak, spróbuj ponownie później';
          }
          console.error(error);
          return EMPTY;
        })
      );
    this.ordersToShow$ = this.orders$;
    this.orders$.subscribe();
  }

  public search(term: string): void {
    this.ordersToShow$ = this.orders$.pipe(
      map((orders) => orders.filter(p => p.orderNumber.toLocaleLowerCase().startsWith(term.toLocaleLowerCase())))
    );
  }
}
