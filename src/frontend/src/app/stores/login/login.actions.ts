import { createAction, props } from "@ngrx/store";
import { Credentials } from "src/app/models/credentials";
import { User } from "src/app/models/user";

export const initializeLogin = createAction('[Login] Initialize Login', props<{ path: string }>());
export const loginFormUpdate = createAction('[Login] Login From Update Request', props<{ credentials: Credentials }>());
export const loginRequest = createAction('[Login] Login Request');
export const loginRequestSuccess = createAction('[Login] Login Request Success', props<{ user: User }>());
export const loginRequestFailed = createAction('[Login] Login Request Failed', props<{ error: string }>());
export const loginSuccess = createAction('[Login] Login Success');
