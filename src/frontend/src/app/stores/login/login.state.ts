import { User } from "../../models/user";

export interface LoginState {
  user: User | null;
  authenticated: boolean;
  error: any;
}
