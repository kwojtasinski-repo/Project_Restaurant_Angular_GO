import { createAction, props } from '@ngrx/store';
import { Credentials } from 'src/app/models/credentials';
import { User } from 'src/app/models/user';

export const reloginRequestSuccess = createAction('[Login] ReLogin Request Success', props<{ user: User }>());
export const reloginRequestFailed = createAction('[Login] ReLogin Request Failed', props<{ error: string }>());
export const loginFormUpdate = createAction('[Login] Login Form Update Request', props<{ credentials: Credentials }>());
export const loginRequest = createAction('[Login] Login Request');
export const loginRequestSuccess = createAction('[Login] Login Request Success', props<{ user: User }>());
export const loginRequestFailed = createAction('[Login] Login Request Failed', props<{ error: string }>());
export const loginSuccess = createAction('[Login] Login Success');
export const logoutRequest = createAction('[Login] Logout Request');
export const logoutRequestSuccess = createAction('[Login] Logout Request Success');
export const logoutRequestFailed = createAction('[Login] Logout Request Failed', props<{ error: string }>());
export const setTargetPath = createAction('[Login] Set Target Path', props<{ path: string }>());
