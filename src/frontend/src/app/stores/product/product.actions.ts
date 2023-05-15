import { createAction, props } from "@ngrx/store";
import { Product } from "src/app/models/product";

export const productFormClear = createAction('[Product] Product From Clear');
export const productFormUpdate = createAction('[Product] Product From Update', props<{ product: Product }>());
export const productRequestBegin = createAction('[Product] Product Request Begin');
export const productRequestSuccess = createAction('[Product] Product Request Success');
export const productRequestFailed = createAction('[Product] Product Request Failed', props<{ error: string }>());
