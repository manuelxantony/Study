export const reducer = (state, action) => {
  const { validationResult, inputId, inputValue } = action;

  const updateValues = {
    ...state.inputValues,
    [inputId]: inputValue,
  };

  const updateInputValidaties = {
    ...state.inputValidaties,
    [inputId]: validationResult,
  };

  let updatedFormIsValid = true;

  for (const key in updateInputValidaties) {
    if (updateInputValidaties[key] !== undefined) {
      // undefined from validation is true
      updatedFormIsValid = false;
      break;
    }
  }
  return {
    inputValues: updateValues,
    inputValidaties: updateInputValidaties,
    formIsValid: updatedFormIsValid,
  };
};
