import { validateEmail, validatePassword, validateString } from "../validation";

export const validateInput = (inputId, inputText) => {
  if (inputId === "firstName" || inputId === "lastName") {
    return validateString(inputId, inputText);
  } else if (inputId === "email") {
    return validateEmail(inputId, inputText);
  } else if (inputId === "password") {
    return validatePassword(inputId, inputText);
  }
};
