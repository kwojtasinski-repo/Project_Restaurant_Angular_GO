import { Component } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { getError, getRegisterRequestState } from 'src/app/stores/register/register.selectors';
import { RegisterState } from 'src/app/stores/register/register.state';
import { SpinnerVersion } from '../spinner-button/spinner-version';
import { Store } from '@ngrx/store';
import { Subscription } from 'rxjs';
import { checkMatchValidator, getValidationMessage } from 'src/app/validations/validations';
import { Actions, ofType } from '@ngrx/effects';
import { registerRequestBegin, registerRequestFailed } from 'src/app/stores/register/register.actions';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.scss']
})
export class RegisterComponent {
  public registerForm: FormGroup;
  public error$ = this.store.select(getError);
  public loginRequestState$ = this.store.select(getRegisterRequestState);
  public spinnerVersion = SpinnerVersion.grow;
  private loginError$: Subscription = new Subscription();

  constructor(private store: Store<RegisterState>, private actions$: Actions) { 
    this.registerForm = new FormGroup({
        emailAddress: new FormControl('', Validators.compose([Validators.required, Validators.email])),
        password: new FormControl('', Validators.compose([ 
          Validators.required, 
          Validators.minLength(12), 
          Validators.maxLength(64), 
          Validators.pattern("^(.[^a-z]{1,}|[^A-Z]{1,}|[^\\d]{1,}|[^\\W]{1,})$|[\\s]")])
        ),
        confirmPassword: new FormControl('', Validators.compose([ 
          Validators.required, 
          Validators.minLength(12), 
          Validators.maxLength(64), 
          Validators.pattern("^(.[^a-z]{1,}|[^A-Z]{1,}|[^\\d]{1,}|[^\\W]{1,})$|[\\s]")]),
        )
      },
      {
        validators: checkMatchValidator('password', 'confirmPassword')
      }
    );
  }

  public ngOnInit(): void {
    this.loginError$ = this.actions$
      .pipe(ofType(registerRequestFailed))
      .subscribe(() => 
        this.registerForm.setValue({
          password: '',
          confirmPassword: ''
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
    if (this.registerForm.invalid) {
      Object.keys(this.registerForm.controls).forEach(key => {
        this.registerForm.get(key)?.markAsDirty();
      });
      return;
    }
    this.store.dispatch(registerRequestBegin());
  }
}
