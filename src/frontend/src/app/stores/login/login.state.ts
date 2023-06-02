import { Credentials } from "src/app/models/credentials";
import { User } from "../../models/user";
import { RequestState } from "src/app/models/request-state";

export interface LoginState {
  user: User | null;
  authenticated: boolean;
  error: string | null;
  credentials: Credentials,
  path: string,
  loginRequestState: RequestState,
  logoutRequestState: RequestState
}
