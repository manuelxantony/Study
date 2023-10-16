import { NavigationContainer } from "@react-navigation/native";

import AppStack from "./AppStack";
import AuthStack from "./AuthStack";
import { useSelector } from "react-redux";
import StartUpScreen from "../Screens/StartUpScreen";
import { SafeAreaProvider } from "react-native-safe-area-context";

export default function Routes() {
  const isAuth = useSelector(
    (state) => state.auth.token !== null && state.auth.token !== ""
  );
  const didTryAutoLogin = useSelector((state) => state.auth.didTryAutoLogin);
  return (
    <SafeAreaProvider>
      <NavigationContainer>
        {isAuth ? (
          <AppStack />
        ) : didTryAutoLogin ? (
          <AuthStack />
        ) : (
          <StartUpScreen />
        )}
      </NavigationContainer>
    </SafeAreaProvider>
  );
}
