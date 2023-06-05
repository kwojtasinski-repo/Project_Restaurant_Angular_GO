export function getValidationMessage(code: any): string | null {
    return mapCodeToMessage(code)
}

export const PATTERN_ONE_UPPER_ONE_LOWER_ONE_SPECIAL_CHARACTER = "^(?=.*\\d)(?=.*[a-z])(?=.*[A-Z])((?=.*\\W)|(?=.*_))^[^ ]+$";

function mapCodeToMessage(code: any): string | null {
    if (code.key === 'required') {
        return 'Pole jest wymagane';
    } else if (code.key === 'email') {
        return 'Niepoprawny adres email';
    } else if (code.key === 'minlength') {
        if (code.value.requiredLength > 1) {
            return `Pole powinno zawierać ${code.value.requiredLength} znaki`;
        } else {
            return `Pole powinno zawierać ${code.value.requiredLength} znak`;
        }
    } else if (code.key === 'pattern') {
        if (code.value.requiredPattern === PATTERN_ONE_UPPER_ONE_LOWER_ONE_SPECIAL_CHARACTER) {
            return 'Pole powinno zawierać małą i dużą literę oraz specjalny znak znak';
        } else {
            return 'Niepoprawny format';
        }
    }
    return null;
}

export function checkMatchValidator(field1: { fieldName: string; labelName: string }, field2: { fieldName: string; labelName: string }) {
    return function (form: any) {
      let field1Value = form.get(field1.fieldName).value;
      let field2Value = form.get(field2.fieldName).value;

      if (field1Value !== '' && field1Value !== field2Value) {
        return { 'match': `Pola '${field1.labelName}' i '${field2.labelName}' nie są identyczne` }
      }
      return null;
    }
  }