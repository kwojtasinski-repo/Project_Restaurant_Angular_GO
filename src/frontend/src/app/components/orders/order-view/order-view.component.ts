import { Component, OnInit } from '@angular/core';
import { Store } from '@ngrx/store';
import { OrderState } from 'src/app/stores/order/order.state';
import { fetchOrder } from 'src/app/stores/order/order.actions';
import { getOrder } from 'src/app/stores/order/order.selectors';
import { ActivatedRoute } from '@angular/router';
import { getFetchState } from 'src/app/stores/order/order.selectors';

@Component({
  selector: 'app-order-view',
  templateUrl: './order-view.component.html',
  styleUrls: ['./order-view.component.scss'],
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
