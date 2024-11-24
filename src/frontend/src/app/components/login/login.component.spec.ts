import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LoginComponent } from './login.component';
import { Observable } from 'rxjs';
import { provideMockActions } from '@ngrx/effects/testing';
import { SpinnerButtonComponent } from '../spinner-button/spinner-button.component';
import { TestSharedModule } from 'src/app/unit-test-fixtures/test-share-module';

describe('LoginComponent', () => {
  let actions$: Observable<any>;
  let component: LoginComponent;
  let fixture: ComponentFixture<LoginComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        LoginComponent,
        SpinnerButtonComponent,
        TestSharedModule
      ],
      providers: [
          provideMockActions(() => actions$)
      ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(LoginComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
