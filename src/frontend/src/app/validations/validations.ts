export function getValidationMessage(code: any): string | null {
    return mapCodeToMessage(code)
}

function mapCodeToMessage(code: any): string | null {
    if (code.key === 'required') {
        return 'Pole jest wymagane';
    } else if (code.key === 'email') {
        return 'Niepoprawny adres email';
    } else if (code.key === 'minlength') {
        if (code.value.requiredLength > 1) {
            return `Pole powinno zawierać ${code.value.requiredLength} znaki`
        } else {
            return `Pole powinno zawierać ${code.value.requiredLength} znak`
        }
    }
    return null;
}