import { useState, useReducer, useCallback } from "react";
import {
  View,
  Text,
  StyleSheet,
  Image,
  Pressable,
  Keyboard,
  KeyboardAvoidingView,
  ActivityIndicator,
  Platform,
} from "react-native";

import SubmitButton from "../Components/SubmitButton";
import ValueInput from "../Components/ValueInput";
import { COLORS } from "../constants";
import { windowWidth } from "../utils/screenDimensions";
import { validateInput } from "../utils/actions/inputValidation";
import { reducer } from "../utils/reducers/inputFormReducers";
import { signIn } from "../utils/actions/authActions";
import { useDispatch } from "react-redux";
import { isTestMode } from "../chatify-config";
import { SafeAreaView } from "react-native-safe-area-context";

const intialState = {
  inputValues: {
    email: isTestMode
      ? Platform.OS == "ios"
        ? "lufy@monkey.com"
        : "levy@akraman.com"
      : "",
    password: isTestMode ? (Platform.OS == "ios" ? "imlufy" : "levylevy") : "",
  },
  inputValidaties: {
    email: isTestMode,
    password: isTestMode,
  },
  formIsValid: isTestMode,
};

const SignIn = ({ navigation }) => {
  const dispatch = useDispatch();

  const [formState, dispatchFormState] = useReducer(reducer, intialState);
  const [error, setError] = useState("");
  const [isloading, setIsLoading] = useState(false);

  const signInHandler = useCallback(async () => {
    try {
      setIsLoading(true);
      const action = signIn(
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
      <KeyboardAvoidingView
        style={styles.innerContainer}
        behavior={Platform.OS === "ios" ? "padding" : "height"}
        keyboardVerticalOffset={Platform.OS === "ios" ? 90 : 0}
      >
        <View style={styles.signInContainer}>
          <Text style={styles.banner}>Chatify</Text>
          <View>
            <Image
              source={require("../assets/icon.jpg")}
              style={styles.image}
            />
          </View>
          <ValueInput
            id="email"
            iconName="mail-outline"
            placeholder="Enter Your Email"
            keyboardType="email-address"
            onChangeInputText={onChangeTextHandler}
            value={formState.inputValues.email}
          />

          <ValueInput
            id="password"
            iconName="key-outline"
            placeholder="Enter Your Password"
            secureTextEntry={true}
            onChangeInputText={onChangeTextHandler}
            value={formState.inputValues.password}
          />
          {isloading ? (
            <ActivityIndicator
              size="small"
              color={COLORS.mainOrange}
              style={styles.indicator}
            />
          ) : (
            <SubmitButton
              submitLabel="Sign In"
              onPress={signInHandler}
              visible={formState.formIsValid}
            />
          )}
          <Text style={styles.error}>{error}</Text>
          <Pressable
            style={styles.signUpContainer}
            onPress={() => {
              navigation.navigate("sign-up");
            }}
          >
            <Text style={styles.signUpContainerText}>
              Don't have an account? Create here
            </Text>
          </Pressable>
        </View>
      </KeyboardAvoidingView>
    </Pressable>
  );
};

export default SignIn;

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: COLORS.white,
  },
  innerContainer: {
    flex: 1,
    marginTop: 30,
  },
  signInContainer: {
    flex: 1,
    justifyContent: "center",
    alignItems: "center",
    backgroundColor: COLORS.white,
    marginBottom: 30,
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
  signUpContainer: {
    marginTop: 20,
  },
  signUpContainerText: {
    color: "#347ec7",
  },
  error: {
    marginTop: 20,
    color: "red",
  },
  indicator: {
    marginTop: 20,
  },
});
