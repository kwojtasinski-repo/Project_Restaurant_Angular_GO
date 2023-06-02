import { Credentials } from "src/app/models/credentials";
import { User } from "../../models/user";

export interface LoginState {
  user: User | null;
  authenticated: boolean;
  error: string | null;
  credentials: Credentials,
  path: string,
  loginRequestState: RequestState,
  logoutRequestState: RequestState
}

export enum RequestState {
  init = 'init', loading = 'loading', success = 'success', failed = 'failed'
}
