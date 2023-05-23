import { ComponentFixture, TestBed } from '@angular/core/testing';

import { MyOrdersComponent } from './my-orders.component';
import { SearchBarComponent } from '../../search-bar/search-bar.component';
import { NgxSpinnerModule } from 'ngx-spinner';
import { MoneyPipe } from 'src/app/pipes/money-pipe';
import { RouterTestingModule } from '@angular/router/testing';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';

describe('MyOrdersComponent', () => {
  let component: MyOrdersComponent;
  let fixture: ComponentFixture<MyOrdersComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ 
        MyOrdersComponent, 
        SearchBarComponent, 
        MoneyPipe
      ],
      imports: [
        RouterTestingModule,
        NgxSpinnerModule,
        ReactiveFormsModule,
        FormsModule
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
