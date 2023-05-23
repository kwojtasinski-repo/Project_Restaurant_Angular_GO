import { Order } from "src/app/models/order";

export interface OrderState {
  order: Order | undefined;
  fetchState: FetchState;
  error: string | null;
}

export enum FetchState {
  init = 'init', loading = 'loading', success = 'success', failed = 'failed'
}
