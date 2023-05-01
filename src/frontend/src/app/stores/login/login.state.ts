import { Credentials } from "src/app/models/credentials";
import { User } from "../../models/user";

export interface LoginState {
  user: User | null;
  authenticated: boolean;
  error: any;
  credentials: Credentials,
  path: string
}
