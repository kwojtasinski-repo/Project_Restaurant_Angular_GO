import { ComponentFixture, TestBed } from '@angular/core/testing';
import { initialState } from 'src/app/stores/register/register.reducers';
import { provideMockStore } from '@ngrx/store/testing';

import { RegisterComponent } from './register.component';
import { RouterModule } from '@angular/router';
import { provideMockActions } from '@ngrx/effects/testing';
import { Actions } from '@ngrx/effects';
import { SpinnerButtonComponent } from '../spinner-button/spinner-button.component';
import { ReactiveFormsModule } from '@angular/forms';
import { RegisterState } from 'src/app/stores/register/register.state';
import { Store } from '@ngrx/store';
import { changeInputValue } from 'src/app/unit-test-fixtures/test-utils';
import { registerRequestBegin } from 'src/app/stores/register/register.actions';

describe('RegisterComponent', () => {
  let component: RegisterComponent;
  let fixture: ComponentFixture<RegisterComponent>;
  let actions: Actions;
  let store : Store<RegisterState>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ 
        RegisterComponent,
        SpinnerButtonComponent
      ],
      imports: [
        ReactiveFormsModule,
        RouterModule.forRoot([]),
      ],
      providers: [
        provideMockStore({ initialState }),
        provideMockActions(() => actions)
      ]
    })
    .compileComponents();

    store = TestBed.inject(Store<RegisterState>);
    fixture = TestBed.createComponent(RegisterComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should create register form', () => {
    const emailInput = fixture.nativeElement.querySelector('#email-input');
    const passwordInput = fixture.nativeElement.querySelector('#password-input');
    const confirmPasswordInput = fixture.nativeElement.querySelector('#confirm-password-input');

    expect(emailInput).not.toBeUndefined();
    expect(emailInput).not.toBeNull();
    expect(passwordInput).not.toBeUndefined();
    expect(passwordInput).not.toBeNull();
    expect(confirmPasswordInput).not.toBeUndefined();
    expect(confirmPasswordInput).not.toBeNull();
  });

  it('should update values in form', () => {
    const expectedEmail = 'email@email.com';
    const expectedPassword = 'PAsW0RD!123456789';

    fixture.componentInstance.registerForm.setValue({
      'emailAddress': expectedEmail,
      'password': expectedPassword,
      'confirmPassword': expectedPassword
    });

    
    const emailInput = fixture.nativeElement.querySelector('#email-input');
    const passwordInput = fixture.nativeElement.querySelector('#password-input');
    const confirmPasswordInput = fixture.nativeElement.querySelector('#confirm-password-input');
    expect(emailInput.value).toEqual(expectedEmail);
    expect(passwordInput.value).toEqual(expectedPassword);
    expect(confirmPasswordInput.value).toEqual(expectedPassword);
  })

  it('should show error while passed invalid input values', () => {
    const emailInput = fixture.nativeElement.querySelector('#email-input');
    const passwordInput = fixture.nativeElement.querySelector('#password-input');
    const confirmPasswordInput = fixture.nativeElement.querySelector('#confirm-password-input');

    changeInputValue(emailInput, 'abc');
    changeInputValue(passwordInput, 'pas');
    changeInputValue(confirmPasswordInput, 'conf');
    fixture.detectChanges();

    const validationErrors = Array.from(fixture.nativeElement.querySelectorAll('.invalid-feedback'));
    expect(validationErrors).not.toBeNull();
    expect(validationErrors).not.toBeUndefined();
    expect(validationErrors.length).toBeGreaterThan(0);
    expect(validationErrors.some((v: any) => v.innerHTML === 'Pole powinno zawierać małą i dużą literę oraz specjalny znak znak'))
      .not.toBeNull();
    expect(validationErrors.some((v: any) => v.innerHTML === 'Pole powinno zawierać małą i dużą literę oraz specjalny znak znak'))
      .not.toBeUndefined();
    expect(validationErrors.some((v: any) => v.innerHTML === 'nie są identyczne')).not.toBeNull();
    expect(validationErrors.some((v: any) => v.innerHTML === 'nie są identyczne')).not.toBeUndefined();
    expect(validationErrors.some((v: any) => v.innerHTML === 'Niepoprawny adres email')).not.toBeNull();
    expect(validationErrors.some((v: any) => v.innerHTML === 'Niepoprawny adres email')).not.toBeUndefined();
  });

  it('should show error while send empty form', () => {
    const submitButton = fixture.nativeElement.querySelector('.btn.btn-primary.me-2');
    
    submitButton.click();
    fixture.detectChanges();

    const validationErrors = Array.from(fixture.nativeElement.querySelectorAll('.invalid-feedback'));
    expect(validationErrors).not.toBeNull();
    expect(validationErrors).not.toBeUndefined();
    expect(validationErrors.length).toBeGreaterThan(0);
    expect(validationErrors.some((v: any) => v.innerHTML === 'Pole jest wymagane')).not.toBeNull();
    expect(validationErrors.some((v: any) => v.innerHTML === 'Pole jest wymagane')).not.toBeUndefined();
  });

  it('should send form', () => {
    const submitButton = fixture.nativeElement.querySelector('.btn.btn-primary.me-2');
    fixture.componentInstance.registerForm.setValue({
      'emailAddress': 'expectedEmail@email',
      'password': 'expectedPassword1234!aV',
      'confirmPassword': 'expectedPassword1234!aV'
    });
    const storeSpy = spyOn(store, 'dispatch');

    submitButton.click();
    fixture.detectChanges();

    expect(storeSpy).toHaveBeenCalled();
    expect(storeSpy).toHaveBeenCalledWith(registerRequestBegin());
  });
});
