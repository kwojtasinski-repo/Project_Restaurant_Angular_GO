export interface User {
    id: number;
    email: string;
    role: string;
    deleted: boolean | null
}

export const roles = {
    admin: 'admin',
    user: 'user'
}
