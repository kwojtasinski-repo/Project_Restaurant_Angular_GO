import { ComponentFixture, TestBed } from '@angular/core/testing';

import { MyOrdersComponent } from './my-orders.component';
import { SearchBarComponent } from '../../search-bar/search-bar.component';
import { TestSharedModule } from 'src/app/unit-test-fixtures/test-share-module';

describe('MyOrdersComponent', () => {
  let component: MyOrdersComponent;
  let fixture: ComponentFixture<MyOrdersComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        MyOrdersComponent,
        SearchBarComponent,
        TestSharedModule
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
