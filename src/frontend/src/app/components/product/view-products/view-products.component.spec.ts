import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ViewProductsComponent } from './view-products.component';
import { RouterTestingModule } from '@angular/router/testing';
import { NgxSpinnerModule } from 'ngx-spinner';
import { initialState } from 'src/app/stores/login/login.reducers';
import { provideMockStore } from '@ngrx/store/testing';

describe('ViewProductsComponent', () => {
  let component: ViewProductsComponent;
  let fixture: ComponentFixture<ViewProductsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ViewProductsComponent ],
      imports: [
        RouterTestingModule,
        NgxSpinnerModule
      ],
      providers: [
        provideMockStore({ initialState })
      ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ViewProductsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
