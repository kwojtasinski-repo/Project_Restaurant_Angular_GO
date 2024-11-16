import { RequestState } from 'src/app/models/request-state';

export interface RegisterState {
  email: string;
  password: string;
  passwordConfirm: string;
  error: string | null;
  registerRequestState: RequestState
}
