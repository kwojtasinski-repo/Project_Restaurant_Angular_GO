import { Component, OnInit } from '@angular/core';
import { Store } from '@ngrx/store';
import { OrderState } from 'src/app/stores/order/order.state';
import { fetchOrder } from 'src/app/stores/order/order.actions';
import { getOrder } from 'src/app/stores/order/order.selectors';
import { ActivatedRoute, RouterLink } from '@angular/router';
import { getFetchState } from 'src/app/stores/order/order.selectors';
import { MoneyPipe } from '../../../pipes/money-pipe';
import { NgIf, NgFor, AsyncPipe } from '@angular/common';

@Component({
    selector: 'app-order-view',
    templateUrl: './order-view.component.html',
    styleUrls: ['./order-view.component.scss'],
    standalone: true,
    imports: [
        NgIf,
        NgFor,
        RouterLink,
        AsyncPipe,
        MoneyPipe,
    ],
})
export class OrderViewComponent implements OnInit {
  public order$ = this.store.select(getOrder);
  public fetchState$ = this.store.select(getFetchState);
  private id = '';

  constructor(private store: Store<OrderState>, private route: ActivatedRoute) { }

  public ngOnInit(): void {
    this.id = this.route.snapshot.paramMap.get('id') ?? '';
    this.store.dispatch(fetchOrder({ id: this.id }));
  }
}
