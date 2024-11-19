import { Component, OnDestroy, OnInit, inject } from '@angular/core';
import { FormGroup, FormControl, Validators, FormsModule, ReactiveFormsModule } from '@angular/forms';
import { Actions, ofType } from '@ngrx/effects';
import { Store } from '@ngrx/store';
import { Subscription } from 'rxjs';
import { loginRequestFailed } from 'src/app/stores/login/login.actions';
import { loginRequest } from 'src/app/stores/login/login.actions';
import { getError, loginRequestState } from 'src/app/stores/login/login.selectors';
import { LoginState } from 'src/app/stores/login/login.state';
import { getValidationMessage } from 'src/app/validations/validations';
import { SpinnerVersion } from '../spinner-button/spinner-version';
import { SpinnerButtonComponent } from '../spinner-button/spinner-button.component';
import { RouterLink } from '@angular/router';
import { LoginFormDirective } from '../../directives/login-form-directive';
import { AsyncPipe, KeyValuePipe } from '@angular/common';

@Component({
    selector: 'app-login',
    templateUrl: './login.component.html',
    styleUrls: ['./login.component.scss'],
    standalone: true,
    imports: [FormsModule, LoginFormDirective, ReactiveFormsModule, RouterLink, SpinnerButtonComponent, AsyncPipe, KeyValuePipe]
})
export class LoginComponent implements OnInit, OnDestroy {
  private store = inject<Store<LoginState>>(Store);
  private actions$ = inject(Actions);

  public loginForm: FormGroup;
  public error$ = this.store.select(getError);
  public loginRequestState$ = this.store.select(loginRequestState);
  public spinnerVersion = SpinnerVersion.grow;
  private loginError$: Subscription = new Subscription();

  constructor() { 
    this.loginForm = new FormGroup({
      emailAddress: new FormControl('', Validators.compose([Validators.required, Validators.email])),
      password: new FormControl('', Validators.required)
    });
  }

  public ngOnInit(): void {
    this.loginError$ = this.actions$
      .pipe(ofType(loginRequestFailed))
      .subscribe(() => 
        this.loginForm.setValue({
          emailAddress: '',
          password: ''
        }, { emitEvent: false })
      );
  }

  public ngOnDestroy(): void {
    this.loginError$.unsubscribe();
  }

  public getErrorMessage(error: any): string | null {
    return getValidationMessage(error);
  }

  public onSubmit() {
    if (this.loginForm.invalid) {
      Object.keys(this.loginForm.controls).forEach(key => {
        this.loginForm.get(key)?.markAsDirty();
      });
      return;
    }
    this.store.dispatch(loginRequest());
  }
}
