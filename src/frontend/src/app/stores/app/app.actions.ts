import { createAction, props } from '@ngrx/store';

export const enableHeader = createAction('[App] Enable Header');
export const disableHeader = createAction('[App] Disable Header');
export const setCurrentUrl = createAction('[App] Set Current Url', props<{ currentUrl: string }>())
