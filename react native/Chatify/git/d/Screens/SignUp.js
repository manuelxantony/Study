import { useCallback, useReducer, useState } from "react";
import {
  View,
  Text,
  StyleSheet,
  Pressable,
  Keyboard,
  KeyboardAvoidingView,
  ActivityIndicator,
} from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import { useDispatch } from "react-redux";

import SubmitButton from "../Components/SubmitButton";
import ValueInput from "../Components/ValueInput";
import { COLORS } from "../constants";
import { signUp } from "../utils/actions/authActions";
import { validateInput } from "../utils/actions/inputValidation";
import { reducer } from "../utils/reducers/inputFormReducers";
import { windowWidth } from "../utils/screenDimensions";

const intialState = {
  inputValues: {
    firstName: false,
    lastName: false,
    email: false,
    password: false,
  },
  inputValidaties: {
    firstName: false,
    lastName: false,
    email: false,
    password: false,
  },
  formIsValid: false,
};

const SignUp = ({ navigation }) => {
  const dispatch = useDispatch();

  const [formState, dispatchFormState] = useReducer(reducer, intialState);
  const [error, setError] = useState("");
  const [isloading, setIsLoading] = useState(false);

  const signInhandler = useCallback(async () => {
    try {
      setIsLoading(true);
      const action = signUp(
        formState.inputValues["firstName"],
        formState.inputValues["lastName"],
        formState.inputValues["email"],
        formState.inputValues["password"]
      );
      setError(null);
      await dispatch(action); // bit confusing ?!!!
    } catch (error) {
      setError(error.message);
      setIsLoading(false);
    }
  }, [dispatch, formState]);

  const onChangeTextHandler = useCallback(
    (inputId, inputValue) => {
      const result = validateInput(inputId, inputValue);
      dispatchFormState({ inputId, validationResult: result, inputValue });
    },
    [formState]
  );

  return (
    <Pressable style={styles.container} onPress={Keyboard.dismiss}>
      <SafeAreaView />
      <KeyboardAvoidingView
        style={styles.container}
        behavior={Platform.OS === "ios" ? "padding" : "height"}
        //keyboardVerticalOffset={Platform.OS === "ios" ? 10 : 0}
      >
        <View style={styles.innerContainer}>
          <Text style={styles.banner}>Chatify</Text>

          <ValueInput
            id="firstName"
            iconName="person-outline"
            placeholder="Enter Your First Name"
            onChangeInputText={onChangeTextHandler}
          />
          <ValueInput
            id="lastName"
            iconName="person-outline"
            placeholder="Enter Your Last Name"
            onChangeInputText={onChangeTextHandler}
          />

          <ValueInput
            id="email"
            iconName="mail-outline"
            placeholder="Enter Your Email"
            keyboardType="email-address"
            onChangeInputText={onChangeTextHandler}
          />

          <ValueInput
            id="password"
            iconName="key-outline"
            placeholder="Enter Your Password"
            secureTextEntry={true}
            onChangeInputText={onChangeTextHandler}
          />

          {isloading ? (
            <ActivityIndicator
              size="small"
              color={COLORS.mainOrange}
              style={styles.indicator}
            />
          ) : (
            <SubmitButton
              submitLabel="Sign Up"
              onPress={signInhandler}
              visible={formState.formIsValid}
            />
          )}
          <Text style={styles.error}>{error}</Text>
          <Pressable
            style={styles.forgotTextContainer}
            onPress={() => {
              navigation.navigate("sign-in");
            }}
          >
            <Text style={styles.forgotText}>
              Having an account? Sign In here
            </Text>
          </Pressable>
        </View>
      </KeyboardAvoidingView>
    </Pressable>
  );
};

export default SignUp;

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: COLORS.white,
  },
  innerContainer: {
    flex: 1,
    backgroundColor: COLORS.white,
    justifyContent: "center",
    alignItems: "center",
  },
  banner: {
    fontSize: 70,
    fontWeight: "800",
    color: COLORS.mainOrange,
    fontFamily: "Raleway-Regular",
  },
  email: {
    flex: 1,
  },
  image: {
    height: 250,
    width: windowWidth,
  },
  forgotTextContainer: {
    marginTop: 20,
  },
  forgotText: {
    color: "#347ec7",
  },
  policy: {
    marginVertical: 20,
    marginHorizontal: 25,
    justifyContent: "center",
    alignItems: "center",
  },
  policyText: {
    textAlign: "center",
    color: "#656464a5",
  },
  conditonText: {
    color: "#e88442",
  },
  error: {
    marginTop: 20,
    color: "red",
  },
  indicator: {
    marginTop: 20,
  },
});
