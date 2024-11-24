import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CartsComponent } from './carts.component';
import { TestSharedModule } from 'src/app/unit-test-fixtures/test-share-module';

describe('CartsComponent', () => {
  let component: CartsComponent;
  let fixture: ComponentFixture<CartsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        CartsComponent,
        TestSharedModule
      ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(CartsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
