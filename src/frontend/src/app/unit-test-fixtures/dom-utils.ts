export const changeInputValue = (htmlInputElement: any, value: any) => {
    htmlInputElement.value = value;
    htmlInputElement.dispatchEvent(new Event('input'));
};

export const changeSelectIndex = (htmlInputElement: any, selectedIndex: any) => {
    htmlInputElement.selectedIndex = selectedIndex;
    htmlInputElement.dispatchEvent(new Event('change'));
};
