import { User } from '../models/user';

export function createUser(id?: string, email?: string, role?: string, deleted?: boolean): User {
  return {
    id: id ?? '1',
    email: email ?? 'email@email',
    role: role ?? 'test',
    deleted: deleted ?? null
  } as User;
}
