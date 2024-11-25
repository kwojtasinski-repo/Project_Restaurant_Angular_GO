import { InjectionToken } from '@angular/core';

export const API_URL = new InjectionToken<string>('API_URL');

export function resolveApiUrl() {
    return (window as any )['__env']['backendUrl'] ?? 'http://localhost:8000';
}
