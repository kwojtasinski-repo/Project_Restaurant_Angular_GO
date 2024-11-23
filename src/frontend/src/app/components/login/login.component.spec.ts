import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LoginComponent } from './login.component';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { provideMockStore } from '@ngrx/store/testing';
import { initialState } from 'src/app/stores/login/login.reducers';
import { provideHttpClient, withInterceptorsFromDi } from '@angular/common/http';
import { provideHttpClientTesting } from '@angular/common/http/testing';
import { Observable } from 'rxjs';
import { provideMockActions } from '@ngrx/effects/testing';
import { SpinnerButtonComponent } from '../spinner-button/spinner-button.component';
import { RouterModule } from '@angular/router';

describe('LoginComponent', () => {
  let actions$: Observable<any>;
  let component: LoginComponent;
  let fixture: ComponentFixture<LoginComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
    imports: [FormsModule,
        ReactiveFormsModule,
        RouterModule.forRoot([]),
        LoginComponent, SpinnerButtonComponent],
    providers: [
        provideMockStore({ initialState }),
        provideMockActions(() => actions$),
        provideHttpClient(withInterceptorsFromDi()),
        provideHttpClientTesting()
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
