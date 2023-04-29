export function getValidationMessage(code :string): string | null {
    return mapCodeToMessage(code)
}

function mapCodeToMessage(code: string): string | null {
    if (code === 'required') {
        return 'Pole jest wymagane';
    } else if (code === 'email') {
        return 'Niepoprawny adres email';
    }
    return null;
}