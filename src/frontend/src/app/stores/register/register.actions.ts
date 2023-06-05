import { createAction, props } from "@ngrx/store";

export const registerFormUpdate = createAction('[Register] Register Form Update Request', props<{ 
    form: {
        email: string;
        password: string;
        confirmPassword: string;
    }
}>());
export const registerRequestBegin = createAction('[Register] Register Request Begin');
export const registerRequestSuccess = createAction('[Register] Register Request Begin');
export const registerRequestFailed = createAction('[Register] Register Request Begin', props<{ error: any }>());
