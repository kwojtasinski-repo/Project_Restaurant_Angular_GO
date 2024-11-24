import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SpinnerButtonComponent } from './spinner-button.component';
import { TestSharedModule } from 'src/app/unit-test-fixtures/test-share-module';

describe('SpinnerButtonComponent', () => {
  let component: SpinnerButtonComponent;
  let fixture: ComponentFixture<SpinnerButtonComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [SpinnerButtonComponent, TestSharedModule]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SpinnerButtonComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
