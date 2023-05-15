import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { Store } from "@ngrx/store";
import { loginRequest } from 'src/app/stores/login/login.actions';
import { getError } from 'src/app/stores/login/login.selectors';
import { LoginState } from 'src/app/stores/login/login.state';
import { getValidationMessage } from 'src/app/validations/validations';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {
  public loginForm: FormGroup;
  public error$ = this.store.select(getError);

  constructor(private store: Store<LoginState>) { 
    this.loginForm = new FormGroup({
      emailAddress: new FormControl('', Validators.compose([Validators.required, Validators.email])),
      password: new FormControl('', Validators.required)
    });
  }

  public ngOnInit(): void {
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
    this.store.dispatch(loginRequest())
  }
}
