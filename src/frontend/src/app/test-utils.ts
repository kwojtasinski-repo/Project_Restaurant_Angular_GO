export const changeInputValue = (htmlInputElement: any, value: any) => {
    htmlInputElement.value = value;
    htmlInputElement.dispatchEvent(new Event('input'));
};
