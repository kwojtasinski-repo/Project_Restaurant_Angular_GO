import { createAction, props } from "@ngrx/store";
import { Product } from "src/app/models/product";

export const productFormClear = createAction('[Product] Product From Clear');
export const productFormUpdate = createAction('[Product] Product From Update', props<{ product: Product }>());
export const productAddRequestBegin = createAction('[Product] Product Add Request Begin');
export const productAddRequestSuccess = createAction('[Product] Product Add Request Success');
export const productAddRequestFailed = createAction('[Product] Product Add Request Failed', props<{ error: string }>());
export const productUpdateRequestBegin = createAction('[Product] Product Update Request Begin');
export const productUpdateRequestSuccess = createAction('[Product] Product Update Request Success');
export const productUpdateRequestFailed = createAction('[Product] Product Update Request Failed', props<{ error: string }>());
export const productCancelOperation = createAction('[Product] Product Cancel operation');
