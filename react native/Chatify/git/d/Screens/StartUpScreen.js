import { View, ActivityIndicator } from "react-native";
import commonStyle from "../utils/styles/commonStyle";
import { COLORS } from "../constants";
import { useDispatch } from "react-redux";
import { useState } from "react";
import AsyncStorage from "@react-native-async-storage/async-storage";
import { authenticate, setDidTryAutoLogin } from "../store/authSlice";
import { getUserData } from "../utils/actions/userAction";

export default StartUpScreen = () => {
  const dispatch = useDispatch();
  useState(() => {
    const tryLogin = async () => {
      const storedAuthInfo = await AsyncStorage.getItem("userData");
      if (!storedAuthInfo) {
        dispatch(setDidTryAutoLogin());
        return;
      }
      const parsedData = JSON.parse(storedAuthInfo);
      const { userId, token, expiryDate: expiryDateString } = parsedData;

      const expiryDate = new Date(expiryDateString);

      if (expiryDate <= new Date() || !token || !userId) {
        dispatch(setDidTryAutoLogin());
        return;
      }
      const userData = await getUserData(userId);
      dispatch(authenticate({ token: token, userData }));
    };
    tryLogin();
  }, [dispatch]);

  return (
    <View style={commonStyle.center}>
      <ActivityIndicator size={"large"} color={COLORS.mainOrange} />
    </View>
  );
};
