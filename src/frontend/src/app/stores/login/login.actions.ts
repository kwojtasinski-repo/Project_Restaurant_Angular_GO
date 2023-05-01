import { createAction, props } from "@ngrx/store";
import { User } from "src/app/models/user";

export const loginRequest = createAction('[Login] Login Request');
export const loginRequestSuccess = createAction('[Login] Login Request Success', props<{ user: User }>());
export const loginRequestFailed = createAction('[Login] Login Request Failed', props<{ error: string }>());
