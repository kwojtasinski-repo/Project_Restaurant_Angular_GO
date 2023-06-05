import { createAction, props } from "@ngrx/store";
import { Category } from "src/app/models/category";

export const categoryFormClear = createAction('[Category] Category Form Clear');
export const categoryFormUpdate = createAction('[Category] Category Form Update', props<{ category: Category }>());
export const categoryAddRequestBegin = createAction('[Category] Category Add Request Begin');
export const categoryAddRequestSuccess = createAction('[Category] Category Add Request Success');
export const categoryAddRequestFailed = createAction('[Category] Category Add Request Failed', props<{ error: string }>());
export const categoryUpdateRequestBegin = createAction('[Category] Category Update Request Begin');
export const categoryUpdateRequestSuccess = createAction('[Category] Category Update Request Success');
export const categoryUpdateRequestFailed = createAction('[Category] Category Update Request Failed', props<{ error: string }>());
export const categoryCancelOperation = createAction('[Category] Category Cancel operation');
