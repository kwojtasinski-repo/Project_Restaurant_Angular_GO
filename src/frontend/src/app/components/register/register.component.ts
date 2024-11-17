import { Component, OnDestroy, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators, FormsModule, ReactiveFormsModule } from '@angular/forms';
import { getEmail, getError, getPassword, getPasswordConfirm, getRegisterRequestState } from 'src/app/stores/register/register.selectors';
import { RegisterState } from 'src/app/stores/register/register.state';
import { SpinnerVersion } from '../spinner-button/spinner-version';
import { Store } from '@ngrx/store';
import { Subscription, debounceTime } from 'rxjs';
import { PATTERN_ONE_UPPER_ONE_LOWER_ONE_SPECIAL_CHARACTER, checkMatchValidator, getValidationMessage } from 'src/app/validations/validations';
import { Actions, ofType } from '@ngrx/effects';
import * as RegisterActions from 'src/app/stores/register/register.actions';
import { SpinnerButtonComponent } from '../spinner-button/spinner-button.component';
import { RouterLink } from '@angular/router';
import { AsyncPipe, KeyValuePipe } from '@angular/common';

@Component({
    selector: 'app-register',
    templateUrl: './register.component.html',
    styleUrls: ['./register.component.scss'],
    standalone: true,
    imports: [FormsModule, ReactiveFormsModule, RouterLink, SpinnerButtonComponent, AsyncPipe, KeyValuePipe]
})
export class RegisterComponent implements OnInit, OnDestroy {
  public registerForm: FormGroup;
  public error$ = this.store.select(getError);
  public loginRequestState$ = this.store.select(getRegisterRequestState);
  public spinnerVersion = SpinnerVersion.grow;
  public email$ = this.store.select(getEmail);
  public password$ = this.store.select(getPassword);
  public confirmPassword$ = this.store.select(getPasswordConfirm);
  private loginError$: Subscription = new Subscription();

  constructor(private store: Store<RegisterState>, private actions$: Actions) {
    this.registerForm = new FormGroup({
        emailAddress: new FormControl('', Validators.compose([Validators.required, Validators.email])),
        password: new FormControl('', Validators.compose([ 
          Validators.required, 
          Validators.minLength(12), 
          Validators.maxLength(64), 
          Validators.pattern(PATTERN_ONE_UPPER_ONE_LOWER_ONE_SPECIAL_CHARACTER)])
        ),
        confirmPassword: new FormControl('', Validators.compose([ 
          Validators.required, 
          Validators.minLength(12), 
          Validators.maxLength(64), 
          Validators.pattern(PATTERN_ONE_UPPER_ONE_LOWER_ONE_SPECIAL_CHARACTER)]),
        )
      },
      {
        validators: checkMatchValidator({ fieldName: 'password', labelName: 'Hasło' } , { fieldName: 'confirmPassword', labelName: 'Powtórz hasło' })
      }
    );
  }

  public ngOnInit(): void {
    this.loginError$ = this.actions$
      .pipe(ofType(RegisterActions.registerRequestFailed))
      .subscribe(() => 
        this.registerForm.setValue({
          emailAddress: this.registerForm.get('emailAddress')?.value ?? '',
          password: '',
          confirmPassword: ''
        }, { emitEvent: false })
      );

      this.registerForm.valueChanges
        .pipe(debounceTime(10))
        .subscribe(val => 
          this.store.dispatch(RegisterActions.registerFormUpdate({
            form: {
              email: val.emailAddress,
              password: val.password,
              confirmPassword: val.confirmPassword,
            }})
          )
        )
  }

  public ngOnDestroy(): void {
    this.loginError$.unsubscribe();
    this.store.dispatch(RegisterActions.clearErrors());
  }

  public getErrorMessage(error: any): string | null {
    return getValidationMessage(error);
  }

  public onSubmit() {
    if (this.registerForm.invalid) {
      Object.keys(this.registerForm.controls).forEach(key => {
        this.registerForm.get(key)?.markAsDirty();
      });
      return;
    }
    this.store.dispatch(RegisterActions.registerRequestBegin());
  }
}
