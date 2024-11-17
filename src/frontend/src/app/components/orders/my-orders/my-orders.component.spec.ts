import { ComponentFixture, TestBed } from '@angular/core/testing';
import { provideMockStore } from '@ngrx/store/testing';

import { MyOrdersComponent } from './my-orders.component';
import { SearchBarComponent } from '../../search-bar/search-bar.component';
import { NgxSpinnerModule } from 'ngx-spinner';
import { MoneyPipe } from 'src/app/pipes/money-pipe';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { initialState } from 'src/app/stores/order/order.reducers';
import { HttpClientModule } from '@angular/common/http';
import { provideRouter } from '@angular/router';

describe('MyOrdersComponent', () => {
  let component: MyOrdersComponent;
  let fixture: ComponentFixture<MyOrdersComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
    imports: [
        NgxSpinnerModule,
        ReactiveFormsModule,
        FormsModule,
        HttpClientModule,
        MyOrdersComponent,
        SearchBarComponent,
        MoneyPipe
    ],
    providers: [
        provideRouter([]),
        provideMockStore({ initialState }),
        provideMockStore({ initialState }),
        {
            provide: 'API_URL', useValue: ''
        }
    ]
})
    .compileComponents();

    fixture = TestBed.createComponent(MyOrdersComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
