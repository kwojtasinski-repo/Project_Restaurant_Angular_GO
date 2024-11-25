export function getValidationMessage(code: any): string | null {
    return mapCodeToMessage(code)
}

export const PATTERN_ONE_UPPER_ONE_LOWER_ONE_SPECIAL_CHARACTER = /^(?=.*\p{Ll})(?=.*\p{Lu})(?=.*\d)(?=.*[\p{P}\p{S}]).+$/u;

function mapCodeToMessage(code: any): string | null {
    if (!code) {
        return null;
    }

    if (code.key === 'required') {
        return 'Pole jest wymagane';
    } else if (code.key === 'email') {
        return 'Niepoprawny adres email';
    } else if (code.key === 'minlength') {
        if (code.value.requiredLength > 1) {
            return `Pole powinno składać się przynajmniej z ${code.value.requiredLength} znaków`;
        } else {
            return `Pole powinno zawierać ${code.value.requiredLength} znak`;
        }
    } else if (code.key === 'maxlength') {
        if (code.value.requiredLength > 1) {
            return `Pole nie powinno przekroczyć ${code.value.requiredLength} znaków`;
        } else {
            return `Pole nie powinno przekroczyć ${code.value.requiredLength} znak`;
        }
    } else if (code.key === 'pattern') {
        if (code.value.requiredPattern === PATTERN_ONE_UPPER_ONE_LOWER_ONE_SPECIAL_CHARACTER.toString()) {
            return 'Pole powinno zawierać małą i dużą literę, jedną liczbę oraz znak specjalny';
        } else {
            return 'Niepoprawny format';
        }
    } else if (code.key === 'min') {
        return `Wartość '${code.value.actual}' powinna być większa niż '${code.value.min}'`;
    }
    return null;
}

export function checkMatchValidator(field1: { fieldName: string; labelName: string }, field2: { fieldName: string; labelName: string }) {
    return function (form: any) {
      const field1Value = form.get(field1.fieldName).value;
      const field2Value = form.get(field2.fieldName).value;

      if (field1Value !== '' && field1Value !== field2Value) {
        return { 'match': `Pola '${field1.labelName}' i '${field2.labelName}' nie są identyczne` }
      }
      return null;
    }
  }