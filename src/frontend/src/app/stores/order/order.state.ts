import { Order } from "src/app/models/order";
import { RequestState } from "src/app/models/request-state";

export interface OrderState {
  order: Order | undefined;
  fetchState: RequestState;
  error: string | null;
}
