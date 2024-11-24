import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AddProductsComponent } from './add-products.component';
import { ProductFormComponent } from '../product-form/product-form.component';
import { TestSharedModule } from 'src/app/unit-test-fixtures/test-share-module';

describe('AddProductsComponent', () => {
  let component: AddProductsComponent;
  let fixture: ComponentFixture<AddProductsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        AddProductsComponent,
        ProductFormComponent,
        TestSharedModule
      ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(AddProductsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
