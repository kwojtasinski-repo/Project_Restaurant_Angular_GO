export function resolveApiUrl() {
    return (window as any )['__env']['backendUrl'] ?? 'http://localhost:8000';
}
