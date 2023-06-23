export interface User {
    id: string;
    email: string;
    role: string;
    deleted: boolean | null
}

export const roles = {
    admin: 'admin',
    user: 'user'
}
