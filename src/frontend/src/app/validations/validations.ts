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

export function checkMatchValidator(field1: string, field2: string) {
    return function (form: any) {
      let field1Value = form.get(field1).value;
      let field2Value = form.get(field2).value;
  
      if (field1Value !== '' && field1Value !== field2Value) {
        return { 'match': `pola '${field1Value}' i '${field2Value}' nie są takie same` }
      }
      return null;
    }
  }